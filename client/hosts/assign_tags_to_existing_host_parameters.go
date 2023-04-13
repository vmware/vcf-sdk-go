// Code generated by go-swagger; DO NOT EDIT.

// Copyright © 2023 VMware, Inc. All Rights Reserved.
// Code licenced under: BSD-2

package hosts

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

// NewAssignTagsToExistingHostParams creates a new AssignTagsToExistingHostParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAssignTagsToExistingHostParams() *AssignTagsToExistingHostParams {
	return &AssignTagsToExistingHostParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAssignTagsToExistingHostParamsWithTimeout creates a new AssignTagsToExistingHostParams object
// with the ability to set a timeout on a request.
func NewAssignTagsToExistingHostParamsWithTimeout(timeout time.Duration) *AssignTagsToExistingHostParams {
	return &AssignTagsToExistingHostParams{
		timeout: timeout,
	}
}

// NewAssignTagsToExistingHostParamsWithContext creates a new AssignTagsToExistingHostParams object
// with the ability to set a context for a request.
func NewAssignTagsToExistingHostParamsWithContext(ctx context.Context) *AssignTagsToExistingHostParams {
	return &AssignTagsToExistingHostParams{
		Context: ctx,
	}
}

// NewAssignTagsToExistingHostParamsWithHTTPClient creates a new AssignTagsToExistingHostParams object
// with the ability to set a custom HTTPClient for a request.
func NewAssignTagsToExistingHostParamsWithHTTPClient(client *http.Client) *AssignTagsToExistingHostParams {
	return &AssignTagsToExistingHostParams{
		HTTPClient: client,
	}
}

/*
AssignTagsToExistingHostParams contains all the parameters to send to the API endpoint

	for the assign tags to existing host operation.

	Typically these are written to a http.Request.
*/
type AssignTagsToExistingHostParams struct {

	/* ID.

	   Host ID
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

// WithDefaults hydrates default values in the assign tags to existing host params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AssignTagsToExistingHostParams) WithDefaults() *AssignTagsToExistingHostParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the assign tags to existing host params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AssignTagsToExistingHostParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the assign tags to existing host params
func (o *AssignTagsToExistingHostParams) WithTimeout(timeout time.Duration) *AssignTagsToExistingHostParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the assign tags to existing host params
func (o *AssignTagsToExistingHostParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the assign tags to existing host params
func (o *AssignTagsToExistingHostParams) WithContext(ctx context.Context) *AssignTagsToExistingHostParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the assign tags to existing host params
func (o *AssignTagsToExistingHostParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the assign tags to existing host params
func (o *AssignTagsToExistingHostParams) WithHTTPClient(client *http.Client) *AssignTagsToExistingHostParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the assign tags to existing host params
func (o *AssignTagsToExistingHostParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the assign tags to existing host params
func (o *AssignTagsToExistingHostParams) WithID(id string) *AssignTagsToExistingHostParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the assign tags to existing host params
func (o *AssignTagsToExistingHostParams) SetID(id string) {
	o.ID = id
}

// WithTagsSpec adds the tagsSpec to the assign tags to existing host params
func (o *AssignTagsToExistingHostParams) WithTagsSpec(tagsSpec *models.TagsSpec) *AssignTagsToExistingHostParams {
	o.SetTagsSpec(tagsSpec)
	return o
}

// SetTagsSpec adds the tagsSpec to the assign tags to existing host params
func (o *AssignTagsToExistingHostParams) SetTagsSpec(tagsSpec *models.TagsSpec) {
	o.TagsSpec = tagsSpec
}

// WriteToRequest writes these params to a swagger request
func (o *AssignTagsToExistingHostParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
