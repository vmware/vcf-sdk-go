// Code generated by go-swagger; DO NOT EDIT.

// Copyright © 2023 VMware, Inc. All Rights Reserved.
// Code licenced under: BSD-2

package vcf_services

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"vcf-sdk-go/models"
)

// GETVcfServiceReader is a Reader for the GETVcfService structure.
type GETVcfServiceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GETVcfServiceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGETVcfServiceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGETVcfServiceNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGETVcfServiceInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGETVcfServiceOK creates a GETVcfServiceOK with default headers values
func NewGETVcfServiceOK() *GETVcfServiceOK {
	return &GETVcfServiceOK{}
}

/*
GETVcfServiceOK describes a response with status code 200, with default header values.

Ok
*/
type GETVcfServiceOK struct {
	Payload *models.VcfService
}

// IsSuccess returns true when this get vcf service o k response has a 2xx status code
func (o *GETVcfServiceOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get vcf service o k response has a 3xx status code
func (o *GETVcfServiceOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get vcf service o k response has a 4xx status code
func (o *GETVcfServiceOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get vcf service o k response has a 5xx status code
func (o *GETVcfServiceOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get vcf service o k response a status code equal to that given
func (o *GETVcfServiceOK) IsCode(code int) bool {
	return code == 200
}

func (o *GETVcfServiceOK) Error() string {
	return fmt.Sprintf("[GET /v1/vcf-services/{id}][%d] getVcfServiceOK  %+v", 200, o.Payload)
}

func (o *GETVcfServiceOK) String() string {
	return fmt.Sprintf("[GET /v1/vcf-services/{id}][%d] getVcfServiceOK  %+v", 200, o.Payload)
}

func (o *GETVcfServiceOK) GetPayload() *models.VcfService {
	return o.Payload
}

func (o *GETVcfServiceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VcfService)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGETVcfServiceNotFound creates a GETVcfServiceNotFound with default headers values
func NewGETVcfServiceNotFound() *GETVcfServiceNotFound {
	return &GETVcfServiceNotFound{}
}

/*
GETVcfServiceNotFound describes a response with status code 404, with default header values.

VcfService not found
*/
type GETVcfServiceNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this get vcf service not found response has a 2xx status code
func (o *GETVcfServiceNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get vcf service not found response has a 3xx status code
func (o *GETVcfServiceNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get vcf service not found response has a 4xx status code
func (o *GETVcfServiceNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get vcf service not found response has a 5xx status code
func (o *GETVcfServiceNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get vcf service not found response a status code equal to that given
func (o *GETVcfServiceNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GETVcfServiceNotFound) Error() string {
	return fmt.Sprintf("[GET /v1/vcf-services/{id}][%d] getVcfServiceNotFound  %+v", 404, o.Payload)
}

func (o *GETVcfServiceNotFound) String() string {
	return fmt.Sprintf("[GET /v1/vcf-services/{id}][%d] getVcfServiceNotFound  %+v", 404, o.Payload)
}

func (o *GETVcfServiceNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GETVcfServiceNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGETVcfServiceInternalServerError creates a GETVcfServiceInternalServerError with default headers values
func NewGETVcfServiceInternalServerError() *GETVcfServiceInternalServerError {
	return &GETVcfServiceInternalServerError{}
}

/*
GETVcfServiceInternalServerError describes a response with status code 500, with default header values.

InternalServerError
*/
type GETVcfServiceInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this get vcf service internal server error response has a 2xx status code
func (o *GETVcfServiceInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get vcf service internal server error response has a 3xx status code
func (o *GETVcfServiceInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get vcf service internal server error response has a 4xx status code
func (o *GETVcfServiceInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get vcf service internal server error response has a 5xx status code
func (o *GETVcfServiceInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get vcf service internal server error response a status code equal to that given
func (o *GETVcfServiceInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GETVcfServiceInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/vcf-services/{id}][%d] getVcfServiceInternalServerError  %+v", 500, o.Payload)
}

func (o *GETVcfServiceInternalServerError) String() string {
	return fmt.Sprintf("[GET /v1/vcf-services/{id}][%d] getVcfServiceInternalServerError  %+v", 500, o.Payload)
}

func (o *GETVcfServiceInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *GETVcfServiceInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
