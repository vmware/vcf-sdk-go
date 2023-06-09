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

	"github.com/vmware/vcf-sdk-go/models"
)

// NewAddLicenseKeyParams creates a new AddLicenseKeyParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAddLicenseKeyParams() *AddLicenseKeyParams {
	return &AddLicenseKeyParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAddLicenseKeyParamsWithTimeout creates a new AddLicenseKeyParams object
// with the ability to set a timeout on a request.
func NewAddLicenseKeyParamsWithTimeout(timeout time.Duration) *AddLicenseKeyParams {
	return &AddLicenseKeyParams{
		timeout: timeout,
	}
}

// NewAddLicenseKeyParamsWithContext creates a new AddLicenseKeyParams object
// with the ability to set a context for a request.
func NewAddLicenseKeyParamsWithContext(ctx context.Context) *AddLicenseKeyParams {
	return &AddLicenseKeyParams{
		Context: ctx,
	}
}

// NewAddLicenseKeyParamsWithHTTPClient creates a new AddLicenseKeyParams object
// with the ability to set a custom HTTPClient for a request.
func NewAddLicenseKeyParamsWithHTTPClient(client *http.Client) *AddLicenseKeyParams {
	return &AddLicenseKeyParams{
		HTTPClient: client,
	}
}

/*
AddLicenseKeyParams contains all the parameters to send to the API endpoint

	for the add license key operation.

	Typically these are written to a http.Request.
*/
type AddLicenseKeyParams struct {

	/* LicenseKey.

	   License key with other attributes
	*/
	LicenseKey *models.LicenseKey

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the add license key params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AddLicenseKeyParams) WithDefaults() *AddLicenseKeyParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the add license key params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AddLicenseKeyParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the add license key params
func (o *AddLicenseKeyParams) WithTimeout(timeout time.Duration) *AddLicenseKeyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the add license key params
func (o *AddLicenseKeyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the add license key params
func (o *AddLicenseKeyParams) WithContext(ctx context.Context) *AddLicenseKeyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the add license key params
func (o *AddLicenseKeyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the add license key params
func (o *AddLicenseKeyParams) WithHTTPClient(client *http.Client) *AddLicenseKeyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the add license key params
func (o *AddLicenseKeyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLicenseKey adds the licenseKey to the add license key params
func (o *AddLicenseKeyParams) WithLicenseKey(licenseKey *models.LicenseKey) *AddLicenseKeyParams {
	o.SetLicenseKey(licenseKey)
	return o
}

// SetLicenseKey adds the licenseKey to the add license key params
func (o *AddLicenseKeyParams) SetLicenseKey(licenseKey *models.LicenseKey) {
	o.LicenseKey = licenseKey
}

// WriteToRequest writes these params to a swagger request
func (o *AddLicenseKeyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.LicenseKey != nil {
		if err := r.SetBodyParam(o.LicenseKey); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
