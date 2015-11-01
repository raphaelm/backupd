package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/raphaelm/backupd/backupd/datastore"
	"github.com/raphaelm/backupd/backupd/model"
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

func RemoteAdd(store datastore.DataStore, w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var obj model.Remote
	err = json.Unmarshal(body, &obj)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("Error: %s", err)))
		return
	}

	if obj.ID != 0 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Error: You cannot pass an ID"))
		return
	}
	created, err := store.SaveRemote(&obj)

	if !created || err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("Error: %s", err)))
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func RemoteUpdate(store datastore.DataStore, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.Write([]byte(fmt.Sprintf("Error: %s", err)))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	obj, err := store.Remote(int64(id))
	if obj.ID == 0 || err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(fmt.Sprintf("Error: %s", err)))
		return
	}

	oldid := obj.ID
	err = json.Unmarshal(body, &obj)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("Error: %s", err)))
		return
	}

	if obj.ID != oldid {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Error: You cannot change an ID"))
		return
	}
	_, err = store.SaveRemote(&obj)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("Error: %s", err)))
		return
	}
	w.WriteHeader(http.StatusOK)
}

func RemoteDelete(store datastore.DataStore, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.Write([]byte(fmt.Sprintf("Error: %s", err)))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	obj, err := store.Remote(int64(id))
	if obj.ID == 0 || err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(fmt.Sprintf("Error: %s", err)))
		return
	}

	deleted, err := store.DeleteRemote(&obj)

	if !deleted || err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("Error: %s", err)))
		return
	}

	w.WriteHeader(http.StatusOK)
}
