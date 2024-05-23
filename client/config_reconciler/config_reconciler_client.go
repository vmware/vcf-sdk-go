// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2023 VMware, Inc.
// SPDX-License-Identifier: BSD-2-Clause

package config_reconciler

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new config reconciler API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for config reconciler API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	GetConfigs(params *GetConfigsParams, opts ...ClientOption) (*GetConfigsOK, error)

	GetReconciliationTask(params *GetReconciliationTaskParams, opts ...ClientOption) (*GetReconciliationTaskOK, error)

	ReconcileConfigs(params *ReconcileConfigsParams, opts ...ClientOption) (*ReconcileConfigsOK, *ReconcileConfigsAccepted, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
GetConfigs gets configs associated with the given criteria

Get configs associated with the given criteria, all if no criteria is provided
*/
func (a *Client) GetConfigs(params *GetConfigsParams, opts ...ClientOption) (*GetConfigsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetConfigsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getConfigs",
		Method:             "GET",
		PathPattern:        "/v1/config-drifts",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetConfigsReader{formats: a.formats},
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
	success, ok := result.(*GetConfigsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getConfigs: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetReconciliationTask gets config reconciliation task associated with the given task Id

Get config reconciliation task associated with the given task Id
*/
func (a *Client) GetReconciliationTask(params *GetReconciliationTaskParams, opts ...ClientOption) (*GetReconciliationTaskOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetReconciliationTaskParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getReconciliationTask",
		Method:             "GET",
		PathPattern:        "/v1/config-drift-reconciliations/{taskId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetReconciliationTaskReader{formats: a.formats},
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
	success, ok := result.(*GetReconciliationTaskOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getReconciliationTask: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
ReconcileConfigs reconciles configs

For selective reconciliation, provide a config spec.
*/
func (a *Client) ReconcileConfigs(params *ReconcileConfigsParams, opts ...ClientOption) (*ReconcileConfigsOK, *ReconcileConfigsAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewReconcileConfigsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "reconcileConfigs",
		Method:             "POST",
		PathPattern:        "/v1/config-drift-reconciliations",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ReconcileConfigsReader{formats: a.formats},
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
	case *ReconcileConfigsOK:
		return value, nil, nil
	case *ReconcileConfigsAccepted:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for config_reconciler: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
