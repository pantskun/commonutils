package osutils

import (
	"context"
	"path"
	"testing"

	"github.com/pantskun/commonutils/pathutils"
)

func TestCommand(t *testing.T) {
	modulePath := pathutils.GetModulePath("commonutils")
	testPath := path.Join(modulePath, "test/testcmd.exe")

	cmd := NewCommand(testPath)

	if err := cmd.SetStdin("1\n"); err != nil {
		t.Fatal(err)
	}

	if err := cmd.SetStdin("12.22\n"); err != nil {
		t.Fatal(err)
	}

	if err := cmd.Run(); err != nil {
		t.Fatal(err)
	}

	if stdout, err := cmd.GetStdout(context.TODO()); err != nil {
		t.Fatal(err)
	} else {
		t.Log(stdout)
	}
}
