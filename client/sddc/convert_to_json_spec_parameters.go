// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2023 VMware, Inc.
// SPDX-License-Identifier: BSD-2-Clause

package sddc

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

// NewConvertToJSONSpecParams creates a new ConvertToJSONSpecParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewConvertToJSONSpecParams() *ConvertToJSONSpecParams {
	return &ConvertToJSONSpecParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewConvertToJSONSpecParamsWithTimeout creates a new ConvertToJSONSpecParams object
// with the ability to set a timeout on a request.
func NewConvertToJSONSpecParamsWithTimeout(timeout time.Duration) *ConvertToJSONSpecParams {
	return &ConvertToJSONSpecParams{
		timeout: timeout,
	}
}

// NewConvertToJSONSpecParamsWithContext creates a new ConvertToJSONSpecParams object
// with the ability to set a context for a request.
func NewConvertToJSONSpecParamsWithContext(ctx context.Context) *ConvertToJSONSpecParams {
	return &ConvertToJSONSpecParams{
		Context: ctx,
	}
}

// NewConvertToJSONSpecParamsWithHTTPClient creates a new ConvertToJSONSpecParams object
// with the ability to set a custom HTTPClient for a request.
func NewConvertToJSONSpecParamsWithHTTPClient(client *http.Client) *ConvertToJSONSpecParams {
	return &ConvertToJSONSpecParams{
		HTTPClient: client,
	}
}

/*
ConvertToJSONSpecParams contains all the parameters to send to the API endpoint

	for the convert to Json spec operation.

	Typically these are written to a http.Request.
*/
type ConvertToJSONSpecParams struct {

	/* Design.

	   Supported bringup designs - EMS, VXRAIL

	   Default: "EMS"
	*/
	Design *string

	/* SpecFile.

	   SDDC specification file which is either a JSON or xls file
	*/
	SpecFile runtime.NamedReadCloser

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the convert to Json spec params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ConvertToJSONSpecParams) WithDefaults() *ConvertToJSONSpecParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the convert to Json spec params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ConvertToJSONSpecParams) SetDefaults() {
	var (
		designDefault = string("EMS")
	)

	val := ConvertToJSONSpecParams{
		Design: &designDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the convert to Json spec params
func (o *ConvertToJSONSpecParams) WithTimeout(timeout time.Duration) *ConvertToJSONSpecParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the convert to Json spec params
func (o *ConvertToJSONSpecParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the convert to Json spec params
func (o *ConvertToJSONSpecParams) WithContext(ctx context.Context) *ConvertToJSONSpecParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the convert to Json spec params
func (o *ConvertToJSONSpecParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the convert to Json spec params
func (o *ConvertToJSONSpecParams) WithHTTPClient(client *http.Client) *ConvertToJSONSpecParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the convert to Json spec params
func (o *ConvertToJSONSpecParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDesign adds the design to the convert to Json spec params
func (o *ConvertToJSONSpecParams) WithDesign(design *string) *ConvertToJSONSpecParams {
	o.SetDesign(design)
	return o
}

// SetDesign adds the design to the convert to Json spec params
func (o *ConvertToJSONSpecParams) SetDesign(design *string) {
	o.Design = design
}

// WithSpecFile adds the specFile to the convert to Json spec params
func (o *ConvertToJSONSpecParams) WithSpecFile(specFile runtime.NamedReadCloser) *ConvertToJSONSpecParams {
	o.SetSpecFile(specFile)
	return o
}

// SetSpecFile adds the specFile to the convert to Json spec params
func (o *ConvertToJSONSpecParams) SetSpecFile(specFile runtime.NamedReadCloser) {
	o.SpecFile = specFile
}

// WriteToRequest writes these params to a swagger request
func (o *ConvertToJSONSpecParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Design != nil {

		// query param design
		var qrDesign string

		if o.Design != nil {
			qrDesign = *o.Design
		}
		qDesign := qrDesign
		if qDesign != "" {

			if err := r.SetQueryParam("design", qDesign); err != nil {
				return err
			}
		}
	}
	// form file param specFile
	if err := r.SetFileParam("specFile", o.SpecFile); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
