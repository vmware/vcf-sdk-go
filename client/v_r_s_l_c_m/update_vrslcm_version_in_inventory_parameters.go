// Code generated by go-swagger; DO NOT EDIT.

package v_r_s_l_c_m

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

// NewUpdateVrslcmVersionInInventoryParams creates a new UpdateVrslcmVersionInInventoryParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateVrslcmVersionInInventoryParams() *UpdateVrslcmVersionInInventoryParams {
	return &UpdateVrslcmVersionInInventoryParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateVrslcmVersionInInventoryParamsWithTimeout creates a new UpdateVrslcmVersionInInventoryParams object
// with the ability to set a timeout on a request.
func NewUpdateVrslcmVersionInInventoryParamsWithTimeout(timeout time.Duration) *UpdateVrslcmVersionInInventoryParams {
	return &UpdateVrslcmVersionInInventoryParams{
		timeout: timeout,
	}
}

// NewUpdateVrslcmVersionInInventoryParamsWithContext creates a new UpdateVrslcmVersionInInventoryParams object
// with the ability to set a context for a request.
func NewUpdateVrslcmVersionInInventoryParamsWithContext(ctx context.Context) *UpdateVrslcmVersionInInventoryParams {
	return &UpdateVrslcmVersionInInventoryParams{
		Context: ctx,
	}
}

// NewUpdateVrslcmVersionInInventoryParamsWithHTTPClient creates a new UpdateVrslcmVersionInInventoryParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateVrslcmVersionInInventoryParamsWithHTTPClient(client *http.Client) *UpdateVrslcmVersionInInventoryParams {
	return &UpdateVrslcmVersionInInventoryParams{
		HTTPClient: client,
	}
}

/*
UpdateVrslcmVersionInInventoryParams contains all the parameters to send to the API endpoint

	for the update vrslcm version in inventory operation.

	Typically these are written to a http.Request.
*/
type UpdateVrslcmVersionInInventoryParams struct {

	/* VrslcmDto.

	   vrslcmDto
	*/
	VrslcmDto *models.Vrslcm

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update vrslcm version in inventory params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateVrslcmVersionInInventoryParams) WithDefaults() *UpdateVrslcmVersionInInventoryParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update vrslcm version in inventory params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateVrslcmVersionInInventoryParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update vrslcm version in inventory params
func (o *UpdateVrslcmVersionInInventoryParams) WithTimeout(timeout time.Duration) *UpdateVrslcmVersionInInventoryParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update vrslcm version in inventory params
func (o *UpdateVrslcmVersionInInventoryParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update vrslcm version in inventory params
func (o *UpdateVrslcmVersionInInventoryParams) WithContext(ctx context.Context) *UpdateVrslcmVersionInInventoryParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update vrslcm version in inventory params
func (o *UpdateVrslcmVersionInInventoryParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update vrslcm version in inventory params
func (o *UpdateVrslcmVersionInInventoryParams) WithHTTPClient(client *http.Client) *UpdateVrslcmVersionInInventoryParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update vrslcm version in inventory params
func (o *UpdateVrslcmVersionInInventoryParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithVrslcmDto adds the vrslcmDto to the update vrslcm version in inventory params
func (o *UpdateVrslcmVersionInInventoryParams) WithVrslcmDto(vrslcmDto *models.Vrslcm) *UpdateVrslcmVersionInInventoryParams {
	o.SetVrslcmDto(vrslcmDto)
	return o
}

// SetVrslcmDto adds the vrslcmDto to the update vrslcm version in inventory params
func (o *UpdateVrslcmVersionInInventoryParams) SetVrslcmDto(vrslcmDto *models.Vrslcm) {
	o.VrslcmDto = vrslcmDto
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateVrslcmVersionInInventoryParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.VrslcmDto != nil {
		if err := r.SetBodyParam(o.VrslcmDto); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}