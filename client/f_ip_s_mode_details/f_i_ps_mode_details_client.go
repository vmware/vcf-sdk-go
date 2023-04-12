// Code generated by go-swagger; DO NOT EDIT.

package f_ip_s_mode_details

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new f ip s mode details API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for f ip s mode details API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	GetSecurityConfig(params *GetSecurityConfigParams, opts ...ClientOption) (*GetSecurityConfigOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
GetSecurityConfig gets v c f security configuration

Retrieve VCF security FIPS mode.
*/
func (a *Client) GetSecurityConfig(params *GetSecurityConfigParams, opts ...ClientOption) (*GetSecurityConfigOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetSecurityConfigParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getSecurityConfig",
		Method:             "GET",
		PathPattern:        "/v1/system/security/fips",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetSecurityConfigReader{formats: a.formats},
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
	success, ok := result.(*GetSecurityConfigOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getSecurityConfig: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}