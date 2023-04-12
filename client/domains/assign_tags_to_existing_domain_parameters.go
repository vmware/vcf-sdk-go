// Code generated by go-swagger; DO NOT EDIT.

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

// NewAssignTagsToExistingDomainParams creates a new AssignTagsToExistingDomainParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAssignTagsToExistingDomainParams() *AssignTagsToExistingDomainParams {
	return &AssignTagsToExistingDomainParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAssignTagsToExistingDomainParamsWithTimeout creates a new AssignTagsToExistingDomainParams object
// with the ability to set a timeout on a request.
func NewAssignTagsToExistingDomainParamsWithTimeout(timeout time.Duration) *AssignTagsToExistingDomainParams {
	return &AssignTagsToExistingDomainParams{
		timeout: timeout,
	}
}

// NewAssignTagsToExistingDomainParamsWithContext creates a new AssignTagsToExistingDomainParams object
// with the ability to set a context for a request.
func NewAssignTagsToExistingDomainParamsWithContext(ctx context.Context) *AssignTagsToExistingDomainParams {
	return &AssignTagsToExistingDomainParams{
		Context: ctx,
	}
}

// NewAssignTagsToExistingDomainParamsWithHTTPClient creates a new AssignTagsToExistingDomainParams object
// with the ability to set a custom HTTPClient for a request.
func NewAssignTagsToExistingDomainParamsWithHTTPClient(client *http.Client) *AssignTagsToExistingDomainParams {
	return &AssignTagsToExistingDomainParams{
		HTTPClient: client,
	}
}

/*
AssignTagsToExistingDomainParams contains all the parameters to send to the API endpoint

	for the assign tags to existing domain operation.

	Typically these are written to a http.Request.
*/
type AssignTagsToExistingDomainParams struct {

	/* ID.

	   Domain ID
	*/
	ID string

	/* TagsSpec.

	   Tags Spec
	*/
	TagsSpec *models.TagsSpec

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the assign tags to existing domain params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AssignTagsToExistingDomainParams) WithDefaults() *AssignTagsToExistingDomainParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the assign tags to existing domain params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AssignTagsToExistingDomainParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the assign tags to existing domain params
func (o *AssignTagsToExistingDomainParams) WithTimeout(timeout time.Duration) *AssignTagsToExistingDomainParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the assign tags to existing domain params
func (o *AssignTagsToExistingDomainParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the assign tags to existing domain params
func (o *AssignTagsToExistingDomainParams) WithContext(ctx context.Context) *AssignTagsToExistingDomainParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the assign tags to existing domain params
func (o *AssignTagsToExistingDomainParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the assign tags to existing domain params
func (o *AssignTagsToExistingDomainParams) WithHTTPClient(client *http.Client) *AssignTagsToExistingDomainParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the assign tags to existing domain params
func (o *AssignTagsToExistingDomainParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the assign tags to existing domain params
func (o *AssignTagsToExistingDomainParams) WithID(id string) *AssignTagsToExistingDomainParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the assign tags to existing domain params
func (o *AssignTagsToExistingDomainParams) SetID(id string) {
	o.ID = id
}

// WithTagsSpec adds the tagsSpec to the assign tags to existing domain params
func (o *AssignTagsToExistingDomainParams) WithTagsSpec(tagsSpec *models.TagsSpec) *AssignTagsToExistingDomainParams {
	o.SetTagsSpec(tagsSpec)
	return o
}

// SetTagsSpec adds the tagsSpec to the assign tags to existing domain params
func (o *AssignTagsToExistingDomainParams) SetTagsSpec(tagsSpec *models.TagsSpec) {
	o.TagsSpec = tagsSpec
}

// WriteToRequest writes these params to a swagger request
func (o *AssignTagsToExistingDomainParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}
	if o.TagsSpec != nil {
		if err := r.SetBodyParam(o.TagsSpec); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
