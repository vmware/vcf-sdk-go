// Code generated by go-swagger; DO NOT EDIT.

// Copyright © 2023 VMware, Inc. All Rights Reserved.
// Code licenced under: BSD-2

package identity_providers

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

	"vcf-sdk-go/models"
)

// NewAddExternalIdentityProviderParams creates a new AddExternalIdentityProviderParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAddExternalIdentityProviderParams() *AddExternalIdentityProviderParams {
	return &AddExternalIdentityProviderParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAddExternalIdentityProviderParamsWithTimeout creates a new AddExternalIdentityProviderParams object
// with the ability to set a timeout on a request.
func NewAddExternalIdentityProviderParamsWithTimeout(timeout time.Duration) *AddExternalIdentityProviderParams {
	return &AddExternalIdentityProviderParams{
		timeout: timeout,
	}
}

// NewAddExternalIdentityProviderParamsWithContext creates a new AddExternalIdentityProviderParams object
// with the ability to set a context for a request.
func NewAddExternalIdentityProviderParamsWithContext(ctx context.Context) *AddExternalIdentityProviderParams {
	return &AddExternalIdentityProviderParams{
		Context: ctx,
	}
}

// NewAddExternalIdentityProviderParamsWithHTTPClient creates a new AddExternalIdentityProviderParams object
// with the ability to set a custom HTTPClient for a request.
func NewAddExternalIdentityProviderParamsWithHTTPClient(client *http.Client) *AddExternalIdentityProviderParams {
	return &AddExternalIdentityProviderParams{
		HTTPClient: client,
	}
}

/*
AddExternalIdentityProviderParams contains all the parameters to send to the API endpoint

	for the add external identity provider operation.

	Typically these are written to a http.Request.
*/
type AddExternalIdentityProviderParams struct {

	/* IdentityProviderSpec.

	   Identity Provider Spec
	*/
	IdentityProviderSpec *models.IdentityProviderSpec

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the add external identity provider params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AddExternalIdentityProviderParams) WithDefaults() *AddExternalIdentityProviderParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the add external identity provider params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AddExternalIdentityProviderParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the add external identity provider params
func (o *AddExternalIdentityProviderParams) WithTimeout(timeout time.Duration) *AddExternalIdentityProviderParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the add external identity provider params
func (o *AddExternalIdentityProviderParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the add external identity provider params
func (o *AddExternalIdentityProviderParams) WithContext(ctx context.Context) *AddExternalIdentityProviderParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the add external identity provider params
func (o *AddExternalIdentityProviderParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the add external identity provider params
func (o *AddExternalIdentityProviderParams) WithHTTPClient(client *http.Client) *AddExternalIdentityProviderParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the add external identity provider params
func (o *AddExternalIdentityProviderParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIdentityProviderSpec adds the identityProviderSpec to the add external identity provider params
func (o *AddExternalIdentityProviderParams) WithIdentityProviderSpec(identityProviderSpec *models.IdentityProviderSpec) *AddExternalIdentityProviderParams {
	o.SetIdentityProviderSpec(identityProviderSpec)
	return o
}

// SetIdentityProviderSpec adds the identityProviderSpec to the add external identity provider params
func (o *AddExternalIdentityProviderParams) SetIdentityProviderSpec(identityProviderSpec *models.IdentityProviderSpec) {
	o.IdentityProviderSpec = identityProviderSpec
}

// WriteToRequest writes these params to a swagger request
func (o *AddExternalIdentityProviderParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.IdentityProviderSpec != nil {
		if err := r.SetBodyParam(o.IdentityProviderSpec); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
