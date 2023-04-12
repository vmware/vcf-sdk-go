// Code generated by go-swagger; DO NOT EDIT.

package clusters

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

// NewRemoveTagsFromExistingClusterParams creates a new RemoveTagsFromExistingClusterParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewRemoveTagsFromExistingClusterParams() *RemoveTagsFromExistingClusterParams {
	return &RemoveTagsFromExistingClusterParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewRemoveTagsFromExistingClusterParamsWithTimeout creates a new RemoveTagsFromExistingClusterParams object
// with the ability to set a timeout on a request.
func NewRemoveTagsFromExistingClusterParamsWithTimeout(timeout time.Duration) *RemoveTagsFromExistingClusterParams {
	return &RemoveTagsFromExistingClusterParams{
		timeout: timeout,
	}
}

// NewRemoveTagsFromExistingClusterParamsWithContext creates a new RemoveTagsFromExistingClusterParams object
// with the ability to set a context for a request.
func NewRemoveTagsFromExistingClusterParamsWithContext(ctx context.Context) *RemoveTagsFromExistingClusterParams {
	return &RemoveTagsFromExistingClusterParams{
		Context: ctx,
	}
}

// NewRemoveTagsFromExistingClusterParamsWithHTTPClient creates a new RemoveTagsFromExistingClusterParams object
// with the ability to set a custom HTTPClient for a request.
func NewRemoveTagsFromExistingClusterParamsWithHTTPClient(client *http.Client) *RemoveTagsFromExistingClusterParams {
	return &RemoveTagsFromExistingClusterParams{
		HTTPClient: client,
	}
}

/*
RemoveTagsFromExistingClusterParams contains all the parameters to send to the API endpoint

	for the remove tags from existing cluster operation.

	Typically these are written to a http.Request.
*/
type RemoveTagsFromExistingClusterParams struct {

	/* ID.

	   Cluster ID
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

// WithDefaults hydrates default values in the remove tags from existing cluster params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RemoveTagsFromExistingClusterParams) WithDefaults() *RemoveTagsFromExistingClusterParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the remove tags from existing cluster params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RemoveTagsFromExistingClusterParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the remove tags from existing cluster params
func (o *RemoveTagsFromExistingClusterParams) WithTimeout(timeout time.Duration) *RemoveTagsFromExistingClusterParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the remove tags from existing cluster params
func (o *RemoveTagsFromExistingClusterParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the remove tags from existing cluster params
func (o *RemoveTagsFromExistingClusterParams) WithContext(ctx context.Context) *RemoveTagsFromExistingClusterParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the remove tags from existing cluster params
func (o *RemoveTagsFromExistingClusterParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the remove tags from existing cluster params
func (o *RemoveTagsFromExistingClusterParams) WithHTTPClient(client *http.Client) *RemoveTagsFromExistingClusterParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the remove tags from existing cluster params
func (o *RemoveTagsFromExistingClusterParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the remove tags from existing cluster params
func (o *RemoveTagsFromExistingClusterParams) WithID(id string) *RemoveTagsFromExistingClusterParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the remove tags from existing cluster params
func (o *RemoveTagsFromExistingClusterParams) SetID(id string) {
	o.ID = id
}

// WithTagsSpec adds the tagsSpec to the remove tags from existing cluster params
func (o *RemoveTagsFromExistingClusterParams) WithTagsSpec(tagsSpec *models.TagsSpec) *RemoveTagsFromExistingClusterParams {
	o.SetTagsSpec(tagsSpec)
	return o
}

// SetTagsSpec adds the tagsSpec to the remove tags from existing cluster params
func (o *RemoveTagsFromExistingClusterParams) SetTagsSpec(tagsSpec *models.TagsSpec) {
	o.TagsSpec = tagsSpec
}

// WriteToRequest writes these params to a swagger request
func (o *RemoveTagsFromExistingClusterParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
