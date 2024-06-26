// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2023 VMware, Inc.
// SPDX-License-Identifier: BSD-2-Clause

package suite_lifecycle

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

	"github.com/vmware/vcf-sdk-go/models"
)

// NewUpdateVRSLCMVersionByIDInInventoryParams creates a new UpdateVRSLCMVersionByIDInInventoryParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateVRSLCMVersionByIDInInventoryParams() *UpdateVRSLCMVersionByIDInInventoryParams {
	return &UpdateVRSLCMVersionByIDInInventoryParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateVRSLCMVersionByIDInInventoryParamsWithTimeout creates a new UpdateVRSLCMVersionByIDInInventoryParams object
// with the ability to set a timeout on a request.
func NewUpdateVRSLCMVersionByIDInInventoryParamsWithTimeout(timeout time.Duration) *UpdateVRSLCMVersionByIDInInventoryParams {
	return &UpdateVRSLCMVersionByIDInInventoryParams{
		timeout: timeout,
	}
}

// NewUpdateVRSLCMVersionByIDInInventoryParamsWithContext creates a new UpdateVRSLCMVersionByIDInInventoryParams object
// with the ability to set a context for a request.
func NewUpdateVRSLCMVersionByIDInInventoryParamsWithContext(ctx context.Context) *UpdateVRSLCMVersionByIDInInventoryParams {
	return &UpdateVRSLCMVersionByIDInInventoryParams{
		Context: ctx,
	}
}

// NewUpdateVRSLCMVersionByIDInInventoryParamsWithHTTPClient creates a new UpdateVRSLCMVersionByIDInInventoryParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateVRSLCMVersionByIDInInventoryParamsWithHTTPClient(client *http.Client) *UpdateVRSLCMVersionByIDInInventoryParams {
	return &UpdateVRSLCMVersionByIDInInventoryParams{
		HTTPClient: client,
	}
}

/*
UpdateVRSLCMVersionByIDInInventoryParams contains all the parameters to send to the API endpoint

	for the update Vrslcm version by Id in inventory operation.

	Typically these are written to a http.Request.
*/
type UpdateVRSLCMVersionByIDInInventoryParams struct {

	/* ID.

	   the ID of VMware Aria Suite Lifecycle instance
	*/
	ID string

	/* VRSLCMDto.

	   vrslcmDto
	*/
	VRSLCMDto *models.VRSLCM

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update Vrslcm version by Id in inventory params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateVRSLCMVersionByIDInInventoryParams) WithDefaults() *UpdateVRSLCMVersionByIDInInventoryParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update Vrslcm version by Id in inventory params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateVRSLCMVersionByIDInInventoryParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update Vrslcm version by Id in inventory params
func (o *UpdateVRSLCMVersionByIDInInventoryParams) WithTimeout(timeout time.Duration) *UpdateVRSLCMVersionByIDInInventoryParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update Vrslcm version by Id in inventory params
func (o *UpdateVRSLCMVersionByIDInInventoryParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update Vrslcm version by Id in inventory params
func (o *UpdateVRSLCMVersionByIDInInventoryParams) WithContext(ctx context.Context) *UpdateVRSLCMVersionByIDInInventoryParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update Vrslcm version by Id in inventory params
func (o *UpdateVRSLCMVersionByIDInInventoryParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update Vrslcm version by Id in inventory params
func (o *UpdateVRSLCMVersionByIDInInventoryParams) WithHTTPClient(client *http.Client) *UpdateVRSLCMVersionByIDInInventoryParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update Vrslcm version by Id in inventory params
func (o *UpdateVRSLCMVersionByIDInInventoryParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the update Vrslcm version by Id in inventory params
func (o *UpdateVRSLCMVersionByIDInInventoryParams) WithID(id string) *UpdateVRSLCMVersionByIDInInventoryParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the update Vrslcm version by Id in inventory params
func (o *UpdateVRSLCMVersionByIDInInventoryParams) SetID(id string) {
	o.ID = id
}

// WithVRSLCMDto adds the vRSLCMDto to the update Vrslcm version by Id in inventory params
func (o *UpdateVRSLCMVersionByIDInInventoryParams) WithVRSLCMDto(vRSLCMDto *models.VRSLCM) *UpdateVRSLCMVersionByIDInInventoryParams {
	o.SetVRSLCMDto(vRSLCMDto)
	return o
}

// SetVRSLCMDto adds the vrslcmDto to the update Vrslcm version by Id in inventory params
func (o *UpdateVRSLCMVersionByIDInInventoryParams) SetVRSLCMDto(vRSLCMDto *models.VRSLCM) {
	o.VRSLCMDto = vRSLCMDto
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateVRSLCMVersionByIDInInventoryParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}
	if o.VRSLCMDto != nil {
		if err := r.SetBodyParam(o.VRSLCMDto); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
