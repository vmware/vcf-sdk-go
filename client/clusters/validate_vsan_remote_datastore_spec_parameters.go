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

// NewValidateVSANRemoteDatastoreSpecParams creates a new ValidateVSANRemoteDatastoreSpecParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewValidateVSANRemoteDatastoreSpecParams() *ValidateVSANRemoteDatastoreSpecParams {
	return &ValidateVSANRemoteDatastoreSpecParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewValidateVSANRemoteDatastoreSpecParamsWithTimeout creates a new ValidateVSANRemoteDatastoreSpecParams object
// with the ability to set a timeout on a request.
func NewValidateVSANRemoteDatastoreSpecParamsWithTimeout(timeout time.Duration) *ValidateVSANRemoteDatastoreSpecParams {
	return &ValidateVSANRemoteDatastoreSpecParams{
		timeout: timeout,
	}
}

// NewValidateVSANRemoteDatastoreSpecParamsWithContext creates a new ValidateVSANRemoteDatastoreSpecParams object
// with the ability to set a context for a request.
func NewValidateVSANRemoteDatastoreSpecParamsWithContext(ctx context.Context) *ValidateVSANRemoteDatastoreSpecParams {
	return &ValidateVSANRemoteDatastoreSpecParams{
		Context: ctx,
	}
}

// NewValidateVSANRemoteDatastoreSpecParamsWithHTTPClient creates a new ValidateVSANRemoteDatastoreSpecParams object
// with the ability to set a custom HTTPClient for a request.
func NewValidateVSANRemoteDatastoreSpecParamsWithHTTPClient(client *http.Client) *ValidateVSANRemoteDatastoreSpecParams {
	return &ValidateVSANRemoteDatastoreSpecParams{
		HTTPClient: client,
	}
}

/*
ValidateVSANRemoteDatastoreSpecParams contains all the parameters to send to the API endpoint

	for the validate Vsan remote datastore spec operation.

	Typically these are written to a http.Request.
*/
type ValidateVSANRemoteDatastoreSpecParams struct {

	/* ClusterID.

	   Cluster ID
	*/
	ClusterID string

	/* DatastoreMountSpec.

	   Datastore Mount Spec
	*/
	DatastoreMountSpec *models.DatastoreMountSpec

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the validate Vsan remote datastore spec params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ValidateVSANRemoteDatastoreSpecParams) WithDefaults() *ValidateVSANRemoteDatastoreSpecParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the validate Vsan remote datastore spec params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ValidateVSANRemoteDatastoreSpecParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the validate Vsan remote datastore spec params
func (o *ValidateVSANRemoteDatastoreSpecParams) WithTimeout(timeout time.Duration) *ValidateVSANRemoteDatastoreSpecParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the validate Vsan remote datastore spec params
func (o *ValidateVSANRemoteDatastoreSpecParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the validate Vsan remote datastore spec params
func (o *ValidateVSANRemoteDatastoreSpecParams) WithContext(ctx context.Context) *ValidateVSANRemoteDatastoreSpecParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the validate Vsan remote datastore spec params
func (o *ValidateVSANRemoteDatastoreSpecParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the validate Vsan remote datastore spec params
func (o *ValidateVSANRemoteDatastoreSpecParams) WithHTTPClient(client *http.Client) *ValidateVSANRemoteDatastoreSpecParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the validate Vsan remote datastore spec params
func (o *ValidateVSANRemoteDatastoreSpecParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClusterID adds the clusterID to the validate Vsan remote datastore spec params
func (o *ValidateVSANRemoteDatastoreSpecParams) WithClusterID(clusterID string) *ValidateVSANRemoteDatastoreSpecParams {
	o.SetClusterID(clusterID)
	return o
}

// SetClusterID adds the clusterId to the validate Vsan remote datastore spec params
func (o *ValidateVSANRemoteDatastoreSpecParams) SetClusterID(clusterID string) {
	o.ClusterID = clusterID
}

// WithDatastoreMountSpec adds the datastoreMountSpec to the validate Vsan remote datastore spec params
func (o *ValidateVSANRemoteDatastoreSpecParams) WithDatastoreMountSpec(datastoreMountSpec *models.DatastoreMountSpec) *ValidateVSANRemoteDatastoreSpecParams {
	o.SetDatastoreMountSpec(datastoreMountSpec)
	return o
}

// SetDatastoreMountSpec adds the datastoreMountSpec to the validate Vsan remote datastore spec params
func (o *ValidateVSANRemoteDatastoreSpecParams) SetDatastoreMountSpec(datastoreMountSpec *models.DatastoreMountSpec) {
	o.DatastoreMountSpec = datastoreMountSpec
}

// WriteToRequest writes these params to a swagger request
func (o *ValidateVSANRemoteDatastoreSpecParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param clusterId
	if err := r.SetPathParam("clusterId", o.ClusterID); err != nil {
		return err
	}
	if o.DatastoreMountSpec != nil {
		if err := r.SetBodyParam(o.DatastoreMountSpec); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}