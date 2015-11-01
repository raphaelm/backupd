package datastore

import "github.com/raphaelm/backupd/backupd/model"

type DataStore interface {
	Remote(id int64) (remote model.Remote, err error)
	Remotes() (remotes model.Remotes, err error)
	SaveRemote(remote *model.Remote) (created bool, err error)
	DeleteRemote(remote *model.Remote) (deleted bool, err error)

	Job(id int64) (job model.Job, err error)
	Jobs() (jobs model.Jobs, err error)
	SaveJob(job *model.Job) (created bool, err error)
	DeleteJob(job *model.Job) (deleted bool, err error)

	JobsForRemote(remote *model.Remote) (jobs model.Jobs, err error)

	Backup(id int64) (backup model.Backup, err error)
	Backups() (backup model.Backups, err error)
	SaveBackup(backup *model.Backup) (created bool, err error)
	DeleteBackup(backup *model.Backup) (deleted bool, err error)

	BackupsForJob(job *model.Job) (backups model.Backups, err error)
}
