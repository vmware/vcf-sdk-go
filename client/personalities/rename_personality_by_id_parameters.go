// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2023 VMware, Inc.
// SPDX-License-Identifier: BSD-2-Clause

package personalities

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

// NewRenamePersonalityByIDParams creates a new RenamePersonalityByIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewRenamePersonalityByIDParams() *RenamePersonalityByIDParams {
	return &RenamePersonalityByIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewRenamePersonalityByIDParamsWithTimeout creates a new RenamePersonalityByIDParams object
// with the ability to set a timeout on a request.
func NewRenamePersonalityByIDParamsWithTimeout(timeout time.Duration) *RenamePersonalityByIDParams {
	return &RenamePersonalityByIDParams{
		timeout: timeout,
	}
}

// NewRenamePersonalityByIDParamsWithContext creates a new RenamePersonalityByIDParams object
// with the ability to set a context for a request.
func NewRenamePersonalityByIDParamsWithContext(ctx context.Context) *RenamePersonalityByIDParams {
	return &RenamePersonalityByIDParams{
		Context: ctx,
	}
}

// NewRenamePersonalityByIDParamsWithHTTPClient creates a new RenamePersonalityByIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewRenamePersonalityByIDParamsWithHTTPClient(client *http.Client) *RenamePersonalityByIDParams {
	return &RenamePersonalityByIDParams{
		HTTPClient: client,
	}
}

/*
RenamePersonalityByIDParams contains all the parameters to send to the API endpoint

	for the rename personality by Id operation.

	Typically these are written to a http.Request.
*/
type RenamePersonalityByIDParams struct {

	/* Personality.

	   personality
	*/
	Personality *models.Personality

	/* PersonalityID.

	   The personality id
	*/
	PersonalityID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the rename personality by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RenamePersonalityByIDParams) WithDefaults() *RenamePersonalityByIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the rename personality by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RenamePersonalityByIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the rename personality by Id params
func (o *RenamePersonalityByIDParams) WithTimeout(timeout time.Duration) *RenamePersonalityByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the rename personality by Id params
func (o *RenamePersonalityByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the rename personality by Id params
func (o *RenamePersonalityByIDParams) WithContext(ctx context.Context) *RenamePersonalityByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the rename personality by Id params
func (o *RenamePersonalityByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the rename personality by Id params
func (o *RenamePersonalityByIDParams) WithHTTPClient(client *http.Client) *RenamePersonalityByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the rename personality by Id params
func (o *RenamePersonalityByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPersonality adds the personality to the rename personality by Id params
func (o *RenamePersonalityByIDParams) WithPersonality(personality *models.Personality) *RenamePersonalityByIDParams {
	o.SetPersonality(personality)
	return o
}

// SetPersonality adds the personality to the rename personality by Id params
func (o *RenamePersonalityByIDParams) SetPersonality(personality *models.Personality) {
	o.Personality = personality
}

// WithPersonalityID adds the personalityID to the rename personality by Id params
func (o *RenamePersonalityByIDParams) WithPersonalityID(personalityID string) *RenamePersonalityByIDParams {
	o.SetPersonalityID(personalityID)
	return o
}

// SetPersonalityID adds the personalityId to the rename personality by Id params
func (o *RenamePersonalityByIDParams) SetPersonalityID(personalityID string) {
	o.PersonalityID = personalityID
}

// WriteToRequest writes these params to a swagger request
func (o *RenamePersonalityByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Personality != nil {
		if err := r.SetBodyParam(o.Personality); err != nil {
			return err
		}
	}

	// path param personalityId
	if err := r.SetPathParam("personalityId", o.PersonalityID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}