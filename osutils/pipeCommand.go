package osutils

import (
	"io"
	"os/exec"
)

type pipeCommand struct {
	cmd   *exec.Cmd
	state ECmdState

	stdinPipe  io.WriteCloser
	stdoutPipe io.ReadCloser
	stderrPipe io.ReadCloser
}
