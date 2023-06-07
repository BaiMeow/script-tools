package ssh

import (
	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
	"io"
	"os"
)

type SFTP struct {
	*sftp.Client
}

func NewSFTP(c *ssh.Client) (*SFTP, error) {
	newClient, err := sftp.NewClient(c)
	if err != nil {
		return nil, err
	}
	return &SFTP{newClient}, nil
}

func (s *SFTP) Upload(reader io.Reader, remotePath string) error {
	remoteFile, err := s.OpenFile(remotePath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC)
	if err != nil {
		return err
	}
	defer remoteFile.Close()
	_, err = io.Copy(remoteFile, reader)
	if err != nil {
		return err
	}
	return nil
}

func (s *SFTP) Download(remotePath string, writer io.Writer) error {
	remoteFile, err := s.Open(remotePath)
	if err != nil {
		return err
	}
	defer remoteFile.Close()
	_, err = io.Copy(writer, remoteFile)
	if err != nil {
		return err
	}
	return nil
}
