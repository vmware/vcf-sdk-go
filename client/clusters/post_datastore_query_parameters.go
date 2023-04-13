// Code generated by go-swagger; DO NOT EDIT.

// Copyright © 2023 VMware, Inc. All Rights Reserved.
// Code licenced under: BSD-2

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

	"vcf-sdk-go/models"
)

// NewPOSTDatastoreQueryParams creates a new POSTDatastoreQueryParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPOSTDatastoreQueryParams() *POSTDatastoreQueryParams {
	return &POSTDatastoreQueryParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPOSTDatastoreQueryParamsWithTimeout creates a new POSTDatastoreQueryParams object
// with the ability to set a timeout on a request.
func NewPOSTDatastoreQueryParamsWithTimeout(timeout time.Duration) *POSTDatastoreQueryParams {
	return &POSTDatastoreQueryParams{
		timeout: timeout,
	}
}

// NewPOSTDatastoreQueryParamsWithContext creates a new POSTDatastoreQueryParams object
// with the ability to set a context for a request.
func NewPOSTDatastoreQueryParamsWithContext(ctx context.Context) *POSTDatastoreQueryParams {
	return &POSTDatastoreQueryParams{
		Context: ctx,
	}
}

// NewPOSTDatastoreQueryParamsWithHTTPClient creates a new POSTDatastoreQueryParams object
// with the ability to set a custom HTTPClient for a request.
func NewPOSTDatastoreQueryParamsWithHTTPClient(client *http.Client) *POSTDatastoreQueryParams {
	return &POSTDatastoreQueryParams{
		HTTPClient: client,
	}
}

/*
POSTDatastoreQueryParams contains all the parameters to send to the API endpoint

	for the post datastore query operation.

	Typically these are written to a http.Request.
*/
type POSTDatastoreQueryParams struct {

	/* DsCriterion.

	   dsCriterion
	*/
	DsCriterion *models.DatastoreCriterion

	/* ID.

	   Cluster ID
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post datastore query params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *POSTDatastoreQueryParams) WithDefaults() *POSTDatastoreQueryParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post datastore query params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *POSTDatastoreQueryParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post datastore query params
func (o *POSTDatastoreQueryParams) WithTimeout(timeout time.Duration) *POSTDatastoreQueryParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post datastore query params
func (o *POSTDatastoreQueryParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post datastore query params
func (o *POSTDatastoreQueryParams) WithContext(ctx context.Context) *POSTDatastoreQueryParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post datastore query params
func (o *POSTDatastoreQueryParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post datastore query params
func (o *POSTDatastoreQueryParams) WithHTTPClient(client *http.Client) *POSTDatastoreQueryParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post datastore query params
func (o *POSTDatastoreQueryParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDsCriterion adds the dsCriterion to the post datastore query params
func (o *POSTDatastoreQueryParams) WithDsCriterion(dsCriterion *models.DatastoreCriterion) *POSTDatastoreQueryParams {
	o.SetDsCriterion(dsCriterion)
	return o
}

// SetDsCriterion adds the dsCriterion to the post datastore query params
func (o *POSTDatastoreQueryParams) SetDsCriterion(dsCriterion *models.DatastoreCriterion) {
	o.DsCriterion = dsCriterion
}

// WithID adds the id to the post datastore query params
func (o *POSTDatastoreQueryParams) WithID(id string) *POSTDatastoreQueryParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the post datastore query params
func (o *POSTDatastoreQueryParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *POSTDatastoreQueryParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.DsCriterion != nil {
		if err := r.SetBodyParam(o.DsCriterion); err != nil {
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
