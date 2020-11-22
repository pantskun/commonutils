package taskutils

import (
	"github.com/pantskun/commonutils/container"
)

type ETaskPoolState int

const (
	ETaskPoolStateError   ETaskPoolState = 0
	ETaskPoolStateRunning ETaskPoolState = 1
	ETaskPoolStateClosing ETaskPoolState = 2
	ETaskPoolStateClosed  ETaskPoolState = 3
)

type TaskPool struct {
	allTaskList      container.Vector
	waitingTaskList  []*Task
	errorTaskList    []*Task
	readyTaskQueue   container.Queue
	finishedTaskList []*Task

	state ETaskPoolState
}

func NewTaskPool() *TaskPool {
	return &TaskPool{state: ETaskPoolStateRunning}
}

func (p *TaskPool) GetAllTaskNum() int {
	return p.allTaskList.Size()
}

func (p *TaskPool) GetFinishedTaskNum() int {
	return len(p.finishedTaskList)
}

func (p *TaskPool) GetErrorTaskNum() int {
	return len(p.errorTaskList)
}

func (p *TaskPool) GetWaitingTaskNum() int {
	return len(p.waitingTaskList)
}

func (p *TaskPool) GetReadyTaskNum() int {
	return p.readyTaskQueue.Size()
}

func (p *TaskPool) Run() {
	// 开始循环，直到TaskPool状态不为Running
	go func() {
		for p.state == ETaskPoolStateRunning {
			// 检测WaitingTaskList，将Ready的Task放入ReadyTaskQueue
			p.checkWaitingTask()

			if p.readyTaskQueue.IsEmpty() {
				continue
			}

			// 从ReadyTaskQueue中取出头部任务进行执行
			t := p.readyTaskQueue.Pop().(*Task)

			t.Run()

			// 根据任务执行后的状态，放入ErrorTaskList和FinishedTaskList
			if t.state == ETaskStateError {
				p.errorTaskList = append(p.errorTaskList, t)
			}

			if t.state == ETaskStateFinished {
				p.finishedTaskList = append(p.finishedTaskList, t)
			}
		}

		// 结束循环，Pool关闭
		p.state = ETaskPoolStateClosed
	}()
}

func (p *TaskPool) Close() {
	// 通知关闭
	p.state = ETaskPoolStateClosing

	// 等待关闭
	for {
		if p.state == ETaskPoolStateClosed {
			return
		}
	}
}

func (p *TaskPool) AddTask(task *Task) {
	p.waitingTaskList = append(p.waitingTaskList, task)
	p.allTaskList.Add(task)
}

func (p *TaskPool) checkWaitingTask() {
	for i, task := range p.waitingTaskList {
		if task.state == ETaskStateReady {
			p.readyTaskQueue.Push(task)

			if i == len(p.waitingTaskList)-1 {
				p.waitingTaskList = p.waitingTaskList[:i]
			} else {
				p.waitingTaskList = append(p.waitingTaskList[:i], p.waitingTaskList[i+1:]...)
			}

		}
	}
}
