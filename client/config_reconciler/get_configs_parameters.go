// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2023 VMware, Inc.
// SPDX-License-Identifier: BSD-2-Clause

package config_reconciler

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
	"github.com/go-openapi/swag"
)

// NewGetConfigsParams creates a new GetConfigsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetConfigsParams() *GetConfigsParams {
	return &GetConfigsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetConfigsParamsWithTimeout creates a new GetConfigsParams object
// with the ability to set a timeout on a request.
func NewGetConfigsParamsWithTimeout(timeout time.Duration) *GetConfigsParams {
	return &GetConfigsParams{
		timeout: timeout,
	}
}

// NewGetConfigsParamsWithContext creates a new GetConfigsParams object
// with the ability to set a context for a request.
func NewGetConfigsParamsWithContext(ctx context.Context) *GetConfigsParams {
	return &GetConfigsParams{
		Context: ctx,
	}
}

// NewGetConfigsParamsWithHTTPClient creates a new GetConfigsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetConfigsParamsWithHTTPClient(client *http.Client) *GetConfigsParams {
	return &GetConfigsParams{
		HTTPClient: client,
	}
}

/*
GetConfigsParams contains all the parameters to send to the API endpoint

	for the get configs operation.

	Typically these are written to a http.Request.
*/
type GetConfigsParams struct {

	/* ConfigID.

	   Config Id
	*/
	ConfigID *string

	/* DriftType.

	   Drift Type
	*/
	DriftType *string

	/* Page.

	   Page number to retrieve. Default page 0 will retrieve all elements. Optional

	   Format: int32
	*/
	Page *int32

	/* ResourceID.

	   Resource Id
	*/
	ResourceID *string

	/* ResourceType.

	   Resource Type
	*/
	ResourceType *string

	/* Size.

	   Size of the page to retrieve. Default page size is 10. Optional

	   Format: int32
	   Default: 10
	*/
	Size *int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get configs params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetConfigsParams) WithDefaults() *GetConfigsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get configs params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetConfigsParams) SetDefaults() {
	var (
		pageDefault = int32(0)

		sizeDefault = int32(10)
	)

	val := GetConfigsParams{
		Page: &pageDefault,
		Size: &sizeDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get configs params
func (o *GetConfigsParams) WithTimeout(timeout time.Duration) *GetConfigsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get configs params
func (o *GetConfigsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get configs params
func (o *GetConfigsParams) WithContext(ctx context.Context) *GetConfigsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get configs params
func (o *GetConfigsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get configs params
func (o *GetConfigsParams) WithHTTPClient(client *http.Client) *GetConfigsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get configs params
func (o *GetConfigsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithConfigID adds the configID to the get configs params
func (o *GetConfigsParams) WithConfigID(configID *string) *GetConfigsParams {
	o.SetConfigID(configID)
	return o
}

// SetConfigID adds the configId to the get configs params
func (o *GetConfigsParams) SetConfigID(configID *string) {
	o.ConfigID = configID
}

// WithDriftType adds the driftType to the get configs params
func (o *GetConfigsParams) WithDriftType(driftType *string) *GetConfigsParams {
	o.SetDriftType(driftType)
	return o
}

// SetDriftType adds the driftType to the get configs params
func (o *GetConfigsParams) SetDriftType(driftType *string) {
	o.DriftType = driftType
}

// WithPage adds the page to the get configs params
func (o *GetConfigsParams) WithPage(page *int32) *GetConfigsParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the get configs params
func (o *GetConfigsParams) SetPage(page *int32) {
	o.Page = page
}

// WithResourceID adds the resourceID to the get configs params
func (o *GetConfigsParams) WithResourceID(resourceID *string) *GetConfigsParams {
	o.SetResourceID(resourceID)
	return o
}

// SetResourceID adds the resourceId to the get configs params
func (o *GetConfigsParams) SetResourceID(resourceID *string) {
	o.ResourceID = resourceID
}

// WithResourceType adds the resourceType to the get configs params
func (o *GetConfigsParams) WithResourceType(resourceType *string) *GetConfigsParams {
	o.SetResourceType(resourceType)
	return o
}

// SetResourceType adds the resourceType to the get configs params
func (o *GetConfigsParams) SetResourceType(resourceType *string) {
	o.ResourceType = resourceType
}

// WithSize adds the size to the get configs params
func (o *GetConfigsParams) WithSize(size *int32) *GetConfigsParams {
	o.SetSize(size)
	return o
}

// SetSize adds the size to the get configs params
func (o *GetConfigsParams) SetSize(size *int32) {
	o.Size = size
}

// WriteToRequest writes these params to a swagger request
func (o *GetConfigsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ConfigID != nil {

		// query param configId
		var qrConfigID string

		if o.ConfigID != nil {
			qrConfigID = *o.ConfigID
		}
		qConfigID := qrConfigID
		if qConfigID != "" {

			if err := r.SetQueryParam("configId", qConfigID); err != nil {
				return err
			}
		}
	}

	if o.DriftType != nil {

		// query param driftType
		var qrDriftType string

		if o.DriftType != nil {
			qrDriftType = *o.DriftType
		}
		qDriftType := qrDriftType
		if qDriftType != "" {

			if err := r.SetQueryParam("driftType", qDriftType); err != nil {
				return err
			}
		}
	}

	if o.Page != nil {

		// query param page
		var qrPage int32

		if o.Page != nil {
			qrPage = *o.Page
		}
		qPage := swag.FormatInt32(qrPage)
		if qPage != "" {

			if err := r.SetQueryParam("page", qPage); err != nil {
				return err
			}
		}
	}

	if o.ResourceID != nil {

		// query param resourceId
		var qrResourceID string

		if o.ResourceID != nil {
			qrResourceID = *o.ResourceID
		}
		qResourceID := qrResourceID
		if qResourceID != "" {

			if err := r.SetQueryParam("resourceId", qResourceID); err != nil {
				return err
			}
		}
	}

	if o.ResourceType != nil {

		// query param resourceType
		var qrResourceType string

		if o.ResourceType != nil {
			qrResourceType = *o.ResourceType
		}
		qResourceType := qrResourceType
		if qResourceType != "" {

			if err := r.SetQueryParam("resourceType", qResourceType); err != nil {
				return err
			}
		}
	}

	if o.Size != nil {

		// query param size
		var qrSize int32

		if o.Size != nil {
			qrSize = *o.Size
		}
		qSize := swag.FormatInt32(qrSize)
		if qSize != "" {

			if err := r.SetQueryParam("size", qSize); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}