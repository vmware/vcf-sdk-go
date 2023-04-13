// Code generated by go-swagger; DO NOT EDIT.

// Copyright © 2023 VMware, Inc. All Rights Reserved.
// Code licenced under: BSD-2

package credentials

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new credentials API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for credentials API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CancelCredentialsTask(params *CancelCredentialsTaskParams, opts ...ClientOption) (*CancelCredentialsTaskOK, *CancelCredentialsTaskNoContent, error)

	FetchPasswordExpiration(params *FetchPasswordExpirationParams, opts ...ClientOption) (*FetchPasswordExpirationOK, *FetchPasswordExpirationAccepted, error)

	GETCredential(params *GETCredentialParams, opts ...ClientOption) (*GETCredentialOK, error)

	GETCredentials(params *GETCredentialsParams, opts ...ClientOption) (*GETCredentialsOK, error)

	GETCredentialsSubTask(params *GETCredentialsSubTaskParams, opts ...ClientOption) (*GETCredentialsSubTaskOK, error)

	GETCredentialsTask(params *GETCredentialsTaskParams, opts ...ClientOption) (*GETCredentialsTaskOK, error)

	GETCredentialsTaskResourcesCredentials(params *GETCredentialsTaskResourcesCredentialsParams, opts ...ClientOption) (*GETCredentialsTaskResourcesCredentialsOK, error)

	GETCredentialsTasks(params *GETCredentialsTasksParams, opts ...ClientOption) (*GETCredentialsTasksOK, error)

	GETExpirationsForPasswords(params *GETExpirationsForPasswordsParams, opts ...ClientOption) (*GETExpirationsForPasswordsOK, error)

	RetryCredentialsTask(params *RetryCredentialsTaskParams, opts ...ClientOption) (*RetryCredentialsTaskOK, *RetryCredentialsTaskAccepted, error)

	UpdateOrRotatePasswords(params *UpdateOrRotatePasswordsParams, opts ...ClientOption) (*UpdateOrRotatePasswordsOK, *UpdateOrRotatePasswordsAccepted, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
CancelCredentialsTask cancels a failed credentials task for a given ID

Cancel a failed credentials task for a given ID
*/
func (a *Client) CancelCredentialsTask(params *CancelCredentialsTaskParams, opts ...ClientOption) (*CancelCredentialsTaskOK, *CancelCredentialsTaskNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCancelCredentialsTaskParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "cancelCredentialsTask",
		Method:             "DELETE",
		PathPattern:        "/v1/credentials/tasks/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CancelCredentialsTaskReader{formats: a.formats},
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
	case *CancelCredentialsTaskOK:
		return value, nil, nil
	case *CancelCredentialsTaskNoContent:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for credentials: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
FetchPasswordExpiration fetches expiration details of passwords for a list of credentials

Fetch expiration details of passwords for a list of credentials
*/
func (a *Client) FetchPasswordExpiration(params *FetchPasswordExpirationParams, opts ...ClientOption) (*FetchPasswordExpirationOK, *FetchPasswordExpirationAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewFetchPasswordExpirationParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "fetchPasswordExpiration",
		Method:             "POST",
		PathPattern:        "/v1/credentials/expirations",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &FetchPasswordExpirationReader{formats: a.formats},
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
	case *FetchPasswordExpirationOK:
		return value, nil, nil
	case *FetchPasswordExpirationAccepted:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for credentials: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GETCredential Gets credential for the given ID

Get Credential for the given ID
*/
func (a *Client) GETCredential(params *GETCredentialParams, opts ...ClientOption) (*GETCredentialOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGETCredentialParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getCredential",
		Method:             "GET",
		PathPattern:        "/v1/credentials/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GETCredentialReader{formats: a.formats},
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
	success, ok := result.(*GETCredentialOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getCredential: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GETCredentials Gets the credentials

Get the Credentials
*/
func (a *Client) GETCredentials(params *GETCredentialsParams, opts ...ClientOption) (*GETCredentialsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGETCredentialsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getCredentials",
		Method:             "GET",
		PathPattern:        "/v1/credentials",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GETCredentialsReader{formats: a.formats},
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
	success, ok := result.(*GETCredentialsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getCredentials: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GETCredentialsSubTask fetches details of a subtask for a given credentials task ID and sub task ID

Fetch details of a subtask for a given credentials task ID and sub-task ID.
*/
func (a *Client) GETCredentialsSubTask(params *GETCredentialsSubTaskParams, opts ...ClientOption) (*GETCredentialsSubTaskOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGETCredentialsSubTaskParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getCredentialsSubTask",
		Method:             "GET",
		PathPattern:        "/v1/credentials/tasks/{id}/subtasks/{subtaskId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GETCredentialsSubTaskReader{formats: a.formats},
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
	success, ok := result.(*GETCredentialsSubTaskOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getCredentialsSubTask: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GETCredentialsTask fetches a credentials task

Fetch credentials task for a given ID
*/
func (a *Client) GETCredentialsTask(params *GETCredentialsTaskParams, opts ...ClientOption) (*GETCredentialsTaskOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGETCredentialsTaskParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getCredentialsTask",
		Method:             "GET",
		PathPattern:        "/v1/credentials/tasks/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GETCredentialsTaskReader{formats: a.formats},
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
	success, ok := result.(*GETCredentialsTaskOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getCredentialsTask: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GETCredentialsTaskResourcesCredentials fetches resource credentials for a given credentials task ID

Fetch resource credentials for a given credentials task ID
*/
func (a *Client) GETCredentialsTaskResourcesCredentials(params *GETCredentialsTaskResourcesCredentialsParams, opts ...ClientOption) (*GETCredentialsTaskResourcesCredentialsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGETCredentialsTaskResourcesCredentialsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getCredentialsTaskResourcesCredentials",
		Method:             "GET",
		PathPattern:        "/v1/credentials/tasks/{id}/resource-credentials",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GETCredentialsTaskResourcesCredentialsReader{formats: a.formats},
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
	success, ok := result.(*GETCredentialsTaskResourcesCredentialsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getCredentialsTaskResourcesCredentials: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GETCredentialsTasks fetches the credentials tasks

Fetch all credentials tasks in reverse chronological order
*/
func (a *Client) GETCredentialsTasks(params *GETCredentialsTasksParams, opts ...ClientOption) (*GETCredentialsTasksOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGETCredentialsTasksParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getCredentialsTasks",
		Method:             "GET",
		PathPattern:        "/v1/credentials/tasks",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GETCredentialsTasksReader{formats: a.formats},
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
	success, ok := result.(*GETCredentialsTasksOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getCredentialsTasks: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GETExpirationsForPasswords Gets the status of the password expiration fetch

Get the status of the password expiration fetch
*/
func (a *Client) GETExpirationsForPasswords(params *GETExpirationsForPasswordsParams, opts ...ClientOption) (*GETExpirationsForPasswordsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGETExpirationsForPasswordsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getExpirationsForPasswords",
		Method:             "GET",
		PathPattern:        "/v1/credentials/expirations/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GETExpirationsForPasswordsReader{formats: a.formats},
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
	success, ok := result.(*GETExpirationsForPasswordsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getExpirationsForPasswords: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
RetryCredentialsTask retries a failed credentials task for a given ID

Retry a failed credentials task for a given ID
*/
func (a *Client) RetryCredentialsTask(params *RetryCredentialsTaskParams, opts ...ClientOption) (*RetryCredentialsTaskOK, *RetryCredentialsTaskAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRetryCredentialsTaskParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "retryCredentialsTask",
		Method:             "PATCH",
		PathPattern:        "/v1/credentials/tasks/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &RetryCredentialsTaskReader{formats: a.formats},
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
	case *RetryCredentialsTaskOK:
		return value, nil, nil
	case *RetryCredentialsTaskAccepted:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for credentials: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
UpdateOrRotatePasswords updates or rotate passwords for a list of resources

Update passwords for given list of resources by supplying new passwords or rotate the passwords using system generated passwords
*/
func (a *Client) UpdateOrRotatePasswords(params *UpdateOrRotatePasswordsParams, opts ...ClientOption) (*UpdateOrRotatePasswordsOK, *UpdateOrRotatePasswordsAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateOrRotatePasswordsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "updateOrRotatePasswords",
		Method:             "PATCH",
		PathPattern:        "/v1/credentials",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateOrRotatePasswordsReader{formats: a.formats},
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
	case *UpdateOrRotatePasswordsOK:
		return value, nil, nil
	case *UpdateOrRotatePasswordsAccepted:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for credentials: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
