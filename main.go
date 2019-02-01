package main

import (
	"github.com/zachtaylor/search-demo/server"
	"ztaylor.me/gops"
	"ztaylor.me/gops/http"
)

// Plugin is to be imported by GoPS
var Plugin = gops.New(
	gops.RouterDomain("search-demo.ztaylor.me"),
	handler,
)

// inject services into server
var handler = server.New(
	Things,
	http.Dir("/srv/www/search-demo.ztaylor.me/"),
)

func main() {
}
