// Code generated by go-swagger; DO NOT EDIT.

// Copyright © 2023 VMware, Inc. All Rights Reserved.
// Code licenced under: BSD-2

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

// GETNSXTIPAddressPoolReader is a Reader for the GETNSXTIPAddressPool structure.
type GETNSXTIPAddressPoolReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GETNSXTIPAddressPoolReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGETNSXTIPAddressPoolOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGETNSXTIPAddressPoolNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGETNSXTIPAddressPoolInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGETNSXTIPAddressPoolOK creates a GETNSXTIPAddressPoolOK with default headers values
func NewGETNSXTIPAddressPoolOK() *GETNSXTIPAddressPoolOK {
	return &GETNSXTIPAddressPoolOK{}
}

/*
GETNSXTIPAddressPoolOK describes a response with status code 200, with default header values.

Ok
*/
type GETNSXTIPAddressPoolOK struct {
	Payload *models.NSXTIPAddressPool
}

// IsSuccess returns true when this get Nsxt Ip address pool o k response has a 2xx status code
func (o *GETNSXTIPAddressPoolOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get Nsxt Ip address pool o k response has a 3xx status code
func (o *GETNSXTIPAddressPoolOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get Nsxt Ip address pool o k response has a 4xx status code
func (o *GETNSXTIPAddressPoolOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get Nsxt Ip address pool o k response has a 5xx status code
func (o *GETNSXTIPAddressPoolOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get Nsxt Ip address pool o k response a status code equal to that given
func (o *GETNSXTIPAddressPoolOK) IsCode(code int) bool {
	return code == 200
}

func (o *GETNSXTIPAddressPoolOK) Error() string {
	return fmt.Sprintf("[GET /v1/nsxt-clusters/{nsxt-cluster-id}/ip-address-pools/{name}][%d] getNsxtIpAddressPoolOK  %+v", 200, o.Payload)
}

func (o *GETNSXTIPAddressPoolOK) String() string {
	return fmt.Sprintf("[GET /v1/nsxt-clusters/{nsxt-cluster-id}/ip-address-pools/{name}][%d] getNsxtIpAddressPoolOK  %+v", 200, o.Payload)
}

func (o *GETNSXTIPAddressPoolOK) GetPayload() *models.NSXTIPAddressPool {
	return o.Payload
}

func (o *GETNSXTIPAddressPoolOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NSXTIPAddressPool)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGETNSXTIPAddressPoolNotFound creates a GETNSXTIPAddressPoolNotFound with default headers values
func NewGETNSXTIPAddressPoolNotFound() *GETNSXTIPAddressPoolNotFound {
	return &GETNSXTIPAddressPoolNotFound{}
}

/*
GETNSXTIPAddressPoolNotFound describes a response with status code 404, with default header values.

IP address pool not found
*/
type GETNSXTIPAddressPoolNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this get Nsxt Ip address pool not found response has a 2xx status code
func (o *GETNSXTIPAddressPoolNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get Nsxt Ip address pool not found response has a 3xx status code
func (o *GETNSXTIPAddressPoolNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get Nsxt Ip address pool not found response has a 4xx status code
func (o *GETNSXTIPAddressPoolNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get Nsxt Ip address pool not found response has a 5xx status code
func (o *GETNSXTIPAddressPoolNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get Nsxt Ip address pool not found response a status code equal to that given
func (o *GETNSXTIPAddressPoolNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GETNSXTIPAddressPoolNotFound) Error() string {
	return fmt.Sprintf("[GET /v1/nsxt-clusters/{nsxt-cluster-id}/ip-address-pools/{name}][%d] getNsxtIpAddressPoolNotFound  %+v", 404, o.Payload)
}

func (o *GETNSXTIPAddressPoolNotFound) String() string {
	return fmt.Sprintf("[GET /v1/nsxt-clusters/{nsxt-cluster-id}/ip-address-pools/{name}][%d] getNsxtIpAddressPoolNotFound  %+v", 404, o.Payload)
}

func (o *GETNSXTIPAddressPoolNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GETNSXTIPAddressPoolNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGETNSXTIPAddressPoolInternalServerError creates a GETNSXTIPAddressPoolInternalServerError with default headers values
func NewGETNSXTIPAddressPoolInternalServerError() *GETNSXTIPAddressPoolInternalServerError {
	return &GETNSXTIPAddressPoolInternalServerError{}
}

/*
GETNSXTIPAddressPoolInternalServerError describes a response with status code 500, with default header values.

InternalServerError
*/
type GETNSXTIPAddressPoolInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this get Nsxt Ip address pool internal server error response has a 2xx status code
func (o *GETNSXTIPAddressPoolInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get Nsxt Ip address pool internal server error response has a 3xx status code
func (o *GETNSXTIPAddressPoolInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get Nsxt Ip address pool internal server error response has a 4xx status code
func (o *GETNSXTIPAddressPoolInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get Nsxt Ip address pool internal server error response has a 5xx status code
func (o *GETNSXTIPAddressPoolInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get Nsxt Ip address pool internal server error response a status code equal to that given
func (o *GETNSXTIPAddressPoolInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GETNSXTIPAddressPoolInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/nsxt-clusters/{nsxt-cluster-id}/ip-address-pools/{name}][%d] getNsxtIpAddressPoolInternalServerError  %+v", 500, o.Payload)
}

func (o *GETNSXTIPAddressPoolInternalServerError) String() string {
	return fmt.Sprintf("[GET /v1/nsxt-clusters/{nsxt-cluster-id}/ip-address-pools/{name}][%d] getNsxtIpAddressPoolInternalServerError  %+v", 500, o.Payload)
}

func (o *GETNSXTIPAddressPoolInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *GETNSXTIPAddressPoolInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}