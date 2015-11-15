package api

import (
	"encoding/json"
	"net/http"

	"github.com/raphaelm/backupd/backupd/datastore"
)

func BackupIndex(store datastore.DataStore, w http.ResponseWriter, r *http.Request) {
	jobs, err := store.Backups()
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(jobs); err != nil {
		panic(err)
	}
}
