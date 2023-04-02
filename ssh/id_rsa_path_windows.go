package ssh

import "os"

var id_rsa_path = os.Getenv("HOMEPATH") + "/.ssh/id_rsa"
