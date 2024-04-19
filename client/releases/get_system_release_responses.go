// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2023 VMware, Inc.
// SPDX-License-Identifier: BSD-2-Clause

package releases

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/vmware/vcf-sdk-go/models"
)

// GetSystemReleaseReader is a Reader for the GetSystemRelease structure.
type GetSystemReleaseReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSystemReleaseReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetSystemReleaseOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetSystemReleaseNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetSystemReleaseInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /v1/releases/system] getSystemRelease", response, response.Code())
	}
}

// NewGetSystemReleaseOK creates a GetSystemReleaseOK with default headers values
func NewGetSystemReleaseOK() *GetSystemReleaseOK {
	return &GetSystemReleaseOK{}
}

/*
GetSystemReleaseOK describes a response with status code 200, with default header values.

Ok
*/
type GetSystemReleaseOK struct {
	Payload *models.Release
}

// IsSuccess returns true when this get system release o k response has a 2xx status code
func (o *GetSystemReleaseOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get system release o k response has a 3xx status code
func (o *GetSystemReleaseOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get system release o k response has a 4xx status code
func (o *GetSystemReleaseOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get system release o k response has a 5xx status code
func (o *GetSystemReleaseOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get system release o k response a status code equal to that given
func (o *GetSystemReleaseOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get system release o k response
func (o *GetSystemReleaseOK) Code() int {
	return 200
}

func (o *GetSystemReleaseOK) Error() string {
	return fmt.Sprintf("[GET /v1/releases/system][%d] getSystemReleaseOK  %+v", 200, o.Payload)
}

func (o *GetSystemReleaseOK) String() string {
	return fmt.Sprintf("[GET /v1/releases/system][%d] getSystemReleaseOK  %+v", 200, o.Payload)
}

func (o *GetSystemReleaseOK) GetPayload() *models.Release {
	return o.Payload
}

func (o *GetSystemReleaseOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Release)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSystemReleaseNotFound creates a GetSystemReleaseNotFound with default headers values
func NewGetSystemReleaseNotFound() *GetSystemReleaseNotFound {
	return &GetSystemReleaseNotFound{}
}

/*
GetSystemReleaseNotFound describes a response with status code 404, with default header values.

Release not found
*/
type GetSystemReleaseNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this get system release not found response has a 2xx status code
func (o *GetSystemReleaseNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get system release not found response has a 3xx status code
func (o *GetSystemReleaseNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get system release not found response has a 4xx status code
func (o *GetSystemReleaseNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get system release not found response has a 5xx status code
func (o *GetSystemReleaseNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get system release not found response a status code equal to that given
func (o *GetSystemReleaseNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get system release not found response
func (o *GetSystemReleaseNotFound) Code() int {
	return 404
}

func (o *GetSystemReleaseNotFound) Error() string {
	return fmt.Sprintf("[GET /v1/releases/system][%d] getSystemReleaseNotFound  %+v", 404, o.Payload)
}

func (o *GetSystemReleaseNotFound) String() string {
	return fmt.Sprintf("[GET /v1/releases/system][%d] getSystemReleaseNotFound  %+v", 404, o.Payload)
}

func (o *GetSystemReleaseNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetSystemReleaseNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSystemReleaseInternalServerError creates a GetSystemReleaseInternalServerError with default headers values
func NewGetSystemReleaseInternalServerError() *GetSystemReleaseInternalServerError {
	return &GetSystemReleaseInternalServerError{}
}

/*
GetSystemReleaseInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetSystemReleaseInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this get system release internal server error response has a 2xx status code
func (o *GetSystemReleaseInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get system release internal server error response has a 3xx status code
func (o *GetSystemReleaseInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get system release internal server error response has a 4xx status code
func (o *GetSystemReleaseInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get system release internal server error response has a 5xx status code
func (o *GetSystemReleaseInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get system release internal server error response a status code equal to that given
func (o *GetSystemReleaseInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get system release internal server error response
func (o *GetSystemReleaseInternalServerError) Code() int {
	return 500
}

func (o *GetSystemReleaseInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/releases/system][%d] getSystemReleaseInternalServerError  %+v", 500, o.Payload)
}

func (o *GetSystemReleaseInternalServerError) String() string {
	return fmt.Sprintf("[GET /v1/releases/system][%d] getSystemReleaseInternalServerError  %+v", 500, o.Payload)
}

func (o *GetSystemReleaseInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetSystemReleaseInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
