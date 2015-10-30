package datastore_test

import (
	"testing"

	"github.com/raphaelm/backupd/backupd/datastore"
)

func TestSuite(t *testing.T) {
	s := datastore.MockStore()
	runTestSuite(s, t)
}
