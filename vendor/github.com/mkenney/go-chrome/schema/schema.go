package Schema

/*
GetDomainsResult represents the result of calls to Schema.getDomains.
*/
type GetDomainsResult struct {
	// List of supported domains.
	Domains []Domain `json:"domains"`
}

/*
Domain is a description of the protocol domain.
*/
type Domain struct {
	// Domain name.
	Name string `json:"name"`

	// Domain version.
	Version string `json:"version"`
}
