// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2023 VMware, Inc.
// SPDX-License-Identifier: BSD-2-Clause

package trusted_certificates

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new trusted certificates API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for trusted certificates API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	AddTrustedCertificate(params *AddTrustedCertificateParams, opts ...ClientOption) (*AddTrustedCertificateOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
AddTrustedCertificate adds trusted certificate to the SDDC manager

Add trusted certificate to the SDDC manager
*/
func (a *Client) AddTrustedCertificate(params *AddTrustedCertificateParams, opts ...ClientOption) (*AddTrustedCertificateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAddTrustedCertificateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "addTrustedCertificate",
		Method:             "POST",
		PathPattern:        "/v1/sddc-manager/trusted-certificates",
		ProducesMediaTypes: []string{"*/*"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &AddTrustedCertificateReader{formats: a.formats},
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
	success, ok := result.(*AddTrustedCertificateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for addTrustedCertificate: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
