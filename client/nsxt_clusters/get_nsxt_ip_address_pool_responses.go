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

// GetNSXTIPAddressPoolReader is a Reader for the GetNSXTIPAddressPool structure.
type GetNSXTIPAddressPoolReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetNSXTIPAddressPoolReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetNSXTIPAddressPoolOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetNSXTIPAddressPoolNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetNSXTIPAddressPoolInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetNSXTIPAddressPoolOK creates a GetNSXTIPAddressPoolOK with default headers values
func NewGetNSXTIPAddressPoolOK() *GetNSXTIPAddressPoolOK {
	return &GetNSXTIPAddressPoolOK{}
}

/*
GetNSXTIPAddressPoolOK describes a response with status code 200, with default header values.

Ok
*/
type GetNSXTIPAddressPoolOK struct {
	Payload *models.NSXTIPAddressPool
}

// IsSuccess returns true when this get Nsxt Ip address pool o k response has a 2xx status code
func (o *GetNSXTIPAddressPoolOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get Nsxt Ip address pool o k response has a 3xx status code
func (o *GetNSXTIPAddressPoolOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get Nsxt Ip address pool o k response has a 4xx status code
func (o *GetNSXTIPAddressPoolOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get Nsxt Ip address pool o k response has a 5xx status code
func (o *GetNSXTIPAddressPoolOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get Nsxt Ip address pool o k response a status code equal to that given
func (o *GetNSXTIPAddressPoolOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetNSXTIPAddressPoolOK) Error() string {
	return fmt.Sprintf("[GET /v1/nsxt-clusters/{nsxt-cluster-id}/ip-address-pools/{name}][%d] getNsxtIpAddressPoolOK  %+v", 200, o.Payload)
}

func (o *GetNSXTIPAddressPoolOK) String() string {
	return fmt.Sprintf("[GET /v1/nsxt-clusters/{nsxt-cluster-id}/ip-address-pools/{name}][%d] getNsxtIpAddressPoolOK  %+v", 200, o.Payload)
}

func (o *GetNSXTIPAddressPoolOK) GetPayload() *models.NSXTIPAddressPool {
	return o.Payload
}

func (o *GetNSXTIPAddressPoolOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NSXTIPAddressPool)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetNSXTIPAddressPoolNotFound creates a GetNSXTIPAddressPoolNotFound with default headers values
func NewGetNSXTIPAddressPoolNotFound() *GetNSXTIPAddressPoolNotFound {
	return &GetNSXTIPAddressPoolNotFound{}
}

/*
GetNSXTIPAddressPoolNotFound describes a response with status code 404, with default header values.

IP address pool not found
*/
type GetNSXTIPAddressPoolNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this get Nsxt Ip address pool not found response has a 2xx status code
func (o *GetNSXTIPAddressPoolNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get Nsxt Ip address pool not found response has a 3xx status code
func (o *GetNSXTIPAddressPoolNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get Nsxt Ip address pool not found response has a 4xx status code
func (o *GetNSXTIPAddressPoolNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get Nsxt Ip address pool not found response has a 5xx status code
func (o *GetNSXTIPAddressPoolNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get Nsxt Ip address pool not found response a status code equal to that given
func (o *GetNSXTIPAddressPoolNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetNSXTIPAddressPoolNotFound) Error() string {
	return fmt.Sprintf("[GET /v1/nsxt-clusters/{nsxt-cluster-id}/ip-address-pools/{name}][%d] getNsxtIpAddressPoolNotFound  %+v", 404, o.Payload)
}

func (o *GetNSXTIPAddressPoolNotFound) String() string {
	return fmt.Sprintf("[GET /v1/nsxt-clusters/{nsxt-cluster-id}/ip-address-pools/{name}][%d] getNsxtIpAddressPoolNotFound  %+v", 404, o.Payload)
}

func (o *GetNSXTIPAddressPoolNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetNSXTIPAddressPoolNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetNSXTIPAddressPoolInternalServerError creates a GetNSXTIPAddressPoolInternalServerError with default headers values
func NewGetNSXTIPAddressPoolInternalServerError() *GetNSXTIPAddressPoolInternalServerError {
	return &GetNSXTIPAddressPoolInternalServerError{}
}

/*
GetNSXTIPAddressPoolInternalServerError describes a response with status code 500, with default header values.

InternalServerError
*/
type GetNSXTIPAddressPoolInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this get Nsxt Ip address pool internal server error response has a 2xx status code
func (o *GetNSXTIPAddressPoolInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get Nsxt Ip address pool internal server error response has a 3xx status code
func (o *GetNSXTIPAddressPoolInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get Nsxt Ip address pool internal server error response has a 4xx status code
func (o *GetNSXTIPAddressPoolInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get Nsxt Ip address pool internal server error response has a 5xx status code
func (o *GetNSXTIPAddressPoolInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get Nsxt Ip address pool internal server error response a status code equal to that given
func (o *GetNSXTIPAddressPoolInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetNSXTIPAddressPoolInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/nsxt-clusters/{nsxt-cluster-id}/ip-address-pools/{name}][%d] getNsxtIpAddressPoolInternalServerError  %+v", 500, o.Payload)
}

func (o *GetNSXTIPAddressPoolInternalServerError) String() string {
	return fmt.Sprintf("[GET /v1/nsxt-clusters/{nsxt-cluster-id}/ip-address-pools/{name}][%d] getNsxtIpAddressPoolInternalServerError  %+v", 500, o.Payload)
}

func (o *GetNSXTIPAddressPoolInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetNSXTIPAddressPoolInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
