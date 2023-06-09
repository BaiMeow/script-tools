package ssh

import (
	"errors"
	"github.com/BaiMeow/script-tools/cmd"
	"golang.org/x/crypto/ssh"
	"net"
	"os"
	"reflect"
	"strings"
	"time"
)

type SSH struct {
	// ip:port
	Addr   string
	Config *ssh.ClientConfig
}

func ConnectWithPassword(addr string, user string, password string) (*ssh.Client, error) {
	return (&SSH{
		Addr: tryAppendDefaultPort(addr),
		Config: &ssh.ClientConfig{
			User: user,
			Auth: []ssh.AuthMethod{
				ssh.Password(password),
			},
			HostKeyCallback: ssh.InsecureIgnoreHostKey(),
			Timeout:         30 * time.Second,
		},
	}).Connect()
}

// ConnectWithPrivateKey connects to a remote host using a private key,
// privateKey should be string or byte slice, if it is a string, it will be treated as an id_rsa_path to the private key file.
// if it is a byte slice, it will be treated as the private key itself.
// you can also not to provide privateKey, it will be treated as an id_rsa_path to the private key file, and the id_rsa_path is "~/.ssh/id_rsa" or "%HOME%/.ssh/id_rsa" (on windows).
func ConnectWithPrivateKey(addr string, user string, privateKey ...any) (*ssh.Client, error) {
	var bytes []byte
	var err error
	if privateKey == nil || len(privateKey) == 0 {
		if bytes, err = os.ReadFile(id_rsa_path); err != nil {
			return nil, err
		}
	} else if reflect.TypeOf(privateKey).Kind() == reflect.String {
		if bytes, err = os.ReadFile(privateKey[0].(string)); err != nil {
			return nil, err
		}
	} else if reflect.TypeOf(privateKey[0]).Kind() == reflect.Slice {
		var ok bool
		if bytes, ok = privateKey[0].([]byte); !ok {
			return nil, errors.New("field privateKey should be string or byte slice")
		}
	} else {
		return nil, errors.New("field privateKey should be string or byte slice")
	}

	signer, err := ssh.ParsePrivateKey(bytes)
	if err != nil {
		return nil, err
	}
	return (&SSH{
		Addr: tryAppendDefaultPort(addr),
		Config: &ssh.ClientConfig{
			User: user,
			Auth: []ssh.AuthMethod{
				ssh.PublicKeys(signer),
			},
			HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		},
	}).Connect()
}

// Connect connects to the remote host, if you want to use custom ssh config, you can use it.
func (s *SSH) Connect() (*ssh.Client, error) {
	c, err := ssh.Dial("tcp", s.Addr, s.Config)
	if err != nil {
		return nil, err
	}
	return c, nil
}

func tryAppendDefaultPort(addr string) string {
	if strings.Contains(addr, ":") {
		return addr
	}
	return net.JoinHostPort(addr, "22")
}

func NewCommander(c *ssh.Client) cmd.Commander {
	return &Client{c}
}
