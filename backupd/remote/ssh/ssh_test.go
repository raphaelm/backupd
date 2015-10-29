package ssh_test

import (
	"github.com/raphaelm/backupd/backupd/model"
	"github.com/raphaelm/backupd/backupd/remote/ssh"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseLocation(t *testing.T) {
	assert := assert.New(t)

	r := model.Remote{
		Driver:   "ssh",
		Location: "hostname",
	}
	rd := ssh.Load(r)
	assert.Equal(rd.Host, "hostname")
	assert.Equal(rd.Port, 22)
	assert.Equal(rd.User, "")

	r = model.Remote{
		Driver:   "ssh",
		Location: "user@hostname",
	}
	rd = ssh.Load(r)
	assert.Equal(rd.Host, "hostname")
	assert.Equal(rd.Port, 22)
	assert.Equal(rd.User, "user")

	r = model.Remote{
		Driver:   "ssh",
		Location: "hostname:42",
	}
	rd = ssh.Load(r)
	assert.Equal(rd.Host, "hostname")
	assert.Equal(rd.Port, 42)
	assert.Equal(rd.User, "")

	r = model.Remote{
		Driver:   "ssh",
		Location: "user@hostname:42",
	}
	rd = ssh.Load(r)
	assert.Equal(rd.Host, "hostname")
	assert.Equal(rd.Port, 42)
	assert.Equal(rd.User, "user")

	r = model.Remote{
		Driver:   "ssh",
		Location: "user@hostname:foo",
	}
	rd = ssh.Load(r)
	assert.Equal(rd.Host, "hostname")
	assert.Equal(rd.Port, 22)
	assert.Equal(rd.User, "user")
}
