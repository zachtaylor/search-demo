package server

import (
	"encoding/json"
	"net/http"

	"github.com/zachtaylor/search-demo/things"
	"ztaylor.me/http/mux"
	"ztaylor.me/log"
)

// New returns a *mux.Mux with provided services injected
func New() *mux.Mux {
	server := mux.NewMux()
	server.Map(mux.Matches(mux.MatcherGET, mux.MatcherLit("/api/things")), GetThings(ThingService))
	server.MapRgx("/*", AssetService)
	return server
}

// GetThings is the endpoint for GET /api/things
//
// Load the things, and write them
func GetThings(thingService things.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		things, err := thingService.GetThings()
		if err != nil {
			log.Add("Request", r).Add("Error", err).Error("GET /api/things: failed to thingService.GetThings()")
			w.WriteHeader(500)
		} else if json, err := json.Marshal(things); err != nil {
			log.Add("Request", r).Add("Error", err).Error("GET /api/things: failed to json.Marshal(things)")
			w.WriteHeader(500)
		} else {
			w.Header().Add("Content-Type", "application/json")
			w.Write(json)
		}
	}
}
