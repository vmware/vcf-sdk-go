// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2023 VMware, Inc.
// SPDX-License-Identifier: BSD-2-Clause

package certificates

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
)

// NewGetResourceCertificatesValidationResultParams creates a new GetResourceCertificatesValidationResultParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetResourceCertificatesValidationResultParams() *GetResourceCertificatesValidationResultParams {
	return &GetResourceCertificatesValidationResultParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetResourceCertificatesValidationResultParamsWithTimeout creates a new GetResourceCertificatesValidationResultParams object
// with the ability to set a timeout on a request.
func NewGetResourceCertificatesValidationResultParamsWithTimeout(timeout time.Duration) *GetResourceCertificatesValidationResultParams {
	return &GetResourceCertificatesValidationResultParams{
		timeout: timeout,
	}
}

// NewGetResourceCertificatesValidationResultParamsWithContext creates a new GetResourceCertificatesValidationResultParams object
// with the ability to set a context for a request.
func NewGetResourceCertificatesValidationResultParamsWithContext(ctx context.Context) *GetResourceCertificatesValidationResultParams {
	return &GetResourceCertificatesValidationResultParams{
		Context: ctx,
	}
}

// NewGetResourceCertificatesValidationResultParamsWithHTTPClient creates a new GetResourceCertificatesValidationResultParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetResourceCertificatesValidationResultParamsWithHTTPClient(client *http.Client) *GetResourceCertificatesValidationResultParams {
	return &GetResourceCertificatesValidationResultParams{
		HTTPClient: client,
	}
}

/*
GetResourceCertificatesValidationResultParams contains all the parameters to send to the API endpoint

	for the get resource certificates validation result operation.

	Typically these are written to a http.Request.
*/
type GetResourceCertificatesValidationResultParams struct {

	/* ID.

	   Domain ID
	*/
	ID string

	/* ValidationID.

	   Validation ID
	*/
	ValidationID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get resource certificates validation result params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetResourceCertificatesValidationResultParams) WithDefaults() *GetResourceCertificatesValidationResultParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get resource certificates validation result params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetResourceCertificatesValidationResultParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get resource certificates validation result params
func (o *GetResourceCertificatesValidationResultParams) WithTimeout(timeout time.Duration) *GetResourceCertificatesValidationResultParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get resource certificates validation result params
func (o *GetResourceCertificatesValidationResultParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get resource certificates validation result params
func (o *GetResourceCertificatesValidationResultParams) WithContext(ctx context.Context) *GetResourceCertificatesValidationResultParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get resource certificates validation result params
func (o *GetResourceCertificatesValidationResultParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get resource certificates validation result params
func (o *GetResourceCertificatesValidationResultParams) WithHTTPClient(client *http.Client) *GetResourceCertificatesValidationResultParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get resource certificates validation result params
func (o *GetResourceCertificatesValidationResultParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get resource certificates validation result params
func (o *GetResourceCertificatesValidationResultParams) WithID(id string) *GetResourceCertificatesValidationResultParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get resource certificates validation result params
func (o *GetResourceCertificatesValidationResultParams) SetID(id string) {
	o.ID = id
}

// WithValidationID adds the validationID to the get resource certificates validation result params
func (o *GetResourceCertificatesValidationResultParams) WithValidationID(validationID string) *GetResourceCertificatesValidationResultParams {
	o.SetValidationID(validationID)
	return o
}

// SetValidationID adds the validationId to the get resource certificates validation result params
func (o *GetResourceCertificatesValidationResultParams) SetValidationID(validationID string) {
	o.ValidationID = validationID
}

// WriteToRequest writes these params to a swagger request
func (o *GetResourceCertificatesValidationResultParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	// path param validationId
	if err := r.SetPathParam("validationId", o.ValidationID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
