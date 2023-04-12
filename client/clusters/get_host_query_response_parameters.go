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

// NewGetHostQueryResponseParams creates a new GetHostQueryResponseParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetHostQueryResponseParams() *GetHostQueryResponseParams {
	return &GetHostQueryResponseParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetHostQueryResponseParamsWithTimeout creates a new GetHostQueryResponseParams object
// with the ability to set a timeout on a request.
func NewGetHostQueryResponseParamsWithTimeout(timeout time.Duration) *GetHostQueryResponseParams {
	return &GetHostQueryResponseParams{
		timeout: timeout,
	}
}

// NewGetHostQueryResponseParamsWithContext creates a new GetHostQueryResponseParams object
// with the ability to set a context for a request.
func NewGetHostQueryResponseParamsWithContext(ctx context.Context) *GetHostQueryResponseParams {
	return &GetHostQueryResponseParams{
		Context: ctx,
	}
}

// NewGetHostQueryResponseParamsWithHTTPClient creates a new GetHostQueryResponseParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetHostQueryResponseParamsWithHTTPClient(client *http.Client) *GetHostQueryResponseParams {
	return &GetHostQueryResponseParams{
		HTTPClient: client,
	}
}

/*
GetHostQueryResponseParams contains all the parameters to send to the API endpoint

	for the get host query response operation.

	Typically these are written to a http.Request.
*/
type GetHostQueryResponseParams struct {

	/* ClusterID.

	   Cluster ID
	*/
	ClusterID string

	/* QueryID.

	   Query ID
	*/
	QueryID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get host query response params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetHostQueryResponseParams) WithDefaults() *GetHostQueryResponseParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get host query response params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetHostQueryResponseParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get host query response params
func (o *GetHostQueryResponseParams) WithTimeout(timeout time.Duration) *GetHostQueryResponseParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get host query response params
func (o *GetHostQueryResponseParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get host query response params
func (o *GetHostQueryResponseParams) WithContext(ctx context.Context) *GetHostQueryResponseParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get host query response params
func (o *GetHostQueryResponseParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get host query response params
func (o *GetHostQueryResponseParams) WithHTTPClient(client *http.Client) *GetHostQueryResponseParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get host query response params
func (o *GetHostQueryResponseParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClusterID adds the clusterID to the get host query response params
func (o *GetHostQueryResponseParams) WithClusterID(clusterID string) *GetHostQueryResponseParams {
	o.SetClusterID(clusterID)
	return o
}

// SetClusterID adds the clusterId to the get host query response params
func (o *GetHostQueryResponseParams) SetClusterID(clusterID string) {
	o.ClusterID = clusterID
}

// WithQueryID adds the queryID to the get host query response params
func (o *GetHostQueryResponseParams) WithQueryID(queryID string) *GetHostQueryResponseParams {
	o.SetQueryID(queryID)
	return o
}

// SetQueryID adds the queryId to the get host query response params
func (o *GetHostQueryResponseParams) SetQueryID(queryID string) {
	o.QueryID = queryID
}

// WriteToRequest writes these params to a swagger request
func (o *GetHostQueryResponseParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param clusterId
	if err := r.SetPathParam("clusterId", o.ClusterID); err != nil {
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