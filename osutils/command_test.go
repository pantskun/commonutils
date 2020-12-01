package osutils

import (
	"log"
	"path"
	"testing"

	"github.com/pantskun/commonutils/pathutils"
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

	log.Println(cmd.GetStdout())
	log.Println(cmd.GetStderr())

	// assert.Equal(t, cmd.GetStderr(), "")
	// assert.Equal(t, cmd.GetStdout(), "test")
}
