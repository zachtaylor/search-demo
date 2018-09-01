package things

// Service provides Things from somewhere
type Service interface {
	// GetThings returns the Things, and error if any
	GetThings() ([]Thing, error)
}
