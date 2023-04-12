// Code generated by go-swagger; DO NOT EDIT.

package tasks

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

// NewCancelTaskParams creates a new CancelTaskParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCancelTaskParams() *CancelTaskParams {
	return &CancelTaskParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCancelTaskParamsWithTimeout creates a new CancelTaskParams object
// with the ability to set a timeout on a request.
func NewCancelTaskParamsWithTimeout(timeout time.Duration) *CancelTaskParams {
	return &CancelTaskParams{
		timeout: timeout,
	}
}

// NewCancelTaskParamsWithContext creates a new CancelTaskParams object
// with the ability to set a context for a request.
func NewCancelTaskParamsWithContext(ctx context.Context) *CancelTaskParams {
	return &CancelTaskParams{
		Context: ctx,
	}
}

// NewCancelTaskParamsWithHTTPClient creates a new CancelTaskParams object
// with the ability to set a custom HTTPClient for a request.
func NewCancelTaskParamsWithHTTPClient(client *http.Client) *CancelTaskParams {
	return &CancelTaskParams{
		HTTPClient: client,
	}
}

/*
CancelTaskParams contains all the parameters to send to the API endpoint

	for the cancel task operation.

	Typically these are written to a http.Request.
*/
type CancelTaskParams struct {

	/* ID.

	   Task id for cancelling
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the cancel task params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CancelTaskParams) WithDefaults() *CancelTaskParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the cancel task params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CancelTaskParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the cancel task params
func (o *CancelTaskParams) WithTimeout(timeout time.Duration) *CancelTaskParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the cancel task params
func (o *CancelTaskParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the cancel task params
func (o *CancelTaskParams) WithContext(ctx context.Context) *CancelTaskParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the cancel task params
func (o *CancelTaskParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the cancel task params
func (o *CancelTaskParams) WithHTTPClient(client *http.Client) *CancelTaskParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the cancel task params
func (o *CancelTaskParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the cancel task params
func (o *CancelTaskParams) WithID(id string) *CancelTaskParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the cancel task params
func (o *CancelTaskParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *CancelTaskParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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