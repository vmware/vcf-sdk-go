// Code generated by go-swagger; DO NOT EDIT.

package tasks

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
	"github.com/go-openapi/swag"
)

// NewGetTasksParams creates a new GetTasksParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetTasksParams() *GetTasksParams {
	return &GetTasksParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetTasksParamsWithTimeout creates a new GetTasksParams object
// with the ability to set a timeout on a request.
func NewGetTasksParamsWithTimeout(timeout time.Duration) *GetTasksParams {
	return &GetTasksParams{
		timeout: timeout,
	}
}

// NewGetTasksParamsWithContext creates a new GetTasksParams object
// with the ability to set a context for a request.
func NewGetTasksParamsWithContext(ctx context.Context) *GetTasksParams {
	return &GetTasksParams{
		Context: ctx,
	}
}

// NewGetTasksParamsWithHTTPClient creates a new GetTasksParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetTasksParamsWithHTTPClient(client *http.Client) *GetTasksParams {
	return &GetTasksParams{
		HTTPClient: client,
	}
}

/*
GetTasksParams contains all the parameters to send to the API endpoint

	for the get tasks operation.

	Typically these are written to a http.Request.
*/
type GetTasksParams struct {

	/* CompletedAfter.

	   A time based filter to get tasks which are completed after the given timestamp. A task is completed if its status is 'Successsful' or 'Failed'. Time is in milliseconds.

	   Format: int64
	*/
	CompletedAfter *int64

	/* Limit.

	   The number of elements to be returned in the result

	   Format: int32
	*/
	Limit *int32

	/* ResourceID.

	   resourceId
	*/
	ResourceID *string

	/* ResourceType.

	   resourceType
	*/
	ResourceType *string

	/* TaskStatus.

	   taskStatus
	*/
	TaskStatus *string

	/* TaskType.

	   taskType
	*/
	TaskType *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get tasks params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetTasksParams) WithDefaults() *GetTasksParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get tasks params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetTasksParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get tasks params
func (o *GetTasksParams) WithTimeout(timeout time.Duration) *GetTasksParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get tasks params
func (o *GetTasksParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get tasks params
func (o *GetTasksParams) WithContext(ctx context.Context) *GetTasksParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get tasks params
func (o *GetTasksParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get tasks params
func (o *GetTasksParams) WithHTTPClient(client *http.Client) *GetTasksParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get tasks params
func (o *GetTasksParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCompletedAfter adds the completedAfter to the get tasks params
func (o *GetTasksParams) WithCompletedAfter(completedAfter *int64) *GetTasksParams {
	o.SetCompletedAfter(completedAfter)
	return o
}

// SetCompletedAfter adds the completedAfter to the get tasks params
func (o *GetTasksParams) SetCompletedAfter(completedAfter *int64) {
	o.CompletedAfter = completedAfter
}

// WithLimit adds the limit to the get tasks params
func (o *GetTasksParams) WithLimit(limit *int32) *GetTasksParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the get tasks params
func (o *GetTasksParams) SetLimit(limit *int32) {
	o.Limit = limit
}

// WithResourceID adds the resourceID to the get tasks params
func (o *GetTasksParams) WithResourceID(resourceID *string) *GetTasksParams {
	o.SetResourceID(resourceID)
	return o
}

// SetResourceID adds the resourceId to the get tasks params
func (o *GetTasksParams) SetResourceID(resourceID *string) {
	o.ResourceID = resourceID
}

// WithResourceType adds the resourceType to the get tasks params
func (o *GetTasksParams) WithResourceType(resourceType *string) *GetTasksParams {
	o.SetResourceType(resourceType)
	return o
}

// SetResourceType adds the resourceType to the get tasks params
func (o *GetTasksParams) SetResourceType(resourceType *string) {
	o.ResourceType = resourceType
}

// WithTaskStatus adds the taskStatus to the get tasks params
func (o *GetTasksParams) WithTaskStatus(taskStatus *string) *GetTasksParams {
	o.SetTaskStatus(taskStatus)
	return o
}

// SetTaskStatus adds the taskStatus to the get tasks params
func (o *GetTasksParams) SetTaskStatus(taskStatus *string) {
	o.TaskStatus = taskStatus
}

// WithTaskType adds the taskType to the get tasks params
func (o *GetTasksParams) WithTaskType(taskType *string) *GetTasksParams {
	o.SetTaskType(taskType)
	return o
}

// SetTaskType adds the taskType to the get tasks params
func (o *GetTasksParams) SetTaskType(taskType *string) {
	o.TaskType = taskType
}

// WriteToRequest writes these params to a swagger request
func (o *GetTasksParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.CompletedAfter != nil {

		// query param completedAfter
		var qrCompletedAfter int64

		if o.CompletedAfter != nil {
			qrCompletedAfter = *o.CompletedAfter
		}
		qCompletedAfter := swag.FormatInt64(qrCompletedAfter)
		if qCompletedAfter != "" {

			if err := r.SetQueryParam("completedAfter", qCompletedAfter); err != nil {
				return err
			}
		}
	}

	if o.Limit != nil {

		// query param limit
		var qrLimit int32

		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := swag.FormatInt32(qrLimit)
		if qLimit != "" {

			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}
	}

	if o.ResourceID != nil {

		// query param resourceId
		var qrResourceID string

		if o.ResourceID != nil {
			qrResourceID = *o.ResourceID
		}
		qResourceID := qrResourceID
		if qResourceID != "" {

			if err := r.SetQueryParam("resourceId", qResourceID); err != nil {
				return err
			}
		}
	}

	if o.ResourceType != nil {

		// query param resourceType
		var qrResourceType string

		if o.ResourceType != nil {
			qrResourceType = *o.ResourceType
		}
		qResourceType := qrResourceType
		if qResourceType != "" {

			if err := r.SetQueryParam("resourceType", qResourceType); err != nil {
				return err
			}
		}
	}

	if o.TaskStatus != nil {

		// query param taskStatus
		var qrTaskStatus string

		if o.TaskStatus != nil {
			qrTaskStatus = *o.TaskStatus
		}
		qTaskStatus := qrTaskStatus
		if qTaskStatus != "" {

			if err := r.SetQueryParam("taskStatus", qTaskStatus); err != nil {
				return err
			}
		}
	}

	if o.TaskType != nil {

		// query param taskType
		var qrTaskType string

		if o.TaskType != nil {
			qrTaskType = *o.TaskType
		}
		qTaskType := qrTaskType
		if qTaskType != "" {

			if err := r.SetQueryParam("taskType", qTaskType); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}