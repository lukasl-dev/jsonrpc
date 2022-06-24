package jsonrpc

import (
	"bytes"
	"encoding/json"
	"github.com/google/uuid"
	"io"
	"net/http"
)

type Client struct {
	endpoint string
	opts     Options
}

// NewClient returns a new JSON-RPC client that operates on the given HTTP
// endpoint.
func NewClient(endpoint string, opts Options) *Client {
	opts.defaults()
	return &Client{endpoint: endpoint, opts: opts}
}

// Call calls the method with the given params and returns its Result.
func (c *Client) Call(method string, params any) (*Response, error) {
	req := Request{
		ID:      ID(uuid.New().String()),
		JsonRPC: "2.0",
		Method:  method,
		Params:  params,
	}
	return c.Request(req)
}

// Request dispatches the Request and returns the Response.
func (c *Client) Request(req Request) (*Response, error) {
	httpReq, err := c.createHTTPRequest(req)
	if err != nil {
		return nil, err
	}

	if c.opts.Before != nil {
		c.opts.Before(req, httpReq)
	}

	httpResp, err := c.opts.Client.Do(httpReq)
	if err != nil {
		return nil, err
	}

	return c.unmarshalResponse(httpResp)
}

func (c *Client) createHTTPRequest(req Request) (*http.Request, error) {
	body, err := c.body(req)
	if err != nil {
		return nil, err
	}
	return http.NewRequest("GET", c.endpoint, body)
}

func (c *Client) body(req Request) (io.Reader, error) {
	b, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	return bytes.NewReader(b), nil
}

func (c *Client) unmarshalResponse(resp *http.Response) (*Response, error) {
	var r Response
	if err := json.NewDecoder(resp.Body).Decode(&r); err != nil {
		return nil, err
	}

	if r.Error != nil {
		return nil, r.Error
	}

	r.hr = resp
	return &r, nil
}
