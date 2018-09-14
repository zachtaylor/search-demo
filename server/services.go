package server

import (
	"net/http"

	"github.com/zachtaylor/search-demo/things"
)

var ThingService things.Service

var AssetService http.Handler
