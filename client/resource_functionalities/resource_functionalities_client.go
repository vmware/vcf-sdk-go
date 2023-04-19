// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2023 VMware, Inc.
// SPDX-License-Identifier: BSD-2-Clause

package resource_functionalities

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new resource functionalities API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for resource functionalities API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	GETResourceFunctionalities(params *GETResourceFunctionalitiesParams, opts ...ClientOption) (*GETResourceFunctionalitiesOK, error)

	GETResourcesFunctionalitiesAllowedGlobal(params *GETResourcesFunctionalitiesAllowedGlobalParams, opts ...ClientOption) (*GETResourcesFunctionalitiesAllowedGlobalOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
GETResourceFunctionalities Gets resource functionalities
*/
func (a *Client) GETResourceFunctionalities(params *GETResourceFunctionalitiesParams, opts ...ClientOption) (*GETResourceFunctionalitiesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGETResourceFunctionalitiesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getResourceFunctionalities",
		Method:             "GET",
		PathPattern:        "/v1/resource-functionalities",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GETResourceFunctionalitiesReader{formats: a.formats},
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
	success, ok := result.(*GETResourceFunctionalitiesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getResourceFunctionalities: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GETResourcesFunctionalitiesAllowedGlobal Gets resource functionalities allowed global configuration
*/
func (a *Client) GETResourcesFunctionalitiesAllowedGlobal(params *GETResourcesFunctionalitiesAllowedGlobalParams, opts ...ClientOption) (*GETResourcesFunctionalitiesAllowedGlobalOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGETResourcesFunctionalitiesAllowedGlobalParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getResourcesFunctionalitiesAllowedGlobal",
		Method:             "GET",
		PathPattern:        "/v1/resource-functionalities/global",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GETResourcesFunctionalitiesAllowedGlobalReader{formats: a.formats},
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
	success, ok := result.(*GETResourcesFunctionalitiesAllowedGlobalOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getResourcesFunctionalitiesAllowedGlobal: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
