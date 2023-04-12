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

// NewAssignableTagsToClusterParams creates a new AssignableTagsToClusterParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAssignableTagsToClusterParams() *AssignableTagsToClusterParams {
	return &AssignableTagsToClusterParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAssignableTagsToClusterParamsWithTimeout creates a new AssignableTagsToClusterParams object
// with the ability to set a timeout on a request.
func NewAssignableTagsToClusterParamsWithTimeout(timeout time.Duration) *AssignableTagsToClusterParams {
	return &AssignableTagsToClusterParams{
		timeout: timeout,
	}
}

// NewAssignableTagsToClusterParamsWithContext creates a new AssignableTagsToClusterParams object
// with the ability to set a context for a request.
func NewAssignableTagsToClusterParamsWithContext(ctx context.Context) *AssignableTagsToClusterParams {
	return &AssignableTagsToClusterParams{
		Context: ctx,
	}
}

// NewAssignableTagsToClusterParamsWithHTTPClient creates a new AssignableTagsToClusterParams object
// with the ability to set a custom HTTPClient for a request.
func NewAssignableTagsToClusterParamsWithHTTPClient(client *http.Client) *AssignableTagsToClusterParams {
	return &AssignableTagsToClusterParams{
		HTTPClient: client,
	}
}

/*
AssignableTagsToClusterParams contains all the parameters to send to the API endpoint

	for the assignable tags to cluster operation.

	Typically these are written to a http.Request.
*/
type AssignableTagsToClusterParams struct {

	/* ID.

	   Cluster ID
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the assignable tags to cluster params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AssignableTagsToClusterParams) WithDefaults() *AssignableTagsToClusterParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the assignable tags to cluster params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AssignableTagsToClusterParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the assignable tags to cluster params
func (o *AssignableTagsToClusterParams) WithTimeout(timeout time.Duration) *AssignableTagsToClusterParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the assignable tags to cluster params
func (o *AssignableTagsToClusterParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the assignable tags to cluster params
func (o *AssignableTagsToClusterParams) WithContext(ctx context.Context) *AssignableTagsToClusterParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the assignable tags to cluster params
func (o *AssignableTagsToClusterParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the assignable tags to cluster params
func (o *AssignableTagsToClusterParams) WithHTTPClient(client *http.Client) *AssignableTagsToClusterParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the assignable tags to cluster params
func (o *AssignableTagsToClusterParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the assignable tags to cluster params
func (o *AssignableTagsToClusterParams) WithID(id string) *AssignableTagsToClusterParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the assignable tags to cluster params
func (o *AssignableTagsToClusterParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *AssignableTagsToClusterParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
