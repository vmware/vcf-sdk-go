// Code generated by go-swagger; DO NOT EDIT.

// Copyright © 2023 VMware, Inc. All Rights Reserved.
// Code licenced under: BSD-2

package hosts

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

// NewAssignableTagsToHostParams creates a new AssignableTagsToHostParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAssignableTagsToHostParams() *AssignableTagsToHostParams {
	return &AssignableTagsToHostParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAssignableTagsToHostParamsWithTimeout creates a new AssignableTagsToHostParams object
// with the ability to set a timeout on a request.
func NewAssignableTagsToHostParamsWithTimeout(timeout time.Duration) *AssignableTagsToHostParams {
	return &AssignableTagsToHostParams{
		timeout: timeout,
	}
}

// NewAssignableTagsToHostParamsWithContext creates a new AssignableTagsToHostParams object
// with the ability to set a context for a request.
func NewAssignableTagsToHostParamsWithContext(ctx context.Context) *AssignableTagsToHostParams {
	return &AssignableTagsToHostParams{
		Context: ctx,
	}
}

// NewAssignableTagsToHostParamsWithHTTPClient creates a new AssignableTagsToHostParams object
// with the ability to set a custom HTTPClient for a request.
func NewAssignableTagsToHostParamsWithHTTPClient(client *http.Client) *AssignableTagsToHostParams {
	return &AssignableTagsToHostParams{
		HTTPClient: client,
	}
}

/*
AssignableTagsToHostParams contains all the parameters to send to the API endpoint

	for the assignable tags to host operation.

	Typically these are written to a http.Request.
*/
type AssignableTagsToHostParams struct {

	/* ID.

	   Host ID
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the assignable tags to host params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AssignableTagsToHostParams) WithDefaults() *AssignableTagsToHostParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the assignable tags to host params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AssignableTagsToHostParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the assignable tags to host params
func (o *AssignableTagsToHostParams) WithTimeout(timeout time.Duration) *AssignableTagsToHostParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the assignable tags to host params
func (o *AssignableTagsToHostParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the assignable tags to host params
func (o *AssignableTagsToHostParams) WithContext(ctx context.Context) *AssignableTagsToHostParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the assignable tags to host params
func (o *AssignableTagsToHostParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the assignable tags to host params
func (o *AssignableTagsToHostParams) WithHTTPClient(client *http.Client) *AssignableTagsToHostParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the assignable tags to host params
func (o *AssignableTagsToHostParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the assignable tags to host params
func (o *AssignableTagsToHostParams) WithID(id string) *AssignableTagsToHostParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the assignable tags to host params
func (o *AssignableTagsToHostParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *AssignableTagsToHostParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
