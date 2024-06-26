// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2023 VMware, Inc.
// SPDX-License-Identifier: BSD-2-Clause

package backup_restore

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

// NewStartRestoreParams creates a new StartRestoreParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewStartRestoreParams() *StartRestoreParams {
	return &StartRestoreParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewStartRestoreParamsWithTimeout creates a new StartRestoreParams object
// with the ability to set a timeout on a request.
func NewStartRestoreParamsWithTimeout(timeout time.Duration) *StartRestoreParams {
	return &StartRestoreParams{
		timeout: timeout,
	}
}

// NewStartRestoreParamsWithContext creates a new StartRestoreParams object
// with the ability to set a context for a request.
func NewStartRestoreParamsWithContext(ctx context.Context) *StartRestoreParams {
	return &StartRestoreParams{
		Context: ctx,
	}
}

// NewStartRestoreParamsWithHTTPClient creates a new StartRestoreParams object
// with the ability to set a custom HTTPClient for a request.
func NewStartRestoreParamsWithHTTPClient(client *http.Client) *StartRestoreParams {
	return &StartRestoreParams{
		HTTPClient: client,
	}
}

/*
StartRestoreParams contains all the parameters to send to the API endpoint

	for the start restore operation.

	Typically these are written to a http.Request.
*/
type StartRestoreParams struct {

	/* RestoreSpec.

	   restoreSpec
	*/
	RestoreSpec *models.RestoreSpec

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the start restore params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *StartRestoreParams) WithDefaults() *StartRestoreParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the start restore params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *StartRestoreParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the start restore params
func (o *StartRestoreParams) WithTimeout(timeout time.Duration) *StartRestoreParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the start restore params
func (o *StartRestoreParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the start restore params
func (o *StartRestoreParams) WithContext(ctx context.Context) *StartRestoreParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the start restore params
func (o *StartRestoreParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the start restore params
func (o *StartRestoreParams) WithHTTPClient(client *http.Client) *StartRestoreParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the start restore params
func (o *StartRestoreParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRestoreSpec adds the restoreSpec to the start restore params
func (o *StartRestoreParams) WithRestoreSpec(restoreSpec *models.RestoreSpec) *StartRestoreParams {
	o.SetRestoreSpec(restoreSpec)
	return o
}

// SetRestoreSpec adds the restoreSpec to the start restore params
func (o *StartRestoreParams) SetRestoreSpec(restoreSpec *models.RestoreSpec) {
	o.RestoreSpec = restoreSpec
}

// WriteToRequest writes these params to a swagger request
func (o *StartRestoreParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.RestoreSpec != nil {
		if err := r.SetBodyParam(o.RestoreSpec); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
