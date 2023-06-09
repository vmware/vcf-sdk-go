// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2023 VMware, Inc.
// SPDX-License-Identifier: BSD-2-Clause

package vsan_health_check

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

// NewGetHealthCheckStatusTaskParams creates a new GetHealthCheckStatusTaskParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetHealthCheckStatusTaskParams() *GetHealthCheckStatusTaskParams {
	return &GetHealthCheckStatusTaskParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetHealthCheckStatusTaskParamsWithTimeout creates a new GetHealthCheckStatusTaskParams object
// with the ability to set a timeout on a request.
func NewGetHealthCheckStatusTaskParamsWithTimeout(timeout time.Duration) *GetHealthCheckStatusTaskParams {
	return &GetHealthCheckStatusTaskParams{
		timeout: timeout,
	}
}

// NewGetHealthCheckStatusTaskParamsWithContext creates a new GetHealthCheckStatusTaskParams object
// with the ability to set a context for a request.
func NewGetHealthCheckStatusTaskParamsWithContext(ctx context.Context) *GetHealthCheckStatusTaskParams {
	return &GetHealthCheckStatusTaskParams{
		Context: ctx,
	}
}

// NewGetHealthCheckStatusTaskParamsWithHTTPClient creates a new GetHealthCheckStatusTaskParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetHealthCheckStatusTaskParamsWithHTTPClient(client *http.Client) *GetHealthCheckStatusTaskParams {
	return &GetHealthCheckStatusTaskParams{
		HTTPClient: client,
	}
}

/*
GetHealthCheckStatusTaskParams contains all the parameters to send to the API endpoint

	for the get health check status task operation.

	Typically these are written to a http.Request.
*/
type GetHealthCheckStatusTaskParams struct {

	/* DomainID.

	   Domain ID
	*/
	DomainID string

	/* TaskID.

	   Health check task id
	*/
	TaskID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get health check status task params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetHealthCheckStatusTaskParams) WithDefaults() *GetHealthCheckStatusTaskParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get health check status task params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetHealthCheckStatusTaskParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get health check status task params
func (o *GetHealthCheckStatusTaskParams) WithTimeout(timeout time.Duration) *GetHealthCheckStatusTaskParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get health check status task params
func (o *GetHealthCheckStatusTaskParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get health check status task params
func (o *GetHealthCheckStatusTaskParams) WithContext(ctx context.Context) *GetHealthCheckStatusTaskParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get health check status task params
func (o *GetHealthCheckStatusTaskParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get health check status task params
func (o *GetHealthCheckStatusTaskParams) WithHTTPClient(client *http.Client) *GetHealthCheckStatusTaskParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get health check status task params
func (o *GetHealthCheckStatusTaskParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDomainID adds the domainID to the get health check status task params
func (o *GetHealthCheckStatusTaskParams) WithDomainID(domainID string) *GetHealthCheckStatusTaskParams {
	o.SetDomainID(domainID)
	return o
}

// SetDomainID adds the domainId to the get health check status task params
func (o *GetHealthCheckStatusTaskParams) SetDomainID(domainID string) {
	o.DomainID = domainID
}

// WithTaskID adds the taskID to the get health check status task params
func (o *GetHealthCheckStatusTaskParams) WithTaskID(taskID string) *GetHealthCheckStatusTaskParams {
	o.SetTaskID(taskID)
	return o
}

// SetTaskID adds the taskId to the get health check status task params
func (o *GetHealthCheckStatusTaskParams) SetTaskID(taskID string) {
	o.TaskID = taskID
}

// WriteToRequest writes these params to a swagger request
func (o *GetHealthCheckStatusTaskParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param domainId
	if err := r.SetPathParam("domainId", o.DomainID); err != nil {
		return err
	}

	// path param taskId
	if err := r.SetPathParam("taskId", o.TaskID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
