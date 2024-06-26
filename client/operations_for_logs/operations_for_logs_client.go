// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2023 VMware, Inc.
// SPDX-License-Identifier: BSD-2-Clause

package operations_for_logs

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new operations for logs API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for operations for logs API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	ConnectVrliWithDomain(params *ConnectVrliWithDomainParams, opts ...ClientOption) (*ConnectVrliWithDomainOK, *ConnectVrliWithDomainAccepted, error)

	GetVrliIntegratedDomains(params *GetVrliIntegratedDomainsParams, opts ...ClientOption) (*GetVrliIntegratedDomainsOK, error)

	GetVrlis(params *GetVrlisParams, opts ...ClientOption) (*GetVrlisOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
ConnectVrliWithDomain connects or disconnect a domain with v mware aria operations for logs
*/
func (a *Client) ConnectVrliWithDomain(params *ConnectVrliWithDomainParams, opts ...ClientOption) (*ConnectVrliWithDomainOK, *ConnectVrliWithDomainAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewConnectVrliWithDomainParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "connectVrliWithDomain",
		Method:             "PUT",
		PathPattern:        "/v1/vrli/domains",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ConnectVrliWithDomainReader{formats: a.formats},
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
	case *ConnectVrliWithDomainOK:
		return value, nil, nil
	case *ConnectVrliWithDomainAccepted:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for operations_for_logs: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetVrliIntegratedDomains retrieves a list of v mware aria operations for logs integration status for domains
*/
func (a *Client) GetVrliIntegratedDomains(params *GetVrliIntegratedDomainsParams, opts ...ClientOption) (*GetVrliIntegratedDomainsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetVrliIntegratedDomainsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getVrliIntegratedDomains",
		Method:             "GET",
		PathPattern:        "/v1/vrli/domains",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetVrliIntegratedDomainsReader{formats: a.formats},
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
	success, ok := result.(*GetVrliIntegratedDomainsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getVrliIntegratedDomains: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetVrlis retrieves a list of v mware aria operations for logs instances
*/
func (a *Client) GetVrlis(params *GetVrlisParams, opts ...ClientOption) (*GetVrlisOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetVrlisParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getVrlis",
		Method:             "GET",
		PathPattern:        "/v1/vrlis",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetVrlisReader{formats: a.formats},
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
	success, ok := result.(*GetVrlisOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getVrlis: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
