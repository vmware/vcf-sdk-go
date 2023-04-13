// Code generated by go-swagger; DO NOT EDIT.

// Copyright © 2023 VMware, Inc. All Rights Reserved.
// Code licenced under: BSD-2

package nsx_t_edge_clusters

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

// NewValidateEdgeClusterSpecParams creates a new ValidateEdgeClusterSpecParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewValidateEdgeClusterSpecParams() *ValidateEdgeClusterSpecParams {
	return &ValidateEdgeClusterSpecParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewValidateEdgeClusterSpecParamsWithTimeout creates a new ValidateEdgeClusterSpecParams object
// with the ability to set a timeout on a request.
func NewValidateEdgeClusterSpecParamsWithTimeout(timeout time.Duration) *ValidateEdgeClusterSpecParams {
	return &ValidateEdgeClusterSpecParams{
		timeout: timeout,
	}
}

// NewValidateEdgeClusterSpecParamsWithContext creates a new ValidateEdgeClusterSpecParams object
// with the ability to set a context for a request.
func NewValidateEdgeClusterSpecParamsWithContext(ctx context.Context) *ValidateEdgeClusterSpecParams {
	return &ValidateEdgeClusterSpecParams{
		Context: ctx,
	}
}

// NewValidateEdgeClusterSpecParamsWithHTTPClient creates a new ValidateEdgeClusterSpecParams object
// with the ability to set a custom HTTPClient for a request.
func NewValidateEdgeClusterSpecParamsWithHTTPClient(client *http.Client) *ValidateEdgeClusterSpecParams {
	return &ValidateEdgeClusterSpecParams{
		HTTPClient: client,
	}
}

/*
ValidateEdgeClusterSpecParams contains all the parameters to send to the API endpoint

	for the validate edge cluster spec operation.

	Typically these are written to a http.Request.
*/
type ValidateEdgeClusterSpecParams struct {

	/* EdgeCreationSpec.

	   NSX-T Edge cluster creation data to be validated
	*/
	EdgeCreationSpec *models.EdgeClusterCreationSpec

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the validate edge cluster spec params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ValidateEdgeClusterSpecParams) WithDefaults() *ValidateEdgeClusterSpecParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the validate edge cluster spec params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ValidateEdgeClusterSpecParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the validate edge cluster spec params
func (o *ValidateEdgeClusterSpecParams) WithTimeout(timeout time.Duration) *ValidateEdgeClusterSpecParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the validate edge cluster spec params
func (o *ValidateEdgeClusterSpecParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the validate edge cluster spec params
func (o *ValidateEdgeClusterSpecParams) WithContext(ctx context.Context) *ValidateEdgeClusterSpecParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the validate edge cluster spec params
func (o *ValidateEdgeClusterSpecParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the validate edge cluster spec params
func (o *ValidateEdgeClusterSpecParams) WithHTTPClient(client *http.Client) *ValidateEdgeClusterSpecParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the validate edge cluster spec params
func (o *ValidateEdgeClusterSpecParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEdgeCreationSpec adds the edgeCreationSpec to the validate edge cluster spec params
func (o *ValidateEdgeClusterSpecParams) WithEdgeCreationSpec(edgeCreationSpec *models.EdgeClusterCreationSpec) *ValidateEdgeClusterSpecParams {
	o.SetEdgeCreationSpec(edgeCreationSpec)
	return o
}

// SetEdgeCreationSpec adds the edgeCreationSpec to the validate edge cluster spec params
func (o *ValidateEdgeClusterSpecParams) SetEdgeCreationSpec(edgeCreationSpec *models.EdgeClusterCreationSpec) {
	o.EdgeCreationSpec = edgeCreationSpec
}

// WriteToRequest writes these params to a swagger request
func (o *ValidateEdgeClusterSpecParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.EdgeCreationSpec != nil {
		if err := r.SetBodyParam(o.EdgeCreationSpec); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
