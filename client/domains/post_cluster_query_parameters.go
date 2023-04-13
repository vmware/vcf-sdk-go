// Code generated by go-swagger; DO NOT EDIT.

// Copyright © 2023 VMware, Inc. All Rights Reserved.
// Code licenced under: BSD-2

package domains

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

// NewPOSTClusterQueryParams creates a new POSTClusterQueryParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPOSTClusterQueryParams() *POSTClusterQueryParams {
	return &POSTClusterQueryParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPOSTClusterQueryParamsWithTimeout creates a new POSTClusterQueryParams object
// with the ability to set a timeout on a request.
func NewPOSTClusterQueryParamsWithTimeout(timeout time.Duration) *POSTClusterQueryParams {
	return &POSTClusterQueryParams{
		timeout: timeout,
	}
}

// NewPOSTClusterQueryParamsWithContext creates a new POSTClusterQueryParams object
// with the ability to set a context for a request.
func NewPOSTClusterQueryParamsWithContext(ctx context.Context) *POSTClusterQueryParams {
	return &POSTClusterQueryParams{
		Context: ctx,
	}
}

// NewPOSTClusterQueryParamsWithHTTPClient creates a new POSTClusterQueryParams object
// with the ability to set a custom HTTPClient for a request.
func NewPOSTClusterQueryParamsWithHTTPClient(client *http.Client) *POSTClusterQueryParams {
	return &POSTClusterQueryParams{
		HTTPClient: client,
	}
}

/*
POSTClusterQueryParams contains all the parameters to send to the API endpoint

	for the post cluster query operation.

	Typically these are written to a http.Request.
*/
type POSTClusterQueryParams struct {

	/* ClusterCriterion.

	   clusterCriterion
	*/
	ClusterCriterion *models.ClusterCriterion

	/* ClusterName.

	   Cluster Name
	*/
	ClusterName string

	/* DomainID.

	   Domain ID
	*/
	DomainID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post cluster query params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *POSTClusterQueryParams) WithDefaults() *POSTClusterQueryParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post cluster query params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *POSTClusterQueryParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post cluster query params
func (o *POSTClusterQueryParams) WithTimeout(timeout time.Duration) *POSTClusterQueryParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post cluster query params
func (o *POSTClusterQueryParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post cluster query params
func (o *POSTClusterQueryParams) WithContext(ctx context.Context) *POSTClusterQueryParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post cluster query params
func (o *POSTClusterQueryParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post cluster query params
func (o *POSTClusterQueryParams) WithHTTPClient(client *http.Client) *POSTClusterQueryParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post cluster query params
func (o *POSTClusterQueryParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClusterCriterion adds the clusterCriterion to the post cluster query params
func (o *POSTClusterQueryParams) WithClusterCriterion(clusterCriterion *models.ClusterCriterion) *POSTClusterQueryParams {
	o.SetClusterCriterion(clusterCriterion)
	return o
}

// SetClusterCriterion adds the clusterCriterion to the post cluster query params
func (o *POSTClusterQueryParams) SetClusterCriterion(clusterCriterion *models.ClusterCriterion) {
	o.ClusterCriterion = clusterCriterion
}

// WithClusterName adds the clusterName to the post cluster query params
func (o *POSTClusterQueryParams) WithClusterName(clusterName string) *POSTClusterQueryParams {
	o.SetClusterName(clusterName)
	return o
}

// SetClusterName adds the clusterName to the post cluster query params
func (o *POSTClusterQueryParams) SetClusterName(clusterName string) {
	o.ClusterName = clusterName
}

// WithDomainID adds the domainID to the post cluster query params
func (o *POSTClusterQueryParams) WithDomainID(domainID string) *POSTClusterQueryParams {
	o.SetDomainID(domainID)
	return o
}

// SetDomainID adds the domainId to the post cluster query params
func (o *POSTClusterQueryParams) SetDomainID(domainID string) {
	o.DomainID = domainID
}

// WriteToRequest writes these params to a swagger request
func (o *POSTClusterQueryParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.ClusterCriterion != nil {
		if err := r.SetBodyParam(o.ClusterCriterion); err != nil {
			return err
		}
	}

	// path param clusterName
	if err := r.SetPathParam("clusterName", o.ClusterName); err != nil {
		return err
	}

	// path param domainId
	if err := r.SetPathParam("domainId", o.DomainID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
