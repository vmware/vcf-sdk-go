// Code generated by go-swagger; DO NOT EDIT.

package identity_providers

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

// NewUpdateEmbeddedIdentitySourceParams creates a new UpdateEmbeddedIdentitySourceParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateEmbeddedIdentitySourceParams() *UpdateEmbeddedIdentitySourceParams {
	return &UpdateEmbeddedIdentitySourceParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateEmbeddedIdentitySourceParamsWithTimeout creates a new UpdateEmbeddedIdentitySourceParams object
// with the ability to set a timeout on a request.
func NewUpdateEmbeddedIdentitySourceParamsWithTimeout(timeout time.Duration) *UpdateEmbeddedIdentitySourceParams {
	return &UpdateEmbeddedIdentitySourceParams{
		timeout: timeout,
	}
}

// NewUpdateEmbeddedIdentitySourceParamsWithContext creates a new UpdateEmbeddedIdentitySourceParams object
// with the ability to set a context for a request.
func NewUpdateEmbeddedIdentitySourceParamsWithContext(ctx context.Context) *UpdateEmbeddedIdentitySourceParams {
	return &UpdateEmbeddedIdentitySourceParams{
		Context: ctx,
	}
}

// NewUpdateEmbeddedIdentitySourceParamsWithHTTPClient creates a new UpdateEmbeddedIdentitySourceParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateEmbeddedIdentitySourceParamsWithHTTPClient(client *http.Client) *UpdateEmbeddedIdentitySourceParams {
	return &UpdateEmbeddedIdentitySourceParams{
		HTTPClient: client,
	}
}

/*
UpdateEmbeddedIdentitySourceParams contains all the parameters to send to the API endpoint

	for the update embedded identity source operation.

	Typically these are written to a http.Request.
*/
type UpdateEmbeddedIdentitySourceParams struct {

	/* DomainName.

	   Domain Name associated with the identity source
	*/
	DomainName string

	/* ID.

	   ID of Identity Provider
	*/
	ID string

	/* IdentitySourceSpec.

	   Identity Source Spec
	*/
	IdentitySourceSpec *models.IdentitySourceSpec

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update embedded identity source params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateEmbeddedIdentitySourceParams) WithDefaults() *UpdateEmbeddedIdentitySourceParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update embedded identity source params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateEmbeddedIdentitySourceParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update embedded identity source params
func (o *UpdateEmbeddedIdentitySourceParams) WithTimeout(timeout time.Duration) *UpdateEmbeddedIdentitySourceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update embedded identity source params
func (o *UpdateEmbeddedIdentitySourceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update embedded identity source params
func (o *UpdateEmbeddedIdentitySourceParams) WithContext(ctx context.Context) *UpdateEmbeddedIdentitySourceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update embedded identity source params
func (o *UpdateEmbeddedIdentitySourceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update embedded identity source params
func (o *UpdateEmbeddedIdentitySourceParams) WithHTTPClient(client *http.Client) *UpdateEmbeddedIdentitySourceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update embedded identity source params
func (o *UpdateEmbeddedIdentitySourceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDomainName adds the domainName to the update embedded identity source params
func (o *UpdateEmbeddedIdentitySourceParams) WithDomainName(domainName string) *UpdateEmbeddedIdentitySourceParams {
	o.SetDomainName(domainName)
	return o
}

// SetDomainName adds the domainName to the update embedded identity source params
func (o *UpdateEmbeddedIdentitySourceParams) SetDomainName(domainName string) {
	o.DomainName = domainName
}

// WithID adds the id to the update embedded identity source params
func (o *UpdateEmbeddedIdentitySourceParams) WithID(id string) *UpdateEmbeddedIdentitySourceParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the update embedded identity source params
func (o *UpdateEmbeddedIdentitySourceParams) SetID(id string) {
	o.ID = id
}

// WithIdentitySourceSpec adds the identitySourceSpec to the update embedded identity source params
func (o *UpdateEmbeddedIdentitySourceParams) WithIdentitySourceSpec(identitySourceSpec *models.IdentitySourceSpec) *UpdateEmbeddedIdentitySourceParams {
	o.SetIdentitySourceSpec(identitySourceSpec)
	return o
}

// SetIdentitySourceSpec adds the identitySourceSpec to the update embedded identity source params
func (o *UpdateEmbeddedIdentitySourceParams) SetIdentitySourceSpec(identitySourceSpec *models.IdentitySourceSpec) {
	o.IdentitySourceSpec = identitySourceSpec
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateEmbeddedIdentitySourceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param domainName
	if err := r.SetPathParam("domainName", o.DomainName); err != nil {
		return err
	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}
	if o.IdentitySourceSpec != nil {
		if err := r.SetBodyParam(o.IdentitySourceSpec); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
