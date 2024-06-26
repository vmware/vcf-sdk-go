// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2023 VMware, Inc.
// SPDX-License-Identifier: BSD-2-Clause

package personalities

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/vmware/vcf-sdk-go/models"
)

// DeletePersonalityReader is a Reader for the DeletePersonality structure.
type DeletePersonalityReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeletePersonalityReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeletePersonalityOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewDeletePersonalityNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeletePersonalityInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[DELETE /v1/personalities] deletePersonality", response, response.Code())
	}
}

// NewDeletePersonalityOK creates a DeletePersonalityOK with default headers values
func NewDeletePersonalityOK() *DeletePersonalityOK {
	return &DeletePersonalityOK{}
}

/*
DeletePersonalityOK describes a response with status code 200, with default header values.

OK
*/
type DeletePersonalityOK struct {
}

// IsSuccess returns true when this delete personality o k response has a 2xx status code
func (o *DeletePersonalityOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete personality o k response has a 3xx status code
func (o *DeletePersonalityOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete personality o k response has a 4xx status code
func (o *DeletePersonalityOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete personality o k response has a 5xx status code
func (o *DeletePersonalityOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete personality o k response a status code equal to that given
func (o *DeletePersonalityOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the delete personality o k response
func (o *DeletePersonalityOK) Code() int {
	return 200
}

func (o *DeletePersonalityOK) Error() string {
	return fmt.Sprintf("[DELETE /v1/personalities][%d] deletePersonalityOK ", 200)
}

func (o *DeletePersonalityOK) String() string {
	return fmt.Sprintf("[DELETE /v1/personalities][%d] deletePersonalityOK ", 200)
}

func (o *DeletePersonalityOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeletePersonalityNotFound creates a DeletePersonalityNotFound with default headers values
func NewDeletePersonalityNotFound() *DeletePersonalityNotFound {
	return &DeletePersonalityNotFound{}
}

/*
DeletePersonalityNotFound describes a response with status code 404, with default header values.

Personality by name or id not found
*/
type DeletePersonalityNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this delete personality not found response has a 2xx status code
func (o *DeletePersonalityNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete personality not found response has a 3xx status code
func (o *DeletePersonalityNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete personality not found response has a 4xx status code
func (o *DeletePersonalityNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete personality not found response has a 5xx status code
func (o *DeletePersonalityNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this delete personality not found response a status code equal to that given
func (o *DeletePersonalityNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the delete personality not found response
func (o *DeletePersonalityNotFound) Code() int {
	return 404
}

func (o *DeletePersonalityNotFound) Error() string {
	return fmt.Sprintf("[DELETE /v1/personalities][%d] deletePersonalityNotFound  %+v", 404, o.Payload)
}

func (o *DeletePersonalityNotFound) String() string {
	return fmt.Sprintf("[DELETE /v1/personalities][%d] deletePersonalityNotFound  %+v", 404, o.Payload)
}

func (o *DeletePersonalityNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeletePersonalityNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeletePersonalityInternalServerError creates a DeletePersonalityInternalServerError with default headers values
func NewDeletePersonalityInternalServerError() *DeletePersonalityInternalServerError {
	return &DeletePersonalityInternalServerError{}
}

/*
DeletePersonalityInternalServerError describes a response with status code 500, with default header values.

Unexpected error
*/
type DeletePersonalityInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this delete personality internal server error response has a 2xx status code
func (o *DeletePersonalityInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete personality internal server error response has a 3xx status code
func (o *DeletePersonalityInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete personality internal server error response has a 4xx status code
func (o *DeletePersonalityInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete personality internal server error response has a 5xx status code
func (o *DeletePersonalityInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this delete personality internal server error response a status code equal to that given
func (o *DeletePersonalityInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the delete personality internal server error response
func (o *DeletePersonalityInternalServerError) Code() int {
	return 500
}

func (o *DeletePersonalityInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /v1/personalities][%d] deletePersonalityInternalServerError  %+v", 500, o.Payload)
}

func (o *DeletePersonalityInternalServerError) String() string {
	return fmt.Sprintf("[DELETE /v1/personalities][%d] deletePersonalityInternalServerError  %+v", 500, o.Payload)
}

func (o *DeletePersonalityInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeletePersonalityInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
