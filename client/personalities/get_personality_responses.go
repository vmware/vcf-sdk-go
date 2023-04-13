// Code generated by go-swagger; DO NOT EDIT.

// Copyright © 2023 VMware, Inc. All Rights Reserved.
// Code licenced under: BSD-2

package personalities

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"vcf-sdk-go/models"
)

// GETPersonalityReader is a Reader for the GETPersonality structure.
type GETPersonalityReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GETPersonalityReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGETPersonalityOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGETPersonalityNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGETPersonalityInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGETPersonalityOK creates a GETPersonalityOK with default headers values
func NewGETPersonalityOK() *GETPersonalityOK {
	return &GETPersonalityOK{}
}

/*
GETPersonalityOK describes a response with status code 200, with default header values.

OK
*/
type GETPersonalityOK struct {
	Payload *models.Personality
}

// IsSuccess returns true when this get personality o k response has a 2xx status code
func (o *GETPersonalityOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get personality o k response has a 3xx status code
func (o *GETPersonalityOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get personality o k response has a 4xx status code
func (o *GETPersonalityOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get personality o k response has a 5xx status code
func (o *GETPersonalityOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get personality o k response a status code equal to that given
func (o *GETPersonalityOK) IsCode(code int) bool {
	return code == 200
}

func (o *GETPersonalityOK) Error() string {
	return fmt.Sprintf("[GET /v1/personalities/{personalityId}][%d] getPersonalityOK  %+v", 200, o.Payload)
}

func (o *GETPersonalityOK) String() string {
	return fmt.Sprintf("[GET /v1/personalities/{personalityId}][%d] getPersonalityOK  %+v", 200, o.Payload)
}

func (o *GETPersonalityOK) GetPayload() *models.Personality {
	return o.Payload
}

func (o *GETPersonalityOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Personality)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGETPersonalityNotFound creates a GETPersonalityNotFound with default headers values
func NewGETPersonalityNotFound() *GETPersonalityNotFound {
	return &GETPersonalityNotFound{}
}

/*
GETPersonalityNotFound describes a response with status code 404, with default header values.

Personality not found
*/
type GETPersonalityNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this get personality not found response has a 2xx status code
func (o *GETPersonalityNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get personality not found response has a 3xx status code
func (o *GETPersonalityNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get personality not found response has a 4xx status code
func (o *GETPersonalityNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get personality not found response has a 5xx status code
func (o *GETPersonalityNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get personality not found response a status code equal to that given
func (o *GETPersonalityNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GETPersonalityNotFound) Error() string {
	return fmt.Sprintf("[GET /v1/personalities/{personalityId}][%d] getPersonalityNotFound  %+v", 404, o.Payload)
}

func (o *GETPersonalityNotFound) String() string {
	return fmt.Sprintf("[GET /v1/personalities/{personalityId}][%d] getPersonalityNotFound  %+v", 404, o.Payload)
}

func (o *GETPersonalityNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GETPersonalityNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGETPersonalityInternalServerError creates a GETPersonalityInternalServerError with default headers values
func NewGETPersonalityInternalServerError() *GETPersonalityInternalServerError {
	return &GETPersonalityInternalServerError{}
}

/*
GETPersonalityInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GETPersonalityInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this get personality internal server error response has a 2xx status code
func (o *GETPersonalityInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get personality internal server error response has a 3xx status code
func (o *GETPersonalityInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get personality internal server error response has a 4xx status code
func (o *GETPersonalityInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get personality internal server error response has a 5xx status code
func (o *GETPersonalityInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get personality internal server error response a status code equal to that given
func (o *GETPersonalityInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GETPersonalityInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/personalities/{personalityId}][%d] getPersonalityInternalServerError  %+v", 500, o.Payload)
}

func (o *GETPersonalityInternalServerError) String() string {
	return fmt.Sprintf("[GET /v1/personalities/{personalityId}][%d] getPersonalityInternalServerError  %+v", 500, o.Payload)
}

func (o *GETPersonalityInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *GETPersonalityInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
