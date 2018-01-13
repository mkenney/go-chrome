package chrome

/*
ChromiumFlags provides an interface for managing CLI arguments to the Chromium binary.
*/
type ChromiumFlags interface {

	// Get returns the specified flag values
	Get(flag string) ([]interface{}, error)

	// Has checks to see if an flag is present.
	Has(flag string) bool

	// List returns an array of each flag for use in os.StartProcess
	List() []string

	// Set sets a flag's values.
	Set(flag string, values []interface{}) error

	// String implments Stringer. It returns the set parameters formatted to be
	// passed to the command line.
	String() string
}
