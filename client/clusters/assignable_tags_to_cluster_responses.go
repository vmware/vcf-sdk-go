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

// AssignableTagsToClusterReader is a Reader for the AssignableTagsToCluster structure.
type AssignableTagsToClusterReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AssignableTagsToClusterReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAssignableTagsToClusterOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewAssignableTagsToClusterBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewAssignableTagsToClusterInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewAssignableTagsToClusterOK creates a AssignableTagsToClusterOK with default headers values
func NewAssignableTagsToClusterOK() *AssignableTagsToClusterOK {
	return &AssignableTagsToClusterOK{}
}

/*
AssignableTagsToClusterOK describes a response with status code 200, with default header values.

Ok
*/
type AssignableTagsToClusterOK struct {
	Payload *models.PageOfTag
}

// IsSuccess returns true when this assignable tags to cluster o k response has a 2xx status code
func (o *AssignableTagsToClusterOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this assignable tags to cluster o k response has a 3xx status code
func (o *AssignableTagsToClusterOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this assignable tags to cluster o k response has a 4xx status code
func (o *AssignableTagsToClusterOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this assignable tags to cluster o k response has a 5xx status code
func (o *AssignableTagsToClusterOK) IsServerError() bool {
	return false
}

// IsCode returns true when this assignable tags to cluster o k response a status code equal to that given
func (o *AssignableTagsToClusterOK) IsCode(code int) bool {
	return code == 200
}

func (o *AssignableTagsToClusterOK) Error() string {
	return fmt.Sprintf("[GET /v1/clusters/{id}/tags/assignable-tags][%d] assignableTagsToClusterOK  %+v", 200, o.Payload)
}

func (o *AssignableTagsToClusterOK) String() string {
	return fmt.Sprintf("[GET /v1/clusters/{id}/tags/assignable-tags][%d] assignableTagsToClusterOK  %+v", 200, o.Payload)
}

func (o *AssignableTagsToClusterOK) GetPayload() *models.PageOfTag {
	return o.Payload
}

func (o *AssignableTagsToClusterOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PageOfTag)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAssignableTagsToClusterBadRequest creates a AssignableTagsToClusterBadRequest with default headers values
func NewAssignableTagsToClusterBadRequest() *AssignableTagsToClusterBadRequest {
	return &AssignableTagsToClusterBadRequest{}
}

/*
AssignableTagsToClusterBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type AssignableTagsToClusterBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this assignable tags to cluster bad request response has a 2xx status code
func (o *AssignableTagsToClusterBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this assignable tags to cluster bad request response has a 3xx status code
func (o *AssignableTagsToClusterBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this assignable tags to cluster bad request response has a 4xx status code
func (o *AssignableTagsToClusterBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this assignable tags to cluster bad request response has a 5xx status code
func (o *AssignableTagsToClusterBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this assignable tags to cluster bad request response a status code equal to that given
func (o *AssignableTagsToClusterBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *AssignableTagsToClusterBadRequest) Error() string {
	return fmt.Sprintf("[GET /v1/clusters/{id}/tags/assignable-tags][%d] assignableTagsToClusterBadRequest  %+v", 400, o.Payload)
}

func (o *AssignableTagsToClusterBadRequest) String() string {
	return fmt.Sprintf("[GET /v1/clusters/{id}/tags/assignable-tags][%d] assignableTagsToClusterBadRequest  %+v", 400, o.Payload)
}

func (o *AssignableTagsToClusterBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *AssignableTagsToClusterBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAssignableTagsToClusterInternalServerError creates a AssignableTagsToClusterInternalServerError with default headers values
func NewAssignableTagsToClusterInternalServerError() *AssignableTagsToClusterInternalServerError {
	return &AssignableTagsToClusterInternalServerError{}
}

/*
AssignableTagsToClusterInternalServerError describes a response with status code 500, with default header values.

InternalServerError
*/
type AssignableTagsToClusterInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this assignable tags to cluster internal server error response has a 2xx status code
func (o *AssignableTagsToClusterInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this assignable tags to cluster internal server error response has a 3xx status code
func (o *AssignableTagsToClusterInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this assignable tags to cluster internal server error response has a 4xx status code
func (o *AssignableTagsToClusterInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this assignable tags to cluster internal server error response has a 5xx status code
func (o *AssignableTagsToClusterInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this assignable tags to cluster internal server error response a status code equal to that given
func (o *AssignableTagsToClusterInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *AssignableTagsToClusterInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/clusters/{id}/tags/assignable-tags][%d] assignableTagsToClusterInternalServerError  %+v", 500, o.Payload)
}

func (o *AssignableTagsToClusterInternalServerError) String() string {
	return fmt.Sprintf("[GET /v1/clusters/{id}/tags/assignable-tags][%d] assignableTagsToClusterInternalServerError  %+v", 500, o.Payload)
}

func (o *AssignableTagsToClusterInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *AssignableTagsToClusterInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
