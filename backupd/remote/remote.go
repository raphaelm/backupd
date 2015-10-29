package remote

import (
	"github.com/raphaelm/backupd/backupd/model"
	"github.com/raphaelm/backupd/backupd/remote/ssh"
	"io"
)

type RemoteDriver interface {
	GetPipe(module string) *io.Writer
}

func LoadDriver(r model.Remote) (driver RemoteDriver, ok bool) {
	switch r.Driver {
	case "ssh":
		return ssh.Load(r), true
	default:
		return nil, false
	}
}
