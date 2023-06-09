// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2023 VMware, Inc.
// SPDX-License-Identifier: BSD-2-Clause

package clusters

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/vmware/vcf-sdk-go/models"
)

// ValidateClusterOperationsReader is a Reader for the ValidateClusterOperations structure.
type ValidateClusterOperationsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ValidateClusterOperationsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewValidateClusterOperationsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewValidateClusterOperationsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewValidateClusterOperationsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewValidateClusterOperationsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewValidateClusterOperationsOK creates a ValidateClusterOperationsOK with default headers values
func NewValidateClusterOperationsOK() *ValidateClusterOperationsOK {
	return &ValidateClusterOperationsOK{}
}

/*
ValidateClusterOperationsOK describes a response with status code 200, with default header values.

Ok
*/
type ValidateClusterOperationsOK struct {
	Payload *models.Validation
}

// IsSuccess returns true when this validate cluster operations o k response has a 2xx status code
func (o *ValidateClusterOperationsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this validate cluster operations o k response has a 3xx status code
func (o *ValidateClusterOperationsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this validate cluster operations o k response has a 4xx status code
func (o *ValidateClusterOperationsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this validate cluster operations o k response has a 5xx status code
func (o *ValidateClusterOperationsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this validate cluster operations o k response a status code equal to that given
func (o *ValidateClusterOperationsOK) IsCode(code int) bool {
	return code == 200
}

func (o *ValidateClusterOperationsOK) Error() string {
	return fmt.Sprintf("[POST /v1/clusters/{id}/validations][%d] validateClusterOperationsOK  %+v", 200, o.Payload)
}

func (o *ValidateClusterOperationsOK) String() string {
	return fmt.Sprintf("[POST /v1/clusters/{id}/validations][%d] validateClusterOperationsOK  %+v", 200, o.Payload)
}

func (o *ValidateClusterOperationsOK) GetPayload() *models.Validation {
	return o.Payload
}

func (o *ValidateClusterOperationsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Validation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewValidateClusterOperationsBadRequest creates a ValidateClusterOperationsBadRequest with default headers values
func NewValidateClusterOperationsBadRequest() *ValidateClusterOperationsBadRequest {
	return &ValidateClusterOperationsBadRequest{}
}

/*
ValidateClusterOperationsBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type ValidateClusterOperationsBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this validate cluster operations bad request response has a 2xx status code
func (o *ValidateClusterOperationsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this validate cluster operations bad request response has a 3xx status code
func (o *ValidateClusterOperationsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this validate cluster operations bad request response has a 4xx status code
func (o *ValidateClusterOperationsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this validate cluster operations bad request response has a 5xx status code
func (o *ValidateClusterOperationsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this validate cluster operations bad request response a status code equal to that given
func (o *ValidateClusterOperationsBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *ValidateClusterOperationsBadRequest) Error() string {
	return fmt.Sprintf("[POST /v1/clusters/{id}/validations][%d] validateClusterOperationsBadRequest  %+v", 400, o.Payload)
}

func (o *ValidateClusterOperationsBadRequest) String() string {
	return fmt.Sprintf("[POST /v1/clusters/{id}/validations][%d] validateClusterOperationsBadRequest  %+v", 400, o.Payload)
}

func (o *ValidateClusterOperationsBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *ValidateClusterOperationsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewValidateClusterOperationsNotFound creates a ValidateClusterOperationsNotFound with default headers values
func NewValidateClusterOperationsNotFound() *ValidateClusterOperationsNotFound {
	return &ValidateClusterOperationsNotFound{}
}

/*
ValidateClusterOperationsNotFound describes a response with status code 404, with default header values.

Not Found
*/
type ValidateClusterOperationsNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this validate cluster operations not found response has a 2xx status code
func (o *ValidateClusterOperationsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this validate cluster operations not found response has a 3xx status code
func (o *ValidateClusterOperationsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this validate cluster operations not found response has a 4xx status code
func (o *ValidateClusterOperationsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this validate cluster operations not found response has a 5xx status code
func (o *ValidateClusterOperationsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this validate cluster operations not found response a status code equal to that given
func (o *ValidateClusterOperationsNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *ValidateClusterOperationsNotFound) Error() string {
	return fmt.Sprintf("[POST /v1/clusters/{id}/validations][%d] validateClusterOperationsNotFound  %+v", 404, o.Payload)
}

func (o *ValidateClusterOperationsNotFound) String() string {
	return fmt.Sprintf("[POST /v1/clusters/{id}/validations][%d] validateClusterOperationsNotFound  %+v", 404, o.Payload)
}

func (o *ValidateClusterOperationsNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *ValidateClusterOperationsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewValidateClusterOperationsInternalServerError creates a ValidateClusterOperationsInternalServerError with default headers values
func NewValidateClusterOperationsInternalServerError() *ValidateClusterOperationsInternalServerError {
	return &ValidateClusterOperationsInternalServerError{}
}

/*
ValidateClusterOperationsInternalServerError describes a response with status code 500, with default header values.

InternalServerError
*/
type ValidateClusterOperationsInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this validate cluster operations internal server error response has a 2xx status code
func (o *ValidateClusterOperationsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this validate cluster operations internal server error response has a 3xx status code
func (o *ValidateClusterOperationsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this validate cluster operations internal server error response has a 4xx status code
func (o *ValidateClusterOperationsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this validate cluster operations internal server error response has a 5xx status code
func (o *ValidateClusterOperationsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this validate cluster operations internal server error response a status code equal to that given
func (o *ValidateClusterOperationsInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *ValidateClusterOperationsInternalServerError) Error() string {
	return fmt.Sprintf("[POST /v1/clusters/{id}/validations][%d] validateClusterOperationsInternalServerError  %+v", 500, o.Payload)
}

func (o *ValidateClusterOperationsInternalServerError) String() string {
	return fmt.Sprintf("[POST /v1/clusters/{id}/validations][%d] validateClusterOperationsInternalServerError  %+v", 500, o.Payload)
}

func (o *ValidateClusterOperationsInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *ValidateClusterOperationsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
