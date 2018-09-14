package main

import (
	"net/http"

	"github.com/zachtaylor/search-demo/server"
	"github.com/zachtaylor/search-demo/things"
)

func main() {
	// inject services into server
	// change thingService to real provider probably
	server.ThingService = things.Default
	server.AssetService = http.FileServer(http.Dir("www/dist/www/"))

	server := server.New()

	server.ListenAndServe(":80")
}
