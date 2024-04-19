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

// NewSetLicenseKeyForResourceParams creates a new SetLicenseKeyForResourceParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSetLicenseKeyForResourceParams() *SetLicenseKeyForResourceParams {
	return &SetLicenseKeyForResourceParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSetLicenseKeyForResourceParamsWithTimeout creates a new SetLicenseKeyForResourceParams object
// with the ability to set a timeout on a request.
func NewSetLicenseKeyForResourceParamsWithTimeout(timeout time.Duration) *SetLicenseKeyForResourceParams {
	return &SetLicenseKeyForResourceParams{
		timeout: timeout,
	}
}

// NewSetLicenseKeyForResourceParamsWithContext creates a new SetLicenseKeyForResourceParams object
// with the ability to set a context for a request.
func NewSetLicenseKeyForResourceParamsWithContext(ctx context.Context) *SetLicenseKeyForResourceParams {
	return &SetLicenseKeyForResourceParams{
		Context: ctx,
	}
}

// NewSetLicenseKeyForResourceParamsWithHTTPClient creates a new SetLicenseKeyForResourceParams object
// with the ability to set a custom HTTPClient for a request.
func NewSetLicenseKeyForResourceParamsWithHTTPClient(client *http.Client) *SetLicenseKeyForResourceParams {
	return &SetLicenseKeyForResourceParams{
		HTTPClient: client,
	}
}

/*
SetLicenseKeyForResourceParams contains all the parameters to send to the API endpoint

	for the set license key for resource operation.

	Typically these are written to a http.Request.
*/
type SetLicenseKeyForResourceParams struct {

	/* LicensingSpec.

	   License Information of resources
	*/
	LicensingSpec *models.LicensingSpec

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the set license key for resource params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SetLicenseKeyForResourceParams) WithDefaults() *SetLicenseKeyForResourceParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the set license key for resource params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SetLicenseKeyForResourceParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the set license key for resource params
func (o *SetLicenseKeyForResourceParams) WithTimeout(timeout time.Duration) *SetLicenseKeyForResourceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the set license key for resource params
func (o *SetLicenseKeyForResourceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the set license key for resource params
func (o *SetLicenseKeyForResourceParams) WithContext(ctx context.Context) *SetLicenseKeyForResourceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the set license key for resource params
func (o *SetLicenseKeyForResourceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the set license key for resource params
func (o *SetLicenseKeyForResourceParams) WithHTTPClient(client *http.Client) *SetLicenseKeyForResourceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the set license key for resource params
func (o *SetLicenseKeyForResourceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLicensingSpec adds the licensingSpec to the set license key for resource params
func (o *SetLicenseKeyForResourceParams) WithLicensingSpec(licensingSpec *models.LicensingSpec) *SetLicenseKeyForResourceParams {
	o.SetLicensingSpec(licensingSpec)
	return o
}

// SetLicensingSpec adds the licensingSpec to the set license key for resource params
func (o *SetLicenseKeyForResourceParams) SetLicensingSpec(licensingSpec *models.LicensingSpec) {
	o.LicensingSpec = licensingSpec
}

// WriteToRequest writes these params to a swagger request
func (o *SetLicenseKeyForResourceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.LicensingSpec != nil {
		if err := r.SetBodyParam(o.LicensingSpec); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}