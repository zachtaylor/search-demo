package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/zachtaylor/search-demo/things"
	"ztaylor.me/log"
)

// Server returns a http.Handler with provided services injected
func Server(thingService things.Service) http.Handler {
	mux := mux.NewRouter()
	mux.HandleFunc("/api/things", GetThings(thingService)).Methods("GET")
	mux.Handle("/*", http.FileServer(http.Dir("www/dist/")))
	return mux
}

// GetThings is the endpoint for GET /api/things
//
// Load the things, and write them
func GetThings(thingService things.Service) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		things, err := thingService.GetThings()
		if err != nil {
			log.Add("Request", r).Add("Error", err).Error("GET /api/things: failed to thingService.GetThings()")
			w.WriteHeader(500)
		} else if json, err := json.Marshal(things); err != nil {
			log.Add("Request", r).Add("Error", err).Error("GET /api/things: failed to json.Marshal(things)")
			w.WriteHeader(500)
		} else {
			w.Header().Add("Content-Type", "text/json")
			w.Write(json)
		}
	}
}
