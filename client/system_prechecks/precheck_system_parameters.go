// Code generated by go-swagger; DO NOT EDIT.

// Copyright © 2023 VMware, Inc. All Rights Reserved.
// Code licenced under: BSD-2

package system_prechecks

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

	"vcf-sdk-go/models"
)

// NewPrecheckSystemParams creates a new PrecheckSystemParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPrecheckSystemParams() *PrecheckSystemParams {
	return &PrecheckSystemParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPrecheckSystemParamsWithTimeout creates a new PrecheckSystemParams object
// with the ability to set a timeout on a request.
func NewPrecheckSystemParamsWithTimeout(timeout time.Duration) *PrecheckSystemParams {
	return &PrecheckSystemParams{
		timeout: timeout,
	}
}

// NewPrecheckSystemParamsWithContext creates a new PrecheckSystemParams object
// with the ability to set a context for a request.
func NewPrecheckSystemParamsWithContext(ctx context.Context) *PrecheckSystemParams {
	return &PrecheckSystemParams{
		Context: ctx,
	}
}

// NewPrecheckSystemParamsWithHTTPClient creates a new PrecheckSystemParams object
// with the ability to set a custom HTTPClient for a request.
func NewPrecheckSystemParamsWithHTTPClient(client *http.Client) *PrecheckSystemParams {
	return &PrecheckSystemParams{
		HTTPClient: client,
	}
}

/*
PrecheckSystemParams contains all the parameters to send to the API endpoint

	for the precheck system operation.

	Typically these are written to a http.Request.
*/
type PrecheckSystemParams struct {

	/* PrecheckSpec.

	   Precheck System Spec
	*/
	PrecheckSpec *models.PrecheckSpec

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the precheck system params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PrecheckSystemParams) WithDefaults() *PrecheckSystemParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the precheck system params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PrecheckSystemParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the precheck system params
func (o *PrecheckSystemParams) WithTimeout(timeout time.Duration) *PrecheckSystemParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the precheck system params
func (o *PrecheckSystemParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the precheck system params
func (o *PrecheckSystemParams) WithContext(ctx context.Context) *PrecheckSystemParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the precheck system params
func (o *PrecheckSystemParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the precheck system params
func (o *PrecheckSystemParams) WithHTTPClient(client *http.Client) *PrecheckSystemParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the precheck system params
func (o *PrecheckSystemParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPrecheckSpec adds the precheckSpec to the precheck system params
func (o *PrecheckSystemParams) WithPrecheckSpec(precheckSpec *models.PrecheckSpec) *PrecheckSystemParams {
	o.SetPrecheckSpec(precheckSpec)
	return o
}

// SetPrecheckSpec adds the precheckSpec to the precheck system params
func (o *PrecheckSystemParams) SetPrecheckSpec(precheckSpec *models.PrecheckSpec) {
	o.PrecheckSpec = precheckSpec
}

// WriteToRequest writes these params to a swagger request
func (o *PrecheckSystemParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.PrecheckSpec != nil {
		if err := r.SetBodyParam(o.PrecheckSpec); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
