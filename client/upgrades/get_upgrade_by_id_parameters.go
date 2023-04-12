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

// NewGetUpgradeByIDParams creates a new GetUpgradeByIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetUpgradeByIDParams() *GetUpgradeByIDParams {
	return &GetUpgradeByIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetUpgradeByIDParamsWithTimeout creates a new GetUpgradeByIDParams object
// with the ability to set a timeout on a request.
func NewGetUpgradeByIDParamsWithTimeout(timeout time.Duration) *GetUpgradeByIDParams {
	return &GetUpgradeByIDParams{
		timeout: timeout,
	}
}

// NewGetUpgradeByIDParamsWithContext creates a new GetUpgradeByIDParams object
// with the ability to set a context for a request.
func NewGetUpgradeByIDParamsWithContext(ctx context.Context) *GetUpgradeByIDParams {
	return &GetUpgradeByIDParams{
		Context: ctx,
	}
}

// NewGetUpgradeByIDParamsWithHTTPClient creates a new GetUpgradeByIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetUpgradeByIDParamsWithHTTPClient(client *http.Client) *GetUpgradeByIDParams {
	return &GetUpgradeByIDParams{
		HTTPClient: client,
	}
}

/*
GetUpgradeByIDParams contains all the parameters to send to the API endpoint

	for the get upgrade by Id operation.

	Typically these are written to a http.Request.
*/
type GetUpgradeByIDParams struct {

	/* UpgradeID.

	   upgradeId
	*/
	UpgradeID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get upgrade by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetUpgradeByIDParams) WithDefaults() *GetUpgradeByIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get upgrade by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetUpgradeByIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get upgrade by Id params
func (o *GetUpgradeByIDParams) WithTimeout(timeout time.Duration) *GetUpgradeByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get upgrade by Id params
func (o *GetUpgradeByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get upgrade by Id params
func (o *GetUpgradeByIDParams) WithContext(ctx context.Context) *GetUpgradeByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get upgrade by Id params
func (o *GetUpgradeByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get upgrade by Id params
func (o *GetUpgradeByIDParams) WithHTTPClient(client *http.Client) *GetUpgradeByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get upgrade by Id params
func (o *GetUpgradeByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUpgradeID adds the upgradeID to the get upgrade by Id params
func (o *GetUpgradeByIDParams) WithUpgradeID(upgradeID string) *GetUpgradeByIDParams {
	o.SetUpgradeID(upgradeID)
	return o
}

// SetUpgradeID adds the upgradeId to the get upgrade by Id params
func (o *GetUpgradeByIDParams) SetUpgradeID(upgradeID string) {
	o.UpgradeID = upgradeID
}

// WriteToRequest writes these params to a swagger request
func (o *GetUpgradeByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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