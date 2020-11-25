package osutils

import (
	"bytes"
	"os/exec"
)

type ECmdState int

const (
	ECmdStateError ECmdState = iota
	ECmdStateReady
	ECmdStateRunning
	ECmdStateFinish
)

type Command interface {
	GetCmdState() ECmdState
	GetCmdError() error
	Run()
	RunAsyn()
	Kill() error
	GetStdout() string
	GetStderr() string
	SetStdin(in string) error
}

type command struct {
	cmd   *exec.Cmd
	state ECmdState
	err   error

	stdin  bytes.Buffer
	stdout bytes.Buffer
	stderr bytes.Buffer

	// stdinPipe  io.WriteCloser
	// stdoutPipe io.ReadCloser
	// stderrPipe io.ReadCloser
}

var _ Command = (*command)(nil)

type CmdStateError struct {
	msg string
}

var _ error = (*CmdStateError)(nil)

func (e *CmdStateError) Error() string {
	return e.msg
}

type CmdTimeoutError struct{}

var _ error = (*CmdTimeoutError)(nil)

func (e *CmdTimeoutError) Error() string {
	return "timeout"
}

func NewCommand(name string, args ...string) Command {
	cmd := new(command)
	cmd.cmd = exec.Command(name, args...)
	cmd.state = ECmdStateReady
	cmd.cmd.Stdin = &cmd.stdin
	cmd.cmd.Stdout = &cmd.stdout
	cmd.cmd.Stderr = &cmd.stderr

	// if stdoutPipe, err := cmd.cmd.StdoutPipe(); err != nil {
	// 	return nil, err
	// } else {
	// 	cmd.stdoutPipe = stdoutPipe
	// }

	// if stdinPipe, err := cmd.cmd.StdinPipe(); err != nil {
	// 	return nil, err
	// } else {
	// 	cmd.stdinPipe = stdinPipe
	// }

	// if stderrPipe, err := cmd.cmd.StderrPipe(); err != nil {
	// 	return nil, err
	// } else {
	// 	cmd.stderrPipe = stderrPipe
	// }

	return cmd
}

func (c *command) GetCmdState() ECmdState {
	return c.state
}

func (c *command) GetCmdError() error {
	return c.err
}

func (c *command) Run() {
	c.state = ECmdStateRunning
	if err := c.cmd.Run(); err != nil {
		c.state = ECmdStateError
		c.err = err

		return
	}

	c.state = ECmdStateFinish
}

func (c *command) RunAsyn() {
	c.state = ECmdStateRunning

	go func() {
		if err := c.cmd.Run(); err != nil {
			c.state = ECmdStateError
			c.err = err

			return
		}

		c.state = ECmdStateFinish
	}()
}

func (c *command) Kill() error {
	if c.state == ECmdStateRunning {
		return c.cmd.Process.Kill()
	}

	return &CmdStateError{msg: "cmd not running"}
}

func (c *command) GetStdout() string {
	return c.stdout.String()
}

// func (c *command) GetStderr(ctx context.Context) (string, error) {
// 	for {
// 		select {
// 		case <-ctx.Done():
// 			{
// 				return "", &CmdTimeoutError{}
// 			}
// 		default:
// 			{
// 				if c.state == ECmdStateFinish || c.state == ECmdStateError {
// 					// _, err := c.stderr.ReadFrom(c.stderrPipe)
// 					// if err != nil {
// 					// 	return "", err
// 					// }
// 					return c.stderr.String(), nil
// 				}
// 			}
// 		}
// 	}
// }

func (c *command) GetStderr() string {
	return c.stderr.String()
}

func (c *command) SetStdin(in string) error {
	if c.state != ECmdStateRunning && c.state != ECmdStateReady {
		return &CmdStateError{msg: "set stdin need cmd in running or ready state"}
	}

	// _, err := c.stdinPipe.Write([]byte(in))
	// if err != nil {
	// 	return err
	// }

	_, err := c.stdin.WriteString(in)
	if err != nil {
		return err
	}

	return nil
}
