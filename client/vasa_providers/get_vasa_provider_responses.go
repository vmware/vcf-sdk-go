// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2023 VMware, Inc.
// SPDX-License-Identifier: BSD-2-Clause

package vasa_providers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/vmware/vcf-sdk-go/models"
)

// GETVasaProviderReader is a Reader for the GETVasaProvider structure.
type GETVasaProviderReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GETVasaProviderReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGETVasaProviderOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGETVasaProviderNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGETVasaProviderInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGETVasaProviderOK creates a GETVasaProviderOK with default headers values
func NewGETVasaProviderOK() *GETVasaProviderOK {
	return &GETVasaProviderOK{}
}

/*
GETVasaProviderOK describes a response with status code 200, with default header values.

Ok
*/
type GETVasaProviderOK struct {
	Payload *models.VasaProvider
}

// IsSuccess returns true when this get vasa provider o k response has a 2xx status code
func (o *GETVasaProviderOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get vasa provider o k response has a 3xx status code
func (o *GETVasaProviderOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get vasa provider o k response has a 4xx status code
func (o *GETVasaProviderOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get vasa provider o k response has a 5xx status code
func (o *GETVasaProviderOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get vasa provider o k response a status code equal to that given
func (o *GETVasaProviderOK) IsCode(code int) bool {
	return code == 200
}

func (o *GETVasaProviderOK) Error() string {
	return fmt.Sprintf("[GET /v1/vasa-providers/{id}][%d] getVasaProviderOK  %+v", 200, o.Payload)
}

func (o *GETVasaProviderOK) String() string {
	return fmt.Sprintf("[GET /v1/vasa-providers/{id}][%d] getVasaProviderOK  %+v", 200, o.Payload)
}

func (o *GETVasaProviderOK) GetPayload() *models.VasaProvider {
	return o.Payload
}

func (o *GETVasaProviderOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VasaProvider)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGETVasaProviderNotFound creates a GETVasaProviderNotFound with default headers values
func NewGETVasaProviderNotFound() *GETVasaProviderNotFound {
	return &GETVasaProviderNotFound{}
}

/*
GETVasaProviderNotFound describes a response with status code 404, with default header values.

VASA Provider not found
*/
type GETVasaProviderNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this get vasa provider not found response has a 2xx status code
func (o *GETVasaProviderNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get vasa provider not found response has a 3xx status code
func (o *GETVasaProviderNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get vasa provider not found response has a 4xx status code
func (o *GETVasaProviderNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get vasa provider not found response has a 5xx status code
func (o *GETVasaProviderNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get vasa provider not found response a status code equal to that given
func (o *GETVasaProviderNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GETVasaProviderNotFound) Error() string {
	return fmt.Sprintf("[GET /v1/vasa-providers/{id}][%d] getVasaProviderNotFound  %+v", 404, o.Payload)
}

func (o *GETVasaProviderNotFound) String() string {
	return fmt.Sprintf("[GET /v1/vasa-providers/{id}][%d] getVasaProviderNotFound  %+v", 404, o.Payload)
}

func (o *GETVasaProviderNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GETVasaProviderNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGETVasaProviderInternalServerError creates a GETVasaProviderInternalServerError with default headers values
func NewGETVasaProviderInternalServerError() *GETVasaProviderInternalServerError {
	return &GETVasaProviderInternalServerError{}
}

/*
GETVasaProviderInternalServerError describes a response with status code 500, with default header values.

InternalServerError
*/
type GETVasaProviderInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this get vasa provider internal server error response has a 2xx status code
func (o *GETVasaProviderInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get vasa provider internal server error response has a 3xx status code
func (o *GETVasaProviderInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get vasa provider internal server error response has a 4xx status code
func (o *GETVasaProviderInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get vasa provider internal server error response has a 5xx status code
func (o *GETVasaProviderInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get vasa provider internal server error response a status code equal to that given
func (o *GETVasaProviderInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GETVasaProviderInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/vasa-providers/{id}][%d] getVasaProviderInternalServerError  %+v", 500, o.Payload)
}

func (o *GETVasaProviderInternalServerError) String() string {
	return fmt.Sprintf("[GET /v1/vasa-providers/{id}][%d] getVasaProviderInternalServerError  %+v", 500, o.Payload)
}

func (o *GETVasaProviderInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *GETVasaProviderInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
