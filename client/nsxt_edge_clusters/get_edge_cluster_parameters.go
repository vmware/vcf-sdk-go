// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2023 VMware, Inc.
// SPDX-License-Identifier: BSD-2-Clause

package nsxt_edge_clusters

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

// NewGetEdgeClusterParams creates a new GetEdgeClusterParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetEdgeClusterParams() *GetEdgeClusterParams {
	return &GetEdgeClusterParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetEdgeClusterParamsWithTimeout creates a new GetEdgeClusterParams object
// with the ability to set a timeout on a request.
func NewGetEdgeClusterParamsWithTimeout(timeout time.Duration) *GetEdgeClusterParams {
	return &GetEdgeClusterParams{
		timeout: timeout,
	}
}

// NewGetEdgeClusterParamsWithContext creates a new GetEdgeClusterParams object
// with the ability to set a context for a request.
func NewGetEdgeClusterParamsWithContext(ctx context.Context) *GetEdgeClusterParams {
	return &GetEdgeClusterParams{
		Context: ctx,
	}
}

// NewGetEdgeClusterParamsWithHTTPClient creates a new GetEdgeClusterParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetEdgeClusterParamsWithHTTPClient(client *http.Client) *GetEdgeClusterParams {
	return &GetEdgeClusterParams{
		HTTPClient: client,
	}
}

/*
GetEdgeClusterParams contains all the parameters to send to the API endpoint

	for the get edge cluster operation.

	Typically these are written to a http.Request.
*/
type GetEdgeClusterParams struct {

	/* ID.

	   Edge Cluster ID
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get edge cluster params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetEdgeClusterParams) WithDefaults() *GetEdgeClusterParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get edge cluster params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetEdgeClusterParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get edge cluster params
func (o *GetEdgeClusterParams) WithTimeout(timeout time.Duration) *GetEdgeClusterParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get edge cluster params
func (o *GetEdgeClusterParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get edge cluster params
func (o *GetEdgeClusterParams) WithContext(ctx context.Context) *GetEdgeClusterParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get edge cluster params
func (o *GetEdgeClusterParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get edge cluster params
func (o *GetEdgeClusterParams) WithHTTPClient(client *http.Client) *GetEdgeClusterParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get edge cluster params
func (o *GetEdgeClusterParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get edge cluster params
func (o *GetEdgeClusterParams) WithID(id string) *GetEdgeClusterParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get edge cluster params
func (o *GetEdgeClusterParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetEdgeClusterParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
