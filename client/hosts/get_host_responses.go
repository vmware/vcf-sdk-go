// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2023 VMware, Inc.
// SPDX-License-Identifier: BSD-2-Clause

package hosts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/vmware/vcf-sdk-go/models"
)

// GetHostReader is a Reader for the GetHost structure.
type GetHostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetHostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetHostOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetHostBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetHostNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetHostInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /v1/hosts/{id}] getHost", response, response.Code())
	}
}

// NewGetHostOK creates a GetHostOK with default headers values
func NewGetHostOK() *GetHostOK {
	return &GetHostOK{}
}

/*
GetHostOK describes a response with status code 200, with default header values.

Ok
*/
type GetHostOK struct {
	Payload *models.Host
}

// IsSuccess returns true when this get host o k response has a 2xx status code
func (o *GetHostOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get host o k response has a 3xx status code
func (o *GetHostOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get host o k response has a 4xx status code
func (o *GetHostOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get host o k response has a 5xx status code
func (o *GetHostOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get host o k response a status code equal to that given
func (o *GetHostOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get host o k response
func (o *GetHostOK) Code() int {
	return 200
}

func (o *GetHostOK) Error() string {
	return fmt.Sprintf("[GET /v1/hosts/{id}][%d] getHostOK  %+v", 200, o.Payload)
}

func (o *GetHostOK) String() string {
	return fmt.Sprintf("[GET /v1/hosts/{id}][%d] getHostOK  %+v", 200, o.Payload)
}

func (o *GetHostOK) GetPayload() *models.Host {
	return o.Payload
}

func (o *GetHostOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Host)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetHostBadRequest creates a GetHostBadRequest with default headers values
func NewGetHostBadRequest() *GetHostBadRequest {
	return &GetHostBadRequest{}
}

/*
GetHostBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetHostBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this get host bad request response has a 2xx status code
func (o *GetHostBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get host bad request response has a 3xx status code
func (o *GetHostBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get host bad request response has a 4xx status code
func (o *GetHostBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get host bad request response has a 5xx status code
func (o *GetHostBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get host bad request response a status code equal to that given
func (o *GetHostBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get host bad request response
func (o *GetHostBadRequest) Code() int {
	return 400
}

func (o *GetHostBadRequest) Error() string {
	return fmt.Sprintf("[GET /v1/hosts/{id}][%d] getHostBadRequest  %+v", 400, o.Payload)
}

func (o *GetHostBadRequest) String() string {
	return fmt.Sprintf("[GET /v1/hosts/{id}][%d] getHostBadRequest  %+v", 400, o.Payload)
}

func (o *GetHostBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetHostBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetHostNotFound creates a GetHostNotFound with default headers values
func NewGetHostNotFound() *GetHostNotFound {
	return &GetHostNotFound{}
}

/*
GetHostNotFound describes a response with status code 404, with default header values.

Host Not Found
*/
type GetHostNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this get host not found response has a 2xx status code
func (o *GetHostNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get host not found response has a 3xx status code
func (o *GetHostNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get host not found response has a 4xx status code
func (o *GetHostNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get host not found response has a 5xx status code
func (o *GetHostNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get host not found response a status code equal to that given
func (o *GetHostNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get host not found response
func (o *GetHostNotFound) Code() int {
	return 404
}

func (o *GetHostNotFound) Error() string {
	return fmt.Sprintf("[GET /v1/hosts/{id}][%d] getHostNotFound  %+v", 404, o.Payload)
}

func (o *GetHostNotFound) String() string {
	return fmt.Sprintf("[GET /v1/hosts/{id}][%d] getHostNotFound  %+v", 404, o.Payload)
}

func (o *GetHostNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetHostNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetHostInternalServerError creates a GetHostInternalServerError with default headers values
func NewGetHostInternalServerError() *GetHostInternalServerError {
	return &GetHostInternalServerError{}
}

/*
GetHostInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetHostInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this get host internal server error response has a 2xx status code
func (o *GetHostInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get host internal server error response has a 3xx status code
func (o *GetHostInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get host internal server error response has a 4xx status code
func (o *GetHostInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get host internal server error response has a 5xx status code
func (o *GetHostInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get host internal server error response a status code equal to that given
func (o *GetHostInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get host internal server error response
func (o *GetHostInternalServerError) Code() int {
	return 500
}

func (o *GetHostInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/hosts/{id}][%d] getHostInternalServerError  %+v", 500, o.Payload)
}

func (o *GetHostInternalServerError) String() string {
	return fmt.Sprintf("[GET /v1/hosts/{id}][%d] getHostInternalServerError  %+v", 500, o.Payload)
}

func (o *GetHostInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetHostInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
