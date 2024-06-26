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

	"github.com/vmware/vcf-sdk-go/models"
)

// NewUpdateVSANHealthCheckByDomainParams creates a new UpdateVSANHealthCheckByDomainParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateVSANHealthCheckByDomainParams() *UpdateVSANHealthCheckByDomainParams {
	return &UpdateVSANHealthCheckByDomainParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateVSANHealthCheckByDomainParamsWithTimeout creates a new UpdateVSANHealthCheckByDomainParams object
// with the ability to set a timeout on a request.
func NewUpdateVSANHealthCheckByDomainParamsWithTimeout(timeout time.Duration) *UpdateVSANHealthCheckByDomainParams {
	return &UpdateVSANHealthCheckByDomainParams{
		timeout: timeout,
	}
}

// NewUpdateVSANHealthCheckByDomainParamsWithContext creates a new UpdateVSANHealthCheckByDomainParams object
// with the ability to set a context for a request.
func NewUpdateVSANHealthCheckByDomainParamsWithContext(ctx context.Context) *UpdateVSANHealthCheckByDomainParams {
	return &UpdateVSANHealthCheckByDomainParams{
		Context: ctx,
	}
}

// NewUpdateVSANHealthCheckByDomainParamsWithHTTPClient creates a new UpdateVSANHealthCheckByDomainParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateVSANHealthCheckByDomainParamsWithHTTPClient(client *http.Client) *UpdateVSANHealthCheckByDomainParams {
	return &UpdateVSANHealthCheckByDomainParams{
		HTTPClient: client,
	}
}

/*
UpdateVSANHealthCheckByDomainParams contains all the parameters to send to the API endpoint

	for the update Vsan health check by domain operation.

	Typically these are written to a http.Request.
*/
type UpdateVSANHealthCheckByDomainParams struct {

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

// WithDefaults hydrates default values in the update Vsan health check by domain params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateVSANHealthCheckByDomainParams) WithDefaults() *UpdateVSANHealthCheckByDomainParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update Vsan health check by domain params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateVSANHealthCheckByDomainParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update Vsan health check by domain params
func (o *UpdateVSANHealthCheckByDomainParams) WithTimeout(timeout time.Duration) *UpdateVSANHealthCheckByDomainParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update Vsan health check by domain params
func (o *UpdateVSANHealthCheckByDomainParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update Vsan health check by domain params
func (o *UpdateVSANHealthCheckByDomainParams) WithContext(ctx context.Context) *UpdateVSANHealthCheckByDomainParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update Vsan health check by domain params
func (o *UpdateVSANHealthCheckByDomainParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update Vsan health check by domain params
func (o *UpdateVSANHealthCheckByDomainParams) WithHTTPClient(client *http.Client) *UpdateVSANHealthCheckByDomainParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update Vsan health check by domain params
func (o *UpdateVSANHealthCheckByDomainParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDomainID adds the domainID to the update Vsan health check by domain params
func (o *UpdateVSANHealthCheckByDomainParams) WithDomainID(domainID string) *UpdateVSANHealthCheckByDomainParams {
	o.SetDomainID(domainID)
	return o
}

// SetDomainID adds the domainId to the update Vsan health check by domain params
func (o *UpdateVSANHealthCheckByDomainParams) SetDomainID(domainID string) {
	o.DomainID = domainID
}

// WithHealthCheckUpdateSpec adds the healthCheckUpdateSpec to the update Vsan health check by domain params
func (o *UpdateVSANHealthCheckByDomainParams) WithHealthCheckUpdateSpec(healthCheckUpdateSpec []*models.HealthCheckSpec) *UpdateVSANHealthCheckByDomainParams {
	o.SetHealthCheckUpdateSpec(healthCheckUpdateSpec)
	return o
}

// SetHealthCheckUpdateSpec adds the healthCheckUpdateSpec to the update Vsan health check by domain params
func (o *UpdateVSANHealthCheckByDomainParams) SetHealthCheckUpdateSpec(healthCheckUpdateSpec []*models.HealthCheckSpec) {
	o.HealthCheckUpdateSpec = healthCheckUpdateSpec
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateVSANHealthCheckByDomainParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
