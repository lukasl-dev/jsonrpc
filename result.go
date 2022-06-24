package jsonrpc

import "encoding/json"

type Result struct {
	json.RawMessage
}

// Text returns the result as a string.
func (r Result) Text() string {
	return string(r.RawMessage)
}

// Unmarshal unmarshalls the result into the given value.
func (r Result) Unmarshal(dst any) error {
	return json.Unmarshal(r.RawMessage, dst)
}
