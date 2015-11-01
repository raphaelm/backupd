package api_test

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/raphaelm/backupd/backupd/api"
	"github.com/raphaelm/backupd/backupd/datastore"
	"github.com/raphaelm/backupd/backupd/model"
	"github.com/stretchr/testify/assert"
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

func TestRemotes(t *testing.T) {
	// Fetch list of remotes
	remotesUrl := fmt.Sprintf("%s/remotes", server.URL)

	r := model.Remote{Driver: "ssh", Location: "foo"}
	store.SaveRemote(&r)

	request, err := http.NewRequest("GET", remotesUrl, nil)
	res, err := http.DefaultClient.Do(request)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, 200, res.StatusCode)
	defer res.Body.Close()
	var target model.Remotes
	json.NewDecoder(res.Body).Decode(&target)
	assert.Equal(t, 1, len(target))
	assert.Equal(t, r, target[0])
}
