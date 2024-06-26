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

// NewAssignTagsToClusterParams creates a new AssignTagsToClusterParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAssignTagsToClusterParams() *AssignTagsToClusterParams {
	return &AssignTagsToClusterParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAssignTagsToClusterParamsWithTimeout creates a new AssignTagsToClusterParams object
// with the ability to set a timeout on a request.
func NewAssignTagsToClusterParamsWithTimeout(timeout time.Duration) *AssignTagsToClusterParams {
	return &AssignTagsToClusterParams{
		timeout: timeout,
	}
}

// NewAssignTagsToClusterParamsWithContext creates a new AssignTagsToClusterParams object
// with the ability to set a context for a request.
func NewAssignTagsToClusterParamsWithContext(ctx context.Context) *AssignTagsToClusterParams {
	return &AssignTagsToClusterParams{
		Context: ctx,
	}
}

// NewAssignTagsToClusterParamsWithHTTPClient creates a new AssignTagsToClusterParams object
// with the ability to set a custom HTTPClient for a request.
func NewAssignTagsToClusterParamsWithHTTPClient(client *http.Client) *AssignTagsToClusterParams {
	return &AssignTagsToClusterParams{
		HTTPClient: client,
	}
}

/*
AssignTagsToClusterParams contains all the parameters to send to the API endpoint

	for the assign tags to cluster operation.

	Typically these are written to a http.Request.
*/
type AssignTagsToClusterParams struct {

	/* ID.

	   Cluster ID
	*/
	ID string

	/* TagsSpec.

	   Tags Spec
	*/
	TagsSpec *models.TagsSpec

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the assign tags to cluster params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AssignTagsToClusterParams) WithDefaults() *AssignTagsToClusterParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the assign tags to cluster params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AssignTagsToClusterParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the assign tags to cluster params
func (o *AssignTagsToClusterParams) WithTimeout(timeout time.Duration) *AssignTagsToClusterParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the assign tags to cluster params
func (o *AssignTagsToClusterParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the assign tags to cluster params
func (o *AssignTagsToClusterParams) WithContext(ctx context.Context) *AssignTagsToClusterParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the assign tags to cluster params
func (o *AssignTagsToClusterParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the assign tags to cluster params
func (o *AssignTagsToClusterParams) WithHTTPClient(client *http.Client) *AssignTagsToClusterParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the assign tags to cluster params
func (o *AssignTagsToClusterParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the assign tags to cluster params
func (o *AssignTagsToClusterParams) WithID(id string) *AssignTagsToClusterParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the assign tags to cluster params
func (o *AssignTagsToClusterParams) SetID(id string) {
	o.ID = id
}

// WithTagsSpec adds the tagsSpec to the assign tags to cluster params
func (o *AssignTagsToClusterParams) WithTagsSpec(tagsSpec *models.TagsSpec) *AssignTagsToClusterParams {
	o.SetTagsSpec(tagsSpec)
	return o
}

// SetTagsSpec adds the tagsSpec to the assign tags to cluster params
func (o *AssignTagsToClusterParams) SetTagsSpec(tagsSpec *models.TagsSpec) {
	o.TagsSpec = tagsSpec
}

// WriteToRequest writes these params to a swagger request
func (o *AssignTagsToClusterParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}
	if o.TagsSpec != nil {
		if err := r.SetBodyParam(o.TagsSpec); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
