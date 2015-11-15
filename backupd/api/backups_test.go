package api_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/raphaelm/backupd/backupd/model"
	"github.com/stretchr/testify/assert"
)

func TestBackups(t *testing.T) {
	remote := &model.Remote{Driver: "ssh", Location: "foo"}
	store.SaveRemote(remote)
	job := &model.Job{RemoteID: remote.ID, JobName: "mysql"}
	store.SaveJob(job)
	backup := &model.Backup{
		JobID:  job.ID,
		Result: model.BackupFailed,
		Start:  time.Date(2015, 11, 10, 12, 23, 0, 0, time.UTC),
		End:    time.Date(2015, 11, 10, 12, 24, 0, 0, time.UTC),
	}
	store.SaveBackup(backup)

	bUrl := fmt.Sprintf("%s/backups", server.URL)

	request, err := http.NewRequest("GET", bUrl, nil)
	res, err := http.DefaultClient.Do(request)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, 200, res.StatusCode)
	defer res.Body.Close()
	var target model.Backups
	json.NewDecoder(res.Body).Decode(&target)
	assert.Equal(t, 1, len(target))
	assert.Equal(t, *backup, target[0])

	store.DeleteBackup(backup)
	store.DeleteJob(job)
	store.DeleteRemote(remote)
}
