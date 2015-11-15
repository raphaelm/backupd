package api_test

import (
	"io"
	"net/http/httptest"

	"github.com/raphaelm/backupd/backupd/api"
	"github.com/raphaelm/backupd/backupd/datastore"
)

var (
	server *httptest.Server
	store  datastore.DataStore
	reader io.Reader
)

func init() {
	store = datastore.MockStore()
	server = httptest.NewServer(api.Router(store))
}
