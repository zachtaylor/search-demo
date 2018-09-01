package main

import (
	"net/http"

	"github.com/zachtaylor/search-demo/things"
)

var thingService things.Service

func main() {
	// find service providers

	// change thingService to real provider probably
	thingService = DefaultThings

	// inject services into server
	server := Server(thingService)

	// listen and serve
	http.ListenAndServe(":80", server)
}
