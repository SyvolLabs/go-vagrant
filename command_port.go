package vagrant

// PortCommand specifies the options and output of vagrant port.
type PortCommand struct {
	BaseCommand
	PortResponse

	// MachineName is the vagrant machine you are interested in. If your
	// Vagrantfile only brings up a single machine, you do not need to specify
	// this. However, if your Vagrantfile brings up multiple machines, you MUST
	// specify this! For some reason, this is the only vagrant command that
	// cannot handle multiple machines.
	MachineName string
}

// Port will return information about ports forwarded from the host to the
// guest machine. After setting options as appropriate, you must call Run()
// or Start() followed by Wait() to execute. Output will be in ForwardedPorts
// with any error in Error.
//
// This appears to be the only vagrant command that cannot handle multi-vm
// Vagrantfiles for some reason. If your Vagrantfile brings up multiple
// machines, you MUST specify which machine you are interested in by specifying
// the PortCommand.MachineName option!
func (client *VagrantClient) Port() *PortCommand {
	return &PortCommand{
		BaseCommand:  newBaseCommand(client),
		PortResponse: newPortResponse(),
	}
}

func (cmd *PortCommand) init() error {
	if cmd.MachineName != "" {
		return cmd.BaseCommand.init(&cmd.PortResponse, "port", cmd.MachineName)
	}
	return cmd.BaseCommand.init(&cmd.PortResponse, "port")
}

// Run the command
func (cmd *PortCommand) Run() error {
	if err := cmd.Start(); err != nil {
		return err
	}
	return cmd.Wait()
}

// Start the command. You must call Wait() to complete execution.
func (cmd *PortCommand) Start() error {
	if err := cmd.init(); err != nil {
		return err
	}
	return cmd.BaseCommand.Start()
}
