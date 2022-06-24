package jsonrpc

type Request struct {
	ID ID `json:"id,omitempty"`

	// JsonRPC is a string specifying the version of the JSON-RPC protocol which
	// must be exactly "2.0".
	JsonRPC string `json:"jsonrpc,omitempty"`

	// Method is the name of the method to be called.
	Method string `json:"method,omitempty"`

	// Params the parameters to be passed to the method.
	Params any `json:"params,omitempty"`
}
