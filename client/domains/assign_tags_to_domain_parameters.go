// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2023 VMware, Inc.
// SPDX-License-Identifier: BSD-2-Clause

package domains

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

// NewAssignTagsToDomainParams creates a new AssignTagsToDomainParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAssignTagsToDomainParams() *AssignTagsToDomainParams {
	return &AssignTagsToDomainParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAssignTagsToDomainParamsWithTimeout creates a new AssignTagsToDomainParams object
// with the ability to set a timeout on a request.
func NewAssignTagsToDomainParamsWithTimeout(timeout time.Duration) *AssignTagsToDomainParams {
	return &AssignTagsToDomainParams{
		timeout: timeout,
	}
}

// NewAssignTagsToDomainParamsWithContext creates a new AssignTagsToDomainParams object
// with the ability to set a context for a request.
func NewAssignTagsToDomainParamsWithContext(ctx context.Context) *AssignTagsToDomainParams {
	return &AssignTagsToDomainParams{
		Context: ctx,
	}
}

// NewAssignTagsToDomainParamsWithHTTPClient creates a new AssignTagsToDomainParams object
// with the ability to set a custom HTTPClient for a request.
func NewAssignTagsToDomainParamsWithHTTPClient(client *http.Client) *AssignTagsToDomainParams {
	return &AssignTagsToDomainParams{
		HTTPClient: client,
	}
}

/*
AssignTagsToDomainParams contains all the parameters to send to the API endpoint

	for the assign tags to domain operation.

	Typically these are written to a http.Request.
*/
type AssignTagsToDomainParams struct {

	/* ID.

	   Domain ID
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

// WithDefaults hydrates default values in the assign tags to domain params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AssignTagsToDomainParams) WithDefaults() *AssignTagsToDomainParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the assign tags to domain params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AssignTagsToDomainParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the assign tags to domain params
func (o *AssignTagsToDomainParams) WithTimeout(timeout time.Duration) *AssignTagsToDomainParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the assign tags to domain params
func (o *AssignTagsToDomainParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the assign tags to domain params
func (o *AssignTagsToDomainParams) WithContext(ctx context.Context) *AssignTagsToDomainParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the assign tags to domain params
func (o *AssignTagsToDomainParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the assign tags to domain params
func (o *AssignTagsToDomainParams) WithHTTPClient(client *http.Client) *AssignTagsToDomainParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the assign tags to domain params
func (o *AssignTagsToDomainParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the assign tags to domain params
func (o *AssignTagsToDomainParams) WithID(id string) *AssignTagsToDomainParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the assign tags to domain params
func (o *AssignTagsToDomainParams) SetID(id string) {
	o.ID = id
}

// WithTagsSpec adds the tagsSpec to the assign tags to domain params
func (o *AssignTagsToDomainParams) WithTagsSpec(tagsSpec *models.TagsSpec) *AssignTagsToDomainParams {
	o.SetTagsSpec(tagsSpec)
	return o
}

// SetTagsSpec adds the tagsSpec to the assign tags to domain params
func (o *AssignTagsToDomainParams) SetTagsSpec(tagsSpec *models.TagsSpec) {
	o.TagsSpec = tagsSpec
}

// WriteToRequest writes these params to a swagger request
func (o *AssignTagsToDomainParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
