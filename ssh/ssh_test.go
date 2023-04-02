package ssh

import "testing"

func TestConnectWithPrivateKey(t *testing.T) {
	session, err := ConnectWithPrivateKey("baimeow.cn", "root")
	if err != nil {
		t.Error(err)
	}
	output, err := session.CombinedOutput("id")
	if err != nil {
		t.Error()
	}
	t.Log(string(output))
}
