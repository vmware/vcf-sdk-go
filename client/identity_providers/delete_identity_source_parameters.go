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
)

// NewDeleteIdentitySourceParams creates a new DeleteIdentitySourceParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteIdentitySourceParams() *DeleteIdentitySourceParams {
	return &DeleteIdentitySourceParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteIdentitySourceParamsWithTimeout creates a new DeleteIdentitySourceParams object
// with the ability to set a timeout on a request.
func NewDeleteIdentitySourceParamsWithTimeout(timeout time.Duration) *DeleteIdentitySourceParams {
	return &DeleteIdentitySourceParams{
		timeout: timeout,
	}
}

// NewDeleteIdentitySourceParamsWithContext creates a new DeleteIdentitySourceParams object
// with the ability to set a context for a request.
func NewDeleteIdentitySourceParamsWithContext(ctx context.Context) *DeleteIdentitySourceParams {
	return &DeleteIdentitySourceParams{
		Context: ctx,
	}
}

// NewDeleteIdentitySourceParamsWithHTTPClient creates a new DeleteIdentitySourceParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteIdentitySourceParamsWithHTTPClient(client *http.Client) *DeleteIdentitySourceParams {
	return &DeleteIdentitySourceParams{
		HTTPClient: client,
	}
}

/*
DeleteIdentitySourceParams contains all the parameters to send to the API endpoint

	for the delete identity source operation.

	Typically these are written to a http.Request.
*/
type DeleteIdentitySourceParams struct {

	/* DomainName.

	   Domain Name associated with the identity source
	*/
	DomainName string

	/* ID.

	   ID of Identity Provider
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete identity source params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteIdentitySourceParams) WithDefaults() *DeleteIdentitySourceParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete identity source params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteIdentitySourceParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete identity source params
func (o *DeleteIdentitySourceParams) WithTimeout(timeout time.Duration) *DeleteIdentitySourceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete identity source params
func (o *DeleteIdentitySourceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete identity source params
func (o *DeleteIdentitySourceParams) WithContext(ctx context.Context) *DeleteIdentitySourceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete identity source params
func (o *DeleteIdentitySourceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete identity source params
func (o *DeleteIdentitySourceParams) WithHTTPClient(client *http.Client) *DeleteIdentitySourceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete identity source params
func (o *DeleteIdentitySourceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDomainName adds the domainName to the delete identity source params
func (o *DeleteIdentitySourceParams) WithDomainName(domainName string) *DeleteIdentitySourceParams {
	o.SetDomainName(domainName)
	return o
}

// SetDomainName adds the domainName to the delete identity source params
func (o *DeleteIdentitySourceParams) SetDomainName(domainName string) {
	o.DomainName = domainName
}

// WithID adds the id to the delete identity source params
func (o *DeleteIdentitySourceParams) WithID(id string) *DeleteIdentitySourceParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete identity source params
func (o *DeleteIdentitySourceParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteIdentitySourceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}