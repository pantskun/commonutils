package osutils

type MultiProcCmd interface {
	GetCmds() []Command
	GetCmdStates() []ECmdState
	GetCmdErrors() []error
	Run()
	RunAsyn()
	Kill() []error
	GetStdouts() []string
	GetStderrs() []string
	SetStdin(in string) []error
}

type multiProcCmd struct {
	cmds []Command
}

var _ MultiProcCmd = (*multiProcCmd)(nil)

func NewMultiProcCmd(procNum int, name string, args ...string) MultiProcCmd {
	cmds := make([]Command, procNum)

	for i := 0; i < procNum; i++ {
		cmds[i] = NewCommand(name, args...)
	}

	return &multiProcCmd{cmds: cmds}
}

func (m *multiProcCmd) GetCmds() []Command {
	return m.cmds
}

func (m *multiProcCmd) GetCmdStates() []ECmdState {
	states := make([]ECmdState, len(m.cmds))

	for i, cmd := range m.cmds {
		states[i] = cmd.GetCmdState()
	}

	return states
}

func (m *multiProcCmd) GetCmdErrors() []error {
	errors := make([]error, len(m.cmds))

	for i, cmd := range m.cmds {
		errors[i] = cmd.GetCmdError()
	}

	return errors
}

func (m *multiProcCmd) Run() {
	for _, cmd := range m.cmds {
		cmd.RunAsyn()
	}

	for _, cmd := range m.cmds {
		for cmd.GetCmdState() == ECmdStateFinish {
		}
	}
}

func (m *multiProcCmd) RunAsyn() {
	go m.Run()
}

func (m *multiProcCmd) Kill() []error {
	errors := make([]error, len(m.cmds))

	for i, cmd := range m.cmds {
		errors[i] = cmd.Kill()
	}

	return errors
}

func (m *multiProcCmd) GetStdouts() []string {
	stdouts := make([]string, len(m.cmds))

	for i, cmd := range m.cmds {
		stdouts[i] = cmd.GetStdout()
	}

	return stdouts
}

func (m *multiProcCmd) GetStderrs() []string {
	stderrs := make([]string, len(m.cmds))

	for i, cmd := range m.cmds {
		stderrs[i] = cmd.GetStderr()
	}

	return stderrs
}
func (m *multiProcCmd) SetStdin(in string) []error {
	errors := make([]error, len(m.cmds))

	for i, cmd := range m.cmds {
		errors[i] = cmd.SetStdin(in)
	}

	return errors
}
