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

	"vcf-sdk-go/models"
)

// RemoveTagsFromHostReader is a Reader for the RemoveTagsFromHost structure.
type RemoveTagsFromHostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RemoveTagsFromHostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRemoveTagsFromHostOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewRemoveTagsFromHostBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewRemoveTagsFromHostInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewRemoveTagsFromHostOK creates a RemoveTagsFromHostOK with default headers values
func NewRemoveTagsFromHostOK() *RemoveTagsFromHostOK {
	return &RemoveTagsFromHostOK{}
}

/*
RemoveTagsFromHostOK describes a response with status code 200, with default header values.

Ok
*/
type RemoveTagsFromHostOK struct {
	Payload *models.TagAssignmentResult
}

// IsSuccess returns true when this remove tags from host o k response has a 2xx status code
func (o *RemoveTagsFromHostOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this remove tags from host o k response has a 3xx status code
func (o *RemoveTagsFromHostOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this remove tags from host o k response has a 4xx status code
func (o *RemoveTagsFromHostOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this remove tags from host o k response has a 5xx status code
func (o *RemoveTagsFromHostOK) IsServerError() bool {
	return false
}

// IsCode returns true when this remove tags from host o k response a status code equal to that given
func (o *RemoveTagsFromHostOK) IsCode(code int) bool {
	return code == 200
}

func (o *RemoveTagsFromHostOK) Error() string {
	return fmt.Sprintf("[DELETE /v1/hosts/{id}/tags][%d] removeTagsFromHostOK  %+v", 200, o.Payload)
}

func (o *RemoveTagsFromHostOK) String() string {
	return fmt.Sprintf("[DELETE /v1/hosts/{id}/tags][%d] removeTagsFromHostOK  %+v", 200, o.Payload)
}

func (o *RemoveTagsFromHostOK) GetPayload() *models.TagAssignmentResult {
	return o.Payload
}

func (o *RemoveTagsFromHostOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TagAssignmentResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRemoveTagsFromHostBadRequest creates a RemoveTagsFromHostBadRequest with default headers values
func NewRemoveTagsFromHostBadRequest() *RemoveTagsFromHostBadRequest {
	return &RemoveTagsFromHostBadRequest{}
}

/*
RemoveTagsFromHostBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type RemoveTagsFromHostBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this remove tags from host bad request response has a 2xx status code
func (o *RemoveTagsFromHostBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this remove tags from host bad request response has a 3xx status code
func (o *RemoveTagsFromHostBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this remove tags from host bad request response has a 4xx status code
func (o *RemoveTagsFromHostBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this remove tags from host bad request response has a 5xx status code
func (o *RemoveTagsFromHostBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this remove tags from host bad request response a status code equal to that given
func (o *RemoveTagsFromHostBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *RemoveTagsFromHostBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /v1/hosts/{id}/tags][%d] removeTagsFromHostBadRequest  %+v", 400, o.Payload)
}

func (o *RemoveTagsFromHostBadRequest) String() string {
	return fmt.Sprintf("[DELETE /v1/hosts/{id}/tags][%d] removeTagsFromHostBadRequest  %+v", 400, o.Payload)
}

func (o *RemoveTagsFromHostBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *RemoveTagsFromHostBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRemoveTagsFromHostInternalServerError creates a RemoveTagsFromHostInternalServerError with default headers values
func NewRemoveTagsFromHostInternalServerError() *RemoveTagsFromHostInternalServerError {
	return &RemoveTagsFromHostInternalServerError{}
}

/*
RemoveTagsFromHostInternalServerError describes a response with status code 500, with default header values.

InternalServerError
*/
type RemoveTagsFromHostInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this remove tags from host internal server error response has a 2xx status code
func (o *RemoveTagsFromHostInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this remove tags from host internal server error response has a 3xx status code
func (o *RemoveTagsFromHostInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this remove tags from host internal server error response has a 4xx status code
func (o *RemoveTagsFromHostInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this remove tags from host internal server error response has a 5xx status code
func (o *RemoveTagsFromHostInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this remove tags from host internal server error response a status code equal to that given
func (o *RemoveTagsFromHostInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *RemoveTagsFromHostInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /v1/hosts/{id}/tags][%d] removeTagsFromHostInternalServerError  %+v", 500, o.Payload)
}

func (o *RemoveTagsFromHostInternalServerError) String() string {
	return fmt.Sprintf("[DELETE /v1/hosts/{id}/tags][%d] removeTagsFromHostInternalServerError  %+v", 500, o.Payload)
}

func (o *RemoveTagsFromHostInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *RemoveTagsFromHostInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
