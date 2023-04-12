// Code generated by go-swagger; DO NOT EDIT.

package upgradables

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

// NewGetUpgradablesByDomainParams creates a new GetUpgradablesByDomainParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetUpgradablesByDomainParams() *GetUpgradablesByDomainParams {
	return &GetUpgradablesByDomainParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetUpgradablesByDomainParamsWithTimeout creates a new GetUpgradablesByDomainParams object
// with the ability to set a timeout on a request.
func NewGetUpgradablesByDomainParamsWithTimeout(timeout time.Duration) *GetUpgradablesByDomainParams {
	return &GetUpgradablesByDomainParams{
		timeout: timeout,
	}
}

// NewGetUpgradablesByDomainParamsWithContext creates a new GetUpgradablesByDomainParams object
// with the ability to set a context for a request.
func NewGetUpgradablesByDomainParamsWithContext(ctx context.Context) *GetUpgradablesByDomainParams {
	return &GetUpgradablesByDomainParams{
		Context: ctx,
	}
}

// NewGetUpgradablesByDomainParamsWithHTTPClient creates a new GetUpgradablesByDomainParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetUpgradablesByDomainParamsWithHTTPClient(client *http.Client) *GetUpgradablesByDomainParams {
	return &GetUpgradablesByDomainParams{
		HTTPClient: client,
	}
}

/*
GetUpgradablesByDomainParams contains all the parameters to send to the API endpoint

	for the get upgradables by domain operation.

	Typically these are written to a http.Request.
*/
type GetUpgradablesByDomainParams struct {

	/* DomainID.

	   Domain ID
	*/
	DomainID string

	/* TargetVersion.

	   Target Version to get Upgradables for a given Target Release
	*/
	TargetVersion *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get upgradables by domain params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetUpgradablesByDomainParams) WithDefaults() *GetUpgradablesByDomainParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get upgradables by domain params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetUpgradablesByDomainParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get upgradables by domain params
func (o *GetUpgradablesByDomainParams) WithTimeout(timeout time.Duration) *GetUpgradablesByDomainParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get upgradables by domain params
func (o *GetUpgradablesByDomainParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get upgradables by domain params
func (o *GetUpgradablesByDomainParams) WithContext(ctx context.Context) *GetUpgradablesByDomainParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get upgradables by domain params
func (o *GetUpgradablesByDomainParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get upgradables by domain params
func (o *GetUpgradablesByDomainParams) WithHTTPClient(client *http.Client) *GetUpgradablesByDomainParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get upgradables by domain params
func (o *GetUpgradablesByDomainParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDomainID adds the domainID to the get upgradables by domain params
func (o *GetUpgradablesByDomainParams) WithDomainID(domainID string) *GetUpgradablesByDomainParams {
	o.SetDomainID(domainID)
	return o
}

// SetDomainID adds the domainId to the get upgradables by domain params
func (o *GetUpgradablesByDomainParams) SetDomainID(domainID string) {
	o.DomainID = domainID
}

// WithTargetVersion adds the targetVersion to the get upgradables by domain params
func (o *GetUpgradablesByDomainParams) WithTargetVersion(targetVersion *string) *GetUpgradablesByDomainParams {
	o.SetTargetVersion(targetVersion)
	return o
}

// SetTargetVersion adds the targetVersion to the get upgradables by domain params
func (o *GetUpgradablesByDomainParams) SetTargetVersion(targetVersion *string) {
	o.TargetVersion = targetVersion
}

// WriteToRequest writes these params to a swagger request
func (o *GetUpgradablesByDomainParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param domainId
	if err := r.SetPathParam("domainId", o.DomainID); err != nil {
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
