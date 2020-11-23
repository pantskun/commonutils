package taskutils

import (
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestTaskPool(t *testing.T) {
	do := func() error {
		t.Log("test task")
		return nil
	}

	errorDo := func() error {
		return errors.New("test error")
	}

	// test NewTaskPool
	t.Log("test NewTaskPool")

	taskPool := NewTaskPool()
	assert.Equal(t, taskPool.GetTaskPoolState(), ETaskPoolStateStop)

	// test AddTask
	t.Log("test AddTask")

	taskWithoutPre := NewTask("testTask1", do)
	taskPool.AddTask(taskWithoutPre)

	errorTask := NewTask("errorTask", errorDo)
	taskPool.AddTask(errorTask)

	assert.Equal(t, taskPool.GetAllTaskNum(), 2)
	assert.Equal(t, taskPool.GetWaitingTaskNum(), 2)
	assert.Equal(t, taskPool.GetErrorTaskNum(), 0)
	assert.Equal(t, taskPool.GetFinishedTaskNum(), 0)

	// test Run
	t.Log("test Run")

	taskPool.Run()
	assert.Equal(t, taskPool.GetTaskPoolState(), ETaskPoolStateRunning)

	time.Sleep(1 * time.Second)
	assert.Equal(t, taskWithoutPre.GetState(), ETaskStateFinished)
	assert.Equal(t, errorTask.GetState(), ETaskStateError)
	assert.Equal(t, taskPool.GetReadyTaskNum(), 0)
	assert.Equal(t, taskPool.GetErrorTaskNum(), 1)
	assert.Equal(t, taskPool.GetFinishedTaskNum(), 1)

	// test Close
	t.Log("test Close")

	taskPool.Close()
	assert.Equal(t, taskPool.GetTaskPoolState(), ETaskPoolStateClosed)
}
