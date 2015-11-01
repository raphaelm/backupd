package datastore_test

import (
	"testing"

	"github.com/raphaelm/backupd/backupd/datastore"
	"github.com/raphaelm/backupd/backupd/model"
	"github.com/stretchr/testify/assert"
)

// Runs the test suite for a given datastore. This does not use Go's
// testing module in a standard way, but runs many small test in one
// big wrapper function. The idea is that the same tests can be run
// for multiple backends (mock, SQL, ...) without duplicating all the
// test code
func runTestSuite(s datastore.DataStore, t *testing.T) {
	testRemotes(s, t)
	testJobs(s, t)
	testBackups(s, t)
	testDeleteCascade(s, t)
}

func testRemotes(s datastore.DataStore, t *testing.T) {
	r := model.Remote{Driver: "ssh", Location: "foo"}
	created, err := s.SaveRemote(&r)
	assert.Equal(t, true, created)
	assert.Nil(t, err)

	r2, err := s.Remote(r.ID)
	assert.Nil(t, err)
	assert.Equal(t, r.Location, r2.Location)
	assert.Equal(t, r.ID, r2.ID)

	r2.Location = "bar"
	created, err = s.SaveRemote(&r2)
	assert.Equal(t, false, created)
	assert.Nil(t, err)

	rslice, err := s.Remotes()
	assert.Nil(t, err)
	assert.Equal(t, 1, len(rslice))
	assert.Equal(t, r2.Location, rslice[0].Location)
	assert.Equal(t, r2.ID, rslice[0].ID)

	s.DeleteRemote(&r2)

	rslice, err = s.Remotes()
	assert.Nil(t, err)
	assert.Equal(t, 0, len(rslice))

	_, err = s.Remote(r.ID)
	assert.NotNil(t, err)
}

func testJobs(s datastore.DataStore, t *testing.T) {
	// Test only valid RemoteIDs are accepted
	j := model.Job{JobName: "foo", RemoteID: 0}
	created, err := s.SaveJob(&j)
	assert.NotNil(t, err)

	r := model.Remote{Driver: "ssh", Location: "foo"}
	created, err = s.SaveRemote(&r)
	r2 := model.Remote{Driver: "ssh", Location: "bar"}
	created, err = s.SaveRemote(&r2)

	j = model.Job{JobName: "foo", RemoteID: r.ID}
	created, err = s.SaveJob(&j)
	assert.Equal(t, true, created)
	assert.Nil(t, err)

	j2, err := s.Job(j.ID)
	assert.Nil(t, err)
	assert.Equal(t, j.JobName, j2.JobName)
	assert.Equal(t, j.ID, j2.ID)

	j2.JobName = "bar"
	created, err = s.SaveJob(&j2)
	assert.Equal(t, false, created)
	assert.Nil(t, err)

	jslice, err := s.Jobs()
	assert.Nil(t, err)
	assert.Equal(t, 1, len(jslice))
	assert.Equal(t, j2.JobName, jslice[0].JobName)
	assert.Equal(t, j2.ID, jslice[0].ID)

	jslice, err = s.JobsForRemote(&r)
	assert.Nil(t, err)
	assert.Equal(t, 1, len(jslice))

	jslice, err = s.JobsForRemote(&r2)
	assert.Nil(t, err)
	assert.Equal(t, 0, len(jslice))

	s.DeleteJob(&j2)

	jslice, err = s.Jobs()
	assert.Nil(t, err)
	assert.Equal(t, 0, len(jslice))

	_, err = s.Job(j.ID)
	assert.NotNil(t, err)

	s.DeleteRemote(&r)
	s.DeleteRemote(&r2)
}

func testBackups(s datastore.DataStore, t *testing.T) {
	// Test only valid JobIDs are accepted
	b := model.Backup{Result: model.BackupCompleted, JobID: 0}
	created, err := s.SaveBackup(&b)
	assert.NotNil(t, err)

	r := model.Remote{Driver: "ssh", Location: "foo"}
	created, err = s.SaveRemote(&r)
	j := model.Job{RemoteID: r.ID}
	created, err = s.SaveJob(&j)
	j2 := model.Job{RemoteID: r.ID}
	created, err = s.SaveJob(&j2)

	b = model.Backup{Result: model.BackupCompleted, JobID: j.ID}
	created, err = s.SaveBackup(&b)
	assert.Equal(t, true, created)
	assert.Nil(t, err)

	b2, err := s.Backup(b.ID)
	assert.Nil(t, err)
	assert.Equal(t, b.Result, b2.Result)
	assert.Equal(t, b.ID, b2.ID)

	b2.Result = model.BackupFailed
	created, err = s.SaveBackup(&b2)
	assert.Equal(t, false, created)
	assert.Nil(t, err)

	bslice, err := s.Backups()
	assert.Nil(t, err)
	assert.Equal(t, 1, len(bslice))
	assert.Equal(t, b2.Result, bslice[0].Result)
	assert.Equal(t, b2.ID, bslice[0].ID)

	bslice, err = s.BackupsForJob(&j)
	assert.Nil(t, err)
	assert.Equal(t, 1, len(bslice))

	bslice, err = s.BackupsForJob(&j2)
	assert.Nil(t, err)
	assert.Equal(t, 0, len(bslice))

	s.DeleteBackup(&b2)

	bslice, err = s.Backups()
	assert.Nil(t, err)
	assert.Equal(t, 0, len(bslice))

	_, err = s.Backup(b.ID)
	assert.NotNil(t, err)

	s.DeleteRemote(&r)
}

func testDeleteCascade(s datastore.DataStore, t *testing.T) {
	r := model.Remote{Driver: "ssh", Location: "foo"}
	created, err := s.SaveRemote(&r)
	assert.Equal(t, true, created)
	assert.Nil(t, err)

	j := model.Job{JobName: "foo", RemoteID: r.ID}
	created, err = s.SaveJob(&j)
	assert.Equal(t, true, created)
	assert.Nil(t, err)

	b := model.Backup{Result: model.BackupCompleted, JobID: j.ID}
	created, err = s.SaveBackup(&b)
	assert.Equal(t, true, created)
	assert.Nil(t, err)

	jslice, err := s.Jobs()
	assert.Nil(t, err)
	assert.Equal(t, 1, len(jslice))

	bslice, err := s.Backups()
	assert.Nil(t, err)
	assert.Equal(t, 1, len(bslice))

	s.DeleteRemote(&r)

	jslice, err = s.Jobs()
	assert.Nil(t, err)
	assert.Equal(t, 0, len(jslice))

	bslice, err = s.Backups()
	assert.Nil(t, err)
	assert.Equal(t, 0, len(bslice))
}
