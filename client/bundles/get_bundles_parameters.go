// Code generated by go-swagger; DO NOT EDIT.

// Copyright © 2023 VMware, Inc. All Rights Reserved.
// Code licenced under: BSD-2

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
	"github.com/go-openapi/swag"
)

// NewGETBundlesParams creates a new GETBundlesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGETBundlesParams() *GETBundlesParams {
	return &GETBundlesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGETBundlesParamsWithTimeout creates a new GETBundlesParams object
// with the ability to set a timeout on a request.
func NewGETBundlesParamsWithTimeout(timeout time.Duration) *GETBundlesParams {
	return &GETBundlesParams{
		timeout: timeout,
	}
}

// NewGETBundlesParamsWithContext creates a new GETBundlesParams object
// with the ability to set a context for a request.
func NewGETBundlesParamsWithContext(ctx context.Context) *GETBundlesParams {
	return &GETBundlesParams{
		Context: ctx,
	}
}

// NewGETBundlesParamsWithHTTPClient creates a new GETBundlesParams object
// with the ability to set a custom HTTPClient for a request.
func NewGETBundlesParamsWithHTTPClient(client *http.Client) *GETBundlesParams {
	return &GETBundlesParams{
		HTTPClient: client,
	}
}

/*
GETBundlesParams contains all the parameters to send to the API endpoint

	for the get bundles operation.

	Typically these are written to a http.Request.
*/
type GETBundlesParams struct {

	/* BundleType.

	   The type of the bundle
	*/
	BundleType *string

	/* IsCompliant.

	   Is compliant with the current VCF version
	*/
	IsCompliant *bool

	/* ProductType.

	   The type of the product
	*/
	ProductType *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get bundles params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GETBundlesParams) WithDefaults() *GETBundlesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get bundles params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GETBundlesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get bundles params
func (o *GETBundlesParams) WithTimeout(timeout time.Duration) *GETBundlesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get bundles params
func (o *GETBundlesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get bundles params
func (o *GETBundlesParams) WithContext(ctx context.Context) *GETBundlesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get bundles params
func (o *GETBundlesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get bundles params
func (o *GETBundlesParams) WithHTTPClient(client *http.Client) *GETBundlesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get bundles params
func (o *GETBundlesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBundleType adds the bundleType to the get bundles params
func (o *GETBundlesParams) WithBundleType(bundleType *string) *GETBundlesParams {
	o.SetBundleType(bundleType)
	return o
}

// SetBundleType adds the bundleType to the get bundles params
func (o *GETBundlesParams) SetBundleType(bundleType *string) {
	o.BundleType = bundleType
}

// WithIsCompliant adds the isCompliant to the get bundles params
func (o *GETBundlesParams) WithIsCompliant(isCompliant *bool) *GETBundlesParams {
	o.SetIsCompliant(isCompliant)
	return o
}

// SetIsCompliant adds the isCompliant to the get bundles params
func (o *GETBundlesParams) SetIsCompliant(isCompliant *bool) {
	o.IsCompliant = isCompliant
}

// WithProductType adds the productType to the get bundles params
func (o *GETBundlesParams) WithProductType(productType *string) *GETBundlesParams {
	o.SetProductType(productType)
	return o
}

// SetProductType adds the productType to the get bundles params
func (o *GETBundlesParams) SetProductType(productType *string) {
	o.ProductType = productType
}

// WriteToRequest writes these params to a swagger request
func (o *GETBundlesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.BundleType != nil {

		// query param bundleType
		var qrBundleType string

		if o.BundleType != nil {
			qrBundleType = *o.BundleType
		}
		qBundleType := qrBundleType
		if qBundleType != "" {

			if err := r.SetQueryParam("bundleType", qBundleType); err != nil {
				return err
			}
		}
	}

	if o.IsCompliant != nil {

		// query param isCompliant
		var qrIsCompliant bool

		if o.IsCompliant != nil {
			qrIsCompliant = *o.IsCompliant
		}
		qIsCompliant := swag.FormatBool(qrIsCompliant)
		if qIsCompliant != "" {

			if err := r.SetQueryParam("isCompliant", qIsCompliant); err != nil {
				return err
			}
		}
	}

	if o.ProductType != nil {

		// query param productType
		var qrProductType string

		if o.ProductType != nil {
			qrProductType = *o.ProductType
		}
		qProductType := qrProductType
		if qProductType != "" {

			if err := r.SetQueryParam("productType", qProductType); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
