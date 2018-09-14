package things

// Default is a default set of Things
//
// TODO: Replace with some kind of service connection
// this should be provided by a package outside the scope of this project
var Default = Things{
	Thing{Name: "dance", Data: "swing"},
	Thing{Name: "dance", Data: "salsa"},
	Thing{Name: "dance", Data: "foxtrot"},
	Thing{Name: "dance", Data: "latin"},
	Thing{Name: "food", Data: "latin"},
	Thing{Name: "food", Data: "sushi"},
	Thing{Name: "food", Data: "curry"},
	Thing{Name: "food", Data: "matzo"},
}
