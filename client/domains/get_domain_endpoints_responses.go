// Code generated by go-swagger; DO NOT EDIT.

// Copyright © 2023 VMware, Inc. All Rights Reserved.
// Code licenced under: BSD-2

package domains

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/vmware/vcf-sdk-go/models"
)

// GETDomainEndpointsReader is a Reader for the GETDomainEndpoints structure.
type GETDomainEndpointsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GETDomainEndpointsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGETDomainEndpointsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGETDomainEndpointsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGETDomainEndpointsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGETDomainEndpointsOK creates a GETDomainEndpointsOK with default headers values
func NewGETDomainEndpointsOK() *GETDomainEndpointsOK {
	return &GETDomainEndpointsOK{}
}

/*
GETDomainEndpointsOK describes a response with status code 200, with default header values.

Ok
*/
type GETDomainEndpointsOK struct {
	Payload *models.PageOfEndpoint
}

// IsSuccess returns true when this get domain endpoints o k response has a 2xx status code
func (o *GETDomainEndpointsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get domain endpoints o k response has a 3xx status code
func (o *GETDomainEndpointsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get domain endpoints o k response has a 4xx status code
func (o *GETDomainEndpointsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get domain endpoints o k response has a 5xx status code
func (o *GETDomainEndpointsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get domain endpoints o k response a status code equal to that given
func (o *GETDomainEndpointsOK) IsCode(code int) bool {
	return code == 200
}

func (o *GETDomainEndpointsOK) Error() string {
	return fmt.Sprintf("[GET /v1/domains/{id}/endpoints][%d] getDomainEndpointsOK  %+v", 200, o.Payload)
}

func (o *GETDomainEndpointsOK) String() string {
	return fmt.Sprintf("[GET /v1/domains/{id}/endpoints][%d] getDomainEndpointsOK  %+v", 200, o.Payload)
}

func (o *GETDomainEndpointsOK) GetPayload() *models.PageOfEndpoint {
	return o.Payload
}

func (o *GETDomainEndpointsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PageOfEndpoint)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGETDomainEndpointsNotFound creates a GETDomainEndpointsNotFound with default headers values
func NewGETDomainEndpointsNotFound() *GETDomainEndpointsNotFound {
	return &GETDomainEndpointsNotFound{}
}

/*
GETDomainEndpointsNotFound describes a response with status code 404, with default header values.

Domain not found
*/
type GETDomainEndpointsNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this get domain endpoints not found response has a 2xx status code
func (o *GETDomainEndpointsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get domain endpoints not found response has a 3xx status code
func (o *GETDomainEndpointsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get domain endpoints not found response has a 4xx status code
func (o *GETDomainEndpointsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get domain endpoints not found response has a 5xx status code
func (o *GETDomainEndpointsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get domain endpoints not found response a status code equal to that given
func (o *GETDomainEndpointsNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GETDomainEndpointsNotFound) Error() string {
	return fmt.Sprintf("[GET /v1/domains/{id}/endpoints][%d] getDomainEndpointsNotFound  %+v", 404, o.Payload)
}

func (o *GETDomainEndpointsNotFound) String() string {
	return fmt.Sprintf("[GET /v1/domains/{id}/endpoints][%d] getDomainEndpointsNotFound  %+v", 404, o.Payload)
}

func (o *GETDomainEndpointsNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GETDomainEndpointsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGETDomainEndpointsInternalServerError creates a GETDomainEndpointsInternalServerError with default headers values
func NewGETDomainEndpointsInternalServerError() *GETDomainEndpointsInternalServerError {
	return &GETDomainEndpointsInternalServerError{}
}

/*
GETDomainEndpointsInternalServerError describes a response with status code 500, with default header values.

InternalServerError
*/
type GETDomainEndpointsInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this get domain endpoints internal server error response has a 2xx status code
func (o *GETDomainEndpointsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get domain endpoints internal server error response has a 3xx status code
func (o *GETDomainEndpointsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get domain endpoints internal server error response has a 4xx status code
func (o *GETDomainEndpointsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get domain endpoints internal server error response has a 5xx status code
func (o *GETDomainEndpointsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get domain endpoints internal server error response a status code equal to that given
func (o *GETDomainEndpointsInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GETDomainEndpointsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/domains/{id}/endpoints][%d] getDomainEndpointsInternalServerError  %+v", 500, o.Payload)
}

func (o *GETDomainEndpointsInternalServerError) String() string {
	return fmt.Sprintf("[GET /v1/domains/{id}/endpoints][%d] getDomainEndpointsInternalServerError  %+v", 500, o.Payload)
}

func (o *GETDomainEndpointsInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *GETDomainEndpointsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
