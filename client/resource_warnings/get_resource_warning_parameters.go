// Code generated by go-swagger; DO NOT EDIT.

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
)

// NewGetResourceWarningParams creates a new GetResourceWarningParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetResourceWarningParams() *GetResourceWarningParams {
	return &GetResourceWarningParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetResourceWarningParamsWithTimeout creates a new GetResourceWarningParams object
// with the ability to set a timeout on a request.
func NewGetResourceWarningParamsWithTimeout(timeout time.Duration) *GetResourceWarningParams {
	return &GetResourceWarningParams{
		timeout: timeout,
	}
}

// NewGetResourceWarningParamsWithContext creates a new GetResourceWarningParams object
// with the ability to set a context for a request.
func NewGetResourceWarningParamsWithContext(ctx context.Context) *GetResourceWarningParams {
	return &GetResourceWarningParams{
		Context: ctx,
	}
}

// NewGetResourceWarningParamsWithHTTPClient creates a new GetResourceWarningParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetResourceWarningParamsWithHTTPClient(client *http.Client) *GetResourceWarningParams {
	return &GetResourceWarningParams{
		HTTPClient: client,
	}
}

/*
GetResourceWarningParams contains all the parameters to send to the API endpoint

	for the get resource warning operation.

	Typically these are written to a http.Request.
*/
type GetResourceWarningParams struct {

	/* ID.

	   id
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get resource warning params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetResourceWarningParams) WithDefaults() *GetResourceWarningParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get resource warning params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetResourceWarningParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get resource warning params
func (o *GetResourceWarningParams) WithTimeout(timeout time.Duration) *GetResourceWarningParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get resource warning params
func (o *GetResourceWarningParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get resource warning params
func (o *GetResourceWarningParams) WithContext(ctx context.Context) *GetResourceWarningParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get resource warning params
func (o *GetResourceWarningParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get resource warning params
func (o *GetResourceWarningParams) WithHTTPClient(client *http.Client) *GetResourceWarningParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get resource warning params
func (o *GetResourceWarningParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get resource warning params
func (o *GetResourceWarningParams) WithID(id string) *GetResourceWarningParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get resource warning params
func (o *GetResourceWarningParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetResourceWarningParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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