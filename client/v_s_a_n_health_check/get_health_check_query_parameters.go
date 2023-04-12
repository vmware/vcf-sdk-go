// Code generated by go-swagger; DO NOT EDIT.

package v_s_a_n_health_check

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

// NewGetHealthCheckQueryParams creates a new GetHealthCheckQueryParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetHealthCheckQueryParams() *GetHealthCheckQueryParams {
	return &GetHealthCheckQueryParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetHealthCheckQueryParamsWithTimeout creates a new GetHealthCheckQueryParams object
// with the ability to set a timeout on a request.
func NewGetHealthCheckQueryParamsWithTimeout(timeout time.Duration) *GetHealthCheckQueryParams {
	return &GetHealthCheckQueryParams{
		timeout: timeout,
	}
}

// NewGetHealthCheckQueryParamsWithContext creates a new GetHealthCheckQueryParams object
// with the ability to set a context for a request.
func NewGetHealthCheckQueryParamsWithContext(ctx context.Context) *GetHealthCheckQueryParams {
	return &GetHealthCheckQueryParams{
		Context: ctx,
	}
}

// NewGetHealthCheckQueryParamsWithHTTPClient creates a new GetHealthCheckQueryParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetHealthCheckQueryParamsWithHTTPClient(client *http.Client) *GetHealthCheckQueryParams {
	return &GetHealthCheckQueryParams{
		HTTPClient: client,
	}
}

/*
GetHealthCheckQueryParams contains all the parameters to send to the API endpoint

	for the get health check query operation.

	Typically these are written to a http.Request.
*/
type GetHealthCheckQueryParams struct {

	/* DomainID.

	   Domain ID
	*/
	DomainID string

	/* Status.

	   Status of health check to filter by
	*/
	Status *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get health check query params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetHealthCheckQueryParams) WithDefaults() *GetHealthCheckQueryParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get health check query params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetHealthCheckQueryParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get health check query params
func (o *GetHealthCheckQueryParams) WithTimeout(timeout time.Duration) *GetHealthCheckQueryParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get health check query params
func (o *GetHealthCheckQueryParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get health check query params
func (o *GetHealthCheckQueryParams) WithContext(ctx context.Context) *GetHealthCheckQueryParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get health check query params
func (o *GetHealthCheckQueryParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get health check query params
func (o *GetHealthCheckQueryParams) WithHTTPClient(client *http.Client) *GetHealthCheckQueryParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get health check query params
func (o *GetHealthCheckQueryParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDomainID adds the domainID to the get health check query params
func (o *GetHealthCheckQueryParams) WithDomainID(domainID string) *GetHealthCheckQueryParams {
	o.SetDomainID(domainID)
	return o
}

// SetDomainID adds the domainId to the get health check query params
func (o *GetHealthCheckQueryParams) SetDomainID(domainID string) {
	o.DomainID = domainID
}

// WithStatus adds the status to the get health check query params
func (o *GetHealthCheckQueryParams) WithStatus(status *string) *GetHealthCheckQueryParams {
	o.SetStatus(status)
	return o
}

// SetStatus adds the status to the get health check query params
func (o *GetHealthCheckQueryParams) SetStatus(status *string) {
	o.Status = status
}

// WriteToRequest writes these params to a swagger request
func (o *GetHealthCheckQueryParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param domainId
	if err := r.SetPathParam("domainId", o.DomainID); err != nil {
		return err
	}

	if o.Status != nil {

		// query param status
		var qrStatus string

		if o.Status != nil {
			qrStatus = *o.Status
		}
		qStatus := qrStatus
		if qStatus != "" {

			if err := r.SetQueryParam("status", qStatus); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
