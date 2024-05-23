// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2023 VMware, Inc.
// SPDX-License-Identifier: BSD-2-Clause

package vcenters

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/vmware/vcf-sdk-go/models"
)

// GetVcenterReader is a Reader for the GetVcenter structure.
type GetVcenterReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetVcenterReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetVcenterOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetVcenterNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetVcenterInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /v1/vcenters/{id}] getVcenter", response, response.Code())
	}
}

// NewGetVcenterOK creates a GetVcenterOK with default headers values
func NewGetVcenterOK() *GetVcenterOK {
	return &GetVcenterOK{}
}

/*
GetVcenterOK describes a response with status code 200, with default header values.

Ok
*/
type GetVcenterOK struct {
	Payload *models.Vcenter
}

// IsSuccess returns true when this get vcenter o k response has a 2xx status code
func (o *GetVcenterOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get vcenter o k response has a 3xx status code
func (o *GetVcenterOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get vcenter o k response has a 4xx status code
func (o *GetVcenterOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get vcenter o k response has a 5xx status code
func (o *GetVcenterOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get vcenter o k response a status code equal to that given
func (o *GetVcenterOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get vcenter o k response
func (o *GetVcenterOK) Code() int {
	return 200
}

func (o *GetVcenterOK) Error() string {
	return fmt.Sprintf("[GET /v1/vcenters/{id}][%d] getVcenterOK  %+v", 200, o.Payload)
}

func (o *GetVcenterOK) String() string {
	return fmt.Sprintf("[GET /v1/vcenters/{id}][%d] getVcenterOK  %+v", 200, o.Payload)
}

func (o *GetVcenterOK) GetPayload() *models.Vcenter {
	return o.Payload
}

func (o *GetVcenterOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Vcenter)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetVcenterNotFound creates a GetVcenterNotFound with default headers values
func NewGetVcenterNotFound() *GetVcenterNotFound {
	return &GetVcenterNotFound{}
}

/*
GetVcenterNotFound describes a response with status code 404, with default header values.

vCenter not found
*/
type GetVcenterNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this get vcenter not found response has a 2xx status code
func (o *GetVcenterNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get vcenter not found response has a 3xx status code
func (o *GetVcenterNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get vcenter not found response has a 4xx status code
func (o *GetVcenterNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get vcenter not found response has a 5xx status code
func (o *GetVcenterNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get vcenter not found response a status code equal to that given
func (o *GetVcenterNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get vcenter not found response
func (o *GetVcenterNotFound) Code() int {
	return 404
}

func (o *GetVcenterNotFound) Error() string {
	return fmt.Sprintf("[GET /v1/vcenters/{id}][%d] getVcenterNotFound  %+v", 404, o.Payload)
}

func (o *GetVcenterNotFound) String() string {
	return fmt.Sprintf("[GET /v1/vcenters/{id}][%d] getVcenterNotFound  %+v", 404, o.Payload)
}

func (o *GetVcenterNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetVcenterNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetVcenterInternalServerError creates a GetVcenterInternalServerError with default headers values
func NewGetVcenterInternalServerError() *GetVcenterInternalServerError {
	return &GetVcenterInternalServerError{}
}

/*
GetVcenterInternalServerError describes a response with status code 500, with default header values.

InternalServerError
*/
type GetVcenterInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this get vcenter internal server error response has a 2xx status code
func (o *GetVcenterInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get vcenter internal server error response has a 3xx status code
func (o *GetVcenterInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get vcenter internal server error response has a 4xx status code
func (o *GetVcenterInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get vcenter internal server error response has a 5xx status code
func (o *GetVcenterInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get vcenter internal server error response a status code equal to that given
func (o *GetVcenterInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get vcenter internal server error response
func (o *GetVcenterInternalServerError) Code() int {
	return 500
}

func (o *GetVcenterInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/vcenters/{id}][%d] getVcenterInternalServerError  %+v", 500, o.Payload)
}

func (o *GetVcenterInternalServerError) String() string {
	return fmt.Sprintf("[GET /v1/vcenters/{id}][%d] getVcenterInternalServerError  %+v", 500, o.Payload)
}

func (o *GetVcenterInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetVcenterInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
