// Code generated by go-swagger; DO NOT EDIT.

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

	"vcf-sdk-go/models"
)

// NewPostQuery1Params creates a new PostQuery1Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostQuery1Params() *PostQuery1Params {
	return &PostQuery1Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostQuery1ParamsWithTimeout creates a new PostQuery1Params object
// with the ability to set a timeout on a request.
func NewPostQuery1ParamsWithTimeout(timeout time.Duration) *PostQuery1Params {
	return &PostQuery1Params{
		timeout: timeout,
	}
}

// NewPostQuery1ParamsWithContext creates a new PostQuery1Params object
// with the ability to set a context for a request.
func NewPostQuery1ParamsWithContext(ctx context.Context) *PostQuery1Params {
	return &PostQuery1Params{
		Context: ctx,
	}
}

// NewPostQuery1ParamsWithHTTPClient creates a new PostQuery1Params object
// with the ability to set a custom HTTPClient for a request.
func NewPostQuery1ParamsWithHTTPClient(client *http.Client) *PostQuery1Params {
	return &PostQuery1Params{
		HTTPClient: client,
	}
}

/*
PostQuery1Params contains all the parameters to send to the API endpoint

	for the post query 1 operation.

	Typically these are written to a http.Request.
*/
type PostQuery1Params struct {

	/* NsxtCriterion.

	   nsxtCriterion
	*/
	NsxtCriterion *models.NsxTCriterion

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post query 1 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostQuery1Params) WithDefaults() *PostQuery1Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post query 1 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostQuery1Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post query 1 params
func (o *PostQuery1Params) WithTimeout(timeout time.Duration) *PostQuery1Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post query 1 params
func (o *PostQuery1Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post query 1 params
func (o *PostQuery1Params) WithContext(ctx context.Context) *PostQuery1Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post query 1 params
func (o *PostQuery1Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post query 1 params
func (o *PostQuery1Params) WithHTTPClient(client *http.Client) *PostQuery1Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post query 1 params
func (o *PostQuery1Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNsxtCriterion adds the nsxtCriterion to the post query 1 params
func (o *PostQuery1Params) WithNsxtCriterion(nsxtCriterion *models.NsxTCriterion) *PostQuery1Params {
	o.SetNsxtCriterion(nsxtCriterion)
	return o
}

// SetNsxtCriterion adds the nsxtCriterion to the post query 1 params
func (o *PostQuery1Params) SetNsxtCriterion(nsxtCriterion *models.NsxTCriterion) {
	o.NsxtCriterion = nsxtCriterion
}

// WriteToRequest writes these params to a swagger request
func (o *PostQuery1Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.NsxtCriterion != nil {
		if err := r.SetBodyParam(o.NsxtCriterion); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}