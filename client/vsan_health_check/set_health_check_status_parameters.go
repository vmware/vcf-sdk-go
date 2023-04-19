// Code generated by go-swagger; DO NOT EDIT.

// Copyright © 2023 VMware, Inc. All Rights Reserved.
// Code licenced under: BSD-2

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

	"github.com/vmware/vcf-sdk-go/models"
)

// NewSetHealthCheckStatusParams creates a new SetHealthCheckStatusParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSetHealthCheckStatusParams() *SetHealthCheckStatusParams {
	return &SetHealthCheckStatusParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSetHealthCheckStatusParamsWithTimeout creates a new SetHealthCheckStatusParams object
// with the ability to set a timeout on a request.
func NewSetHealthCheckStatusParamsWithTimeout(timeout time.Duration) *SetHealthCheckStatusParams {
	return &SetHealthCheckStatusParams{
		timeout: timeout,
	}
}

// NewSetHealthCheckStatusParamsWithContext creates a new SetHealthCheckStatusParams object
// with the ability to set a context for a request.
func NewSetHealthCheckStatusParamsWithContext(ctx context.Context) *SetHealthCheckStatusParams {
	return &SetHealthCheckStatusParams{
		Context: ctx,
	}
}

// NewSetHealthCheckStatusParamsWithHTTPClient creates a new SetHealthCheckStatusParams object
// with the ability to set a custom HTTPClient for a request.
func NewSetHealthCheckStatusParamsWithHTTPClient(client *http.Client) *SetHealthCheckStatusParams {
	return &SetHealthCheckStatusParams{
		HTTPClient: client,
	}
}

/*
SetHealthCheckStatusParams contains all the parameters to send to the API endpoint

	for the set health check status operation.

	Typically these are written to a http.Request.
*/
type SetHealthCheckStatusParams struct {

	/* DomainID.

	   Domain ID
	*/
	DomainID string

	/* HealthCheckUpdateSpec.

	   Health check update spec
	*/
	HealthCheckUpdateSpec []*models.HealthCheckSpec

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the set health check status params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SetHealthCheckStatusParams) WithDefaults() *SetHealthCheckStatusParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the set health check status params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SetHealthCheckStatusParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the set health check status params
func (o *SetHealthCheckStatusParams) WithTimeout(timeout time.Duration) *SetHealthCheckStatusParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the set health check status params
func (o *SetHealthCheckStatusParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the set health check status params
func (o *SetHealthCheckStatusParams) WithContext(ctx context.Context) *SetHealthCheckStatusParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the set health check status params
func (o *SetHealthCheckStatusParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the set health check status params
func (o *SetHealthCheckStatusParams) WithHTTPClient(client *http.Client) *SetHealthCheckStatusParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the set health check status params
func (o *SetHealthCheckStatusParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDomainID adds the domainID to the set health check status params
func (o *SetHealthCheckStatusParams) WithDomainID(domainID string) *SetHealthCheckStatusParams {
	o.SetDomainID(domainID)
	return o
}

// SetDomainID adds the domainId to the set health check status params
func (o *SetHealthCheckStatusParams) SetDomainID(domainID string) {
	o.DomainID = domainID
}

// WithHealthCheckUpdateSpec adds the healthCheckUpdateSpec to the set health check status params
func (o *SetHealthCheckStatusParams) WithHealthCheckUpdateSpec(healthCheckUpdateSpec []*models.HealthCheckSpec) *SetHealthCheckStatusParams {
	o.SetHealthCheckUpdateSpec(healthCheckUpdateSpec)
	return o
}

// SetHealthCheckUpdateSpec adds the healthCheckUpdateSpec to the set health check status params
func (o *SetHealthCheckStatusParams) SetHealthCheckUpdateSpec(healthCheckUpdateSpec []*models.HealthCheckSpec) {
	o.HealthCheckUpdateSpec = healthCheckUpdateSpec
}

// WriteToRequest writes these params to a swagger request
func (o *SetHealthCheckStatusParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param domainId
	if err := r.SetPathParam("domainId", o.DomainID); err != nil {
		return err
	}
	if o.HealthCheckUpdateSpec != nil {
		if err := r.SetBodyParam(o.HealthCheckUpdateSpec); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}