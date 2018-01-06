package schema

/*
GetDomainsResult represents the result of calls to Schema.getDomains.

https://chromedevtools.github.io/devtools-protocol/tot/Schema/#method-getDomains
*/
type GetDomainsResult struct {
	// List of supported domains.
	Domains []*Domain `json:"domains"`

	// Error information related to executing this method
	Err error `json:"-"`
}
