package data

import (
	"encoding/json"
	"time"
)

type TaskType string

var (
	MITTask  TaskType = "MIT"
	TodoTask TaskType = "TODO"
)

type TaskSet struct {
	MITTasks  TaskList `json:"mit_tasks"`
	TodoTasks TaskList `json:"todo_tasks"`
}

func NewTaskSet() *TaskSet {
	return &TaskSet{
		MITTasks: TaskList{
			TasksType: MITTask,
			Tasks:     make([]Task, 0),
		},
		TodoTasks: TaskList{
			TasksType: TodoTask,
			Tasks:     make([]Task, 0),
		},
	}
}

func (ts *TaskSet) AddTask(tt TaskType, task string) error {
	if tt == MITTask {
		return ts.MITTasks.AddTask(task)
	}

	return ts.TodoTasks.AddTask(task)
}

func (ts *TaskSet) RemoveTask(i int) error {
	if i >= ts.Len() {
		return ErrOutOfBounds
	}

	mitTaskCount := len(ts.MITTasks.Tasks)

	if i >= mitTaskCount {
		// remove from Todos
		ts.TodoTasks.RemoveTask(i - mitTaskCount)
	}

	// remove from MITs
	ts.MITTasks.RemoveTask(i)

	return nil
}

func (ts *TaskSet) Len() int {
	return len(ts.MITTasks.Tasks) + len(ts.TodoTasks.Tasks)
}

func (ts *TaskSet) ToJSON() ([]byte, error) {
	return json.Marshal(ts)
}

type TaskList struct {
	TasksType TaskType `json:"tasks_type"`
	Tasks     []Task   `json:"tasks"`
}

func (tl *TaskList) IsEmpty() bool {
	return len(tl.Tasks) == 0
}

func (tl *TaskList) AddTask(task string) error {
	if tl.IsAtLimit() {
		return ErrAtTaskLimit
	}

	tl.Tasks = append(tl.Tasks, Task{
		CreatedAt: time.Now(),
		Content:   task,
	})

	return nil
}

func (tl *TaskList) RemoveTask(i int) error {
	if  i < 0 || len(tl.Tasks) <= i {
		return ErrOutOfBounds
	}

	tl.Tasks = append(tl.Tasks[:i], tl.Tasks[i+1:]...)

	return nil
}

func (tl *TaskList) IsAtLimit() bool {
	return len(tl.Tasks) >= tl.Limit()
}

// @todo config control of task limits
func (tl *TaskList) Limit() int {
	if tl.TasksType == MITTask {
		return MITLimit
	}

	return TodoLimit
}

type Task struct {
	CreatedAt time.Time `json:"created_at"`
	Content   string    `json:"content"`
}

func (t *Task) IsOutstanding() bool {
	return time.Now().Sub(t.CreatedAt) >= (time.Hour * 12)
}
