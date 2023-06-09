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

	"github.com/vmware/vcf-sdk-go/models"
)

// NewAddIPPoolToNetworkOfNetworkPoolParams creates a new AddIPPoolToNetworkOfNetworkPoolParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAddIPPoolToNetworkOfNetworkPoolParams() *AddIPPoolToNetworkOfNetworkPoolParams {
	return &AddIPPoolToNetworkOfNetworkPoolParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAddIPPoolToNetworkOfNetworkPoolParamsWithTimeout creates a new AddIPPoolToNetworkOfNetworkPoolParams object
// with the ability to set a timeout on a request.
func NewAddIPPoolToNetworkOfNetworkPoolParamsWithTimeout(timeout time.Duration) *AddIPPoolToNetworkOfNetworkPoolParams {
	return &AddIPPoolToNetworkOfNetworkPoolParams{
		timeout: timeout,
	}
}

// NewAddIPPoolToNetworkOfNetworkPoolParamsWithContext creates a new AddIPPoolToNetworkOfNetworkPoolParams object
// with the ability to set a context for a request.
func NewAddIPPoolToNetworkOfNetworkPoolParamsWithContext(ctx context.Context) *AddIPPoolToNetworkOfNetworkPoolParams {
	return &AddIPPoolToNetworkOfNetworkPoolParams{
		Context: ctx,
	}
}

// NewAddIPPoolToNetworkOfNetworkPoolParamsWithHTTPClient creates a new AddIPPoolToNetworkOfNetworkPoolParams object
// with the ability to set a custom HTTPClient for a request.
func NewAddIPPoolToNetworkOfNetworkPoolParamsWithHTTPClient(client *http.Client) *AddIPPoolToNetworkOfNetworkPoolParams {
	return &AddIPPoolToNetworkOfNetworkPoolParams{
		HTTPClient: client,
	}
}

/*
AddIPPoolToNetworkOfNetworkPoolParams contains all the parameters to send to the API endpoint

	for the add Ip pool to network of network pool operation.

	Typically these are written to a http.Request.
*/
type AddIPPoolToNetworkOfNetworkPoolParams struct {

	/* ID.

	   Id of the networkpoolk
	*/
	ID string

	/* IPPool.

	   ipPool
	*/
	IPPool *models.IPPool

	/* NetworkID.

	   Id of the network
	*/
	NetworkID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the add Ip pool to network of network pool params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AddIPPoolToNetworkOfNetworkPoolParams) WithDefaults() *AddIPPoolToNetworkOfNetworkPoolParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the add Ip pool to network of network pool params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AddIPPoolToNetworkOfNetworkPoolParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the add Ip pool to network of network pool params
func (o *AddIPPoolToNetworkOfNetworkPoolParams) WithTimeout(timeout time.Duration) *AddIPPoolToNetworkOfNetworkPoolParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the add Ip pool to network of network pool params
func (o *AddIPPoolToNetworkOfNetworkPoolParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the add Ip pool to network of network pool params
func (o *AddIPPoolToNetworkOfNetworkPoolParams) WithContext(ctx context.Context) *AddIPPoolToNetworkOfNetworkPoolParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the add Ip pool to network of network pool params
func (o *AddIPPoolToNetworkOfNetworkPoolParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the add Ip pool to network of network pool params
func (o *AddIPPoolToNetworkOfNetworkPoolParams) WithHTTPClient(client *http.Client) *AddIPPoolToNetworkOfNetworkPoolParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the add Ip pool to network of network pool params
func (o *AddIPPoolToNetworkOfNetworkPoolParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the add Ip pool to network of network pool params
func (o *AddIPPoolToNetworkOfNetworkPoolParams) WithID(id string) *AddIPPoolToNetworkOfNetworkPoolParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the add Ip pool to network of network pool params
func (o *AddIPPoolToNetworkOfNetworkPoolParams) SetID(id string) {
	o.ID = id
}

// WithIPPool adds the iPPool to the add Ip pool to network of network pool params
func (o *AddIPPoolToNetworkOfNetworkPoolParams) WithIPPool(iPPool *models.IPPool) *AddIPPoolToNetworkOfNetworkPoolParams {
	o.SetIPPool(iPPool)
	return o
}

// SetIPPool adds the ipPool to the add Ip pool to network of network pool params
func (o *AddIPPoolToNetworkOfNetworkPoolParams) SetIPPool(iPPool *models.IPPool) {
	o.IPPool = iPPool
}

// WithNetworkID adds the networkID to the add Ip pool to network of network pool params
func (o *AddIPPoolToNetworkOfNetworkPoolParams) WithNetworkID(networkID string) *AddIPPoolToNetworkOfNetworkPoolParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the add Ip pool to network of network pool params
func (o *AddIPPoolToNetworkOfNetworkPoolParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WriteToRequest writes these params to a swagger request
func (o *AddIPPoolToNetworkOfNetworkPoolParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}
	if o.IPPool != nil {
		if err := r.SetBodyParam(o.IPPool); err != nil {
			return err
		}
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
