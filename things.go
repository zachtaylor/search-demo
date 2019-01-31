package main

import (
	"github.com/zachtaylor/search-demo/things"
)

// defaultThings is slice of things.Thing, which provides things.Service
type defaultThings []things.Thing

// GetThings provides ThingService
//
// Returns itself and nil, always
func (things defaultThings) GetThings() ([]things.Thing, error) {
	return things, nil
}

// Things is a default set of things.Thing, implementing things.Service
//
// TODO: Replace with some kind of service connection
// this should be provided by a package outside the scope of this project
var Things = defaultThings{
	things.Thing{Name: "dance", Data: "swing"},
	things.Thing{Name: "dance", Data: "salsa"},
	things.Thing{Name: "dance", Data: "foxtrot"},
	things.Thing{Name: "dance", Data: "latin"},
	things.Thing{Name: "food", Data: "latin"},
	things.Thing{Name: "food", Data: "sushi"},
	things.Thing{Name: "food", Data: "curry"},
	things.Thing{Name: "food", Data: "matzo"},
}
