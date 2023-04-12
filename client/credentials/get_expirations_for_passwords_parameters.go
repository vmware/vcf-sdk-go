// Code generated by go-swagger; DO NOT EDIT.

package credentials

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

// NewGetExpirationsForPasswordsParams creates a new GetExpirationsForPasswordsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetExpirationsForPasswordsParams() *GetExpirationsForPasswordsParams {
	return &GetExpirationsForPasswordsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetExpirationsForPasswordsParamsWithTimeout creates a new GetExpirationsForPasswordsParams object
// with the ability to set a timeout on a request.
func NewGetExpirationsForPasswordsParamsWithTimeout(timeout time.Duration) *GetExpirationsForPasswordsParams {
	return &GetExpirationsForPasswordsParams{
		timeout: timeout,
	}
}

// NewGetExpirationsForPasswordsParamsWithContext creates a new GetExpirationsForPasswordsParams object
// with the ability to set a context for a request.
func NewGetExpirationsForPasswordsParamsWithContext(ctx context.Context) *GetExpirationsForPasswordsParams {
	return &GetExpirationsForPasswordsParams{
		Context: ctx,
	}
}

// NewGetExpirationsForPasswordsParamsWithHTTPClient creates a new GetExpirationsForPasswordsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetExpirationsForPasswordsParamsWithHTTPClient(client *http.Client) *GetExpirationsForPasswordsParams {
	return &GetExpirationsForPasswordsParams{
		HTTPClient: client,
	}
}

/*
GetExpirationsForPasswordsParams contains all the parameters to send to the API endpoint

	for the get expirations for passwords operation.

	Typically these are written to a http.Request.
*/
type GetExpirationsForPasswordsParams struct {

	/* ID.

	   The expiration fetch workflow ID
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get expirations for passwords params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetExpirationsForPasswordsParams) WithDefaults() *GetExpirationsForPasswordsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get expirations for passwords params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetExpirationsForPasswordsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get expirations for passwords params
func (o *GetExpirationsForPasswordsParams) WithTimeout(timeout time.Duration) *GetExpirationsForPasswordsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get expirations for passwords params
func (o *GetExpirationsForPasswordsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get expirations for passwords params
func (o *GetExpirationsForPasswordsParams) WithContext(ctx context.Context) *GetExpirationsForPasswordsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get expirations for passwords params
func (o *GetExpirationsForPasswordsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get expirations for passwords params
func (o *GetExpirationsForPasswordsParams) WithHTTPClient(client *http.Client) *GetExpirationsForPasswordsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get expirations for passwords params
func (o *GetExpirationsForPasswordsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get expirations for passwords params
func (o *GetExpirationsForPasswordsParams) WithID(id string) *GetExpirationsForPasswordsParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get expirations for passwords params
func (o *GetExpirationsForPasswordsParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetExpirationsForPasswordsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
