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

// NewGetValidationForCommissionHostsParams creates a new GetValidationForCommissionHostsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetValidationForCommissionHostsParams() *GetValidationForCommissionHostsParams {
	return &GetValidationForCommissionHostsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetValidationForCommissionHostsParamsWithTimeout creates a new GetValidationForCommissionHostsParams object
// with the ability to set a timeout on a request.
func NewGetValidationForCommissionHostsParamsWithTimeout(timeout time.Duration) *GetValidationForCommissionHostsParams {
	return &GetValidationForCommissionHostsParams{
		timeout: timeout,
	}
}

// NewGetValidationForCommissionHostsParamsWithContext creates a new GetValidationForCommissionHostsParams object
// with the ability to set a context for a request.
func NewGetValidationForCommissionHostsParamsWithContext(ctx context.Context) *GetValidationForCommissionHostsParams {
	return &GetValidationForCommissionHostsParams{
		Context: ctx,
	}
}

// NewGetValidationForCommissionHostsParamsWithHTTPClient creates a new GetValidationForCommissionHostsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetValidationForCommissionHostsParamsWithHTTPClient(client *http.Client) *GetValidationForCommissionHostsParams {
	return &GetValidationForCommissionHostsParams{
		HTTPClient: client,
	}
}

/*
GetValidationForCommissionHostsParams contains all the parameters to send to the API endpoint

	for the get validation for commission hosts operation.

	Typically these are written to a http.Request.
*/
type GetValidationForCommissionHostsParams struct {

	/* ID.

	   The validation ID
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get validation for commission hosts params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetValidationForCommissionHostsParams) WithDefaults() *GetValidationForCommissionHostsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get validation for commission hosts params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetValidationForCommissionHostsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get validation for commission hosts params
func (o *GetValidationForCommissionHostsParams) WithTimeout(timeout time.Duration) *GetValidationForCommissionHostsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get validation for commission hosts params
func (o *GetValidationForCommissionHostsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get validation for commission hosts params
func (o *GetValidationForCommissionHostsParams) WithContext(ctx context.Context) *GetValidationForCommissionHostsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get validation for commission hosts params
func (o *GetValidationForCommissionHostsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get validation for commission hosts params
func (o *GetValidationForCommissionHostsParams) WithHTTPClient(client *http.Client) *GetValidationForCommissionHostsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get validation for commission hosts params
func (o *GetValidationForCommissionHostsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get validation for commission hosts params
func (o *GetValidationForCommissionHostsParams) WithID(id string) *GetValidationForCommissionHostsParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get validation for commission hosts params
func (o *GetValidationForCommissionHostsParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetValidationForCommissionHostsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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