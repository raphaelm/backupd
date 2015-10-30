package datastore

import "github.com/raphaelm/backupd/backupd/model"

type DataStore interface {
	Remote(id int64) (remote model.Remote, err error)
	Remotes() (remotes []model.Remote, err error)
	SaveRemote(remote *model.Remote) (created bool, err error)
	DeleteRemote(remote *model.Remote) (deleted bool, err error)
}
