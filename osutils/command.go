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
	RunAsyn() <-chan bool
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

// NewCommand 创建Command对象.
func NewCommand(name string, args ...string) Command {
	cmd := new(command)
	cmd.cmd = exec.Command(name, args...)
	cmd.state = ECmdStateReady
	cmd.cmd.Stdin = &cmd.stdin
	cmd.cmd.Stdout = &cmd.stdout
	cmd.cmd.Stderr = &cmd.stderr

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

func (c *command) RunAsyn() <-chan bool {
	c.state = ECmdStateRunning

	endChan := make(chan bool, 1)

	go func() {
		defer func() { endChan <- true }()

		if err := c.cmd.Run(); err != nil {
			c.state = ECmdStateError
			c.err = err

			return
		}

		c.state = ECmdStateFinish
	}()

	return endChan
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

func (c *command) GetStderr() string {
	return c.stderr.String()
}

func (c *command) SetStdin(in string) error {
	if c.state != ECmdStateRunning && c.state != ECmdStateReady {
		return &CmdStateError{msg: "set stdin need cmd in running or ready state"}
	}

	_, err := c.stdin.WriteString(in)
	if err != nil {
		return err
	}

	return nil
}
