// Code generated by go-swagger; DO NOT EDIT.

// Copyright © 2023 VMware, Inc. All Rights Reserved.
// Code licenced under: BSD-2

package credentials

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
)

// NewGETCredentialsTaskResourcesCredentialsParams creates a new GETCredentialsTaskResourcesCredentialsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGETCredentialsTaskResourcesCredentialsParams() *GETCredentialsTaskResourcesCredentialsParams {
	return &GETCredentialsTaskResourcesCredentialsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGETCredentialsTaskResourcesCredentialsParamsWithTimeout creates a new GETCredentialsTaskResourcesCredentialsParams object
// with the ability to set a timeout on a request.
func NewGETCredentialsTaskResourcesCredentialsParamsWithTimeout(timeout time.Duration) *GETCredentialsTaskResourcesCredentialsParams {
	return &GETCredentialsTaskResourcesCredentialsParams{
		timeout: timeout,
	}
}

// NewGETCredentialsTaskResourcesCredentialsParamsWithContext creates a new GETCredentialsTaskResourcesCredentialsParams object
// with the ability to set a context for a request.
func NewGETCredentialsTaskResourcesCredentialsParamsWithContext(ctx context.Context) *GETCredentialsTaskResourcesCredentialsParams {
	return &GETCredentialsTaskResourcesCredentialsParams{
		Context: ctx,
	}
}

// NewGETCredentialsTaskResourcesCredentialsParamsWithHTTPClient creates a new GETCredentialsTaskResourcesCredentialsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGETCredentialsTaskResourcesCredentialsParamsWithHTTPClient(client *http.Client) *GETCredentialsTaskResourcesCredentialsParams {
	return &GETCredentialsTaskResourcesCredentialsParams{
		HTTPClient: client,
	}
}

/*
GETCredentialsTaskResourcesCredentialsParams contains all the parameters to send to the API endpoint

	for the get credentials task resources credentials operation.

	Typically these are written to a http.Request.
*/
type GETCredentialsTaskResourcesCredentialsParams struct {

	/* ID.

	   The ID of the credentials task
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get credentials task resources credentials params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GETCredentialsTaskResourcesCredentialsParams) WithDefaults() *GETCredentialsTaskResourcesCredentialsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get credentials task resources credentials params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GETCredentialsTaskResourcesCredentialsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get credentials task resources credentials params
func (o *GETCredentialsTaskResourcesCredentialsParams) WithTimeout(timeout time.Duration) *GETCredentialsTaskResourcesCredentialsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get credentials task resources credentials params
func (o *GETCredentialsTaskResourcesCredentialsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get credentials task resources credentials params
func (o *GETCredentialsTaskResourcesCredentialsParams) WithContext(ctx context.Context) *GETCredentialsTaskResourcesCredentialsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get credentials task resources credentials params
func (o *GETCredentialsTaskResourcesCredentialsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get credentials task resources credentials params
func (o *GETCredentialsTaskResourcesCredentialsParams) WithHTTPClient(client *http.Client) *GETCredentialsTaskResourcesCredentialsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get credentials task resources credentials params
func (o *GETCredentialsTaskResourcesCredentialsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get credentials task resources credentials params
func (o *GETCredentialsTaskResourcesCredentialsParams) WithID(id string) *GETCredentialsTaskResourcesCredentialsParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get credentials task resources credentials params
func (o *GETCredentialsTaskResourcesCredentialsParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GETCredentialsTaskResourcesCredentialsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
