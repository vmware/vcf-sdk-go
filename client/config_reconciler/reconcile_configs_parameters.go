// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2023 VMware, Inc.
// SPDX-License-Identifier: BSD-2-Clause

package config_reconciler

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

// NewReconcileConfigsParams creates a new ReconcileConfigsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewReconcileConfigsParams() *ReconcileConfigsParams {
	return &ReconcileConfigsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewReconcileConfigsParamsWithTimeout creates a new ReconcileConfigsParams object
// with the ability to set a timeout on a request.
func NewReconcileConfigsParamsWithTimeout(timeout time.Duration) *ReconcileConfigsParams {
	return &ReconcileConfigsParams{
		timeout: timeout,
	}
}

// NewReconcileConfigsParamsWithContext creates a new ReconcileConfigsParams object
// with the ability to set a context for a request.
func NewReconcileConfigsParamsWithContext(ctx context.Context) *ReconcileConfigsParams {
	return &ReconcileConfigsParams{
		Context: ctx,
	}
}

// NewReconcileConfigsParamsWithHTTPClient creates a new ReconcileConfigsParams object
// with the ability to set a custom HTTPClient for a request.
func NewReconcileConfigsParamsWithHTTPClient(client *http.Client) *ReconcileConfigsParams {
	return &ReconcileConfigsParams{
		HTTPClient: client,
	}
}

/*
ReconcileConfigsParams contains all the parameters to send to the API endpoint

	for the reconcile configs operation.

	Typically these are written to a http.Request.
*/
type ReconcileConfigsParams struct {

	/* ConfigDriftApplySpec.

	   configDriftApplySpec
	*/
	ConfigDriftApplySpec *models.ConfigDriftApplySpec

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the reconcile configs params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ReconcileConfigsParams) WithDefaults() *ReconcileConfigsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the reconcile configs params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ReconcileConfigsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the reconcile configs params
func (o *ReconcileConfigsParams) WithTimeout(timeout time.Duration) *ReconcileConfigsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the reconcile configs params
func (o *ReconcileConfigsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the reconcile configs params
func (o *ReconcileConfigsParams) WithContext(ctx context.Context) *ReconcileConfigsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the reconcile configs params
func (o *ReconcileConfigsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the reconcile configs params
func (o *ReconcileConfigsParams) WithHTTPClient(client *http.Client) *ReconcileConfigsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the reconcile configs params
func (o *ReconcileConfigsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithConfigDriftApplySpec adds the configDriftApplySpec to the reconcile configs params
func (o *ReconcileConfigsParams) WithConfigDriftApplySpec(configDriftApplySpec *models.ConfigDriftApplySpec) *ReconcileConfigsParams {
	o.SetConfigDriftApplySpec(configDriftApplySpec)
	return o
}

// SetConfigDriftApplySpec adds the configDriftApplySpec to the reconcile configs params
func (o *ReconcileConfigsParams) SetConfigDriftApplySpec(configDriftApplySpec *models.ConfigDriftApplySpec) {
	o.ConfigDriftApplySpec = configDriftApplySpec
}

// WriteToRequest writes these params to a swagger request
func (o *ReconcileConfigsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.ConfigDriftApplySpec != nil {
		if err := r.SetBodyParam(o.ConfigDriftApplySpec); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
