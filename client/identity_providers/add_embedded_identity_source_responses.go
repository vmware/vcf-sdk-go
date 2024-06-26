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

// AddEmbeddedIdentitySourceReader is a Reader for the AddEmbeddedIdentitySource structure.
type AddEmbeddedIdentitySourceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddEmbeddedIdentitySourceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAddEmbeddedIdentitySourceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewAddEmbeddedIdentitySourceNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewAddEmbeddedIdentitySourceBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewAddEmbeddedIdentitySourceNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewAddEmbeddedIdentitySourceInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /v1/identity-providers/{id}/identity-sources] addEmbeddedIdentitySource", response, response.Code())
	}
}

// NewAddEmbeddedIdentitySourceOK creates a AddEmbeddedIdentitySourceOK with default headers values
func NewAddEmbeddedIdentitySourceOK() *AddEmbeddedIdentitySourceOK {
	return &AddEmbeddedIdentitySourceOK{}
}

/*
AddEmbeddedIdentitySourceOK describes a response with status code 200, with default header values.

OK
*/
type AddEmbeddedIdentitySourceOK struct {
	Payload interface{}
}

// IsSuccess returns true when this add embedded identity source o k response has a 2xx status code
func (o *AddEmbeddedIdentitySourceOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this add embedded identity source o k response has a 3xx status code
func (o *AddEmbeddedIdentitySourceOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this add embedded identity source o k response has a 4xx status code
func (o *AddEmbeddedIdentitySourceOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this add embedded identity source o k response has a 5xx status code
func (o *AddEmbeddedIdentitySourceOK) IsServerError() bool {
	return false
}

// IsCode returns true when this add embedded identity source o k response a status code equal to that given
func (o *AddEmbeddedIdentitySourceOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the add embedded identity source o k response
func (o *AddEmbeddedIdentitySourceOK) Code() int {
	return 200
}

func (o *AddEmbeddedIdentitySourceOK) Error() string {
	return fmt.Sprintf("[POST /v1/identity-providers/{id}/identity-sources][%d] addEmbeddedIdentitySourceOK  %+v", 200, o.Payload)
}

func (o *AddEmbeddedIdentitySourceOK) String() string {
	return fmt.Sprintf("[POST /v1/identity-providers/{id}/identity-sources][%d] addEmbeddedIdentitySourceOK  %+v", 200, o.Payload)
}

func (o *AddEmbeddedIdentitySourceOK) GetPayload() interface{} {
	return o.Payload
}

func (o *AddEmbeddedIdentitySourceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddEmbeddedIdentitySourceNoContent creates a AddEmbeddedIdentitySourceNoContent with default headers values
func NewAddEmbeddedIdentitySourceNoContent() *AddEmbeddedIdentitySourceNoContent {
	return &AddEmbeddedIdentitySourceNoContent{}
}

/*
AddEmbeddedIdentitySourceNoContent describes a response with status code 204, with default header values.

No content
*/
type AddEmbeddedIdentitySourceNoContent struct {
	Payload interface{}
}

// IsSuccess returns true when this add embedded identity source no content response has a 2xx status code
func (o *AddEmbeddedIdentitySourceNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this add embedded identity source no content response has a 3xx status code
func (o *AddEmbeddedIdentitySourceNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this add embedded identity source no content response has a 4xx status code
func (o *AddEmbeddedIdentitySourceNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this add embedded identity source no content response has a 5xx status code
func (o *AddEmbeddedIdentitySourceNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this add embedded identity source no content response a status code equal to that given
func (o *AddEmbeddedIdentitySourceNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the add embedded identity source no content response
func (o *AddEmbeddedIdentitySourceNoContent) Code() int {
	return 204
}

func (o *AddEmbeddedIdentitySourceNoContent) Error() string {
	return fmt.Sprintf("[POST /v1/identity-providers/{id}/identity-sources][%d] addEmbeddedIdentitySourceNoContent  %+v", 204, o.Payload)
}

func (o *AddEmbeddedIdentitySourceNoContent) String() string {
	return fmt.Sprintf("[POST /v1/identity-providers/{id}/identity-sources][%d] addEmbeddedIdentitySourceNoContent  %+v", 204, o.Payload)
}

func (o *AddEmbeddedIdentitySourceNoContent) GetPayload() interface{} {
	return o.Payload
}

func (o *AddEmbeddedIdentitySourceNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddEmbeddedIdentitySourceBadRequest creates a AddEmbeddedIdentitySourceBadRequest with default headers values
func NewAddEmbeddedIdentitySourceBadRequest() *AddEmbeddedIdentitySourceBadRequest {
	return &AddEmbeddedIdentitySourceBadRequest{}
}

/*
AddEmbeddedIdentitySourceBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type AddEmbeddedIdentitySourceBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this add embedded identity source bad request response has a 2xx status code
func (o *AddEmbeddedIdentitySourceBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this add embedded identity source bad request response has a 3xx status code
func (o *AddEmbeddedIdentitySourceBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this add embedded identity source bad request response has a 4xx status code
func (o *AddEmbeddedIdentitySourceBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this add embedded identity source bad request response has a 5xx status code
func (o *AddEmbeddedIdentitySourceBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this add embedded identity source bad request response a status code equal to that given
func (o *AddEmbeddedIdentitySourceBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the add embedded identity source bad request response
func (o *AddEmbeddedIdentitySourceBadRequest) Code() int {
	return 400
}

func (o *AddEmbeddedIdentitySourceBadRequest) Error() string {
	return fmt.Sprintf("[POST /v1/identity-providers/{id}/identity-sources][%d] addEmbeddedIdentitySourceBadRequest  %+v", 400, o.Payload)
}

func (o *AddEmbeddedIdentitySourceBadRequest) String() string {
	return fmt.Sprintf("[POST /v1/identity-providers/{id}/identity-sources][%d] addEmbeddedIdentitySourceBadRequest  %+v", 400, o.Payload)
}

func (o *AddEmbeddedIdentitySourceBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *AddEmbeddedIdentitySourceBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddEmbeddedIdentitySourceNotFound creates a AddEmbeddedIdentitySourceNotFound with default headers values
func NewAddEmbeddedIdentitySourceNotFound() *AddEmbeddedIdentitySourceNotFound {
	return &AddEmbeddedIdentitySourceNotFound{}
}

/*
AddEmbeddedIdentitySourceNotFound describes a response with status code 404, with default header values.

Identity Provider not found
*/
type AddEmbeddedIdentitySourceNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this add embedded identity source not found response has a 2xx status code
func (o *AddEmbeddedIdentitySourceNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this add embedded identity source not found response has a 3xx status code
func (o *AddEmbeddedIdentitySourceNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this add embedded identity source not found response has a 4xx status code
func (o *AddEmbeddedIdentitySourceNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this add embedded identity source not found response has a 5xx status code
func (o *AddEmbeddedIdentitySourceNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this add embedded identity source not found response a status code equal to that given
func (o *AddEmbeddedIdentitySourceNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the add embedded identity source not found response
func (o *AddEmbeddedIdentitySourceNotFound) Code() int {
	return 404
}

func (o *AddEmbeddedIdentitySourceNotFound) Error() string {
	return fmt.Sprintf("[POST /v1/identity-providers/{id}/identity-sources][%d] addEmbeddedIdentitySourceNotFound  %+v", 404, o.Payload)
}

func (o *AddEmbeddedIdentitySourceNotFound) String() string {
	return fmt.Sprintf("[POST /v1/identity-providers/{id}/identity-sources][%d] addEmbeddedIdentitySourceNotFound  %+v", 404, o.Payload)
}

func (o *AddEmbeddedIdentitySourceNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *AddEmbeddedIdentitySourceNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddEmbeddedIdentitySourceInternalServerError creates a AddEmbeddedIdentitySourceInternalServerError with default headers values
func NewAddEmbeddedIdentitySourceInternalServerError() *AddEmbeddedIdentitySourceInternalServerError {
	return &AddEmbeddedIdentitySourceInternalServerError{}
}

/*
AddEmbeddedIdentitySourceInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type AddEmbeddedIdentitySourceInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this add embedded identity source internal server error response has a 2xx status code
func (o *AddEmbeddedIdentitySourceInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this add embedded identity source internal server error response has a 3xx status code
func (o *AddEmbeddedIdentitySourceInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this add embedded identity source internal server error response has a 4xx status code
func (o *AddEmbeddedIdentitySourceInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this add embedded identity source internal server error response has a 5xx status code
func (o *AddEmbeddedIdentitySourceInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this add embedded identity source internal server error response a status code equal to that given
func (o *AddEmbeddedIdentitySourceInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the add embedded identity source internal server error response
func (o *AddEmbeddedIdentitySourceInternalServerError) Code() int {
	return 500
}

func (o *AddEmbeddedIdentitySourceInternalServerError) Error() string {
	return fmt.Sprintf("[POST /v1/identity-providers/{id}/identity-sources][%d] addEmbeddedIdentitySourceInternalServerError  %+v", 500, o.Payload)
}

func (o *AddEmbeddedIdentitySourceInternalServerError) String() string {
	return fmt.Sprintf("[POST /v1/identity-providers/{id}/identity-sources][%d] addEmbeddedIdentitySourceInternalServerError  %+v", 500, o.Payload)
}

func (o *AddEmbeddedIdentitySourceInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *AddEmbeddedIdentitySourceInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
