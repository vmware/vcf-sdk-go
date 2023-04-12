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

// NewGetDatastoreQueryResponseParams creates a new GetDatastoreQueryResponseParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetDatastoreQueryResponseParams() *GetDatastoreQueryResponseParams {
	return &GetDatastoreQueryResponseParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetDatastoreQueryResponseParamsWithTimeout creates a new GetDatastoreQueryResponseParams object
// with the ability to set a timeout on a request.
func NewGetDatastoreQueryResponseParamsWithTimeout(timeout time.Duration) *GetDatastoreQueryResponseParams {
	return &GetDatastoreQueryResponseParams{
		timeout: timeout,
	}
}

// NewGetDatastoreQueryResponseParamsWithContext creates a new GetDatastoreQueryResponseParams object
// with the ability to set a context for a request.
func NewGetDatastoreQueryResponseParamsWithContext(ctx context.Context) *GetDatastoreQueryResponseParams {
	return &GetDatastoreQueryResponseParams{
		Context: ctx,
	}
}

// NewGetDatastoreQueryResponseParamsWithHTTPClient creates a new GetDatastoreQueryResponseParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetDatastoreQueryResponseParamsWithHTTPClient(client *http.Client) *GetDatastoreQueryResponseParams {
	return &GetDatastoreQueryResponseParams{
		HTTPClient: client,
	}
}

/*
GetDatastoreQueryResponseParams contains all the parameters to send to the API endpoint

	for the get datastore query response operation.

	Typically these are written to a http.Request.
*/
type GetDatastoreQueryResponseParams struct {

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

// WithDefaults hydrates default values in the get datastore query response params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDatastoreQueryResponseParams) WithDefaults() *GetDatastoreQueryResponseParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get datastore query response params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDatastoreQueryResponseParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get datastore query response params
func (o *GetDatastoreQueryResponseParams) WithTimeout(timeout time.Duration) *GetDatastoreQueryResponseParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get datastore query response params
func (o *GetDatastoreQueryResponseParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get datastore query response params
func (o *GetDatastoreQueryResponseParams) WithContext(ctx context.Context) *GetDatastoreQueryResponseParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get datastore query response params
func (o *GetDatastoreQueryResponseParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get datastore query response params
func (o *GetDatastoreQueryResponseParams) WithHTTPClient(client *http.Client) *GetDatastoreQueryResponseParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get datastore query response params
func (o *GetDatastoreQueryResponseParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClusterID adds the clusterID to the get datastore query response params
func (o *GetDatastoreQueryResponseParams) WithClusterID(clusterID string) *GetDatastoreQueryResponseParams {
	o.SetClusterID(clusterID)
	return o
}

// SetClusterID adds the clusterId to the get datastore query response params
func (o *GetDatastoreQueryResponseParams) SetClusterID(clusterID string) {
	o.ClusterID = clusterID
}

// WithQueryID adds the queryID to the get datastore query response params
func (o *GetDatastoreQueryResponseParams) WithQueryID(queryID string) *GetDatastoreQueryResponseParams {
	o.SetQueryID(queryID)
	return o
}

// SetQueryID adds the queryId to the get datastore query response params
func (o *GetDatastoreQueryResponseParams) SetQueryID(queryID string) {
	o.QueryID = queryID
}

// WriteToRequest writes these params to a swagger request
func (o *GetDatastoreQueryResponseParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
