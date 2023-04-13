// Code generated by go-swagger; DO NOT EDIT.

// Copyright © 2023 VMware, Inc. All Rights Reserved.
// Code licenced under: BSD-2

package system

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"vcf-sdk-go/models"
)

// GETValidationsOfNtpConfigurationReader is a Reader for the GETValidationsOfNtpConfiguration structure.
type GETValidationsOfNtpConfigurationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GETValidationsOfNtpConfigurationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGETValidationsOfNtpConfigurationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGETValidationsOfNtpConfigurationBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGETValidationsOfNtpConfigurationInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGETValidationsOfNtpConfigurationOK creates a GETValidationsOfNtpConfigurationOK with default headers values
func NewGETValidationsOfNtpConfigurationOK() *GETValidationsOfNtpConfigurationOK {
	return &GETValidationsOfNtpConfigurationOK{}
}

/*
GETValidationsOfNtpConfigurationOK describes a response with status code 200, with default header values.

OK
*/
type GETValidationsOfNtpConfigurationOK struct {
	Payload *models.Validation
}

// IsSuccess returns true when this get validations of ntp configuration o k response has a 2xx status code
func (o *GETValidationsOfNtpConfigurationOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get validations of ntp configuration o k response has a 3xx status code
func (o *GETValidationsOfNtpConfigurationOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get validations of ntp configuration o k response has a 4xx status code
func (o *GETValidationsOfNtpConfigurationOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get validations of ntp configuration o k response has a 5xx status code
func (o *GETValidationsOfNtpConfigurationOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get validations of ntp configuration o k response a status code equal to that given
func (o *GETValidationsOfNtpConfigurationOK) IsCode(code int) bool {
	return code == 200
}

func (o *GETValidationsOfNtpConfigurationOK) Error() string {
	return fmt.Sprintf("[GET /v1/system/ntp-configuration/validations][%d] getValidationsOfNtpConfigurationOK  %+v", 200, o.Payload)
}

func (o *GETValidationsOfNtpConfigurationOK) String() string {
	return fmt.Sprintf("[GET /v1/system/ntp-configuration/validations][%d] getValidationsOfNtpConfigurationOK  %+v", 200, o.Payload)
}

func (o *GETValidationsOfNtpConfigurationOK) GetPayload() *models.Validation {
	return o.Payload
}

func (o *GETValidationsOfNtpConfigurationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Validation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGETValidationsOfNtpConfigurationBadRequest creates a GETValidationsOfNtpConfigurationBadRequest with default headers values
func NewGETValidationsOfNtpConfigurationBadRequest() *GETValidationsOfNtpConfigurationBadRequest {
	return &GETValidationsOfNtpConfigurationBadRequest{}
}

/*
GETValidationsOfNtpConfigurationBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GETValidationsOfNtpConfigurationBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this get validations of ntp configuration bad request response has a 2xx status code
func (o *GETValidationsOfNtpConfigurationBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get validations of ntp configuration bad request response has a 3xx status code
func (o *GETValidationsOfNtpConfigurationBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get validations of ntp configuration bad request response has a 4xx status code
func (o *GETValidationsOfNtpConfigurationBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get validations of ntp configuration bad request response has a 5xx status code
func (o *GETValidationsOfNtpConfigurationBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get validations of ntp configuration bad request response a status code equal to that given
func (o *GETValidationsOfNtpConfigurationBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GETValidationsOfNtpConfigurationBadRequest) Error() string {
	return fmt.Sprintf("[GET /v1/system/ntp-configuration/validations][%d] getValidationsOfNtpConfigurationBadRequest  %+v", 400, o.Payload)
}

func (o *GETValidationsOfNtpConfigurationBadRequest) String() string {
	return fmt.Sprintf("[GET /v1/system/ntp-configuration/validations][%d] getValidationsOfNtpConfigurationBadRequest  %+v", 400, o.Payload)
}

func (o *GETValidationsOfNtpConfigurationBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *GETValidationsOfNtpConfigurationBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGETValidationsOfNtpConfigurationInternalServerError creates a GETValidationsOfNtpConfigurationInternalServerError with default headers values
func NewGETValidationsOfNtpConfigurationInternalServerError() *GETValidationsOfNtpConfigurationInternalServerError {
	return &GETValidationsOfNtpConfigurationInternalServerError{}
}

/*
GETValidationsOfNtpConfigurationInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type GETValidationsOfNtpConfigurationInternalServerError struct {
}

// IsSuccess returns true when this get validations of ntp configuration internal server error response has a 2xx status code
func (o *GETValidationsOfNtpConfigurationInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get validations of ntp configuration internal server error response has a 3xx status code
func (o *GETValidationsOfNtpConfigurationInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get validations of ntp configuration internal server error response has a 4xx status code
func (o *GETValidationsOfNtpConfigurationInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get validations of ntp configuration internal server error response has a 5xx status code
func (o *GETValidationsOfNtpConfigurationInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get validations of ntp configuration internal server error response a status code equal to that given
func (o *GETValidationsOfNtpConfigurationInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GETValidationsOfNtpConfigurationInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/system/ntp-configuration/validations][%d] getValidationsOfNtpConfigurationInternalServerError ", 500)
}

func (o *GETValidationsOfNtpConfigurationInternalServerError) String() string {
	return fmt.Sprintf("[GET /v1/system/ntp-configuration/validations][%d] getValidationsOfNtpConfigurationInternalServerError ", 500)
}

func (o *GETValidationsOfNtpConfigurationInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
