// Code generated by go-swagger; DO NOT EDIT.

package hosts

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

// NewGetCriterionParams creates a new GetCriterionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetCriterionParams() *GetCriterionParams {
	return &GetCriterionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetCriterionParamsWithTimeout creates a new GetCriterionParams object
// with the ability to set a timeout on a request.
func NewGetCriterionParamsWithTimeout(timeout time.Duration) *GetCriterionParams {
	return &GetCriterionParams{
		timeout: timeout,
	}
}

// NewGetCriterionParamsWithContext creates a new GetCriterionParams object
// with the ability to set a context for a request.
func NewGetCriterionParamsWithContext(ctx context.Context) *GetCriterionParams {
	return &GetCriterionParams{
		Context: ctx,
	}
}

// NewGetCriterionParamsWithHTTPClient creates a new GetCriterionParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetCriterionParamsWithHTTPClient(client *http.Client) *GetCriterionParams {
	return &GetCriterionParams{
		HTTPClient: client,
	}
}

/*
GetCriterionParams contains all the parameters to send to the API endpoint

	for the get criterion operation.

	Typically these are written to a http.Request.
*/
type GetCriterionParams struct {

	/* Name.

	   name
	*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get criterion params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetCriterionParams) WithDefaults() *GetCriterionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get criterion params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetCriterionParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get criterion params
func (o *GetCriterionParams) WithTimeout(timeout time.Duration) *GetCriterionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get criterion params
func (o *GetCriterionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get criterion params
func (o *GetCriterionParams) WithContext(ctx context.Context) *GetCriterionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get criterion params
func (o *GetCriterionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get criterion params
func (o *GetCriterionParams) WithHTTPClient(client *http.Client) *GetCriterionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get criterion params
func (o *GetCriterionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the get criterion params
func (o *GetCriterionParams) WithName(name string) *GetCriterionParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the get criterion params
func (o *GetCriterionParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *GetCriterionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
