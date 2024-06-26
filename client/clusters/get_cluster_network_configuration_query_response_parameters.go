// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2023 VMware, Inc.
// SPDX-License-Identifier: BSD-2-Clause

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

// NewGetClusterNetworkConfigurationQueryResponseParams creates a new GetClusterNetworkConfigurationQueryResponseParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetClusterNetworkConfigurationQueryResponseParams() *GetClusterNetworkConfigurationQueryResponseParams {
	return &GetClusterNetworkConfigurationQueryResponseParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetClusterNetworkConfigurationQueryResponseParamsWithTimeout creates a new GetClusterNetworkConfigurationQueryResponseParams object
// with the ability to set a timeout on a request.
func NewGetClusterNetworkConfigurationQueryResponseParamsWithTimeout(timeout time.Duration) *GetClusterNetworkConfigurationQueryResponseParams {
	return &GetClusterNetworkConfigurationQueryResponseParams{
		timeout: timeout,
	}
}

// NewGetClusterNetworkConfigurationQueryResponseParamsWithContext creates a new GetClusterNetworkConfigurationQueryResponseParams object
// with the ability to set a context for a request.
func NewGetClusterNetworkConfigurationQueryResponseParamsWithContext(ctx context.Context) *GetClusterNetworkConfigurationQueryResponseParams {
	return &GetClusterNetworkConfigurationQueryResponseParams{
		Context: ctx,
	}
}

// NewGetClusterNetworkConfigurationQueryResponseParamsWithHTTPClient creates a new GetClusterNetworkConfigurationQueryResponseParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetClusterNetworkConfigurationQueryResponseParamsWithHTTPClient(client *http.Client) *GetClusterNetworkConfigurationQueryResponseParams {
	return &GetClusterNetworkConfigurationQueryResponseParams{
		HTTPClient: client,
	}
}

/*
GetClusterNetworkConfigurationQueryResponseParams contains all the parameters to send to the API endpoint

	for the get cluster network configuration query response operation.

	Typically these are written to a http.Request.
*/
type GetClusterNetworkConfigurationQueryResponseParams struct {

	/* ID.

	   Cluster ID
	*/
	ID string

	/* QueryID.

	   Query ID
	*/
	QueryID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get cluster network configuration query response params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetClusterNetworkConfigurationQueryResponseParams) WithDefaults() *GetClusterNetworkConfigurationQueryResponseParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get cluster network configuration query response params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetClusterNetworkConfigurationQueryResponseParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get cluster network configuration query response params
func (o *GetClusterNetworkConfigurationQueryResponseParams) WithTimeout(timeout time.Duration) *GetClusterNetworkConfigurationQueryResponseParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get cluster network configuration query response params
func (o *GetClusterNetworkConfigurationQueryResponseParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get cluster network configuration query response params
func (o *GetClusterNetworkConfigurationQueryResponseParams) WithContext(ctx context.Context) *GetClusterNetworkConfigurationQueryResponseParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get cluster network configuration query response params
func (o *GetClusterNetworkConfigurationQueryResponseParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get cluster network configuration query response params
func (o *GetClusterNetworkConfigurationQueryResponseParams) WithHTTPClient(client *http.Client) *GetClusterNetworkConfigurationQueryResponseParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get cluster network configuration query response params
func (o *GetClusterNetworkConfigurationQueryResponseParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get cluster network configuration query response params
func (o *GetClusterNetworkConfigurationQueryResponseParams) WithID(id string) *GetClusterNetworkConfigurationQueryResponseParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get cluster network configuration query response params
func (o *GetClusterNetworkConfigurationQueryResponseParams) SetID(id string) {
	o.ID = id
}

// WithQueryID adds the queryID to the get cluster network configuration query response params
func (o *GetClusterNetworkConfigurationQueryResponseParams) WithQueryID(queryID string) *GetClusterNetworkConfigurationQueryResponseParams {
	o.SetQueryID(queryID)
	return o
}

// SetQueryID adds the queryId to the get cluster network configuration query response params
func (o *GetClusterNetworkConfigurationQueryResponseParams) SetQueryID(queryID string) {
	o.QueryID = queryID
}

// WriteToRequest writes these params to a swagger request
func (o *GetClusterNetworkConfigurationQueryResponseParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	// path param queryId
	if err := r.SetPathParam("queryId", o.QueryID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
