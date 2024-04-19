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

// GetValidationOfDNSConfigurationReader is a Reader for the GetValidationOfDNSConfiguration structure.
type GetValidationOfDNSConfigurationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetValidationOfDNSConfigurationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetValidationOfDNSConfigurationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetValidationOfDNSConfigurationBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetValidationOfDNSConfigurationInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /v1/system/dns-configuration/validations/{id}] getValidationOfDnsConfiguration", response, response.Code())
	}
}

// NewGetValidationOfDNSConfigurationOK creates a GetValidationOfDNSConfigurationOK with default headers values
func NewGetValidationOfDNSConfigurationOK() *GetValidationOfDNSConfigurationOK {
	return &GetValidationOfDNSConfigurationOK{}
}

/*
GetValidationOfDNSConfigurationOK describes a response with status code 200, with default header values.

OK
*/
type GetValidationOfDNSConfigurationOK struct {
	Payload *models.Validation
}

// IsSuccess returns true when this get validation of Dns configuration o k response has a 2xx status code
func (o *GetValidationOfDNSConfigurationOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get validation of Dns configuration o k response has a 3xx status code
func (o *GetValidationOfDNSConfigurationOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get validation of Dns configuration o k response has a 4xx status code
func (o *GetValidationOfDNSConfigurationOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get validation of Dns configuration o k response has a 5xx status code
func (o *GetValidationOfDNSConfigurationOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get validation of Dns configuration o k response a status code equal to that given
func (o *GetValidationOfDNSConfigurationOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get validation of Dns configuration o k response
func (o *GetValidationOfDNSConfigurationOK) Code() int {
	return 200
}

func (o *GetValidationOfDNSConfigurationOK) Error() string {
	return fmt.Sprintf("[GET /v1/system/dns-configuration/validations/{id}][%d] getValidationOfDnsConfigurationOK  %+v", 200, o.Payload)
}

func (o *GetValidationOfDNSConfigurationOK) String() string {
	return fmt.Sprintf("[GET /v1/system/dns-configuration/validations/{id}][%d] getValidationOfDnsConfigurationOK  %+v", 200, o.Payload)
}

func (o *GetValidationOfDNSConfigurationOK) GetPayload() *models.Validation {
	return o.Payload
}

func (o *GetValidationOfDNSConfigurationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Validation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetValidationOfDNSConfigurationBadRequest creates a GetValidationOfDNSConfigurationBadRequest with default headers values
func NewGetValidationOfDNSConfigurationBadRequest() *GetValidationOfDNSConfigurationBadRequest {
	return &GetValidationOfDNSConfigurationBadRequest{}
}

/*
GetValidationOfDNSConfigurationBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetValidationOfDNSConfigurationBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this get validation of Dns configuration bad request response has a 2xx status code
func (o *GetValidationOfDNSConfigurationBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get validation of Dns configuration bad request response has a 3xx status code
func (o *GetValidationOfDNSConfigurationBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get validation of Dns configuration bad request response has a 4xx status code
func (o *GetValidationOfDNSConfigurationBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get validation of Dns configuration bad request response has a 5xx status code
func (o *GetValidationOfDNSConfigurationBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get validation of Dns configuration bad request response a status code equal to that given
func (o *GetValidationOfDNSConfigurationBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get validation of Dns configuration bad request response
func (o *GetValidationOfDNSConfigurationBadRequest) Code() int {
	return 400
}

func (o *GetValidationOfDNSConfigurationBadRequest) Error() string {
	return fmt.Sprintf("[GET /v1/system/dns-configuration/validations/{id}][%d] getValidationOfDnsConfigurationBadRequest  %+v", 400, o.Payload)
}

func (o *GetValidationOfDNSConfigurationBadRequest) String() string {
	return fmt.Sprintf("[GET /v1/system/dns-configuration/validations/{id}][%d] getValidationOfDnsConfigurationBadRequest  %+v", 400, o.Payload)
}

func (o *GetValidationOfDNSConfigurationBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetValidationOfDNSConfigurationBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetValidationOfDNSConfigurationInternalServerError creates a GetValidationOfDNSConfigurationInternalServerError with default headers values
func NewGetValidationOfDNSConfigurationInternalServerError() *GetValidationOfDNSConfigurationInternalServerError {
	return &GetValidationOfDNSConfigurationInternalServerError{}
}

/*
GetValidationOfDNSConfigurationInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type GetValidationOfDNSConfigurationInternalServerError struct {
}

// IsSuccess returns true when this get validation of Dns configuration internal server error response has a 2xx status code
func (o *GetValidationOfDNSConfigurationInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get validation of Dns configuration internal server error response has a 3xx status code
func (o *GetValidationOfDNSConfigurationInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get validation of Dns configuration internal server error response has a 4xx status code
func (o *GetValidationOfDNSConfigurationInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get validation of Dns configuration internal server error response has a 5xx status code
func (o *GetValidationOfDNSConfigurationInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get validation of Dns configuration internal server error response a status code equal to that given
func (o *GetValidationOfDNSConfigurationInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get validation of Dns configuration internal server error response
func (o *GetValidationOfDNSConfigurationInternalServerError) Code() int {
	return 500
}

func (o *GetValidationOfDNSConfigurationInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/system/dns-configuration/validations/{id}][%d] getValidationOfDnsConfigurationInternalServerError ", 500)
}

func (o *GetValidationOfDNSConfigurationInternalServerError) String() string {
	return fmt.Sprintf("[GET /v1/system/dns-configuration/validations/{id}][%d] getValidationOfDnsConfigurationInternalServerError ", 500)
}

func (o *GetValidationOfDNSConfigurationInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
