// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2023 VMware, Inc.
// SPDX-License-Identifier: BSD-2-Clause

package clusters

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

// NewValidateClusterUpdateSpecParams creates a new ValidateClusterUpdateSpecParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewValidateClusterUpdateSpecParams() *ValidateClusterUpdateSpecParams {
	return &ValidateClusterUpdateSpecParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewValidateClusterUpdateSpecParamsWithTimeout creates a new ValidateClusterUpdateSpecParams object
// with the ability to set a timeout on a request.
func NewValidateClusterUpdateSpecParamsWithTimeout(timeout time.Duration) *ValidateClusterUpdateSpecParams {
	return &ValidateClusterUpdateSpecParams{
		timeout: timeout,
	}
}

// NewValidateClusterUpdateSpecParamsWithContext creates a new ValidateClusterUpdateSpecParams object
// with the ability to set a context for a request.
func NewValidateClusterUpdateSpecParamsWithContext(ctx context.Context) *ValidateClusterUpdateSpecParams {
	return &ValidateClusterUpdateSpecParams{
		Context: ctx,
	}
}

// NewValidateClusterUpdateSpecParamsWithHTTPClient creates a new ValidateClusterUpdateSpecParams object
// with the ability to set a custom HTTPClient for a request.
func NewValidateClusterUpdateSpecParamsWithHTTPClient(client *http.Client) *ValidateClusterUpdateSpecParams {
	return &ValidateClusterUpdateSpecParams{
		HTTPClient: client,
	}
}

/*
ValidateClusterUpdateSpecParams contains all the parameters to send to the API endpoint

	for the validate cluster update spec operation.

	Typically these are written to a http.Request.
*/
type ValidateClusterUpdateSpecParams struct {

	/* ClusterUpdateSpec.

	   clusterUpdateSpec
	*/
	ClusterUpdateSpec *models.ClusterUpdateSpec

	/* ID.

	   Cluster ID
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the validate cluster update spec params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ValidateClusterUpdateSpecParams) WithDefaults() *ValidateClusterUpdateSpecParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the validate cluster update spec params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ValidateClusterUpdateSpecParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the validate cluster update spec params
func (o *ValidateClusterUpdateSpecParams) WithTimeout(timeout time.Duration) *ValidateClusterUpdateSpecParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the validate cluster update spec params
func (o *ValidateClusterUpdateSpecParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the validate cluster update spec params
func (o *ValidateClusterUpdateSpecParams) WithContext(ctx context.Context) *ValidateClusterUpdateSpecParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the validate cluster update spec params
func (o *ValidateClusterUpdateSpecParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the validate cluster update spec params
func (o *ValidateClusterUpdateSpecParams) WithHTTPClient(client *http.Client) *ValidateClusterUpdateSpecParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the validate cluster update spec params
func (o *ValidateClusterUpdateSpecParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClusterUpdateSpec adds the clusterUpdateSpec to the validate cluster update spec params
func (o *ValidateClusterUpdateSpecParams) WithClusterUpdateSpec(clusterUpdateSpec *models.ClusterUpdateSpec) *ValidateClusterUpdateSpecParams {
	o.SetClusterUpdateSpec(clusterUpdateSpec)
	return o
}

// SetClusterUpdateSpec adds the clusterUpdateSpec to the validate cluster update spec params
func (o *ValidateClusterUpdateSpecParams) SetClusterUpdateSpec(clusterUpdateSpec *models.ClusterUpdateSpec) {
	o.ClusterUpdateSpec = clusterUpdateSpec
}

// WithID adds the id to the validate cluster update spec params
func (o *ValidateClusterUpdateSpecParams) WithID(id string) *ValidateClusterUpdateSpecParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the validate cluster update spec params
func (o *ValidateClusterUpdateSpecParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *ValidateClusterUpdateSpecParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.ClusterUpdateSpec != nil {
		if err := r.SetBodyParam(o.ClusterUpdateSpec); err != nil {
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
