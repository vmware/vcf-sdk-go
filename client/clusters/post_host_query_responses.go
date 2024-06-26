// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2023 VMware, Inc.
// SPDX-License-Identifier: BSD-2-Clause

package clusters

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/vmware/vcf-sdk-go/models"
)

// PostHostQueryReader is a Reader for the PostHostQuery structure.
type PostHostQueryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostHostQueryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostHostQueryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostHostQueryBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostHostQueryNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostHostQueryInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /v1/clusters/{id}/hosts/queries] postHostQuery", response, response.Code())
	}
}

// NewPostHostQueryOK creates a PostHostQueryOK with default headers values
func NewPostHostQueryOK() *PostHostQueryOK {
	return &PostHostQueryOK{}
}

/*
PostHostQueryOK describes a response with status code 200, with default header values.

Ok
*/
type PostHostQueryOK struct {
	Payload *models.HostQueryResponse
}

// IsSuccess returns true when this post host query o k response has a 2xx status code
func (o *PostHostQueryOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post host query o k response has a 3xx status code
func (o *PostHostQueryOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post host query o k response has a 4xx status code
func (o *PostHostQueryOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this post host query o k response has a 5xx status code
func (o *PostHostQueryOK) IsServerError() bool {
	return false
}

// IsCode returns true when this post host query o k response a status code equal to that given
func (o *PostHostQueryOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the post host query o k response
func (o *PostHostQueryOK) Code() int {
	return 200
}

func (o *PostHostQueryOK) Error() string {
	return fmt.Sprintf("[POST /v1/clusters/{id}/hosts/queries][%d] postHostQueryOK  %+v", 200, o.Payload)
}

func (o *PostHostQueryOK) String() string {
	return fmt.Sprintf("[POST /v1/clusters/{id}/hosts/queries][%d] postHostQueryOK  %+v", 200, o.Payload)
}

func (o *PostHostQueryOK) GetPayload() *models.HostQueryResponse {
	return o.Payload
}

func (o *PostHostQueryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HostQueryResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostHostQueryBadRequest creates a PostHostQueryBadRequest with default headers values
func NewPostHostQueryBadRequest() *PostHostQueryBadRequest {
	return &PostHostQueryBadRequest{}
}

/*
PostHostQueryBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PostHostQueryBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this post host query bad request response has a 2xx status code
func (o *PostHostQueryBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post host query bad request response has a 3xx status code
func (o *PostHostQueryBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post host query bad request response has a 4xx status code
func (o *PostHostQueryBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this post host query bad request response has a 5xx status code
func (o *PostHostQueryBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this post host query bad request response a status code equal to that given
func (o *PostHostQueryBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the post host query bad request response
func (o *PostHostQueryBadRequest) Code() int {
	return 400
}

func (o *PostHostQueryBadRequest) Error() string {
	return fmt.Sprintf("[POST /v1/clusters/{id}/hosts/queries][%d] postHostQueryBadRequest  %+v", 400, o.Payload)
}

func (o *PostHostQueryBadRequest) String() string {
	return fmt.Sprintf("[POST /v1/clusters/{id}/hosts/queries][%d] postHostQueryBadRequest  %+v", 400, o.Payload)
}

func (o *PostHostQueryBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostHostQueryBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostHostQueryNotFound creates a PostHostQueryNotFound with default headers values
func NewPostHostQueryNotFound() *PostHostQueryNotFound {
	return &PostHostQueryNotFound{}
}

/*
PostHostQueryNotFound describes a response with status code 404, with default header values.

Cluster Not Found
*/
type PostHostQueryNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this post host query not found response has a 2xx status code
func (o *PostHostQueryNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post host query not found response has a 3xx status code
func (o *PostHostQueryNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post host query not found response has a 4xx status code
func (o *PostHostQueryNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this post host query not found response has a 5xx status code
func (o *PostHostQueryNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this post host query not found response a status code equal to that given
func (o *PostHostQueryNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the post host query not found response
func (o *PostHostQueryNotFound) Code() int {
	return 404
}

func (o *PostHostQueryNotFound) Error() string {
	return fmt.Sprintf("[POST /v1/clusters/{id}/hosts/queries][%d] postHostQueryNotFound  %+v", 404, o.Payload)
}

func (o *PostHostQueryNotFound) String() string {
	return fmt.Sprintf("[POST /v1/clusters/{id}/hosts/queries][%d] postHostQueryNotFound  %+v", 404, o.Payload)
}

func (o *PostHostQueryNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostHostQueryNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostHostQueryInternalServerError creates a PostHostQueryInternalServerError with default headers values
func NewPostHostQueryInternalServerError() *PostHostQueryInternalServerError {
	return &PostHostQueryInternalServerError{}
}

/*
PostHostQueryInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PostHostQueryInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this post host query internal server error response has a 2xx status code
func (o *PostHostQueryInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post host query internal server error response has a 3xx status code
func (o *PostHostQueryInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post host query internal server error response has a 4xx status code
func (o *PostHostQueryInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this post host query internal server error response has a 5xx status code
func (o *PostHostQueryInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this post host query internal server error response a status code equal to that given
func (o *PostHostQueryInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the post host query internal server error response
func (o *PostHostQueryInternalServerError) Code() int {
	return 500
}

func (o *PostHostQueryInternalServerError) Error() string {
	return fmt.Sprintf("[POST /v1/clusters/{id}/hosts/queries][%d] postHostQueryInternalServerError  %+v", 500, o.Payload)
}

func (o *PostHostQueryInternalServerError) String() string {
	return fmt.Sprintf("[POST /v1/clusters/{id}/hosts/queries][%d] postHostQueryInternalServerError  %+v", 500, o.Payload)
}

func (o *PostHostQueryInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostHostQueryInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
