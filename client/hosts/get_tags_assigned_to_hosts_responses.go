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

// GETTagsAssignedToHostsReader is a Reader for the GETTagsAssignedToHosts structure.
type GETTagsAssignedToHostsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GETTagsAssignedToHostsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGETTagsAssignedToHostsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGETTagsAssignedToHostsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGETTagsAssignedToHostsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGETTagsAssignedToHostsOK creates a GETTagsAssignedToHostsOK with default headers values
func NewGETTagsAssignedToHostsOK() *GETTagsAssignedToHostsOK {
	return &GETTagsAssignedToHostsOK{}
}

/*
GETTagsAssignedToHostsOK describes a response with status code 200, with default header values.

Ok
*/
type GETTagsAssignedToHostsOK struct {
	Payload *models.PageOfTagsForResource
}

// IsSuccess returns true when this get tags assigned to hosts o k response has a 2xx status code
func (o *GETTagsAssignedToHostsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get tags assigned to hosts o k response has a 3xx status code
func (o *GETTagsAssignedToHostsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get tags assigned to hosts o k response has a 4xx status code
func (o *GETTagsAssignedToHostsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get tags assigned to hosts o k response has a 5xx status code
func (o *GETTagsAssignedToHostsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get tags assigned to hosts o k response a status code equal to that given
func (o *GETTagsAssignedToHostsOK) IsCode(code int) bool {
	return code == 200
}

func (o *GETTagsAssignedToHostsOK) Error() string {
	return fmt.Sprintf("[GET /v1/hosts/tags][%d] getTagsAssignedToHostsOK  %+v", 200, o.Payload)
}

func (o *GETTagsAssignedToHostsOK) String() string {
	return fmt.Sprintf("[GET /v1/hosts/tags][%d] getTagsAssignedToHostsOK  %+v", 200, o.Payload)
}

func (o *GETTagsAssignedToHostsOK) GetPayload() *models.PageOfTagsForResource {
	return o.Payload
}

func (o *GETTagsAssignedToHostsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PageOfTagsForResource)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGETTagsAssignedToHostsBadRequest creates a GETTagsAssignedToHostsBadRequest with default headers values
func NewGETTagsAssignedToHostsBadRequest() *GETTagsAssignedToHostsBadRequest {
	return &GETTagsAssignedToHostsBadRequest{}
}

/*
GETTagsAssignedToHostsBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GETTagsAssignedToHostsBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this get tags assigned to hosts bad request response has a 2xx status code
func (o *GETTagsAssignedToHostsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get tags assigned to hosts bad request response has a 3xx status code
func (o *GETTagsAssignedToHostsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get tags assigned to hosts bad request response has a 4xx status code
func (o *GETTagsAssignedToHostsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get tags assigned to hosts bad request response has a 5xx status code
func (o *GETTagsAssignedToHostsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get tags assigned to hosts bad request response a status code equal to that given
func (o *GETTagsAssignedToHostsBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GETTagsAssignedToHostsBadRequest) Error() string {
	return fmt.Sprintf("[GET /v1/hosts/tags][%d] getTagsAssignedToHostsBadRequest  %+v", 400, o.Payload)
}

func (o *GETTagsAssignedToHostsBadRequest) String() string {
	return fmt.Sprintf("[GET /v1/hosts/tags][%d] getTagsAssignedToHostsBadRequest  %+v", 400, o.Payload)
}

func (o *GETTagsAssignedToHostsBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *GETTagsAssignedToHostsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGETTagsAssignedToHostsInternalServerError creates a GETTagsAssignedToHostsInternalServerError with default headers values
func NewGETTagsAssignedToHostsInternalServerError() *GETTagsAssignedToHostsInternalServerError {
	return &GETTagsAssignedToHostsInternalServerError{}
}

/*
GETTagsAssignedToHostsInternalServerError describes a response with status code 500, with default header values.

InternalServerError
*/
type GETTagsAssignedToHostsInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this get tags assigned to hosts internal server error response has a 2xx status code
func (o *GETTagsAssignedToHostsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get tags assigned to hosts internal server error response has a 3xx status code
func (o *GETTagsAssignedToHostsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get tags assigned to hosts internal server error response has a 4xx status code
func (o *GETTagsAssignedToHostsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get tags assigned to hosts internal server error response has a 5xx status code
func (o *GETTagsAssignedToHostsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get tags assigned to hosts internal server error response a status code equal to that given
func (o *GETTagsAssignedToHostsInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GETTagsAssignedToHostsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/hosts/tags][%d] getTagsAssignedToHostsInternalServerError  %+v", 500, o.Payload)
}

func (o *GETTagsAssignedToHostsInternalServerError) String() string {
	return fmt.Sprintf("[GET /v1/hosts/tags][%d] getTagsAssignedToHostsInternalServerError  %+v", 500, o.Payload)
}

func (o *GETTagsAssignedToHostsInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *GETTagsAssignedToHostsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
