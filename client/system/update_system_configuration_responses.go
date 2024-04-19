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

// UpdateSystemConfigurationReader is a Reader for the UpdateSystemConfiguration structure.
type UpdateSystemConfigurationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateSystemConfigurationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateSystemConfigurationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateSystemConfigurationBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUpdateSystemConfigurationInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[PATCH /v1/system] updateSystemConfiguration", response, response.Code())
	}
}

// NewUpdateSystemConfigurationOK creates a UpdateSystemConfigurationOK with default headers values
func NewUpdateSystemConfigurationOK() *UpdateSystemConfigurationOK {
	return &UpdateSystemConfigurationOK{}
}

/*
UpdateSystemConfigurationOK describes a response with status code 200, with default header values.

OK
*/
type UpdateSystemConfigurationOK struct {
}

// IsSuccess returns true when this update system configuration o k response has a 2xx status code
func (o *UpdateSystemConfigurationOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update system configuration o k response has a 3xx status code
func (o *UpdateSystemConfigurationOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update system configuration o k response has a 4xx status code
func (o *UpdateSystemConfigurationOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update system configuration o k response has a 5xx status code
func (o *UpdateSystemConfigurationOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update system configuration o k response a status code equal to that given
func (o *UpdateSystemConfigurationOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the update system configuration o k response
func (o *UpdateSystemConfigurationOK) Code() int {
	return 200
}

func (o *UpdateSystemConfigurationOK) Error() string {
	return fmt.Sprintf("[PATCH /v1/system][%d] updateSystemConfigurationOK ", 200)
}

func (o *UpdateSystemConfigurationOK) String() string {
	return fmt.Sprintf("[PATCH /v1/system][%d] updateSystemConfigurationOK ", 200)
}

func (o *UpdateSystemConfigurationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateSystemConfigurationBadRequest creates a UpdateSystemConfigurationBadRequest with default headers values
func NewUpdateSystemConfigurationBadRequest() *UpdateSystemConfigurationBadRequest {
	return &UpdateSystemConfigurationBadRequest{}
}

/*
UpdateSystemConfigurationBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type UpdateSystemConfigurationBadRequest struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this update system configuration bad request response has a 2xx status code
func (o *UpdateSystemConfigurationBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update system configuration bad request response has a 3xx status code
func (o *UpdateSystemConfigurationBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update system configuration bad request response has a 4xx status code
func (o *UpdateSystemConfigurationBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this update system configuration bad request response has a 5xx status code
func (o *UpdateSystemConfigurationBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this update system configuration bad request response a status code equal to that given
func (o *UpdateSystemConfigurationBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the update system configuration bad request response
func (o *UpdateSystemConfigurationBadRequest) Code() int {
	return 400
}

func (o *UpdateSystemConfigurationBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /v1/system][%d] updateSystemConfigurationBadRequest  %+v", 400, o.Payload)
}

func (o *UpdateSystemConfigurationBadRequest) String() string {
	return fmt.Sprintf("[PATCH /v1/system][%d] updateSystemConfigurationBadRequest  %+v", 400, o.Payload)
}

func (o *UpdateSystemConfigurationBadRequest) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *UpdateSystemConfigurationBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateSystemConfigurationInternalServerError creates a UpdateSystemConfigurationInternalServerError with default headers values
func NewUpdateSystemConfigurationInternalServerError() *UpdateSystemConfigurationInternalServerError {
	return &UpdateSystemConfigurationInternalServerError{}
}

/*
UpdateSystemConfigurationInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type UpdateSystemConfigurationInternalServerError struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this update system configuration internal server error response has a 2xx status code
func (o *UpdateSystemConfigurationInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update system configuration internal server error response has a 3xx status code
func (o *UpdateSystemConfigurationInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update system configuration internal server error response has a 4xx status code
func (o *UpdateSystemConfigurationInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this update system configuration internal server error response has a 5xx status code
func (o *UpdateSystemConfigurationInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this update system configuration internal server error response a status code equal to that given
func (o *UpdateSystemConfigurationInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the update system configuration internal server error response
func (o *UpdateSystemConfigurationInternalServerError) Code() int {
	return 500
}

func (o *UpdateSystemConfigurationInternalServerError) Error() string {
	return fmt.Sprintf("[PATCH /v1/system][%d] updateSystemConfigurationInternalServerError  %+v", 500, o.Payload)
}

func (o *UpdateSystemConfigurationInternalServerError) String() string {
	return fmt.Sprintf("[PATCH /v1/system][%d] updateSystemConfigurationInternalServerError  %+v", 500, o.Payload)
}

func (o *UpdateSystemConfigurationInternalServerError) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *UpdateSystemConfigurationInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
