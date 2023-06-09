// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2023 VMware, Inc.
// SPDX-License-Identifier: BSD-2-Clause

package releases

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
	"github.com/go-openapi/swag"
)

// NewGetReleasesParams creates a new GetReleasesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetReleasesParams() *GetReleasesParams {
	return &GetReleasesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetReleasesParamsWithTimeout creates a new GetReleasesParams object
// with the ability to set a timeout on a request.
func NewGetReleasesParamsWithTimeout(timeout time.Duration) *GetReleasesParams {
	return &GetReleasesParams{
		timeout: timeout,
	}
}

// NewGetReleasesParamsWithContext creates a new GetReleasesParams object
// with the ability to set a context for a request.
func NewGetReleasesParamsWithContext(ctx context.Context) *GetReleasesParams {
	return &GetReleasesParams{
		Context: ctx,
	}
}

// NewGetReleasesParamsWithHTTPClient creates a new GetReleasesParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetReleasesParamsWithHTTPClient(client *http.Client) *GetReleasesParams {
	return &GetReleasesParams{
		HTTPClient: client,
	}
}

/*
GetReleasesParams contains all the parameters to send to the API endpoint

	for the get releases operation.

	Typically these are written to a http.Request.
*/
type GetReleasesParams struct {

	/* ApplicableForVersion.

	   Release version to get applicable releases
	*/
	ApplicableForVersion *string

	/* ApplicableForVxRailVersion.

	   Release VxRail version to get applicable releases
	*/
	ApplicableForVxRailVersion *string

	/* DomainID.

	   Domain ID to get current release of the domain
	*/
	DomainID *string

	/* GetFutureReleases.

	   [Deprecated] Get all future releases for a given domain
	*/
	GetFutureReleases *bool

	/* VersionEq.

	   Release version to get its release
	*/
	VersionEq *string

	/* VersionGt.

	   Release version to get its future releases
	*/
	VersionGt *string

	/* VxRailVersionEq.

	   Release VxRail version to get its release
	*/
	VxRailVersionEq *string

	/* VxRailVersionGt.

	   Release vxrail version to get its future releases
	*/
	VxRailVersionGt *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get releases params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetReleasesParams) WithDefaults() *GetReleasesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get releases params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetReleasesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get releases params
