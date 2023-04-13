// Code generated by go-swagger; DO NOT EDIT.

// Copyright © 2023 VMware, Inc. All Rights Reserved.
// Code licenced under: BSD-2

package identity_providers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"vcf-sdk-go/models"
)

// GETIdentityProviderByIDReader is a Reader for the GETIdentityProviderByID structure.
type GETIdentityProviderByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GETIdentityProviderByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGETIdentityProviderByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGETIdentityProviderByIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGETIdentityProviderByIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGETIdentityProviderByIDOK creates a GETIdentityProviderByIDOK with default headers values
func NewGETIdentityProviderByIDOK() *GETIdentityProviderByIDOK {
	return &GETIdentityProviderByIDOK{}
}

/*
GETIdentityProviderByIDOK describes a response with status code 200, with default header values.

OK
*/
type GETIdentityProviderByIDOK struct {
	Payload *models.IdentityProvider
}

// IsSuccess returns true when this get identity provider by Id o k response has a 2xx status code
func (o *GETIdentityProviderByIDOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get identity provider by Id o k response has a 3xx status code
func (o *GETIdentityProviderByIDOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get identity provider by Id o k response has a 4xx status code
func (o *GETIdentityProviderByIDOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get identity provider by Id o k response has a 5xx status code
func (o *GETIdentityProviderByIDOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get identity provider by Id o k response a status code equal to that given
func (o *GETIdentityProviderByIDOK) IsCode(code int) bool {
	return code == 200
}

func (o *GETIdentityProviderByIDOK) Error() string {
	return fmt.Sprintf("[GET /v1/identity-providers/{id}][%d] getIdentityProviderByIdOK  %+v", 200, o.Payload)
}

func (o *GETIdentityProviderByIDOK) String() string {
	return fmt.Sprintf("[GET /v1/identity-providers/{id}][%d] getIdentityProviderByIdOK  %+v", 200, o.Payload)
}

func (o *GETIdentityProviderByIDOK) GetPayload() *models.IdentityProvider {
	return o.Payload
}

func (o *GETIdentityProviderByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IdentityProvider)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGETIdentityProviderByIDUnauthorized creates a GETIdentityProviderByIDUnauthorized with default headers values
func NewGETIdentityProviderByIDUnauthorized() *GETIdentityProviderByIDUnauthorized {
	return &GETIdentityProviderByIDUnauthorized{}
}

/*
GETIdentityProviderByIDUnauthorized describes a response with status code 401, with default header values.

Unauthorized Request
*/
type GETIdentityProviderByIDUnauthorized struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this get identity provider by Id unauthorized response has a 2xx status code
func (o *GETIdentityProviderByIDUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get identity provider by Id unauthorized response has a 3xx status code
func (o *GETIdentityProviderByIDUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get identity provider by Id unauthorized response has a 4xx status code
func (o *GETIdentityProviderByIDUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get identity provider by Id unauthorized response has a 5xx status code
func (o *GETIdentityProviderByIDUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get identity provider by Id unauthorized response a status code equal to that given
func (o *GETIdentityProviderByIDUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GETIdentityProviderByIDUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v1/identity-providers/{id}][%d] getIdentityProviderByIdUnauthorized  %+v", 401, o.Payload)
}

func (o *GETIdentityProviderByIDUnauthorized) String() string {
	return fmt.Sprintf("[GET /v1/identity-providers/{id}][%d] getIdentityProviderByIdUnauthorized  %+v", 401, o.Payload)
}

func (o *GETIdentityProviderByIDUnauthorized) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GETIdentityProviderByIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGETIdentityProviderByIDInternalServerError creates a GETIdentityProviderByIDInternalServerError with default headers values
func NewGETIdentityProviderByIDInternalServerError() *GETIdentityProviderByIDInternalServerError {
	return &GETIdentityProviderByIDInternalServerError{}
}

/*
GETIdentityProviderByIDInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type GETIdentityProviderByIDInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this get identity provider by Id internal server error response has a 2xx status code
func (o *GETIdentityProviderByIDInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get identity provider by Id internal server error response has a 3xx status code
func (o *GETIdentityProviderByIDInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get identity provider by Id internal server error response has a 4xx status code
func (o *GETIdentityProviderByIDInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get identity provider by Id internal server error response has a 5xx status code
func (o *GETIdentityProviderByIDInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get identity provider by Id internal server error response a status code equal to that given
func (o *GETIdentityProviderByIDInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GETIdentityProviderByIDInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/identity-providers/{id}][%d] getIdentityProviderByIdInternalServerError  %+v", 500, o.Payload)
}

func (o *GETIdentityProviderByIDInternalServerError) String() string {
	return fmt.Sprintf("[GET /v1/identity-providers/{id}][%d] getIdentityProviderByIdInternalServerError  %+v", 500, o.Payload)
}

func (o *GETIdentityProviderByIDInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *GETIdentityProviderByIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
