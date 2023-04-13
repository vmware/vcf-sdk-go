// Code generated by go-swagger; DO NOT EDIT.

// Copyright © 2023 VMware, Inc. All Rights Reserved.
// Code licenced under: BSD-2

package credentials

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

// NewRetryCredentialsTaskParams creates a new RetryCredentialsTaskParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewRetryCredentialsTaskParams() *RetryCredentialsTaskParams {
	return &RetryCredentialsTaskParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewRetryCredentialsTaskParamsWithTimeout creates a new RetryCredentialsTaskParams object
// with the ability to set a timeout on a request.
func NewRetryCredentialsTaskParamsWithTimeout(timeout time.Duration) *RetryCredentialsTaskParams {
	return &RetryCredentialsTaskParams{
		timeout: timeout,
	}
}

// NewRetryCredentialsTaskParamsWithContext creates a new RetryCredentialsTaskParams object
// with the ability to set a context for a request.
func NewRetryCredentialsTaskParamsWithContext(ctx context.Context) *RetryCredentialsTaskParams {
	return &RetryCredentialsTaskParams{
		Context: ctx,
	}
}

// NewRetryCredentialsTaskParamsWithHTTPClient creates a new RetryCredentialsTaskParams object
// with the ability to set a custom HTTPClient for a request.
func NewRetryCredentialsTaskParamsWithHTTPClient(client *http.Client) *RetryCredentialsTaskParams {
	return &RetryCredentialsTaskParams{
		HTTPClient: client,
	}
}

/*
RetryCredentialsTaskParams contains all the parameters to send to the API endpoint

	for the retry credentials task operation.

	Typically these are written to a http.Request.
*/
type RetryCredentialsTaskParams struct {

	/* CredentialsUpdateSpec.

	   credentialsUpdateSpec
	*/
	CredentialsUpdateSpec *models.CredentialsUpdateSpec

	/* ID.

	   Task ID of the failed operation that is to be retried
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the retry credentials task params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RetryCredentialsTaskParams) WithDefaults() *RetryCredentialsTaskParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the retry credentials task params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RetryCredentialsTaskParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the retry credentials task params
func (o *RetryCredentialsTaskParams) WithTimeout(timeout time.Duration) *RetryCredentialsTaskParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the retry credentials task params
func (o *RetryCredentialsTaskParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the retry credentials task params
func (o *RetryCredentialsTaskParams) WithContext(ctx context.Context) *RetryCredentialsTaskParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the retry credentials task params
func (o *RetryCredentialsTaskParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the retry credentials task params
func (o *RetryCredentialsTaskParams) WithHTTPClient(client *http.Client) *RetryCredentialsTaskParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the retry credentials task params
func (o *RetryCredentialsTaskParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCredentialsUpdateSpec adds the credentialsUpdateSpec to the retry credentials task params
func (o *RetryCredentialsTaskParams) WithCredentialsUpdateSpec(credentialsUpdateSpec *models.CredentialsUpdateSpec) *RetryCredentialsTaskParams {
	o.SetCredentialsUpdateSpec(credentialsUpdateSpec)
	return o
}

// SetCredentialsUpdateSpec adds the credentialsUpdateSpec to the retry credentials task params
func (o *RetryCredentialsTaskParams) SetCredentialsUpdateSpec(credentialsUpdateSpec *models.CredentialsUpdateSpec) {
	o.CredentialsUpdateSpec = credentialsUpdateSpec
}

// WithID adds the id to the retry credentials task params
func (o *RetryCredentialsTaskParams) WithID(id string) *RetryCredentialsTaskParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the retry credentials task params
func (o *RetryCredentialsTaskParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *RetryCredentialsTaskParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.CredentialsUpdateSpec != nil {
		if err := r.SetBodyParam(o.CredentialsUpdateSpec); err != nil {
			return err
		}
	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
