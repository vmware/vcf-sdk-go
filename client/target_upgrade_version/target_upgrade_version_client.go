// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2023 VMware, Inc.
// SPDX-License-Identifier: BSD-2-Clause

package target_upgrade_version

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new target upgrade version API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for target upgrade version API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	GetReleaseByDomain(params *GetReleaseByDomainParams, opts ...ClientOption) (*GetReleaseByDomainOK, error)

	GetReleaseByDomain1(params *GetReleaseByDomain1Params, opts ...ClientOption) (*GetReleaseByDomain1OK, error)

	UpdateReleaseByDomainID(params *UpdateReleaseByDomainIDParams, opts ...ClientOption) (*UpdateReleaseByDomainIDOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
GetReleaseByDomain retrieves a release view for a domain by its ID

Get last selected upgrade version for the domain.
*/
func (a *Client) GetReleaseByDomain(params *GetReleaseByDomainParams, opts ...ClientOption) (*GetReleaseByDomainOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetReleaseByDomainParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getReleaseByDomain",
		Method:             "GET",
		PathPattern:        "/v1/releases/domains/{domainId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetReleaseByDomainReader{formats: a.formats},
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
	success, ok := result.(*GetReleaseByDomainOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getReleaseByDomain: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetReleaseByDomain1 retrieves a release view for all domains

Get last selected upgrade version for WLDs.
*/
func (a *Client) GetReleaseByDomain1(params *GetReleaseByDomain1Params, opts ...ClientOption) (*GetReleaseByDomain1OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetReleaseByDomain1Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "getReleaseByDomain_1",
		Method:             "GET",
		PathPattern:        "/v1/releases/domains",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetReleaseByDomain1Reader{formats: a.formats},
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
	success, ok := result.(*GetReleaseByDomain1OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getReleaseByDomain_1: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
UpdateReleaseByDomainID modifies the target upgrade release for a domain by its ID

Update last selected upgrade version for the domain
*/
func (a *Client) UpdateReleaseByDomainID(params *UpdateReleaseByDomainIDParams, opts ...ClientOption) (*UpdateReleaseByDomainIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateReleaseByDomainIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "updateReleaseByDomainID",
		Method:             "PATCH",
		PathPattern:        "/v1/releases/domains/{domainId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateReleaseByDomainIDReader{formats: a.formats},
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
	success, ok := result.(*UpdateReleaseByDomainIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for updateReleaseByDomainID: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
