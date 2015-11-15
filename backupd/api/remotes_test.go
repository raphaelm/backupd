package api_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"testing"

	"github.com/raphaelm/backupd/backupd/model"
	"github.com/stretchr/testify/assert"
)

func TestRemotes(t *testing.T) {
	// Fetch list of remotes
	remotesUrl := fmt.Sprintf("%s/remotes", server.URL)

	r := model.Remote{Driver: "ssh", Location: "foo"}
	rJson := `{"driver": "ssh", "location": "foo"}`
	request, err := http.NewRequest("POST", remotesUrl, strings.NewReader(rJson))
	res, err := http.DefaultClient.Do(request)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, 201, res.StatusCode)
	rslice, err := store.Remotes()
	assert.Equal(t, 1, len(rslice))
	assert.Equal(t, "ssh", rslice[0].Driver)
	assert.Equal(t, "foo", rslice[0].Location)
	r.ID = rslice[0].ID

	request, err = http.NewRequest("GET", remotesUrl, nil)
	res, err = http.DefaultClient.Do(request)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, 200, res.StatusCode)
	defer res.Body.Close()
	var target model.Remotes
	json.NewDecoder(res.Body).Decode(&target)
	assert.Equal(t, 1, len(target))
	assert.Equal(t, r, target[0])

	rJson = `{"driver": "ssh", "location": "bar"}`
	request, err = http.NewRequest("PUT", remotesUrl+"/"+strconv.Itoa(int(r.ID)), strings.NewReader(rJson))
	res, err = http.DefaultClient.Do(request)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, 200, res.StatusCode)
	rslice, err = store.Remotes()
	assert.Equal(t, 1, len(rslice))
	assert.Equal(t, "ssh", rslice[0].Driver)
	assert.Equal(t, "bar", rslice[0].Location)

	request, err = http.NewRequest("DELETE", remotesUrl+"/"+strconv.Itoa(int(r.ID)), nil)
	res, err = http.DefaultClient.Do(request)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, 200, res.StatusCode)
	rslice, err = store.Remotes()
	assert.Equal(t, 0, len(rslice))
}
