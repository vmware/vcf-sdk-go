// Code generated by go-swagger; DO NOT EDIT.

// Copyright © 2023 VMware, Inc. All Rights Reserved.
// Code licenced under: BSD-2

package tokens

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

// NewCreateTokenParams creates a new CreateTokenParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateTokenParams() *CreateTokenParams {
	return &CreateTokenParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateTokenParamsWithTimeout creates a new CreateTokenParams object
// with the ability to set a timeout on a request.
func NewCreateTokenParamsWithTimeout(timeout time.Duration) *CreateTokenParams {
	return &CreateTokenParams{
		timeout: timeout,
	}
}

// NewCreateTokenParamsWithContext creates a new CreateTokenParams object
// with the ability to set a context for a request.
func NewCreateTokenParamsWithContext(ctx context.Context) *CreateTokenParams {
	return &CreateTokenParams{
		Context: ctx,
	}
}

// NewCreateTokenParamsWithHTTPClient creates a new CreateTokenParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateTokenParamsWithHTTPClient(client *http.Client) *CreateTokenParams {
	return &CreateTokenParams{
		HTTPClient: client,
	}
}

/*
CreateTokenParams contains all the parameters to send to the API endpoint

	for the create token operation.

	Typically these are written to a http.Request.
*/
type CreateTokenParams struct {

	/* TokenCreationSpec.

	   The spec used to sign the token
	*/
	TokenCreationSpec *models.TokenCreationSpec

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create token params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateTokenParams) WithDefaults() *CreateTokenParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create token params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateTokenParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create token params
func (o *CreateTokenParams) WithTimeout(timeout time.Duration) *CreateTokenParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create token params
func (o *CreateTokenParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create token params
func (o *CreateTokenParams) WithContext(ctx context.Context) *CreateTokenParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create token params
func (o *CreateTokenParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create token params
func (o *CreateTokenParams) WithHTTPClient(client *http.Client) *CreateTokenParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create token params
func (o *CreateTokenParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithTokenCreationSpec adds the tokenCreationSpec to the create token params
func (o *CreateTokenParams) WithTokenCreationSpec(tokenCreationSpec *models.TokenCreationSpec) *CreateTokenParams {
	o.SetTokenCreationSpec(tokenCreationSpec)
	return o
}

// SetTokenCreationSpec adds the tokenCreationSpec to the create token params
func (o *CreateTokenParams) SetTokenCreationSpec(tokenCreationSpec *models.TokenCreationSpec) {
	o.TokenCreationSpec = tokenCreationSpec
}

// WriteToRequest writes these params to a swagger request
func (o *CreateTokenParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.TokenCreationSpec != nil {
		if err := r.SetBodyParam(o.TokenCreationSpec); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
