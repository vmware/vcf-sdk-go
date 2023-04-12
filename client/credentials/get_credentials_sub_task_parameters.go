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

// NewGetCredentialsSubTaskParams creates a new GetCredentialsSubTaskParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetCredentialsSubTaskParams() *GetCredentialsSubTaskParams {
	return &GetCredentialsSubTaskParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetCredentialsSubTaskParamsWithTimeout creates a new GetCredentialsSubTaskParams object
// with the ability to set a timeout on a request.
func NewGetCredentialsSubTaskParamsWithTimeout(timeout time.Duration) *GetCredentialsSubTaskParams {
	return &GetCredentialsSubTaskParams{
		timeout: timeout,
	}
}

// NewGetCredentialsSubTaskParamsWithContext creates a new GetCredentialsSubTaskParams object
// with the ability to set a context for a request.
func NewGetCredentialsSubTaskParamsWithContext(ctx context.Context) *GetCredentialsSubTaskParams {
	return &GetCredentialsSubTaskParams{
		Context: ctx,
	}
}

// NewGetCredentialsSubTaskParamsWithHTTPClient creates a new GetCredentialsSubTaskParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetCredentialsSubTaskParamsWithHTTPClient(client *http.Client) *GetCredentialsSubTaskParams {
	return &GetCredentialsSubTaskParams{
		HTTPClient: client,
	}
}

/*
GetCredentialsSubTaskParams contains all the parameters to send to the API endpoint

	for the get credentials sub task operation.

	Typically these are written to a http.Request.
*/
type GetCredentialsSubTaskParams struct {

	/* ID.

	   The ID of the credentials task
	*/
	ID string

	/* SubtaskID.

	   The ID of the credentials sub-task
	*/
	SubtaskID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get credentials sub task params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetCredentialsSubTaskParams) WithDefaults() *GetCredentialsSubTaskParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get credentials sub task params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetCredentialsSubTaskParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get credentials sub task params
func (o *GetCredentialsSubTaskParams) WithTimeout(timeout time.Duration) *GetCredentialsSubTaskParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get credentials sub task params
func (o *GetCredentialsSubTaskParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get credentials sub task params
func (o *GetCredentialsSubTaskParams) WithContext(ctx context.Context) *GetCredentialsSubTaskParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get credentials sub task params
func (o *GetCredentialsSubTaskParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get credentials sub task params
func (o *GetCredentialsSubTaskParams) WithHTTPClient(client *http.Client) *GetCredentialsSubTaskParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get credentials sub task params
func (o *GetCredentialsSubTaskParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get credentials sub task params
func (o *GetCredentialsSubTaskParams) WithID(id string) *GetCredentialsSubTaskParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get credentials sub task params
func (o *GetCredentialsSubTaskParams) SetID(id string) {
	o.ID = id
}

// WithSubtaskID adds the subtaskID to the get credentials sub task params
func (o *GetCredentialsSubTaskParams) WithSubtaskID(subtaskID string) *GetCredentialsSubTaskParams {
	o.SetSubtaskID(subtaskID)
	return o
}

// SetSubtaskID adds the subtaskId to the get credentials sub task params
func (o *GetCredentialsSubTaskParams) SetSubtaskID(subtaskID string) {
	o.SubtaskID = subtaskID
}

// WriteToRequest writes these params to a swagger request
func (o *GetCredentialsSubTaskParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	// path param subtaskId
	if err := r.SetPathParam("subtaskId", o.SubtaskID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}