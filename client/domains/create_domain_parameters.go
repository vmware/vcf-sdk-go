// Code generated by go-swagger; DO NOT EDIT.

// Copyright © 2023 VMware, Inc. All Rights Reserved.
// Code licenced under: BSD-2

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

	"vcf-sdk-go/models"
)

// NewCreateDomainParams creates a new CreateDomainParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateDomainParams() *CreateDomainParams {
	return &CreateDomainParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateDomainParamsWithTimeout creates a new CreateDomainParams object
// with the ability to set a timeout on a request.
func NewCreateDomainParamsWithTimeout(timeout time.Duration) *CreateDomainParams {
	return &CreateDomainParams{
		timeout: timeout,
	}
}

// NewCreateDomainParamsWithContext creates a new CreateDomainParams object
// with the ability to set a context for a request.
func NewCreateDomainParamsWithContext(ctx context.Context) *CreateDomainParams {
	return &CreateDomainParams{
		Context: ctx,
	}
}

// NewCreateDomainParamsWithHTTPClient creates a new CreateDomainParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateDomainParamsWithHTTPClient(client *http.Client) *CreateDomainParams {
	return &CreateDomainParams{
		HTTPClient: client,
	}
}

/*
CreateDomainParams contains all the parameters to send to the API endpoint

	for the create domain operation.

	Typically these are written to a http.Request.
*/
type CreateDomainParams struct {

	/* DomainCreationSpec.

	   Domain creation data
	*/
	DomainCreationSpec *models.DomainCreationSpec

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create domain params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateDomainParams) WithDefaults() *CreateDomainParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create domain params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateDomainParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create domain params
func (o *CreateDomainParams) WithTimeout(timeout time.Duration) *CreateDomainParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create domain params
func (o *CreateDomainParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create domain params
func (o *CreateDomainParams) WithContext(ctx context.Context) *CreateDomainParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create domain params
func (o *CreateDomainParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create domain params
func (o *CreateDomainParams) WithHTTPClient(client *http.Client) *CreateDomainParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create domain params
func (o *CreateDomainParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDomainCreationSpec adds the domainCreationSpec to the create domain params
func (o *CreateDomainParams) WithDomainCreationSpec(domainCreationSpec *models.DomainCreationSpec) *CreateDomainParams {
	o.SetDomainCreationSpec(domainCreationSpec)
	return o
}

// SetDomainCreationSpec adds the domainCreationSpec to the create domain params
func (o *CreateDomainParams) SetDomainCreationSpec(domainCreationSpec *models.DomainCreationSpec) {
	o.DomainCreationSpec = domainCreationSpec
}

// WriteToRequest writes these params to a swagger request
func (o *CreateDomainParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.DomainCreationSpec != nil {
		if err := r.SetBodyParam(o.DomainCreationSpec); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
