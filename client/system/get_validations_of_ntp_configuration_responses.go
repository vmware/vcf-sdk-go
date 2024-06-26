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

// GetValidationsOfNtpConfigurationReader is a Reader for the GetValidationsOfNtpConfiguration structure.
type GetValidationsOfNtpConfigurationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetValidationsOfNtpConfigurationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetValidationsOfNtpConfigurationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetValidationsOfNtpConfigurationBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetValidationsOfNtpConfigurationInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /v1/system/ntp-configuration/validations] getValidationsOfNtpConfiguration", response, response.Code())
	}
}

// NewGetValidationsOfNtpConfigurationOK creates a GetValidationsOfNtpConfigurationOK with default headers values
func NewGetValidationsOfNtpConfigurationOK() *GetValidationsOfNtpConfigurationOK {
	return &GetValidationsOfNtpConfigurationOK{}
}

/*
GetValidationsOfNtpConfigurationOK describes a response with status code 200, with default header values.

OK
*/
type GetValidationsOfNtpConfigurationOK struct {
	Payload *models.Validation
}

// IsSuccess returns true when this get validations of ntp configuration o k response has a 2xx status code
func (o *GetValidationsOfNtpConfigurationOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get validations of ntp configuration o k response has a 3xx status code
func (o *GetValidationsOfNtpConfigurationOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get validations of ntp configuration o k response has a 4xx status code
func (o *GetValidationsOfNtpConfigurationOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get validations of ntp configuration o k response has a 5xx status code
func (o *GetValidationsOfNtpConfigurationOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get validations of ntp configuration o k response a status code equal to that given
func (o *GetValidationsOfNtpConfigurationOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get validations of ntp configuration o k response
func (o *GetValidationsOfNtpConfigurationOK) Code() int {
	return 200
}

func (o *GetValidationsOfNtpConfigurationOK) Error() string {
	return fmt.Sprintf("[GET /v1/system/ntp-configuration/validations][%d] getValidationsOfNtpConfigurationOK  %+v", 200, o.Payload)
}

func (o *GetValidationsOfNtpConfigurationOK) String() string {
	return fmt.Sprintf("[GET /v1/system/ntp-configuration/validations][%d] getValidationsOfNtpConfigurationOK  %+v", 200, o.Payload)
}

func (o *GetValidationsOfNtpConfigurationOK) GetPayload() *models.Validation {
	return o.Payload
}

func (o *GetValidationsOfNtpConfigurationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Validation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetValidationsOfNtpConfigurationBadRequest creates a GetValidationsOfNtpConfigurationBadRequest with default headers values
func NewGetValidationsOfNtpConfigurationBadRequest() *GetValidationsOfNtpConfigurationBadRequest {
	return &GetValidationsOfNtpConfigurationBadRequest{}
}

/*
GetValidationsOfNtpConfigurationBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetValidationsOfNtpConfigurationBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this get validations of ntp configuration bad request response has a 2xx status code
func (o *GetValidationsOfNtpConfigurationBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get validations of ntp configuration bad request response has a 3xx status code
func (o *GetValidationsOfNtpConfigurationBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get validations of ntp configuration bad request response has a 4xx status code
func (o *GetValidationsOfNtpConfigurationBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get validations of ntp configuration bad request response has a 5xx status code
func (o *GetValidationsOfNtpConfigurationBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get validations of ntp configuration bad request response a status code equal to that given
func (o *GetValidationsOfNtpConfigurationBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get validations of ntp configuration bad request response
func (o *GetValidationsOfNtpConfigurationBadRequest) Code() int {
	return 400
}

func (o *GetValidationsOfNtpConfigurationBadRequest) Error() string {
	return fmt.Sprintf("[GET /v1/system/ntp-configuration/validations][%d] getValidationsOfNtpConfigurationBadRequest  %+v", 400, o.Payload)
}

func (o *GetValidationsOfNtpConfigurationBadRequest) String() string {
	return fmt.Sprintf("[GET /v1/system/ntp-configuration/validations][%d] getValidationsOfNtpConfigurationBadRequest  %+v", 400, o.Payload)
}

func (o *GetValidationsOfNtpConfigurationBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetValidationsOfNtpConfigurationBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetValidationsOfNtpConfigurationInternalServerError creates a GetValidationsOfNtpConfigurationInternalServerError with default headers values
func NewGetValidationsOfNtpConfigurationInternalServerError() *GetValidationsOfNtpConfigurationInternalServerError {
	return &GetValidationsOfNtpConfigurationInternalServerError{}
}

/*
GetValidationsOfNtpConfigurationInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type GetValidationsOfNtpConfigurationInternalServerError struct {
}

// IsSuccess returns true when this get validations of ntp configuration internal server error response has a 2xx status code
func (o *GetValidationsOfNtpConfigurationInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get validations of ntp configuration internal server error response has a 3xx status code
func (o *GetValidationsOfNtpConfigurationInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get validations of ntp configuration internal server error response has a 4xx status code
func (o *GetValidationsOfNtpConfigurationInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get validations of ntp configuration internal server error response has a 5xx status code
func (o *GetValidationsOfNtpConfigurationInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get validations of ntp configuration internal server error response a status code equal to that given
func (o *GetValidationsOfNtpConfigurationInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get validations of ntp configuration internal server error response
func (o *GetValidationsOfNtpConfigurationInternalServerError) Code() int {
	return 500
}

func (o *GetValidationsOfNtpConfigurationInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/system/ntp-configuration/validations][%d] getValidationsOfNtpConfigurationInternalServerError ", 500)
}

func (o *GetValidationsOfNtpConfigurationInternalServerError) String() string {
	return fmt.Sprintf("[GET /v1/system/ntp-configuration/validations][%d] getValidationsOfNtpConfigurationInternalServerError ", 500)
}

func (o *GetValidationsOfNtpConfigurationInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
