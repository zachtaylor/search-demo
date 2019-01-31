package things

// Thing is an item in a list
type Thing struct {
	// Name is the category or subject
	Name string
	// Data is the value
	Data string
}

// Service provides Things from somewhere
type Service interface {
	// GetThings returns the Things, and error if any
	GetThings() ([]Thing, error)
}
