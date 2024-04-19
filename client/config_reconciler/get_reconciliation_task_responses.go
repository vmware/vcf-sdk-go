// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2023 VMware, Inc.
// SPDX-License-Identifier: BSD-2-Clause

package config_reconciler

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/vmware/vcf-sdk-go/models"
)

// GetReconciliationTaskReader is a Reader for the GetReconciliationTask structure.
type GetReconciliationTaskReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetReconciliationTaskReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetReconciliationTaskOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetReconciliationTaskBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetReconciliationTaskNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetReconciliationTaskInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /v1/config-drift-reconciliations/{taskId}] getReconciliationTask", response, response.Code())
	}
}

// NewGetReconciliationTaskOK creates a GetReconciliationTaskOK with default headers values
func NewGetReconciliationTaskOK() *GetReconciliationTaskOK {
	return &GetReconciliationTaskOK{}
}

/*
GetReconciliationTaskOK describes a response with status code 200, with default header values.

OK
*/
type GetReconciliationTaskOK struct {
	Payload *models.Task
}

// IsSuccess returns true when this get reconciliation task o k response has a 2xx status code
func (o *GetReconciliationTaskOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get reconciliation task o k response has a 3xx status code
func (o *GetReconciliationTaskOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get reconciliation task o k response has a 4xx status code
func (o *GetReconciliationTaskOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get reconciliation task o k response has a 5xx status code
func (o *GetReconciliationTaskOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get reconciliation task o k response a status code equal to that given
func (o *GetReconciliationTaskOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get reconciliation task o k response
func (o *GetReconciliationTaskOK) Code() int {
	return 200
}

func (o *GetReconciliationTaskOK) Error() string {
	return fmt.Sprintf("[GET /v1/config-drift-reconciliations/{taskId}][%d] getReconciliationTaskOK  %+v", 200, o.Payload)
}

func (o *GetReconciliationTaskOK) String() string {
	return fmt.Sprintf("[GET /v1/config-drift-reconciliations/{taskId}][%d] getReconciliationTaskOK  %+v", 200, o.Payload)
}

func (o *GetReconciliationTaskOK) GetPayload() *models.Task {
	return o.Payload
}

func (o *GetReconciliationTaskOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Task)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetReconciliationTaskBadRequest creates a GetReconciliationTaskBadRequest with default headers values
func NewGetReconciliationTaskBadRequest() *GetReconciliationTaskBadRequest {
	return &GetReconciliationTaskBadRequest{}
}

/*
GetReconciliationTaskBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetReconciliationTaskBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this get reconciliation task bad request response has a 2xx status code
func (o *GetReconciliationTaskBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get reconciliation task bad request response has a 3xx status code
func (o *GetReconciliationTaskBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get reconciliation task bad request response has a 4xx status code
func (o *GetReconciliationTaskBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get reconciliation task bad request response has a 5xx status code
func (o *GetReconciliationTaskBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get reconciliation task bad request response a status code equal to that given
func (o *GetReconciliationTaskBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get reconciliation task bad request response
func (o *GetReconciliationTaskBadRequest) Code() int {
	return 400
}

func (o *GetReconciliationTaskBadRequest) Error() string {
	return fmt.Sprintf("[GET /v1/config-drift-reconciliations/{taskId}][%d] getReconciliationTaskBadRequest  %+v", 400, o.Payload)
}

func (o *GetReconciliationTaskBadRequest) String() string {
	return fmt.Sprintf("[GET /v1/config-drift-reconciliations/{taskId}][%d] getReconciliationTaskBadRequest  %+v", 400, o.Payload)
}

func (o *GetReconciliationTaskBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetReconciliationTaskBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetReconciliationTaskNotFound creates a GetReconciliationTaskNotFound with default headers values
func NewGetReconciliationTaskNotFound() *GetReconciliationTaskNotFound {
	return &GetReconciliationTaskNotFound{}
}

/*
GetReconciliationTaskNotFound describes a response with status code 404, with default header values.

Reconciliation task not found
*/
type GetReconciliationTaskNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this get reconciliation task not found response has a 2xx status code
func (o *GetReconciliationTaskNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get reconciliation task not found response has a 3xx status code
func (o *GetReconciliationTaskNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get reconciliation task not found response has a 4xx status code
func (o *GetReconciliationTaskNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get reconciliation task not found response has a 5xx status code
func (o *GetReconciliationTaskNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get reconciliation task not found response a status code equal to that given
func (o *GetReconciliationTaskNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get reconciliation task not found response
func (o *GetReconciliationTaskNotFound) Code() int {
	return 404
}

func (o *GetReconciliationTaskNotFound) Error() string {
	return fmt.Sprintf("[GET /v1/config-drift-reconciliations/{taskId}][%d] getReconciliationTaskNotFound  %+v", 404, o.Payload)
}

func (o *GetReconciliationTaskNotFound) String() string {
	return fmt.Sprintf("[GET /v1/config-drift-reconciliations/{taskId}][%d] getReconciliationTaskNotFound  %+v", 404, o.Payload)
}

func (o *GetReconciliationTaskNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetReconciliationTaskNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetReconciliationTaskInternalServerError creates a GetReconciliationTaskInternalServerError with default headers values
func NewGetReconciliationTaskInternalServerError() *GetReconciliationTaskInternalServerError {
	return &GetReconciliationTaskInternalServerError{}
}

/*
GetReconciliationTaskInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetReconciliationTaskInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this get reconciliation task internal server error response has a 2xx status code
func (o *GetReconciliationTaskInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get reconciliation task internal server error response has a 3xx status code
func (o *GetReconciliationTaskInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get reconciliation task internal server error response has a 4xx status code
func (o *GetReconciliationTaskInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get reconciliation task internal server error response has a 5xx status code
func (o *GetReconciliationTaskInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get reconciliation task internal server error response a status code equal to that given
func (o *GetReconciliationTaskInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get reconciliation task internal server error response
func (o *GetReconciliationTaskInternalServerError) Code() int {
	return 500
}

func (o *GetReconciliationTaskInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/config-drift-reconciliations/{taskId}][%d] getReconciliationTaskInternalServerError  %+v", 500, o.Payload)
}

func (o *GetReconciliationTaskInternalServerError) String() string {
	return fmt.Sprintf("[GET /v1/config-drift-reconciliations/{taskId}][%d] getReconciliationTaskInternalServerError  %+v", 500, o.Payload)
}

func (o *GetReconciliationTaskInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetReconciliationTaskInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
