// Code generated by go-swagger; DO NOT EDIT.

// Copyright © 2023 VMware, Inc. All Rights Reserved.
// Code licenced under: BSD-2

package bundles

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

// NewGETBundlesForSkipUpgradeUsingGETParams creates a new GETBundlesForSkipUpgradeUsingGETParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGETBundlesForSkipUpgradeUsingGETParams() *GETBundlesForSkipUpgradeUsingGETParams {
	return &GETBundlesForSkipUpgradeUsingGETParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGETBundlesForSkipUpgradeUsingGETParamsWithTimeout creates a new GETBundlesForSkipUpgradeUsingGETParams object
// with the ability to set a timeout on a request.
func NewGETBundlesForSkipUpgradeUsingGETParamsWithTimeout(timeout time.Duration) *GETBundlesForSkipUpgradeUsingGETParams {
	return &GETBundlesForSkipUpgradeUsingGETParams{
		timeout: timeout,
	}
}

// NewGETBundlesForSkipUpgradeUsingGETParamsWithContext creates a new GETBundlesForSkipUpgradeUsingGETParams object
// with the ability to set a context for a request.
func NewGETBundlesForSkipUpgradeUsingGETParamsWithContext(ctx context.Context) *GETBundlesForSkipUpgradeUsingGETParams {
	return &GETBundlesForSkipUpgradeUsingGETParams{
		Context: ctx,
	}
}

// NewGETBundlesForSkipUpgradeUsingGETParamsWithHTTPClient creates a new GETBundlesForSkipUpgradeUsingGETParams object
// with the ability to set a custom HTTPClient for a request.
func NewGETBundlesForSkipUpgradeUsingGETParamsWithHTTPClient(client *http.Client) *GETBundlesForSkipUpgradeUsingGETParams {
	return &GETBundlesForSkipUpgradeUsingGETParams{
		HTTPClient: client,
	}
}

/*
GETBundlesForSkipUpgradeUsingGETParams contains all the parameters to send to the API endpoint

	for the get bundles for skip upgrade using GET operation.

	Typically these are written to a http.Request.
*/
type GETBundlesForSkipUpgradeUsingGETParams struct {

	/* ID.

	   Domain ID
	*/
	ID string

	/* TargetVersion.

	   [Deprecated] Target domain VCF version
	*/
	TargetVersion *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get bundles for skip upgrade using GET params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GETBundlesForSkipUpgradeUsingGETParams) WithDefaults() *GETBundlesForSkipUpgradeUsingGETParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get bundles for skip upgrade using GET params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GETBundlesForSkipUpgradeUsingGETParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get bundles for skip upgrade using GET params
func (o *GETBundlesForSkipUpgradeUsingGETParams) WithTimeout(timeout time.Duration) *GETBundlesForSkipUpgradeUsingGETParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get bundles for skip upgrade using GET params
func (o *GETBundlesForSkipUpgradeUsingGETParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get bundles for skip upgrade using GET params
func (o *GETBundlesForSkipUpgradeUsingGETParams) WithContext(ctx context.Context) *GETBundlesForSkipUpgradeUsingGETParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get bundles for skip upgrade using GET params
func (o *GETBundlesForSkipUpgradeUsingGETParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get bundles for skip upgrade using GET params
func (o *GETBundlesForSkipUpgradeUsingGETParams) WithHTTPClient(client *http.Client) *GETBundlesForSkipUpgradeUsingGETParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get bundles for skip upgrade using GET params
func (o *GETBundlesForSkipUpgradeUsingGETParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get bundles for skip upgrade using GET params
func (o *GETBundlesForSkipUpgradeUsingGETParams) WithID(id string) *GETBundlesForSkipUpgradeUsingGETParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get bundles for skip upgrade using GET params
func (o *GETBundlesForSkipUpgradeUsingGETParams) SetID(id string) {
	o.ID = id
}

// WithTargetVersion adds the targetVersion to the get bundles for skip upgrade using GET params
func (o *GETBundlesForSkipUpgradeUsingGETParams) WithTargetVersion(targetVersion *string) *GETBundlesForSkipUpgradeUsingGETParams {
	o.SetTargetVersion(targetVersion)
	return o
}

// SetTargetVersion adds the targetVersion to the get bundles for skip upgrade using GET params
func (o *GETBundlesForSkipUpgradeUsingGETParams) SetTargetVersion(targetVersion *string) {
	o.TargetVersion = targetVersion
}

// WriteToRequest writes these params to a swagger request
func (o *GETBundlesForSkipUpgradeUsingGETParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if o.TargetVersion != nil {

		// query param targetVersion
		var qrTargetVersion string

		if o.TargetVersion != nil {
			qrTargetVersion = *o.TargetVersion
		}
		qTargetVersion := qrTargetVersion
		if qTargetVersion != "" {

			if err := r.SetQueryParam("targetVersion", qTargetVersion); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
