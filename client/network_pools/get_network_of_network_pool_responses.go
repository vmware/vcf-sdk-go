// Code generated by go-swagger; DO NOT EDIT.

// Copyright © 2023 VMware, Inc. All Rights Reserved.
// Code licenced under: BSD-2

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

// GETNetworkOfNetworkPoolReader is a Reader for the GETNetworkOfNetworkPool structure.
type GETNetworkOfNetworkPoolReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GETNetworkOfNetworkPoolReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGETNetworkOfNetworkPoolOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGETNetworkOfNetworkPoolNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGETNetworkOfNetworkPoolInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGETNetworkOfNetworkPoolOK creates a GETNetworkOfNetworkPoolOK with default headers values
func NewGETNetworkOfNetworkPoolOK() *GETNetworkOfNetworkPoolOK {
	return &GETNetworkOfNetworkPoolOK{}
}

/*
GETNetworkOfNetworkPoolOK describes a response with status code 200, with default header values.

Network for referenced network pool
*/
type GETNetworkOfNetworkPoolOK struct {
	Payload *models.Network
}

// IsSuccess returns true when this get network of network pool o k response has a 2xx status code
func (o *GETNetworkOfNetworkPoolOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get network of network pool o k response has a 3xx status code
func (o *GETNetworkOfNetworkPoolOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get network of network pool o k response has a 4xx status code
func (o *GETNetworkOfNetworkPoolOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get network of network pool o k response has a 5xx status code
func (o *GETNetworkOfNetworkPoolOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get network of network pool o k response a status code equal to that given
func (o *GETNetworkOfNetworkPoolOK) IsCode(code int) bool {
	return code == 200
}

func (o *GETNetworkOfNetworkPoolOK) Error() string {
	return fmt.Sprintf("[GET /v1/network-pools/{id}/networks/{networkId}][%d] getNetworkOfNetworkPoolOK  %+v", 200, o.Payload)
}

func (o *GETNetworkOfNetworkPoolOK) String() string {
	return fmt.Sprintf("[GET /v1/network-pools/{id}/networks/{networkId}][%d] getNetworkOfNetworkPoolOK  %+v", 200, o.Payload)
}

func (o *GETNetworkOfNetworkPoolOK) GetPayload() *models.Network {
	return o.Payload
}

func (o *GETNetworkOfNetworkPoolOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Network)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGETNetworkOfNetworkPoolNotFound creates a GETNetworkOfNetworkPoolNotFound with default headers values
func NewGETNetworkOfNetworkPoolNotFound() *GETNetworkOfNetworkPoolNotFound {
	return &GETNetworkOfNetworkPoolNotFound{}
}

/*
GETNetworkOfNetworkPoolNotFound describes a response with status code 404, with default header values.

Networkpool not found
*/
type GETNetworkOfNetworkPoolNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this get network of network pool not found response has a 2xx status code
func (o *GETNetworkOfNetworkPoolNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get network of network pool not found response has a 3xx status code
func (o *GETNetworkOfNetworkPoolNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get network of network pool not found response has a 4xx status code
func (o *GETNetworkOfNetworkPoolNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get network of network pool not found response has a 5xx status code
func (o *GETNetworkOfNetworkPoolNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get network of network pool not found response a status code equal to that given
func (o *GETNetworkOfNetworkPoolNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GETNetworkOfNetworkPoolNotFound) Error() string {
	return fmt.Sprintf("[GET /v1/network-pools/{id}/networks/{networkId}][%d] getNetworkOfNetworkPoolNotFound  %+v", 404, o.Payload)
}

func (o *GETNetworkOfNetworkPoolNotFound) String() string {
	return fmt.Sprintf("[GET /v1/network-pools/{id}/networks/{networkId}][%d] getNetworkOfNetworkPoolNotFound  %+v", 404, o.Payload)
}

func (o *GETNetworkOfNetworkPoolNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GETNetworkOfNetworkPoolNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGETNetworkOfNetworkPoolInternalServerError creates a GETNetworkOfNetworkPoolInternalServerError with default headers values
func NewGETNetworkOfNetworkPoolInternalServerError() *GETNetworkOfNetworkPoolInternalServerError {
	return &GETNetworkOfNetworkPoolInternalServerError{}
}

/*
GETNetworkOfNetworkPoolInternalServerError describes a response with status code 500, with default header values.

Unexpected error
*/
type GETNetworkOfNetworkPoolInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this get network of network pool internal server error response has a 2xx status code
func (o *GETNetworkOfNetworkPoolInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get network of network pool internal server error response has a 3xx status code
func (o *GETNetworkOfNetworkPoolInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get network of network pool internal server error response has a 4xx status code
func (o *GETNetworkOfNetworkPoolInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get network of network pool internal server error response has a 5xx status code
func (o *GETNetworkOfNetworkPoolInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get network of network pool internal server error response a status code equal to that given
func (o *GETNetworkOfNetworkPoolInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GETNetworkOfNetworkPoolInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/network-pools/{id}/networks/{networkId}][%d] getNetworkOfNetworkPoolInternalServerError  %+v", 500, o.Payload)
}

func (o *GETNetworkOfNetworkPoolInternalServerError) String() string {
	return fmt.Sprintf("[GET /v1/network-pools/{id}/networks/{networkId}][%d] getNetworkOfNetworkPoolInternalServerError  %+v", 500, o.Payload)
}

func (o *GETNetworkOfNetworkPoolInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *GETNetworkOfNetworkPoolInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
