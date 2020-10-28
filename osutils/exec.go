package osutils

import (
	"bytes"
	"context"
	"os/exec"
)

type ECmdState int

const (
	ECmdStateError   = 0
	ECmdStateReady   = 1
	ECmdStateRunning = 2
	ECmdStateFinish  = 3
)

type Command interface {
	GetCmdState() ECmdState
	Run() error
	GetStdout(ctx context.Context) (string, error)
	GetStderr(ctx context.Context) (string, error)
	SetStdin(in string) error
}

type command struct {
	cmd    *exec.Cmd
	state  ECmdState
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

func (c *command) Run() error {
	c.state = ECmdStateRunning
	if err := c.cmd.Run(); err != nil {
		c.state = ECmdStateError
		return err
	}

	c.state = ECmdStateFinish

	return nil
}

func (c *command) GetStdout(ctx context.Context) (string, error) {
	for {
		select {
		case <-ctx.Done():
			{
				return "", nil
			}
		default:
			{
				if c.state == ECmdStateFinish || c.state == ECmdStateError {
					return c.stdout.String(), nil
				}
			}
		}
	}
}

func (c *command) GetStderr(ctx context.Context) (string, error) {
	for {
		select {
		case <-ctx.Done():
			{
				return "", nil
			}
		default:
			{
				if c.state == ECmdStateFinish || c.state == ECmdStateError {
					return c.stderr.String(), nil
				}
			}
		}
	}
}

func (c *command) SetStdin(in string) error {
	if c.state != ECmdStateRunning && c.state != ECmdStateReady {
		return &CmdStateError{msg: "set stdin need cmd in running or ready state"}
	}

	c.stdin.WriteString(in)

	return nil
}
