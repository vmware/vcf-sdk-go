// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2023 VMware, Inc.
// SPDX-License-Identifier: BSD-2-Clause

package vasa_providers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/vmware/vcf-sdk-go/models"
)

// ValidateVasaProviderSpecReader is a Reader for the ValidateVasaProviderSpec structure.
type ValidateVasaProviderSpecReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ValidateVasaProviderSpecReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewValidateVasaProviderSpecOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewValidateVasaProviderSpecAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewValidateVasaProviderSpecBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewValidateVasaProviderSpecInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /v1/vasa-providers/validations] validateVasaProviderSpec", response, response.Code())
	}
}

// NewValidateVasaProviderSpecOK creates a ValidateVasaProviderSpecOK with default headers values
func NewValidateVasaProviderSpecOK() *ValidateVasaProviderSpecOK {
	return &ValidateVasaProviderSpecOK{}
}

/*
ValidateVasaProviderSpecOK describes a response with status code 200, with default header values.

OK
*/
type ValidateVasaProviderSpecOK struct {
	Payload *models.Validation
}

// IsSuccess returns true when this validate vasa provider spec o k response has a 2xx status code
func (o *ValidateVasaProviderSpecOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this validate vasa provider spec o k response has a 3xx status code
func (o *ValidateVasaProviderSpecOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this validate vasa provider spec o k response has a 4xx status code
func (o *ValidateVasaProviderSpecOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this validate vasa provider spec o k response has a 5xx status code
func (o *ValidateVasaProviderSpecOK) IsServerError() bool {
	return false
}

// IsCode returns true when this validate vasa provider spec o k response a status code equal to that given
func (o *ValidateVasaProviderSpecOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the validate vasa provider spec o k response
func (o *ValidateVasaProviderSpecOK) Code() int {
	return 200
}

func (o *ValidateVasaProviderSpecOK) Error() string {
	return fmt.Sprintf("[POST /v1/vasa-providers/validations][%d] validateVasaProviderSpecOK  %+v", 200, o.Payload)
}

func (o *ValidateVasaProviderSpecOK) String() string {
	return fmt.Sprintf("[POST /v1/vasa-providers/validations][%d] validateVasaProviderSpecOK  %+v", 200, o.Payload)
}

func (o *ValidateVasaProviderSpecOK) GetPayload() *models.Validation {
	return o.Payload
}

func (o *ValidateVasaProviderSpecOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Validation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewValidateVasaProviderSpecAccepted creates a ValidateVasaProviderSpecAccepted with default headers values
func NewValidateVasaProviderSpecAccepted() *ValidateVasaProviderSpecAccepted {
	return &ValidateVasaProviderSpecAccepted{}
}

/*
ValidateVasaProviderSpecAccepted describes a response with status code 202, with default header values.

Accepted
*/
type ValidateVasaProviderSpecAccepted struct {
	Payload *models.Validation
}

// IsSuccess returns true when this validate vasa provider spec accepted response has a 2xx status code
func (o *ValidateVasaProviderSpecAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this validate vasa provider spec accepted response has a 3xx status code
func (o *ValidateVasaProviderSpecAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this validate vasa provider spec accepted response has a 4xx status code
func (o *ValidateVasaProviderSpecAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this validate vasa provider spec accepted response has a 5xx status code
func (o *ValidateVasaProviderSpecAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this validate vasa provider spec accepted response a status code equal to that given
func (o *ValidateVasaProviderSpecAccepted) IsCode(code int) bool {
	return code == 202
}

// Code gets the status code for the validate vasa provider spec accepted response
func (o *ValidateVasaProviderSpecAccepted) Code() int {
	return 202
}

func (o *ValidateVasaProviderSpecAccepted) Error() string {
	return fmt.Sprintf("[POST /v1/vasa-providers/validations][%d] validateVasaProviderSpecAccepted  %+v", 202, o.Payload)
}

func (o *ValidateVasaProviderSpecAccepted) String() string {
	return fmt.Sprintf("[POST /v1/vasa-providers/validations][%d] validateVasaProviderSpecAccepted  %+v", 202, o.Payload)
}

func (o *ValidateVasaProviderSpecAccepted) GetPayload() *models.Validation {
	return o.Payload
}

func (o *ValidateVasaProviderSpecAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Validation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewValidateVasaProviderSpecBadRequest creates a ValidateVasaProviderSpecBadRequest with default headers values
func NewValidateVasaProviderSpecBadRequest() *ValidateVasaProviderSpecBadRequest {
	return &ValidateVasaProviderSpecBadRequest{}
}

/*
ValidateVasaProviderSpecBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type ValidateVasaProviderSpecBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this validate vasa provider spec bad request response has a 2xx status code
func (o *ValidateVasaProviderSpecBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this validate vasa provider spec bad request response has a 3xx status code
func (o *ValidateVasaProviderSpecBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this validate vasa provider spec bad request response has a 4xx status code
func (o *ValidateVasaProviderSpecBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this validate vasa provider spec bad request response has a 5xx status code
func (o *ValidateVasaProviderSpecBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this validate vasa provider spec bad request response a status code equal to that given
func (o *ValidateVasaProviderSpecBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the validate vasa provider spec bad request response
func (o *ValidateVasaProviderSpecBadRequest) Code() int {
	return 400
}

func (o *ValidateVasaProviderSpecBadRequest) Error() string {
	return fmt.Sprintf("[POST /v1/vasa-providers/validations][%d] validateVasaProviderSpecBadRequest  %+v", 400, o.Payload)
}

func (o *ValidateVasaProviderSpecBadRequest) String() string {
	return fmt.Sprintf("[POST /v1/vasa-providers/validations][%d] validateVasaProviderSpecBadRequest  %+v", 400, o.Payload)
}

func (o *ValidateVasaProviderSpecBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *ValidateVasaProviderSpecBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewValidateVasaProviderSpecInternalServerError creates a ValidateVasaProviderSpecInternalServerError with default headers values
func NewValidateVasaProviderSpecInternalServerError() *ValidateVasaProviderSpecInternalServerError {
	return &ValidateVasaProviderSpecInternalServerError{}
}

/*
ValidateVasaProviderSpecInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type ValidateVasaProviderSpecInternalServerError struct {
}

// IsSuccess returns true when this validate vasa provider spec internal server error response has a 2xx status code
func (o *ValidateVasaProviderSpecInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this validate vasa provider spec internal server error response has a 3xx status code
func (o *ValidateVasaProviderSpecInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this validate vasa provider spec internal server error response has a 4xx status code
func (o *ValidateVasaProviderSpecInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this validate vasa provider spec internal server error response has a 5xx status code
func (o *ValidateVasaProviderSpecInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this validate vasa provider spec internal server error response a status code equal to that given
func (o *ValidateVasaProviderSpecInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the validate vasa provider spec internal server error response
func (o *ValidateVasaProviderSpecInternalServerError) Code() int {
	return 500
}

func (o *ValidateVasaProviderSpecInternalServerError) Error() string {
	return fmt.Sprintf("[POST /v1/vasa-providers/validations][%d] validateVasaProviderSpecInternalServerError ", 500)
}

func (o *ValidateVasaProviderSpecInternalServerError) String() string {
	return fmt.Sprintf("[POST /v1/vasa-providers/validations][%d] validateVasaProviderSpecInternalServerError ", 500)
}

func (o *ValidateVasaProviderSpecInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}