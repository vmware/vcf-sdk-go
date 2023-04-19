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
	"github.com/go-openapi/swag"
)

// NewGETLicenseKeysParams creates a new GETLicenseKeysParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGETLicenseKeysParams() *GETLicenseKeysParams {
	return &GETLicenseKeysParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGETLicenseKeysParamsWithTimeout creates a new GETLicenseKeysParams object
// with the ability to set a timeout on a request.
func NewGETLicenseKeysParamsWithTimeout(timeout time.Duration) *GETLicenseKeysParams {
	return &GETLicenseKeysParams{
		timeout: timeout,
	}
}

// NewGETLicenseKeysParamsWithContext creates a new GETLicenseKeysParams object
// with the ability to set a context for a request.
func NewGETLicenseKeysParamsWithContext(ctx context.Context) *GETLicenseKeysParams {
	return &GETLicenseKeysParams{
		Context: ctx,
	}
}

// NewGETLicenseKeysParamsWithHTTPClient creates a new GETLicenseKeysParams object
// with the ability to set a custom HTTPClient for a request.
func NewGETLicenseKeysParamsWithHTTPClient(client *http.Client) *GETLicenseKeysParams {
	return &GETLicenseKeysParams{
		HTTPClient: client,
	}
}

/*
GETLicenseKeysParams contains all the parameters to send to the API endpoint

	for the get license keys operation.

	Typically these are written to a http.Request.
*/
type GETLicenseKeysParams struct {

	/* LicenseKeyStatus.

	   Status of a License Key
	*/
	LicenseKeyStatus []string

	/* ProductType.

	   Type of a Product
	*/
	ProductType []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get license keys params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GETLicenseKeysParams) WithDefaults() *GETLicenseKeysParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get license keys params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GETLicenseKeysParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get license keys params
func (o *GETLicenseKeysParams) WithTimeout(timeout time.Duration) *GETLicenseKeysParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get license keys params
func (o *GETLicenseKeysParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get license keys params
func (o *GETLicenseKeysParams) WithContext(ctx context.Context) *GETLicenseKeysParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get license keys params
func (o *GETLicenseKeysParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get license keys params
func (o *GETLicenseKeysParams) WithHTTPClient(client *http.Client) *GETLicenseKeysParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get license keys params
func (o *GETLicenseKeysParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLicenseKeyStatus adds the licenseKeyStatus to the get license keys params
func (o *GETLicenseKeysParams) WithLicenseKeyStatus(licenseKeyStatus []string) *GETLicenseKeysParams {
	o.SetLicenseKeyStatus(licenseKeyStatus)
	return o
}

// SetLicenseKeyStatus adds the licenseKeyStatus to the get license keys params
func (o *GETLicenseKeysParams) SetLicenseKeyStatus(licenseKeyStatus []string) {
	o.LicenseKeyStatus = licenseKeyStatus
}

// WithProductType adds the productType to the get license keys params
func (o *GETLicenseKeysParams) WithProductType(productType []string) *GETLicenseKeysParams {
	o.SetProductType(productType)
	return o
}

// SetProductType adds the productType to the get license keys params
func (o *GETLicenseKeysParams) SetProductType(productType []string) {
	o.ProductType = productType
}

// WriteToRequest writes these params to a swagger request
func (o *GETLicenseKeysParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.LicenseKeyStatus != nil {

		// binding items for licenseKeyStatus
		joinedLicenseKeyStatus := o.bindParamLicenseKeyStatus(reg)

		// query array param licenseKeyStatus
		if err := r.SetQueryParam("licenseKeyStatus", joinedLicenseKeyStatus...); err != nil {
			return err
		}
	}

	if o.ProductType != nil {

		// binding items for productType
		joinedProductType := o.bindParamProductType(reg)

		// query array param productType
		if err := r.SetQueryParam("productType", joinedProductType...); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamGETLicenseKeys binds the parameter licenseKeyStatus
func (o *GETLicenseKeysParams) bindParamLicenseKeyStatus(formats strfmt.Registry) []string {
	licenseKeyStatusIR := o.LicenseKeyStatus

	var licenseKeyStatusIC []string
	for _, licenseKeyStatusIIR := range licenseKeyStatusIR { // explode []string

		licenseKeyStatusIIV := licenseKeyStatusIIR // string as string
		licenseKeyStatusIC = append(licenseKeyStatusIC, licenseKeyStatusIIV)
	}

	// items.CollectionFormat: "multi"
	licenseKeyStatusIS := swag.JoinByFormat(licenseKeyStatusIC, "multi")

	return licenseKeyStatusIS
}

// bindParamGETLicenseKeys binds the parameter productType
func (o *GETLicenseKeysParams) bindParamProductType(formats strfmt.Registry) []string {
	productTypeIR := o.ProductType

	var productTypeIC []string
	for _, productTypeIIR := range productTypeIR { // explode []string

		productTypeIIV := productTypeIIR // string as string
		productTypeIC = append(productTypeIC, productTypeIIV)
	}

	// items.CollectionFormat: "multi"
	productTypeIS := swag.JoinByFormat(productTypeIC, "multi")

	return productTypeIS
}
