package main

import (
	"github.com/zachtaylor/search-demo/things"
)

// DefaultThings is a default set of Things
//
// TODO: Replace with some kind of service connection
// this should be provided by a package outside the scope of this project
var DefaultThings = things.Things{
	things.Thing{Name: "dance", Data: "swing"},
	things.Thing{Name: "dance", Data: "salsa"},
	things.Thing{Name: "dance", Data: "foxtrot"},
	things.Thing{Name: "dance", Data: "latin"},
	things.Thing{Name: "food", Data: "latin"},
	things.Thing{Name: "food", Data: "sushi"},
	things.Thing{Name: "food", Data: "curry"},
	things.Thing{Name: "food", Data: "matzo"},
}
