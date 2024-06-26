// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2023 VMware, Inc.
// SPDX-License-Identifier: BSD-2-Clause

package nsxt_clusters

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

// NewGetNsxTransportZonesParams creates a new GetNsxTransportZonesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetNsxTransportZonesParams() *GetNsxTransportZonesParams {
	return &GetNsxTransportZonesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetNsxTransportZonesParamsWithTimeout creates a new GetNsxTransportZonesParams object
// with the ability to set a timeout on a request.
func NewGetNsxTransportZonesParamsWithTimeout(timeout time.Duration) *GetNsxTransportZonesParams {
	return &GetNsxTransportZonesParams{
		timeout: timeout,
	}
}

// NewGetNsxTransportZonesParamsWithContext creates a new GetNsxTransportZonesParams object
// with the ability to set a context for a request.
func NewGetNsxTransportZonesParamsWithContext(ctx context.Context) *GetNsxTransportZonesParams {
	return &GetNsxTransportZonesParams{
		Context: ctx,
	}
}

// NewGetNsxTransportZonesParamsWithHTTPClient creates a new GetNsxTransportZonesParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetNsxTransportZonesParamsWithHTTPClient(client *http.Client) *GetNsxTransportZonesParams {
	return &GetNsxTransportZonesParams{
		HTTPClient: client,
	}
}

/*
GetNsxTransportZonesParams contains all the parameters to send to the API endpoint

	for the get nsx transport zones operation.

	Typically these are written to a http.Request.
*/
type GetNsxTransportZonesParams struct {

	/* NSXTClusterID.

	   NSX cluster ID
	*/
	NSXTClusterID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get nsx transport zones params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetNsxTransportZonesParams) WithDefaults() *GetNsxTransportZonesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get nsx transport zones params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetNsxTransportZonesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get nsx transport zones params
func (o *GetNsxTransportZonesParams) WithTimeout(timeout time.Duration) *GetNsxTransportZonesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get nsx transport zones params
func (o *GetNsxTransportZonesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get nsx transport zones params
func (o *GetNsxTransportZonesParams) WithContext(ctx context.Context) *GetNsxTransportZonesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get nsx transport zones params
func (o *GetNsxTransportZonesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get nsx transport zones params
func (o *GetNsxTransportZonesParams) WithHTTPClient(client *http.Client) *GetNsxTransportZonesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get nsx transport zones params
func (o *GetNsxTransportZonesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNSXTClusterID adds the nSXTClusterID to the get nsx transport zones params
func (o *GetNsxTransportZonesParams) WithNSXTClusterID(nSXTClusterID string) *GetNsxTransportZonesParams {
	o.SetNSXTClusterID(nSXTClusterID)
	return o
}

// SetNSXTClusterID adds the nsxtClusterId to the get nsx transport zones params
func (o *GetNsxTransportZonesParams) SetNSXTClusterID(nSXTClusterID string) {
	o.NSXTClusterID = nSXTClusterID
}

// WriteToRequest writes these params to a swagger request
func (o *GetNsxTransportZonesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param nsxt-cluster-id
	if err := r.SetPathParam("nsxt-cluster-id", o.NSXTClusterID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
