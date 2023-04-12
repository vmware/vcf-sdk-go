// Code generated by go-swagger; DO NOT EDIT.

package network_pools

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

// NewCreateNetworkPoolParams creates a new CreateNetworkPoolParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateNetworkPoolParams() *CreateNetworkPoolParams {
	return &CreateNetworkPoolParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateNetworkPoolParamsWithTimeout creates a new CreateNetworkPoolParams object
// with the ability to set a timeout on a request.
func NewCreateNetworkPoolParamsWithTimeout(timeout time.Duration) *CreateNetworkPoolParams {
	return &CreateNetworkPoolParams{
		timeout: timeout,
	}
}

// NewCreateNetworkPoolParamsWithContext creates a new CreateNetworkPoolParams object
// with the ability to set a context for a request.
func NewCreateNetworkPoolParamsWithContext(ctx context.Context) *CreateNetworkPoolParams {
	return &CreateNetworkPoolParams{
		Context: ctx,
	}
}

// NewCreateNetworkPoolParamsWithHTTPClient creates a new CreateNetworkPoolParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateNetworkPoolParamsWithHTTPClient(client *http.Client) *CreateNetworkPoolParams {
	return &CreateNetworkPoolParams{
		HTTPClient: client,
	}
}

/*
CreateNetworkPoolParams contains all the parameters to send to the API endpoint

	for the create network pool operation.

	Typically these are written to a http.Request.
*/
type CreateNetworkPoolParams struct {

	/* NetworkPool.

	   Specification of the Network pool to create
	*/
	NetworkPool *models.NetworkPool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create network pool params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateNetworkPoolParams) WithDefaults() *CreateNetworkPoolParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create network pool params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateNetworkPoolParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create network pool params
func (o *CreateNetworkPoolParams) WithTimeout(timeout time.Duration) *CreateNetworkPoolParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create network pool params
func (o *CreateNetworkPoolParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create network pool params
func (o *CreateNetworkPoolParams) WithContext(ctx context.Context) *CreateNetworkPoolParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create network pool params
func (o *CreateNetworkPoolParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create network pool params
func (o *CreateNetworkPoolParams) WithHTTPClient(client *http.Client) *CreateNetworkPoolParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create network pool params
func (o *CreateNetworkPoolParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNetworkPool adds the networkPool to the create network pool params
func (o *CreateNetworkPoolParams) WithNetworkPool(networkPool *models.NetworkPool) *CreateNetworkPoolParams {
	o.SetNetworkPool(networkPool)
	return o
}

// SetNetworkPool adds the networkPool to the create network pool params
func (o *CreateNetworkPoolParams) SetNetworkPool(networkPool *models.NetworkPool) {
	o.NetworkPool = networkPool
}

// WriteToRequest writes these params to a swagger request
func (o *CreateNetworkPoolParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.NetworkPool != nil {
		if err := r.SetBodyParam(o.NetworkPool); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}