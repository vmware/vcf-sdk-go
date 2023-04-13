// Code generated by go-swagger; DO NOT EDIT.

// Copyright © 2023 VMware, Inc. All Rights Reserved.
// Code licenced under: BSD-2

package clusters

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

// NewGETVdsesParams creates a new GETVdsesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGETVdsesParams() *GETVdsesParams {
	return &GETVdsesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGETVdsesParamsWithTimeout creates a new GETVdsesParams object
// with the ability to set a timeout on a request.
func NewGETVdsesParamsWithTimeout(timeout time.Duration) *GETVdsesParams {
	return &GETVdsesParams{
		timeout: timeout,
	}
}

// NewGETVdsesParamsWithContext creates a new GETVdsesParams object
// with the ability to set a context for a request.
func NewGETVdsesParamsWithContext(ctx context.Context) *GETVdsesParams {
	return &GETVdsesParams{
		Context: ctx,
	}
}

// NewGETVdsesParamsWithHTTPClient creates a new GETVdsesParams object
// with the ability to set a custom HTTPClient for a request.
func NewGETVdsesParamsWithHTTPClient(client *http.Client) *GETVdsesParams {
	return &GETVdsesParams{
		HTTPClient: client,
	}
}

/*
GETVdsesParams contains all the parameters to send to the API endpoint

	for the get vdses operation.

	Typically these are written to a http.Request.
*/
type GETVdsesParams struct {

	/* ClusterID.

	   Cluster ID
	*/
	ClusterID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get vdses params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GETVdsesParams) WithDefaults() *GETVdsesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get vdses params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GETVdsesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get vdses params
func (o *GETVdsesParams) WithTimeout(timeout time.Duration) *GETVdsesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get vdses params
func (o *GETVdsesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get vdses params
func (o *GETVdsesParams) WithContext(ctx context.Context) *GETVdsesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get vdses params
func (o *GETVdsesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get vdses params
func (o *GETVdsesParams) WithHTTPClient(client *http.Client) *GETVdsesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get vdses params
func (o *GETVdsesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClusterID adds the clusterID to the get vdses params
func (o *GETVdsesParams) WithClusterID(clusterID string) *GETVdsesParams {
	o.SetClusterID(clusterID)
	return o
}

// SetClusterID adds the clusterId to the get vdses params
func (o *GETVdsesParams) SetClusterID(clusterID string) {
	o.ClusterID = clusterID
}

// WriteToRequest writes these params to a swagger request
func (o *GETVdsesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param clusterId
	if err := r.SetPathParam("clusterId", o.ClusterID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
