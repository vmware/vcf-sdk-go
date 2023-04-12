// Code generated by go-swagger; DO NOT EDIT.

package domains

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

// NewGetClusterCriteriaParams creates a new GetClusterCriteriaParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetClusterCriteriaParams() *GetClusterCriteriaParams {
	return &GetClusterCriteriaParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetClusterCriteriaParamsWithTimeout creates a new GetClusterCriteriaParams object
// with the ability to set a timeout on a request.
func NewGetClusterCriteriaParamsWithTimeout(timeout time.Duration) *GetClusterCriteriaParams {
	return &GetClusterCriteriaParams{
		timeout: timeout,
	}
}

// NewGetClusterCriteriaParamsWithContext creates a new GetClusterCriteriaParams object
// with the ability to set a context for a request.
func NewGetClusterCriteriaParamsWithContext(ctx context.Context) *GetClusterCriteriaParams {
	return &GetClusterCriteriaParams{
		Context: ctx,
	}
}

// NewGetClusterCriteriaParamsWithHTTPClient creates a new GetClusterCriteriaParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetClusterCriteriaParamsWithHTTPClient(client *http.Client) *GetClusterCriteriaParams {
	return &GetClusterCriteriaParams{
		HTTPClient: client,
	}
}

/*
GetClusterCriteriaParams contains all the parameters to send to the API endpoint

	for the get cluster criteria operation.

	Typically these are written to a http.Request.
*/
type GetClusterCriteriaParams struct {

	/* DomainID.

	   Domain ID
	*/
	DomainID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get cluster criteria params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetClusterCriteriaParams) WithDefaults() *GetClusterCriteriaParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get cluster criteria params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetClusterCriteriaParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get cluster criteria params
func (o *GetClusterCriteriaParams) WithTimeout(timeout time.Duration) *GetClusterCriteriaParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get cluster criteria params
func (o *GetClusterCriteriaParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get cluster criteria params
func (o *GetClusterCriteriaParams) WithContext(ctx context.Context) *GetClusterCriteriaParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get cluster criteria params
func (o *GetClusterCriteriaParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get cluster criteria params
func (o *GetClusterCriteriaParams) WithHTTPClient(client *http.Client) *GetClusterCriteriaParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get cluster criteria params
func (o *GetClusterCriteriaParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDomainID adds the domainID to the get cluster criteria params
func (o *GetClusterCriteriaParams) WithDomainID(domainID string) *GetClusterCriteriaParams {
	o.SetDomainID(domainID)
	return o
}

// SetDomainID adds the domainId to the get cluster criteria params
func (o *GetClusterCriteriaParams) SetDomainID(domainID string) {
	o.DomainID = domainID
}

// WriteToRequest writes these params to a swagger request
func (o *GetClusterCriteriaParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
