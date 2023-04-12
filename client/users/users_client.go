// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new users API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for users API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	AddUsers(params *AddUsersParams, opts ...ClientOption) (*AddUsersOK, *AddUsersCreated, error)

	DeleteUser(params *DeleteUserParams, opts ...ClientOption) (*DeleteUserNoContent, error)

	GetAllUIUsersUsingGET(params *GetAllUIUsersUsingGETParams, opts ...ClientOption) (*GetAllUIUsersUsingGETOK, error)

	GetLocalAccount(params *GetLocalAccountParams, opts ...ClientOption) (*GetLocalAccountOK, error)

	GetRoles(params *GetRolesParams, opts ...ClientOption) (*GetRolesOK, error)

	GetSSODomainEntities(params *GetSSODomainEntitiesParams, opts ...ClientOption) (*GetSSODomainEntitiesOK, error)

	GetSSODomains(params *GetSSODomainsParams, opts ...ClientOption) (*GetSSODomainsOK, error)

	GetUsers(params *GetUsersParams, opts ...ClientOption) (*GetUsersOK, error)

	UpdateLocalUserPassword(params *UpdateLocalUserPasswordParams, opts ...ClientOption) (*UpdateLocalUserPasswordNoContent, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
AddUsers adds users

Add list of users
*/
func (a *Client) AddUsers(params *AddUsersParams, opts ...ClientOption) (*AddUsersOK, *AddUsersCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAddUsersParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "addUsers",
		Method:             "POST",
		PathPattern:        "/v1/users",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &AddUsersReader{formats: a.formats},
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
	case *AddUsersOK:
		return value, nil, nil
	case *AddUsersCreated:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for users: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
DeleteUser deletes a user

Delete the user by the ID, if it exists
*/
func (a *Client) DeleteUser(params *DeleteUserParams, opts ...ClientOption) (*DeleteUserNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteUserParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "deleteUser",
		Method:             "DELETE",
		PathPattern:        "/v1/users/{id}",
		ProducesMediaTypes: []string{"*/*"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteUserReader{formats: a.formats},
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
	success, ok := result.(*DeleteUserNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteUser: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetAllUIUsersUsingGET gets all Ui users
*/
func (a *Client) GetAllUIUsersUsingGET(params *GetAllUIUsersUsingGETParams, opts ...ClientOption) (*GetAllUIUsersUsingGETOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAllUIUsersUsingGETParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getAllUiUsersUsingGET",
		Method:             "GET",
		PathPattern:        "/v1/users/ui",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetAllUIUsersUsingGETReader{formats: a.formats},
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
	success, ok := result.(*GetAllUIUsersUsingGETOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getAllUiUsersUsingGET: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetLocalAccount gets local account details

Get information on the local account
*/
func (a *Client) GetLocalAccount(params *GetLocalAccountParams, opts ...ClientOption) (*GetLocalAccountOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetLocalAccountParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getLocalAccount",
		Method:             "GET",
		PathPattern:        "/v1/users/local/admin",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetLocalAccountReader{formats: a.formats},
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
	success, ok := result.(*GetLocalAccountOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getLocalAccount: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetRoles gets all roles

Get a list of all roles
*/
func (a *Client) GetRoles(params *GetRolesParams, opts ...ClientOption) (*GetRolesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetRolesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getRoles",
		Method:             "GET",
		PathPattern:        "/v1/roles",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetRolesReader{formats: a.formats},
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
	success, ok := result.(*GetRolesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getRoles: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetSSODomainEntities gets all entities of s s o domain

Get a list of all entities in the SSO domain
*/
func (a *Client) GetSSODomainEntities(params *GetSSODomainEntitiesParams, opts ...ClientOption) (*GetSSODomainEntitiesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetSSODomainEntitiesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getSSODomainEntities",
		Method:             "GET",
		PathPattern:        "/v1/sso-domains/{sso-domain}/entities",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetSSODomainEntitiesReader{formats: a.formats},
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
	success, ok := result.(*GetSSODomainEntitiesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getSSODomainEntities: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetSSODomains gets all s s o domains

Get a list of all SSO domains
*/
func (a *Client) GetSSODomains(params *GetSSODomainsParams, opts ...ClientOption) (*GetSSODomainsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetSSODomainsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getSSODomains",
		Method:             "GET",
		PathPattern:        "/v1/sso-domains",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetSSODomainsReader{formats: a.formats},
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
	success, ok := result.(*GetSSODomainsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getSSODomains: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetUsers gets all users

Get a list of all users
*/
func (a *Client) GetUsers(params *GetUsersParams, opts ...ClientOption) (*GetUsersOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetUsersParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getUsers",
		Method:             "GET",
		PathPattern:        "/v1/users",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetUsersReader{formats: a.formats},
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
	success, ok := result.(*GetUsersOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getUsers: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
UpdateLocalUserPassword updates password for local account

Update the password for local account only if the old password is correct, or if user configures the local account for the first time
*/
func (a *Client) UpdateLocalUserPassword(params *UpdateLocalUserPasswordParams, opts ...ClientOption) (*UpdateLocalUserPasswordNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateLocalUserPasswordParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "updateLocalUserPassword",
		Method:             "PATCH",
		PathPattern:        "/v1/users/local/admin",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateLocalUserPasswordReader{formats: a.formats},
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
	success, ok := result.(*UpdateLocalUserPasswordNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for updateLocalUserPassword: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
