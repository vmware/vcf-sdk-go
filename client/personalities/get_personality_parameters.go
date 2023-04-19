// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2023 VMware, Inc.
// SPDX-License-Identifier: BSD-2-Clause

package personalities

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

// NewGETPersonalityParams creates a new GETPersonalityParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGETPersonalityParams() *GETPersonalityParams {
	return &GETPersonalityParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGETPersonalityParamsWithTimeout creates a new GETPersonalityParams object
// with the ability to set a timeout on a request.
func NewGETPersonalityParamsWithTimeout(timeout time.Duration) *GETPersonalityParams {
	return &GETPersonalityParams{
		timeout: timeout,
	}
}

// NewGETPersonalityParamsWithContext creates a new GETPersonalityParams object
// with the ability to set a context for a request.
func NewGETPersonalityParamsWithContext(ctx context.Context) *GETPersonalityParams {
	return &GETPersonalityParams{
		Context: ctx,
	}
}

// NewGETPersonalityParamsWithHTTPClient creates a new GETPersonalityParams object
// with the ability to set a custom HTTPClient for a request.
func NewGETPersonalityParamsWithHTTPClient(client *http.Client) *GETPersonalityParams {
	return &GETPersonalityParams{
		HTTPClient: client,
	}
}

/*
GETPersonalityParams contains all the parameters to send to the API endpoint

	for the get personality operation.

	Typically these are written to a http.Request.
*/
type GETPersonalityParams struct {

	/* PersonalityID.

	   Personality ID
	*/
	PersonalityID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get personality params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GETPersonalityParams) WithDefaults() *GETPersonalityParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get personality params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GETPersonalityParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get personality params
func (o *GETPersonalityParams) WithTimeout(timeout time.Duration) *GETPersonalityParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get personality params
func (o *GETPersonalityParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get personality params
func (o *GETPersonalityParams) WithContext(ctx context.Context) *GETPersonalityParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get personality params
func (o *GETPersonalityParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get personality params
func (o *GETPersonalityParams) WithHTTPClient(client *http.Client) *GETPersonalityParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get personality params
func (o *GETPersonalityParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPersonalityID adds the personalityID to the get personality params
func (o *GETPersonalityParams) WithPersonalityID(personalityID string) *GETPersonalityParams {
	o.SetPersonalityID(personalityID)
	return o
}

// SetPersonalityID adds the personalityId to the get personality params
func (o *GETPersonalityParams) SetPersonalityID(personalityID string) {
	o.PersonalityID = personalityID
}

// WriteToRequest writes these params to a swagger request
func (o *GETPersonalityParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param personalityId
	if err := r.SetPathParam("personalityId", o.PersonalityID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
