// Code generated by go-swagger; DO NOT EDIT.

package version_aliases_for_bundle_component_type

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

// NewDeleteAliasVersionsBySoftwareTypeAndBaseVersionParams creates a new DeleteAliasVersionsBySoftwareTypeAndBaseVersionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteAliasVersionsBySoftwareTypeAndBaseVersionParams() *DeleteAliasVersionsBySoftwareTypeAndBaseVersionParams {
	return &DeleteAliasVersionsBySoftwareTypeAndBaseVersionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteAliasVersionsBySoftwareTypeAndBaseVersionParamsWithTimeout creates a new DeleteAliasVersionsBySoftwareTypeAndBaseVersionParams object
// with the ability to set a timeout on a request.
func NewDeleteAliasVersionsBySoftwareTypeAndBaseVersionParamsWithTimeout(timeout time.Duration) *DeleteAliasVersionsBySoftwareTypeAndBaseVersionParams {
	return &DeleteAliasVersionsBySoftwareTypeAndBaseVersionParams{
		timeout: timeout,
	}
}

// NewDeleteAliasVersionsBySoftwareTypeAndBaseVersionParamsWithContext creates a new DeleteAliasVersionsBySoftwareTypeAndBaseVersionParams object
// with the ability to set a context for a request.
func NewDeleteAliasVersionsBySoftwareTypeAndBaseVersionParamsWithContext(ctx context.Context) *DeleteAliasVersionsBySoftwareTypeAndBaseVersionParams {
	return &DeleteAliasVersionsBySoftwareTypeAndBaseVersionParams{
		Context: ctx,
	}
}

// NewDeleteAliasVersionsBySoftwareTypeAndBaseVersionParamsWithHTTPClient creates a new DeleteAliasVersionsBySoftwareTypeAndBaseVersionParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteAliasVersionsBySoftwareTypeAndBaseVersionParamsWithHTTPClient(client *http.Client) *DeleteAliasVersionsBySoftwareTypeAndBaseVersionParams {
	return &DeleteAliasVersionsBySoftwareTypeAndBaseVersionParams{
		HTTPClient: client,
	}
}

/*
DeleteAliasVersionsBySoftwareTypeAndBaseVersionParams contains all the parameters to send to the API endpoint

	for the delete alias versions by software type and base version operation.

	Typically these are written to a http.Request.
*/
type DeleteAliasVersionsBySoftwareTypeAndBaseVersionParams struct {

	/* AliasVersions.

	   List of alias versions
	*/
	AliasVersions []string

	/* BundleComponentType.

	   Bundle Component Type
	*/
	BundleComponentType string

	/* Version.

	   Version
	*/
	Version string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete alias versions by software type and base version params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteAliasVersionsBySoftwareTypeAndBaseVersionParams) WithDefaults() *DeleteAliasVersionsBySoftwareTypeAndBaseVersionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete alias versions by software type and base version params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteAliasVersionsBySoftwareTypeAndBaseVersionParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete alias versions by software type and base version params
func (o *DeleteAliasVersionsBySoftwareTypeAndBaseVersionParams) WithTimeout(timeout time.Duration) *DeleteAliasVersionsBySoftwareTypeAndBaseVersionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete alias versions by software type and base version params
func (o *DeleteAliasVersionsBySoftwareTypeAndBaseVersionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete alias versions by software type and base version params
func (o *DeleteAliasVersionsBySoftwareTypeAndBaseVersionParams) WithContext(ctx context.Context) *DeleteAliasVersionsBySoftwareTypeAndBaseVersionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete alias versions by software type and base version params
func (o *DeleteAliasVersionsBySoftwareTypeAndBaseVersionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete alias versions by software type and base version params
func (o *DeleteAliasVersionsBySoftwareTypeAndBaseVersionParams) WithHTTPClient(client *http.Client) *DeleteAliasVersionsBySoftwareTypeAndBaseVersionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete alias versions by software type and base version params
func (o *DeleteAliasVersionsBySoftwareTypeAndBaseVersionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAliasVersions adds the aliasVersions to the delete alias versions by software type and base version params
func (o *DeleteAliasVersionsBySoftwareTypeAndBaseVersionParams) WithAliasVersions(aliasVersions []string) *DeleteAliasVersionsBySoftwareTypeAndBaseVersionParams {
	o.SetAliasVersions(aliasVersions)
	return o
}

// SetAliasVersions adds the aliasVersions to the delete alias versions by software type and base version params
func (o *DeleteAliasVersionsBySoftwareTypeAndBaseVersionParams) SetAliasVersions(aliasVersions []string) {
	o.AliasVersions = aliasVersions
}

// WithBundleComponentType adds the bundleComponentType to the delete alias versions by software type and base version params
func (o *DeleteAliasVersionsBySoftwareTypeAndBaseVersionParams) WithBundleComponentType(bundleComponentType string) *DeleteAliasVersionsBySoftwareTypeAndBaseVersionParams {
	o.SetBundleComponentType(bundleComponentType)
	return o
}

// SetBundleComponentType adds the bundleComponentType to the delete alias versions by software type and base version params
func (o *DeleteAliasVersionsBySoftwareTypeAndBaseVersionParams) SetBundleComponentType(bundleComponentType string) {
	o.BundleComponentType = bundleComponentType
}

// WithVersion adds the version to the delete alias versions by software type and base version params
func (o *DeleteAliasVersionsBySoftwareTypeAndBaseVersionParams) WithVersion(version string) *DeleteAliasVersionsBySoftwareTypeAndBaseVersionParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the delete alias versions by software type and base version params
func (o *DeleteAliasVersionsBySoftwareTypeAndBaseVersionParams) SetVersion(version string) {
	o.Version = version
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteAliasVersionsBySoftwareTypeAndBaseVersionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.AliasVersions != nil {
		if err := r.SetBodyParam(o.AliasVersions); err != nil {
			return err
		}
	}

	// path param bundleComponentType
	if err := r.SetPathParam("bundleComponentType", o.BundleComponentType); err != nil {
		return err
	}

	// path param version
	if err := r.SetPathParam("version", o.Version); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
