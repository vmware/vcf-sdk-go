// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2023 VMware, Inc.
// SPDX-License-Identifier: BSD-2-Clause

package identity_providers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/vmware/vcf-sdk-go/models"
)

// GetIdentityProviderByIDReader is a Reader for the GetIdentityProviderByID structure.
type GetIdentityProviderByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetIdentityProviderByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetIdentityProviderByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetIdentityProviderByIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetIdentityProviderByIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /v1/identity-providers/{id}] getIdentityProviderById", response, response.Code())
	}
}

// NewGetIdentityProviderByIDOK creates a GetIdentityProviderByIDOK with default headers values
func NewGetIdentityProviderByIDOK() *GetIdentityProviderByIDOK {
	return &GetIdentityProviderByIDOK{}
}

/*
GetIdentityProviderByIDOK describes a response with status code 200, with default header values.

OK
*/
type GetIdentityProviderByIDOK struct {
	Payload *models.IdentityProvider
}

// IsSuccess returns true when this get identity provider by Id o k response has a 2xx status code
func (o *GetIdentityProviderByIDOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get identity provider by Id o k response has a 3xx status code
func (o *GetIdentityProviderByIDOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get identity provider by Id o k response has a 4xx status code
func (o *GetIdentityProviderByIDOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get identity provider by Id o k response has a 5xx status code
func (o *GetIdentityProviderByIDOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get identity provider by Id o k response a status code equal to that given
func (o *GetIdentityProviderByIDOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get identity provider by Id o k response
func (o *GetIdentityProviderByIDOK) Code() int {
	return 200
}

func (o *GetIdentityProviderByIDOK) Error() string {
	return fmt.Sprintf("[GET /v1/identity-providers/{id}][%d] getIdentityProviderByIdOK  %+v", 200, o.Payload)
}

func (o *GetIdentityProviderByIDOK) String() string {
	return fmt.Sprintf("[GET /v1/identity-providers/{id}][%d] getIdentityProviderByIdOK  %+v", 200, o.Payload)
}

func (o *GetIdentityProviderByIDOK) GetPayload() *models.IdentityProvider {
	return o.Payload
}

func (o *GetIdentityProviderByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IdentityProvider)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIdentityProviderByIDUnauthorized creates a GetIdentityProviderByIDUnauthorized with default headers values
func NewGetIdentityProviderByIDUnauthorized() *GetIdentityProviderByIDUnauthorized {
	return &GetIdentityProviderByIDUnauthorized{}
}

/*
GetIdentityProviderByIDUnauthorized describes a response with status code 401, with default header values.

Unauthorized Request
*/
type GetIdentityProviderByIDUnauthorized struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this get identity provider by Id unauthorized response has a 2xx status code
func (o *GetIdentityProviderByIDUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get identity provider by Id unauthorized response has a 3xx status code
func (o *GetIdentityProviderByIDUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get identity provider by Id unauthorized response has a 4xx status code
func (o *GetIdentityProviderByIDUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get identity provider by Id unauthorized response has a 5xx status code
func (o *GetIdentityProviderByIDUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get identity provider by Id unauthorized response a status code equal to that given
func (o *GetIdentityProviderByIDUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get identity provider by Id unauthorized response
func (o *GetIdentityProviderByIDUnauthorized) Code() int {
	return 401
}

func (o *GetIdentityProviderByIDUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v1/identity-providers/{id}][%d] getIdentityProviderByIdUnauthorized  %+v", 401, o.Payload)
}

func (o *GetIdentityProviderByIDUnauthorized) String() string {
	return fmt.Sprintf("[GET /v1/identity-providers/{id}][%d] getIdentityProviderByIdUnauthorized  %+v", 401, o.Payload)
}

func (o *GetIdentityProviderByIDUnauthorized) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetIdentityProviderByIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIdentityProviderByIDInternalServerError creates a GetIdentityProviderByIDInternalServerError with default headers values
func NewGetIdentityProviderByIDInternalServerError() *GetIdentityProviderByIDInternalServerError {
	return &GetIdentityProviderByIDInternalServerError{}
}

/*
GetIdentityProviderByIDInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type GetIdentityProviderByIDInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this get identity provider by Id internal server error response has a 2xx status code
func (o *GetIdentityProviderByIDInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get identity provider by Id internal server error response has a 3xx status code
func (o *GetIdentityProviderByIDInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get identity provider by Id internal server error response has a 4xx status code
func (o *GetIdentityProviderByIDInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get identity provider by Id internal server error response has a 5xx status code
func (o *GetIdentityProviderByIDInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get identity provider by Id internal server error response a status code equal to that given
func (o *GetIdentityProviderByIDInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get identity provider by Id internal server error response
func (o *GetIdentityProviderByIDInternalServerError) Code() int {
	return 500
}

func (o *GetIdentityProviderByIDInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/identity-providers/{id}][%d] getIdentityProviderByIdInternalServerError  %+v", 500, o.Payload)
}

func (o *GetIdentityProviderByIDInternalServerError) String() string {
	return fmt.Sprintf("[GET /v1/identity-providers/{id}][%d] getIdentityProviderByIdInternalServerError  %+v", 500, o.Payload)
}

func (o *GetIdentityProviderByIDInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetIdentityProviderByIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
