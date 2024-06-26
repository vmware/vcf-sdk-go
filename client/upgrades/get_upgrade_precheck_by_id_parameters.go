// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2023 VMware, Inc.
// SPDX-License-Identifier: BSD-2-Clause

package upgrades

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

// NewGetUpgradePrecheckByIDParams creates a new GetUpgradePrecheckByIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetUpgradePrecheckByIDParams() *GetUpgradePrecheckByIDParams {
	return &GetUpgradePrecheckByIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetUpgradePrecheckByIDParamsWithTimeout creates a new GetUpgradePrecheckByIDParams object
// with the ability to set a timeout on a request.
func NewGetUpgradePrecheckByIDParamsWithTimeout(timeout time.Duration) *GetUpgradePrecheckByIDParams {
	return &GetUpgradePrecheckByIDParams{
		timeout: timeout,
	}
}

// NewGetUpgradePrecheckByIDParamsWithContext creates a new GetUpgradePrecheckByIDParams object
// with the ability to set a context for a request.
func NewGetUpgradePrecheckByIDParamsWithContext(ctx context.Context) *GetUpgradePrecheckByIDParams {
	return &GetUpgradePrecheckByIDParams{
		Context: ctx,
	}
}

// NewGetUpgradePrecheckByIDParamsWithHTTPClient creates a new GetUpgradePrecheckByIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetUpgradePrecheckByIDParamsWithHTTPClient(client *http.Client) *GetUpgradePrecheckByIDParams {
	return &GetUpgradePrecheckByIDParams{
		HTTPClient: client,
	}
}

/*
GetUpgradePrecheckByIDParams contains all the parameters to send to the API endpoint

	for the get upgrade precheck by ID operation.

	Typically these are written to a http.Request.
*/
type GetUpgradePrecheckByIDParams struct {

	/* PrecheckID.

	   precheckId
	*/
	PrecheckID string

	/* UpgradeID.

	   upgradeId
	*/
	UpgradeID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get upgrade precheck by ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetUpgradePrecheckByIDParams) WithDefaults() *GetUpgradePrecheckByIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get upgrade precheck by ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetUpgradePrecheckByIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get upgrade precheck by ID params
func (o *GetUpgradePrecheckByIDParams) WithTimeout(timeout time.Duration) *GetUpgradePrecheckByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get upgrade precheck by ID params
func (o *GetUpgradePrecheckByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get upgrade precheck by ID params
func (o *GetUpgradePrecheckByIDParams) WithContext(ctx context.Context) *GetUpgradePrecheckByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get upgrade precheck by ID params
func (o *GetUpgradePrecheckByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get upgrade precheck by ID params
func (o *GetUpgradePrecheckByIDParams) WithHTTPClient(client *http.Client) *GetUpgradePrecheckByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get upgrade precheck by ID params
func (o *GetUpgradePrecheckByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPrecheckID adds the precheckID to the get upgrade precheck by ID params
func (o *GetUpgradePrecheckByIDParams) WithPrecheckID(precheckID string) *GetUpgradePrecheckByIDParams {
	o.SetPrecheckID(precheckID)
	return o
}

// SetPrecheckID adds the precheckId to the get upgrade precheck by ID params
func (o *GetUpgradePrecheckByIDParams) SetPrecheckID(precheckID string) {
	o.PrecheckID = precheckID
}

// WithUpgradeID adds the upgradeID to the get upgrade precheck by ID params
func (o *GetUpgradePrecheckByIDParams) WithUpgradeID(upgradeID string) *GetUpgradePrecheckByIDParams {
	o.SetUpgradeID(upgradeID)
	return o
}

// SetUpgradeID adds the upgradeId to the get upgrade precheck by ID params
func (o *GetUpgradePrecheckByIDParams) SetUpgradeID(upgradeID string) {
	o.UpgradeID = upgradeID
}

// WriteToRequest writes these params to a swagger request
func (o *GetUpgradePrecheckByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param precheckId
	if err := r.SetPathParam("precheckId", o.PrecheckID); err != nil {
		return err
	}

	// path param upgradeId
	if err := r.SetPathParam("upgradeId", o.UpgradeID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
