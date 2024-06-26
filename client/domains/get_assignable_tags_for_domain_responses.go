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

// GetAssignableTagsForDomainReader is a Reader for the GetAssignableTagsForDomain structure.
type GetAssignableTagsForDomainReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAssignableTagsForDomainReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAssignableTagsForDomainOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetAssignableTagsForDomainBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetAssignableTagsForDomainInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /v1/domains/{id}/tags/assignable-tags] getAssignableTagsForDomain", response, response.Code())
	}
}

// NewGetAssignableTagsForDomainOK creates a GetAssignableTagsForDomainOK with default headers values
func NewGetAssignableTagsForDomainOK() *GetAssignableTagsForDomainOK {
	return &GetAssignableTagsForDomainOK{}
}

/*
GetAssignableTagsForDomainOK describes a response with status code 200, with default header values.

Ok
*/
type GetAssignableTagsForDomainOK struct {
	Payload *models.PageOfTag
}

// IsSuccess returns true when this get assignable tags for domain o k response has a 2xx status code
func (o *GetAssignableTagsForDomainOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get assignable tags for domain o k response has a 3xx status code
func (o *GetAssignableTagsForDomainOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get assignable tags for domain o k response has a 4xx status code
func (o *GetAssignableTagsForDomainOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get assignable tags for domain o k response has a 5xx status code
func (o *GetAssignableTagsForDomainOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get assignable tags for domain o k response a status code equal to that given
func (o *GetAssignableTagsForDomainOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get assignable tags for domain o k response
func (o *GetAssignableTagsForDomainOK) Code() int {
	return 200
}

func (o *GetAssignableTagsForDomainOK) Error() string {
	return fmt.Sprintf("[GET /v1/domains/{id}/tags/assignable-tags][%d] getAssignableTagsForDomainOK  %+v", 200, o.Payload)
}

func (o *GetAssignableTagsForDomainOK) String() string {
	return fmt.Sprintf("[GET /v1/domains/{id}/tags/assignable-tags][%d] getAssignableTagsForDomainOK  %+v", 200, o.Payload)
}

func (o *GetAssignableTagsForDomainOK) GetPayload() *models.PageOfTag {
	return o.Payload
}

func (o *GetAssignableTagsForDomainOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PageOfTag)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAssignableTagsForDomainBadRequest creates a GetAssignableTagsForDomainBadRequest with default headers values
func NewGetAssignableTagsForDomainBadRequest() *GetAssignableTagsForDomainBadRequest {
	return &GetAssignableTagsForDomainBadRequest{}
}

/*
GetAssignableTagsForDomainBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetAssignableTagsForDomainBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this get assignable tags for domain bad request response has a 2xx status code
func (o *GetAssignableTagsForDomainBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get assignable tags for domain bad request response has a 3xx status code
func (o *GetAssignableTagsForDomainBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get assignable tags for domain bad request response has a 4xx status code
func (o *GetAssignableTagsForDomainBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get assignable tags for domain bad request response has a 5xx status code
func (o *GetAssignableTagsForDomainBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get assignable tags for domain bad request response a status code equal to that given
func (o *GetAssignableTagsForDomainBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get assignable tags for domain bad request response
func (o *GetAssignableTagsForDomainBadRequest) Code() int {
	return 400
}

func (o *GetAssignableTagsForDomainBadRequest) Error() string {
	return fmt.Sprintf("[GET /v1/domains/{id}/tags/assignable-tags][%d] getAssignableTagsForDomainBadRequest  %+v", 400, o.Payload)
}

func (o *GetAssignableTagsForDomainBadRequest) String() string {
	return fmt.Sprintf("[GET /v1/domains/{id}/tags/assignable-tags][%d] getAssignableTagsForDomainBadRequest  %+v", 400, o.Payload)
}

func (o *GetAssignableTagsForDomainBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetAssignableTagsForDomainBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAssignableTagsForDomainInternalServerError creates a GetAssignableTagsForDomainInternalServerError with default headers values
func NewGetAssignableTagsForDomainInternalServerError() *GetAssignableTagsForDomainInternalServerError {
	return &GetAssignableTagsForDomainInternalServerError{}
}

/*
GetAssignableTagsForDomainInternalServerError describes a response with status code 500, with default header values.

InternalServerError
*/
type GetAssignableTagsForDomainInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this get assignable tags for domain internal server error response has a 2xx status code
func (o *GetAssignableTagsForDomainInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get assignable tags for domain internal server error response has a 3xx status code
func (o *GetAssignableTagsForDomainInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get assignable tags for domain internal server error response has a 4xx status code
func (o *GetAssignableTagsForDomainInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get assignable tags for domain internal server error response has a 5xx status code
func (o *GetAssignableTagsForDomainInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get assignable tags for domain internal server error response a status code equal to that given
func (o *GetAssignableTagsForDomainInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get assignable tags for domain internal server error response
func (o *GetAssignableTagsForDomainInternalServerError) Code() int {
	return 500
}

func (o *GetAssignableTagsForDomainInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/domains/{id}/tags/assignable-tags][%d] getAssignableTagsForDomainInternalServerError  %+v", 500, o.Payload)
}

func (o *GetAssignableTagsForDomainInternalServerError) String() string {
	return fmt.Sprintf("[GET /v1/domains/{id}/tags/assignable-tags][%d] getAssignableTagsForDomainInternalServerError  %+v", 500, o.Payload)
}

func (o *GetAssignableTagsForDomainInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetAssignableTagsForDomainInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
