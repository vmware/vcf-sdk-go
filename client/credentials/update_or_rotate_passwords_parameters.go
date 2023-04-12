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

	"vcf-sdk-go/models"
)

// NewUpdateOrRotatePasswordsParams creates a new UpdateOrRotatePasswordsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateOrRotatePasswordsParams() *UpdateOrRotatePasswordsParams {
	return &UpdateOrRotatePasswordsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateOrRotatePasswordsParamsWithTimeout creates a new UpdateOrRotatePasswordsParams object
// with the ability to set a timeout on a request.
func NewUpdateOrRotatePasswordsParamsWithTimeout(timeout time.Duration) *UpdateOrRotatePasswordsParams {
	return &UpdateOrRotatePasswordsParams{
		timeout: timeout,
	}
}

// NewUpdateOrRotatePasswordsParamsWithContext creates a new UpdateOrRotatePasswordsParams object
// with the ability to set a context for a request.
func NewUpdateOrRotatePasswordsParamsWithContext(ctx context.Context) *UpdateOrRotatePasswordsParams {
	return &UpdateOrRotatePasswordsParams{
		Context: ctx,
	}
}

// NewUpdateOrRotatePasswordsParamsWithHTTPClient creates a new UpdateOrRotatePasswordsParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateOrRotatePasswordsParamsWithHTTPClient(client *http.Client) *UpdateOrRotatePasswordsParams {
	return &UpdateOrRotatePasswordsParams{
		HTTPClient: client,
	}
}

/*
UpdateOrRotatePasswordsParams contains all the parameters to send to the API endpoint

	for the update or rotate passwords operation.

	Typically these are written to a http.Request.
*/
type UpdateOrRotatePasswordsParams struct {

	/* CredentialsUpdateSpec.

	   credentialsUpdateSpec
	*/
	CredentialsUpdateSpec *models.CredentialsUpdateSpec

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update or rotate passwords params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateOrRotatePasswordsParams) WithDefaults() *UpdateOrRotatePasswordsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update or rotate passwords params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateOrRotatePasswordsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update or rotate passwords params
func (o *UpdateOrRotatePasswordsParams) WithTimeout(timeout time.Duration) *UpdateOrRotatePasswordsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update or rotate passwords params
func (o *UpdateOrRotatePasswordsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update or rotate passwords params
func (o *UpdateOrRotatePasswordsParams) WithContext(ctx context.Context) *UpdateOrRotatePasswordsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update or rotate passwords params
func (o *UpdateOrRotatePasswordsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update or rotate passwords params
func (o *UpdateOrRotatePasswordsParams) WithHTTPClient(client *http.Client) *UpdateOrRotatePasswordsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update or rotate passwords params
func (o *UpdateOrRotatePasswordsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCredentialsUpdateSpec adds the credentialsUpdateSpec to the update or rotate passwords params
func (o *UpdateOrRotatePasswordsParams) WithCredentialsUpdateSpec(credentialsUpdateSpec *models.CredentialsUpdateSpec) *UpdateOrRotatePasswordsParams {
	o.SetCredentialsUpdateSpec(credentialsUpdateSpec)
	return o
}

// SetCredentialsUpdateSpec adds the credentialsUpdateSpec to the update or rotate passwords params
func (o *UpdateOrRotatePasswordsParams) SetCredentialsUpdateSpec(credentialsUpdateSpec *models.CredentialsUpdateSpec) {
	o.CredentialsUpdateSpec = credentialsUpdateSpec
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateOrRotatePasswordsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.CredentialsUpdateSpec != nil {
		if err := r.SetBodyParam(o.CredentialsUpdateSpec); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
