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

// GETNetworkPoolReader is a Reader for the GETNetworkPool structure.
type GETNetworkPoolReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GETNetworkPoolReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGETNetworkPoolOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGETNetworkPoolNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGETNetworkPoolInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGETNetworkPoolOK creates a GETNetworkPoolOK with default headers values
func NewGETNetworkPoolOK() *GETNetworkPoolOK {
	return &GETNetworkPoolOK{}
}

/*
GETNetworkPoolOK describes a response with status code 200, with default header values.

Referenced network pool
*/
type GETNetworkPoolOK struct {
	Payload *models.NetworkPool
}

// IsSuccess returns true when this get network pool o k response has a 2xx status code
func (o *GETNetworkPoolOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get network pool o k response has a 3xx status code
func (o *GETNetworkPoolOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get network pool o k response has a 4xx status code
func (o *GETNetworkPoolOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get network pool o k response has a 5xx status code
func (o *GETNetworkPoolOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get network pool o k response a status code equal to that given
func (o *GETNetworkPoolOK) IsCode(code int) bool {
	return code == 200
}

func (o *GETNetworkPoolOK) Error() string {
	return fmt.Sprintf("[GET /v1/network-pools/{id}][%d] getNetworkPoolOK  %+v", 200, o.Payload)
}

func (o *GETNetworkPoolOK) String() string {
	return fmt.Sprintf("[GET /v1/network-pools/{id}][%d] getNetworkPoolOK  %+v", 200, o.Payload)
}

func (o *GETNetworkPoolOK) GetPayload() *models.NetworkPool {
	return o.Payload
}

func (o *GETNetworkPoolOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NetworkPool)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGETNetworkPoolNotFound creates a GETNetworkPoolNotFound with default headers values
func NewGETNetworkPoolNotFound() *GETNetworkPoolNotFound {
	return &GETNetworkPoolNotFound{}
}

/*
GETNetworkPoolNotFound describes a response with status code 404, with default header values.

Referenced network pool not found
*/
type GETNetworkPoolNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this get network pool not found response has a 2xx status code
func (o *GETNetworkPoolNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get network pool not found response has a 3xx status code
func (o *GETNetworkPoolNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get network pool not found response has a 4xx status code
func (o *GETNetworkPoolNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get network pool not found response has a 5xx status code
func (o *GETNetworkPoolNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get network pool not found response a status code equal to that given
func (o *GETNetworkPoolNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GETNetworkPoolNotFound) Error() string {
	return fmt.Sprintf("[GET /v1/network-pools/{id}][%d] getNetworkPoolNotFound  %+v", 404, o.Payload)
}

func (o *GETNetworkPoolNotFound) String() string {
	return fmt.Sprintf("[GET /v1/network-pools/{id}][%d] getNetworkPoolNotFound  %+v", 404, o.Payload)
}

func (o *GETNetworkPoolNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GETNetworkPoolNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGETNetworkPoolInternalServerError creates a GETNetworkPoolInternalServerError with default headers values
func NewGETNetworkPoolInternalServerError() *GETNetworkPoolInternalServerError {
	return &GETNetworkPoolInternalServerError{}
}

/*
GETNetworkPoolInternalServerError describes a response with status code 500, with default header values.

Unexpected error
*/
type GETNetworkPoolInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this get network pool internal server error response has a 2xx status code
func (o *GETNetworkPoolInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get network pool internal server error response has a 3xx status code
func (o *GETNetworkPoolInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get network pool internal server error response has a 4xx status code
func (o *GETNetworkPoolInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get network pool internal server error response has a 5xx status code
func (o *GETNetworkPoolInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get network pool internal server error response a status code equal to that given
func (o *GETNetworkPoolInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GETNetworkPoolInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/network-pools/{id}][%d] getNetworkPoolInternalServerError  %+v", 500, o.Payload)
}

func (o *GETNetworkPoolInternalServerError) String() string {
	return fmt.Sprintf("[GET /v1/network-pools/{id}][%d] getNetworkPoolInternalServerError  %+v", 500, o.Payload)
}

func (o *GETNetworkPoolInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *GETNetworkPoolInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
