// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2023 VMware, Inc.
// SPDX-License-Identifier: BSD-2-Clause

package sddc

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/vmware/vcf-sdk-go/models"
)

// RetrySDDCValidationReader is a Reader for the RetrySDDCValidation structure.
type RetrySDDCValidationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RetrySDDCValidationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRetrySDDCValidationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewRetrySDDCValidationBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 405:
		result := NewRetrySDDCValidationMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewRetrySDDCValidationInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewRetrySDDCValidationOK creates a RetrySDDCValidationOK with default headers values
func NewRetrySDDCValidationOK() *RetrySDDCValidationOK {
	return &RetrySDDCValidationOK{}
}

/*
RetrySDDCValidationOK describes a response with status code 200, with default header values.

Completed
*/
type RetrySDDCValidationOK struct {
	Payload *models.Validation
}

// IsSuccess returns true when this retry Sddc validation o k response has a 2xx status code
func (o *RetrySDDCValidationOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this retry Sddc validation o k response has a 3xx status code
func (o *RetrySDDCValidationOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this retry Sddc validation o k response has a 4xx status code
func (o *RetrySDDCValidationOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this retry Sddc validation o k response has a 5xx status code
func (o *RetrySDDCValidationOK) IsServerError() bool {
	return false
}

// IsCode returns true when this retry Sddc validation o k response a status code equal to that given
func (o *RetrySDDCValidationOK) IsCode(code int) bool {
	return code == 200
}

func (o *RetrySDDCValidationOK) Error() string {
	return fmt.Sprintf("[PATCH /v1/sddcs/validations/{id}][%d] retrySddcValidationOK  %+v", 200, o.Payload)
}

func (o *RetrySDDCValidationOK) String() string {
	return fmt.Sprintf("[PATCH /v1/sddcs/validations/{id}][%d] retrySddcValidationOK  %+v", 200, o.Payload)
}

func (o *RetrySDDCValidationOK) GetPayload() *models.Validation {
	return o.Payload
}

func (o *RetrySDDCValidationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Validation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRetrySDDCValidationBadRequest creates a RetrySDDCValidationBadRequest with default headers values
func NewRetrySDDCValidationBadRequest() *RetrySDDCValidationBadRequest {
	return &RetrySDDCValidationBadRequest{}
}

/*
RetrySDDCValidationBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type RetrySDDCValidationBadRequest struct {
}

// IsSuccess returns true when this retry Sddc validation bad request response has a 2xx status code
func (o *RetrySDDCValidationBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this retry Sddc validation bad request response has a 3xx status code
func (o *RetrySDDCValidationBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this retry Sddc validation bad request response has a 4xx status code
func (o *RetrySDDCValidationBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this retry Sddc validation bad request response has a 5xx status code
func (o *RetrySDDCValidationBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this retry Sddc validation bad request response a status code equal to that given
func (o *RetrySDDCValidationBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *RetrySDDCValidationBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /v1/sddcs/validations/{id}][%d] retrySddcValidationBadRequest ", 400)
}

func (o *RetrySDDCValidationBadRequest) String() string {
	return fmt.Sprintf("[PATCH /v1/sddcs/validations/{id}][%d] retrySddcValidationBadRequest ", 400)
}

func (o *RetrySDDCValidationBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRetrySDDCValidationMethodNotAllowed creates a RetrySDDCValidationMethodNotAllowed with default headers values
func NewRetrySDDCValidationMethodNotAllowed() *RetrySDDCValidationMethodNotAllowed {
	return &RetrySDDCValidationMethodNotAllowed{}
}

/*
RetrySDDCValidationMethodNotAllowed describes a response with status code 405, with default header values.

Method Not Allowed
*/
type RetrySDDCValidationMethodNotAllowed struct {
}

// IsSuccess returns true when this retry Sddc validation method not allowed response has a 2xx status code
func (o *RetrySDDCValidationMethodNotAllowed) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this retry Sddc validation method not allowed response has a 3xx status code
func (o *RetrySDDCValidationMethodNotAllowed) IsRedirect() bool {
	return false
}

// IsClientError returns true when this retry Sddc validation method not allowed response has a 4xx status code
func (o *RetrySDDCValidationMethodNotAllowed) IsClientError() bool {
	return true
}

// IsServerError returns true when this retry Sddc validation method not allowed response has a 5xx status code
func (o *RetrySDDCValidationMethodNotAllowed) IsServerError() bool {
	return false
}

// IsCode returns true when this retry Sddc validation method not allowed response a status code equal to that given
func (o *RetrySDDCValidationMethodNotAllowed) IsCode(code int) bool {
	return code == 405
}

func (o *RetrySDDCValidationMethodNotAllowed) Error() string {
	return fmt.Sprintf("[PATCH /v1/sddcs/validations/{id}][%d] retrySddcValidationMethodNotAllowed ", 405)
}

func (o *RetrySDDCValidationMethodNotAllowed) String() string {
	return fmt.Sprintf("[PATCH /v1/sddcs/validations/{id}][%d] retrySddcValidationMethodNotAllowed ", 405)
}

func (o *RetrySDDCValidationMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRetrySDDCValidationInternalServerError creates a RetrySDDCValidationInternalServerError with default headers values
func NewRetrySDDCValidationInternalServerError() *RetrySDDCValidationInternalServerError {
	return &RetrySDDCValidationInternalServerError{}
}

/*
RetrySDDCValidationInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type RetrySDDCValidationInternalServerError struct {
}

// IsSuccess returns true when this retry Sddc validation internal server error response has a 2xx status code
func (o *RetrySDDCValidationInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this retry Sddc validation internal server error response has a 3xx status code
func (o *RetrySDDCValidationInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this retry Sddc validation internal server error response has a 4xx status code
func (o *RetrySDDCValidationInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this retry Sddc validation internal server error response has a 5xx status code
func (o *RetrySDDCValidationInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this retry Sddc validation internal server error response a status code equal to that given
func (o *RetrySDDCValidationInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *RetrySDDCValidationInternalServerError) Error() string {
	return fmt.Sprintf("[PATCH /v1/sddcs/validations/{id}][%d] retrySddcValidationInternalServerError ", 500)
}

func (o *RetrySDDCValidationInternalServerError) String() string {
	return fmt.Sprintf("[PATCH /v1/sddcs/validations/{id}][%d] retrySddcValidationInternalServerError ", 500)
}

func (o *RetrySDDCValidationInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
