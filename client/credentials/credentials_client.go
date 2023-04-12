// Code generated by go-swagger; DO NOT EDIT.

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

	GetCredential(params *GetCredentialParams, opts ...ClientOption) (*GetCredentialOK, error)

	GetCredentials(params *GetCredentialsParams, opts ...ClientOption) (*GetCredentialsOK, error)

	GetCredentialsSubTask(params *GetCredentialsSubTaskParams, opts ...ClientOption) (*GetCredentialsSubTaskOK, error)

	GetCredentialsTask(params *GetCredentialsTaskParams, opts ...ClientOption) (*GetCredentialsTaskOK, error)

	GetCredentialsTaskResourcesCredentials(params *GetCredentialsTaskResourcesCredentialsParams, opts ...ClientOption) (*GetCredentialsTaskResourcesCredentialsOK, error)

	GetCredentialsTasks(params *GetCredentialsTasksParams, opts ...ClientOption) (*GetCredentialsTasksOK, error)

	GetExpirationsForPasswords(params *GetExpirationsForPasswordsParams, opts ...ClientOption) (*GetExpirationsForPasswordsOK, error)

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
GetCredential gets credential for the given ID

Get Credential for the given ID
*/
func (a *Client) GetCredential(params *GetCredentialParams, opts ...ClientOption) (*GetCredentialOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCredentialParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getCredential",
		Method:             "GET",
		PathPattern:        "/v1/credentials/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetCredentialReader{formats: a.formats},
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
	success, ok := result.(*GetCredentialOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getCredential: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetCredentials gets the credentials

Get the Credentials
*/
func (a *Client) GetCredentials(params *GetCredentialsParams, opts ...ClientOption) (*GetCredentialsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCredentialsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getCredentials",
		Method:             "GET",
		PathPattern:        "/v1/credentials",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetCredentialsReader{formats: a.formats},
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
	success, ok := result.(*GetCredentialsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getCredentials: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetCredentialsSubTask fetches details of a subtask for a given credentials task ID and sub task ID

Fetch details of a subtask for a given credentials task ID and sub-task ID.
*/
func (a *Client) GetCredentialsSubTask(params *GetCredentialsSubTaskParams, opts ...ClientOption) (*GetCredentialsSubTaskOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCredentialsSubTaskParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getCredentialsSubTask",
		Method:             "GET",
		PathPattern:        "/v1/credentials/tasks/{id}/subtasks/{subtaskId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetCredentialsSubTaskReader{formats: a.formats},
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
	success, ok := result.(*GetCredentialsSubTaskOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getCredentialsSubTask: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetCredentialsTask fetches a credentials task

Fetch credentials task for a given ID
*/
func (a *Client) GetCredentialsTask(params *GetCredentialsTaskParams, opts ...ClientOption) (*GetCredentialsTaskOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCredentialsTaskParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getCredentialsTask",
		Method:             "GET",
		PathPattern:        "/v1/credentials/tasks/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetCredentialsTaskReader{formats: a.formats},
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
	success, ok := result.(*GetCredentialsTaskOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getCredentialsTask: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetCredentialsTaskResourcesCredentials fetches resource credentials for a given credentials task ID

Fetch resource credentials for a given credentials task ID
*/
func (a *Client) GetCredentialsTaskResourcesCredentials(params *GetCredentialsTaskResourcesCredentialsParams, opts ...ClientOption) (*GetCredentialsTaskResourcesCredentialsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCredentialsTaskResourcesCredentialsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getCredentialsTaskResourcesCredentials",
		Method:             "GET",
		PathPattern:        "/v1/credentials/tasks/{id}/resource-credentials",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetCredentialsTaskResourcesCredentialsReader{formats: a.formats},
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
	success, ok := result.(*GetCredentialsTaskResourcesCredentialsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getCredentialsTaskResourcesCredentials: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetCredentialsTasks fetches the credentials tasks

Fetch all credentials tasks in reverse chronological order
*/
func (a *Client) GetCredentialsTasks(params *GetCredentialsTasksParams, opts ...ClientOption) (*GetCredentialsTasksOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCredentialsTasksParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getCredentialsTasks",
		Method:             "GET",
		PathPattern:        "/v1/credentials/tasks",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetCredentialsTasksReader{formats: a.formats},
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
	success, ok := result.(*GetCredentialsTasksOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getCredentialsTasks: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetExpirationsForPasswords gets the status of the password expiration fetch

Get the status of the password expiration fetch
*/
func (a *Client) GetExpirationsForPasswords(params *GetExpirationsForPasswordsParams, opts ...ClientOption) (*GetExpirationsForPasswordsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetExpirationsForPasswordsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getExpirationsForPasswords",
		Method:             "GET",
		PathPattern:        "/v1/credentials/expirations/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetExpirationsForPasswordsReader{formats: a.formats},
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
	success, ok := result.(*GetExpirationsForPasswordsOK)
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