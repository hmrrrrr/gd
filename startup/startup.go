// Package startup provides a runtime for connecting to the graphics engine.
package startup

import (
	"iter"

	"graphics.gd/classdb"
	EngineClass "graphics.gd/classdb/Engine"
	MainLoopClass "graphics.gd/classdb/MainLoop"
	"graphics.gd/classdb/ProjectSettings"
	gd "graphics.gd/internal"
	"graphics.gd/internal/pointers"
	"graphics.gd/variant/Float"
)

var mainloop MainLoopClass.Interface

var intialized = make(chan struct{})
var shutdown = make(chan struct{})

// MainLoop uses the given struct as the main loop implementation. This will take care of initialising
// the Go runtime correctly, blocks until the main loop has shutdown.
func MainLoop(loop MainLoopClass.Interface) {
	if EngineClass.IsEditorHint() {
		stop_main()
	}
	theMainFunctionIsWaitingForTheEngineToShutDown = true
	classdb.Register[goMainLoop]()
	className := classdb.NameFor[goMainLoop]()
	ProjectSettings.SetInitialValue("application/run/main_loop_type", className)
	ProjectSettings.SetSetting("application/run/main_loop_type", className)
	mainloop = loop
	pause_main(false) // We pause here until the engine has fully started up.
}

var theMainFunctionIsWaitingForTheEngineToShutDown = false

// Engine starts up the engine and blocks until it shuts down.
func Engine() {
	if pause_main != nil {
		theMainFunctionIsWaitingForTheEngineToShutDown = true
		pause_main(false)
	} else {
		<-shutdown
	}
}

// Loader starts up the loading and initialization process of the graphics engine.
// After this function is called, all graphics functions will be available to use.
// A subsequent call to [Engine] is required to continue to the default main loop.
func Loader() {
	<-intialized
	if pause_main != nil {
		gd.NewCallable(func() {
			resume_main()
		}).CallDeferred()
		pause_main(false)
		if EngineClass.IsEditorHint() {
			stop_main()
		}
	}
}

type goMainLoop struct {
	classdb.Extension[goMainLoop, MainLoopClass.Instance] `gd:"GoMainLoop"`
}

// Called once during initialization.
func (loop goMainLoop) Initialize() {
	if mainloop != nil {
		mainloop.Initialize()
	} else {
		resume_main()
	}
}

// Called each physics frame with the time since the last physics frame as argument ([param delta], in seconds). Equivalent to [method Node._physics_process].
// If implemented, the method must return a boolean value. [code]true[/code] ends the main loop, while [code]false[/code] lets it proceed to the next frame.
func (loop goMainLoop) PhysicsProcess(delta Float.X) bool {
	if mainloop != nil {
		return mainloop.PhysicsProcess(delta)
	}
	return false
}

var dt Float.X

// Called each process (idle) frame with the time since the last process frame as argument (in seconds). Equivalent to [method Node._process].
// If implemented, the method must return a boolean value. [code]true[/code] ends the main loop, while [code]false[/code] lets it proceed to the next frame.
func (loop goMainLoop) Process(delta Float.X) bool {
	defer pointers.Cycle()
	if mainloop != nil {
		return mainloop.Process(delta)
	}
	dt = delta
	close, _ := resume_main()
	return close
}

// Called before the program exits.
func (loop goMainLoop) Finalize() {
	if mainloop != nil {
		mainloop.Finalize()
	} else {
		resume_main()
	}
}

// Rendering waits for the engine to startup and returns a frame iterator for the primary viewport that is
// ready for rendering. The iterator will block until the engine shuts down.
//
//		func main() {
//			frames := startup.Rendering()
//	    	// init.
//			for frame := range frames {
//				// render frame
//			}
//			// finalize
//		}
func Rendering() iter.Seq[Float.X] {
	if EngineClass.IsEditorHint() {
		stop_main()
	}
	classdb.Register[goMainLoop]()
	className := classdb.NameFor[goMainLoop]()
	ProjectSettings.SetInitialValue("application/run/main_loop_type", className)
	ProjectSettings.SetSetting("application/run/main_loop_type", className)
	pause_main(false) // We pause here until the engine has fully started up.
	return func(yield func(Float.X) bool) {
		pause_main(false) // we pause here until the MainLoop initialize function is called.
		for {
			pause_main(false) // we pause here until the next frame is ready (next Process callback).
			if !yield(dt) {
				break
			}
		}
		pause_main(true) // we pause here until the engine has fully shut down.
	}
}

var (
	pause_main  func(bool) bool
	resume_main func() (bool, bool)
	stop_main   func()
)
