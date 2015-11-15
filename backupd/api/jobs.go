package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/raphaelm/backupd/backupd/datastore"
	"github.com/raphaelm/backupd/backupd/model"
)

func JobIndex(store datastore.DataStore, w http.ResponseWriter, r *http.Request) {
	jobs, err := store.Jobs()
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(jobs); err != nil {
		panic(err)
	}
}

func JobAdd(store datastore.DataStore, w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var obj model.Job
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
	created, err := store.SaveJob(&obj)

	if !created || err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("Error: %s", err)))
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func JobUpdate(store datastore.DataStore, w http.ResponseWriter, r *http.Request) {
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

	obj, err := store.Job(int64(id))
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
	_, err = store.SaveJob(&obj)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("Error: %s", err)))
		return
	}
	w.WriteHeader(http.StatusOK)
}

func JobDelete(store datastore.DataStore, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.Write([]byte(fmt.Sprintf("Error: %s", err)))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	obj, err := store.Job(int64(id))
	if obj.ID == 0 || err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(fmt.Sprintf("Error: %s", err)))
		return
	}

	deleted, err := store.DeleteJob(&obj)

	if !deleted || err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("Error: %s", err)))
		return
	}

	w.WriteHeader(http.StatusOK)
}
