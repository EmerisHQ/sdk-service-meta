// Code generated by goa v3.7.5, DO NOT EDIT.
//
// sdk-utilities client HTTP transport
//
// Command:
// $ goa gen github.com/emerishq/sdk-service-meta

package client

import (
	"net/http"

	goahttp "goa.design/goa/v3/http"
)

// Client lists the sdk-utilities service endpoint HTTP clients.
type Client struct {
	// RestoreResponseBody controls whether the response bodies are reset after
	// decoding so they can be read again.
	RestoreResponseBody bool

	scheme  string
	host    string
	encoder func(*http.Request) goahttp.Encoder
	decoder func(*http.Response) goahttp.Decoder
}

// NewClient instantiates HTTP clients for all the sdk-utilities service
// servers.
func NewClient(
	scheme string,
	host string,
	doer goahttp.Doer,
	enc func(*http.Request) goahttp.Encoder,
	dec func(*http.Response) goahttp.Decoder,
	restoreBody bool,
) *Client {
	return &Client{
		RestoreResponseBody: restoreBody,
		scheme:              scheme,
		host:                host,
		decoder:             dec,
		encoder:             enc,
	}
}
