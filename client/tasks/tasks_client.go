// Code generated by go-swagger; DO NOT EDIT.

package tasks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new tasks API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for tasks API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CancelTask(params *CancelTaskParams, opts ...ClientOption) (*CancelTaskOK, error)

	GetTask(params *GetTaskParams, opts ...ClientOption) (*GetTaskOK, error)

	GetTasks(params *GetTasksParams, opts ...ClientOption) (*GetTasksOK, error)

	RetryTask(params *RetryTaskParams, opts ...ClientOption) (*RetryTaskOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
CancelTask cancels a task

Cancel a Task by ID, if it exists
*/
func (a *Client) CancelTask(params *CancelTaskParams, opts ...ClientOption) (*CancelTaskOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCancelTaskParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "cancelTask",
		Method:             "DELETE",
		PathPattern:        "/v1/tasks/{id}",
		ProducesMediaTypes: []string{"*/*"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CancelTaskReader{formats: a.formats},
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
	success, ok := result.(*CancelTaskOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for cancelTask: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetTask gets a task

Get a Task by ID, if it exists
*/
func (a *Client) GetTask(params *GetTaskParams, opts ...ClientOption) (*GetTaskOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetTaskParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getTask",
		Method:             "GET",
		PathPattern:        "/v1/tasks/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetTaskReader{formats: a.formats},
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
	success, ok := result.(*GetTaskOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getTask: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetTasks gets the tasks

Get the Tasks
*/
func (a *Client) GetTasks(params *GetTasksParams, opts ...ClientOption) (*GetTasksOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetTasksParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getTasks",
		Method:             "GET",
		PathPattern:        "/v1/tasks",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetTasksReader{formats: a.formats},
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
	success, ok := result.(*GetTasksOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getTasks: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
RetryTask retries a task

Retry a failed Task by ID, if it exists
*/
func (a *Client) RetryTask(params *RetryTaskParams, opts ...ClientOption) (*RetryTaskOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRetryTaskParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "retryTask",
		Method:             "PATCH",
		PathPattern:        "/v1/tasks/{id}",
		ProducesMediaTypes: []string{"*/*"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &RetryTaskReader{formats: a.formats},
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
	success, ok := result.(*RetryTaskOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for retryTask: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}