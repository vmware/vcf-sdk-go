// Code generated by go-swagger; DO NOT EDIT.

package nsx_t_edge_clusters

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new nsx t edge clusters API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for nsx t edge clusters API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CreateEdge(params *CreateEdgeParams, opts ...ClientOption) (*CreateEdgeOK, *CreateEdgeAccepted, error)

	GetEdgeCluster(params *GetEdgeClusterParams, opts ...ClientOption) (*GetEdgeClusterOK, error)

	GetEdgeClusters(params *GetEdgeClustersParams, opts ...ClientOption) (*GetEdgeClustersOK, error)

	GetValidationForCreateEdgeCluster(params *GetValidationForCreateEdgeClusterParams, opts ...ClientOption) (*GetValidationForCreateEdgeClusterOK, error)

	UpdateEdgeCluster(params *UpdateEdgeClusterParams, opts ...ClientOption) (*UpdateEdgeClusterOK, *UpdateEdgeClusterAccepted, error)

	ValidateEdgeClusterSpec(params *ValidateEdgeClusterSpecParams, opts ...ClientOption) (*ValidateEdgeClusterSpecOK, *ValidateEdgeClusterSpecAccepted, error)

	ValidateEdgeClusterUpdateSpec(params *ValidateEdgeClusterUpdateSpecParams, opts ...ClientOption) (*ValidateEdgeClusterUpdateSpecOK, *ValidateEdgeClusterUpdateSpecAccepted, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
CreateEdge creates an edge cluster
*/
func (a *Client) CreateEdge(params *CreateEdgeParams, opts ...ClientOption) (*CreateEdgeOK, *CreateEdgeAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateEdgeParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "createEdge",
		Method:             "POST",
		PathPattern:        "/v1/edge-clusters",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateEdgeReader{formats: a.formats},
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
	case *CreateEdgeOK:
		return value, nil, nil
	case *CreateEdgeAccepted:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for nsx_t_edge_clusters: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetEdgeCluster gets an edge cluster
*/
func (a *Client) GetEdgeCluster(params *GetEdgeClusterParams, opts ...ClientOption) (*GetEdgeClusterOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetEdgeClusterParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getEdgeCluster",
		Method:             "GET",
		PathPattern:        "/v1/edge-clusters/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetEdgeClusterReader{formats: a.formats},
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
	success, ok := result.(*GetEdgeClusterOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getEdgeCluster: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetEdgeClusters gets the edge clusters
*/
func (a *Client) GetEdgeClusters(params *GetEdgeClustersParams, opts ...ClientOption) (*GetEdgeClustersOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetEdgeClustersParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getEdgeClusters",
		Method:             "GET",
		PathPattern:        "/v1/edge-clusters",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetEdgeClustersReader{formats: a.formats},
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
	success, ok := result.(*GetEdgeClustersOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getEdgeClusters: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetValidationForCreateEdgeCluster gets the edge cluster spec validation
*/
func (a *Client) GetValidationForCreateEdgeCluster(params *GetValidationForCreateEdgeClusterParams, opts ...ClientOption) (*GetValidationForCreateEdgeClusterOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetValidationForCreateEdgeClusterParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getValidationForCreateEdgeCluster",
		Method:             "GET",
		PathPattern:        "/v1/edge-clusters/validations/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetValidationForCreateEdgeClusterReader{formats: a.formats},
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
	success, ok := result.(*GetValidationForCreateEdgeClusterOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getValidationForCreateEdgeCluster: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
UpdateEdgeCluster expands or shrink an n s x t edge cluster
*/
func (a *Client) UpdateEdgeCluster(params *UpdateEdgeClusterParams, opts ...ClientOption) (*UpdateEdgeClusterOK, *UpdateEdgeClusterAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateEdgeClusterParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "updateEdgeCluster",
		Method:             "PATCH",
		PathPattern:        "/v1/edge-clusters/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateEdgeClusterReader{formats: a.formats},
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
	case *UpdateEdgeClusterOK:
		return value, nil, nil
	case *UpdateEdgeClusterAccepted:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for nsx_t_edge_clusters: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
ValidateEdgeClusterSpec validates an edge cluster spec
*/
func (a *Client) ValidateEdgeClusterSpec(params *ValidateEdgeClusterSpecParams, opts ...ClientOption) (*ValidateEdgeClusterSpecOK, *ValidateEdgeClusterSpecAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewValidateEdgeClusterSpecParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "validateEdgeClusterSpec",
		Method:             "POST",
		PathPattern:        "/v1/edge-clusters/validations",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ValidateEdgeClusterSpecReader{formats: a.formats},
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
	case *ValidateEdgeClusterSpecOK:
		return value, nil, nil
	case *ValidateEdgeClusterSpecAccepted:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for nsx_t_edge_clusters: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
ValidateEdgeClusterUpdateSpec validates an n s x t edge cluster update spec
*/
func (a *Client) ValidateEdgeClusterUpdateSpec(params *ValidateEdgeClusterUpdateSpecParams, opts ...ClientOption) (*ValidateEdgeClusterUpdateSpecOK, *ValidateEdgeClusterUpdateSpecAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewValidateEdgeClusterUpdateSpecParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "validateEdgeClusterUpdateSpec",
		Method:             "POST",
		PathPattern:        "/v1/edge-clusters/{id}/validations",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ValidateEdgeClusterUpdateSpecReader{formats: a.formats},
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
	case *ValidateEdgeClusterUpdateSpecOK:
		return value, nil, nil
	case *ValidateEdgeClusterUpdateSpecAccepted:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for nsx_t_edge_clusters: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}