// Code generated by go-swagger; DO NOT EDIT.

// Copyright © 2023 VMware, Inc. All Rights Reserved.
// Code licenced under: BSD-2

package n_s_x_t_clusters

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

// NewGETNSXTClusterParams creates a new GETNSXTClusterParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGETNSXTClusterParams() *GETNSXTClusterParams {
	return &GETNSXTClusterParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGETNSXTClusterParamsWithTimeout creates a new GETNSXTClusterParams object
// with the ability to set a timeout on a request.
func NewGETNSXTClusterParamsWithTimeout(timeout time.Duration) *GETNSXTClusterParams {
	return &GETNSXTClusterParams{
		timeout: timeout,
	}
}

// NewGETNSXTClusterParamsWithContext creates a new GETNSXTClusterParams object
// with the ability to set a context for a request.
func NewGETNSXTClusterParamsWithContext(ctx context.Context) *GETNSXTClusterParams {
	return &GETNSXTClusterParams{
		Context: ctx,
	}
}

// NewGETNSXTClusterParamsWithHTTPClient creates a new GETNSXTClusterParams object
// with the ability to set a custom HTTPClient for a request.
func NewGETNSXTClusterParamsWithHTTPClient(client *http.Client) *GETNSXTClusterParams {
	return &GETNSXTClusterParams{
		HTTPClient: client,
	}
}

/*
GETNSXTClusterParams contains all the parameters to send to the API endpoint

	for the get Nsxt cluster operation.

	Typically these are written to a http.Request.
*/
type GETNSXTClusterParams struct {

	/* ID.

	   NSX-T cluster ID
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get Nsxt cluster params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GETNSXTClusterParams) WithDefaults() *GETNSXTClusterParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get Nsxt cluster params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GETNSXTClusterParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get Nsxt cluster params
func (o *GETNSXTClusterParams) WithTimeout(timeout time.Duration) *GETNSXTClusterParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get Nsxt cluster params
func (o *GETNSXTClusterParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get Nsxt cluster params
func (o *GETNSXTClusterParams) WithContext(ctx context.Context) *GETNSXTClusterParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get Nsxt cluster params
func (o *GETNSXTClusterParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get Nsxt cluster params
func (o *GETNSXTClusterParams) WithHTTPClient(client *http.Client) *GETNSXTClusterParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get Nsxt cluster params
func (o *GETNSXTClusterParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get Nsxt cluster params
func (o *GETNSXTClusterParams) WithID(id string) *GETNSXTClusterParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get Nsxt cluster params
func (o *GETNSXTClusterParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GETNSXTClusterParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
