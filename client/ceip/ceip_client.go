// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2023 VMware, Inc.
// SPDX-License-Identifier: BSD-2-Clause

package ceip

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new ceip API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for ceip API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	GetCEIPStatus(params *GetCEIPStatusParams, opts ...ClientOption) (*GetCEIPStatusOK, error)

	SetCEIPStatus(params *SetCEIPStatusParams, opts ...ClientOption) (*SetCEIPStatusOK, *SetCEIPStatusAccepted, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
GetCEIPStatus retrieves the CEIP status

Get CEIP status and instance id
*/
func (a *Client) GetCEIPStatus(params *GetCEIPStatusParams, opts ...ClientOption) (*GetCEIPStatusOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCEIPStatusParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getCeipStatus",
		Method:             "GET",
		PathPattern:        "/v1/system/ceip",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetCEIPStatusReader{formats: a.formats},
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
	success, ok := result.(*GetCEIPStatusOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getCeipStatus: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
SetCEIPStatus configures CEIP to opt in or opt out

Opt-in or Opt-out of CEIP
*/
func (a *Client) SetCEIPStatus(params *SetCEIPStatusParams, opts ...ClientOption) (*SetCEIPStatusOK, *SetCEIPStatusAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSetCEIPStatusParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "setCeipStatus",
		Method:             "PATCH",
		PathPattern:        "/v1/system/ceip",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &SetCEIPStatusReader{formats: a.formats},
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
	case *SetCEIPStatusOK:
		return value, nil, nil
	case *SetCEIPStatusAccepted:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ceip: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
