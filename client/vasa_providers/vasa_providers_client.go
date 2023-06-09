// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2023 VMware, Inc.
// SPDX-License-Identifier: BSD-2-Clause

package vasa_providers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new vasa providers API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for vasa providers API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	AddStorageContainersToVasaProvider(params *AddStorageContainersToVasaProviderParams, opts ...ClientOption) (*AddStorageContainersToVasaProviderCreated, error)

	AddUsersToVasaProvider(params *AddUsersToVasaProviderParams, opts ...ClientOption) (*AddUsersToVasaProviderCreated, error)

	AddVasaProvider(params *AddVasaProviderParams, opts ...ClientOption) (*AddVasaProviderCreated, error)

	DeleteStorageContainerOfVasaProvider(params *DeleteStorageContainerOfVasaProviderParams, opts ...ClientOption) (*DeleteStorageContainerOfVasaProviderNoContent, error)

	DeleteVasaProvider(params *DeleteVasaProviderParams, opts ...ClientOption) (*DeleteVasaProviderNoContent, error)

	GetStorageContainersOfVasaProvider(params *GetStorageContainersOfVasaProviderParams, opts ...ClientOption) (*GetStorageContainersOfVasaProviderOK, error)

	GetUsersOfVasaProvider(params *GetUsersOfVasaProviderParams, opts ...ClientOption) (*GetUsersOfVasaProviderOK, error)

	GetValidationOfVasaProvider(params *GetValidationOfVasaProviderParams, opts ...ClientOption) (*GetValidationOfVasaProviderOK, error)

	GetVasaProvider(params *GetVasaProviderParams, opts ...ClientOption) (*GetVasaProviderOK, error)

	GetVasaProviders(params *GetVasaProvidersParams, opts ...ClientOption) (*GetVasaProvidersOK, error)

	UpdateStorageContainerOfVasaProvider(params *UpdateStorageContainerOfVasaProviderParams, opts ...ClientOption) (*UpdateStorageContainerOfVasaProviderOK, error)

	UpdateUserOfVasaProvider(params *UpdateUserOfVasaProviderParams, opts ...ClientOption) (*UpdateUserOfVasaProviderOK, error)

	UpdateVasaProvider(params *UpdateVasaProviderParams, opts ...ClientOption) (*UpdateVasaProviderOK, error)

	ValidateVasaProvider(params *ValidateVasaProviderParams, opts ...ClientOption) (*ValidateVasaProviderOK, *ValidateVasaProviderAccepted, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
AddStorageContainersToVasaProvider adds the storage containers to a v a s a provider
*/
func (a *Client) AddStorageContainersToVasaProvider(params *AddStorageContainersToVasaProviderParams, opts ...ClientOption) (*AddStorageContainersToVasaProviderCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAddStorageContainersToVasaProviderParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "addStorageContainersToVasaProvider",
		Method:             "POST",
		PathPattern:        "/v1/vasa-providers/{id}/storage-containers",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &AddStorageContainersToVasaProviderReader{formats: a.formats},
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
	success, ok := result.(*AddStorageContainersToVasaProviderCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for addStorageContainersToVasaProvider: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
AddUsersToVasaProvider adds the users to a v a s a provider
*/
func (a *Client) AddUsersToVasaProvider(params *AddUsersToVasaProviderParams, opts ...ClientOption) (*AddUsersToVasaProviderCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAddUsersToVasaProviderParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "addUsersToVasaProvider",
		Method:             "POST",
		PathPattern:        "/v1/vasa-providers/{id}/users",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &AddUsersToVasaProviderReader{formats: a.formats},
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
	success, ok := result.(*AddUsersToVasaProviderCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for addUsersToVasaProvider: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
AddVasaProvider adds a v a s a provider
*/
func (a *Client) AddVasaProvider(params *AddVasaProviderParams, opts ...ClientOption) (*AddVasaProviderCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAddVasaProviderParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "addVasaProvider",
		Method:             "POST",
		PathPattern:        "/v1/vasa-providers",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &AddVasaProviderReader{formats: a.formats},
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
	success, ok := result.(*AddVasaProviderCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for addVasaProvider: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
DeleteStorageContainerOfVasaProvider deletes a storage container of a v a s a provider
*/
func (a *Client) DeleteStorageContainerOfVasaProvider(params *DeleteStorageContainerOfVasaProviderParams, opts ...ClientOption) (*DeleteStorageContainerOfVasaProviderNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteStorageContainerOfVasaProviderParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "deleteStorageContainerOfVasaProvider",
		Method:             "DELETE",
		PathPattern:        "/v1/vasa-providers/{id}/storage-containers/{storageContainerId}",
		ProducesMediaTypes: []string{"*/*"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteStorageContainerOfVasaProviderReader{formats: a.formats},
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
	success, ok := result.(*DeleteStorageContainerOfVasaProviderNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteStorageContainerOfVasaProvider: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
DeleteVasaProvider deletes a v a s a provider
*/
func (a *Client) DeleteVasaProvider(params *DeleteVasaProviderParams, opts ...ClientOption) (*DeleteVasaProviderNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteVasaProviderParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "deleteVasaProvider",
		Method:             "DELETE",
		PathPattern:        "/v1/vasa-providers/{id}",
		ProducesMediaTypes: []string{"*/*"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteVasaProviderReader{formats: a.formats},
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
	success, ok := result.(*DeleteVasaProviderNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteVasaProvider: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetStorageContainersOfVasaProvider gets the storage containers of a v a s a provider
*/
func (a *Client) GetStorageContainersOfVasaProvider(params *GetStorageContainersOfVasaProviderParams, opts ...ClientOption) (*GetStorageContainersOfVasaProviderOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetStorageContainersOfVasaProviderParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getStorageContainersOfVasaProvider",
		Method:             "GET",
		PathPattern:        "/v1/vasa-providers/{id}/storage-containers",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetStorageContainersOfVasaProviderReader{formats: a.formats},
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
	success, ok := result.(*GetStorageContainersOfVasaProviderOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getStorageContainersOfVasaProvider: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetUsersOfVasaProvider gets the users of a v a s a provider
*/
func (a *Client) GetUsersOfVasaProvider(params *GetUsersOfVasaProviderParams, opts ...ClientOption) (*GetUsersOfVasaProviderOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetUsersOfVasaProviderParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getUsersOfVasaProvider",
		Method:             "GET",
		PathPattern:        "/v1/vasa-providers/{id}/users",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetUsersOfVasaProviderReader{formats: a.formats},
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
	success, ok := result.(*GetUsersOfVasaProviderOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getUsersOfVasaProvider: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetValidationOfVasaProvider gets the status of the validation of the v a s a provider
*/
func (a *Client) GetValidationOfVasaProvider(params *GetValidationOfVasaProviderParams, opts ...ClientOption) (*GetValidationOfVasaProviderOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetValidationOfVasaProviderParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getValidationOfVasaProvider",
		Method:             "GET",
		PathPattern:        "/v1/vasa-providers/validations/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetValidationOfVasaProviderReader{formats: a.formats},
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
	success, ok := result.(*GetValidationOfVasaProviderOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getValidationOfVasaProvider: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetVasaProvider gets a v a s a provider
*/
func (a *Client) GetVasaProvider(params *GetVasaProviderParams, opts ...ClientOption) (*GetVasaProviderOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetVasaProviderParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getVasaProvider",
		Method:             "GET",
		PathPattern:        "/v1/vasa-providers/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetVasaProviderReader{formats: a.formats},
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
	success, ok := result.(*GetVasaProviderOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getVasaProvider: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetVasaProviders gets the v a s a providers
*/
func (a *Client) GetVasaProviders(params *GetVasaProvidersParams, opts ...ClientOption) (*GetVasaProvidersOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetVasaProvidersParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getVasaProviders",
		Method:             "GET",
		PathPattern:        "/v1/vasa-providers",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetVasaProvidersReader{formats: a.formats},
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
	success, ok := result.(*GetVasaProvidersOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getVasaProviders: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
UpdateStorageContainerOfVasaProvider updates the storage container of a v a s a provider
*/
func (a *Client) UpdateStorageContainerOfVasaProvider(params *UpdateStorageContainerOfVasaProviderParams, opts ...ClientOption) (*UpdateStorageContainerOfVasaProviderOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateStorageContainerOfVasaProviderParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "updateStorageContainerOfVasaProvider",
		Method:             "PATCH",
		PathPattern:        "/v1/vasa-providers/{id}/storage-containers/{storageContainerId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateStorageContainerOfVasaProviderReader{formats: a.formats},
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
	success, ok := result.(*UpdateStorageContainerOfVasaProviderOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for updateStorageContainerOfVasaProvider: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
UpdateUserOfVasaProvider updates the user of a v a s a provider
*/
func (a *Client) UpdateUserOfVasaProvider(params *UpdateUserOfVasaProviderParams, opts ...ClientOption) (*UpdateUserOfVasaProviderOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateUserOfVasaProviderParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "updateUserOfVasaProvider",
		Method:             "PATCH",
		PathPattern:        "/v1/vasa-providers/{id}/users/{userId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateUserOfVasaProviderReader{formats: a.formats},
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
	success, ok := result.(*UpdateUserOfVasaProviderOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for updateUserOfVasaProvider: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
UpdateVasaProvider updates a v a s a provider
*/
func (a *Client) UpdateVasaProvider(params *UpdateVasaProviderParams, opts ...ClientOption) (*UpdateVasaProviderOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateVasaProviderParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "updateVasaProvider",
		Method:             "PATCH",
		PathPattern:        "/v1/vasa-providers/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateVasaProviderReader{formats: a.formats},
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
	success, ok := result.(*UpdateVasaProviderOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for updateVasaProvider: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
ValidateVasaProvider validates vasa provider input specification
*/
func (a *Client) ValidateVasaProvider(params *ValidateVasaProviderParams, opts ...ClientOption) (*ValidateVasaProviderOK, *ValidateVasaProviderAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewValidateVasaProviderParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "validateVasaProvider",
		Method:             "POST",
		PathPattern:        "/v1/vasa-providers/validations",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ValidateVasaProviderReader{formats: a.formats},
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
	case *ValidateVasaProviderOK:
		return value, nil, nil
	case *ValidateVasaProviderAccepted:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for vasa_providers: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
