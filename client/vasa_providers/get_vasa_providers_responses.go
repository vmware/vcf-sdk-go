// Code generated by go-swagger; DO NOT EDIT.

// Copyright © 2023 VMware, Inc. All Rights Reserved.
// Code licenced under: BSD-2

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

// GETVasaProvidersReader is a Reader for the GETVasaProviders structure.
type GETVasaProvidersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GETVasaProvidersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGETVasaProvidersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGETVasaProvidersBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGETVasaProvidersInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGETVasaProvidersOK creates a GETVasaProvidersOK with default headers values
func NewGETVasaProvidersOK() *GETVasaProvidersOK {
	return &GETVasaProvidersOK{}
}

/*
GETVasaProvidersOK describes a response with status code 200, with default header values.

Ok
*/
type GETVasaProvidersOK struct {
	Payload *models.PageOfVasaProvider
}

// IsSuccess returns true when this get vasa providers o k response has a 2xx status code
func (o *GETVasaProvidersOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get vasa providers o k response has a 3xx status code
func (o *GETVasaProvidersOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get vasa providers o k response has a 4xx status code
func (o *GETVasaProvidersOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get vasa providers o k response has a 5xx status code
func (o *GETVasaProvidersOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get vasa providers o k response a status code equal to that given
func (o *GETVasaProvidersOK) IsCode(code int) bool {
	return code == 200
}

func (o *GETVasaProvidersOK) Error() string {
	return fmt.Sprintf("[GET /v1/vasa-providers][%d] getVasaProvidersOK  %+v", 200, o.Payload)
}

func (o *GETVasaProvidersOK) String() string {
	return fmt.Sprintf("[GET /v1/vasa-providers][%d] getVasaProvidersOK  %+v", 200, o.Payload)
}

func (o *GETVasaProvidersOK) GetPayload() *models.PageOfVasaProvider {
	return o.Payload
}

func (o *GETVasaProvidersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PageOfVasaProvider)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGETVasaProvidersBadRequest creates a GETVasaProvidersBadRequest with default headers values
func NewGETVasaProvidersBadRequest() *GETVasaProvidersBadRequest {
	return &GETVasaProvidersBadRequest{}
}

/*
GETVasaProvidersBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GETVasaProvidersBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this get vasa providers bad request response has a 2xx status code
func (o *GETVasaProvidersBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get vasa providers bad request response has a 3xx status code
func (o *GETVasaProvidersBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get vasa providers bad request response has a 4xx status code
func (o *GETVasaProvidersBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get vasa providers bad request response has a 5xx status code
func (o *GETVasaProvidersBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get vasa providers bad request response a status code equal to that given
func (o *GETVasaProvidersBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GETVasaProvidersBadRequest) Error() string {
	return fmt.Sprintf("[GET /v1/vasa-providers][%d] getVasaProvidersBadRequest  %+v", 400, o.Payload)
}

func (o *GETVasaProvidersBadRequest) String() string {
	return fmt.Sprintf("[GET /v1/vasa-providers][%d] getVasaProvidersBadRequest  %+v", 400, o.Payload)
}

func (o *GETVasaProvidersBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *GETVasaProvidersBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGETVasaProvidersInternalServerError creates a GETVasaProvidersInternalServerError with default headers values
func NewGETVasaProvidersInternalServerError() *GETVasaProvidersInternalServerError {
	return &GETVasaProvidersInternalServerError{}
}

/*
GETVasaProvidersInternalServerError describes a response with status code 500, with default header values.

InternalServerError
*/
type GETVasaProvidersInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this get vasa providers internal server error response has a 2xx status code
func (o *GETVasaProvidersInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get vasa providers internal server error response has a 3xx status code
func (o *GETVasaProvidersInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get vasa providers internal server error response has a 4xx status code
func (o *GETVasaProvidersInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get vasa providers internal server error response has a 5xx status code
func (o *GETVasaProvidersInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get vasa providers internal server error response a status code equal to that given
func (o *GETVasaProvidersInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GETVasaProvidersInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/vasa-providers][%d] getVasaProvidersInternalServerError  %+v", 500, o.Payload)
}

func (o *GETVasaProvidersInternalServerError) String() string {
	return fmt.Sprintf("[GET /v1/vasa-providers][%d] getVasaProvidersInternalServerError  %+v", 500, o.Payload)
}

func (o *GETVasaProvidersInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *GETVasaProvidersInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
