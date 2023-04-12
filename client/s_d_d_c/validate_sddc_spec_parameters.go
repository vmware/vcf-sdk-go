// Code generated by go-swagger; DO NOT EDIT.

package s_d_d_c

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

	"vcf-sdk-go/models"
)

// NewValidateSddcSpecParams creates a new ValidateSddcSpecParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewValidateSddcSpecParams() *ValidateSddcSpecParams {
	return &ValidateSddcSpecParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewValidateSddcSpecParamsWithTimeout creates a new ValidateSddcSpecParams object
// with the ability to set a timeout on a request.
func NewValidateSddcSpecParamsWithTimeout(timeout time.Duration) *ValidateSddcSpecParams {
	return &ValidateSddcSpecParams{
		timeout: timeout,
	}
}

// NewValidateSddcSpecParamsWithContext creates a new ValidateSddcSpecParams object
// with the ability to set a context for a request.
func NewValidateSddcSpecParamsWithContext(ctx context.Context) *ValidateSddcSpecParams {
	return &ValidateSddcSpecParams{
		Context: ctx,
	}
}

// NewValidateSddcSpecParamsWithHTTPClient creates a new ValidateSddcSpecParams object
// with the ability to set a custom HTTPClient for a request.
func NewValidateSddcSpecParamsWithHTTPClient(client *http.Client) *ValidateSddcSpecParams {
	return &ValidateSddcSpecParams{
		HTTPClient: client,
	}
}

/*
ValidateSddcSpecParams contains all the parameters to send to the API endpoint

	for the validate sddc spec operation.

	Typically these are written to a http.Request.
*/
type ValidateSddcSpecParams struct {

	/* Name.

	     Validation name
	Deprecated: ESXI_VERSION_VALIDATION
	*/
	Name *string

	/* Redo.

	   redo
	*/
	Redo *bool

	/* SddcSpec.

	   SDDC specification
	*/
	SddcSpec *models.SddcSpec

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the validate sddc spec params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ValidateSddcSpecParams) WithDefaults() *ValidateSddcSpecParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the validate sddc spec params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ValidateSddcSpecParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the validate sddc spec params
func (o *ValidateSddcSpecParams) WithTimeout(timeout time.Duration) *ValidateSddcSpecParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the validate sddc spec params
func (o *ValidateSddcSpecParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the validate sddc spec params
func (o *ValidateSddcSpecParams) WithContext(ctx context.Context) *ValidateSddcSpecParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the validate sddc spec params
func (o *ValidateSddcSpecParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the validate sddc spec params
func (o *ValidateSddcSpecParams) WithHTTPClient(client *http.Client) *ValidateSddcSpecParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the validate sddc spec params
func (o *ValidateSddcSpecParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the validate sddc spec params
func (o *ValidateSddcSpecParams) WithName(name *string) *ValidateSddcSpecParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the validate sddc spec params
func (o *ValidateSddcSpecParams) SetName(name *string) {
	o.Name = name
}

// WithRedo adds the redo to the validate sddc spec params
func (o *ValidateSddcSpecParams) WithRedo(redo *bool) *ValidateSddcSpecParams {
	o.SetRedo(redo)
	return o
}

// SetRedo adds the redo to the validate sddc spec params
func (o *ValidateSddcSpecParams) SetRedo(redo *bool) {
	o.Redo = redo
}

// WithSddcSpec adds the sddcSpec to the validate sddc spec params
func (o *ValidateSddcSpecParams) WithSddcSpec(sddcSpec *models.SddcSpec) *ValidateSddcSpecParams {
	o.SetSddcSpec(sddcSpec)
	return o
}

// SetSddcSpec adds the sddcSpec to the validate sddc spec params
func (o *ValidateSddcSpecParams) SetSddcSpec(sddcSpec *models.SddcSpec) {
	o.SddcSpec = sddcSpec
}

// WriteToRequest writes these params to a swagger request
func (o *ValidateSddcSpecParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Name != nil {

		// query param name
		var qrName string

		if o.Name != nil {
			qrName = *o.Name
		}
		qName := qrName
		if qName != "" {

			if err := r.SetQueryParam("name", qName); err != nil {
				return err
			}
		}
	}

	if o.Redo != nil {

		// query param redo
		var qrRedo bool

		if o.Redo != nil {
			qrRedo = *o.Redo
		}
		qRedo := swag.FormatBool(qrRedo)
		if qRedo != "" {

			if err := r.SetQueryParam("redo", qRedo); err != nil {
				return err
			}
		}
	}
	if o.SddcSpec != nil {
		if err := r.SetBodyParam(o.SddcSpec); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
