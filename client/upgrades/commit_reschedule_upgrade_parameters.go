// Code generated by go-swagger; DO NOT EDIT.

// Copyright © 2023 VMware, Inc. All Rights Reserved.
// Code licenced under: BSD-2

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

	"vcf-sdk-go/models"
)

// NewCommitRescheduleUpgradeParams creates a new CommitRescheduleUpgradeParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCommitRescheduleUpgradeParams() *CommitRescheduleUpgradeParams {
	return &CommitRescheduleUpgradeParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCommitRescheduleUpgradeParamsWithTimeout creates a new CommitRescheduleUpgradeParams object
// with the ability to set a timeout on a request.
func NewCommitRescheduleUpgradeParamsWithTimeout(timeout time.Duration) *CommitRescheduleUpgradeParams {
	return &CommitRescheduleUpgradeParams{
		timeout: timeout,
	}
}

// NewCommitRescheduleUpgradeParamsWithContext creates a new CommitRescheduleUpgradeParams object
// with the ability to set a context for a request.
func NewCommitRescheduleUpgradeParamsWithContext(ctx context.Context) *CommitRescheduleUpgradeParams {
	return &CommitRescheduleUpgradeParams{
		Context: ctx,
	}
}

// NewCommitRescheduleUpgradeParamsWithHTTPClient creates a new CommitRescheduleUpgradeParams object
// with the ability to set a custom HTTPClient for a request.
func NewCommitRescheduleUpgradeParamsWithHTTPClient(client *http.Client) *CommitRescheduleUpgradeParams {
	return &CommitRescheduleUpgradeParams{
		HTTPClient: client,
	}
}

/*
CommitRescheduleUpgradeParams contains all the parameters to send to the API endpoint

	for the commit reschedule upgrade operation.

	Typically these are written to a http.Request.
*/
type CommitRescheduleUpgradeParams struct {

	/* UpgradeCommitSpec.

	   Upgrade Commit/Reschedule Specification
	*/
	UpgradeCommitSpec *models.UpgradeCommitSpec

	/* UpgradeID.

	   upgradeId
	*/
	UpgradeID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the commit reschedule upgrade params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CommitRescheduleUpgradeParams) WithDefaults() *CommitRescheduleUpgradeParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the commit reschedule upgrade params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CommitRescheduleUpgradeParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the commit reschedule upgrade params
func (o *CommitRescheduleUpgradeParams) WithTimeout(timeout time.Duration) *CommitRescheduleUpgradeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the commit reschedule upgrade params
func (o *CommitRescheduleUpgradeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the commit reschedule upgrade params
func (o *CommitRescheduleUpgradeParams) WithContext(ctx context.Context) *CommitRescheduleUpgradeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the commit reschedule upgrade params
func (o *CommitRescheduleUpgradeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the commit reschedule upgrade params
func (o *CommitRescheduleUpgradeParams) WithHTTPClient(client *http.Client) *CommitRescheduleUpgradeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the commit reschedule upgrade params
func (o *CommitRescheduleUpgradeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUpgradeCommitSpec adds the upgradeCommitSpec to the commit reschedule upgrade params
func (o *CommitRescheduleUpgradeParams) WithUpgradeCommitSpec(upgradeCommitSpec *models.UpgradeCommitSpec) *CommitRescheduleUpgradeParams {
	o.SetUpgradeCommitSpec(upgradeCommitSpec)
	return o
}

// SetUpgradeCommitSpec adds the upgradeCommitSpec to the commit reschedule upgrade params
func (o *CommitRescheduleUpgradeParams) SetUpgradeCommitSpec(upgradeCommitSpec *models.UpgradeCommitSpec) {
	o.UpgradeCommitSpec = upgradeCommitSpec
}

// WithUpgradeID adds the upgradeID to the commit reschedule upgrade params
func (o *CommitRescheduleUpgradeParams) WithUpgradeID(upgradeID string) *CommitRescheduleUpgradeParams {
	o.SetUpgradeID(upgradeID)
	return o
}

// SetUpgradeID adds the upgradeId to the commit reschedule upgrade params
func (o *CommitRescheduleUpgradeParams) SetUpgradeID(upgradeID string) {
	o.UpgradeID = upgradeID
}

// WriteToRequest writes these params to a swagger request
func (o *CommitRescheduleUpgradeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.UpgradeCommitSpec != nil {
		if err := r.SetBodyParam(o.UpgradeCommitSpec); err != nil {
			return err
		}
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
