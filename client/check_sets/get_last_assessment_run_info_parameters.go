// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2023 VMware, Inc.
// SPDX-License-Identifier: BSD-2-Clause

package check_sets

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

// NewGetLastAssessmentRunInfoParams creates a new GetLastAssessmentRunInfoParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetLastAssessmentRunInfoParams() *GetLastAssessmentRunInfoParams {
	return &GetLastAssessmentRunInfoParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetLastAssessmentRunInfoParamsWithTimeout creates a new GetLastAssessmentRunInfoParams object
// with the ability to set a timeout on a request.
func NewGetLastAssessmentRunInfoParamsWithTimeout(timeout time.Duration) *GetLastAssessmentRunInfoParams {
	return &GetLastAssessmentRunInfoParams{
		timeout: timeout,
	}
}

// NewGetLastAssessmentRunInfoParamsWithContext creates a new GetLastAssessmentRunInfoParams object
// with the ability to set a context for a request.
func NewGetLastAssessmentRunInfoParamsWithContext(ctx context.Context) *GetLastAssessmentRunInfoParams {
	return &GetLastAssessmentRunInfoParams{
		Context: ctx,
	}
}

// NewGetLastAssessmentRunInfoParamsWithHTTPClient creates a new GetLastAssessmentRunInfoParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetLastAssessmentRunInfoParamsWithHTTPClient(client *http.Client) *GetLastAssessmentRunInfoParams {
	return &GetLastAssessmentRunInfoParams{
		HTTPClient: client,
	}
}

/*
GetLastAssessmentRunInfoParams contains all the parameters to send to the API endpoint

	for the get last assessment run info operation.

	Typically these are written to a http.Request.
*/
type GetLastAssessmentRunInfoParams struct {

	/* DomainID.

	   Id of the domain to filter tasks for
	*/
	DomainID *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get last assessment run info params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetLastAssessmentRunInfoParams) WithDefaults() *GetLastAssessmentRunInfoParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get last assessment run info params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetLastAssessmentRunInfoParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get last assessment run info params
func (o *GetLastAssessmentRunInfoParams) WithTimeout(timeout time.Duration) *GetLastAssessmentRunInfoParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get last assessment run info params
func (o *GetLastAssessmentRunInfoParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get last assessment run info params
func (o *GetLastAssessmentRunInfoParams) WithContext(ctx context.Context) *GetLastAssessmentRunInfoParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get last assessment run info params
func (o *GetLastAssessmentRunInfoParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get last assessment run info params
func (o *GetLastAssessmentRunInfoParams) WithHTTPClient(client *http.Client) *GetLastAssessmentRunInfoParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get last assessment run info params
func (o *GetLastAssessmentRunInfoParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDomainID adds the domainID to the get last assessment run info params
func (o *GetLastAssessmentRunInfoParams) WithDomainID(domainID *string) *GetLastAssessmentRunInfoParams {
	o.SetDomainID(domainID)
	return o
}

// SetDomainID adds the domainId to the get last assessment run info params
func (o *GetLastAssessmentRunInfoParams) SetDomainID(domainID *string) {
	o.DomainID = domainID
}

// WriteToRequest writes these params to a swagger request
func (o *GetLastAssessmentRunInfoParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.DomainID != nil {

		// query param domainId
		var qrDomainID string

		if o.DomainID != nil {
			qrDomainID = *o.DomainID
		}
		qDomainID := qrDomainID
		if qDomainID != "" {

			if err := r.SetQueryParam("domainId", qDomainID); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
