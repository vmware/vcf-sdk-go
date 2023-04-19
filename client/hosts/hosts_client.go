// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2023 VMware, Inc.
// SPDX-License-Identifier: BSD-2-Clause

package hosts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new hosts API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for hosts API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	AssignTagsToExistingHost(params *AssignTagsToExistingHostParams, opts ...ClientOption) (*AssignTagsToExistingHostOK, error)

	AssignableTagsToHost(params *AssignableTagsToHostParams, opts ...ClientOption) (*AssignableTagsToHostOK, error)

	CommissionHosts(params *CommissionHostsParams, opts ...ClientOption) (*CommissionHostsOK, *CommissionHostsAccepted, error)

	DecommissionHosts(params *DecommissionHostsParams, opts ...ClientOption) (*DecommissionHostsOK, *DecommissionHostsAccepted, error)

	GETCriteria(params *GETCriteriaParams, opts ...ClientOption) (*GETCriteriaOK, error)

	GETCriterion(params *GETCriterionParams, opts ...ClientOption) (*GETCriterionOK, error)

	GETHost(params *GETHostParams, opts ...ClientOption) (*GETHostOK, error)

	GETHostQueryResponse1(params *GETHostQueryResponse1Params, opts ...ClientOption) (*GETHostQueryResponse1OK, error)

	GETHostTagManagerURL(params *GETHostTagManagerURLParams, opts ...ClientOption) (*GETHostTagManagerURLOK, error)

	GETHosts(params *GETHostsParams, opts ...ClientOption) (*GETHostsOK, error)

	GETTagsAssignedToHost(params *GETTagsAssignedToHostParams, opts ...ClientOption) (*GETTagsAssignedToHostOK, error)

	GETTagsAssignedToHosts(params *GETTagsAssignedToHostsParams, opts ...ClientOption) (*GETTagsAssignedToHostsOK, error)

	GETValidationForCommissionHosts(params *GETValidationForCommissionHostsParams, opts ...ClientOption) (*GETValidationForCommissionHostsOK, error)

	POSTQuery(params *POSTQueryParams, opts ...ClientOption) (*POSTQueryOK, error)

	RemoveTagsFromHost(params *RemoveTagsFromHostParams, opts ...ClientOption) (*RemoveTagsFromHostOK, error)

	ValidateHostsOperations(params *ValidateHostsOperationsParams, opts ...ClientOption) (*ValidateHostsOperationsOK, *ValidateHostsOperationsAccepted, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
AssignTagsToExistingHost assigns tags to host
*/
func (a *Client) AssignTagsToExistingHost(params *AssignTagsToExistingHostParams, opts ...ClientOption) (*AssignTagsToExistingHostOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAssignTagsToExistingHostParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "assignTagsToExistingHost",
		Method:             "PUT",
		PathPattern:        "/v1/hosts/{id}/tags",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &AssignTagsToExistingHostReader{formats: a.formats},
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
	success, ok := result.(*AssignTagsToExistingHostOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for assignTagsToExistingHost: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
AssignableTagsToHost gets assignable tags to host
*/
func (a *Client) AssignableTagsToHost(params *AssignableTagsToHostParams, opts ...ClientOption) (*AssignableTagsToHostOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAssignableTagsToHostParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "assignableTagsToHost",
		Method:             "GET",
		PathPattern:        "/v1/hosts/{id}/tags/assignable-tags",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &AssignableTagsToHostReader{formats: a.formats},
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
	success, ok := result.(*AssignableTagsToHostOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for assignableTagsToHost: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
CommissionHosts commissions the hosts
*/
func (a *Client) CommissionHosts(params *CommissionHostsParams, opts ...ClientOption) (*CommissionHostsOK, *CommissionHostsAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCommissionHostsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "commissionHosts",
		Method:             "POST",
		PathPattern:        "/v1/hosts",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CommissionHostsReader{formats: a.formats},
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
	case *CommissionHostsOK:
		return value, nil, nil
	case *CommissionHostsAccepted:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for hosts: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
DecommissionHosts decommissions the hosts
*/
func (a *Client) DecommissionHosts(params *DecommissionHostsParams, opts ...ClientOption) (*DecommissionHostsOK, *DecommissionHostsAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDecommissionHostsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "decommissionHosts",
		Method:             "DELETE",
		PathPattern:        "/v1/hosts",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DecommissionHostsReader{formats: a.formats},
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
	case *DecommissionHostsOK:
		return value, nil, nil
	case *DecommissionHostsAccepted:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for hosts: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GETCriteria Gets all criteria
*/
func (a *Client) GETCriteria(params *GETCriteriaParams, opts ...ClientOption) (*GETCriteriaOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGETCriteriaParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getCriteria",
		Method:             "GET",
		PathPattern:        "/v1/hosts/criteria",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GETCriteriaReader{formats: a.formats},
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
	success, ok := result.(*GETCriteriaOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getCriteria: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GETCriterion Gets a criterion
*/
func (a *Client) GETCriterion(params *GETCriterionParams, opts ...ClientOption) (*GETCriterionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGETCriterionParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getCriterion",
		Method:             "GET",
		PathPattern:        "/v1/hosts/criteria/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GETCriterionReader{formats: a.formats},
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
	success, ok := result.(*GETCriterionOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getCriterion: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GETHost Gets a host
*/
func (a *Client) GETHost(params *GETHostParams, opts ...ClientOption) (*GETHostOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGETHostParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getHost",
		Method:             "GET",
		PathPattern:        "/v1/hosts/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GETHostReader{formats: a.formats},
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
	success, ok := result.(*GETHostOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getHost: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GETHostQueryResponse1 Gets query response
*/
func (a *Client) GETHostQueryResponse1(params *GETHostQueryResponse1Params, opts ...ClientOption) (*GETHostQueryResponse1OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGETHostQueryResponse1Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "getHostQueryResponse_1",
		Method:             "GET",
		PathPattern:        "/v1/hosts/queries/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GETHostQueryResponse1Reader{formats: a.formats},
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
	success, ok := result.(*GETHostQueryResponse1OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getHostQueryResponse_1: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GETHostTagManagerURL Gets host tag manager Url
*/
func (a *Client) GETHostTagManagerURL(params *GETHostTagManagerURLParams, opts ...ClientOption) (*GETHostTagManagerURLOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGETHostTagManagerURLParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getHostTagManagerUrl",
		Method:             "GET",
		PathPattern:        "/v1/hosts/{id}/tags/tag-manager",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GETHostTagManagerURLReader{formats: a.formats},
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
	success, ok := result.(*GETHostTagManagerURLOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getHostTagManagerUrl: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GETHosts Gets the hosts
*/
func (a *Client) GETHosts(params *GETHostsParams, opts ...ClientOption) (*GETHostsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGETHostsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getHosts",
		Method:             "GET",
		PathPattern:        "/v1/hosts",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GETHostsReader{formats: a.formats},
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
	success, ok := result.(*GETHostsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getHosts: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GETTagsAssignedToHost Gets tags assigned to host
*/
func (a *Client) GETTagsAssignedToHost(params *GETTagsAssignedToHostParams, opts ...ClientOption) (*GETTagsAssignedToHostOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGETTagsAssignedToHostParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getTagsAssignedToHost",
		Method:             "GET",
		PathPattern:        "/v1/hosts/{id}/tags",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GETTagsAssignedToHostReader{formats: a.formats},
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
	success, ok := result.(*GETTagsAssignedToHostOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getTagsAssignedToHost: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GETTagsAssignedToHosts Gets tags assigned to hosts
*/
func (a *Client) GETTagsAssignedToHosts(params *GETTagsAssignedToHostsParams, opts ...ClientOption) (*GETTagsAssignedToHostsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGETTagsAssignedToHostsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getTagsAssignedToHosts",
		Method:             "GET",
		PathPattern:        "/v1/hosts/tags",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GETTagsAssignedToHostsReader{formats: a.formats},
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
	success, ok := result.(*GETTagsAssignedToHostsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getTagsAssignedToHosts: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GETValidationForCommissionHosts Gets the status of the validation of the input specification to commission the hosts
*/
func (a *Client) GETValidationForCommissionHosts(params *GETValidationForCommissionHostsParams, opts ...ClientOption) (*GETValidationForCommissionHostsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGETValidationForCommissionHostsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getValidationForCommissionHosts",
		Method:             "GET",
		PathPattern:        "/v1/hosts/validations/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GETValidationForCommissionHostsReader{formats: a.formats},
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
	success, ok := result.(*GETValidationForCommissionHostsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getValidationForCommissionHosts: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
POSTQuery Posts a query
*/
func (a *Client) POSTQuery(params *POSTQueryParams, opts ...ClientOption) (*POSTQueryOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPOSTQueryParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "postQuery",
		Method:             "POST",
		PathPattern:        "/v1/hosts/queries",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &POSTQueryReader{formats: a.formats},
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
	success, ok := result.(*POSTQueryOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for postQuery: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
RemoveTagsFromHost removes tags from host
*/
func (a *Client) RemoveTagsFromHost(params *RemoveTagsFromHostParams, opts ...ClientOption) (*RemoveTagsFromHostOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRemoveTagsFromHostParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "removeTagsFromHost",
		Method:             "DELETE",
		PathPattern:        "/v1/hosts/{id}/tags",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &RemoveTagsFromHostReader{formats: a.formats},
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
	success, ok := result.(*RemoveTagsFromHostOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for removeTagsFromHost: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
ValidateHostsOperations validates the input spec for hosts operations
*/
func (a *Client) ValidateHostsOperations(params *ValidateHostsOperationsParams, opts ...ClientOption) (*ValidateHostsOperationsOK, *ValidateHostsOperationsAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewValidateHostsOperationsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "validateHostsOperations",
		Method:             "POST",
		PathPattern:        "/v1/hosts/validations",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ValidateHostsOperationsReader{formats: a.formats},
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
	case *ValidateHostsOperationsOK:
		return value, nil, nil
	case *ValidateHostsOperationsAccepted:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for hosts: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
