// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2023 VMware, Inc.
// SPDX-License-Identifier: BSD-2-Clause

package vasa_providers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/vmware/vcf-sdk-go/models"
)

// NewAddUsersToVasaProviderParams creates a new AddUsersToVasaProviderParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAddUsersToVasaProviderParams() *AddUsersToVasaProviderParams {
	return &AddUsersToVasaProviderParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAddUsersToVasaProviderParamsWithTimeout creates a new AddUsersToVasaProviderParams object
// with the ability to set a timeout on a request.
func NewAddUsersToVasaProviderParamsWithTimeout(timeout time.Duration) *AddUsersToVasaProviderParams {
	return &AddUsersToVasaProviderParams{
		timeout: timeout,
	}
}

// NewAddUsersToVasaProviderParamsWithContext creates a new AddUsersToVasaProviderParams object
// with the ability to set a context for a request.
func NewAddUsersToVasaProviderParamsWithContext(ctx context.Context) *AddUsersToVasaProviderParams {
	return &AddUsersToVasaProviderParams{
		Context: ctx,
	}
}

// NewAddUsersToVasaProviderParamsWithHTTPClient creates a new AddUsersToVasaProviderParams object
// with the ability to set a custom HTTPClient for a request.
func NewAddUsersToVasaProviderParamsWithHTTPClient(client *http.Client) *AddUsersToVasaProviderParams {
	return &AddUsersToVasaProviderParams{
		HTTPClient: client,
	}
}

/*
AddUsersToVasaProviderParams contains all the parameters to send to the API endpoint

	for the add users to vasa provider operation.

	Typically these are written to a http.Request.
*/
type AddUsersToVasaProviderParams struct {

	/* ID.

	   VASA Provider ID
	*/
	ID string

	/* VasaUsers.

	   VASA Users data
	*/
	VasaUsers []*models.VasaUser

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the add users to vasa provider params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AddUsersToVasaProviderParams) WithDefaults() *AddUsersToVasaProviderParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the add users to vasa provider params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AddUsersToVasaProviderParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the add users to vasa provider params
func (o *AddUsersToVasaProviderParams) WithTimeout(timeout time.Duration) *AddUsersToVasaProviderParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the add users to vasa provider params
func (o *AddUsersToVasaProviderParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the add users to vasa provider params
func (o *AddUsersToVasaProviderParams) WithContext(ctx context.Context) *AddUsersToVasaProviderParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the add users to vasa provider params
func (o *AddUsersToVasaProviderParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the add users to vasa provider params
func (o *AddUsersToVasaProviderParams) WithHTTPClient(client *http.Client) *AddUsersToVasaProviderParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the add users to vasa provider params
func (o *AddUsersToVasaProviderParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the add users to vasa provider params
func (o *AddUsersToVasaProviderParams) WithID(id string) *AddUsersToVasaProviderParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the add users to vasa provider params
func (o *AddUsersToVasaProviderParams) SetID(id string) {
	o.ID = id
}

// WithVasaUsers adds the vasaUsers to the add users to vasa provider params
func (o *AddUsersToVasaProviderParams) WithVasaUsers(vasaUsers []*models.VasaUser) *AddUsersToVasaProviderParams {
	o.SetVasaUsers(vasaUsers)
	return o
}

// SetVasaUsers adds the vasaUsers to the add users to vasa provider params
func (o *AddUsersToVasaProviderParams) SetVasaUsers(vasaUsers []*models.VasaUser) {
	o.VasaUsers = vasaUsers
}

// WriteToRequest writes these params to a swagger request
func (o *AddUsersToVasaProviderParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}
	if o.VasaUsers != nil {
		if err := r.SetBodyParam(o.VasaUsers); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
