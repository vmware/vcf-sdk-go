// Code generated by go-swagger; DO NOT EDIT.

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
)

// NewGetPersonalitiesParams creates a new GetPersonalitiesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetPersonalitiesParams() *GetPersonalitiesParams {
	return &GetPersonalitiesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetPersonalitiesParamsWithTimeout creates a new GetPersonalitiesParams object
// with the ability to set a timeout on a request.
func NewGetPersonalitiesParamsWithTimeout(timeout time.Duration) *GetPersonalitiesParams {
	return &GetPersonalitiesParams{
		timeout: timeout,
	}
}

// NewGetPersonalitiesParamsWithContext creates a new GetPersonalitiesParams object
// with the ability to set a context for a request.
func NewGetPersonalitiesParamsWithContext(ctx context.Context) *GetPersonalitiesParams {
	return &GetPersonalitiesParams{
		Context: ctx,
	}
}

// NewGetPersonalitiesParamsWithHTTPClient creates a new GetPersonalitiesParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetPersonalitiesParamsWithHTTPClient(client *http.Client) *GetPersonalitiesParams {
	return &GetPersonalitiesParams{
		HTTPClient: client,
	}
}

/*
GetPersonalitiesParams contains all the parameters to send to the API endpoint

	for the get personalities operation.

	Typically these are written to a http.Request.
*/
type GetPersonalitiesParams struct {

	/* AddOnName.

	   The add on name
	*/
	AddOnName *string

	/* AddOnVendorName.

	   The add on vendor name
	*/
	AddOnVendorName *string

	/* BaseOSVersion.

	   The base OS version
	*/
	BaseOSVersion *string

	/* ComponentName.

	   The component name
	*/
	ComponentName *string

	/* ComponentVendorName.

	   The component vendor name
	*/
	ComponentVendorName *string

	/* PersonalityName.

	   personalityName
	*/
	PersonalityName *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get personalities params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetPersonalitiesParams) WithDefaults() *GetPersonalitiesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get personalities params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetPersonalitiesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get personalities params
func (o *GetPersonalitiesParams) WithTimeout(timeout time.Duration) *GetPersonalitiesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get personalities params
func (o *GetPersonalitiesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get personalities params
func (o *GetPersonalitiesParams) WithContext(ctx context.Context) *GetPersonalitiesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get personalities params
func (o *GetPersonalitiesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get personalities params
func (o *GetPersonalitiesParams) WithHTTPClient(client *http.Client) *GetPersonalitiesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get personalities params
func (o *GetPersonalitiesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAddOnName adds the addOnName to the get personalities params
func (o *GetPersonalitiesParams) WithAddOnName(addOnName *string) *GetPersonalitiesParams {
	o.SetAddOnName(addOnName)
	return o
}

// SetAddOnName adds the addOnName to the get personalities params
func (o *GetPersonalitiesParams) SetAddOnName(addOnName *string) {
	o.AddOnName = addOnName
}

// WithAddOnVendorName adds the addOnVendorName to the get personalities params
func (o *GetPersonalitiesParams) WithAddOnVendorName(addOnVendorName *string) *GetPersonalitiesParams {
	o.SetAddOnVendorName(addOnVendorName)
	return o
}

// SetAddOnVendorName adds the addOnVendorName to the get personalities params
func (o *GetPersonalitiesParams) SetAddOnVendorName(addOnVendorName *string) {
	o.AddOnVendorName = addOnVendorName
}

// WithBaseOSVersion adds the baseOSVersion to the get personalities params
func (o *GetPersonalitiesParams) WithBaseOSVersion(baseOSVersion *string) *GetPersonalitiesParams {
	o.SetBaseOSVersion(baseOSVersion)
	return o
}

// SetBaseOSVersion adds the baseOSVersion to the get personalities params
func (o *GetPersonalitiesParams) SetBaseOSVersion(baseOSVersion *string) {
	o.BaseOSVersion = baseOSVersion
}

// WithComponentName adds the componentName to the get personalities params
func (o *GetPersonalitiesParams) WithComponentName(componentName *string) *GetPersonalitiesParams {
	o.SetComponentName(componentName)
	return o
}

// SetComponentName adds the componentName to the get personalities params
func (o *GetPersonalitiesParams) SetComponentName(componentName *string) {
	o.ComponentName = componentName
}

// WithComponentVendorName adds the componentVendorName to the get personalities params
func (o *GetPersonalitiesParams) WithComponentVendorName(componentVendorName *string) *GetPersonalitiesParams {
	o.SetComponentVendorName(componentVendorName)
	return o
}

// SetComponentVendorName adds the componentVendorName to the get personalities params
func (o *GetPersonalitiesParams) SetComponentVendorName(componentVendorName *string) {
	o.ComponentVendorName = componentVendorName
}

// WithPersonalityName adds the personalityName to the get personalities params
func (o *GetPersonalitiesParams) WithPersonalityName(personalityName *string) *GetPersonalitiesParams {
	o.SetPersonalityName(personalityName)
	return o
}

// SetPersonalityName adds the personalityName to the get personalities params
func (o *GetPersonalitiesParams) SetPersonalityName(personalityName *string) {
	o.PersonalityName = personalityName
}

// WriteToRequest writes these params to a swagger request
func (o *GetPersonalitiesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.AddOnName != nil {

		// query param addOnName
		var qrAddOnName string

		if o.AddOnName != nil {
			qrAddOnName = *o.AddOnName
		}
		qAddOnName := qrAddOnName
		if qAddOnName != "" {

			if err := r.SetQueryParam("addOnName", qAddOnName); err != nil {
				return err
			}
		}
	}

	if o.AddOnVendorName != nil {

		// query param addOnVendorName
		var qrAddOnVendorName string

		if o.AddOnVendorName != nil {
			qrAddOnVendorName = *o.AddOnVendorName
		}
		qAddOnVendorName := qrAddOnVendorName
		if qAddOnVendorName != "" {

			if err := r.SetQueryParam("addOnVendorName", qAddOnVendorName); err != nil {
				return err
			}
		}
	}

	if o.BaseOSVersion != nil {

		// query param baseOSVersion
		var qrBaseOSVersion string

		if o.BaseOSVersion != nil {
			qrBaseOSVersion = *o.BaseOSVersion
		}
		qBaseOSVersion := qrBaseOSVersion
		if qBaseOSVersion != "" {

			if err := r.SetQueryParam("baseOSVersion", qBaseOSVersion); err != nil {
				return err
			}
		}
	}

	if o.ComponentName != nil {

		// query param componentName
		var qrComponentName string

		if o.ComponentName != nil {
			qrComponentName = *o.ComponentName
		}
		qComponentName := qrComponentName
		if qComponentName != "" {

			if err := r.SetQueryParam("componentName", qComponentName); err != nil {
				return err
			}
		}
	}

	if o.ComponentVendorName != nil {

		// query param componentVendorName
		var qrComponentVendorName string

		if o.ComponentVendorName != nil {
			qrComponentVendorName = *o.ComponentVendorName
		}
		qComponentVendorName := qrComponentVendorName
		if qComponentVendorName != "" {

			if err := r.SetQueryParam("componentVendorName", qComponentVendorName); err != nil {
				return err
			}
		}
	}

	if o.PersonalityName != nil {

		// query param personalityName
		var qrPersonalityName string

		if o.PersonalityName != nil {
			qrPersonalityName = *o.PersonalityName
		}
		qPersonalityName := qrPersonalityName
		if qPersonalityName != "" {

			if err := r.SetQueryParam("personalityName", qPersonalityName); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
