// Code generated by go-swagger; DO NOT EDIT.

package system

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

// NewGetValidationOfDNSConfigurationParams creates a new GetValidationOfDNSConfigurationParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetValidationOfDNSConfigurationParams() *GetValidationOfDNSConfigurationParams {
	return &GetValidationOfDNSConfigurationParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetValidationOfDNSConfigurationParamsWithTimeout creates a new GetValidationOfDNSConfigurationParams object
// with the ability to set a timeout on a request.
func NewGetValidationOfDNSConfigurationParamsWithTimeout(timeout time.Duration) *GetValidationOfDNSConfigurationParams {
	return &GetValidationOfDNSConfigurationParams{
		timeout: timeout,
	}
}

// NewGetValidationOfDNSConfigurationParamsWithContext creates a new GetValidationOfDNSConfigurationParams object
// with the ability to set a context for a request.
func NewGetValidationOfDNSConfigurationParamsWithContext(ctx context.Context) *GetValidationOfDNSConfigurationParams {
	return &GetValidationOfDNSConfigurationParams{
		Context: ctx,
	}
}

// NewGetValidationOfDNSConfigurationParamsWithHTTPClient creates a new GetValidationOfDNSConfigurationParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetValidationOfDNSConfigurationParamsWithHTTPClient(client *http.Client) *GetValidationOfDNSConfigurationParams {
	return &GetValidationOfDNSConfigurationParams{
		HTTPClient: client,
	}
}

/*
GetValidationOfDNSConfigurationParams contains all the parameters to send to the API endpoint

	for the get validation of Dns configuration operation.

	Typically these are written to a http.Request.
*/
type GetValidationOfDNSConfigurationParams struct {

	/* ID.

	   The validation ID
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get validation of Dns configuration params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetValidationOfDNSConfigurationParams) WithDefaults() *GetValidationOfDNSConfigurationParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get validation of Dns configuration params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetValidationOfDNSConfigurationParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get validation of Dns configuration params
func (o *GetValidationOfDNSConfigurationParams) WithTimeout(timeout time.Duration) *GetValidationOfDNSConfigurationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get validation of Dns configuration params
func (o *GetValidationOfDNSConfigurationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get validation of Dns configuration params
func (o *GetValidationOfDNSConfigurationParams) WithContext(ctx context.Context) *GetValidationOfDNSConfigurationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get validation of Dns configuration params
func (o *GetValidationOfDNSConfigurationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get validation of Dns configuration params
func (o *GetValidationOfDNSConfigurationParams) WithHTTPClient(client *http.Client) *GetValidationOfDNSConfigurationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get validation of Dns configuration params
func (o *GetValidationOfDNSConfigurationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get validation of Dns configuration params
func (o *GetValidationOfDNSConfigurationParams) WithID(id string) *GetValidationOfDNSConfigurationParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get validation of Dns configuration params
func (o *GetValidationOfDNSConfigurationParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetValidationOfDNSConfigurationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
