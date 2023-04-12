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

// NewGetClustersQueryResponseParams creates a new GetClustersQueryResponseParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetClustersQueryResponseParams() *GetClustersQueryResponseParams {
	return &GetClustersQueryResponseParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetClustersQueryResponseParamsWithTimeout creates a new GetClustersQueryResponseParams object
// with the ability to set a timeout on a request.
func NewGetClustersQueryResponseParamsWithTimeout(timeout time.Duration) *GetClustersQueryResponseParams {
	return &GetClustersQueryResponseParams{
		timeout: timeout,
	}
}

// NewGetClustersQueryResponseParamsWithContext creates a new GetClustersQueryResponseParams object
// with the ability to set a context for a request.
func NewGetClustersQueryResponseParamsWithContext(ctx context.Context) *GetClustersQueryResponseParams {
	return &GetClustersQueryResponseParams{
		Context: ctx,
	}
}

// NewGetClustersQueryResponseParamsWithHTTPClient creates a new GetClustersQueryResponseParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetClustersQueryResponseParamsWithHTTPClient(client *http.Client) *GetClustersQueryResponseParams {
	return &GetClustersQueryResponseParams{
		HTTPClient: client,
	}
}

/*
GetClustersQueryResponseParams contains all the parameters to send to the API endpoint

	for the get clusters query response operation.

	Typically these are written to a http.Request.
*/
type GetClustersQueryResponseParams struct {

	/* DomainID.

	   Domain ID
	*/
	DomainID string

	/* QueryID.

	   Query ID
	*/
	QueryID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get clusters query response params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetClustersQueryResponseParams) WithDefaults() *GetClustersQueryResponseParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get clusters query response params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetClustersQueryResponseParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get clusters query response params
func (o *GetClustersQueryResponseParams) WithTimeout(timeout time.Duration) *GetClustersQueryResponseParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get clusters query response params
func (o *GetClustersQueryResponseParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get clusters query response params
func (o *GetClustersQueryResponseParams) WithContext(ctx context.Context) *GetClustersQueryResponseParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get clusters query response params
func (o *GetClustersQueryResponseParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get clusters query response params
func (o *GetClustersQueryResponseParams) WithHTTPClient(client *http.Client) *GetClustersQueryResponseParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get clusters query response params
func (o *GetClustersQueryResponseParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDomainID adds the domainID to the get clusters query response params
func (o *GetClustersQueryResponseParams) WithDomainID(domainID string) *GetClustersQueryResponseParams {
	o.SetDomainID(domainID)
	return o
}

// SetDomainID adds the domainId to the get clusters query response params
func (o *GetClustersQueryResponseParams) SetDomainID(domainID string) {
	o.DomainID = domainID
}

// WithQueryID adds the queryID to the get clusters query response params
func (o *GetClustersQueryResponseParams) WithQueryID(queryID string) *GetClustersQueryResponseParams {
	o.SetQueryID(queryID)
	return o
}

// SetQueryID adds the queryId to the get clusters query response params
func (o *GetClustersQueryResponseParams) SetQueryID(queryID string) {
	o.QueryID = queryID
}

// WriteToRequest writes these params to a swagger request
func (o *GetClustersQueryResponseParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param domainId
	if err := r.SetPathParam("domainId", o.DomainID); err != nil {
		return err
	}

	// path param queryId
	if err := r.SetPathParam("queryId", o.QueryID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
