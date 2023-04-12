// Code generated by go-swagger; DO NOT EDIT.

package upgrades

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

// NewPerformPrechecksUsingPOSTParams creates a new PerformPrechecksUsingPOSTParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPerformPrechecksUsingPOSTParams() *PerformPrechecksUsingPOSTParams {
	return &PerformPrechecksUsingPOSTParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPerformPrechecksUsingPOSTParamsWithTimeout creates a new PerformPrechecksUsingPOSTParams object
// with the ability to set a timeout on a request.
func NewPerformPrechecksUsingPOSTParamsWithTimeout(timeout time.Duration) *PerformPrechecksUsingPOSTParams {
	return &PerformPrechecksUsingPOSTParams{
		timeout: timeout,
	}
}

// NewPerformPrechecksUsingPOSTParamsWithContext creates a new PerformPrechecksUsingPOSTParams object
// with the ability to set a context for a request.
func NewPerformPrechecksUsingPOSTParamsWithContext(ctx context.Context) *PerformPrechecksUsingPOSTParams {
	return &PerformPrechecksUsingPOSTParams{
		Context: ctx,
	}
}

// NewPerformPrechecksUsingPOSTParamsWithHTTPClient creates a new PerformPrechecksUsingPOSTParams object
// with the ability to set a custom HTTPClient for a request.
func NewPerformPrechecksUsingPOSTParamsWithHTTPClient(client *http.Client) *PerformPrechecksUsingPOSTParams {
	return &PerformPrechecksUsingPOSTParams{
		HTTPClient: client,
	}
}

/*
PerformPrechecksUsingPOSTParams contains all the parameters to send to the API endpoint

	for the perform prechecks using p o s t operation.

	Typically these are written to a http.Request.
*/
type PerformPrechecksUsingPOSTParams struct {

	/* UpgradeID.

	   upgradeId
	*/
	UpgradeID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the perform prechecks using p o s t params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PerformPrechecksUsingPOSTParams) WithDefaults() *PerformPrechecksUsingPOSTParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the perform prechecks using p o s t params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PerformPrechecksUsingPOSTParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the perform prechecks using p o s t params
func (o *PerformPrechecksUsingPOSTParams) WithTimeout(timeout time.Duration) *PerformPrechecksUsingPOSTParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the perform prechecks using p o s t params
func (o *PerformPrechecksUsingPOSTParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the perform prechecks using p o s t params
func (o *PerformPrechecksUsingPOSTParams) WithContext(ctx context.Context) *PerformPrechecksUsingPOSTParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the perform prechecks using p o s t params
func (o *PerformPrechecksUsingPOSTParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the perform prechecks using p o s t params
func (o *PerformPrechecksUsingPOSTParams) WithHTTPClient(client *http.Client) *PerformPrechecksUsingPOSTParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the perform prechecks using p o s t params
func (o *PerformPrechecksUsingPOSTParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUpgradeID adds the upgradeID to the perform prechecks using p o s t params
func (o *PerformPrechecksUsingPOSTParams) WithUpgradeID(upgradeID string) *PerformPrechecksUsingPOSTParams {
	o.SetUpgradeID(upgradeID)
	return o
}

// SetUpgradeID adds the upgradeId to the perform prechecks using p o s t params
func (o *PerformPrechecksUsingPOSTParams) SetUpgradeID(upgradeID string) {
	o.UpgradeID = upgradeID
}

// WriteToRequest writes these params to a swagger request
func (o *PerformPrechecksUsingPOSTParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param upgradeId
	if err := r.SetPathParam("upgradeId", o.UpgradeID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
