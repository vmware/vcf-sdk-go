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

// PostClusterQueryReader is a Reader for the PostClusterQuery structure.
type PostClusterQueryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostClusterQueryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostClusterQueryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostClusterQueryBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostClusterQueryNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostClusterQueryInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /v1/domains/{domainId}/clusters/{clusterName}/queries] postClusterQuery", response, response.Code())
	}
}

// NewPostClusterQueryOK creates a PostClusterQueryOK with default headers values
func NewPostClusterQueryOK() *PostClusterQueryOK {
	return &PostClusterQueryOK{}
}

/*
PostClusterQueryOK describes a response with status code 200, with default header values.

Ok
*/
type PostClusterQueryOK struct {
	Payload *models.ClusterQueryResponse
}

// IsSuccess returns true when this post cluster query o k response has a 2xx status code
func (o *PostClusterQueryOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post cluster query o k response has a 3xx status code
func (o *PostClusterQueryOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post cluster query o k response has a 4xx status code
func (o *PostClusterQueryOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this post cluster query o k response has a 5xx status code
func (o *PostClusterQueryOK) IsServerError() bool {
	return false
}

// IsCode returns true when this post cluster query o k response a status code equal to that given
func (o *PostClusterQueryOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the post cluster query o k response
func (o *PostClusterQueryOK) Code() int {
	return 200
}

func (o *PostClusterQueryOK) Error() string {
	return fmt.Sprintf("[POST /v1/domains/{domainId}/clusters/{clusterName}/queries][%d] postClusterQueryOK  %+v", 200, o.Payload)
}

func (o *PostClusterQueryOK) String() string {
	return fmt.Sprintf("[POST /v1/domains/{domainId}/clusters/{clusterName}/queries][%d] postClusterQueryOK  %+v", 200, o.Payload)
}

func (o *PostClusterQueryOK) GetPayload() *models.ClusterQueryResponse {
	return o.Payload
}

func (o *PostClusterQueryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ClusterQueryResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostClusterQueryBadRequest creates a PostClusterQueryBadRequest with default headers values
func NewPostClusterQueryBadRequest() *PostClusterQueryBadRequest {
	return &PostClusterQueryBadRequest{}
}

/*
PostClusterQueryBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PostClusterQueryBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this post cluster query bad request response has a 2xx status code
func (o *PostClusterQueryBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post cluster query bad request response has a 3xx status code
func (o *PostClusterQueryBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post cluster query bad request response has a 4xx status code
func (o *PostClusterQueryBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this post cluster query bad request response has a 5xx status code
func (o *PostClusterQueryBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this post cluster query bad request response a status code equal to that given
func (o *PostClusterQueryBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the post cluster query bad request response
func (o *PostClusterQueryBadRequest) Code() int {
	return 400
}

func (o *PostClusterQueryBadRequest) Error() string {
	return fmt.Sprintf("[POST /v1/domains/{domainId}/clusters/{clusterName}/queries][%d] postClusterQueryBadRequest  %+v", 400, o.Payload)
}

func (o *PostClusterQueryBadRequest) String() string {
	return fmt.Sprintf("[POST /v1/domains/{domainId}/clusters/{clusterName}/queries][%d] postClusterQueryBadRequest  %+v", 400, o.Payload)
}

func (o *PostClusterQueryBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostClusterQueryBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostClusterQueryNotFound creates a PostClusterQueryNotFound with default headers values
func NewPostClusterQueryNotFound() *PostClusterQueryNotFound {
	return &PostClusterQueryNotFound{}
}

/*
PostClusterQueryNotFound describes a response with status code 404, with default header values.

Not Found
*/
type PostClusterQueryNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this post cluster query not found response has a 2xx status code
func (o *PostClusterQueryNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post cluster query not found response has a 3xx status code
func (o *PostClusterQueryNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post cluster query not found response has a 4xx status code
func (o *PostClusterQueryNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this post cluster query not found response has a 5xx status code
func (o *PostClusterQueryNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this post cluster query not found response a status code equal to that given
func (o *PostClusterQueryNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the post cluster query not found response
func (o *PostClusterQueryNotFound) Code() int {
	return 404
}

func (o *PostClusterQueryNotFound) Error() string {
	return fmt.Sprintf("[POST /v1/domains/{domainId}/clusters/{clusterName}/queries][%d] postClusterQueryNotFound  %+v", 404, o.Payload)
}

func (o *PostClusterQueryNotFound) String() string {
	return fmt.Sprintf("[POST /v1/domains/{domainId}/clusters/{clusterName}/queries][%d] postClusterQueryNotFound  %+v", 404, o.Payload)
}

func (o *PostClusterQueryNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostClusterQueryNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostClusterQueryInternalServerError creates a PostClusterQueryInternalServerError with default headers values
func NewPostClusterQueryInternalServerError() *PostClusterQueryInternalServerError {
	return &PostClusterQueryInternalServerError{}
}

/*
PostClusterQueryInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PostClusterQueryInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this post cluster query internal server error response has a 2xx status code
func (o *PostClusterQueryInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post cluster query internal server error response has a 3xx status code
func (o *PostClusterQueryInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post cluster query internal server error response has a 4xx status code
func (o *PostClusterQueryInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this post cluster query internal server error response has a 5xx status code
func (o *PostClusterQueryInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this post cluster query internal server error response a status code equal to that given
func (o *PostClusterQueryInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the post cluster query internal server error response
func (o *PostClusterQueryInternalServerError) Code() int {
	return 500
}

func (o *PostClusterQueryInternalServerError) Error() string {
	return fmt.Sprintf("[POST /v1/domains/{domainId}/clusters/{clusterName}/queries][%d] postClusterQueryInternalServerError  %+v", 500, o.Payload)
}

func (o *PostClusterQueryInternalServerError) String() string {
	return fmt.Sprintf("[POST /v1/domains/{domainId}/clusters/{clusterName}/queries][%d] postClusterQueryInternalServerError  %+v", 500, o.Payload)
}

func (o *PostClusterQueryInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostClusterQueryInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
