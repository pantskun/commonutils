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
		task     Task
		expected ETaskState
	}

	errorTask := task{name: "errorTask", do: do}
	errorTask.state = ETaskStateError

	readyTask := task{name: "readyTask", do: do}
	readyTask.state = ETaskStateReady

	finishedTask := task{name: "finishedTask", do: do}
	finishedTask.state = ETaskStateFinished

	taskWithFinishedPre := NewTask("taskWithFinishedPre", do, &finishedTask)
	taskWithReadyPre := NewTask("taskWithReadyPre", do, &readyTask)
	taskWithErrorPre := NewTask("taskWithErrorPre", do, &errorTask)

	testCases := []TestCase{
		{task: taskWithFinishedPre, expected: ETaskStateReady},
		{task: taskWithReadyPre, expected: ETaskStateWaiting},
		{task: taskWithErrorPre, expected: ETaskStateError},
	}

	for _, testCase := range testCases {
		testCase.task.checkIsReady()
		assert.Equal(t, testCase.task.GetState(), testCase.expected)
	}
}
