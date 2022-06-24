package jsonrpc

type Response struct {
	// ID is the unique ID of the request that caused this response.
	ID ID `json:"id,omitempty"`

	// JsonRPC is a string specifying the version of the JSON-RPC protocol which
	// must be exactly "2.0".
	JsonRPC string `json:"jsonrpc,omitempty"`

	// Result is the result of the request.
	Result *Result `json:"result,omitempty"`

	// Error is the error that occurred during the request.
	Error *Error `json:"error,omitempty"`
}