func (o *GetReleasesParams) WithTimeout(timeout time.Duration) *GetReleasesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get releases params
func (o *GetReleasesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get releases params
func (o *GetReleasesParams) WithContext(ctx context.Context) *GetReleasesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get releases params
func (o *GetReleasesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get releases params
func (o *GetReleasesParams) WithHTTPClient(client *http.Client) *GetReleasesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get releases params
func (o *GetReleasesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithApplicableForVersion adds the applicableForVersion to the get releases params
func (o *GetReleasesParams) WithApplicableForVersion(applicableForVersion *string) *GetReleasesParams {
	o.SetApplicableForVersion(applicableForVersion)
	return o
}

// SetApplicableForVersion adds the applicableForVersion to the get releases params
func (o *GetReleasesParams) SetApplicableForVersion(applicableForVersion *string) {
	o.ApplicableForVersion = applicableForVersion
}

// WithApplicableForVxRailVersion adds the applicableForVxRailVersion to the get releases params
func (o *GetReleasesParams) WithApplicableForVxRailVersion(applicableForVxRailVersion *string) *GetReleasesParams {
	o.SetApplicableForVxRailVersion(applicableForVxRailVersion)
	return o
}

// SetApplicableForVxRailVersion adds the applicableForVxRailVersion to the get releases params
func (o *GetReleasesParams) SetApplicableForVxRailVersion(applicableForVxRailVersion *string) {
	o.ApplicableForVxRailVersion = applicableForVxRailVersion
}

// WithDomainID adds the domainID to the get releases params
func (o *GetReleasesParams) WithDomainID(domainID *string) *GetReleasesParams {
	o.SetDomainID(domainID)
	return o
}

// SetDomainID adds the domainId to the get releases params
func (o *GetReleasesParams) SetDomainID(domainID *string) {
	o.DomainID = domainID
}

// WithGetFutureReleases adds the getFutureReleases to the get releases params
func (o *GetReleasesParams) WithGetFutureReleases(getFutureReleases *bool) *GetReleasesParams {
	o.SetGetFutureReleases(getFutureReleases)
	return o
}

// SetGetFutureReleases adds the getFutureReleases to the get releases params
func (o *GetReleasesParams) SetGetFutureReleases(getFutureReleases *bool) {
	o.GetFutureReleases = getFutureReleases
}

// WithVersionEq adds the versionEq to the get releases params
func (o *GetReleasesParams) WithVersionEq(versionEq *string) *GetReleasesParams {
	o.SetVersionEq(versionEq)
	return o
}

// SetVersionEq adds the versionEq to the get releases params
func (o *GetReleasesParams) SetVersionEq(versionEq *string) {
	o.VersionEq = versionEq
}

// WithVersionGt adds the versionGt to the get releases params
func (o *GetReleasesParams) WithVersionGt(versionGt *string) *GetReleasesParams {
	o.SetVersionGt(versionGt)
	return o
}

// SetVersionGt adds the versionGt to the get releases params
func (o *GetReleasesParams) SetVersionGt(versionGt *string) {
	o.VersionGt = versionGt
}

// WithVxRailVersionEq adds the vxRailVersionEq to the get releases params
func (o *GetReleasesParams) WithVxRailVersionEq(vxRailVersionEq *string) *GetReleasesParams {
	o.SetVxRailVersionEq(vxRailVersionEq)
	return o
}

// SetVxRailVersionEq adds the vxRailVersionEq to the get releases params
func (o *GetReleasesParams) SetVxRailVersionEq(vxRailVersionEq *string) {
	o.VxRailVersionEq = vxRailVersionEq
}

// WithVxRailVersionGt adds the vxRailVersionGt to the get releases params
func (o *GetReleasesParams) WithVxRailVersionGt(vxRailVersionGt *string) *GetReleasesParams {
	o.SetVxRailVersionGt(vxRailVersionGt)
	return o
}

// SetVxRailVersionGt adds the vxRailVersionGt to the get releases params
func (o *GetReleasesParams) SetVxRailVersionGt(vxRailVersionGt *string) {
	o.VxRailVersionGt = vxRailVersionGt
}

// WriteToRequest writes these params to a swagger request
func (o *GetReleasesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ApplicableForVersion != nil {

		// query param applicableForVersion
		var qrApplicableForVersion string

		if o.ApplicableForVersion != nil {
			qrApplicableForVersion = *o.ApplicableForVersion
		}
		qApplicableForVersion := qrApplicableForVersion
		if qApplicableForVersion != "" {

			if err := r.SetQueryParam("applicableForVersion", qApplicableForVersion); err != nil {
				return err
			}
		}
	}

	if o.ApplicableForVxRailVersion != nil {

		// query param applicableForVxRailVersion
		var qrApplicableForVxRailVersion string

		if o.ApplicableForVxRailVersion != nil {
			qrApplicableForVxRailVersion = *o.ApplicableForVxRailVersion
		}
		qApplicableForVxRailVersion := qrApplicableForVxRailVersion
		if qApplicableForVxRailVersion != "" {

			if err := r.SetQueryParam("applicableForVxRailVersion", qApplicableForVxRailVersion); err != nil {
				return err
			}
		}
	}

	if o.DomainID != nil {

		// query param domainId
		var qrDomainID string

		if o.DomainID != nil {
			qrDomainID = *o.DomainID
		}
		qDomainID := qrDomainID
		if qDomainID != "" {

			if err := r.SetQueryParam("domainId", qDomainID); err != nil {
				return err
			}
		}
	}

	if o.GetFutureReleases != nil {

		// query param getFutureReleases
		var qrGetFutureReleases bool

		if o.GetFutureReleases != nil {
			qrGetFutureReleases = *o.GetFutureReleases
		}
		qGetFutureReleases := swag.FormatBool(qrGetFutureReleases)
		if qGetFutureReleases != "" {

			if err := r.SetQueryParam("getFutureReleases", qGetFutureReleases); err != nil {
				return err
			}
		}
	}

	if o.VersionEq != nil {

		// query param versionEq
		var qrVersionEq string

		if o.VersionEq != nil {
			qrVersionEq = *o.VersionEq
		}
		qVersionEq := qrVersionEq
		if qVersionEq != "" {

			if err := r.SetQueryParam("versionEq", qVersionEq); err != nil {
				return err
			}
		}
	}

	if o.VersionGt != nil {

		// query param versionGt
		var qrVersionGt string

		if o.VersionGt != nil {
			qrVersionGt = *o.VersionGt
		}
		qVersionGt := qrVersionGt
		if qVersionGt != "" {

			if err := r.SetQueryParam("versionGt", qVersionGt); err != nil {
				return err
			}
		}
	}

	if o.VxRailVersionEq != nil {

		// query param vxRailVersionEq
		var qrVxRailVersionEq string

		if o.VxRailVersionEq != nil {
			qrVxRailVersionEq = *o.VxRailVersionEq
		}
		qVxRailVersionEq := qrVxRailVersionEq
		if qVxRailVersionEq != "" {

			if err := r.SetQueryParam("vxRailVersionEq", qVxRailVersionEq); err != nil {
				return err
			}
		}
	}

	if o.VxRailVersionGt != nil {

		// query param vxRailVersionGt
		var qrVxRailVersionGt string

		if o.VxRailVersionGt != nil {
			qrVxRailVersionGt = *o.VxRailVersionGt
		}
		qVxRailVersionGt := qrVxRailVersionGt
		if qVxRailVersionGt != "" {

			if err := r.SetQueryParam("vxRailVersionGt", qVxRailVersionGt); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
