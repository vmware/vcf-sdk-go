// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2023 VMware, Inc.
// SPDX-License-Identifier: BSD-2-Clause

package releases

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/vmware/vcf-sdk-go/models"
)

// GetReleasesReader is a Reader for the GetReleases structure.
type GetReleasesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetReleasesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetReleasesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetReleasesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetReleasesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /v1/releases] getReleases", response, response.Code())
	}
}

// NewGetReleasesOK creates a GetReleasesOK with default headers values
func NewGetReleasesOK() *GetReleasesOK {
	return &GetReleasesOK{}
}

/*
GetReleasesOK describes a response with status code 200, with default header values.

Ok
*/
type GetReleasesOK struct {
	Payload *models.PageOfRelease
}

// IsSuccess returns true when this get releases o k response has a 2xx status code
func (o *GetReleasesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get releases o k response has a 3xx status code
func (o *GetReleasesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get releases o k response has a 4xx status code
func (o *GetReleasesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get releases o k response has a 5xx status code
func (o *GetReleasesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get releases o k response a status code equal to that given
func (o *GetReleasesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get releases o k response
func (o *GetReleasesOK) Code() int {
	return 200
}

func (o *GetReleasesOK) Error() string {
	return fmt.Sprintf("[GET /v1/releases][%d] getReleasesOK  %+v", 200, o.Payload)
}

func (o *GetReleasesOK) String() string {
	return fmt.Sprintf("[GET /v1/releases][%d] getReleasesOK  %+v", 200, o.Payload)
}

func (o *GetReleasesOK) GetPayload() *models.PageOfRelease {
	return o.Payload
}

func (o *GetReleasesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PageOfRelease)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetReleasesNotFound creates a GetReleasesNotFound with default headers values
func NewGetReleasesNotFound() *GetReleasesNotFound {
	return &GetReleasesNotFound{}
}

/*
GetReleasesNotFound describes a response with status code 404, with default header values.

Release not found
*/
type GetReleasesNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this get releases not found response has a 2xx status code
func (o *GetReleasesNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get releases not found response has a 3xx status code
func (o *GetReleasesNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get releases not found response has a 4xx status code
func (o *GetReleasesNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get releases not found response has a 5xx status code
func (o *GetReleasesNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get releases not found response a status code equal to that given
func (o *GetReleasesNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get releases not found response
func (o *GetReleasesNotFound) Code() int {
	return 404
}

func (o *GetReleasesNotFound) Error() string {
	return fmt.Sprintf("[GET /v1/releases][%d] getReleasesNotFound  %+v", 404, o.Payload)
}

func (o *GetReleasesNotFound) String() string {
	return fmt.Sprintf("[GET /v1/releases][%d] getReleasesNotFound  %+v", 404, o.Payload)
}

func (o *GetReleasesNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetReleasesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetReleasesInternalServerError creates a GetReleasesInternalServerError with default headers values
func NewGetReleasesInternalServerError() *GetReleasesInternalServerError {
	return &GetReleasesInternalServerError{}
}

/*
GetReleasesInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetReleasesInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this get releases internal server error response has a 2xx status code
func (o *GetReleasesInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get releases internal server error response has a 3xx status code
func (o *GetReleasesInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get releases internal server error response has a 4xx status code
func (o *GetReleasesInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get releases internal server error response has a 5xx status code
func (o *GetReleasesInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get releases internal server error response a status code equal to that given
func (o *GetReleasesInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get releases internal server error response
func (o *GetReleasesInternalServerError) Code() int {
	return 500
}

func (o *GetReleasesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/releases][%d] getReleasesInternalServerError  %+v", 500, o.Payload)
}

func (o *GetReleasesInternalServerError) String() string {
	return fmt.Sprintf("[GET /v1/releases][%d] getReleasesInternalServerError  %+v", 500, o.Payload)
}

func (o *GetReleasesInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetReleasesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
