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

// GetNetworksOfNetworkPoolReader is a Reader for the GetNetworksOfNetworkPool structure.
type GetNetworksOfNetworkPoolReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetNetworksOfNetworkPoolReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetNetworksOfNetworkPoolOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetNetworksOfNetworkPoolNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetNetworksOfNetworkPoolInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetNetworksOfNetworkPoolOK creates a GetNetworksOfNetworkPoolOK with default headers values
func NewGetNetworksOfNetworkPoolOK() *GetNetworksOfNetworkPoolOK {
	return &GetNetworksOfNetworkPoolOK{}
}

/*
GetNetworksOfNetworkPoolOK describes a response with status code 200, with default header values.

Networks for referenced network pool
*/
type GetNetworksOfNetworkPoolOK struct {
	Payload *models.PageOfNetwork
}

// IsSuccess returns true when this get networks of network pool o k response has a 2xx status code
func (o *GetNetworksOfNetworkPoolOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get networks of network pool o k response has a 3xx status code
func (o *GetNetworksOfNetworkPoolOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get networks of network pool o k response has a 4xx status code
func (o *GetNetworksOfNetworkPoolOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get networks of network pool o k response has a 5xx status code
func (o *GetNetworksOfNetworkPoolOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get networks of network pool o k response a status code equal to that given
func (o *GetNetworksOfNetworkPoolOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetNetworksOfNetworkPoolOK) Error() string {
	return fmt.Sprintf("[GET /v1/network-pools/{id}/networks][%d] getNetworksOfNetworkPoolOK  %+v", 200, o.Payload)
}

func (o *GetNetworksOfNetworkPoolOK) String() string {
	return fmt.Sprintf("[GET /v1/network-pools/{id}/networks][%d] getNetworksOfNetworkPoolOK  %+v", 200, o.Payload)
}

func (o *GetNetworksOfNetworkPoolOK) GetPayload() *models.PageOfNetwork {
	return o.Payload
}

func (o *GetNetworksOfNetworkPoolOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PageOfNetwork)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetNetworksOfNetworkPoolNotFound creates a GetNetworksOfNetworkPoolNotFound with default headers values
func NewGetNetworksOfNetworkPoolNotFound() *GetNetworksOfNetworkPoolNotFound {
	return &GetNetworksOfNetworkPoolNotFound{}
}

/*
GetNetworksOfNetworkPoolNotFound describes a response with status code 404, with default header values.

Network pool not found
*/
type GetNetworksOfNetworkPoolNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this get networks of network pool not found response has a 2xx status code
func (o *GetNetworksOfNetworkPoolNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get networks of network pool not found response has a 3xx status code
func (o *GetNetworksOfNetworkPoolNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get networks of network pool not found response has a 4xx status code
func (o *GetNetworksOfNetworkPoolNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get networks of network pool not found response has a 5xx status code
func (o *GetNetworksOfNetworkPoolNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get networks of network pool not found response a status code equal to that given
func (o *GetNetworksOfNetworkPoolNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetNetworksOfNetworkPoolNotFound) Error() string {
	return fmt.Sprintf("[GET /v1/network-pools/{id}/networks][%d] getNetworksOfNetworkPoolNotFound  %+v", 404, o.Payload)
}

func (o *GetNetworksOfNetworkPoolNotFound) String() string {
	return fmt.Sprintf("[GET /v1/network-pools/{id}/networks][%d] getNetworksOfNetworkPoolNotFound  %+v", 404, o.Payload)
}

func (o *GetNetworksOfNetworkPoolNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetNetworksOfNetworkPoolNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetNetworksOfNetworkPoolInternalServerError creates a GetNetworksOfNetworkPoolInternalServerError with default headers values
func NewGetNetworksOfNetworkPoolInternalServerError() *GetNetworksOfNetworkPoolInternalServerError {
	return &GetNetworksOfNetworkPoolInternalServerError{}
}

/*
GetNetworksOfNetworkPoolInternalServerError describes a response with status code 500, with default header values.

Unexpected error
*/
type GetNetworksOfNetworkPoolInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this get networks of network pool internal server error response has a 2xx status code
func (o *GetNetworksOfNetworkPoolInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get networks of network pool internal server error response has a 3xx status code
func (o *GetNetworksOfNetworkPoolInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get networks of network pool internal server error response has a 4xx status code
func (o *GetNetworksOfNetworkPoolInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get networks of network pool internal server error response has a 5xx status code
func (o *GetNetworksOfNetworkPoolInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get networks of network pool internal server error response a status code equal to that given
func (o *GetNetworksOfNetworkPoolInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetNetworksOfNetworkPoolInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/network-pools/{id}/networks][%d] getNetworksOfNetworkPoolInternalServerError  %+v", 500, o.Payload)
}

func (o *GetNetworksOfNetworkPoolInternalServerError) String() string {
	return fmt.Sprintf("[GET /v1/network-pools/{id}/networks][%d] getNetworksOfNetworkPoolInternalServerError  %+v", 500, o.Payload)
}

func (o *GetNetworksOfNetworkPoolInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetNetworksOfNetworkPoolInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
