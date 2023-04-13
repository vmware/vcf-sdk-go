// Code generated by go-swagger; DO NOT EDIT.

// Copyright © 2023 VMware, Inc. All Rights Reserved.
// Code licenced under: BSD-2

package resource_warnings

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

// NewGETResourceWarningsParams creates a new GETResourceWarningsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGETResourceWarningsParams() *GETResourceWarningsParams {
	return &GETResourceWarningsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGETResourceWarningsParamsWithTimeout creates a new GETResourceWarningsParams object
// with the ability to set a timeout on a request.
func NewGETResourceWarningsParamsWithTimeout(timeout time.Duration) *GETResourceWarningsParams {
	return &GETResourceWarningsParams{
		timeout: timeout,
	}
}

// NewGETResourceWarningsParamsWithContext creates a new GETResourceWarningsParams object
// with the ability to set a context for a request.
func NewGETResourceWarningsParamsWithContext(ctx context.Context) *GETResourceWarningsParams {
	return &GETResourceWarningsParams{
		Context: ctx,
	}
}

// NewGETResourceWarningsParamsWithHTTPClient creates a new GETResourceWarningsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGETResourceWarningsParamsWithHTTPClient(client *http.Client) *GETResourceWarningsParams {
	return &GETResourceWarningsParams{
		HTTPClient: client,
	}
}

/*
GETResourceWarningsParams contains all the parameters to send to the API endpoint

	for the get resource warnings operation.

	Typically these are written to a http.Request.
*/
type GETResourceWarningsParams struct {

	/* ResourceIds.

	   Resource IDs
	*/
	ResourceIds []string

	/* ResourceNames.

	   Resource Names
	*/
	ResourceNames []string

	/* ResourceType.

	   Resource type
	*/
	ResourceType *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get resource warnings params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GETResourceWarningsParams) WithDefaults() *GETResourceWarningsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get resource warnings params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GETResourceWarningsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get resource warnings params
func (o *GETResourceWarningsParams) WithTimeout(timeout time.Duration) *GETResourceWarningsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get resource warnings params
func (o *GETResourceWarningsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get resource warnings params
func (o *GETResourceWarningsParams) WithContext(ctx context.Context) *GETResourceWarningsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get resource warnings params
func (o *GETResourceWarningsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get resource warnings params
func (o *GETResourceWarningsParams) WithHTTPClient(client *http.Client) *GETResourceWarningsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get resource warnings params
func (o *GETResourceWarningsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithResourceIds adds the resourceIds to the get resource warnings params
func (o *GETResourceWarningsParams) WithResourceIds(resourceIds []string) *GETResourceWarningsParams {
	o.SetResourceIds(resourceIds)
	return o
}

// SetResourceIds adds the resourceIds to the get resource warnings params
func (o *GETResourceWarningsParams) SetResourceIds(resourceIds []string) {
	o.ResourceIds = resourceIds
}

// WithResourceNames adds the resourceNames to the get resource warnings params
func (o *GETResourceWarningsParams) WithResourceNames(resourceNames []string) *GETResourceWarningsParams {
	o.SetResourceNames(resourceNames)
	return o
}

// SetResourceNames adds the resourceNames to the get resource warnings params
func (o *GETResourceWarningsParams) SetResourceNames(resourceNames []string) {
	o.ResourceNames = resourceNames
}

// WithResourceType adds the resourceType to the get resource warnings params
func (o *GETResourceWarningsParams) WithResourceType(resourceType *string) *GETResourceWarningsParams {
	o.SetResourceType(resourceType)
	return o
}

// SetResourceType adds the resourceType to the get resource warnings params
func (o *GETResourceWarningsParams) SetResourceType(resourceType *string) {
	o.ResourceType = resourceType
}

// WriteToRequest writes these params to a swagger request
func (o *GETResourceWarningsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ResourceIds != nil {

		// binding items for resourceIds
		joinedResourceIds := o.bindParamResourceIds(reg)

		// query array param resourceIds
		if err := r.SetQueryParam("resourceIds", joinedResourceIds...); err != nil {
			return err
		}
	}

	if o.ResourceNames != nil {

		// binding items for resourceNames
		joinedResourceNames := o.bindParamResourceNames(reg)

		// query array param resourceNames
		if err := r.SetQueryParam("resourceNames", joinedResourceNames...); err != nil {
			return err
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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamGETResourceWarnings binds the parameter resourceIds
func (o *GETResourceWarningsParams) bindParamResourceIds(formats strfmt.Registry) []string {
	resourceIdsIR := o.ResourceIds

	var resourceIdsIC []string
	for _, resourceIdsIIR := range resourceIdsIR { // explode []string

		resourceIdsIIV := resourceIdsIIR // string as string
		resourceIdsIC = append(resourceIdsIC, resourceIdsIIV)
	}

	// items.CollectionFormat: "multi"
	resourceIdsIS := swag.JoinByFormat(resourceIdsIC, "multi")

	return resourceIdsIS
}

// bindParamGETResourceWarnings binds the parameter resourceNames
func (o *GETResourceWarningsParams) bindParamResourceNames(formats strfmt.Registry) []string {
	resourceNamesIR := o.ResourceNames

	var resourceNamesIC []string
	for _, resourceNamesIIR := range resourceNamesIR { // explode []string

		resourceNamesIIV := resourceNamesIIR // string as string
		resourceNamesIC = append(resourceNamesIC, resourceNamesIIV)
	}

	// items.CollectionFormat: "multi"
	resourceNamesIS := swag.JoinByFormat(resourceNamesIC, "multi")

	return resourceNamesIS
}
