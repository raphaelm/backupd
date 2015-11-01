package api

import (
	"encoding/json"
	"net/http"

	"github.com/raphaelm/backupd/backupd/datastore"
)

func RemoteIndex(store datastore.DataStore, w http.ResponseWriter, r *http.Request) {
	remotes, err := store.Remotes()
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(remotes); err != nil {
		panic(err)
	}
}
