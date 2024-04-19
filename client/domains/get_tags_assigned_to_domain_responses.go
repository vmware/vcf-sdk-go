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

// GetTagsAssignedToDomainReader is a Reader for the GetTagsAssignedToDomain structure.
type GetTagsAssignedToDomainReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetTagsAssignedToDomainReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetTagsAssignedToDomainOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetTagsAssignedToDomainBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetTagsAssignedToDomainInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /v1/domains/{id}/tags] getTagsAssignedToDomain", response, response.Code())
	}
}

// NewGetTagsAssignedToDomainOK creates a GetTagsAssignedToDomainOK with default headers values
func NewGetTagsAssignedToDomainOK() *GetTagsAssignedToDomainOK {
	return &GetTagsAssignedToDomainOK{}
}

/*
GetTagsAssignedToDomainOK describes a response with status code 200, with default header values.

Ok
*/
type GetTagsAssignedToDomainOK struct {
	Payload *models.PageOfTag
}

// IsSuccess returns true when this get tags assigned to domain o k response has a 2xx status code
func (o *GetTagsAssignedToDomainOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get tags assigned to domain o k response has a 3xx status code
func (o *GetTagsAssignedToDomainOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get tags assigned to domain o k response has a 4xx status code
func (o *GetTagsAssignedToDomainOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get tags assigned to domain o k response has a 5xx status code
func (o *GetTagsAssignedToDomainOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get tags assigned to domain o k response a status code equal to that given
func (o *GetTagsAssignedToDomainOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get tags assigned to domain o k response
func (o *GetTagsAssignedToDomainOK) Code() int {
	return 200
}

func (o *GetTagsAssignedToDomainOK) Error() string {
	return fmt.Sprintf("[GET /v1/domains/{id}/tags][%d] getTagsAssignedToDomainOK  %+v", 200, o.Payload)
}

func (o *GetTagsAssignedToDomainOK) String() string {
	return fmt.Sprintf("[GET /v1/domains/{id}/tags][%d] getTagsAssignedToDomainOK  %+v", 200, o.Payload)
}

func (o *GetTagsAssignedToDomainOK) GetPayload() *models.PageOfTag {
	return o.Payload
}

func (o *GetTagsAssignedToDomainOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PageOfTag)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTagsAssignedToDomainBadRequest creates a GetTagsAssignedToDomainBadRequest with default headers values
func NewGetTagsAssignedToDomainBadRequest() *GetTagsAssignedToDomainBadRequest {
	return &GetTagsAssignedToDomainBadRequest{}
}

/*
GetTagsAssignedToDomainBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetTagsAssignedToDomainBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this get tags assigned to domain bad request response has a 2xx status code
func (o *GetTagsAssignedToDomainBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get tags assigned to domain bad request response has a 3xx status code
func (o *GetTagsAssignedToDomainBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get tags assigned to domain bad request response has a 4xx status code
func (o *GetTagsAssignedToDomainBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get tags assigned to domain bad request response has a 5xx status code
func (o *GetTagsAssignedToDomainBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get tags assigned to domain bad request response a status code equal to that given
func (o *GetTagsAssignedToDomainBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get tags assigned to domain bad request response
func (o *GetTagsAssignedToDomainBadRequest) Code() int {
	return 400
}

func (o *GetTagsAssignedToDomainBadRequest) Error() string {
	return fmt.Sprintf("[GET /v1/domains/{id}/tags][%d] getTagsAssignedToDomainBadRequest  %+v", 400, o.Payload)
}

func (o *GetTagsAssignedToDomainBadRequest) String() string {
	return fmt.Sprintf("[GET /v1/domains/{id}/tags][%d] getTagsAssignedToDomainBadRequest  %+v", 400, o.Payload)
}

func (o *GetTagsAssignedToDomainBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetTagsAssignedToDomainBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTagsAssignedToDomainInternalServerError creates a GetTagsAssignedToDomainInternalServerError with default headers values
func NewGetTagsAssignedToDomainInternalServerError() *GetTagsAssignedToDomainInternalServerError {
	return &GetTagsAssignedToDomainInternalServerError{}
}

/*
GetTagsAssignedToDomainInternalServerError describes a response with status code 500, with default header values.

InternalServerError
*/
type GetTagsAssignedToDomainInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this get tags assigned to domain internal server error response has a 2xx status code
func (o *GetTagsAssignedToDomainInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get tags assigned to domain internal server error response has a 3xx status code
func (o *GetTagsAssignedToDomainInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get tags assigned to domain internal server error response has a 4xx status code
func (o *GetTagsAssignedToDomainInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get tags assigned to domain internal server error response has a 5xx status code
func (o *GetTagsAssignedToDomainInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get tags assigned to domain internal server error response a status code equal to that given
func (o *GetTagsAssignedToDomainInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get tags assigned to domain internal server error response
func (o *GetTagsAssignedToDomainInternalServerError) Code() int {
	return 500
}

func (o *GetTagsAssignedToDomainInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/domains/{id}/tags][%d] getTagsAssignedToDomainInternalServerError  %+v", 500, o.Payload)
}

func (o *GetTagsAssignedToDomainInternalServerError) String() string {
	return fmt.Sprintf("[GET /v1/domains/{id}/tags][%d] getTagsAssignedToDomainInternalServerError  %+v", 500, o.Payload)
}

func (o *GetTagsAssignedToDomainInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetTagsAssignedToDomainInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
