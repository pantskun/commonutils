package taskutils

import (
	"context"
	"time"

	"github.com/pantskun/commonutils/container"
)

const closeTimeout = 5 * time.Second

type ETaskPoolState int

const (
	ETaskPoolStateError ETaskPoolState = iota
	ETaskPoolStateClosed
	ETaskPoolStateRunning
	ETaskPoolStateClosing
)

type TaskPool interface {
	GetTaskPoolState() ETaskPoolState
	GetAllTaskNum() int
	GetFinishedTaskNum() int
	GetErrorTaskNum() int
	GetWaitingTaskNum() int
	GetReadyTaskNum() int
	Run()
	Close()
	AddTask(task Task)
}

type taskPool struct {
	allTaskList      container.Vector
	waitingTaskList  container.Vector
	errorTaskList    container.Vector
	readyTaskQueue   container.Queue
	finishedTaskList container.Vector

	state ETaskPoolState
}

func NewTaskPool() TaskPool {
	return &taskPool{
		state: ETaskPoolStateClosed,
	}
}

func (p *taskPool) GetTaskPoolState() ETaskPoolState {
	return p.state
}

func (p *taskPool) GetAllTaskNum() int {
	return p.allTaskList.Size()
}

func (p *taskPool) GetFinishedTaskNum() int {
	return p.finishedTaskList.Size()
}

func (p *taskPool) GetErrorTaskNum() int {
	return p.errorTaskList.Size()
}

func (p *taskPool) GetWaitingTaskNum() int {
	return p.waitingTaskList.Size()
}

func (p *taskPool) GetReadyTaskNum() int {
	return p.readyTaskQueue.Size()
}

func (p *taskPool) Run() {
	p.state = ETaskPoolStateRunning

	// 开始循环，直到TaskPool状态不为Running
	go func() {
		for p.state == ETaskPoolStateRunning {
			// 检测WaitingTaskList，将Ready的Task放入ReadyTaskQueue
			p.getReadyTaskFromWaitingList()

			if p.readyTaskQueue.IsEmpty() {
				continue
			}

			// 从ReadyTaskQueue中取出头部任务进行执行
			t := p.readyTaskQueue.Pop().(Task)

			go func() {
				t.Run()

				// 根据任务执行后的状态，放入ErrorTaskList和FinishedTaskList
				if t.GetState() == ETaskStateError {
					p.errorTaskList.Add(t)
				}

				if t.GetState() == ETaskStateFinished {
					p.finishedTaskList.Add(t)
				}
			}()
		}

		// 结束循环，Pool关闭
		p.state = ETaskPoolStateClosed
	}()
}

func (p *taskPool) Close() {
	// 通知关闭
	p.state = ETaskPoolStateClosing

	ctx, cancel := context.WithTimeout(context.TODO(), closeTimeout)
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

func (p *taskPool) AddTask(task Task) {
	p.waitingTaskList.Add(task)
	p.allTaskList.Add(task)
}

func (p *taskPool) getReadyTaskFromWaitingList() {
	for i := 0; ; {
		if !(i < p.waitingTaskList.Size()) {
			break
		}

		task := p.waitingTaskList.Get(i).(Task)
		if task.GetState() == ETaskStateReady {
			p.readyTaskQueue.Push(task)
			p.waitingTaskList.Remove(i)

			continue
		}

		if task.GetState() == ETaskStateError {
			p.errorTaskList.Add(task)
			p.waitingTaskList.Remove(i)

			continue
		}

		i += 1
	}
}
