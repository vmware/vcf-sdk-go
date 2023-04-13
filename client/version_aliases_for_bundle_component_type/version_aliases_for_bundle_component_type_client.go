// Code generated by go-swagger; DO NOT EDIT.

// Copyright © 2023 VMware, Inc. All Rights Reserved.
// Code licenced under: BSD-2

package version_aliases_for_bundle_component_type

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new version aliases for bundle component type API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for version aliases for bundle component type API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	DeleteAliasVersionsBySoftwareTypeAndBaseVersion(params *DeleteAliasVersionsBySoftwareTypeAndBaseVersionParams, opts ...ClientOption) (*DeleteAliasVersionsBySoftwareTypeAndBaseVersionOK, *DeleteAliasVersionsBySoftwareTypeAndBaseVersionNoContent, error)

	DeleteVersionAliasBySoftwareType(params *DeleteVersionAliasBySoftwareTypeParams, opts ...ClientOption) (*DeleteVersionAliasBySoftwareTypeOK, *DeleteVersionAliasBySoftwareTypeNoContent, error)

	GETVersionAliasConfiguration(params *GETVersionAliasConfigurationParams, opts ...ClientOption) (*GETVersionAliasConfigurationOK, error)

	UpdateVersionAliasConfiguration(params *UpdateVersionAliasConfigurationParams, opts ...ClientOption) (*UpdateVersionAliasConfigurationOK, error)

	UpdateVersionAliasConfigurations(params *UpdateVersionAliasConfigurationsParams, opts ...ClientOption) (*UpdateVersionAliasConfigurationsOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
DeleteAliasVersionsBySoftwareTypeAndBaseVersion deletes version alias configuration

Delete Version Alias Configuration by bundle component type, version, and aliases.
*/
func (a *Client) DeleteAliasVersionsBySoftwareTypeAndBaseVersion(params *DeleteAliasVersionsBySoftwareTypeAndBaseVersionParams, opts ...ClientOption) (*DeleteAliasVersionsBySoftwareTypeAndBaseVersionOK, *DeleteAliasVersionsBySoftwareTypeAndBaseVersionNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteAliasVersionsBySoftwareTypeAndBaseVersionParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "deleteAliasVersionsBySoftwareTypeAndBaseVersion",
		Method:             "DELETE",
		PathPattern:        "/v1/system/settings/version-aliases/{bundleComponentType}/{version}",
		ProducesMediaTypes: []string{"*/*"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteAliasVersionsBySoftwareTypeAndBaseVersionReader{formats: a.formats},
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
	case *DeleteAliasVersionsBySoftwareTypeAndBaseVersionOK:
		return value, nil, nil
	case *DeleteAliasVersionsBySoftwareTypeAndBaseVersionNoContent:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for version_aliases_for_bundle_component_type: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
DeleteVersionAliasBySoftwareType deletes version alias for a bundle component type

Delete Version Alias for a bundle component type.
*/
func (a *Client) DeleteVersionAliasBySoftwareType(params *DeleteVersionAliasBySoftwareTypeParams, opts ...ClientOption) (*DeleteVersionAliasBySoftwareTypeOK, *DeleteVersionAliasBySoftwareTypeNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteVersionAliasBySoftwareTypeParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "deleteVersionAliasBySoftwareType",
		Method:             "DELETE",
		PathPattern:        "/v1/system/settings/version-aliases/{bundleComponentType}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteVersionAliasBySoftwareTypeReader{formats: a.formats},
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
	case *DeleteVersionAliasBySoftwareTypeOK:
		return value, nil, nil
	case *DeleteVersionAliasBySoftwareTypeNoContent:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for version_aliases_for_bundle_component_type: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GETVersionAliasConfiguration Gets version alias configuration

Get the Version Alias Configuration.
*/
func (a *Client) GETVersionAliasConfiguration(params *GETVersionAliasConfigurationParams, opts ...ClientOption) (*GETVersionAliasConfigurationOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGETVersionAliasConfigurationParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getVersionAliasConfiguration",
		Method:             "GET",
		PathPattern:        "/v1/system/settings/version-aliases",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GETVersionAliasConfigurationReader{formats: a.formats},
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
	success, ok := result.(*GETVersionAliasConfigurationOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getVersionAliasConfiguration: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
UpdateVersionAliasConfiguration updates version alias configuration

Update Version Alias Configuration.
*/
func (a *Client) UpdateVersionAliasConfiguration(params *UpdateVersionAliasConfigurationParams, opts ...ClientOption) (*UpdateVersionAliasConfigurationOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateVersionAliasConfigurationParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "updateVersionAliasConfiguration",
		Method:             "PUT",
		PathPattern:        "/v1/system/settings/version-aliases/{bundleComponentType}/{version}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateVersionAliasConfigurationReader{formats: a.formats},
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
	success, ok := result.(*UpdateVersionAliasConfigurationOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for updateVersionAliasConfiguration: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
UpdateVersionAliasConfigurations updates version alias configurations

Update Version Alias Configurations.
*/
func (a *Client) UpdateVersionAliasConfigurations(params *UpdateVersionAliasConfigurationsParams, opts ...ClientOption) (*UpdateVersionAliasConfigurationsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateVersionAliasConfigurationsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "updateVersionAliasConfigurations",
		Method:             "PUT",
		PathPattern:        "/v1/system/settings/version-aliases",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateVersionAliasConfigurationsReader{formats: a.formats},
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
	success, ok := result.(*UpdateVersionAliasConfigurationsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for updateVersionAliasConfigurations: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
