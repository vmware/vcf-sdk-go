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

// NewAddDatastoreToClusterParams creates a new AddDatastoreToClusterParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAddDatastoreToClusterParams() *AddDatastoreToClusterParams {
	return &AddDatastoreToClusterParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAddDatastoreToClusterParamsWithTimeout creates a new AddDatastoreToClusterParams object
// with the ability to set a timeout on a request.
func NewAddDatastoreToClusterParamsWithTimeout(timeout time.Duration) *AddDatastoreToClusterParams {
	return &AddDatastoreToClusterParams{
		timeout: timeout,
	}
}

// NewAddDatastoreToClusterParamsWithContext creates a new AddDatastoreToClusterParams object
// with the ability to set a context for a request.
func NewAddDatastoreToClusterParamsWithContext(ctx context.Context) *AddDatastoreToClusterParams {
	return &AddDatastoreToClusterParams{
		Context: ctx,
	}
}

// NewAddDatastoreToClusterParamsWithHTTPClient creates a new AddDatastoreToClusterParams object
// with the ability to set a custom HTTPClient for a request.
func NewAddDatastoreToClusterParamsWithHTTPClient(client *http.Client) *AddDatastoreToClusterParams {
	return &AddDatastoreToClusterParams{
		HTTPClient: client,
	}
}

/*
AddDatastoreToClusterParams contains all the parameters to send to the API endpoint

	for the add datastore to cluster operation.

	Typically these are written to a http.Request.
*/
type AddDatastoreToClusterParams struct {

	/* DatastoreMountSpec.

	   Datastore Mount Spec
	*/
	DatastoreMountSpec *models.DatastoreMountSpec

	/* ID.

	   Cluster ID
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the add datastore to cluster params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AddDatastoreToClusterParams) WithDefaults() *AddDatastoreToClusterParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the add datastore to cluster params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AddDatastoreToClusterParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the add datastore to cluster params
func (o *AddDatastoreToClusterParams) WithTimeout(timeout time.Duration) *AddDatastoreToClusterParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the add datastore to cluster params
func (o *AddDatastoreToClusterParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the add datastore to cluster params
func (o *AddDatastoreToClusterParams) WithContext(ctx context.Context) *AddDatastoreToClusterParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the add datastore to cluster params
func (o *AddDatastoreToClusterParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the add datastore to cluster params
func (o *AddDatastoreToClusterParams) WithHTTPClient(client *http.Client) *AddDatastoreToClusterParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the add datastore to cluster params
func (o *AddDatastoreToClusterParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDatastoreMountSpec adds the datastoreMountSpec to the add datastore to cluster params
func (o *AddDatastoreToClusterParams) WithDatastoreMountSpec(datastoreMountSpec *models.DatastoreMountSpec) *AddDatastoreToClusterParams {
	o.SetDatastoreMountSpec(datastoreMountSpec)
	return o
}

// SetDatastoreMountSpec adds the datastoreMountSpec to the add datastore to cluster params
func (o *AddDatastoreToClusterParams) SetDatastoreMountSpec(datastoreMountSpec *models.DatastoreMountSpec) {
	o.DatastoreMountSpec = datastoreMountSpec
}

// WithID adds the id to the add datastore to cluster params
func (o *AddDatastoreToClusterParams) WithID(id string) *AddDatastoreToClusterParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the add datastore to cluster params
func (o *AddDatastoreToClusterParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *AddDatastoreToClusterParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.DatastoreMountSpec != nil {
		if err := r.SetBodyParam(o.DatastoreMountSpec); err != nil {
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
