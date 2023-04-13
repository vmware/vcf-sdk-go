// Code generated by go-swagger; DO NOT EDIT.

// Copyright © 2023 VMware, Inc. All Rights Reserved.
// Code licenced under: BSD-2

package resource_warnings

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new resource warnings API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for resource warnings API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	GETResourceWarning(params *GETResourceWarningParams, opts ...ClientOption) (*GETResourceWarningOK, error)

	GETResourceWarnings(params *GETResourceWarningsParams, opts ...ClientOption) (*GETResourceWarningsOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
GETResourceWarning Gets a resource warning by ID
*/
func (a *Client) GETResourceWarning(params *GETResourceWarningParams, opts ...ClientOption) (*GETResourceWarningOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGETResourceWarningParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getResourceWarning",
		Method:             "GET",
		PathPattern:        "/v1/resource-warnings/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GETResourceWarningReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GETResourceWarningOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getResourceWarning: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GETResourceWarnings Gets resource warnings
*/
func (a *Client) GETResourceWarnings(params *GETResourceWarningsParams, opts ...ClientOption) (*GETResourceWarningsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGETResourceWarningsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getResourceWarnings",
		Method:             "GET",
		PathPattern:        "/v1/resource-warnings",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GETResourceWarningsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GETResourceWarningsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getResourceWarnings: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
