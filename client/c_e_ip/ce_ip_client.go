// Code generated by go-swagger; DO NOT EDIT.

package c_e_ip

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new c e ip API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for c e ip API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	GetCeipStatus(params *GetCeipStatusParams, opts ...ClientOption) (*GetCeipStatusOK, error)

	UpdateCeipStatus(params *UpdateCeipStatusParams, opts ...ClientOption) (*UpdateCeipStatusOK, *UpdateCeipStatusAccepted, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
GetCeipStatus gets c e IP status

Get CEIP status and instance id
*/
func (a *Client) GetCeipStatus(params *GetCeipStatusParams, opts ...ClientOption) (*GetCeipStatusOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCeipStatusParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getCeipStatus",
		Method:             "GET",
		PathPattern:        "/v1/system/ceip",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetCeipStatusReader{formats: a.formats},
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
	success, ok := result.(*GetCeipStatusOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getCeipStatus: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
UpdateCeipStatus opts in or opt out of c e IP

Opt-in or Opt-out of CEIP
*/
func (a *Client) UpdateCeipStatus(params *UpdateCeipStatusParams, opts ...ClientOption) (*UpdateCeipStatusOK, *UpdateCeipStatusAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateCeipStatusParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "updateCeipStatus",
		Method:             "PATCH",
		PathPattern:        "/v1/system/ceip",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateCeipStatusReader{formats: a.formats},
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
	case *UpdateCeipStatusOK:
		return value, nil, nil
	case *UpdateCeipStatusAccepted:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for c_e_ip: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}