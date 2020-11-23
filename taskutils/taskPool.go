package taskutils

import (
	"context"
	"time"

	"github.com/pantskun/commonutils/container"
)

type ETaskPoolState int

const (
	ETaskPoolStateError ETaskPoolState = iota
	ETaskPoolStateStop
	ETaskPoolStateRunning
	ETaskPoolStateClosing
	ETaskPoolStateClosed
)

type TaskPool struct {
	allTaskList      container.Vector
	waitingTaskList  container.Vector
	errorTaskList    container.Vector
	readyTaskQueue   container.Queue
	finishedTaskList container.Vector

	state ETaskPoolState
}

func NewTaskPool() *TaskPool {
	return &TaskPool{
		state: ETaskPoolStateStop,
	}
}

func (p *TaskPool) GetTaskPoolState() ETaskPoolState {
	return p.state
}

func (p *TaskPool) GetAllTaskNum() int {
	return p.allTaskList.Size()
}

func (p *TaskPool) GetFinishedTaskNum() int {
	return p.finishedTaskList.Size()
}

func (p *TaskPool) GetErrorTaskNum() int {
	return p.errorTaskList.Size()
}

func (p *TaskPool) GetWaitingTaskNum() int {
	return p.waitingTaskList.Size()
}

func (p *TaskPool) GetReadyTaskNum() int {
	return p.readyTaskQueue.Size()
}

func (p *TaskPool) Run() {
	p.state = ETaskPoolStateRunning

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
				p.errorTaskList.Add(t)
			}

			if t.state == ETaskStateFinished {
				p.finishedTaskList.Add(t)
			}
		}

		// 结束循环，Pool关闭
		p.state = ETaskPoolStateClosed
	}()
}

func (p *TaskPool) Close() {
	// 通知关闭
	p.state = ETaskPoolStateClosing

	ctx, cancel := context.WithTimeout(context.TODO(), 5*time.Second)
	defer cancel()

	// 等待关闭
	for {
		select {
		case <-ctx.Done():
			return
		default:
		}

		if p.state == ETaskPoolStateClosed {
			return
		}
	}
}

func (p *TaskPool) AddTask(task *Task) {
	p.waitingTaskList.Add(task)
	p.allTaskList.Add(task)
}

func (p *TaskPool) checkWaitingTask() {
	for i := 0; ; {
		if !(i < p.waitingTaskList.Size()) {
			break
		}

		task := p.waitingTaskList.Get(i).(*Task)
		if task.state == ETaskStateReady {
			p.readyTaskQueue.Push(task)
			p.waitingTaskList.Remove(i)

			continue
		}

		i += 1
	}
}
