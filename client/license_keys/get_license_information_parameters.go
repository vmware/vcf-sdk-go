// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2023 VMware, Inc.
// SPDX-License-Identifier: BSD-2-Clause

package license_keys

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

// NewGetLicenseInformationParams creates a new GetLicenseInformationParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetLicenseInformationParams() *GetLicenseInformationParams {
	return &GetLicenseInformationParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetLicenseInformationParamsWithTimeout creates a new GetLicenseInformationParams object
// with the ability to set a timeout on a request.
func NewGetLicenseInformationParamsWithTimeout(timeout time.Duration) *GetLicenseInformationParams {
	return &GetLicenseInformationParams{
		timeout: timeout,
	}
}

// NewGetLicenseInformationParamsWithContext creates a new GetLicenseInformationParams object
// with the ability to set a context for a request.
func NewGetLicenseInformationParamsWithContext(ctx context.Context) *GetLicenseInformationParams {
	return &GetLicenseInformationParams{
		Context: ctx,
	}
}

// NewGetLicenseInformationParamsWithHTTPClient creates a new GetLicenseInformationParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetLicenseInformationParamsWithHTTPClient(client *http.Client) *GetLicenseInformationParams {
	return &GetLicenseInformationParams{
		HTTPClient: client,
	}
}

/*
GetLicenseInformationParams contains all the parameters to send to the API endpoint

	for the get license information operation.

	Typically these are written to a http.Request.
*/
type GetLicenseInformationParams struct {

	/* ResourceIds.

	   Resource IDs
	*/
	ResourceIds []string

	/* ResourceType.

	   Resource type
	*/
	ResourceType *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get license information params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetLicenseInformationParams) WithDefaults() *GetLicenseInformationParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get license information params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetLicenseInformationParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get license information params
func (o *GetLicenseInformationParams) WithTimeout(timeout time.Duration) *GetLicenseInformationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get license information params
func (o *GetLicenseInformationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get license information params
func (o *GetLicenseInformationParams) WithContext(ctx context.Context) *GetLicenseInformationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get license information params
func (o *GetLicenseInformationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get license information params
func (o *GetLicenseInformationParams) WithHTTPClient(client *http.Client) *GetLicenseInformationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get license information params
func (o *GetLicenseInformationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithResourceIds adds the resourceIds to the get license information params
func (o *GetLicenseInformationParams) WithResourceIds(resourceIds []string) *GetLicenseInformationParams {
	o.SetResourceIds(resourceIds)
	return o
}

// SetResourceIds adds the resourceIds to the get license information params
func (o *GetLicenseInformationParams) SetResourceIds(resourceIds []string) {
	o.ResourceIds = resourceIds
}

// WithResourceType adds the resourceType to the get license information params
func (o *GetLicenseInformationParams) WithResourceType(resourceType *string) *GetLicenseInformationParams {
	o.SetResourceType(resourceType)
	return o
}

// SetResourceType adds the resourceType to the get license information params
func (o *GetLicenseInformationParams) SetResourceType(resourceType *string) {
	o.ResourceType = resourceType
}

// WriteToRequest writes these params to a swagger request
func (o *GetLicenseInformationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

// bindParamGetLicenseInformation binds the parameter resourceIds
func (o *GetLicenseInformationParams) bindParamResourceIds(formats strfmt.Registry) []string {
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