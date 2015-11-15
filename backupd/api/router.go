package api

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/raphaelm/backupd/backupd/datastore"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc func(store datastore.DataStore, w http.ResponseWriter, r *http.Request)
}
type Routes []Route

func Router(store datastore.DataStore) *mux.Router {
	r := mux.NewRouter().StrictSlash(true)

	for _, route := range routes {
		// Function scope hacks
		// TODO: Do this in an idiomatic way
		s := store
		routed := route

		r.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(
			http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				routed.HandlerFunc(s, w, r)
			}))
	}
	return r
}

var routes = Routes{
	Route{
		"RemoteAdd",
		"POST",
		"/remotes",
		RemoteAdd,
	},
	Route{
		"RemoteIndex",
		"GET",
		"/remotes",
		RemoteIndex,
	},
	Route{
		"RemoteUpdate",
		"PUT",
		"/remotes/{id}",
		RemoteUpdate,
	},
	Route{
		"RemoteDelete",
		"DELETE",
		"/remotes/{id}",
		RemoteDelete,
	},
	Route{
		"JobAdd",
		"POST",
		"/jobs",
		JobAdd,
	},
	Route{
		"JobIndex",
		"GET",
		"/jobs",
		JobIndex,
	},
	Route{
		"JobUpdate",
		"PUT",
		"/jobs/{id}",
		JobUpdate,
	},
	Route{
		"JobDelete",
		"DELETE",
		"/jobs/{id}",
		JobDelete,
	},
	Route{
		"BackupIndex",
		"GET",
		"/backups",
		BackupIndex,
	},
}
