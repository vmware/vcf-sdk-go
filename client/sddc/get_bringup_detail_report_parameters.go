// Code generated by go-swagger; DO NOT EDIT.

// Copyright © 2023 VMware, Inc. All Rights Reserved.
// Code licenced under: BSD-2

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

// NewGETBringupDetailReportParams creates a new GETBringupDetailReportParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGETBringupDetailReportParams() *GETBringupDetailReportParams {
	return &GETBringupDetailReportParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGETBringupDetailReportParamsWithTimeout creates a new GETBringupDetailReportParams object
// with the ability to set a timeout on a request.
func NewGETBringupDetailReportParamsWithTimeout(timeout time.Duration) *GETBringupDetailReportParams {
	return &GETBringupDetailReportParams{
		timeout: timeout,
	}
}

// NewGETBringupDetailReportParamsWithContext creates a new GETBringupDetailReportParams object
// with the ability to set a context for a request.
func NewGETBringupDetailReportParamsWithContext(ctx context.Context) *GETBringupDetailReportParams {
	return &GETBringupDetailReportParams{
		Context: ctx,
	}
}

// NewGETBringupDetailReportParamsWithHTTPClient creates a new GETBringupDetailReportParams object
// with the ability to set a custom HTTPClient for a request.
func NewGETBringupDetailReportParamsWithHTTPClient(client *http.Client) *GETBringupDetailReportParams {
	return &GETBringupDetailReportParams{
		HTTPClient: client,
	}
}

/*
GETBringupDetailReportParams contains all the parameters to send to the API endpoint

	for the get bringup detail report operation.

	Typically these are written to a http.Request.
*/
type GETBringupDetailReportParams struct {

	/* Format.

	   One among: PDF, CSV

	   Default: "PDF"
	*/
	Format *string

	/* ID.

	   SDDC ID
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get bringup detail report params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GETBringupDetailReportParams) WithDefaults() *GETBringupDetailReportParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get bringup detail report params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GETBringupDetailReportParams) SetDefaults() {
	var (
		formatDefault = string("PDF")
	)

	val := GETBringupDetailReportParams{
		Format: &formatDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get bringup detail report params
func (o *GETBringupDetailReportParams) WithTimeout(timeout time.Duration) *GETBringupDetailReportParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get bringup detail report params
func (o *GETBringupDetailReportParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get bringup detail report params
func (o *GETBringupDetailReportParams) WithContext(ctx context.Context) *GETBringupDetailReportParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get bringup detail report params
func (o *GETBringupDetailReportParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get bringup detail report params
func (o *GETBringupDetailReportParams) WithHTTPClient(client *http.Client) *GETBringupDetailReportParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get bringup detail report params
func (o *GETBringupDetailReportParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFormat adds the format to the get bringup detail report params
func (o *GETBringupDetailReportParams) WithFormat(format *string) *GETBringupDetailReportParams {
	o.SetFormat(format)
	return o
}

// SetFormat adds the format to the get bringup detail report params
func (o *GETBringupDetailReportParams) SetFormat(format *string) {
	o.Format = format
}

// WithID adds the id to the get bringup detail report params
func (o *GETBringupDetailReportParams) WithID(id string) *GETBringupDetailReportParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get bringup detail report params
func (o *GETBringupDetailReportParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GETBringupDetailReportParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Format != nil {

		// query param format
		var qrFormat string

		if o.Format != nil {
			qrFormat = *o.Format
		}
		qFormat := qrFormat
		if qFormat != "" {

			if err := r.SetQueryParam("format", qFormat); err != nil {
				return err
			}
		}
	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}