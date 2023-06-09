// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2023 VMware, Inc.
// SPDX-License-Identifier: BSD-2-Clause

package bundles

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

// NewUploadBundleParams creates a new UploadBundleParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUploadBundleParams() *UploadBundleParams {
	return &UploadBundleParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUploadBundleParamsWithTimeout creates a new UploadBundleParams object
// with the ability to set a timeout on a request.
func NewUploadBundleParamsWithTimeout(timeout time.Duration) *UploadBundleParams {
	return &UploadBundleParams{
		timeout: timeout,
	}
}

// NewUploadBundleParamsWithContext creates a new UploadBundleParams object
// with the ability to set a context for a request.
func NewUploadBundleParamsWithContext(ctx context.Context) *UploadBundleParams {
	return &UploadBundleParams{
		Context: ctx,
	}
}

// NewUploadBundleParamsWithHTTPClient creates a new UploadBundleParams object
// with the ability to set a custom HTTPClient for a request.
func NewUploadBundleParamsWithHTTPClient(client *http.Client) *UploadBundleParams {
	return &UploadBundleParams{
		HTTPClient: client,
	}
}

/*
UploadBundleParams contains all the parameters to send to the API endpoint

	for the upload bundle operation.

	Typically these are written to a http.Request.
*/
type UploadBundleParams struct {

	/* BundleUploadSpec.

	   Bundle Upload Specification
	*/
	BundleUploadSpec *models.BundleUploadSpec

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the upload bundle params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UploadBundleParams) WithDefaults() *UploadBundleParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the upload bundle params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UploadBundleParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the upload bundle params
func (o *UploadBundleParams) WithTimeout(timeout time.Duration) *UploadBundleParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the upload bundle params
func (o *UploadBundleParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the upload bundle params
func (o *UploadBundleParams) WithContext(ctx context.Context) *UploadBundleParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the upload bundle params
func (o *UploadBundleParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the upload bundle params
func (o *UploadBundleParams) WithHTTPClient(client *http.Client) *UploadBundleParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the upload bundle params
func (o *UploadBundleParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBundleUploadSpec adds the bundleUploadSpec to the upload bundle params
func (o *UploadBundleParams) WithBundleUploadSpec(bundleUploadSpec *models.BundleUploadSpec) *UploadBundleParams {
	o.SetBundleUploadSpec(bundleUploadSpec)
	return o
}

// SetBundleUploadSpec adds the bundleUploadSpec to the upload bundle params
func (o *UploadBundleParams) SetBundleUploadSpec(bundleUploadSpec *models.BundleUploadSpec) {
	o.BundleUploadSpec = bundleUploadSpec
}

// WriteToRequest writes these params to a swagger request
func (o *UploadBundleParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.BundleUploadSpec != nil {
		if err := r.SetBodyParam(o.BundleUploadSpec); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
