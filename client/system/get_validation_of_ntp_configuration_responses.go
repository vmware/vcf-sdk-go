// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2023 VMware, Inc.
// SPDX-License-Identifier: BSD-2-Clause

package system

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/vmware/vcf-sdk-go/models"
)

// GetValidationOfNtpConfigurationReader is a Reader for the GetValidationOfNtpConfiguration structure.
type GetValidationOfNtpConfigurationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetValidationOfNtpConfigurationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetValidationOfNtpConfigurationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetValidationOfNtpConfigurationBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetValidationOfNtpConfigurationInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /v1/system/ntp-configuration/validations/{id}] getValidationOfNtpConfiguration", response, response.Code())
	}
}

// NewGetValidationOfNtpConfigurationOK creates a GetValidationOfNtpConfigurationOK with default headers values
func NewGetValidationOfNtpConfigurationOK() *GetValidationOfNtpConfigurationOK {
	return &GetValidationOfNtpConfigurationOK{}
}

/*
GetValidationOfNtpConfigurationOK describes a response with status code 200, with default header values.

OK
*/
type GetValidationOfNtpConfigurationOK struct {
	Payload *models.Validation
}

// IsSuccess returns true when this get validation of ntp configuration o k response has a 2xx status code
func (o *GetValidationOfNtpConfigurationOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get validation of ntp configuration o k response has a 3xx status code
func (o *GetValidationOfNtpConfigurationOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get validation of ntp configuration o k response has a 4xx status code
func (o *GetValidationOfNtpConfigurationOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get validation of ntp configuration o k response has a 5xx status code
func (o *GetValidationOfNtpConfigurationOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get validation of ntp configuration o k response a status code equal to that given
func (o *GetValidationOfNtpConfigurationOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get validation of ntp configuration o k response
func (o *GetValidationOfNtpConfigurationOK) Code() int {
	return 200
}

func (o *GetValidationOfNtpConfigurationOK) Error() string {
	return fmt.Sprintf("[GET /v1/system/ntp-configuration/validations/{id}][%d] getValidationOfNtpConfigurationOK  %+v", 200, o.Payload)
}

func (o *GetValidationOfNtpConfigurationOK) String() string {
	return fmt.Sprintf("[GET /v1/system/ntp-configuration/validations/{id}][%d] getValidationOfNtpConfigurationOK  %+v", 200, o.Payload)
}

func (o *GetValidationOfNtpConfigurationOK) GetPayload() *models.Validation {
	return o.Payload
}

func (o *GetValidationOfNtpConfigurationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Validation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetValidationOfNtpConfigurationBadRequest creates a GetValidationOfNtpConfigurationBadRequest with default headers values
func NewGetValidationOfNtpConfigurationBadRequest() *GetValidationOfNtpConfigurationBadRequest {
	return &GetValidationOfNtpConfigurationBadRequest{}
}

/*
GetValidationOfNtpConfigurationBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetValidationOfNtpConfigurationBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this get validation of ntp configuration bad request response has a 2xx status code
func (o *GetValidationOfNtpConfigurationBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get validation of ntp configuration bad request response has a 3xx status code
func (o *GetValidationOfNtpConfigurationBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get validation of ntp configuration bad request response has a 4xx status code
func (o *GetValidationOfNtpConfigurationBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get validation of ntp configuration bad request response has a 5xx status code
func (o *GetValidationOfNtpConfigurationBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get validation of ntp configuration bad request response a status code equal to that given
func (o *GetValidationOfNtpConfigurationBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get validation of ntp configuration bad request response
func (o *GetValidationOfNtpConfigurationBadRequest) Code() int {
	return 400
}

func (o *GetValidationOfNtpConfigurationBadRequest) Error() string {
	return fmt.Sprintf("[GET /v1/system/ntp-configuration/validations/{id}][%d] getValidationOfNtpConfigurationBadRequest  %+v", 400, o.Payload)
}

func (o *GetValidationOfNtpConfigurationBadRequest) String() string {
	return fmt.Sprintf("[GET /v1/system/ntp-configuration/validations/{id}][%d] getValidationOfNtpConfigurationBadRequest  %+v", 400, o.Payload)
}

func (o *GetValidationOfNtpConfigurationBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetValidationOfNtpConfigurationBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetValidationOfNtpConfigurationInternalServerError creates a GetValidationOfNtpConfigurationInternalServerError with default headers values
func NewGetValidationOfNtpConfigurationInternalServerError() *GetValidationOfNtpConfigurationInternalServerError {
	return &GetValidationOfNtpConfigurationInternalServerError{}
}

/*
GetValidationOfNtpConfigurationInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type GetValidationOfNtpConfigurationInternalServerError struct {
}

// IsSuccess returns true when this get validation of ntp configuration internal server error response has a 2xx status code
func (o *GetValidationOfNtpConfigurationInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get validation of ntp configuration internal server error response has a 3xx status code
func (o *GetValidationOfNtpConfigurationInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get validation of ntp configuration internal server error response has a 4xx status code
func (o *GetValidationOfNtpConfigurationInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get validation of ntp configuration internal server error response has a 5xx status code
func (o *GetValidationOfNtpConfigurationInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get validation of ntp configuration internal server error response a status code equal to that given
func (o *GetValidationOfNtpConfigurationInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get validation of ntp configuration internal server error response
func (o *GetValidationOfNtpConfigurationInternalServerError) Code() int {
	return 500
}

func (o *GetValidationOfNtpConfigurationInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/system/ntp-configuration/validations/{id}][%d] getValidationOfNtpConfigurationInternalServerError ", 500)
}

func (o *GetValidationOfNtpConfigurationInternalServerError) String() string {
	return fmt.Sprintf("[GET /v1/system/ntp-configuration/validations/{id}][%d] getValidationOfNtpConfigurationInternalServerError ", 500)
}

func (o *GetValidationOfNtpConfigurationInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
