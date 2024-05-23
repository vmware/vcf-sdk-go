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

// NewGetClusterNetworkConfigurationCriteriaParams creates a new GetClusterNetworkConfigurationCriteriaParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetClusterNetworkConfigurationCriteriaParams() *GetClusterNetworkConfigurationCriteriaParams {
	return &GetClusterNetworkConfigurationCriteriaParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetClusterNetworkConfigurationCriteriaParamsWithTimeout creates a new GetClusterNetworkConfigurationCriteriaParams object
// with the ability to set a timeout on a request.
func NewGetClusterNetworkConfigurationCriteriaParamsWithTimeout(timeout time.Duration) *GetClusterNetworkConfigurationCriteriaParams {
	return &GetClusterNetworkConfigurationCriteriaParams{
		timeout: timeout,
	}
}

// NewGetClusterNetworkConfigurationCriteriaParamsWithContext creates a new GetClusterNetworkConfigurationCriteriaParams object
// with the ability to set a context for a request.
func NewGetClusterNetworkConfigurationCriteriaParamsWithContext(ctx context.Context) *GetClusterNetworkConfigurationCriteriaParams {
	return &GetClusterNetworkConfigurationCriteriaParams{
		Context: ctx,
	}
}

// NewGetClusterNetworkConfigurationCriteriaParamsWithHTTPClient creates a new GetClusterNetworkConfigurationCriteriaParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetClusterNetworkConfigurationCriteriaParamsWithHTTPClient(client *http.Client) *GetClusterNetworkConfigurationCriteriaParams {
	return &GetClusterNetworkConfigurationCriteriaParams{
		HTTPClient: client,
	}
}

/*
GetClusterNetworkConfigurationCriteriaParams contains all the parameters to send to the API endpoint

	for the get cluster network configuration criteria operation.

	Typically these are written to a http.Request.
*/
type GetClusterNetworkConfigurationCriteriaParams struct {

	/* ID.

	   Cluster ID
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get cluster network configuration criteria params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetClusterNetworkConfigurationCriteriaParams) WithDefaults() *GetClusterNetworkConfigurationCriteriaParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get cluster network configuration criteria params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetClusterNetworkConfigurationCriteriaParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get cluster network configuration criteria params
func (o *GetClusterNetworkConfigurationCriteriaParams) WithTimeout(timeout time.Duration) *GetClusterNetworkConfigurationCriteriaParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get cluster network configuration criteria params
func (o *GetClusterNetworkConfigurationCriteriaParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get cluster network configuration criteria params
func (o *GetClusterNetworkConfigurationCriteriaParams) WithContext(ctx context.Context) *GetClusterNetworkConfigurationCriteriaParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get cluster network configuration criteria params
func (o *GetClusterNetworkConfigurationCriteriaParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get cluster network configuration criteria params
func (o *GetClusterNetworkConfigurationCriteriaParams) WithHTTPClient(client *http.Client) *GetClusterNetworkConfigurationCriteriaParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get cluster network configuration criteria params
func (o *GetClusterNetworkConfigurationCriteriaParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get cluster network configuration criteria params
func (o *GetClusterNetworkConfigurationCriteriaParams) WithID(id string) *GetClusterNetworkConfigurationCriteriaParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get cluster network configuration criteria params
func (o *GetClusterNetworkConfigurationCriteriaParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetClusterNetworkConfigurationCriteriaParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
