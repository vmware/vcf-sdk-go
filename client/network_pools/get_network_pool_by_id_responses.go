// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2023 VMware, Inc.
// SPDX-License-Identifier: BSD-2-Clause

package network_pools

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/vmware/vcf-sdk-go/models"
)

// GetNetworkPoolByIDReader is a Reader for the GetNetworkPoolByID structure.
type GetNetworkPoolByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetNetworkPoolByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetNetworkPoolByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetNetworkPoolByIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetNetworkPoolByIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /v1/network-pools/{id}] getNetworkPoolByID", response, response.Code())
	}
}

// NewGetNetworkPoolByIDOK creates a GetNetworkPoolByIDOK with default headers values
func NewGetNetworkPoolByIDOK() *GetNetworkPoolByIDOK {
	return &GetNetworkPoolByIDOK{}
}

/*
GetNetworkPoolByIDOK describes a response with status code 200, with default header values.

Referenced network pool
*/
type GetNetworkPoolByIDOK struct {
	Payload *models.NetworkPool
}

// IsSuccess returns true when this get network pool by Id o k response has a 2xx status code
func (o *GetNetworkPoolByIDOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get network pool by Id o k response has a 3xx status code
func (o *GetNetworkPoolByIDOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get network pool by Id o k response has a 4xx status code
func (o *GetNetworkPoolByIDOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get network pool by Id o k response has a 5xx status code
func (o *GetNetworkPoolByIDOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get network pool by Id o k response a status code equal to that given
func (o *GetNetworkPoolByIDOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get network pool by Id o k response
func (o *GetNetworkPoolByIDOK) Code() int {
	return 200
}

func (o *GetNetworkPoolByIDOK) Error() string {
	return fmt.Sprintf("[GET /v1/network-pools/{id}][%d] getNetworkPoolByIdOK  %+v", 200, o.Payload)
}

func (o *GetNetworkPoolByIDOK) String() string {
	return fmt.Sprintf("[GET /v1/network-pools/{id}][%d] getNetworkPoolByIdOK  %+v", 200, o.Payload)
}

func (o *GetNetworkPoolByIDOK) GetPayload() *models.NetworkPool {
	return o.Payload
}

func (o *GetNetworkPoolByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NetworkPool)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetNetworkPoolByIDNotFound creates a GetNetworkPoolByIDNotFound with default headers values
func NewGetNetworkPoolByIDNotFound() *GetNetworkPoolByIDNotFound {
	return &GetNetworkPoolByIDNotFound{}
}

/*
GetNetworkPoolByIDNotFound describes a response with status code 404, with default header values.

Referenced network pool not found
*/
type GetNetworkPoolByIDNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this get network pool by Id not found response has a 2xx status code
func (o *GetNetworkPoolByIDNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get network pool by Id not found response has a 3xx status code
func (o *GetNetworkPoolByIDNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get network pool by Id not found response has a 4xx status code
func (o *GetNetworkPoolByIDNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get network pool by Id not found response has a 5xx status code
func (o *GetNetworkPoolByIDNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get network pool by Id not found response a status code equal to that given
func (o *GetNetworkPoolByIDNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get network pool by Id not found response
func (o *GetNetworkPoolByIDNotFound) Code() int {
	return 404
}

func (o *GetNetworkPoolByIDNotFound) Error() string {
	return fmt.Sprintf("[GET /v1/network-pools/{id}][%d] getNetworkPoolByIdNotFound  %+v", 404, o.Payload)
}

func (o *GetNetworkPoolByIDNotFound) String() string {
	return fmt.Sprintf("[GET /v1/network-pools/{id}][%d] getNetworkPoolByIdNotFound  %+v", 404, o.Payload)
}

func (o *GetNetworkPoolByIDNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetNetworkPoolByIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetNetworkPoolByIDInternalServerError creates a GetNetworkPoolByIDInternalServerError with default headers values
func NewGetNetworkPoolByIDInternalServerError() *GetNetworkPoolByIDInternalServerError {
	return &GetNetworkPoolByIDInternalServerError{}
}

/*
GetNetworkPoolByIDInternalServerError describes a response with status code 500, with default header values.

Unexpected error
*/
type GetNetworkPoolByIDInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this get network pool by Id internal server error response has a 2xx status code
func (o *GetNetworkPoolByIDInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get network pool by Id internal server error response has a 3xx status code
func (o *GetNetworkPoolByIDInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get network pool by Id internal server error response has a 4xx status code
func (o *GetNetworkPoolByIDInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get network pool by Id internal server error response has a 5xx status code
func (o *GetNetworkPoolByIDInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get network pool by Id internal server error response a status code equal to that given
func (o *GetNetworkPoolByIDInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get network pool by Id internal server error response
func (o *GetNetworkPoolByIDInternalServerError) Code() int {
	return 500
}

func (o *GetNetworkPoolByIDInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/network-pools/{id}][%d] getNetworkPoolByIdInternalServerError  %+v", 500, o.Payload)
}

func (o *GetNetworkPoolByIDInternalServerError) String() string {
	return fmt.Sprintf("[GET /v1/network-pools/{id}][%d] getNetworkPoolByIdInternalServerError  %+v", 500, o.Payload)
}

func (o *GetNetworkPoolByIDInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetNetworkPoolByIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}