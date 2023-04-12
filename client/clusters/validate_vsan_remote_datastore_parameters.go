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

	"vcf-sdk-go/models"
)

// NewValidateVsanRemoteDatastoreParams creates a new ValidateVsanRemoteDatastoreParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewValidateVsanRemoteDatastoreParams() *ValidateVsanRemoteDatastoreParams {
	return &ValidateVsanRemoteDatastoreParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewValidateVsanRemoteDatastoreParamsWithTimeout creates a new ValidateVsanRemoteDatastoreParams object
// with the ability to set a timeout on a request.
func NewValidateVsanRemoteDatastoreParamsWithTimeout(timeout time.Duration) *ValidateVsanRemoteDatastoreParams {
	return &ValidateVsanRemoteDatastoreParams{
		timeout: timeout,
	}
}

// NewValidateVsanRemoteDatastoreParamsWithContext creates a new ValidateVsanRemoteDatastoreParams object
// with the ability to set a context for a request.
func NewValidateVsanRemoteDatastoreParamsWithContext(ctx context.Context) *ValidateVsanRemoteDatastoreParams {
	return &ValidateVsanRemoteDatastoreParams{
		Context: ctx,
	}
}

// NewValidateVsanRemoteDatastoreParamsWithHTTPClient creates a new ValidateVsanRemoteDatastoreParams object
// with the ability to set a custom HTTPClient for a request.
func NewValidateVsanRemoteDatastoreParamsWithHTTPClient(client *http.Client) *ValidateVsanRemoteDatastoreParams {
	return &ValidateVsanRemoteDatastoreParams{
		HTTPClient: client,
	}
}

/*
ValidateVsanRemoteDatastoreParams contains all the parameters to send to the API endpoint

	for the validate vsan remote datastore operation.

	Typically these are written to a http.Request.
*/
type ValidateVsanRemoteDatastoreParams struct {

	/* ClusterID.

	   Cluster ID
	*/
	ClusterID string

	/* DatastoreMountSpec.

	   Datastore Mount Spec
	*/
	DatastoreMountSpec *models.DatastoreMountSpec

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the validate vsan remote datastore params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ValidateVsanRemoteDatastoreParams) WithDefaults() *ValidateVsanRemoteDatastoreParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the validate vsan remote datastore params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ValidateVsanRemoteDatastoreParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the validate vsan remote datastore params
func (o *ValidateVsanRemoteDatastoreParams) WithTimeout(timeout time.Duration) *ValidateVsanRemoteDatastoreParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the validate vsan remote datastore params
func (o *ValidateVsanRemoteDatastoreParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the validate vsan remote datastore params
func (o *ValidateVsanRemoteDatastoreParams) WithContext(ctx context.Context) *ValidateVsanRemoteDatastoreParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the validate vsan remote datastore params
func (o *ValidateVsanRemoteDatastoreParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the validate vsan remote datastore params
func (o *ValidateVsanRemoteDatastoreParams) WithHTTPClient(client *http.Client) *ValidateVsanRemoteDatastoreParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the validate vsan remote datastore params
func (o *ValidateVsanRemoteDatastoreParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClusterID adds the clusterID to the validate vsan remote datastore params
func (o *ValidateVsanRemoteDatastoreParams) WithClusterID(clusterID string) *ValidateVsanRemoteDatastoreParams {
	o.SetClusterID(clusterID)
	return o
}

// SetClusterID adds the clusterId to the validate vsan remote datastore params
func (o *ValidateVsanRemoteDatastoreParams) SetClusterID(clusterID string) {
	o.ClusterID = clusterID
}

// WithDatastoreMountSpec adds the datastoreMountSpec to the validate vsan remote datastore params
func (o *ValidateVsanRemoteDatastoreParams) WithDatastoreMountSpec(datastoreMountSpec *models.DatastoreMountSpec) *ValidateVsanRemoteDatastoreParams {
	o.SetDatastoreMountSpec(datastoreMountSpec)
	return o
}

// SetDatastoreMountSpec adds the datastoreMountSpec to the validate vsan remote datastore params
func (o *ValidateVsanRemoteDatastoreParams) SetDatastoreMountSpec(datastoreMountSpec *models.DatastoreMountSpec) {
	o.DatastoreMountSpec = datastoreMountSpec
}

// WriteToRequest writes these params to a swagger request
func (o *ValidateVsanRemoteDatastoreParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param clusterId
	if err := r.SetPathParam("clusterId", o.ClusterID); err != nil {
		return err
	}
	if o.DatastoreMountSpec != nil {
		if err := r.SetBodyParam(o.DatastoreMountSpec); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}