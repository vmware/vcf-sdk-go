// Code generated by go-swagger; DO NOT EDIT.

package credentials

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

// NewGetCredentialsTaskResourcesCredentialsParams creates a new GetCredentialsTaskResourcesCredentialsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetCredentialsTaskResourcesCredentialsParams() *GetCredentialsTaskResourcesCredentialsParams {
	return &GetCredentialsTaskResourcesCredentialsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetCredentialsTaskResourcesCredentialsParamsWithTimeout creates a new GetCredentialsTaskResourcesCredentialsParams object
// with the ability to set a timeout on a request.
func NewGetCredentialsTaskResourcesCredentialsParamsWithTimeout(timeout time.Duration) *GetCredentialsTaskResourcesCredentialsParams {
	return &GetCredentialsTaskResourcesCredentialsParams{
		timeout: timeout,
	}
}

// NewGetCredentialsTaskResourcesCredentialsParamsWithContext creates a new GetCredentialsTaskResourcesCredentialsParams object
// with the ability to set a context for a request.
func NewGetCredentialsTaskResourcesCredentialsParamsWithContext(ctx context.Context) *GetCredentialsTaskResourcesCredentialsParams {
	return &GetCredentialsTaskResourcesCredentialsParams{
		Context: ctx,
	}
}

// NewGetCredentialsTaskResourcesCredentialsParamsWithHTTPClient creates a new GetCredentialsTaskResourcesCredentialsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetCredentialsTaskResourcesCredentialsParamsWithHTTPClient(client *http.Client) *GetCredentialsTaskResourcesCredentialsParams {
	return &GetCredentialsTaskResourcesCredentialsParams{
		HTTPClient: client,
	}
}

/*
GetCredentialsTaskResourcesCredentialsParams contains all the parameters to send to the API endpoint

	for the get credentials task resources credentials operation.

	Typically these are written to a http.Request.
*/
type GetCredentialsTaskResourcesCredentialsParams struct {

	/* ID.

	   The ID of the credentials task
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get credentials task resources credentials params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetCredentialsTaskResourcesCredentialsParams) WithDefaults() *GetCredentialsTaskResourcesCredentialsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get credentials task resources credentials params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetCredentialsTaskResourcesCredentialsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get credentials task resources credentials params
func (o *GetCredentialsTaskResourcesCredentialsParams) WithTimeout(timeout time.Duration) *GetCredentialsTaskResourcesCredentialsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get credentials task resources credentials params
func (o *GetCredentialsTaskResourcesCredentialsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get credentials task resources credentials params
func (o *GetCredentialsTaskResourcesCredentialsParams) WithContext(ctx context.Context) *GetCredentialsTaskResourcesCredentialsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get credentials task resources credentials params
func (o *GetCredentialsTaskResourcesCredentialsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get credentials task resources credentials params
func (o *GetCredentialsTaskResourcesCredentialsParams) WithHTTPClient(client *http.Client) *GetCredentialsTaskResourcesCredentialsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get credentials task resources credentials params
func (o *GetCredentialsTaskResourcesCredentialsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get credentials task resources credentials params
func (o *GetCredentialsTaskResourcesCredentialsParams) WithID(id string) *GetCredentialsTaskResourcesCredentialsParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get credentials task resources credentials params
func (o *GetCredentialsTaskResourcesCredentialsParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetCredentialsTaskResourcesCredentialsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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