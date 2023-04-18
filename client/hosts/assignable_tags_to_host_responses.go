// Code generated by go-swagger; DO NOT EDIT.

// Copyright © 2023 VMware, Inc. All Rights Reserved.
// Code licenced under: BSD-2

package hosts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/vmware/vcf-sdk-go/models"
)

// AssignableTagsToHostReader is a Reader for the AssignableTagsToHost structure.
type AssignableTagsToHostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AssignableTagsToHostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAssignableTagsToHostOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewAssignableTagsToHostBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewAssignableTagsToHostInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewAssignableTagsToHostOK creates a AssignableTagsToHostOK with default headers values
func NewAssignableTagsToHostOK() *AssignableTagsToHostOK {
	return &AssignableTagsToHostOK{}
}

/*
AssignableTagsToHostOK describes a response with status code 200, with default header values.

Ok
*/
type AssignableTagsToHostOK struct {
	Payload *models.PageOfTag
}

// IsSuccess returns true when this assignable tags to host o k response has a 2xx status code
func (o *AssignableTagsToHostOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this assignable tags to host o k response has a 3xx status code
func (o *AssignableTagsToHostOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this assignable tags to host o k response has a 4xx status code
func (o *AssignableTagsToHostOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this assignable tags to host o k response has a 5xx status code
func (o *AssignableTagsToHostOK) IsServerError() bool {
	return false
}

// IsCode returns true when this assignable tags to host o k response a status code equal to that given
func (o *AssignableTagsToHostOK) IsCode(code int) bool {
	return code == 200
}

func (o *AssignableTagsToHostOK) Error() string {
	return fmt.Sprintf("[GET /v1/hosts/{id}/tags/assignable-tags][%d] assignableTagsToHostOK  %+v", 200, o.Payload)
}

func (o *AssignableTagsToHostOK) String() string {
	return fmt.Sprintf("[GET /v1/hosts/{id}/tags/assignable-tags][%d] assignableTagsToHostOK  %+v", 200, o.Payload)
}

func (o *AssignableTagsToHostOK) GetPayload() *models.PageOfTag {
	return o.Payload
}

func (o *AssignableTagsToHostOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PageOfTag)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAssignableTagsToHostBadRequest creates a AssignableTagsToHostBadRequest with default headers values
func NewAssignableTagsToHostBadRequest() *AssignableTagsToHostBadRequest {
	return &AssignableTagsToHostBadRequest{}
}

/*
AssignableTagsToHostBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type AssignableTagsToHostBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this assignable tags to host bad request response has a 2xx status code
func (o *AssignableTagsToHostBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this assignable tags to host bad request response has a 3xx status code
func (o *AssignableTagsToHostBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this assignable tags to host bad request response has a 4xx status code
func (o *AssignableTagsToHostBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this assignable tags to host bad request response has a 5xx status code
func (o *AssignableTagsToHostBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this assignable tags to host bad request response a status code equal to that given
func (o *AssignableTagsToHostBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *AssignableTagsToHostBadRequest) Error() string {
	return fmt.Sprintf("[GET /v1/hosts/{id}/tags/assignable-tags][%d] assignableTagsToHostBadRequest  %+v", 400, o.Payload)
}

func (o *AssignableTagsToHostBadRequest) String() string {
	return fmt.Sprintf("[GET /v1/hosts/{id}/tags/assignable-tags][%d] assignableTagsToHostBadRequest  %+v", 400, o.Payload)
}

func (o *AssignableTagsToHostBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *AssignableTagsToHostBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAssignableTagsToHostInternalServerError creates a AssignableTagsToHostInternalServerError with default headers values
func NewAssignableTagsToHostInternalServerError() *AssignableTagsToHostInternalServerError {
	return &AssignableTagsToHostInternalServerError{}
}

/*
AssignableTagsToHostInternalServerError describes a response with status code 500, with default header values.

InternalServerError
*/
type AssignableTagsToHostInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this assignable tags to host internal server error response has a 2xx status code
func (o *AssignableTagsToHostInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this assignable tags to host internal server error response has a 3xx status code
func (o *AssignableTagsToHostInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this assignable tags to host internal server error response has a 4xx status code
func (o *AssignableTagsToHostInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this assignable tags to host internal server error response has a 5xx status code
func (o *AssignableTagsToHostInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this assignable tags to host internal server error response a status code equal to that given
func (o *AssignableTagsToHostInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *AssignableTagsToHostInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/hosts/{id}/tags/assignable-tags][%d] assignableTagsToHostInternalServerError  %+v", 500, o.Payload)
}

func (o *AssignableTagsToHostInternalServerError) String() string {
	return fmt.Sprintf("[GET /v1/hosts/{id}/tags/assignable-tags][%d] assignableTagsToHostInternalServerError  %+v", 500, o.Payload)
}

func (o *AssignableTagsToHostInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *AssignableTagsToHostInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
