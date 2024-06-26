// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2023 VMware, Inc.
// SPDX-License-Identifier: BSD-2-Clause

package vasa_providers

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

// NewAddVasaProviderStorageContainerParams creates a new AddVasaProviderStorageContainerParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAddVasaProviderStorageContainerParams() *AddVasaProviderStorageContainerParams {
	return &AddVasaProviderStorageContainerParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAddVasaProviderStorageContainerParamsWithTimeout creates a new AddVasaProviderStorageContainerParams object
// with the ability to set a timeout on a request.
func NewAddVasaProviderStorageContainerParamsWithTimeout(timeout time.Duration) *AddVasaProviderStorageContainerParams {
	return &AddVasaProviderStorageContainerParams{
		timeout: timeout,
	}
}

// NewAddVasaProviderStorageContainerParamsWithContext creates a new AddVasaProviderStorageContainerParams object
// with the ability to set a context for a request.
func NewAddVasaProviderStorageContainerParamsWithContext(ctx context.Context) *AddVasaProviderStorageContainerParams {
	return &AddVasaProviderStorageContainerParams{
		Context: ctx,
	}
}

// NewAddVasaProviderStorageContainerParamsWithHTTPClient creates a new AddVasaProviderStorageContainerParams object
// with the ability to set a custom HTTPClient for a request.
func NewAddVasaProviderStorageContainerParamsWithHTTPClient(client *http.Client) *AddVasaProviderStorageContainerParams {
	return &AddVasaProviderStorageContainerParams{
		HTTPClient: client,
	}
}

/*
AddVasaProviderStorageContainerParams contains all the parameters to send to the API endpoint

	for the add vasa provider storage container operation.

	Typically these are written to a http.Request.
*/
type AddVasaProviderStorageContainerParams struct {

	/* ID.

	   VASA Provider ID
	*/
	ID string

	/* StorageContainers.

	   Storage containers data
	*/
	StorageContainers []*models.StorageContainer

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the add vasa provider storage container params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AddVasaProviderStorageContainerParams) WithDefaults() *AddVasaProviderStorageContainerParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the add vasa provider storage container params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AddVasaProviderStorageContainerParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the add vasa provider storage container params
func (o *AddVasaProviderStorageContainerParams) WithTimeout(timeout time.Duration) *AddVasaProviderStorageContainerParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the add vasa provider storage container params
func (o *AddVasaProviderStorageContainerParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the add vasa provider storage container params
func (o *AddVasaProviderStorageContainerParams) WithContext(ctx context.Context) *AddVasaProviderStorageContainerParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the add vasa provider storage container params
func (o *AddVasaProviderStorageContainerParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the add vasa provider storage container params
func (o *AddVasaProviderStorageContainerParams) WithHTTPClient(client *http.Client) *AddVasaProviderStorageContainerParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the add vasa provider storage container params
func (o *AddVasaProviderStorageContainerParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the add vasa provider storage container params
func (o *AddVasaProviderStorageContainerParams) WithID(id string) *AddVasaProviderStorageContainerParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the add vasa provider storage container params
func (o *AddVasaProviderStorageContainerParams) SetID(id string) {
	o.ID = id
}

// WithStorageContainers adds the storageContainers to the add vasa provider storage container params
func (o *AddVasaProviderStorageContainerParams) WithStorageContainers(storageContainers []*models.StorageContainer) *AddVasaProviderStorageContainerParams {
	o.SetStorageContainers(storageContainers)
	return o
}

// SetStorageContainers adds the storageContainers to the add vasa provider storage container params
func (o *AddVasaProviderStorageContainerParams) SetStorageContainers(storageContainers []*models.StorageContainer) {
	o.StorageContainers = storageContainers
}

// WriteToRequest writes these params to a swagger request
func (o *AddVasaProviderStorageContainerParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}
	if o.StorageContainers != nil {
		if err := r.SetBodyParam(o.StorageContainers); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
