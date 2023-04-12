// Code generated by go-swagger; DO NOT EDIT.

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

	"vcf-sdk-go/models"
)

// NewAddVasaProviderParams creates a new AddVasaProviderParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAddVasaProviderParams() *AddVasaProviderParams {
	return &AddVasaProviderParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAddVasaProviderParamsWithTimeout creates a new AddVasaProviderParams object
// with the ability to set a timeout on a request.
func NewAddVasaProviderParamsWithTimeout(timeout time.Duration) *AddVasaProviderParams {
	return &AddVasaProviderParams{
		timeout: timeout,
	}
}

// NewAddVasaProviderParamsWithContext creates a new AddVasaProviderParams object
// with the ability to set a context for a request.
func NewAddVasaProviderParamsWithContext(ctx context.Context) *AddVasaProviderParams {
	return &AddVasaProviderParams{
		Context: ctx,
	}
}

// NewAddVasaProviderParamsWithHTTPClient creates a new AddVasaProviderParams object
// with the ability to set a custom HTTPClient for a request.
func NewAddVasaProviderParamsWithHTTPClient(client *http.Client) *AddVasaProviderParams {
	return &AddVasaProviderParams{
		HTTPClient: client,
	}
}

/*
AddVasaProviderParams contains all the parameters to send to the API endpoint

	for the add vasa provider operation.

	Typically these are written to a http.Request.
*/
type AddVasaProviderParams struct {

	/* VasaProvider.

	   VASA Provider data
	*/
	VasaProvider *models.VasaProvider

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the add vasa provider params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AddVasaProviderParams) WithDefaults() *AddVasaProviderParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the add vasa provider params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AddVasaProviderParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the add vasa provider params
func (o *AddVasaProviderParams) WithTimeout(timeout time.Duration) *AddVasaProviderParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the add vasa provider params
func (o *AddVasaProviderParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the add vasa provider params
func (o *AddVasaProviderParams) WithContext(ctx context.Context) *AddVasaProviderParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the add vasa provider params
func (o *AddVasaProviderParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the add vasa provider params
func (o *AddVasaProviderParams) WithHTTPClient(client *http.Client) *AddVasaProviderParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the add vasa provider params
func (o *AddVasaProviderParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithVasaProvider adds the vasaProvider to the add vasa provider params
func (o *AddVasaProviderParams) WithVasaProvider(vasaProvider *models.VasaProvider) *AddVasaProviderParams {
	o.SetVasaProvider(vasaProvider)
	return o
}

// SetVasaProvider adds the vasaProvider to the add vasa provider params
func (o *AddVasaProviderParams) SetVasaProvider(vasaProvider *models.VasaProvider) {
	o.VasaProvider = vasaProvider
}

// WriteToRequest writes these params to a swagger request
func (o *AddVasaProviderParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.VasaProvider != nil {
		if err := r.SetBodyParam(o.VasaProvider); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
