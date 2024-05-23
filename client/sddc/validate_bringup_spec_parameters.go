// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2023 VMware, Inc.
// SPDX-License-Identifier: BSD-2-Clause

package sddc

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

	"github.com/vmware/vcf-sdk-go/models"
)

// NewValidateBringupSpecParams creates a new ValidateBringupSpecParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewValidateBringupSpecParams() *ValidateBringupSpecParams {
	return &ValidateBringupSpecParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewValidateBringupSpecParamsWithTimeout creates a new ValidateBringupSpecParams object
// with the ability to set a timeout on a request.
func NewValidateBringupSpecParamsWithTimeout(timeout time.Duration) *ValidateBringupSpecParams {
	return &ValidateBringupSpecParams{
		timeout: timeout,
	}
}

// NewValidateBringupSpecParamsWithContext creates a new ValidateBringupSpecParams object
// with the ability to set a context for a request.
func NewValidateBringupSpecParamsWithContext(ctx context.Context) *ValidateBringupSpecParams {
	return &ValidateBringupSpecParams{
		Context: ctx,
	}
}

// NewValidateBringupSpecParamsWithHTTPClient creates a new ValidateBringupSpecParams object
// with the ability to set a custom HTTPClient for a request.
func NewValidateBringupSpecParamsWithHTTPClient(client *http.Client) *ValidateBringupSpecParams {
	return &ValidateBringupSpecParams{
		HTTPClient: client,
	}
}

/*
ValidateBringupSpecParams contains all the parameters to send to the API endpoint

	for the validate bringup spec operation.

	Typically these are written to a http.Request.
*/
type ValidateBringupSpecParams struct {

	/* Name.

	     Validation name
	Deprecated: ESXI_VERSION_VALIDATION
	*/
	Name *string

	/* Redo.

	   redo
	*/
	Redo *bool

	/* SDDCSpec.

	   SDDC specification
	*/
	SDDCSpec *models.SDDCSpec

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the validate bringup spec params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ValidateBringupSpecParams) WithDefaults() *ValidateBringupSpecParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the validate bringup spec params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ValidateBringupSpecParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the validate bringup spec params
func (o *ValidateBringupSpecParams) WithTimeout(timeout time.Duration) *ValidateBringupSpecParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the validate bringup spec params
func (o *ValidateBringupSpecParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the validate bringup spec params
func (o *ValidateBringupSpecParams) WithContext(ctx context.Context) *ValidateBringupSpecParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the validate bringup spec params
func (o *ValidateBringupSpecParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the validate bringup spec params
func (o *ValidateBringupSpecParams) WithHTTPClient(client *http.Client) *ValidateBringupSpecParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the validate bringup spec params
func (o *ValidateBringupSpecParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the validate bringup spec params
func (o *ValidateBringupSpecParams) WithName(name *string) *ValidateBringupSpecParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the validate bringup spec params
func (o *ValidateBringupSpecParams) SetName(name *string) {
	o.Name = name
}

// WithRedo adds the redo to the validate bringup spec params
func (o *ValidateBringupSpecParams) WithRedo(redo *bool) *ValidateBringupSpecParams {
	o.SetRedo(redo)
	return o
}

// SetRedo adds the redo to the validate bringup spec params
func (o *ValidateBringupSpecParams) SetRedo(redo *bool) {
	o.Redo = redo
}

// WithSDDCSpec adds the sDDCSpec to the validate bringup spec params
func (o *ValidateBringupSpecParams) WithSDDCSpec(sDDCSpec *models.SDDCSpec) *ValidateBringupSpecParams {
	o.SetSDDCSpec(sDDCSpec)
	return o
}

// SetSDDCSpec adds the sddcSpec to the validate bringup spec params
func (o *ValidateBringupSpecParams) SetSDDCSpec(sDDCSpec *models.SDDCSpec) {
	o.SDDCSpec = sDDCSpec
}

// WriteToRequest writes these params to a swagger request
func (o *ValidateBringupSpecParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Name != nil {

		// query param name
		var qrName string

		if o.Name != nil {
			qrName = *o.Name
		}
		qName := qrName
		if qName != "" {

			if err := r.SetQueryParam("name", qName); err != nil {
				return err
			}
		}
	}

	if o.Redo != nil {

		// query param redo
		var qrRedo bool

		if o.Redo != nil {
			qrRedo = *o.Redo
		}
		qRedo := swag.FormatBool(qrRedo)
		if qRedo != "" {

			if err := r.SetQueryParam("redo", qRedo); err != nil {
				return err
			}
		}
	}
	if o.SDDCSpec != nil {
		if err := r.SetBodyParam(o.SDDCSpec); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
