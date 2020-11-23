package taskutils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewTask(t *testing.T) {
	do := func() error {
		t.Log("testTask")
		return nil
	}

	task1 := NewTask("testTask", do)
	task2 := NewTask("testTask", do, task1)

	// Test NewTask
	assert.NotEqual(t, task1, nil)
	assert.Equal(t, task1.GetState(), ETaskStateReady)
	assert.Equal(t, task2.GetState(), ETaskStateWaiting)

	// Test Run
	task1.Run()
	assert.Equal(t, task1.GetState(), ETaskStateFinished)

	// Test Equal
	assert.Equal(t, task1.Equal(task2), false)
}

func TestCheckIsReady(t *testing.T) {
	do := func() error {
		t.Log("testTask")
		return nil
	}

	type TestCase struct {
		task     *Task
		expected ETaskState
	}

	errorTask := NewTask("testTask", do)
	errorTask.state = ETaskStateError

	readyTask := NewTask("testTask", do)
	readyTask.state = ETaskStateReady

	finishedTask := NewTask("testTask", do)
	finishedTask.state = ETaskStateFinished

	taskWithFinishedPre := NewTask("testTask", do, finishedTask)
	taskWithReadyPre := NewTask("testTask", do, readyTask)
	taskWithErrorPre := NewTask("testTask", do, errorTask)

	testCases := []TestCase{
		{task: taskWithFinishedPre, expected: ETaskStateReady},
		{task: taskWithReadyPre, expected: ETaskStateWaiting},
		{task: taskWithErrorPre, expected: ETaskStateError},
	}

	for _, testCase := range testCases {
		testCase.task.CheckIsReady()
		assert.Equal(t, testCase.task.GetState(), testCase.expected)
	}
}
