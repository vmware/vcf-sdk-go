// Code generated by go-swagger; DO NOT EDIT.

package certificates

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

// NewDeleteCaConfigurationParams creates a new DeleteCaConfigurationParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteCaConfigurationParams() *DeleteCaConfigurationParams {
	return &DeleteCaConfigurationParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteCaConfigurationParamsWithTimeout creates a new DeleteCaConfigurationParams object
// with the ability to set a timeout on a request.
func NewDeleteCaConfigurationParamsWithTimeout(timeout time.Duration) *DeleteCaConfigurationParams {
	return &DeleteCaConfigurationParams{
		timeout: timeout,
	}
}

// NewDeleteCaConfigurationParamsWithContext creates a new DeleteCaConfigurationParams object
// with the ability to set a context for a request.
func NewDeleteCaConfigurationParamsWithContext(ctx context.Context) *DeleteCaConfigurationParams {
	return &DeleteCaConfigurationParams{
		Context: ctx,
	}
}

// NewDeleteCaConfigurationParamsWithHTTPClient creates a new DeleteCaConfigurationParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteCaConfigurationParamsWithHTTPClient(client *http.Client) *DeleteCaConfigurationParams {
	return &DeleteCaConfigurationParams{
		HTTPClient: client,
	}
}

/*
DeleteCaConfigurationParams contains all the parameters to send to the API endpoint

	for the delete ca configuration operation.

	Typically these are written to a http.Request.
*/
type DeleteCaConfigurationParams struct {

	/* CaType.

	   The CA type
	*/
	CaType string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete ca configuration params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteCaConfigurationParams) WithDefaults() *DeleteCaConfigurationParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete ca configuration params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteCaConfigurationParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete ca configuration params
func (o *DeleteCaConfigurationParams) WithTimeout(timeout time.Duration) *DeleteCaConfigurationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete ca configuration params
func (o *DeleteCaConfigurationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete ca configuration params
func (o *DeleteCaConfigurationParams) WithContext(ctx context.Context) *DeleteCaConfigurationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete ca configuration params
func (o *DeleteCaConfigurationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete ca configuration params
func (o *DeleteCaConfigurationParams) WithHTTPClient(client *http.Client) *DeleteCaConfigurationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete ca configuration params
func (o *DeleteCaConfigurationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCaType adds the caType to the delete ca configuration params
func (o *DeleteCaConfigurationParams) WithCaType(caType string) *DeleteCaConfigurationParams {
	o.SetCaType(caType)
	return o
}

// SetCaType adds the caType to the delete ca configuration params
func (o *DeleteCaConfigurationParams) SetCaType(caType string) {
	o.CaType = caType
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteCaConfigurationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param caType
	if err := r.SetPathParam("caType", o.CaType); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}