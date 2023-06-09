// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2023 VMware, Inc.
// SPDX-License-Identifier: BSD-2-Clause

package network_pools

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

// NewGetNetworkOfNetworkPoolParams creates a new GetNetworkOfNetworkPoolParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetNetworkOfNetworkPoolParams() *GetNetworkOfNetworkPoolParams {
	return &GetNetworkOfNetworkPoolParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetNetworkOfNetworkPoolParamsWithTimeout creates a new GetNetworkOfNetworkPoolParams object
// with the ability to set a timeout on a request.
func NewGetNetworkOfNetworkPoolParamsWithTimeout(timeout time.Duration) *GetNetworkOfNetworkPoolParams {
	return &GetNetworkOfNetworkPoolParams{
		timeout: timeout,
	}
}

// NewGetNetworkOfNetworkPoolParamsWithContext creates a new GetNetworkOfNetworkPoolParams object
// with the ability to set a context for a request.
func NewGetNetworkOfNetworkPoolParamsWithContext(ctx context.Context) *GetNetworkOfNetworkPoolParams {
	return &GetNetworkOfNetworkPoolParams{
		Context: ctx,
	}
}

// NewGetNetworkOfNetworkPoolParamsWithHTTPClient creates a new GetNetworkOfNetworkPoolParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetNetworkOfNetworkPoolParamsWithHTTPClient(client *http.Client) *GetNetworkOfNetworkPoolParams {
	return &GetNetworkOfNetworkPoolParams{
		HTTPClient: client,
	}
}

/*
GetNetworkOfNetworkPoolParams contains all the parameters to send to the API endpoint

	for the get network of network pool operation.

	Typically these are written to a http.Request.
*/
type GetNetworkOfNetworkPoolParams struct {

	/* ID.

	   Id of the Network pool
	*/
	ID string

	/* NetworkID.

	   Id of the Network
	*/
	NetworkID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get network of network pool params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetNetworkOfNetworkPoolParams) WithDefaults() *GetNetworkOfNetworkPoolParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get network of network pool params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetNetworkOfNetworkPoolParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get network of network pool params
func (o *GetNetworkOfNetworkPoolParams) WithTimeout(timeout time.Duration) *GetNetworkOfNetworkPoolParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get network of network pool params
func (o *GetNetworkOfNetworkPoolParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get network of network pool params
func (o *GetNetworkOfNetworkPoolParams) WithContext(ctx context.Context) *GetNetworkOfNetworkPoolParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get network of network pool params
func (o *GetNetworkOfNetworkPoolParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get network of network pool params
func (o *GetNetworkOfNetworkPoolParams) WithHTTPClient(client *http.Client) *GetNetworkOfNetworkPoolParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get network of network pool params
func (o *GetNetworkOfNetworkPoolParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get network of network pool params
func (o *GetNetworkOfNetworkPoolParams) WithID(id string) *GetNetworkOfNetworkPoolParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get network of network pool params
func (o *GetNetworkOfNetworkPoolParams) SetID(id string) {
	o.ID = id
}

// WithNetworkID adds the networkID to the get network of network pool params
func (o *GetNetworkOfNetworkPoolParams) WithNetworkID(networkID string) *GetNetworkOfNetworkPoolParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the get network of network pool params
func (o *GetNetworkOfNetworkPoolParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WriteToRequest writes these params to a swagger request
func (o *GetNetworkOfNetworkPoolParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	// path param networkId
	if err := r.SetPathParam("networkId", o.NetworkID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
