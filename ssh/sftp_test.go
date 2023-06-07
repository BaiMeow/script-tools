package ssh

import (
	"bytes"
	"fmt"
	"testing"
)

func TestSFTP(t *testing.T) {
	client, err := ConnectWithPrivateKey("baimeow.cn", "root")
	if err != nil {
		t.Error(t)
		return
	}
	sftp, err := NewSFTP(client)
	if err != nil {
		t.Error(err)
		return
	}
	file := new(bytes.Buffer)
	err = sftp.Download("/etc/bird.conf", file)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(file.String())
}
