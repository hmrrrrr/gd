// Package WorkerThreadPool provides methods for working with WorkerThreadPool object instances.
package WorkerThreadPool

import "unsafe"
import "sync"
import "reflect"
import "slices"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant"
import "graphics.gd/variant/Array"
import "graphics.gd/variant/Callable"
import "graphics.gd/variant/Dictionary"
import "graphics.gd/variant/Error"
import "graphics.gd/variant/Float"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/Packed"
import "graphics.gd/variant/Path"
import "graphics.gd/variant/RID"
import "graphics.gd/variant/RefCounted"
import "graphics.gd/variant/String"

var _ Object.ID
var _ RefCounted.Instance
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle
var _ = Array.Nil
var _ variant.Any
var _ Callable.Function
var _ Dictionary.Any
var _ RID.Any
var _ String.Readable
var _ Path.ToNode
var _ Packed.Bytes
var _ Error.Code
var _ Float.X
var _ = slices.Delete[[]struct{}, struct{}]

/*
The [WorkerThreadPool] singleton allocates a set of [Thread]s (called worker threads) on project startup and provides methods for offloading tasks to them. This can be used for simple multithreading without having to create [Thread]s.
Tasks hold the [Callable] to be run by the threads. [WorkerThreadPool] can be used to create regular tasks, which will be taken by one worker thread, or group tasks, which can be distributed between multiple worker threads. Group tasks execute the [Callable] multiple times, which makes them useful for iterating over a lot of elements, such as the enemies in an arena.
Here's a sample on how to offload an expensive function to worker threads:
[codeblocks]
[gdscript]
var enemies = [] # An array to be filled with enemies.

func process_enemy_ai(enemy_index):

	var processed_enemy = enemies[enemy_index]
	# Expensive logic...

func _process(delta):

	var task_id = WorkerThreadPool.add_group_task(process_enemy_ai, enemies.size())
	# Other code...
	WorkerThreadPool.wait_for_group_task_completion(task_id)
	# Other code that depends on the enemy AI already being processed.

[/gdscript]
[csharp]
private List<Node> _enemies = new List<Node>(); // A list to be filled with enemies.

private void ProcessEnemyAI(int enemyIndex)

	{
	    Node processedEnemy = _enemies[enemyIndex];
	    // Expensive logic here.
	}

public override void _Process(double delta)

	{
	    long taskId = WorkerThreadPool.AddGroupTask(Callable.From<int>(ProcessEnemyAI), _enemies.Count);
	    // Other code...
	    WorkerThreadPool.WaitForGroupTaskCompletion(taskId);
	    // Other code that depends on the enemy AI already being processed.
	}

[/csharp]
[/codeblocks]
The above code relies on the number of elements in the [code]enemies[/code] array remaining constant during the multithreaded part.
[b]Note:[/b] Using this singleton could affect performance negatively if the task being distributed between threads is not computationally expensive.
*/
var self [1]gdclass.WorkerThreadPool
var once sync.Once

func singleton() {
	obj := gd.Global.Object.GetSingleton(gd.Global.Singletons.WorkerThreadPool)
	self = *(*[1]gdclass.WorkerThreadPool)(unsafe.Pointer(&obj))
}

/*
Adds [param action] as a task to be executed by a worker thread. [param high_priority] determines if the task has a high priority or a low priority (default). You can optionally provide a [param description] to help with debugging.
Returns a task ID that can be used by other methods.
[b]Warning:[/b] Every task must be waited for completion using [method wait_for_task_completion] or [method wait_for_group_task_completion] at some point so that any allocated resources inside the task can be cleaned up.
*/
func AddTask(action func()) int { //gd:WorkerThreadPool.add_task
	once.Do(singleton)
	return int(int(class(self).AddTask(Callable.New(action), false, String.New(""))))
}

/*
Returns [code]true[/code] if the task with the given ID is completed.
[b]Note:[/b] You should only call this method between adding the task and awaiting its completion.
*/
func IsTaskCompleted(task_id int) bool { //gd:WorkerThreadPool.is_task_completed
	once.Do(singleton)
	return bool(class(self).IsTaskCompleted(int64(task_id)))
}

