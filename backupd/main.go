package main

import (
	"log"
	"net/http"

	"github.com/raphaelm/backupd/backupd/api"
	"github.com/raphaelm/backupd/backupd/datastore"
)

func main() {
	store := datastore.MockStore()
	router := api.Router(store)
	log.Fatal(http.ListenAndServe(":8080", router))
}
