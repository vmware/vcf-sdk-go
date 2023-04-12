// Code generated by go-swagger; DO NOT EDIT.

package nsx_t_edge_clusters

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

// NewGetEdgeClustersParams creates a new GetEdgeClustersParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetEdgeClustersParams() *GetEdgeClustersParams {
	return &GetEdgeClustersParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetEdgeClustersParamsWithTimeout creates a new GetEdgeClustersParams object
// with the ability to set a timeout on a request.
func NewGetEdgeClustersParamsWithTimeout(timeout time.Duration) *GetEdgeClustersParams {
	return &GetEdgeClustersParams{
		timeout: timeout,
	}
}

// NewGetEdgeClustersParamsWithContext creates a new GetEdgeClustersParams object
// with the ability to set a context for a request.
func NewGetEdgeClustersParamsWithContext(ctx context.Context) *GetEdgeClustersParams {
	return &GetEdgeClustersParams{
		Context: ctx,
	}
}

// NewGetEdgeClustersParamsWithHTTPClient creates a new GetEdgeClustersParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetEdgeClustersParamsWithHTTPClient(client *http.Client) *GetEdgeClustersParams {
	return &GetEdgeClustersParams{
		HTTPClient: client,
	}
}

/*
GetEdgeClustersParams contains all the parameters to send to the API endpoint

	for the get edge clusters operation.

	Typically these are written to a http.Request.
*/
type GetEdgeClustersParams struct {

	/* ClusterID.

	   Pass an optional vSphere Cluster ID to fetch edge clusters associated with the vSphere Cluster
	*/
	ClusterID *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get edge clusters params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetEdgeClustersParams) WithDefaults() *GetEdgeClustersParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get edge clusters params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetEdgeClustersParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get edge clusters params
func (o *GetEdgeClustersParams) WithTimeout(timeout time.Duration) *GetEdgeClustersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get edge clusters params
func (o *GetEdgeClustersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get edge clusters params
func (o *GetEdgeClustersParams) WithContext(ctx context.Context) *GetEdgeClustersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get edge clusters params
func (o *GetEdgeClustersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get edge clusters params
func (o *GetEdgeClustersParams) WithHTTPClient(client *http.Client) *GetEdgeClustersParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get edge clusters params
func (o *GetEdgeClustersParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClusterID adds the clusterID to the get edge clusters params
func (o *GetEdgeClustersParams) WithClusterID(clusterID *string) *GetEdgeClustersParams {
	o.SetClusterID(clusterID)
	return o
}

// SetClusterID adds the clusterId to the get edge clusters params
func (o *GetEdgeClustersParams) SetClusterID(clusterID *string) {
	o.ClusterID = clusterID
}

// WriteToRequest writes these params to a swagger request
func (o *GetEdgeClustersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ClusterID != nil {

		// query param clusterId
		var qrClusterID string

		if o.ClusterID != nil {
			qrClusterID = *o.ClusterID
		}
		qClusterID := qrClusterID
		if qClusterID != "" {

			if err := r.SetQueryParam("clusterId", qClusterID); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
