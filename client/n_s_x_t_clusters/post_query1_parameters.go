// Code generated by go-swagger; DO NOT EDIT.

// Copyright © 2023 VMware, Inc. All Rights Reserved.
// Code licenced under: BSD-2

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

	"github.com/vmware/vcf-sdk-go/models"
)

// NewPOSTQuery1Params creates a new POSTQuery1Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPOSTQuery1Params() *POSTQuery1Params {
	return &POSTQuery1Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewPOSTQuery1ParamsWithTimeout creates a new POSTQuery1Params object
// with the ability to set a timeout on a request.
func NewPOSTQuery1ParamsWithTimeout(timeout time.Duration) *POSTQuery1Params {
	return &POSTQuery1Params{
		timeout: timeout,
	}
}

// NewPOSTQuery1ParamsWithContext creates a new POSTQuery1Params object
// with the ability to set a context for a request.
func NewPOSTQuery1ParamsWithContext(ctx context.Context) *POSTQuery1Params {
	return &POSTQuery1Params{
		Context: ctx,
	}
}

// NewPOSTQuery1ParamsWithHTTPClient creates a new POSTQuery1Params object
// with the ability to set a custom HTTPClient for a request.
func NewPOSTQuery1ParamsWithHTTPClient(client *http.Client) *POSTQuery1Params {
	return &POSTQuery1Params{
		HTTPClient: client,
	}
}

/*
POSTQuery1Params contains all the parameters to send to the API endpoint

	for the post query 1 operation.

	Typically these are written to a http.Request.
*/
type POSTQuery1Params struct {

	/* NSXTCriterion.

	   nsxtCriterion
	*/
	NSXTCriterion *models.NsxTCriterion

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post query 1 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *POSTQuery1Params) WithDefaults() *POSTQuery1Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post query 1 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *POSTQuery1Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post query 1 params
func (o *POSTQuery1Params) WithTimeout(timeout time.Duration) *POSTQuery1Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post query 1 params
func (o *POSTQuery1Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post query 1 params
func (o *POSTQuery1Params) WithContext(ctx context.Context) *POSTQuery1Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post query 1 params
func (o *POSTQuery1Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post query 1 params
func (o *POSTQuery1Params) WithHTTPClient(client *http.Client) *POSTQuery1Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post query 1 params
func (o *POSTQuery1Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNSXTCriterion adds the nSXTCriterion to the post query 1 params
func (o *POSTQuery1Params) WithNSXTCriterion(nSXTCriterion *models.NsxTCriterion) *POSTQuery1Params {
	o.SetNSXTCriterion(nSXTCriterion)
	return o
}

// SetNSXTCriterion adds the nsxtCriterion to the post query 1 params
func (o *POSTQuery1Params) SetNSXTCriterion(nSXTCriterion *models.NsxTCriterion) {
	o.NSXTCriterion = nSXTCriterion
}

// WriteToRequest writes these params to a swagger request
func (o *POSTQuery1Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.NSXTCriterion != nil {
		if err := r.SetBodyParam(o.NSXTCriterion); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
