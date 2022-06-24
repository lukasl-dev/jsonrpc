package jsonrpc

type Error struct {
	Code    ErrorCode `json:"code,omitempty"`
	Message string    `json:"message,omitempty"`
}

type ErrorCode int

const (
	ParseError     ErrorCode = -32700
	InvalidRequest ErrorCode = -32600
	MethodNotFound ErrorCode = -32601
	InvalidParams  ErrorCode = -32602
	InternalError  ErrorCode = -32603
)

// IsServerError returns whether the error is considered a server error.
func (code ErrorCode) IsServerError() bool {
	return code >= -32000 && code <= -32099
}
