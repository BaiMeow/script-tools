package listen

import (
	"github.com/BaiMeow/script-tools/cmd"
	"testing"
)

func TestListener_ListenTCP(t *testing.T) {
	wait := make(chan struct{})
	listener := Default(func(cmd cmd.Commander) {
		rec, err := cmd.Run("ls")
		if err != nil {
			return
		}
		t.Log("rec:", string(rec), '\n')
		wait <- struct{}{}
	})
	err := listener.ListenTCP()
	if err != nil {
		return
	}
	<-wait
	defer listener.Cancel()
}
