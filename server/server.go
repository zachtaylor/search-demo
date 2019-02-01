package server

import (
	"encoding/json"

	"github.com/zachtaylor/search-demo/things"
	"ztaylor.me/gops"
	"ztaylor.me/gops/http"
	"ztaylor.me/log"
)

// New returns a gops.Handler with provided services injected
func New(things things.Service, assets http.FileSystem) gops.Handler {
	return gops.Mux{
		gops.New(
			gops.RouterSet(
				gops.RouterGET,
				gops.RouterPath("/api/things"),
			),
			GetThings(things),
		),
	}
}

// GetThings is the endpoint for GET /api/things
//
// Load the things, and write them
func GetThings(thingService things.Service) gops.HandlerFunc {
	return func(i gops.In, o gops.Out) {
		things, err := thingService.GetThings()
		if err != nil {
			log.Add("Request", i).Add("Error", err).Error("GET /api/things: failed to thingService.GetThings()")
			o.StatusCode(500)
		} else if json, err := json.Marshal(things); err != nil {
			log.Add("Request", i).Add("Error", err).Error("GET /api/things: failed to json.Marshal(things)")
			o.StatusCode(500)
		} else {
			o.Header("Content-Type", "application/json")
			o.Write(json)
		}
	}
}
