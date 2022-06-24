package jsonrpc

import "net/http"

// Options consists of additional options for the NewClient() function.
type Options struct {
	// Client is the HTTP client to use to dispatch requests.
	Client *http.Client `json:"client,omitempty"`

	// Before is the function to call before dispatching a request.
	Before func(r Request, hr *http.Request)
}

// NoOptions are the default options.
var NoOptions = Options{
	Client: http.DefaultClient,
}

// defaults sets the default values for Options.
func (opts *Options) defaults() {
	if opts.Client == nil {
		opts.Client = NoOptions.Client
	}
}
