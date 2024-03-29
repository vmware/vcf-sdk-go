// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2023 VMware, Inc.
// SPDX-License-Identifier: BSD-2-Clause

package license_keys

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

// NewGetDomainLicensingInfoParams creates a new GetDomainLicensingInfoParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetDomainLicensingInfoParams() *GetDomainLicensingInfoParams {
	return &GetDomainLicensingInfoParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetDomainLicensingInfoParamsWithTimeout creates a new GetDomainLicensingInfoParams object
// with the ability to set a timeout on a request.
func NewGetDomainLicensingInfoParamsWithTimeout(timeout time.Duration) *GetDomainLicensingInfoParams {
	return &GetDomainLicensingInfoParams{
		timeout: timeout,
	}
}

// NewGetDomainLicensingInfoParamsWithContext creates a new GetDomainLicensingInfoParams object
// with the ability to set a context for a request.
func NewGetDomainLicensingInfoParamsWithContext(ctx context.Context) *GetDomainLicensingInfoParams {
	return &GetDomainLicensingInfoParams{
		Context: ctx,
	}
}

// NewGetDomainLicensingInfoParamsWithHTTPClient creates a new GetDomainLicensingInfoParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetDomainLicensingInfoParamsWithHTTPClient(client *http.Client) *GetDomainLicensingInfoParams {
	return &GetDomainLicensingInfoParams{
		HTTPClient: client,
	}
}

/*
GetDomainLicensingInfoParams contains all the parameters to send to the API endpoint

	for the get domain licensing info operation.

	Typically these are written to a http.Request.
*/
type GetDomainLicensingInfoParams struct {

	/* ID.

	   The domain ID
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get domain licensing info params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDomainLicensingInfoParams) WithDefaults() *GetDomainLicensingInfoParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get domain licensing info params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDomainLicensingInfoParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get domain licensing info params
func (o *GetDomainLicensingInfoParams) WithTimeout(timeout time.Duration) *GetDomainLicensingInfoParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get domain licensing info params
func (o *GetDomainLicensingInfoParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get domain licensing info params
func (o *GetDomainLicensingInfoParams) WithContext(ctx context.Context) *GetDomainLicensingInfoParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get domain licensing info params
func (o *GetDomainLicensingInfoParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get domain licensing info params
func (o *GetDomainLicensingInfoParams) WithHTTPClient(client *http.Client) *GetDomainLicensingInfoParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get domain licensing info params
func (o *GetDomainLicensingInfoParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get domain licensing info params
func (o *GetDomainLicensingInfoParams) WithID(id string) *GetDomainLicensingInfoParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get domain licensing info params
func (o *GetDomainLicensingInfoParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetDomainLicensingInfoParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
