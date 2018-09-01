package things

// Thing is an item in a list
type Thing struct {
	// Name is the category or subject
	Name string
	// Data is the value
	Data string
}

// Things is array of Thing, which provides Service
type Things []Thing

// GetThings provides ThingService
//
// Returns itself and nil, always
func (things Things) GetThings() ([]Thing, error) {
	return things, nil
}
