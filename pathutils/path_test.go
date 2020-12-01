package pathutils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConvertBackslashToSlash(t *testing.T) {
	assert.Equal(t, ConvertBackslashToSlash("a\\b\\c\\"), "a/b/c/")
}

func TestGetParentPath(t *testing.T) {
	assert.Equal(t, GetParentPath("/a/b/c"), "/a/b")
}
