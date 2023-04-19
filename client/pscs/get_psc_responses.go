// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2023 VMware, Inc.
// SPDX-License-Identifier: BSD-2-Clause

package pscs

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/vmware/vcf-sdk-go/models"
)

// GETPscReader is a Reader for the GETPsc structure.
type GETPscReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GETPscReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGETPscOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGETPscNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGETPscInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGETPscOK creates a GETPscOK with default headers values
func NewGETPscOK() *GETPscOK {
	return &GETPscOK{}
}

/*
GETPscOK describes a response with status code 200, with default header values.

Ok
*/
type GETPscOK struct {
	Payload *models.Psc
}

// IsSuccess returns true when this get psc o k response has a 2xx status code
func (o *GETPscOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get psc o k response has a 3xx status code
func (o *GETPscOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get psc o k response has a 4xx status code
func (o *GETPscOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get psc o k response has a 5xx status code
func (o *GETPscOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get psc o k response a status code equal to that given
func (o *GETPscOK) IsCode(code int) bool {
	return code == 200
}

func (o *GETPscOK) Error() string {
	return fmt.Sprintf("[GET /v1/pscs/{id}][%d] getPscOK  %+v", 200, o.Payload)
}

func (o *GETPscOK) String() string {
	return fmt.Sprintf("[GET /v1/pscs/{id}][%d] getPscOK  %+v", 200, o.Payload)
}

func (o *GETPscOK) GetPayload() *models.Psc {
	return o.Payload
}

func (o *GETPscOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Psc)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGETPscNotFound creates a GETPscNotFound with default headers values
func NewGETPscNotFound() *GETPscNotFound {
	return &GETPscNotFound{}
}

/*
GETPscNotFound describes a response with status code 404, with default header values.

Psc not found
*/
type GETPscNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this get psc not found response has a 2xx status code
func (o *GETPscNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get psc not found response has a 3xx status code
func (o *GETPscNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get psc not found response has a 4xx status code
func (o *GETPscNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get psc not found response has a 5xx status code
func (o *GETPscNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get psc not found response a status code equal to that given
func (o *GETPscNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GETPscNotFound) Error() string {
	return fmt.Sprintf("[GET /v1/pscs/{id}][%d] getPscNotFound  %+v", 404, o.Payload)
}

func (o *GETPscNotFound) String() string {
	return fmt.Sprintf("[GET /v1/pscs/{id}][%d] getPscNotFound  %+v", 404, o.Payload)
}

func (o *GETPscNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GETPscNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGETPscInternalServerError creates a GETPscInternalServerError with default headers values
func NewGETPscInternalServerError() *GETPscInternalServerError {
	return &GETPscInternalServerError{}
}

/*
GETPscInternalServerError describes a response with status code 500, with default header values.

InternalServerError
*/
type GETPscInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this get psc internal server error response has a 2xx status code
func (o *GETPscInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get psc internal server error response has a 3xx status code
func (o *GETPscInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get psc internal server error response has a 4xx status code
func (o *GETPscInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get psc internal server error response has a 5xx status code
func (o *GETPscInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get psc internal server error response a status code equal to that given
func (o *GETPscInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GETPscInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/pscs/{id}][%d] getPscInternalServerError  %+v", 500, o.Payload)
}

func (o *GETPscInternalServerError) String() string {
	return fmt.Sprintf("[GET /v1/pscs/{id}][%d] getPscInternalServerError  %+v", 500, o.Payload)
}

func (o *GETPscInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *GETPscInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
