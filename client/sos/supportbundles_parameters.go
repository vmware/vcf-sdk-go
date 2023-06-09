// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2023 VMware, Inc.
// SPDX-License-Identifier: BSD-2-Clause

package sos

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

// NewSupportbundlesParams creates a new SupportbundlesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSupportbundlesParams() *SupportbundlesParams {
	return &SupportbundlesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSupportbundlesParamsWithTimeout creates a new SupportbundlesParams object
// with the ability to set a timeout on a request.
func NewSupportbundlesParamsWithTimeout(timeout time.Duration) *SupportbundlesParams {
	return &SupportbundlesParams{
		timeout: timeout,
	}
}

// NewSupportbundlesParamsWithContext creates a new SupportbundlesParams object
// with the ability to set a context for a request.
func NewSupportbundlesParamsWithContext(ctx context.Context) *SupportbundlesParams {
	return &SupportbundlesParams{
		Context: ctx,
	}
}

// NewSupportbundlesParamsWithHTTPClient creates a new SupportbundlesParams object
// with the ability to set a custom HTTPClient for a request.
func NewSupportbundlesParamsWithHTTPClient(client *http.Client) *SupportbundlesParams {
	return &SupportbundlesParams{
		HTTPClient: client,
	}
}

/*
SupportbundlesParams contains all the parameters to send to the API endpoint

	for the supportbundles operation.

	Typically these are written to a http.Request.
*/
type SupportbundlesParams struct {

	/* Supportbundlespec.

	   supportbundlespec
	*/
	Supportbundlespec *models.SupportBundleSpec

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the supportbundles params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SupportbundlesParams) WithDefaults() *SupportbundlesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the supportbundles params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SupportbundlesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the supportbundles params
func (o *SupportbundlesParams) WithTimeout(timeout time.Duration) *SupportbundlesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the supportbundles params
func (o *SupportbundlesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the supportbundles params
func (o *SupportbundlesParams) WithContext(ctx context.Context) *SupportbundlesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the supportbundles params
func (o *SupportbundlesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the supportbundles params
func (o *SupportbundlesParams) WithHTTPClient(client *http.Client) *SupportbundlesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the supportbundles params
func (o *SupportbundlesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSupportbundlespec adds the supportbundlespec to the supportbundles params
func (o *SupportbundlesParams) WithSupportbundlespec(supportbundlespec *models.SupportBundleSpec) *SupportbundlesParams {
	o.SetSupportbundlespec(supportbundlespec)
	return o
}

// SetSupportbundlespec adds the supportbundlespec to the supportbundles params
func (o *SupportbundlesParams) SetSupportbundlespec(supportbundlespec *models.SupportBundleSpec) {
	o.Supportbundlespec = supportbundlespec
}

// WriteToRequest writes these params to a swagger request
func (o *SupportbundlesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Supportbundlespec != nil {
		if err := r.SetBodyParam(o.Supportbundlespec); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
