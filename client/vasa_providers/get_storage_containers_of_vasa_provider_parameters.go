// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2023 VMware, Inc.
// SPDX-License-Identifier: BSD-2-Clause

package vasa_providers

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

// NewGetStorageContainersOfVasaProviderParams creates a new GetStorageContainersOfVasaProviderParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetStorageContainersOfVasaProviderParams() *GetStorageContainersOfVasaProviderParams {
	return &GetStorageContainersOfVasaProviderParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetStorageContainersOfVasaProviderParamsWithTimeout creates a new GetStorageContainersOfVasaProviderParams object
// with the ability to set a timeout on a request.
func NewGetStorageContainersOfVasaProviderParamsWithTimeout(timeout time.Duration) *GetStorageContainersOfVasaProviderParams {
	return &GetStorageContainersOfVasaProviderParams{
		timeout: timeout,
	}
}

// NewGetStorageContainersOfVasaProviderParamsWithContext creates a new GetStorageContainersOfVasaProviderParams object
// with the ability to set a context for a request.
func NewGetStorageContainersOfVasaProviderParamsWithContext(ctx context.Context) *GetStorageContainersOfVasaProviderParams {
	return &GetStorageContainersOfVasaProviderParams{
		Context: ctx,
	}
}

// NewGetStorageContainersOfVasaProviderParamsWithHTTPClient creates a new GetStorageContainersOfVasaProviderParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetStorageContainersOfVasaProviderParamsWithHTTPClient(client *http.Client) *GetStorageContainersOfVasaProviderParams {
	return &GetStorageContainersOfVasaProviderParams{
		HTTPClient: client,
	}
}

/*
GetStorageContainersOfVasaProviderParams contains all the parameters to send to the API endpoint

	for the get storage containers of vasa provider operation.

	Typically these are written to a http.Request.
*/
type GetStorageContainersOfVasaProviderParams struct {

	/* ID.

	   VASA Provider ID
	*/
	ID string

	/* ProtocolType.

	   Pass an optional Storage Protocol type
	*/
	ProtocolType *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get storage containers of vasa provider params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetStorageContainersOfVasaProviderParams) WithDefaults() *GetStorageContainersOfVasaProviderParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get storage containers of vasa provider params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetStorageContainersOfVasaProviderParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get storage containers of vasa provider params
func (o *GetStorageContainersOfVasaProviderParams) WithTimeout(timeout time.Duration) *GetStorageContainersOfVasaProviderParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get storage containers of vasa provider params
func (o *GetStorageContainersOfVasaProviderParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get storage containers of vasa provider params
func (o *GetStorageContainersOfVasaProviderParams) WithContext(ctx context.Context) *GetStorageContainersOfVasaProviderParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get storage containers of vasa provider params
func (o *GetStorageContainersOfVasaProviderParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get storage containers of vasa provider params
func (o *GetStorageContainersOfVasaProviderParams) WithHTTPClient(client *http.Client) *GetStorageContainersOfVasaProviderParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get storage containers of vasa provider params
func (o *GetStorageContainersOfVasaProviderParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get storage containers of vasa provider params
func (o *GetStorageContainersOfVasaProviderParams) WithID(id string) *GetStorageContainersOfVasaProviderParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get storage containers of vasa provider params
func (o *GetStorageContainersOfVasaProviderParams) SetID(id string) {
	o.ID = id
}

// WithProtocolType adds the protocolType to the get storage containers of vasa provider params
func (o *GetStorageContainersOfVasaProviderParams) WithProtocolType(protocolType *string) *GetStorageContainersOfVasaProviderParams {
	o.SetProtocolType(protocolType)
	return o
}

// SetProtocolType adds the protocolType to the get storage containers of vasa provider params
func (o *GetStorageContainersOfVasaProviderParams) SetProtocolType(protocolType *string) {
	o.ProtocolType = protocolType
}

// WriteToRequest writes these params to a swagger request
func (o *GetStorageContainersOfVasaProviderParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if o.ProtocolType != nil {

		// query param protocolType
		var qrProtocolType string

		if o.ProtocolType != nil {
			qrProtocolType = *o.ProtocolType
		}
		qProtocolType := qrProtocolType
		if qProtocolType != "" {

			if err := r.SetQueryParam("protocolType", qProtocolType); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
