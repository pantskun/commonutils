package taskutils

import (
	"github.com/pantskun/commonutils/container"
)

// type Task interface {
// 	GetState() ETaskState
// }

type ETaskState int

const (
	ETaskStateError    ETaskState = 0
	ETaskStateWaiting  ETaskState = 1
	ETaskStateReady    ETaskState = 2
	ETaskStateRunning  ETaskState = 3
	ETaskStateFinished ETaskState = 4
)

type Task struct {
	name        string
	do          func() error
	state       ETaskState
	preTaskList []*Task
	subTaskList []*Task
}

// NewTask 创建新的任务.
//  do func() error: 任务的内容
//  preTaskks: 任务的前置任务
func NewTask(name string, do func() error, preTasks ...*Task) *Task {
	newTask := Task{name: name, do: do, state: ETaskStateWaiting, preTaskList: preTasks}

	for _, preTask := range preTasks {
		if preTask != nil {
			preTask.subTaskList = append(preTask.subTaskList, &newTask)
		}
	}

	if len(preTasks) == 0 {
		newTask.state = ETaskStateReady
	}

	return &newTask
}

func (t *Task) Equal(other container.Element) bool {
	value, ok := other.(*Task)
	if !ok {
		return false
	}

	return t == value
}

// GetState 获取任务状态.
func (t *Task) GetState() ETaskState {
	return t.state
}

// CheckIsReady 检查任务前置是否都已完成.
func (t *Task) CheckIsReady() {
	for _, preTask := range t.preTaskList {
		if preTask == nil {
			continue
		}

		if preTask.state == ETaskStateError {
			t.state = ETaskStateError
			return
		}

		if preTask.state != ETaskStateFinished {
			t.state = ETaskStateWaiting
			return
		}
	}

	t.state = ETaskStateReady
}

func (t *Task) Run() {
	if t.state != ETaskStateReady {
		return
	}

	t.state = ETaskStateRunning
	// log.Println("Task:", t.name, " Running")

	err := t.do()
	if err != nil {
		t.state = ETaskStateError
		return
	}

	t.state = ETaskStateFinished
	// log.Println("Task:", t.name, " Finished")

	for _, subTask := range t.subTaskList {
		subTask.CheckIsReady()
	}
}
