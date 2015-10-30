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
