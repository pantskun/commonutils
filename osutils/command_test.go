package osutils

import (
	"path"
	"testing"

	"github.com/pantskun/commonutils/pathutils"
	"github.com/stretchr/testify/assert"
)

func TestCommand(t *testing.T) {
	modulePath := pathutils.GetModulePath("commonutils")
	testPath := path.Join(modulePath, "test/testcmd.go")

	cmd := NewCommand("go", "run", testPath)

	if err := cmd.SetStdin("test\n"); err != nil {
		t.Fatal(err)
	}

	cmd.Run()

	if cmd.GetCmdState() == ECmdStateError {
		t.Fatal(cmd.GetCmdError())
	}

	stdout := cmd.GetStdout()
	stderr := cmd.GetStderr()

	assert.Equal(t, stdout, "test\n")
	assert.Equal(t, stderr, "\n")

	// log.Println(cmd.GetStdout())
	// log.Println(cmd.GetStderr())
}
