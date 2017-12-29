// Code generated by goagen v1.3.0, DO NOT EDIT.
//
// API "cycles": Application Media Types
//
// Command:
// $ goagen
// --design=github.com/rchampourlier/cycles-backend/design
// --out=$(GOPATH)/src/github.com/rchampourlier/cycles-backend
// --version=v1.3.0

package client

import (
	"github.com/goadesign/goa"
	"net/http"
)

// State of the Cycles frontend application (default view)
//
// Identifier: application/cycles.state+json; view=default
type CyclesState struct {
	// JSON state of the application
	State *string `form:"state,omitempty" json:"state,omitempty" xml:"state,omitempty"`
}

// DecodeCyclesState decodes the CyclesState instance encoded in resp body.
func (c *Client) DecodeCyclesState(resp *http.Response) (*CyclesState, error) {
	var decoded CyclesState
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// DecodeErrorResponse decodes the ErrorResponse instance encoded in resp body.
func (c *Client) DecodeErrorResponse(resp *http.Response) (*goa.ErrorResponse, error) {
	var decoded goa.ErrorResponse
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}
