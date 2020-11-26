package taskutils

import (
	"github.com/pantskun/commonutils/container"
)

// type Task interface {
// 	GetState() ETaskState
// }

type ETaskState int

const (
	ETaskStateError ETaskState = iota
	ETaskStateWaiting
	ETaskStateReady
	ETaskStateRunning
	ETaskStateFinished
)

type Task interface {
	Equal(other container.Element) bool
	GetState() ETaskState
	Run()
	GetError() error

	checkIsReady()
}

type task struct {
	name        string
	do          func() error
	err         error
	state       ETaskState
	preTaskList []Task
	subTaskList []Task
}

// NewTask 创建新的任务.
//  do func() error: 任务的内容
//  preTaskks: 任务的前置任务
func NewTask(name string, do func() error, preTasks ...Task) Task {
	newTask := task{name: name, do: do, state: ETaskStateWaiting, preTaskList: preTasks}

	for _, preTask := range preTasks {
		if preTask != nil {
			task, _ := preTask.(*task)
			task.subTaskList = append(task.subTaskList, &newTask)
		}
	}

	if len(preTasks) == 0 {
		newTask.state = ETaskStateReady
	}

	return &newTask
}

// Equal
// 判断是否指向同一个task.
func (t *task) Equal(other container.Element) bool {
	value, ok := other.(Task)
	if !ok {
		return false
	}

	return t == value
}

// GetState 获取任务状态.
func (t *task) GetState() ETaskState {
	return t.state
}

// CheckIsReady 检查任务前置是否都已完成.
func (t *task) checkIsReady() {
	for _, preTask := range t.preTaskList {
		if preTask == nil {
			continue
		}

		if preTask.GetState() == ETaskStateError {
			t.state = ETaskStateError
			return
		}

		if preTask.GetState() != ETaskStateFinished {
			t.state = ETaskStateWaiting
			return
		}
	}

	t.state = ETaskStateReady
}

func (t *task) Run() {
	if t.state != ETaskStateReady {
		return
	}

	t.state = ETaskStateRunning

	t.err = t.do()
	if t.err != nil {
		t.state = ETaskStateError
		return
	}

	t.state = ETaskStateFinished

	for _, subTask := range t.subTaskList {
		subTask.checkIsReady()
	}
}

func (t *task) GetError() error {
	return t.err
}
