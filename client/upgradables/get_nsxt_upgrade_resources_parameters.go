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

// NewGetNsxtUpgradeResourcesParams creates a new GetNsxtUpgradeResourcesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetNsxtUpgradeResourcesParams() *GetNsxtUpgradeResourcesParams {
	return &GetNsxtUpgradeResourcesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetNsxtUpgradeResourcesParamsWithTimeout creates a new GetNsxtUpgradeResourcesParams object
// with the ability to set a timeout on a request.
func NewGetNsxtUpgradeResourcesParamsWithTimeout(timeout time.Duration) *GetNsxtUpgradeResourcesParams {
	return &GetNsxtUpgradeResourcesParams{
		timeout: timeout,
	}
}

// NewGetNsxtUpgradeResourcesParamsWithContext creates a new GetNsxtUpgradeResourcesParams object
// with the ability to set a context for a request.
func NewGetNsxtUpgradeResourcesParamsWithContext(ctx context.Context) *GetNsxtUpgradeResourcesParams {
	return &GetNsxtUpgradeResourcesParams{
		Context: ctx,
	}
}

// NewGetNsxtUpgradeResourcesParamsWithHTTPClient creates a new GetNsxtUpgradeResourcesParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetNsxtUpgradeResourcesParamsWithHTTPClient(client *http.Client) *GetNsxtUpgradeResourcesParams {
	return &GetNsxtUpgradeResourcesParams{
		HTTPClient: client,
	}
}

/*
GetNsxtUpgradeResourcesParams contains all the parameters to send to the API endpoint

	for the get nsxt upgrade resources operation.

	Typically these are written to a http.Request.
*/
type GetNsxtUpgradeResourcesParams struct {

	/* BundleID.

	   bundle Id of the upgrade bundle applicable on the domain
	*/
	BundleID *string

	/* DomainID.

	   Domain ID
	*/
	DomainID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get nsxt upgrade resources params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetNsxtUpgradeResourcesParams) WithDefaults() *GetNsxtUpgradeResourcesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get nsxt upgrade resources params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetNsxtUpgradeResourcesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get nsxt upgrade resources params
func (o *GetNsxtUpgradeResourcesParams) WithTimeout(timeout time.Duration) *GetNsxtUpgradeResourcesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get nsxt upgrade resources params
func (o *GetNsxtUpgradeResourcesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get nsxt upgrade resources params
func (o *GetNsxtUpgradeResourcesParams) WithContext(ctx context.Context) *GetNsxtUpgradeResourcesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get nsxt upgrade resources params
func (o *GetNsxtUpgradeResourcesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get nsxt upgrade resources params
func (o *GetNsxtUpgradeResourcesParams) WithHTTPClient(client *http.Client) *GetNsxtUpgradeResourcesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get nsxt upgrade resources params
func (o *GetNsxtUpgradeResourcesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBundleID adds the bundleID to the get nsxt upgrade resources params
func (o *GetNsxtUpgradeResourcesParams) WithBundleID(bundleID *string) *GetNsxtUpgradeResourcesParams {
	o.SetBundleID(bundleID)
	return o
}

// SetBundleID adds the bundleId to the get nsxt upgrade resources params
func (o *GetNsxtUpgradeResourcesParams) SetBundleID(bundleID *string) {
	o.BundleID = bundleID
}

// WithDomainID adds the domainID to the get nsxt upgrade resources params
func (o *GetNsxtUpgradeResourcesParams) WithDomainID(domainID string) *GetNsxtUpgradeResourcesParams {
	o.SetDomainID(domainID)
	return o
}

// SetDomainID adds the domainId to the get nsxt upgrade resources params
func (o *GetNsxtUpgradeResourcesParams) SetDomainID(domainID string) {
	o.DomainID = domainID
}

// WriteToRequest writes these params to a swagger request
func (o *GetNsxtUpgradeResourcesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.BundleID != nil {

		// query param bundleId
		var qrBundleID string

		if o.BundleID != nil {
			qrBundleID = *o.BundleID
		}
		qBundleID := qrBundleID
		if qBundleID != "" {

			if err := r.SetQueryParam("bundleId", qBundleID); err != nil {
				return err
			}
		}
	}

	// path param domainId
	if err := r.SetPathParam("domainId", o.DomainID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
