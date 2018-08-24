package vagrant

import (
	"testing"
)

func init() {
	successfulOutput["port"] = `
1534865103,default,metadata,provider,virtualbox
1534865103,,ui,info,The forwarded ports for the machine are listed below. Please note that\nthese values may differ from values configured in the Vagrantfile if the\nprovider supports automatic port collision detection and resolution.
1534865103,,ui,info,
1534865103,,ui,info,    22 (guest) => 2222 (host)
1534865103,default,forwarded_port,22,2222
`
}

func TestPortCommand_Run(t *testing.T) {
	client := newMockVagrantClient()
	cmd := client.Port()
	if err := cmd.Run(); err != nil {
		t.Fatalf("Command failed to run: %v", err)
	}
	if cmd.Error != nil {
		t.Fatalf("Command return error: %v", cmd.Error)
	}

	if len(cmd.ForwardedPorts) != 1 {
		t.Fatalf("Expected forwarded ports for 1 VM; got %v", len(cmd.ForwardedPorts))
	}

	forwardedPorts, ok := cmd.ForwardedPorts["default"]
	if !ok {
		t.Fatal("Expected forwarded ports for 'default', but there were none.")
	}

	if len(forwardedPorts) != 1 {
		t.Fatalf("Expected 1 forwarded port; got %v", len(forwardedPorts))
	}
	if forwardedPorts[0].Guest != 22 || forwardedPorts[0].Host != 2222 {
		t.Errorf("Expected guest port 22 -> host 2222; got %v", forwardedPorts[0])
	}
}
