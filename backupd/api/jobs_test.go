package api_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/raphaelm/backupd/backupd/model"
	"github.com/stretchr/testify/assert"
)

func TestJobs(t *testing.T) {
	remote := &model.Remote{Driver: "ssh", Location: "foo"}
	store.SaveRemote(remote)

	jobsUrl := fmt.Sprintf("%s/jobs", server.URL)

	jJson := `{"remote_id": 123, "job_name": "mysql", "interval": 1200000000000, "preferred_time": "03:00"}`
	request, err := http.NewRequest("POST", jobsUrl, strings.NewReader(jJson))
	res, err := http.DefaultClient.Do(request)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, 500, res.StatusCode)

	j := model.Job{
		RemoteID:      remote.ID,
		JobName:       "mysql",
		Interval:      time.Duration(20 * time.Minute),
		PreferredTime: model.ClockTime{Hour: 3, Minute: 0},
	}

	jJson = fmt.Sprintf(`{"remote_id": %d, "job_name": "mysql", "interval": 1200000000000, "preferred_time": "03:00"}`, remote.ID)
	request, err = http.NewRequest("POST", jobsUrl, strings.NewReader(jJson))

	res, err = http.DefaultClient.Do(request)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, 201, res.StatusCode)
	jslice, err := store.Jobs()
	assert.Equal(t, 1, len(jslice))
	assert.Equal(t, "mysql", jslice[0].JobName)
	assert.Equal(t, remote.ID, jslice[0].RemoteID)
	assert.Equal(t, 3, jslice[0].PreferredTime.Hour)
	assert.Equal(t, 0, jslice[0].PreferredTime.Minute)
	j.ID = jslice[0].ID

	request, err = http.NewRequest("GET", jobsUrl, nil)
	res, err = http.DefaultClient.Do(request)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, 200, res.StatusCode)
	defer res.Body.Close()
	var target model.Jobs
	json.NewDecoder(res.Body).Decode(&target)
	assert.Equal(t, 1, len(target))
	assert.Equal(t, j, target[0])

	jJson = `{"job_name": "rsync"}`
	request, err = http.NewRequest("PUT", jobsUrl+"/"+strconv.Itoa(int(j.ID)), strings.NewReader(jJson))
	res, err = http.DefaultClient.Do(request)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, 200, res.StatusCode)
	jslice, err = store.Jobs()
	assert.Equal(t, 1, len(jslice))
	assert.Equal(t, "rsync", jslice[0].JobName)

	request, err = http.NewRequest("DELETE", jobsUrl+"/"+strconv.Itoa(int(j.ID)), nil)
	res, err = http.DefaultClient.Do(request)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, 200, res.StatusCode)
	jslice, err = store.Jobs()
	assert.Equal(t, 0, len(jslice))

	store.DeleteRemote(remote)
}
