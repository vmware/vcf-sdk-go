// Code generated by go-swagger; DO NOT EDIT.

// Copyright © 2023 VMware, Inc. All Rights Reserved.
// Code licenced under: BSD-2

package vrslcm

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

// NewGETVRSLCMValidationParams creates a new GETVRSLCMValidationParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGETVRSLCMValidationParams() *GETVRSLCMValidationParams {
	return &GETVRSLCMValidationParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGETVRSLCMValidationParamsWithTimeout creates a new GETVRSLCMValidationParams object
// with the ability to set a timeout on a request.
func NewGETVRSLCMValidationParamsWithTimeout(timeout time.Duration) *GETVRSLCMValidationParams {
	return &GETVRSLCMValidationParams{
		timeout: timeout,
	}
}

// NewGETVRSLCMValidationParamsWithContext creates a new GETVRSLCMValidationParams object
// with the ability to set a context for a request.
func NewGETVRSLCMValidationParamsWithContext(ctx context.Context) *GETVRSLCMValidationParams {
	return &GETVRSLCMValidationParams{
		Context: ctx,
	}
}

// NewGETVRSLCMValidationParamsWithHTTPClient creates a new GETVRSLCMValidationParams object
// with the ability to set a custom HTTPClient for a request.
func NewGETVRSLCMValidationParamsWithHTTPClient(client *http.Client) *GETVRSLCMValidationParams {
	return &GETVRSLCMValidationParams{
		HTTPClient: client,
	}
}

/*
GETVRSLCMValidationParams contains all the parameters to send to the API endpoint

	for the get Vrslcm validation operation.

	Typically these are written to a http.Request.
*/
type GETVRSLCMValidationParams struct {

	/* ID.

	   vRealize validation id
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get Vrslcm validation params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GETVRSLCMValidationParams) WithDefaults() *GETVRSLCMValidationParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get Vrslcm validation params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GETVRSLCMValidationParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get Vrslcm validation params
func (o *GETVRSLCMValidationParams) WithTimeout(timeout time.Duration) *GETVRSLCMValidationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get Vrslcm validation params
func (o *GETVRSLCMValidationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get Vrslcm validation params
func (o *GETVRSLCMValidationParams) WithContext(ctx context.Context) *GETVRSLCMValidationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get Vrslcm validation params
func (o *GETVRSLCMValidationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get Vrslcm validation params
func (o *GETVRSLCMValidationParams) WithHTTPClient(client *http.Client) *GETVRSLCMValidationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get Vrslcm validation params
func (o *GETVRSLCMValidationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get Vrslcm validation params
func (o *GETVRSLCMValidationParams) WithID(id string) *GETVRSLCMValidationParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get Vrslcm validation params
func (o *GETVRSLCMValidationParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GETVRSLCMValidationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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