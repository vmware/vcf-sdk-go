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

// GetIdentityProvidersReader is a Reader for the GetIdentityProviders structure.
type GetIdentityProvidersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetIdentityProvidersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetIdentityProvidersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetIdentityProvidersBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetIdentityProvidersUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetIdentityProvidersInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /v1/identity-providers] getIdentityProviders", response, response.Code())
	}
}

// NewGetIdentityProvidersOK creates a GetIdentityProvidersOK with default headers values
func NewGetIdentityProvidersOK() *GetIdentityProvidersOK {
	return &GetIdentityProvidersOK{}
}

/*
GetIdentityProvidersOK describes a response with status code 200, with default header values.

OK
*/
type GetIdentityProvidersOK struct {
	Payload *models.PageOfIdentityProvider
}

// IsSuccess returns true when this get identity providers o k response has a 2xx status code
func (o *GetIdentityProvidersOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get identity providers o k response has a 3xx status code
func (o *GetIdentityProvidersOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get identity providers o k response has a 4xx status code
func (o *GetIdentityProvidersOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get identity providers o k response has a 5xx status code
func (o *GetIdentityProvidersOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get identity providers o k response a status code equal to that given
func (o *GetIdentityProvidersOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get identity providers o k response
func (o *GetIdentityProvidersOK) Code() int {
	return 200
}

func (o *GetIdentityProvidersOK) Error() string {
	return fmt.Sprintf("[GET /v1/identity-providers][%d] getIdentityProvidersOK  %+v", 200, o.Payload)
}

func (o *GetIdentityProvidersOK) String() string {
	return fmt.Sprintf("[GET /v1/identity-providers][%d] getIdentityProvidersOK  %+v", 200, o.Payload)
}

func (o *GetIdentityProvidersOK) GetPayload() *models.PageOfIdentityProvider {
	return o.Payload
}

func (o *GetIdentityProvidersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PageOfIdentityProvider)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIdentityProvidersBadRequest creates a GetIdentityProvidersBadRequest with default headers values
func NewGetIdentityProvidersBadRequest() *GetIdentityProvidersBadRequest {
	return &GetIdentityProvidersBadRequest{}
}

/*
GetIdentityProvidersBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type GetIdentityProvidersBadRequest struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this get identity providers bad request response has a 2xx status code
func (o *GetIdentityProvidersBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get identity providers bad request response has a 3xx status code
func (o *GetIdentityProvidersBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get identity providers bad request response has a 4xx status code
func (o *GetIdentityProvidersBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get identity providers bad request response has a 5xx status code
func (o *GetIdentityProvidersBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get identity providers bad request response a status code equal to that given
func (o *GetIdentityProvidersBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get identity providers bad request response
func (o *GetIdentityProvidersBadRequest) Code() int {
	return 400
}

func (o *GetIdentityProvidersBadRequest) Error() string {
	return fmt.Sprintf("[GET /v1/identity-providers][%d] getIdentityProvidersBadRequest  %+v", 400, o.Payload)
}

func (o *GetIdentityProvidersBadRequest) String() string {
	return fmt.Sprintf("[GET /v1/identity-providers][%d] getIdentityProvidersBadRequest  %+v", 400, o.Payload)
}

func (o *GetIdentityProvidersBadRequest) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetIdentityProvidersBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIdentityProvidersUnauthorized creates a GetIdentityProvidersUnauthorized with default headers values
func NewGetIdentityProvidersUnauthorized() *GetIdentityProvidersUnauthorized {
	return &GetIdentityProvidersUnauthorized{}
}

/*
GetIdentityProvidersUnauthorized describes a response with status code 401, with default header values.

Unauthorized Request
*/
type GetIdentityProvidersUnauthorized struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this get identity providers unauthorized response has a 2xx status code
func (o *GetIdentityProvidersUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get identity providers unauthorized response has a 3xx status code
func (o *GetIdentityProvidersUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get identity providers unauthorized response has a 4xx status code
func (o *GetIdentityProvidersUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get identity providers unauthorized response has a 5xx status code
func (o *GetIdentityProvidersUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get identity providers unauthorized response a status code equal to that given
func (o *GetIdentityProvidersUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get identity providers unauthorized response
func (o *GetIdentityProvidersUnauthorized) Code() int {
	return 401
}

func (o *GetIdentityProvidersUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v1/identity-providers][%d] getIdentityProvidersUnauthorized  %+v", 401, o.Payload)
}

func (o *GetIdentityProvidersUnauthorized) String() string {
	return fmt.Sprintf("[GET /v1/identity-providers][%d] getIdentityProvidersUnauthorized  %+v", 401, o.Payload)
}

func (o *GetIdentityProvidersUnauthorized) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetIdentityProvidersUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIdentityProvidersInternalServerError creates a GetIdentityProvidersInternalServerError with default headers values
func NewGetIdentityProvidersInternalServerError() *GetIdentityProvidersInternalServerError {
	return &GetIdentityProvidersInternalServerError{}
}

/*
GetIdentityProvidersInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetIdentityProvidersInternalServerError struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this get identity providers internal server error response has a 2xx status code
func (o *GetIdentityProvidersInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get identity providers internal server error response has a 3xx status code
func (o *GetIdentityProvidersInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get identity providers internal server error response has a 4xx status code
func (o *GetIdentityProvidersInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get identity providers internal server error response has a 5xx status code
func (o *GetIdentityProvidersInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get identity providers internal server error response a status code equal to that given
func (o *GetIdentityProvidersInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get identity providers internal server error response
func (o *GetIdentityProvidersInternalServerError) Code() int {
	return 500
}

func (o *GetIdentityProvidersInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/identity-providers][%d] getIdentityProvidersInternalServerError  %+v", 500, o.Payload)
}

func (o *GetIdentityProvidersInternalServerError) String() string {
	return fmt.Sprintf("[GET /v1/identity-providers][%d] getIdentityProvidersInternalServerError  %+v", 500, o.Payload)
}

func (o *GetIdentityProvidersInternalServerError) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetIdentityProvidersInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
