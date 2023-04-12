// Code generated by go-swagger; DO NOT EDIT.

package n_s_x_t_clusters

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

// NewValidateIPPoolUsingPOSTParams creates a new ValidateIPPoolUsingPOSTParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewValidateIPPoolUsingPOSTParams() *ValidateIPPoolUsingPOSTParams {
	return &ValidateIPPoolUsingPOSTParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewValidateIPPoolUsingPOSTParamsWithTimeout creates a new ValidateIPPoolUsingPOSTParams object
// with the ability to set a timeout on a request.
func NewValidateIPPoolUsingPOSTParamsWithTimeout(timeout time.Duration) *ValidateIPPoolUsingPOSTParams {
	return &ValidateIPPoolUsingPOSTParams{
		timeout: timeout,
	}
}

// NewValidateIPPoolUsingPOSTParamsWithContext creates a new ValidateIPPoolUsingPOSTParams object
// with the ability to set a context for a request.
func NewValidateIPPoolUsingPOSTParamsWithContext(ctx context.Context) *ValidateIPPoolUsingPOSTParams {
	return &ValidateIPPoolUsingPOSTParams{
		Context: ctx,
	}
}

// NewValidateIPPoolUsingPOSTParamsWithHTTPClient creates a new ValidateIPPoolUsingPOSTParams object
// with the ability to set a custom HTTPClient for a request.
func NewValidateIPPoolUsingPOSTParamsWithHTTPClient(client *http.Client) *ValidateIPPoolUsingPOSTParams {
	return &ValidateIPPoolUsingPOSTParams{
		HTTPClient: client,
	}
}

/*
ValidateIPPoolUsingPOSTParams contains all the parameters to send to the API endpoint

	for the validate Ip pool using p o s t operation.

	Typically these are written to a http.Request.
*/
type ValidateIPPoolUsingPOSTParams struct {

	/* NsxtIPAddressPoolValidationSpec.

	   nsxtIpAddressPoolValidationSpec
	*/
	NsxtIPAddressPoolValidationSpec *models.NsxtIPAddressPoolValidationSpec

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the validate Ip pool using p o s t params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ValidateIPPoolUsingPOSTParams) WithDefaults() *ValidateIPPoolUsingPOSTParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the validate Ip pool using p o s t params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ValidateIPPoolUsingPOSTParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the validate Ip pool using p o s t params
func (o *ValidateIPPoolUsingPOSTParams) WithTimeout(timeout time.Duration) *ValidateIPPoolUsingPOSTParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the validate Ip pool using p o s t params
func (o *ValidateIPPoolUsingPOSTParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the validate Ip pool using p o s t params
func (o *ValidateIPPoolUsingPOSTParams) WithContext(ctx context.Context) *ValidateIPPoolUsingPOSTParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the validate Ip pool using p o s t params
func (o *ValidateIPPoolUsingPOSTParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the validate Ip pool using p o s t params
func (o *ValidateIPPoolUsingPOSTParams) WithHTTPClient(client *http.Client) *ValidateIPPoolUsingPOSTParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the validate Ip pool using p o s t params
func (o *ValidateIPPoolUsingPOSTParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNsxtIPAddressPoolValidationSpec adds the nsxtIPAddressPoolValidationSpec to the validate Ip pool using p o s t params
func (o *ValidateIPPoolUsingPOSTParams) WithNsxtIPAddressPoolValidationSpec(nsxtIPAddressPoolValidationSpec *models.NsxtIPAddressPoolValidationSpec) *ValidateIPPoolUsingPOSTParams {
	o.SetNsxtIPAddressPoolValidationSpec(nsxtIPAddressPoolValidationSpec)
	return o
}

// SetNsxtIPAddressPoolValidationSpec adds the nsxtIpAddressPoolValidationSpec to the validate Ip pool using p o s t params
func (o *ValidateIPPoolUsingPOSTParams) SetNsxtIPAddressPoolValidationSpec(nsxtIPAddressPoolValidationSpec *models.NsxtIPAddressPoolValidationSpec) {
	o.NsxtIPAddressPoolValidationSpec = nsxtIPAddressPoolValidationSpec
}

// WriteToRequest writes these params to a swagger request
func (o *ValidateIPPoolUsingPOSTParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.NsxtIPAddressPoolValidationSpec != nil {
		if err := r.SetBodyParam(o.NsxtIPAddressPoolValidationSpec); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
