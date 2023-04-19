// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2023 VMware, Inc.
// SPDX-License-Identifier: BSD-2-Clause

package proxy_configuration

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

// NewUpdateProxyConfigurationParams creates a new UpdateProxyConfigurationParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateProxyConfigurationParams() *UpdateProxyConfigurationParams {
	return &UpdateProxyConfigurationParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateProxyConfigurationParamsWithTimeout creates a new UpdateProxyConfigurationParams object
// with the ability to set a timeout on a request.
func NewUpdateProxyConfigurationParamsWithTimeout(timeout time.Duration) *UpdateProxyConfigurationParams {
	return &UpdateProxyConfigurationParams{
		timeout: timeout,
	}
}

// NewUpdateProxyConfigurationParamsWithContext creates a new UpdateProxyConfigurationParams object
// with the ability to set a context for a request.
func NewUpdateProxyConfigurationParamsWithContext(ctx context.Context) *UpdateProxyConfigurationParams {
	return &UpdateProxyConfigurationParams{
		Context: ctx,
	}
}

// NewUpdateProxyConfigurationParamsWithHTTPClient creates a new UpdateProxyConfigurationParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateProxyConfigurationParamsWithHTTPClient(client *http.Client) *UpdateProxyConfigurationParams {
	return &UpdateProxyConfigurationParams{
		HTTPClient: client,
	}
}

/*
UpdateProxyConfigurationParams contains all the parameters to send to the API endpoint

	for the update proxy configuration operation.

	Typically these are written to a http.Request.
*/
type UpdateProxyConfigurationParams struct {

	/* ProxyConfig.

	   proxyConfig
	*/
	ProxyConfig *models.ProxyConfiguration

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update proxy configuration params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateProxyConfigurationParams) WithDefaults() *UpdateProxyConfigurationParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update proxy configuration params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateProxyConfigurationParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update proxy configuration params
func (o *UpdateProxyConfigurationParams) WithTimeout(timeout time.Duration) *UpdateProxyConfigurationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update proxy configuration params
func (o *UpdateProxyConfigurationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update proxy configuration params
func (o *UpdateProxyConfigurationParams) WithContext(ctx context.Context) *UpdateProxyConfigurationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update proxy configuration params
func (o *UpdateProxyConfigurationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update proxy configuration params
func (o *UpdateProxyConfigurationParams) WithHTTPClient(client *http.Client) *UpdateProxyConfigurationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update proxy configuration params
func (o *UpdateProxyConfigurationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithProxyConfig adds the proxyConfig to the update proxy configuration params
func (o *UpdateProxyConfigurationParams) WithProxyConfig(proxyConfig *models.ProxyConfiguration) *UpdateProxyConfigurationParams {
	o.SetProxyConfig(proxyConfig)
	return o
}

// SetProxyConfig adds the proxyConfig to the update proxy configuration params
func (o *UpdateProxyConfigurationParams) SetProxyConfig(proxyConfig *models.ProxyConfiguration) {
	o.ProxyConfig = proxyConfig
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateProxyConfigurationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.ProxyConfig != nil {
		if err := r.SetBodyParam(o.ProxyConfig); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
