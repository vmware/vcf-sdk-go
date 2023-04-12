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

// NewValidateVasaProviderParams creates a new ValidateVasaProviderParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewValidateVasaProviderParams() *ValidateVasaProviderParams {
	return &ValidateVasaProviderParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewValidateVasaProviderParamsWithTimeout creates a new ValidateVasaProviderParams object
// with the ability to set a timeout on a request.
func NewValidateVasaProviderParamsWithTimeout(timeout time.Duration) *ValidateVasaProviderParams {
	return &ValidateVasaProviderParams{
		timeout: timeout,
	}
}

// NewValidateVasaProviderParamsWithContext creates a new ValidateVasaProviderParams object
// with the ability to set a context for a request.
func NewValidateVasaProviderParamsWithContext(ctx context.Context) *ValidateVasaProviderParams {
	return &ValidateVasaProviderParams{
		Context: ctx,
	}
}

// NewValidateVasaProviderParamsWithHTTPClient creates a new ValidateVasaProviderParams object
// with the ability to set a custom HTTPClient for a request.
func NewValidateVasaProviderParamsWithHTTPClient(client *http.Client) *ValidateVasaProviderParams {
	return &ValidateVasaProviderParams{
		HTTPClient: client,
	}
}

/*
ValidateVasaProviderParams contains all the parameters to send to the API endpoint

	for the validate vasa provider operation.

	Typically these are written to a http.Request.
*/
type ValidateVasaProviderParams struct {

	/* VasaProvider.

	   vasaProvider
	*/
	VasaProvider *models.VasaProvider

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the validate vasa provider params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ValidateVasaProviderParams) WithDefaults() *ValidateVasaProviderParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the validate vasa provider params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ValidateVasaProviderParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the validate vasa provider params
func (o *ValidateVasaProviderParams) WithTimeout(timeout time.Duration) *ValidateVasaProviderParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the validate vasa provider params
func (o *ValidateVasaProviderParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the validate vasa provider params
func (o *ValidateVasaProviderParams) WithContext(ctx context.Context) *ValidateVasaProviderParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the validate vasa provider params
func (o *ValidateVasaProviderParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the validate vasa provider params
func (o *ValidateVasaProviderParams) WithHTTPClient(client *http.Client) *ValidateVasaProviderParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the validate vasa provider params
func (o *ValidateVasaProviderParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithVasaProvider adds the vasaProvider to the validate vasa provider params
func (o *ValidateVasaProviderParams) WithVasaProvider(vasaProvider *models.VasaProvider) *ValidateVasaProviderParams {
	o.SetVasaProvider(vasaProvider)
	return o
}

// SetVasaProvider adds the vasaProvider to the validate vasa provider params
func (o *ValidateVasaProviderParams) SetVasaProvider(vasaProvider *models.VasaProvider) {
	o.VasaProvider = vasaProvider
}

// WriteToRequest writes these params to a swagger request
func (o *ValidateVasaProviderParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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