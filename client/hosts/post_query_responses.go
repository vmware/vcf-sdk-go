// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2023 VMware, Inc.
// SPDX-License-Identifier: BSD-2-Clause

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

// PostQueryReader is a Reader for the PostQuery structure.
type PostQueryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostQueryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostQueryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostQueryBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostQueryInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /v1/hosts/queries] postQuery", response, response.Code())
	}
}

// NewPostQueryOK creates a PostQueryOK with default headers values
func NewPostQueryOK() *PostQueryOK {
	return &PostQueryOK{}
}

/*
PostQueryOK describes a response with status code 200, with default header values.

Ok
*/
type PostQueryOK struct {
	Payload *models.HostQueryResponse
}

// IsSuccess returns true when this post query o k response has a 2xx status code
func (o *PostQueryOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post query o k response has a 3xx status code
func (o *PostQueryOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post query o k response has a 4xx status code
func (o *PostQueryOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this post query o k response has a 5xx status code
func (o *PostQueryOK) IsServerError() bool {
	return false
}

// IsCode returns true when this post query o k response a status code equal to that given
func (o *PostQueryOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the post query o k response
func (o *PostQueryOK) Code() int {
	return 200
}

func (o *PostQueryOK) Error() string {
	return fmt.Sprintf("[POST /v1/hosts/queries][%d] postQueryOK  %+v", 200, o.Payload)
}

func (o *PostQueryOK) String() string {
	return fmt.Sprintf("[POST /v1/hosts/queries][%d] postQueryOK  %+v", 200, o.Payload)
}

func (o *PostQueryOK) GetPayload() *models.HostQueryResponse {
	return o.Payload
}

func (o *PostQueryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HostQueryResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostQueryBadRequest creates a PostQueryBadRequest with default headers values
func NewPostQueryBadRequest() *PostQueryBadRequest {
	return &PostQueryBadRequest{}
}

/*
PostQueryBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PostQueryBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this post query bad request response has a 2xx status code
func (o *PostQueryBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post query bad request response has a 3xx status code
func (o *PostQueryBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post query bad request response has a 4xx status code
func (o *PostQueryBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this post query bad request response has a 5xx status code
func (o *PostQueryBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this post query bad request response a status code equal to that given
func (o *PostQueryBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the post query bad request response
func (o *PostQueryBadRequest) Code() int {
	return 400
}

func (o *PostQueryBadRequest) Error() string {
	return fmt.Sprintf("[POST /v1/hosts/queries][%d] postQueryBadRequest  %+v", 400, o.Payload)
}

func (o *PostQueryBadRequest) String() string {
	return fmt.Sprintf("[POST /v1/hosts/queries][%d] postQueryBadRequest  %+v", 400, o.Payload)
}

func (o *PostQueryBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostQueryBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostQueryInternalServerError creates a PostQueryInternalServerError with default headers values
func NewPostQueryInternalServerError() *PostQueryInternalServerError {
	return &PostQueryInternalServerError{}
}

/*
PostQueryInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PostQueryInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this post query internal server error response has a 2xx status code
func (o *PostQueryInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post query internal server error response has a 3xx status code
func (o *PostQueryInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post query internal server error response has a 4xx status code
func (o *PostQueryInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this post query internal server error response has a 5xx status code
func (o *PostQueryInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this post query internal server error response a status code equal to that given
func (o *PostQueryInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the post query internal server error response
func (o *PostQueryInternalServerError) Code() int {
	return 500
}

func (o *PostQueryInternalServerError) Error() string {
	return fmt.Sprintf("[POST /v1/hosts/queries][%d] postQueryInternalServerError  %+v", 500, o.Payload)
}

func (o *PostQueryInternalServerError) String() string {
	return fmt.Sprintf("[POST /v1/hosts/queries][%d] postQueryInternalServerError  %+v", 500, o.Payload)
}

func (o *PostQueryInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostQueryInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
