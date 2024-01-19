# Godot 4.2.1 + Go [![Go Reference](https://pkg.go.dev/badge/grow.graphics/gd.svg)](https://pkg.go.dev/grow.graphics/gd)

This module provides a safe and simple way to work with Godot 4.2.1, in Go.

You can support the project and prioritise issues [here](https://buy.stripe.com/4gw14maETbnX3vOcMM)

```go
package main

import (
	"fmt"

	"grow.graphics/gd"
	"grow.graphics/gd/gdextension"
)

type HelloWorld struct {
	gd.Class[HelloWorld, gd.Node2D]
}

// Ready implements the Godot Node2D _ready interface (virtual function).
func (h *HelloWorld) Ready(gd.Context) {
	fmt.Println("Hello World from Go!")
}

func main() {
	godot, ok := gdextension.Link()
	if !ok {
		return
	}
	gdextension.RegisterClass[HelloWorld](godot)
}

```

## Getting Started
The module includes a drop-in replacement for the go command called `gd` that 
makes it easy to work with projects that run within Godot. It enables you to
start developing a new project from a single main.go file, to install it, make
sure that your `$GOPATH/bin` is in your `$PATH` and run:

```sh
	go install grow.graphics/gd/cmd/gd@master
```

Now when you can run `gd run`, `gd test` in your project's directory and things
will work as expected.

On linux, `gd` will download the correct version of Godot for you automatically.
Running the command without any arguments will startup the editor so you can 
manage the graphical aspects of your project.

## Design Principles

Godot classes are exported by the `gd` package and can be referred to by 
their standard Godot names, for example `gd.Object` is an 
[Object](https://docs.godotengine.org/en/latest/classes/class_object.html) 
reference. There's no inheritance, so to access the 'super' class, you need 
to call `Super()` on your custom 'Class'. All Godot classes have methods
to cast to the classes they extend for example `AsObject()` or `AsNode2D()`.

Methods have been renamed to follow Go conventions, so instead of
underscores, methods are named as PascalCase. Keep this in mind when
referring to the Godot documentation.

https://docs.godotengine.org/en/latest/index.html

## Semi-Automatic Memory Management

Godot types are preferred over Go types, in order to keep allocations optional. 
All values are tied to a [gd.Context] type, which is either:

    (a) a function argument and any values associated with it will be freed
        when the function returns.
    (b) builtin to the class you are extending, and any values associated 
        with it will be freed when the class is destroyed by Godot.

This module aims to offer memory safety for race-free extensions, if you discover
a way to unintentionally do something unsafe (like double free, use-after-free or
a segfault), using methods on types exported by the root `gd` package please open 
an issue. 

## Performance

No profiling has been completed, however all Go -> Godot calls are optimised
in a way to avoid almost all allocations. Allocations are currently unavoidable
for all Godot -> Go script calls (but not for virtual overrides, which are essentially
allocation free).

## Example

Run `make` in the example/src directory or manually build a C-shared library:

```sh
cd example/src
make # go build -o ./bin/example.so -buildmode=c-shared
```

**NOTE:** the first time you build a project, it will take a little while to compile.
After this, subsequent builds will be quite a bit faster!

Now open the example project in Godot from a terminal and you will be able to 
see Go printing things to the console.

## Testing
To run the go tests for this module, you need to compile the test library:
```sh
cd ./internal/testbed
make # go test .. -o ./testing.lib -c -buildmode=c-shared
```
Then launch godot from the `internal/testbed` directory with the desired Go
test-prefixed test flags (`-test.v`, `-test.bench`, etc) be sure to add the
`--headless` flag to prevent Godot from opening a window!

## Known Limitations

* No support for indexed properties
* No support for Godot functions with varargs.
* No support for script extensions.
* No methods for Godot math types, Vectors, Transforms, etc.
* 64bit support only.
* Untested on Windows.

## Licensing
This project is licensed under an MIT license (the same license as Godot), you can use 
it in any manner you can use the Godot engine. If you use this for a commercially successful
project, please consider [financially supporting us](https://buy.stripe.com/4gw14maETbnX3vOcMM).
