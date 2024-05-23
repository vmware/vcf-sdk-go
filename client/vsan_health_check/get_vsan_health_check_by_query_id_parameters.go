// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2023 VMware, Inc.
// SPDX-License-Identifier: BSD-2-Clause

package vsan_health_check

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

// NewGetVSANHealthCheckByQueryIDParams creates a new GetVSANHealthCheckByQueryIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetVSANHealthCheckByQueryIDParams() *GetVSANHealthCheckByQueryIDParams {
	return &GetVSANHealthCheckByQueryIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetVSANHealthCheckByQueryIDParamsWithTimeout creates a new GetVSANHealthCheckByQueryIDParams object
// with the ability to set a timeout on a request.
func NewGetVSANHealthCheckByQueryIDParamsWithTimeout(timeout time.Duration) *GetVSANHealthCheckByQueryIDParams {
	return &GetVSANHealthCheckByQueryIDParams{
		timeout: timeout,
	}
}

// NewGetVSANHealthCheckByQueryIDParamsWithContext creates a new GetVSANHealthCheckByQueryIDParams object
// with the ability to set a context for a request.
func NewGetVSANHealthCheckByQueryIDParamsWithContext(ctx context.Context) *GetVSANHealthCheckByQueryIDParams {
	return &GetVSANHealthCheckByQueryIDParams{
		Context: ctx,
	}
}

// NewGetVSANHealthCheckByQueryIDParamsWithHTTPClient creates a new GetVSANHealthCheckByQueryIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetVSANHealthCheckByQueryIDParamsWithHTTPClient(client *http.Client) *GetVSANHealthCheckByQueryIDParams {
	return &GetVSANHealthCheckByQueryIDParams{
		HTTPClient: client,
	}
}

/*
GetVSANHealthCheckByQueryIDParams contains all the parameters to send to the API endpoint

	for the get Vsan health check by query ID operation.

	Typically these are written to a http.Request.
*/
type GetVSANHealthCheckByQueryIDParams struct {

	/* DomainID.

	   Domain ID
	*/
	DomainID string

	/* QueryID.

	   Query ID
	*/
	QueryID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get Vsan health check by query ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetVSANHealthCheckByQueryIDParams) WithDefaults() *GetVSANHealthCheckByQueryIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get Vsan health check by query ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetVSANHealthCheckByQueryIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get Vsan health check by query ID params
func (o *GetVSANHealthCheckByQueryIDParams) WithTimeout(timeout time.Duration) *GetVSANHealthCheckByQueryIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get Vsan health check by query ID params
func (o *GetVSANHealthCheckByQueryIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get Vsan health check by query ID params
func (o *GetVSANHealthCheckByQueryIDParams) WithContext(ctx context.Context) *GetVSANHealthCheckByQueryIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get Vsan health check by query ID params
func (o *GetVSANHealthCheckByQueryIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get Vsan health check by query ID params
func (o *GetVSANHealthCheckByQueryIDParams) WithHTTPClient(client *http.Client) *GetVSANHealthCheckByQueryIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get Vsan health check by query ID params
func (o *GetVSANHealthCheckByQueryIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDomainID adds the domainID to the get Vsan health check by query ID params
func (o *GetVSANHealthCheckByQueryIDParams) WithDomainID(domainID string) *GetVSANHealthCheckByQueryIDParams {
	o.SetDomainID(domainID)
	return o
}

// SetDomainID adds the domainId to the get Vsan health check by query ID params
func (o *GetVSANHealthCheckByQueryIDParams) SetDomainID(domainID string) {
	o.DomainID = domainID
}

// WithQueryID adds the queryID to the get Vsan health check by query ID params
func (o *GetVSANHealthCheckByQueryIDParams) WithQueryID(queryID string) *GetVSANHealthCheckByQueryIDParams {
	o.SetQueryID(queryID)
	return o
}

// SetQueryID adds the queryId to the get Vsan health check by query ID params
func (o *GetVSANHealthCheckByQueryIDParams) SetQueryID(queryID string) {
	o.QueryID = queryID
}

// WriteToRequest writes these params to a swagger request
func (o *GetVSANHealthCheckByQueryIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param domainId
	if err := r.SetPathParam("domainId", o.DomainID); err != nil {
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
