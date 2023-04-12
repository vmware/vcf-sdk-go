// Code generated by go-swagger; DO NOT EDIT.

package license_keys

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

// NewDeleteLicenseKeyParams creates a new DeleteLicenseKeyParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteLicenseKeyParams() *DeleteLicenseKeyParams {
	return &DeleteLicenseKeyParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteLicenseKeyParamsWithTimeout creates a new DeleteLicenseKeyParams object
// with the ability to set a timeout on a request.
func NewDeleteLicenseKeyParamsWithTimeout(timeout time.Duration) *DeleteLicenseKeyParams {
	return &DeleteLicenseKeyParams{
		timeout: timeout,
	}
}

// NewDeleteLicenseKeyParamsWithContext creates a new DeleteLicenseKeyParams object
// with the ability to set a context for a request.
func NewDeleteLicenseKeyParamsWithContext(ctx context.Context) *DeleteLicenseKeyParams {
	return &DeleteLicenseKeyParams{
		Context: ctx,
	}
}

// NewDeleteLicenseKeyParamsWithHTTPClient creates a new DeleteLicenseKeyParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteLicenseKeyParamsWithHTTPClient(client *http.Client) *DeleteLicenseKeyParams {
	return &DeleteLicenseKeyParams{
		HTTPClient: client,
	}
}

/*
DeleteLicenseKeyParams contains all the parameters to send to the API endpoint

	for the delete license key operation.

	Typically these are written to a http.Request.
*/
type DeleteLicenseKeyParams struct {

	/* Key.

	   The 29 alpha numeric character license key with hyphens
	*/
	Key string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete license key params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteLicenseKeyParams) WithDefaults() *DeleteLicenseKeyParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete license key params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteLicenseKeyParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete license key params
func (o *DeleteLicenseKeyParams) WithTimeout(timeout time.Duration) *DeleteLicenseKeyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete license key params
func (o *DeleteLicenseKeyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete license key params
func (o *DeleteLicenseKeyParams) WithContext(ctx context.Context) *DeleteLicenseKeyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete license key params
func (o *DeleteLicenseKeyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete license key params
func (o *DeleteLicenseKeyParams) WithHTTPClient(client *http.Client) *DeleteLicenseKeyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete license key params
func (o *DeleteLicenseKeyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithKey adds the key to the delete license key params
func (o *DeleteLicenseKeyParams) WithKey(key string) *DeleteLicenseKeyParams {
	o.SetKey(key)
	return o
}

// SetKey adds the key to the delete license key params
func (o *DeleteLicenseKeyParams) SetKey(key string) {
	o.Key = key
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteLicenseKeyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param key
	if err := r.SetPathParam("key", o.Key); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
