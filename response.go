package jsonrpc

import "encoding/json"

type Response struct {
	ID      ID              `json:"id,omitempty"`
	JsonRPC string          `json:"jsonRpc,omitempty"`
	Result  json.RawMessage `json:"result,omitempty"`
}

func (r Response) Text() string {
	return string(r.Result)
}

func (r Response) Unmarshal(dst any) error {
	return json.Unmarshal(r.Result, dst)
}
