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
)

// NewGetValidationOfVasaProviderParams creates a new GetValidationOfVasaProviderParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetValidationOfVasaProviderParams() *GetValidationOfVasaProviderParams {
	return &GetValidationOfVasaProviderParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetValidationOfVasaProviderParamsWithTimeout creates a new GetValidationOfVasaProviderParams object
// with the ability to set a timeout on a request.
func NewGetValidationOfVasaProviderParamsWithTimeout(timeout time.Duration) *GetValidationOfVasaProviderParams {
	return &GetValidationOfVasaProviderParams{
		timeout: timeout,
	}
}

// NewGetValidationOfVasaProviderParamsWithContext creates a new GetValidationOfVasaProviderParams object
// with the ability to set a context for a request.
func NewGetValidationOfVasaProviderParamsWithContext(ctx context.Context) *GetValidationOfVasaProviderParams {
	return &GetValidationOfVasaProviderParams{
		Context: ctx,
	}
}

// NewGetValidationOfVasaProviderParamsWithHTTPClient creates a new GetValidationOfVasaProviderParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetValidationOfVasaProviderParamsWithHTTPClient(client *http.Client) *GetValidationOfVasaProviderParams {
	return &GetValidationOfVasaProviderParams{
		HTTPClient: client,
	}
}

/*
GetValidationOfVasaProviderParams contains all the parameters to send to the API endpoint

	for the get validation of vasa provider operation.

	Typically these are written to a http.Request.
*/
type GetValidationOfVasaProviderParams struct {

	/* ID.

	   The validation ID
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get validation of vasa provider params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetValidationOfVasaProviderParams) WithDefaults() *GetValidationOfVasaProviderParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get validation of vasa provider params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetValidationOfVasaProviderParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get validation of vasa provider params
func (o *GetValidationOfVasaProviderParams) WithTimeout(timeout time.Duration) *GetValidationOfVasaProviderParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get validation of vasa provider params
func (o *GetValidationOfVasaProviderParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get validation of vasa provider params
func (o *GetValidationOfVasaProviderParams) WithContext(ctx context.Context) *GetValidationOfVasaProviderParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get validation of vasa provider params
func (o *GetValidationOfVasaProviderParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get validation of vasa provider params
func (o *GetValidationOfVasaProviderParams) WithHTTPClient(client *http.Client) *GetValidationOfVasaProviderParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get validation of vasa provider params
func (o *GetValidationOfVasaProviderParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get validation of vasa provider params
func (o *GetValidationOfVasaProviderParams) WithID(id string) *GetValidationOfVasaProviderParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get validation of vasa provider params
func (o *GetValidationOfVasaProviderParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetValidationOfVasaProviderParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}