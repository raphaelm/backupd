package ssh

import (
	"github.com/raphaelm/backupd/backupd/model"
	"io"
	"strconv"
	"strings"
)

type Ssh struct {
	Host string
	Port int
	User string
}

func split2(s, sep string) (a, b string) {
	x := strings.SplitN(s, sep, 2)
	return x[0], x[1]
}

func Load(r model.Remote) *Ssh {
	driver := Ssh{}

	host := r.Location
	if strings.Contains(host, "@") {
		driver.User, host = split2(r.Location, "@")
	}
	driver.Port = 22
	if strings.Contains(host, ":") {
		port := ""
		host, port = split2(host, ":")
		p, err := strconv.Atoi(port)
		if err == nil {
			driver.Port = p
		}
	}
	driver.Host = host

	return &driver
}

func (d *Ssh) GetPipe(module string) *io.Writer {
	return nil
}
