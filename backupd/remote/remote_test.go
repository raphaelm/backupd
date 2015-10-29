package remote_test

import (
	"github.com/raphaelm/backupd/backupd/model"
	"github.com/raphaelm/backupd/backupd/remote"
	"testing"
)

func TestRemoteLoadDriver(t *testing.T) {
	r := model.Remote{Driver: "ssh"}
	_, ok := remote.LoadDriver(r)
	if !ok {
		t.Fatal("Driver not foud.")
	}
}

func TestRemoteLoadUnknownDriver(t *testing.T) {
	r := model.Remote{Driver: "foo"}
	driver, ok := remote.LoadDriver(r)
	if ok || driver != nil {
		t.Fatal("Invalid driver found.")
	}
}
