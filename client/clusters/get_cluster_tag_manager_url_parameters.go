// Code generated by go-swagger; DO NOT EDIT.

package clusters

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

// NewGetClusterTagManagerURLParams creates a new GetClusterTagManagerURLParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetClusterTagManagerURLParams() *GetClusterTagManagerURLParams {
	return &GetClusterTagManagerURLParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetClusterTagManagerURLParamsWithTimeout creates a new GetClusterTagManagerURLParams object
// with the ability to set a timeout on a request.
func NewGetClusterTagManagerURLParamsWithTimeout(timeout time.Duration) *GetClusterTagManagerURLParams {
	return &GetClusterTagManagerURLParams{
		timeout: timeout,
	}
}

// NewGetClusterTagManagerURLParamsWithContext creates a new GetClusterTagManagerURLParams object
// with the ability to set a context for a request.
func NewGetClusterTagManagerURLParamsWithContext(ctx context.Context) *GetClusterTagManagerURLParams {
	return &GetClusterTagManagerURLParams{
		Context: ctx,
	}
}

// NewGetClusterTagManagerURLParamsWithHTTPClient creates a new GetClusterTagManagerURLParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetClusterTagManagerURLParamsWithHTTPClient(client *http.Client) *GetClusterTagManagerURLParams {
	return &GetClusterTagManagerURLParams{
		HTTPClient: client,
	}
}

/*
GetClusterTagManagerURLParams contains all the parameters to send to the API endpoint

	for the get cluster tag manager Url operation.

	Typically these are written to a http.Request.
*/
type GetClusterTagManagerURLParams struct {

	/* ID.

	   Cluster ID
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get cluster tag manager Url params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetClusterTagManagerURLParams) WithDefaults() *GetClusterTagManagerURLParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get cluster tag manager Url params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetClusterTagManagerURLParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get cluster tag manager Url params
func (o *GetClusterTagManagerURLParams) WithTimeout(timeout time.Duration) *GetClusterTagManagerURLParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get cluster tag manager Url params
func (o *GetClusterTagManagerURLParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get cluster tag manager Url params
func (o *GetClusterTagManagerURLParams) WithContext(ctx context.Context) *GetClusterTagManagerURLParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get cluster tag manager Url params
func (o *GetClusterTagManagerURLParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get cluster tag manager Url params
func (o *GetClusterTagManagerURLParams) WithHTTPClient(client *http.Client) *GetClusterTagManagerURLParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get cluster tag manager Url params
func (o *GetClusterTagManagerURLParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get cluster tag manager Url params
func (o *GetClusterTagManagerURLParams) WithID(id string) *GetClusterTagManagerURLParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get cluster tag manager Url params
func (o *GetClusterTagManagerURLParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetClusterTagManagerURLParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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