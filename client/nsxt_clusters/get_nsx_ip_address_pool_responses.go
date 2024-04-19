// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2023 VMware, Inc.
// SPDX-License-Identifier: BSD-2-Clause

package nsxt_clusters

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/vmware/vcf-sdk-go/models"
)

// GetNsxIPAddressPoolReader is a Reader for the GetNsxIPAddressPool structure.
type GetNsxIPAddressPoolReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetNsxIPAddressPoolReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetNsxIPAddressPoolOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetNsxIPAddressPoolNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetNsxIPAddressPoolInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /v1/nsxt-clusters/{nsxt-cluster-id}/ip-address-pools/{name}] getNsxIpAddressPool", response, response.Code())
	}
}

// NewGetNsxIPAddressPoolOK creates a GetNsxIPAddressPoolOK with default headers values
func NewGetNsxIPAddressPoolOK() *GetNsxIPAddressPoolOK {
	return &GetNsxIPAddressPoolOK{}
}

/*
GetNsxIPAddressPoolOK describes a response with status code 200, with default header values.

Ok
*/
type GetNsxIPAddressPoolOK struct {
	Payload *models.NSXTIPAddressPool
}

// IsSuccess returns true when this get nsx Ip address pool o k response has a 2xx status code
func (o *GetNsxIPAddressPoolOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get nsx Ip address pool o k response has a 3xx status code
func (o *GetNsxIPAddressPoolOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get nsx Ip address pool o k response has a 4xx status code
func (o *GetNsxIPAddressPoolOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get nsx Ip address pool o k response has a 5xx status code
func (o *GetNsxIPAddressPoolOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get nsx Ip address pool o k response a status code equal to that given
func (o *GetNsxIPAddressPoolOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get nsx Ip address pool o k response
func (o *GetNsxIPAddressPoolOK) Code() int {
	return 200
}

func (o *GetNsxIPAddressPoolOK) Error() string {
	return fmt.Sprintf("[GET /v1/nsxt-clusters/{nsxt-cluster-id}/ip-address-pools/{name}][%d] getNsxIpAddressPoolOK  %+v", 200, o.Payload)
}

func (o *GetNsxIPAddressPoolOK) String() string {
	return fmt.Sprintf("[GET /v1/nsxt-clusters/{nsxt-cluster-id}/ip-address-pools/{name}][%d] getNsxIpAddressPoolOK  %+v", 200, o.Payload)
}

func (o *GetNsxIPAddressPoolOK) GetPayload() *models.NSXTIPAddressPool {
	return o.Payload
}

func (o *GetNsxIPAddressPoolOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NSXTIPAddressPool)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetNsxIPAddressPoolNotFound creates a GetNsxIPAddressPoolNotFound with default headers values
func NewGetNsxIPAddressPoolNotFound() *GetNsxIPAddressPoolNotFound {
	return &GetNsxIPAddressPoolNotFound{}
}

/*
GetNsxIPAddressPoolNotFound describes a response with status code 404, with default header values.

IP address pool not found
*/
type GetNsxIPAddressPoolNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this get nsx Ip address pool not found response has a 2xx status code
func (o *GetNsxIPAddressPoolNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get nsx Ip address pool not found response has a 3xx status code
func (o *GetNsxIPAddressPoolNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get nsx Ip address pool not found response has a 4xx status code
func (o *GetNsxIPAddressPoolNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get nsx Ip address pool not found response has a 5xx status code
func (o *GetNsxIPAddressPoolNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get nsx Ip address pool not found response a status code equal to that given
func (o *GetNsxIPAddressPoolNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get nsx Ip address pool not found response
func (o *GetNsxIPAddressPoolNotFound) Code() int {
	return 404
}

func (o *GetNsxIPAddressPoolNotFound) Error() string {
	return fmt.Sprintf("[GET /v1/nsxt-clusters/{nsxt-cluster-id}/ip-address-pools/{name}][%d] getNsxIpAddressPoolNotFound  %+v", 404, o.Payload)
}

func (o *GetNsxIPAddressPoolNotFound) String() string {
	return fmt.Sprintf("[GET /v1/nsxt-clusters/{nsxt-cluster-id}/ip-address-pools/{name}][%d] getNsxIpAddressPoolNotFound  %+v", 404, o.Payload)
}

func (o *GetNsxIPAddressPoolNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetNsxIPAddressPoolNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetNsxIPAddressPoolInternalServerError creates a GetNsxIPAddressPoolInternalServerError with default headers values
func NewGetNsxIPAddressPoolInternalServerError() *GetNsxIPAddressPoolInternalServerError {
	return &GetNsxIPAddressPoolInternalServerError{}
}

/*
GetNsxIPAddressPoolInternalServerError describes a response with status code 500, with default header values.

InternalServerError
*/
type GetNsxIPAddressPoolInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this get nsx Ip address pool internal server error response has a 2xx status code
func (o *GetNsxIPAddressPoolInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get nsx Ip address pool internal server error response has a 3xx status code
func (o *GetNsxIPAddressPoolInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get nsx Ip address pool internal server error response has a 4xx status code
func (o *GetNsxIPAddressPoolInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get nsx Ip address pool internal server error response has a 5xx status code
func (o *GetNsxIPAddressPoolInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get nsx Ip address pool internal server error response a status code equal to that given
func (o *GetNsxIPAddressPoolInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get nsx Ip address pool internal server error response
func (o *GetNsxIPAddressPoolInternalServerError) Code() int {
	return 500
}

func (o *GetNsxIPAddressPoolInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/nsxt-clusters/{nsxt-cluster-id}/ip-address-pools/{name}][%d] getNsxIpAddressPoolInternalServerError  %+v", 500, o.Payload)
}

func (o *GetNsxIPAddressPoolInternalServerError) String() string {
	return fmt.Sprintf("[GET /v1/nsxt-clusters/{nsxt-cluster-id}/ip-address-pools/{name}][%d] getNsxIpAddressPoolInternalServerError  %+v", 500, o.Payload)
}

func (o *GetNsxIPAddressPoolInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetNsxIPAddressPoolInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}