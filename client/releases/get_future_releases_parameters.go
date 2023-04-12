// Code generated by go-swagger; DO NOT EDIT.

package releases

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

// NewGetFutureReleasesParams creates a new GetFutureReleasesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetFutureReleasesParams() *GetFutureReleasesParams {
	return &GetFutureReleasesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetFutureReleasesParamsWithTimeout creates a new GetFutureReleasesParams object
// with the ability to set a timeout on a request.
func NewGetFutureReleasesParamsWithTimeout(timeout time.Duration) *GetFutureReleasesParams {
	return &GetFutureReleasesParams{
		timeout: timeout,
	}
}

// NewGetFutureReleasesParamsWithContext creates a new GetFutureReleasesParams object
// with the ability to set a context for a request.
func NewGetFutureReleasesParamsWithContext(ctx context.Context) *GetFutureReleasesParams {
	return &GetFutureReleasesParams{
		Context: ctx,
	}
}

// NewGetFutureReleasesParamsWithHTTPClient creates a new GetFutureReleasesParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetFutureReleasesParamsWithHTTPClient(client *http.Client) *GetFutureReleasesParams {
	return &GetFutureReleasesParams{
		HTTPClient: client,
	}
}

/*
GetFutureReleasesParams contains all the parameters to send to the API endpoint

	for the get future releases operation.

	Typically these are written to a http.Request.
*/
type GetFutureReleasesParams struct {

	/* DomainID.

	   Domain ID to get all feature target versions
	*/
	DomainID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get future releases params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetFutureReleasesParams) WithDefaults() *GetFutureReleasesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get future releases params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetFutureReleasesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get future releases params
func (o *GetFutureReleasesParams) WithTimeout(timeout time.Duration) *GetFutureReleasesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get future releases params
func (o *GetFutureReleasesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get future releases params
func (o *GetFutureReleasesParams) WithContext(ctx context.Context) *GetFutureReleasesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get future releases params
func (o *GetFutureReleasesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get future releases params
func (o *GetFutureReleasesParams) WithHTTPClient(client *http.Client) *GetFutureReleasesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get future releases params
func (o *GetFutureReleasesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDomainID adds the domainID to the get future releases params
func (o *GetFutureReleasesParams) WithDomainID(domainID string) *GetFutureReleasesParams {
	o.SetDomainID(domainID)
	return o
}

// SetDomainID adds the domainId to the get future releases params
func (o *GetFutureReleasesParams) SetDomainID(domainID string) {
	o.DomainID = domainID
}

// WriteToRequest writes these params to a swagger request
func (o *GetFutureReleasesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param domainId
	if err := r.SetPathParam("domainId", o.DomainID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