/*
Pauses the thread that calls this method until the task with the given ID is completed.
Returns [constant @GlobalScope.OK] if the task could be successfully awaited.
Returns [constant @GlobalScope.ERR_INVALID_PARAMETER] if a task with the passed ID does not exist (maybe because it was already awaited and disposed of).
Returns [constant @GlobalScope.ERR_BUSY] if the call is made from another running task and, due to task scheduling, there's potential for deadlocking (e.g., the task to await may be at a lower level in the call stack and therefore can't progress). This is an advanced situation that should only matter when some tasks depend on others (in the current implementation, the tricky case is a task trying to wait on an older one).
*/
func WaitForTaskCompletion(task_id int) error { //gd:WorkerThreadPool.wait_for_task_completion
	once.Do(singleton)
	return error(gd.ToError(class(self).WaitForTaskCompletion(int64(task_id))))
}

/*
Adds [param action] as a group task to be executed by the worker threads. The [Callable] will be called a number of times based on [param elements], with the first thread calling it with the value [code]0[/code] as a parameter, and each consecutive execution incrementing this value by 1 until it reaches [code]element - 1[/code].
The number of threads the task is distributed to is defined by [param tasks_needed], where the default value [code]-1[/code] means it is distributed to all worker threads. [param high_priority] determines if the task has a high priority or a low priority (default). You can optionally provide a [param description] to help with debugging.
Returns a group task ID that can be used by other methods.
[b]Warning:[/b] Every task must be waited for completion using [method wait_for_task_completion] or [method wait_for_group_task_completion] at some point so that any allocated resources inside the task can be cleaned up.
*/
func AddGroupTask(action func(), elements int) int { //gd:WorkerThreadPool.add_group_task
	once.Do(singleton)
	return int(int(class(self).AddGroupTask(Callable.New(action), int64(elements), int64(-1), false, String.New(""))))
}

/*
Returns [code]true[/code] if the group task with the given ID is completed.
[b]Note:[/b] You should only call this method between adding the group task and awaiting its completion.
*/
func IsGroupTaskCompleted(group_id int) bool { //gd:WorkerThreadPool.is_group_task_completed
	once.Do(singleton)
	return bool(class(self).IsGroupTaskCompleted(int64(group_id)))
}

/*
Returns how many times the [Callable] of the group task with the given ID has already been executed by the worker threads.
[b]Note:[/b] If a thread has started executing the [Callable] but is yet to finish, it won't be counted.
*/
func GetGroupProcessedElementCount(group_id int) int { //gd:WorkerThreadPool.get_group_processed_element_count
	once.Do(singleton)
	return int(int(class(self).GetGroupProcessedElementCount(int64(group_id))))
}

/*
Pauses the thread that calls this method until the group task with the given ID is completed.
*/
func WaitForGroupTaskCompletion(group_id int) { //gd:WorkerThreadPool.wait_for_group_task_completion
	once.Do(singleton)
	class(self).WaitForGroupTaskCompletion(int64(group_id))
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
func Advanced() class { once.Do(singleton); return self }

type class [1]gdclass.WorkerThreadPool

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }

