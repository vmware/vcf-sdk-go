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

// GetSystemConfigurationReader is a Reader for the GetSystemConfiguration structure.
type GetSystemConfigurationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSystemConfigurationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetSystemConfigurationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetSystemConfigurationBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetSystemConfigurationInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /v1/system] getSystemConfiguration", response, response.Code())
	}
}

// NewGetSystemConfigurationOK creates a GetSystemConfigurationOK with default headers values
func NewGetSystemConfigurationOK() *GetSystemConfigurationOK {
	return &GetSystemConfigurationOK{}
}

/*
GetSystemConfigurationOK describes a response with status code 200, with default header values.

OK
*/
type GetSystemConfigurationOK struct {
	Payload *models.System
}

// IsSuccess returns true when this get system configuration o k response has a 2xx status code
func (o *GetSystemConfigurationOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get system configuration o k response has a 3xx status code
func (o *GetSystemConfigurationOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get system configuration o k response has a 4xx status code
func (o *GetSystemConfigurationOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get system configuration o k response has a 5xx status code
func (o *GetSystemConfigurationOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get system configuration o k response a status code equal to that given
func (o *GetSystemConfigurationOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get system configuration o k response
func (o *GetSystemConfigurationOK) Code() int {
	return 200
}

func (o *GetSystemConfigurationOK) Error() string {
	return fmt.Sprintf("[GET /v1/system][%d] getSystemConfigurationOK  %+v", 200, o.Payload)
}

func (o *GetSystemConfigurationOK) String() string {
	return fmt.Sprintf("[GET /v1/system][%d] getSystemConfigurationOK  %+v", 200, o.Payload)
}

func (o *GetSystemConfigurationOK) GetPayload() *models.System {
	return o.Payload
}

func (o *GetSystemConfigurationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.System)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSystemConfigurationBadRequest creates a GetSystemConfigurationBadRequest with default headers values
func NewGetSystemConfigurationBadRequest() *GetSystemConfigurationBadRequest {
	return &GetSystemConfigurationBadRequest{}
}

/*
GetSystemConfigurationBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetSystemConfigurationBadRequest struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this get system configuration bad request response has a 2xx status code
func (o *GetSystemConfigurationBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get system configuration bad request response has a 3xx status code
func (o *GetSystemConfigurationBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get system configuration bad request response has a 4xx status code
func (o *GetSystemConfigurationBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get system configuration bad request response has a 5xx status code
func (o *GetSystemConfigurationBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get system configuration bad request response a status code equal to that given
func (o *GetSystemConfigurationBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get system configuration bad request response
func (o *GetSystemConfigurationBadRequest) Code() int {
	return 400
}

func (o *GetSystemConfigurationBadRequest) Error() string {
	return fmt.Sprintf("[GET /v1/system][%d] getSystemConfigurationBadRequest  %+v", 400, o.Payload)
}

func (o *GetSystemConfigurationBadRequest) String() string {
	return fmt.Sprintf("[GET /v1/system][%d] getSystemConfigurationBadRequest  %+v", 400, o.Payload)
}

func (o *GetSystemConfigurationBadRequest) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetSystemConfigurationBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSystemConfigurationInternalServerError creates a GetSystemConfigurationInternalServerError with default headers values
func NewGetSystemConfigurationInternalServerError() *GetSystemConfigurationInternalServerError {
	return &GetSystemConfigurationInternalServerError{}
}

/*
GetSystemConfigurationInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetSystemConfigurationInternalServerError struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this get system configuration internal server error response has a 2xx status code
func (o *GetSystemConfigurationInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get system configuration internal server error response has a 3xx status code
func (o *GetSystemConfigurationInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get system configuration internal server error response has a 4xx status code
func (o *GetSystemConfigurationInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get system configuration internal server error response has a 5xx status code
func (o *GetSystemConfigurationInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get system configuration internal server error response a status code equal to that given
func (o *GetSystemConfigurationInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get system configuration internal server error response
func (o *GetSystemConfigurationInternalServerError) Code() int {
	return 500
}

func (o *GetSystemConfigurationInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/system][%d] getSystemConfigurationInternalServerError  %+v", 500, o.Payload)
}

func (o *GetSystemConfigurationInternalServerError) String() string {
	return fmt.Sprintf("[GET /v1/system][%d] getSystemConfigurationInternalServerError  %+v", 500, o.Payload)
}

func (o *GetSystemConfigurationInternalServerError) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetSystemConfigurationInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
