// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2023 VMware, Inc.
// SPDX-License-Identifier: BSD-2-Clause

package hosts

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

// NewValidateHostsOperationsParams creates a new ValidateHostsOperationsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewValidateHostsOperationsParams() *ValidateHostsOperationsParams {
	return &ValidateHostsOperationsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewValidateHostsOperationsParamsWithTimeout creates a new ValidateHostsOperationsParams object
// with the ability to set a timeout on a request.
func NewValidateHostsOperationsParamsWithTimeout(timeout time.Duration) *ValidateHostsOperationsParams {
	return &ValidateHostsOperationsParams{
		timeout: timeout,
	}
}

// NewValidateHostsOperationsParamsWithContext creates a new ValidateHostsOperationsParams object
// with the ability to set a context for a request.
func NewValidateHostsOperationsParamsWithContext(ctx context.Context) *ValidateHostsOperationsParams {
	return &ValidateHostsOperationsParams{
		Context: ctx,
	}
}

// NewValidateHostsOperationsParamsWithHTTPClient creates a new ValidateHostsOperationsParams object
// with the ability to set a custom HTTPClient for a request.
func NewValidateHostsOperationsParamsWithHTTPClient(client *http.Client) *ValidateHostsOperationsParams {
	return &ValidateHostsOperationsParams{
		HTTPClient: client,
	}
}

/*
ValidateHostsOperationsParams contains all the parameters to send to the API endpoint

	for the validate hosts operations operation.

	Typically these are written to a http.Request.
*/
type ValidateHostsOperationsParams struct {

	/* HostCommissionSpecs.

	   hostCommissionSpecs
	*/
	HostCommissionSpecs []*models.HostCommissionSpec

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the validate hosts operations params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ValidateHostsOperationsParams) WithDefaults() *ValidateHostsOperationsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the validate hosts operations params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ValidateHostsOperationsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the validate hosts operations params
func (o *ValidateHostsOperationsParams) WithTimeout(timeout time.Duration) *ValidateHostsOperationsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the validate hosts operations params
func (o *ValidateHostsOperationsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the validate hosts operations params
func (o *ValidateHostsOperationsParams) WithContext(ctx context.Context) *ValidateHostsOperationsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the validate hosts operations params
func (o *ValidateHostsOperationsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the validate hosts operations params
func (o *ValidateHostsOperationsParams) WithHTTPClient(client *http.Client) *ValidateHostsOperationsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the validate hosts operations params
func (o *ValidateHostsOperationsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithHostCommissionSpecs adds the hostCommissionSpecs to the validate hosts operations params
func (o *ValidateHostsOperationsParams) WithHostCommissionSpecs(hostCommissionSpecs []*models.HostCommissionSpec) *ValidateHostsOperationsParams {
	o.SetHostCommissionSpecs(hostCommissionSpecs)
	return o
}

// SetHostCommissionSpecs adds the hostCommissionSpecs to the validate hosts operations params
func (o *ValidateHostsOperationsParams) SetHostCommissionSpecs(hostCommissionSpecs []*models.HostCommissionSpec) {
	o.HostCommissionSpecs = hostCommissionSpecs
}

// WriteToRequest writes these params to a swagger request
func (o *ValidateHostsOperationsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.HostCommissionSpecs != nil {
		if err := r.SetBodyParam(o.HostCommissionSpecs); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