/*
Adds [param action] as a task to be executed by a worker thread. [param high_priority] determines if the task has a high priority or a low priority (default). You can optionally provide a [param description] to help with debugging.
Returns a task ID that can be used by other methods.
[b]Warning:[/b] Every task must be waited for completion using [method wait_for_task_completion] or [method wait_for_group_task_completion] at some point so that any allocated resources inside the task can be cleaned up.
*/
//go:nosplit
func (self class) AddTask(action Callable.Function, high_priority bool, description String.Readable) int64 { //gd:WorkerThreadPool.add_task
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalCallable(action)))
	callframe.Arg(frame, high_priority)
	callframe.Arg(frame, pointers.Get(gd.InternalString(description)))
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.WorkerThreadPool.Bind_add_task, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns [code]true[/code] if the task with the given ID is completed.
[b]Note:[/b] You should only call this method between adding the task and awaiting its completion.
*/
//go:nosplit
func (self class) IsTaskCompleted(task_id int64) bool { //gd:WorkerThreadPool.is_task_completed
	var frame = callframe.New()
	callframe.Arg(frame, task_id)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.WorkerThreadPool.Bind_is_task_completed, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Pauses the thread that calls this method until the task with the given ID is completed.
Returns [constant @GlobalScope.OK] if the task could be successfully awaited.
Returns [constant @GlobalScope.ERR_INVALID_PARAMETER] if a task with the passed ID does not exist (maybe because it was already awaited and disposed of).
Returns [constant @GlobalScope.ERR_BUSY] if the call is made from another running task and, due to task scheduling, there's potential for deadlocking (e.g., the task to await may be at a lower level in the call stack and therefore can't progress). This is an advanced situation that should only matter when some tasks depend on others (in the current implementation, the tricky case is a task trying to wait on an older one).
*/
//go:nosplit
func (self class) WaitForTaskCompletion(task_id int64) Error.Code { //gd:WorkerThreadPool.wait_for_task_completion
	var frame = callframe.New()
	callframe.Arg(frame, task_id)
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.WorkerThreadPool.Bind_wait_for_task_completion, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Error.Code(r_ret.Get())
	frame.Free()
	return ret
}

/*
Adds [param action] as a group task to be executed by the worker threads. The [Callable] will be called a number of times based on [param elements], with the first thread calling it with the value [code]0[/code] as a parameter, and each consecutive execution incrementing this value by 1 until it reaches [code]element - 1[/code].
The number of threads the task is distributed to is defined by [param tasks_needed], where the default value [code]-1[/code] means it is distributed to all worker threads. [param high_priority] determines if the task has a high priority or a low priority (default). You can optionally provide a [param description] to help with debugging.
Returns a group task ID that can be used by other methods.
[b]Warning:[/b] Every task must be waited for completion using [method wait_for_task_completion] or [method wait_for_group_task_completion] at some point so that any allocated resources inside the task can be cleaned up.
*/
//go:nosplit
func (self class) AddGroupTask(action Callable.Function, elements int64, tasks_needed int64, high_priority bool, description String.Readable) int64 { //gd:WorkerThreadPool.add_group_task
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalCallable(action)))
	callframe.Arg(frame, elements)
	callframe.Arg(frame, tasks_needed)
	callframe.Arg(frame, high_priority)
	callframe.Arg(frame, pointers.Get(gd.InternalString(description)))
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.WorkerThreadPool.Bind_add_group_task, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns [code]true[/code] if the group task with the given ID is completed.
[b]Note:[/b] You should only call this method between adding the group task and awaiting its completion.
*/
//go:nosplit
func (self class) IsGroupTaskCompleted(group_id int64) bool { //gd:WorkerThreadPool.is_group_task_completed
	var frame = callframe.New()
	callframe.Arg(frame, group_id)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.WorkerThreadPool.Bind_is_group_task_completed, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns how many times the [Callable] of the group task with the given ID has already been executed by the worker threads.
[b]Note:[/b] If a thread has started executing the [Callable] but is yet to finish, it won't be counted.
*/
//go:nosplit
func (self class) GetGroupProcessedElementCount(group_id int64) int64 { //gd:WorkerThreadPool.get_group_processed_element_count
	var frame = callframe.New()
	callframe.Arg(frame, group_id)
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.WorkerThreadPool.Bind_get_group_processed_element_count, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Pauses the thread that calls this method until the group task with the given ID is completed.
*/
//go:nosplit
func (self class) WaitForGroupTaskCompletion(group_id int64) { //gd:WorkerThreadPool.wait_for_group_task_completion
	var frame = callframe.New()
	callframe.Arg(frame, group_id)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.WorkerThreadPool.Bind_wait_for_group_task_completion, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}
func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(Object.Advanced(self.AsObject()), name)
	}
}
func init() {
	gdclass.Register("WorkerThreadPool", func(ptr gd.Object) any {
		return [1]gdclass.WorkerThreadPool{*(*gdclass.WorkerThreadPool)(unsafe.Pointer(&ptr))}
	})
}
