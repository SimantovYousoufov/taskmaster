package data

import (
	test "github.com/simantovyousoufov/taskmaster/testing"
	"testing"
	"time"
)

func TestTask_IsOutstanding(t *testing.T) {
	task := &Task{
		CreatedAt: time.Now(),
		Content: "Foo bar bazz",
	}

	test.AssertFalse(t, task.IsOutstanding())
}

func TestTask_IsOutstanding_IfOutstanding(t *testing.T) {
	task := &Task{
		CreatedAt: time.Now().Add(-(time.Hour * 12 + 1)),
		Content: "Foo bar bazz",
	}

	test.AssertTrue(t, task.IsOutstanding())
}

func TestTaskList_IsEmpty(t *testing.T) {
	l := &TaskList{
		TasksType: MITTask,
		Tasks: make([]Task, 0),
	}

	test.AssertTrue(t, l.IsEmpty())
}

func TestTaskList_AddTask(t *testing.T) {
	l := &TaskList{
		TasksType: MITTask,
		Tasks: make([]Task, 0),
	}

	test.AssertNilError(t, l.AddTask("foo string"))

	test.AssertFalse(t, l.IsEmpty())

	test.AssertSame(t, "foo string", l.Tasks[0].Content)
}

func TestTaskList_AddTask_ReturnsErrorIfBeyondLimit(t *testing.T) {
	l := &TaskList{
		TasksType: MITTask,
		Tasks: make([]Task, 0),
	}

	test.AssertNilError(t, l.AddTask("foo string"))
	test.AssertFalse(t, l.IsAtLimit())

	test.AssertNilError(t, l.AddTask("foo string"))
	test.AssertFalse(t, l.IsAtLimit())

	test.AssertNilError(t, l.AddTask("foo string"))
	test.AssertTrue(t, l.IsAtLimit())

	test.AssertErrorValue(t, l.AddTask("foo string"), ErrAtTaskLimit.Error())
}

func TestTaskList_RemoveTask_ReturnsErrorIfOutOfBounds(t *testing.T) {
	l := &TaskList{
		TasksType: MITTask,
		Tasks: make([]Task, 0),
	}

	test.AssertErrorValue(t, l.RemoveTask(0), ErrOutOfBounds.Error())

	test.AssertNilError(t, l.AddTask("foo string"))

	test.AssertErrorValue(t, l.RemoveTask(1), ErrOutOfBounds.Error())
}

func TestTaskList_RemoveTask(t *testing.T) {
	l := &TaskList{
		TasksType: MITTask,
		Tasks: make([]Task, 0),
	}

	test.AssertNilError(t, l.AddTask("foo string"))

	test.AssertNilError(t, l.RemoveTask(0))
}

func TestTaskList_Limit(t *testing.T) {
	l := &TaskList{
		TasksType: MITTask,
		Tasks: make([]Task, 0),
	}

	test.AssertSame(t, 3, l.Limit())

	l.TaskLimit = 10

	test.AssertSame(t, 10, l.Limit())
}

func TestTaskSet_AddTask(t *testing.T) {
	s := NewTaskSet()

	test.AssertNilError(t, s.AddTask(MITTask, "foo string"))

	test.AssertSame(t, "foo string", s.MITTasks.Tasks[0].Content)

	test.AssertSame(t, 1, s.Len())
}

func TestTaskSet_RemoveTask(t *testing.T) {
	s := NewTaskSet()

	test.AssertNilError(t, s.AddTask(MITTask, "foo string"))

	test.AssertSame(t, 1, s.Len())

	test.AssertNilError(t, s.RemoveTask(0))

	test.AssertSame(t, 0, s.Len())
}

func TestTaskSet_RemoveTask_ReturnsLimitErrors(t *testing.T) {
	s := NewTaskSet()

	for i := 0; i < 3; i++ {
		test.AssertNilError(t, s.AddTask(MITTask, "foo string"))
	}

	test.AssertErrorValue(t, s.AddTask(MITTask, "foo string"), ErrAtTaskLimit.Error())

	for i := 0; i < 10; i++ {
		test.AssertNilError(t, s.AddTask(TodoTask, "foo string"))
	}

	test.AssertErrorValue(t, s.AddTask(TodoTask, "foo string"), ErrAtTaskLimit.Error())
}

func TestTaskSet_RemoveTask_RemovesFromMITUnderMITLimit(t *testing.T) {
	s := NewTaskSet()

	for i := 0; i < 3; i++ {
		test.AssertNilError(t, s.AddTask(MITTask, "foo string"))
	}

	for i := 0; i < 10; i++ {
		test.AssertNilError(t, s.AddTask(TodoTask, "foo string"))
	}

	test.AssertSame(t, 3, len(s.MITTasks.Tasks))
	test.AssertSame(t, 10, len(s.TodoTasks.Tasks))

	s.RemoveTask(2)

	test.AssertSame(t, 2, len(s.MITTasks.Tasks))
	test.AssertSame(t, 10, len(s.TodoTasks.Tasks))
}

func TestTaskSet_RemoveTask_RemovesFromTodosAboveMITLimit(t *testing.T) {
	s := NewTaskSet()

	for i := 0; i < 3; i++ {
		test.AssertNilError(t, s.AddTask(MITTask, "foo string"))
	}

	for i := 0; i < 10; i++ {
		test.AssertNilError(t, s.AddTask(TodoTask, "foo string"))
	}

	test.AssertSame(t, 3, len(s.MITTasks.Tasks))
	test.AssertSame(t, 10, len(s.TodoTasks.Tasks))

	s.RemoveTask(4)

	test.AssertSame(t, 3, len(s.MITTasks.Tasks))
	test.AssertSame(t, 9, len(s.TodoTasks.Tasks))
}

func TestTaskSet_RemoveTask_ReturnsOutOfBoundsError(t *testing.T) {
	s := NewTaskSet()

	test.AssertNilError(t, s.AddTask(MITTask, "foo string"))

	test.AssertSame(t, 1, s.Len())

	test.AssertErrorValue(t, s.RemoveTask(1), ErrOutOfBounds.Error())
}
