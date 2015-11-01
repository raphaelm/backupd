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
		r.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(
			http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				route.HandlerFunc(store, w, r)
			}))
	}
	return r
}

var routes = Routes{
	Route{
		"RemoteIndex",
		"GET",
		"/remotes",
		RemoteIndex,
	},
}
