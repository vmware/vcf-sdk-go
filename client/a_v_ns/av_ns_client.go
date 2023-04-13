// Code generated by go-swagger; DO NOT EDIT.

// Copyright © 2023 VMware, Inc. All Rights Reserved.
// Code licenced under: BSD-2

package a_v_ns

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new a v ns API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for a v ns API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CreateAVNS(params *CreateAVNSParams, opts ...ClientOption) (*CreateAVNSOK, *CreateAVNSAccepted, error)

	GETAllAVNS(params *GETAllAVNSParams, opts ...ClientOption) (*GETAllAVNSOK, error)

	ValidateAVNS(params *ValidateAVNSParams, opts ...ClientOption) (*ValidateAVNSOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
CreateAVNS creates a v ns
*/
func (a *Client) CreateAVNS(params *CreateAVNSParams, opts ...ClientOption) (*CreateAVNSOK, *CreateAVNSAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateAVNSParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "createAvns",
		Method:             "POST",
		PathPattern:        "/v1/avns",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateAVNSReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *CreateAVNSOK:
		return value, nil, nil
	case *CreateAVNSAccepted:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for a_v_ns: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GETAllAVNS fetches all a v ns

Returns all matching AVNs
*/
func (a *Client) GETAllAVNS(params *GETAllAVNSParams, opts ...ClientOption) (*GETAllAVNSOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGETAllAVNSParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getAllAvns",
		Method:             "GET",
		PathPattern:        "/v1/avns",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GETAllAVNSReader{formats: a.formats},
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
	success, ok := result.(*GETAllAVNSOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getAllAvns: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
ValidateAVNS validates a v n creation spec

Returns Validation report
*/
func (a *Client) ValidateAVNS(params *ValidateAVNSParams, opts ...ClientOption) (*ValidateAVNSOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewValidateAVNSParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "validateAvns",
		Method:             "POST",
		PathPattern:        "/v1/avns/validations",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ValidateAVNSReader{formats: a.formats},
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
	success, ok := result.(*ValidateAVNSOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for validateAvns: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
