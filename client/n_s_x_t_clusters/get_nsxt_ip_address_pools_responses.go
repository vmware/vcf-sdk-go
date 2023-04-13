// Code generated by go-swagger; DO NOT EDIT.

// Copyright © 2023 VMware, Inc. All Rights Reserved.
// Code licenced under: BSD-2

package n_s_x_t_clusters

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"vcf-sdk-go/models"
)

// GETNSXTIPAddressPoolsReader is a Reader for the GETNSXTIPAddressPools structure.
type GETNSXTIPAddressPoolsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GETNSXTIPAddressPoolsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGETNSXTIPAddressPoolsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGETNSXTIPAddressPoolsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGETNSXTIPAddressPoolsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGETNSXTIPAddressPoolsOK creates a GETNSXTIPAddressPoolsOK with default headers values
func NewGETNSXTIPAddressPoolsOK() *GETNSXTIPAddressPoolsOK {
	return &GETNSXTIPAddressPoolsOK{}
}

/*
GETNSXTIPAddressPoolsOK describes a response with status code 200, with default header values.

Ok
*/
type GETNSXTIPAddressPoolsOK struct {
	Payload *models.NSXTIPAddressPool
}

// IsSuccess returns true when this get Nsxt Ip address pools o k response has a 2xx status code
func (o *GETNSXTIPAddressPoolsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get Nsxt Ip address pools o k response has a 3xx status code
func (o *GETNSXTIPAddressPoolsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get Nsxt Ip address pools o k response has a 4xx status code
func (o *GETNSXTIPAddressPoolsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get Nsxt Ip address pools o k response has a 5xx status code
func (o *GETNSXTIPAddressPoolsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get Nsxt Ip address pools o k response a status code equal to that given
func (o *GETNSXTIPAddressPoolsOK) IsCode(code int) bool {
	return code == 200
}

func (o *GETNSXTIPAddressPoolsOK) Error() string {
	return fmt.Sprintf("[GET /v1/nsxt-clusters/{nsxt-cluster-id}/ip-address-pools][%d] getNsxtIpAddressPoolsOK  %+v", 200, o.Payload)
}

func (o *GETNSXTIPAddressPoolsOK) String() string {
	return fmt.Sprintf("[GET /v1/nsxt-clusters/{nsxt-cluster-id}/ip-address-pools][%d] getNsxtIpAddressPoolsOK  %+v", 200, o.Payload)
}

func (o *GETNSXTIPAddressPoolsOK) GetPayload() *models.NSXTIPAddressPool {
	return o.Payload
}

func (o *GETNSXTIPAddressPoolsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NSXTIPAddressPool)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGETNSXTIPAddressPoolsNotFound creates a GETNSXTIPAddressPoolsNotFound with default headers values
func NewGETNSXTIPAddressPoolsNotFound() *GETNSXTIPAddressPoolsNotFound {
	return &GETNSXTIPAddressPoolsNotFound{}
}

/*
GETNSXTIPAddressPoolsNotFound describes a response with status code 404, with default header values.

IP address pools not found
*/
type GETNSXTIPAddressPoolsNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this get Nsxt Ip address pools not found response has a 2xx status code
func (o *GETNSXTIPAddressPoolsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get Nsxt Ip address pools not found response has a 3xx status code
func (o *GETNSXTIPAddressPoolsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get Nsxt Ip address pools not found response has a 4xx status code
func (o *GETNSXTIPAddressPoolsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get Nsxt Ip address pools not found response has a 5xx status code
func (o *GETNSXTIPAddressPoolsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get Nsxt Ip address pools not found response a status code equal to that given
func (o *GETNSXTIPAddressPoolsNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GETNSXTIPAddressPoolsNotFound) Error() string {
	return fmt.Sprintf("[GET /v1/nsxt-clusters/{nsxt-cluster-id}/ip-address-pools][%d] getNsxtIpAddressPoolsNotFound  %+v", 404, o.Payload)
}

func (o *GETNSXTIPAddressPoolsNotFound) String() string {
	return fmt.Sprintf("[GET /v1/nsxt-clusters/{nsxt-cluster-id}/ip-address-pools][%d] getNsxtIpAddressPoolsNotFound  %+v", 404, o.Payload)
}

func (o *GETNSXTIPAddressPoolsNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GETNSXTIPAddressPoolsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGETNSXTIPAddressPoolsInternalServerError creates a GETNSXTIPAddressPoolsInternalServerError with default headers values
func NewGETNSXTIPAddressPoolsInternalServerError() *GETNSXTIPAddressPoolsInternalServerError {
	return &GETNSXTIPAddressPoolsInternalServerError{}
}

/*
GETNSXTIPAddressPoolsInternalServerError describes a response with status code 500, with default header values.

InternalServerError
*/
type GETNSXTIPAddressPoolsInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this get Nsxt Ip address pools internal server error response has a 2xx status code
func (o *GETNSXTIPAddressPoolsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get Nsxt Ip address pools internal server error response has a 3xx status code
func (o *GETNSXTIPAddressPoolsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get Nsxt Ip address pools internal server error response has a 4xx status code
func (o *GETNSXTIPAddressPoolsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get Nsxt Ip address pools internal server error response has a 5xx status code
func (o *GETNSXTIPAddressPoolsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get Nsxt Ip address pools internal server error response a status code equal to that given
func (o *GETNSXTIPAddressPoolsInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GETNSXTIPAddressPoolsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/nsxt-clusters/{nsxt-cluster-id}/ip-address-pools][%d] getNsxtIpAddressPoolsInternalServerError  %+v", 500, o.Payload)
}

func (o *GETNSXTIPAddressPoolsInternalServerError) String() string {
	return fmt.Sprintf("[GET /v1/nsxt-clusters/{nsxt-cluster-id}/ip-address-pools][%d] getNsxtIpAddressPoolsInternalServerError  %+v", 500, o.Payload)
}

func (o *GETNSXTIPAddressPoolsInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *GETNSXTIPAddressPoolsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
