// Code generated by go-swagger; DO NOT EDIT.

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
	"github.com/go-openapi/swag"
)

// NewGetBundlesParams creates a new GetBundlesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetBundlesParams() *GetBundlesParams {
	return &GetBundlesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetBundlesParamsWithTimeout creates a new GetBundlesParams object
// with the ability to set a timeout on a request.
func NewGetBundlesParamsWithTimeout(timeout time.Duration) *GetBundlesParams {
	return &GetBundlesParams{
		timeout: timeout,
	}
}

// NewGetBundlesParamsWithContext creates a new GetBundlesParams object
// with the ability to set a context for a request.
func NewGetBundlesParamsWithContext(ctx context.Context) *GetBundlesParams {
	return &GetBundlesParams{
		Context: ctx,
	}
}

// NewGetBundlesParamsWithHTTPClient creates a new GetBundlesParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetBundlesParamsWithHTTPClient(client *http.Client) *GetBundlesParams {
	return &GetBundlesParams{
		HTTPClient: client,
	}
}

/*
GetBundlesParams contains all the parameters to send to the API endpoint

	for the get bundles operation.

	Typically these are written to a http.Request.
*/
type GetBundlesParams struct {

	/* BundleType.

	   The type of the bundle
	*/
	BundleType *string

	/* IsCompliant.

	   Is compliant with the current VCF version
	*/
	IsCompliant *bool

	/* ProductType.

	   The type of the product
	*/
	ProductType *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get bundles params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetBundlesParams) WithDefaults() *GetBundlesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get bundles params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetBundlesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get bundles params
func (o *GetBundlesParams) WithTimeout(timeout time.Duration) *GetBundlesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get bundles params
func (o *GetBundlesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get bundles params
func (o *GetBundlesParams) WithContext(ctx context.Context) *GetBundlesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get bundles params
func (o *GetBundlesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get bundles params
func (o *GetBundlesParams) WithHTTPClient(client *http.Client) *GetBundlesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get bundles params
func (o *GetBundlesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBundleType adds the bundleType to the get bundles params
func (o *GetBundlesParams) WithBundleType(bundleType *string) *GetBundlesParams {
	o.SetBundleType(bundleType)
	return o
}

// SetBundleType adds the bundleType to the get bundles params
func (o *GetBundlesParams) SetBundleType(bundleType *string) {
	o.BundleType = bundleType
}

// WithIsCompliant adds the isCompliant to the get bundles params
func (o *GetBundlesParams) WithIsCompliant(isCompliant *bool) *GetBundlesParams {
	o.SetIsCompliant(isCompliant)
	return o
}

// SetIsCompliant adds the isCompliant to the get bundles params
func (o *GetBundlesParams) SetIsCompliant(isCompliant *bool) {
	o.IsCompliant = isCompliant
}

// WithProductType adds the productType to the get bundles params
func (o *GetBundlesParams) WithProductType(productType *string) *GetBundlesParams {
	o.SetProductType(productType)
	return o
}

// SetProductType adds the productType to the get bundles params
func (o *GetBundlesParams) SetProductType(productType *string) {
	o.ProductType = productType
}

// WriteToRequest writes these params to a swagger request
func (o *GetBundlesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.BundleType != nil {

		// query param bundleType
		var qrBundleType string

		if o.BundleType != nil {
			qrBundleType = *o.BundleType
		}
		qBundleType := qrBundleType
		if qBundleType != "" {

			if err := r.SetQueryParam("bundleType", qBundleType); err != nil {
				return err
			}
		}
	}

	if o.IsCompliant != nil {

		// query param isCompliant
		var qrIsCompliant bool

		if o.IsCompliant != nil {
			qrIsCompliant = *o.IsCompliant
		}
		qIsCompliant := swag.FormatBool(qrIsCompliant)
		if qIsCompliant != "" {

			if err := r.SetQueryParam("isCompliant", qIsCompliant); err != nil {
				return err
			}
		}
	}

	if o.ProductType != nil {

		// query param productType
		var qrProductType string

		if o.ProductType != nil {
			qrProductType = *o.ProductType
		}
		qProductType := qrProductType
		if qProductType != "" {

			if err := r.SetQueryParam("productType", qProductType); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
