package ssh

import (
	"errors"
	"golang.org/x/crypto/ssh"
	"net"
	"os"
	"reflect"
	"strings"
)

type SSH struct {
	// ip:port
	Addr   string
	Config *ssh.ClientConfig
}

func ConnectWithPassword(addr string, user string, password string) (*ssh.Session, error) {
	return (&SSH{
		Addr: TryAppendDefaultPort(addr),
		Config: &ssh.ClientConfig{
			User: user,
			Auth: []ssh.AuthMethod{
				ssh.Password(password),
			},
			HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		},
	}).Connect()
}

// ConnectWithPrivateKey connects to a remote host using a private key,
// privateKey should be string or byte slice, if it is a string, it will be treated as a id_rsa_path to the private key file.
// if it is a byte slice, it will be treated as the private key itself.
// you can also not to provide privateKey, it will be treated as a id_rsa_path to the private key file, and the id_rsa_path is "~/.ssh/id_rsa" or "%HOME%/.ssh/id_rsa" (on windows).
func ConnectWithPrivateKey(addr string, user string, privateKey ...any) (*ssh.Session, error) {
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
		Addr: TryAppendDefaultPort(addr),
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
func (s *SSH) Connect() (*ssh.Session, error) {
	c, err := ssh.Dial("tcp", s.Addr, s.Config)
	if err != nil {
		return nil, err
	}
	return c.NewSession()
}

func TryAppendDefaultPort(addr string) string {
	if strings.Contains(addr, ":") {
		return addr
	}
	return net.JoinHostPort(addr, "22")
}
