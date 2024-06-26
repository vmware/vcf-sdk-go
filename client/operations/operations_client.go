// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2023 VMware, Inc.
// SPDX-License-Identifier: BSD-2-Clause

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new operations API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for operations API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	ConnectVROPSWithDomain(params *ConnectVROPSWithDomainParams, opts ...ClientOption) (*ConnectVROPSWithDomainOK, *ConnectVROPSWithDomainAccepted, error)

	GetVROPSIntegratedDomains(params *GetVROPSIntegratedDomainsParams, opts ...ClientOption) (*GetVROPSIntegratedDomainsOK, error)

	GetVropses(params *GetVropsesParams, opts ...ClientOption) (*GetVropsesOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
ConnectVROPSWithDomain connects or disconnect a domain with v mware aria operations

Connects/disconnects a workload domains with VMware Aria Operations
*/
func (a *Client) ConnectVROPSWithDomain(params *ConnectVROPSWithDomainParams, opts ...ClientOption) (*ConnectVROPSWithDomainOK, *ConnectVROPSWithDomainAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewConnectVROPSWithDomainParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "connectVropsWithDomain",
		Method:             "PUT",
		PathPattern:        "/v1/vrops/domains",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ConnectVROPSWithDomainReader{formats: a.formats},
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
	case *ConnectVROPSWithDomainOK:
		return value, nil, nil
	case *ConnectVROPSWithDomainAccepted:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for operations: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetVROPSIntegratedDomains retrieves a list of v mware aria operations integration status for domains

Retrieves the existing  domains and their connection status with VMware Aria Operations.
*/
func (a *Client) GetVROPSIntegratedDomains(params *GetVROPSIntegratedDomainsParams, opts ...ClientOption) (*GetVROPSIntegratedDomainsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetVROPSIntegratedDomainsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getVropsIntegratedDomains",
		Method:             "GET",
		PathPattern:        "/v1/vrops/domains",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetVROPSIntegratedDomainsReader{formats: a.formats},
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
	success, ok := result.(*GetVROPSIntegratedDomainsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getVropsIntegratedDomains: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetVropses retrieves a list of v mware aria operations instances
*/
func (a *Client) GetVropses(params *GetVropsesParams, opts ...ClientOption) (*GetVropsesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetVropsesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getVropses",
		Method:             "GET",
		PathPattern:        "/v1/vropses",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetVropsesReader{formats: a.formats},
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
	success, ok := result.(*GetVropsesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getVropses: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
