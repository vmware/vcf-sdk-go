// Code generated by go-swagger; DO NOT EDIT.

package depot_settings

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new depot settings API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for depot settings API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	GetDepotSettings(params *GetDepotSettingsParams, opts ...ClientOption) (*GetDepotSettingsOK, error)

	UpdateDepotSettings(params *UpdateDepotSettingsParams, opts ...ClientOption) (*UpdateDepotSettingsOK, *UpdateDepotSettingsAccepted, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
GetDepotSettings gets depot settings

Get the Depot Settings, In a fresh setup, this would be empty
*/
func (a *Client) GetDepotSettings(params *GetDepotSettingsParams, opts ...ClientOption) (*GetDepotSettingsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetDepotSettingsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getDepotSettings",
		Method:             "GET",
		PathPattern:        "/v1/system/settings/depot",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetDepotSettingsReader{formats: a.formats},
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
	success, ok := result.(*GetDepotSettingsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getDepotSettings: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
UpdateDepotSettings updates depot settings

Update Depot Settings. Depot Settings can be updated with VMware Depot Account
*/
func (a *Client) UpdateDepotSettings(params *UpdateDepotSettingsParams, opts ...ClientOption) (*UpdateDepotSettingsOK, *UpdateDepotSettingsAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateDepotSettingsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "updateDepotSettings",
		Method:             "PUT",
		PathPattern:        "/v1/system/settings/depot",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateDepotSettingsReader{formats: a.formats},
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
	case *UpdateDepotSettingsOK:
		return value, nil, nil
	case *UpdateDepotSettingsAccepted:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for depot_settings: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
