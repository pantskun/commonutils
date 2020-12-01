package pathutils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetURLPath(t *testing.T) {
	type TestCase struct {
		url      string
		expected string
	}

	testCases := []TestCase{
		{url: "http://test/test", expected: "test/test"},
	}

	for _, testCase := range testCases {
		assert.Equal(t, GetURLPath(testCase.url), testCase.expected)
	}
}
