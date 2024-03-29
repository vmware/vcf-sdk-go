// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2023 VMware, Inc.
// SPDX-License-Identifier: BSD-2-Clause

package license_keys

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new license keys API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for license keys API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	AddLicenseKey(params *AddLicenseKeyParams, opts ...ClientOption) (*AddLicenseKeyOK, *AddLicenseKeyCreated, error)

	DeleteLicenseKey(params *DeleteLicenseKeyParams, opts ...ClientOption) (*DeleteLicenseKeyOK, *DeleteLicenseKeyNoContent, error)

	GetDomainLicensingInfo(params *GetDomainLicensingInfoParams, opts ...ClientOption) (*GetDomainLicensingInfoOK, error)

	GetLicenseKey(params *GetLicenseKeyParams, opts ...ClientOption) (*GetLicenseKeyOK, error)

	GetLicenseKeys(params *GetLicenseKeysParams, opts ...ClientOption) (*GetLicenseKeysOK, error)

	GetLicensingInfo(params *GetLicensingInfoParams, opts ...ClientOption) (*GetLicensingInfoOK, error)

	GetSystemLicensingInfo(params *GetSystemLicensingInfoParams, opts ...ClientOption) (*GetSystemLicensingInfoOK, error)

	UpdateDomainLicensingInfo(params *UpdateDomainLicensingInfoParams, opts ...ClientOption) (*UpdateDomainLicensingInfoOK, *UpdateDomainLicensingInfoAccepted, error)

	UpdateSystemLicensingInfo(params *UpdateSystemLicensingInfoParams, opts ...ClientOption) (*UpdateSystemLicensingInfoOK, *UpdateSystemLicensingInfoAccepted, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
AddLicenseKey adds a license key
*/
func (a *Client) AddLicenseKey(params *AddLicenseKeyParams, opts ...ClientOption) (*AddLicenseKeyOK, *AddLicenseKeyCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAddLicenseKeyParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "addLicenseKey",
		Method:             "POST",
		PathPattern:        "/v1/license-keys",
		ProducesMediaTypes: []string{"*/*"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &AddLicenseKeyReader{formats: a.formats},
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
	case *AddLicenseKeyOK:
		return value, nil, nil
	case *AddLicenseKeyCreated:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for license_keys: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
DeleteLicenseKey deletes a license key
*/
func (a *Client) DeleteLicenseKey(params *DeleteLicenseKeyParams, opts ...ClientOption) (*DeleteLicenseKeyOK, *DeleteLicenseKeyNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteLicenseKeyParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "deleteLicenseKey",
		Method:             "DELETE",
		PathPattern:        "/v1/license-keys/{key}",
		ProducesMediaTypes: []string{"*/*"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteLicenseKeyReader{formats: a.formats},
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
	case *DeleteLicenseKeyOK:
		return value, nil, nil
	case *DeleteLicenseKeyNoContent:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for license_keys: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetDomainLicensingInfo gets the licensing information of a domain
*/
func (a *Client) GetDomainLicensingInfo(params *GetDomainLicensingInfoParams, opts ...ClientOption) (*GetDomainLicensingInfoOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetDomainLicensingInfoParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getDomainLicensingInfo",
		Method:             "GET",
		PathPattern:        "/v1/licensing-info/domains/{id}",
		ProducesMediaTypes: []string{"*/*"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetDomainLicensingInfoReader{formats: a.formats},
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
	success, ok := result.(*GetDomainLicensingInfoOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getDomainLicensingInfo: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetLicenseKey gets a license key by key
*/
func (a *Client) GetLicenseKey(params *GetLicenseKeyParams, opts ...ClientOption) (*GetLicenseKeyOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetLicenseKeyParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getLicenseKey",
		Method:             "GET",
		PathPattern:        "/v1/license-keys/{key}",
		ProducesMediaTypes: []string{"*/*"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetLicenseKeyReader{formats: a.formats},
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
	success, ok := result.(*GetLicenseKeyOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getLicenseKey: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetLicenseKeys gets the license keys
*/
func (a *Client) GetLicenseKeys(params *GetLicenseKeysParams, opts ...ClientOption) (*GetLicenseKeysOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetLicenseKeysParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getLicenseKeys",
		Method:             "GET",
		PathPattern:        "/v1/license-keys",
		ProducesMediaTypes: []string{"*/*"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetLicenseKeysReader{formats: a.formats},
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
	success, ok := result.(*GetLicenseKeysOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getLicenseKeys: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetLicensingInfo gets the licensing information
*/
func (a *Client) GetLicensingInfo(params *GetLicensingInfoParams, opts ...ClientOption) (*GetLicensingInfoOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetLicensingInfoParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getLicensingInfo",
		Method:             "GET",
		PathPattern:        "/v1/licensing-info",
		ProducesMediaTypes: []string{"*/*"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetLicensingInfoReader{formats: a.formats},
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
	success, ok := result.(*GetLicensingInfoOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getLicensingInfo: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetSystemLicensingInfo gets the licensing information of system
*/
func (a *Client) GetSystemLicensingInfo(params *GetSystemLicensingInfoParams, opts ...ClientOption) (*GetSystemLicensingInfoOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetSystemLicensingInfoParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getSystemLicensingInfo",
		Method:             "GET",
		PathPattern:        "/v1/licensing-info/system",
		ProducesMediaTypes: []string{"*/*"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetSystemLicensingInfoReader{formats: a.formats},
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
	success, ok := result.(*GetSystemLicensingInfoOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getSystemLicensingInfo: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
UpdateDomainLicensingInfo updates the licensing information of a domain
*/
func (a *Client) UpdateDomainLicensingInfo(params *UpdateDomainLicensingInfoParams, opts ...ClientOption) (*UpdateDomainLicensingInfoOK, *UpdateDomainLicensingInfoAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateDomainLicensingInfoParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "updateDomainLicensingInfo",
		Method:             "PUT",
		PathPattern:        "/v1/licensing-info/domains/{id}",
		ProducesMediaTypes: []string{"*/*"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateDomainLicensingInfoReader{formats: a.formats},
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
	case *UpdateDomainLicensingInfoOK:
		return value, nil, nil
	case *UpdateDomainLicensingInfoAccepted:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for license_keys: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
UpdateSystemLicensingInfo updates the licensing information of system
*/
func (a *Client) UpdateSystemLicensingInfo(params *UpdateSystemLicensingInfoParams, opts ...ClientOption) (*UpdateSystemLicensingInfoOK, *UpdateSystemLicensingInfoAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateSystemLicensingInfoParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "updateSystemLicensingInfo",
		Method:             "PUT",
		PathPattern:        "/v1/licensing-info/system",
		ProducesMediaTypes: []string{"*/*"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateSystemLicensingInfoReader{formats: a.formats},
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
	case *UpdateSystemLicensingInfoOK:
		return value, nil, nil
	case *UpdateSystemLicensingInfoAccepted:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for license_keys: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
