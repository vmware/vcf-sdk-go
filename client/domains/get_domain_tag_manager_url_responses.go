// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2023 VMware, Inc.
// SPDX-License-Identifier: BSD-2-Clause

package domains

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/vmware/vcf-sdk-go/models"
)

// GETDomainTagManagerURLReader is a Reader for the GETDomainTagManagerURL structure.
type GETDomainTagManagerURLReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GETDomainTagManagerURLReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGETDomainTagManagerURLOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGETDomainTagManagerURLBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGETDomainTagManagerURLInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGETDomainTagManagerURLOK creates a GETDomainTagManagerURLOK with default headers values
func NewGETDomainTagManagerURLOK() *GETDomainTagManagerURLOK {
	return &GETDomainTagManagerURLOK{}
}

/*
GETDomainTagManagerURLOK describes a response with status code 200, with default header values.

Ok
*/
type GETDomainTagManagerURLOK struct {
	Payload *models.TagManagerModel
}

// IsSuccess returns true when this get domain tag manager Url o k response has a 2xx status code
func (o *GETDomainTagManagerURLOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get domain tag manager Url o k response has a 3xx status code
func (o *GETDomainTagManagerURLOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get domain tag manager Url o k response has a 4xx status code
func (o *GETDomainTagManagerURLOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get domain tag manager Url o k response has a 5xx status code
func (o *GETDomainTagManagerURLOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get domain tag manager Url o k response a status code equal to that given
func (o *GETDomainTagManagerURLOK) IsCode(code int) bool {
	return code == 200
}

func (o *GETDomainTagManagerURLOK) Error() string {
	return fmt.Sprintf("[GET /v1/domains/{id}/tags/tag-manager][%d] getDomainTagManagerUrlOK  %+v", 200, o.Payload)
}

func (o *GETDomainTagManagerURLOK) String() string {
	return fmt.Sprintf("[GET /v1/domains/{id}/tags/tag-manager][%d] getDomainTagManagerUrlOK  %+v", 200, o.Payload)
}

func (o *GETDomainTagManagerURLOK) GetPayload() *models.TagManagerModel {
	return o.Payload
}

func (o *GETDomainTagManagerURLOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TagManagerModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGETDomainTagManagerURLBadRequest creates a GETDomainTagManagerURLBadRequest with default headers values
func NewGETDomainTagManagerURLBadRequest() *GETDomainTagManagerURLBadRequest {
	return &GETDomainTagManagerURLBadRequest{}
}

/*
GETDomainTagManagerURLBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GETDomainTagManagerURLBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this get domain tag manager Url bad request response has a 2xx status code
func (o *GETDomainTagManagerURLBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get domain tag manager Url bad request response has a 3xx status code
func (o *GETDomainTagManagerURLBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get domain tag manager Url bad request response has a 4xx status code
func (o *GETDomainTagManagerURLBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get domain tag manager Url bad request response has a 5xx status code
func (o *GETDomainTagManagerURLBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get domain tag manager Url bad request response a status code equal to that given
func (o *GETDomainTagManagerURLBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GETDomainTagManagerURLBadRequest) Error() string {
	return fmt.Sprintf("[GET /v1/domains/{id}/tags/tag-manager][%d] getDomainTagManagerUrlBadRequest  %+v", 400, o.Payload)
}

func (o *GETDomainTagManagerURLBadRequest) String() string {
	return fmt.Sprintf("[GET /v1/domains/{id}/tags/tag-manager][%d] getDomainTagManagerUrlBadRequest  %+v", 400, o.Payload)
}

func (o *GETDomainTagManagerURLBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *GETDomainTagManagerURLBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGETDomainTagManagerURLInternalServerError creates a GETDomainTagManagerURLInternalServerError with default headers values
func NewGETDomainTagManagerURLInternalServerError() *GETDomainTagManagerURLInternalServerError {
	return &GETDomainTagManagerURLInternalServerError{}
}

/*
GETDomainTagManagerURLInternalServerError describes a response with status code 500, with default header values.

InternalServerError
*/
type GETDomainTagManagerURLInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this get domain tag manager Url internal server error response has a 2xx status code
func (o *GETDomainTagManagerURLInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get domain tag manager Url internal server error response has a 3xx status code
func (o *GETDomainTagManagerURLInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get domain tag manager Url internal server error response has a 4xx status code
func (o *GETDomainTagManagerURLInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get domain tag manager Url internal server error response has a 5xx status code
func (o *GETDomainTagManagerURLInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get domain tag manager Url internal server error response a status code equal to that given
func (o *GETDomainTagManagerURLInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GETDomainTagManagerURLInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/domains/{id}/tags/tag-manager][%d] getDomainTagManagerUrlInternalServerError  %+v", 500, o.Payload)
}

func (o *GETDomainTagManagerURLInternalServerError) String() string {
	return fmt.Sprintf("[GET /v1/domains/{id}/tags/tag-manager][%d] getDomainTagManagerUrlInternalServerError  %+v", 500, o.Payload)
}

func (o *GETDomainTagManagerURLInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *GETDomainTagManagerURLInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
