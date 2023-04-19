// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2023 VMware, Inc.
// SPDX-License-Identifier: BSD-2-Clause

package system

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

// NewConfigureNtpParams creates a new ConfigureNtpParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewConfigureNtpParams() *ConfigureNtpParams {
	return &ConfigureNtpParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewConfigureNtpParamsWithTimeout creates a new ConfigureNtpParams object
// with the ability to set a timeout on a request.
func NewConfigureNtpParamsWithTimeout(timeout time.Duration) *ConfigureNtpParams {
	return &ConfigureNtpParams{
		timeout: timeout,
	}
}

// NewConfigureNtpParamsWithContext creates a new ConfigureNtpParams object
// with the ability to set a context for a request.
func NewConfigureNtpParamsWithContext(ctx context.Context) *ConfigureNtpParams {
	return &ConfigureNtpParams{
		Context: ctx,
	}
}

// NewConfigureNtpParamsWithHTTPClient creates a new ConfigureNtpParams object
// with the ability to set a custom HTTPClient for a request.
func NewConfigureNtpParamsWithHTTPClient(client *http.Client) *ConfigureNtpParams {
	return &ConfigureNtpParams{
		HTTPClient: client,
	}
}

/*
ConfigureNtpParams contains all the parameters to send to the API endpoint

	for the configure ntp operation.

	Typically these are written to a http.Request.
*/
type ConfigureNtpParams struct {

	/* NtpConfiguration.

	   ntpConfiguration
	*/
	NtpConfiguration *models.NtpConfiguration

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the configure ntp params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ConfigureNtpParams) WithDefaults() *ConfigureNtpParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the configure ntp params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ConfigureNtpParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the configure ntp params
func (o *ConfigureNtpParams) WithTimeout(timeout time.Duration) *ConfigureNtpParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the configure ntp params
func (o *ConfigureNtpParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the configure ntp params
func (o *ConfigureNtpParams) WithContext(ctx context.Context) *ConfigureNtpParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the configure ntp params
func (o *ConfigureNtpParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the configure ntp params
func (o *ConfigureNtpParams) WithHTTPClient(client *http.Client) *ConfigureNtpParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the configure ntp params
func (o *ConfigureNtpParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNtpConfiguration adds the ntpConfiguration to the configure ntp params
func (o *ConfigureNtpParams) WithNtpConfiguration(ntpConfiguration *models.NtpConfiguration) *ConfigureNtpParams {
	o.SetNtpConfiguration(ntpConfiguration)
	return o
}

// SetNtpConfiguration adds the ntpConfiguration to the configure ntp params
func (o *ConfigureNtpParams) SetNtpConfiguration(ntpConfiguration *models.NtpConfiguration) {
	o.NtpConfiguration = ntpConfiguration
}

// WriteToRequest writes these params to a swagger request
func (o *ConfigureNtpParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.NtpConfiguration != nil {
		if err := r.SetBodyParam(o.NtpConfiguration); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
