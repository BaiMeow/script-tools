package ssh

import "testing"

func TestConnectWithPrivateKey(t *testing.T) {
	cmd, err := ConnectWithPrivateKey("baimeow.cn", "root")
	if err != nil {
		t.Error(err)
	}
	output, err := cmd.Run("id")
	if err != nil {
		t.Error()
	}
	t.Log(string(output))
	output, err = cmd.Run("ls")
	if err != nil {
		t.Error()
	}
	t.Log(string(output))
}
