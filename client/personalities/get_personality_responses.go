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

// GetPersonalityReader is a Reader for the GetPersonality structure.
type GetPersonalityReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPersonalityReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetPersonalityOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetPersonalityNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetPersonalityInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetPersonalityOK creates a GetPersonalityOK with default headers values
func NewGetPersonalityOK() *GetPersonalityOK {
	return &GetPersonalityOK{}
}

/*
GetPersonalityOK describes a response with status code 200, with default header values.

OK
*/
type GetPersonalityOK struct {
	Payload *models.Personality
}

// IsSuccess returns true when this get personality o k response has a 2xx status code
func (o *GetPersonalityOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get personality o k response has a 3xx status code
func (o *GetPersonalityOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get personality o k response has a 4xx status code
func (o *GetPersonalityOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get personality o k response has a 5xx status code
func (o *GetPersonalityOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get personality o k response a status code equal to that given
func (o *GetPersonalityOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetPersonalityOK) Error() string {
	return fmt.Sprintf("[GET /v1/personalities/{personalityId}][%d] getPersonalityOK  %+v", 200, o.Payload)
}

func (o *GetPersonalityOK) String() string {
	return fmt.Sprintf("[GET /v1/personalities/{personalityId}][%d] getPersonalityOK  %+v", 200, o.Payload)
}

func (o *GetPersonalityOK) GetPayload() *models.Personality {
	return o.Payload
}

func (o *GetPersonalityOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Personality)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPersonalityNotFound creates a GetPersonalityNotFound with default headers values
func NewGetPersonalityNotFound() *GetPersonalityNotFound {
	return &GetPersonalityNotFound{}
}

/*
GetPersonalityNotFound describes a response with status code 404, with default header values.

Personality not found
*/
type GetPersonalityNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this get personality not found response has a 2xx status code
func (o *GetPersonalityNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get personality not found response has a 3xx status code
func (o *GetPersonalityNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get personality not found response has a 4xx status code
func (o *GetPersonalityNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get personality not found response has a 5xx status code
func (o *GetPersonalityNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get personality not found response a status code equal to that given
func (o *GetPersonalityNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetPersonalityNotFound) Error() string {
	return fmt.Sprintf("[GET /v1/personalities/{personalityId}][%d] getPersonalityNotFound  %+v", 404, o.Payload)
}

func (o *GetPersonalityNotFound) String() string {
	return fmt.Sprintf("[GET /v1/personalities/{personalityId}][%d] getPersonalityNotFound  %+v", 404, o.Payload)
}

func (o *GetPersonalityNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetPersonalityNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPersonalityInternalServerError creates a GetPersonalityInternalServerError with default headers values
func NewGetPersonalityInternalServerError() *GetPersonalityInternalServerError {
	return &GetPersonalityInternalServerError{}
}

/*
GetPersonalityInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetPersonalityInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this get personality internal server error response has a 2xx status code
func (o *GetPersonalityInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get personality internal server error response has a 3xx status code
func (o *GetPersonalityInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get personality internal server error response has a 4xx status code
func (o *GetPersonalityInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get personality internal server error response has a 5xx status code
func (o *GetPersonalityInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get personality internal server error response a status code equal to that given
func (o *GetPersonalityInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetPersonalityInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/personalities/{personalityId}][%d] getPersonalityInternalServerError  %+v", 500, o.Payload)
}

func (o *GetPersonalityInternalServerError) String() string {
	return fmt.Sprintf("[GET /v1/personalities/{personalityId}][%d] getPersonalityInternalServerError  %+v", 500, o.Payload)
}

func (o *GetPersonalityInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetPersonalityInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
