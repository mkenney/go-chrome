package Schema

/*
Domain is a description of the protocol domain.
*/
type Domain struct {
	// Domain name.
	Name string `json:"name"`

	// Domain version.
	Version string `json:"version"`
}
