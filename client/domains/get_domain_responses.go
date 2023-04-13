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

	"vcf-sdk-go/models"
)

// GETDomainReader is a Reader for the GETDomain structure.
type GETDomainReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GETDomainReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGETDomainOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGETDomainNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGETDomainInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGETDomainOK creates a GETDomainOK with default headers values
func NewGETDomainOK() *GETDomainOK {
	return &GETDomainOK{}
}

/*
GETDomainOK describes a response with status code 200, with default header values.

Ok
*/
type GETDomainOK struct {
	Payload *models.Domain
}

// IsSuccess returns true when this get domain o k response has a 2xx status code
func (o *GETDomainOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get domain o k response has a 3xx status code
func (o *GETDomainOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get domain o k response has a 4xx status code
func (o *GETDomainOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get domain o k response has a 5xx status code
func (o *GETDomainOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get domain o k response a status code equal to that given
func (o *GETDomainOK) IsCode(code int) bool {
	return code == 200
}

func (o *GETDomainOK) Error() string {
	return fmt.Sprintf("[GET /v1/domains/{id}][%d] getDomainOK  %+v", 200, o.Payload)
}

func (o *GETDomainOK) String() string {
	return fmt.Sprintf("[GET /v1/domains/{id}][%d] getDomainOK  %+v", 200, o.Payload)
}

func (o *GETDomainOK) GetPayload() *models.Domain {
	return o.Payload
}

func (o *GETDomainOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Domain)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGETDomainNotFound creates a GETDomainNotFound with default headers values
func NewGETDomainNotFound() *GETDomainNotFound {
	return &GETDomainNotFound{}
}

/*
GETDomainNotFound describes a response with status code 404, with default header values.

Domain not found
*/
type GETDomainNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this get domain not found response has a 2xx status code
func (o *GETDomainNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get domain not found response has a 3xx status code
func (o *GETDomainNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get domain not found response has a 4xx status code
func (o *GETDomainNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get domain not found response has a 5xx status code
func (o *GETDomainNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get domain not found response a status code equal to that given
func (o *GETDomainNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GETDomainNotFound) Error() string {
	return fmt.Sprintf("[GET /v1/domains/{id}][%d] getDomainNotFound  %+v", 404, o.Payload)
}

func (o *GETDomainNotFound) String() string {
	return fmt.Sprintf("[GET /v1/domains/{id}][%d] getDomainNotFound  %+v", 404, o.Payload)
}

func (o *GETDomainNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GETDomainNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGETDomainInternalServerError creates a GETDomainInternalServerError with default headers values
func NewGETDomainInternalServerError() *GETDomainInternalServerError {
	return &GETDomainInternalServerError{}
}

/*
GETDomainInternalServerError describes a response with status code 500, with default header values.

InternalServerError
*/
type GETDomainInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this get domain internal server error response has a 2xx status code
func (o *GETDomainInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get domain internal server error response has a 3xx status code
func (o *GETDomainInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get domain internal server error response has a 4xx status code
func (o *GETDomainInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get domain internal server error response has a 5xx status code
func (o *GETDomainInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get domain internal server error response a status code equal to that given
func (o *GETDomainInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GETDomainInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/domains/{id}][%d] getDomainInternalServerError  %+v", 500, o.Payload)
}

func (o *GETDomainInternalServerError) String() string {
	return fmt.Sprintf("[GET /v1/domains/{id}][%d] getDomainInternalServerError  %+v", 500, o.Payload)
}

func (o *GETDomainInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *GETDomainInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
