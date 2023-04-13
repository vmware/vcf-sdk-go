// Code generated by go-swagger; DO NOT EDIT.

// Copyright © 2023 VMware, Inc. All Rights Reserved.
// Code licenced under: BSD-2

package v_r_o_ps

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

// NewGETIntegratedDomainsVROPSParams creates a new GETIntegratedDomainsVROPSParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGETIntegratedDomainsVROPSParams() *GETIntegratedDomainsVROPSParams {
	return &GETIntegratedDomainsVROPSParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGETIntegratedDomainsVROPSParamsWithTimeout creates a new GETIntegratedDomainsVROPSParams object
// with the ability to set a timeout on a request.
func NewGETIntegratedDomainsVROPSParamsWithTimeout(timeout time.Duration) *GETIntegratedDomainsVROPSParams {
	return &GETIntegratedDomainsVROPSParams{
		timeout: timeout,
	}
}

// NewGETIntegratedDomainsVROPSParamsWithContext creates a new GETIntegratedDomainsVROPSParams object
// with the ability to set a context for a request.
func NewGETIntegratedDomainsVROPSParamsWithContext(ctx context.Context) *GETIntegratedDomainsVROPSParams {
	return &GETIntegratedDomainsVROPSParams{
		Context: ctx,
	}
}

// NewGETIntegratedDomainsVROPSParamsWithHTTPClient creates a new GETIntegratedDomainsVROPSParams object
// with the ability to set a custom HTTPClient for a request.
func NewGETIntegratedDomainsVROPSParamsWithHTTPClient(client *http.Client) *GETIntegratedDomainsVROPSParams {
	return &GETIntegratedDomainsVROPSParams{
		HTTPClient: client,
	}
}

/*
GETIntegratedDomainsVROPSParams contains all the parameters to send to the API endpoint

	for the get integrated domains Vrops operation.

	Typically these are written to a http.Request.
*/
type GETIntegratedDomainsVROPSParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get integrated domains Vrops params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GETIntegratedDomainsVROPSParams) WithDefaults() *GETIntegratedDomainsVROPSParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get integrated domains Vrops params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GETIntegratedDomainsVROPSParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get integrated domains Vrops params
func (o *GETIntegratedDomainsVROPSParams) WithTimeout(timeout time.Duration) *GETIntegratedDomainsVROPSParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get integrated domains Vrops params
func (o *GETIntegratedDomainsVROPSParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get integrated domains Vrops params
func (o *GETIntegratedDomainsVROPSParams) WithContext(ctx context.Context) *GETIntegratedDomainsVROPSParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get integrated domains Vrops params
func (o *GETIntegratedDomainsVROPSParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get integrated domains Vrops params
func (o *GETIntegratedDomainsVROPSParams) WithHTTPClient(client *http.Client) *GETIntegratedDomainsVROPSParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get integrated domains Vrops params
func (o *GETIntegratedDomainsVROPSParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GETIntegratedDomainsVROPSParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
