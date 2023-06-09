// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2023 VMware, Inc.
// SPDX-License-Identifier: BSD-2-Clause

package backup_restore

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

// NewRestoresTasksParams creates a new RestoresTasksParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewRestoresTasksParams() *RestoresTasksParams {
	return &RestoresTasksParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewRestoresTasksParamsWithTimeout creates a new RestoresTasksParams object
// with the ability to set a timeout on a request.
func NewRestoresTasksParamsWithTimeout(timeout time.Duration) *RestoresTasksParams {
	return &RestoresTasksParams{
		timeout: timeout,
	}
}

// NewRestoresTasksParamsWithContext creates a new RestoresTasksParams object
// with the ability to set a context for a request.
func NewRestoresTasksParamsWithContext(ctx context.Context) *RestoresTasksParams {
	return &RestoresTasksParams{
		Context: ctx,
	}
}

// NewRestoresTasksParamsWithHTTPClient creates a new RestoresTasksParams object
// with the ability to set a custom HTTPClient for a request.
func NewRestoresTasksParamsWithHTTPClient(client *http.Client) *RestoresTasksParams {
	return &RestoresTasksParams{
		HTTPClient: client,
	}
}

/*
RestoresTasksParams contains all the parameters to send to the API endpoint

	for the restores tasks operation.

	Typically these are written to a http.Request.
*/
type RestoresTasksParams struct {

	/* ID.

	   The restore task ID
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the restores tasks params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RestoresTasksParams) WithDefaults() *RestoresTasksParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the restores tasks params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RestoresTasksParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the restores tasks params
func (o *RestoresTasksParams) WithTimeout(timeout time.Duration) *RestoresTasksParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the restores tasks params
func (o *RestoresTasksParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the restores tasks params
func (o *RestoresTasksParams) WithContext(ctx context.Context) *RestoresTasksParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the restores tasks params
func (o *RestoresTasksParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the restores tasks params
func (o *RestoresTasksParams) WithHTTPClient(client *http.Client) *RestoresTasksParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the restores tasks params
func (o *RestoresTasksParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the restores tasks params
func (o *RestoresTasksParams) WithID(id string) *RestoresTasksParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the restores tasks params
func (o *RestoresTasksParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *RestoresTasksParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
