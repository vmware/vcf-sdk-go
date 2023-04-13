// Code generated by go-swagger; DO NOT EDIT.

// Copyright © 2023 VMware, Inc. All Rights Reserved.
// Code licenced under: BSD-2

package clusters

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"vcf-sdk-go/models"
)

// GETTagsAssignedToClusterReader is a Reader for the GETTagsAssignedToCluster structure.
type GETTagsAssignedToClusterReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GETTagsAssignedToClusterReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGETTagsAssignedToClusterOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGETTagsAssignedToClusterBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGETTagsAssignedToClusterInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGETTagsAssignedToClusterOK creates a GETTagsAssignedToClusterOK with default headers values
func NewGETTagsAssignedToClusterOK() *GETTagsAssignedToClusterOK {
	return &GETTagsAssignedToClusterOK{}
}

/*
GETTagsAssignedToClusterOK describes a response with status code 200, with default header values.

Ok
*/
type GETTagsAssignedToClusterOK struct {
	Payload *models.PageOfTag
}

// IsSuccess returns true when this get tags assigned to cluster o k response has a 2xx status code
func (o *GETTagsAssignedToClusterOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get tags assigned to cluster o k response has a 3xx status code
func (o *GETTagsAssignedToClusterOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get tags assigned to cluster o k response has a 4xx status code
func (o *GETTagsAssignedToClusterOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get tags assigned to cluster o k response has a 5xx status code
func (o *GETTagsAssignedToClusterOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get tags assigned to cluster o k response a status code equal to that given
func (o *GETTagsAssignedToClusterOK) IsCode(code int) bool {
	return code == 200
}

func (o *GETTagsAssignedToClusterOK) Error() string {
	return fmt.Sprintf("[GET /v1/clusters/{id}/tags][%d] getTagsAssignedToClusterOK  %+v", 200, o.Payload)
}

func (o *GETTagsAssignedToClusterOK) String() string {
	return fmt.Sprintf("[GET /v1/clusters/{id}/tags][%d] getTagsAssignedToClusterOK  %+v", 200, o.Payload)
}

func (o *GETTagsAssignedToClusterOK) GetPayload() *models.PageOfTag {
	return o.Payload
}

func (o *GETTagsAssignedToClusterOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PageOfTag)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGETTagsAssignedToClusterBadRequest creates a GETTagsAssignedToClusterBadRequest with default headers values
func NewGETTagsAssignedToClusterBadRequest() *GETTagsAssignedToClusterBadRequest {
	return &GETTagsAssignedToClusterBadRequest{}
}

/*
GETTagsAssignedToClusterBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GETTagsAssignedToClusterBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this get tags assigned to cluster bad request response has a 2xx status code
func (o *GETTagsAssignedToClusterBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get tags assigned to cluster bad request response has a 3xx status code
func (o *GETTagsAssignedToClusterBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get tags assigned to cluster bad request response has a 4xx status code
func (o *GETTagsAssignedToClusterBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get tags assigned to cluster bad request response has a 5xx status code
func (o *GETTagsAssignedToClusterBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get tags assigned to cluster bad request response a status code equal to that given
func (o *GETTagsAssignedToClusterBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GETTagsAssignedToClusterBadRequest) Error() string {
	return fmt.Sprintf("[GET /v1/clusters/{id}/tags][%d] getTagsAssignedToClusterBadRequest  %+v", 400, o.Payload)
}

func (o *GETTagsAssignedToClusterBadRequest) String() string {
	return fmt.Sprintf("[GET /v1/clusters/{id}/tags][%d] getTagsAssignedToClusterBadRequest  %+v", 400, o.Payload)
}

func (o *GETTagsAssignedToClusterBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *GETTagsAssignedToClusterBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGETTagsAssignedToClusterInternalServerError creates a GETTagsAssignedToClusterInternalServerError with default headers values
func NewGETTagsAssignedToClusterInternalServerError() *GETTagsAssignedToClusterInternalServerError {
	return &GETTagsAssignedToClusterInternalServerError{}
}

/*
GETTagsAssignedToClusterInternalServerError describes a response with status code 500, with default header values.

InternalServerError
*/
type GETTagsAssignedToClusterInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this get tags assigned to cluster internal server error response has a 2xx status code
func (o *GETTagsAssignedToClusterInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get tags assigned to cluster internal server error response has a 3xx status code
func (o *GETTagsAssignedToClusterInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get tags assigned to cluster internal server error response has a 4xx status code
func (o *GETTagsAssignedToClusterInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get tags assigned to cluster internal server error response has a 5xx status code
func (o *GETTagsAssignedToClusterInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get tags assigned to cluster internal server error response a status code equal to that given
func (o *GETTagsAssignedToClusterInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GETTagsAssignedToClusterInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/clusters/{id}/tags][%d] getTagsAssignedToClusterInternalServerError  %+v", 500, o.Payload)
}

func (o *GETTagsAssignedToClusterInternalServerError) String() string {
	return fmt.Sprintf("[GET /v1/clusters/{id}/tags][%d] getTagsAssignedToClusterInternalServerError  %+v", 500, o.Payload)
}

func (o *GETTagsAssignedToClusterInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *GETTagsAssignedToClusterInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
