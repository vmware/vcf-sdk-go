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

// PostClustersQueryReader is a Reader for the PostClustersQuery structure.
type PostClustersQueryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostClustersQueryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostClustersQueryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostClustersQueryBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostClustersQueryInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostClustersQueryOK creates a PostClustersQueryOK with default headers values
func NewPostClustersQueryOK() *PostClustersQueryOK {
	return &PostClustersQueryOK{}
}

/*
PostClustersQueryOK describes a response with status code 200, with default header values.

Ok
*/
type PostClustersQueryOK struct {
	Payload *models.ClusterQueryResponse
}

// IsSuccess returns true when this post clusters query o k response has a 2xx status code
func (o *PostClustersQueryOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post clusters query o k response has a 3xx status code
func (o *PostClustersQueryOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post clusters query o k response has a 4xx status code
func (o *PostClustersQueryOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this post clusters query o k response has a 5xx status code
func (o *PostClustersQueryOK) IsServerError() bool {
	return false
}

// IsCode returns true when this post clusters query o k response a status code equal to that given
func (o *PostClustersQueryOK) IsCode(code int) bool {
	return code == 200
}

func (o *PostClustersQueryOK) Error() string {
	return fmt.Sprintf("[POST /v1/domains/{domainId}/clusters/queries][%d] postClustersQueryOK  %+v", 200, o.Payload)
}

func (o *PostClustersQueryOK) String() string {
	return fmt.Sprintf("[POST /v1/domains/{domainId}/clusters/queries][%d] postClustersQueryOK  %+v", 200, o.Payload)
}

func (o *PostClustersQueryOK) GetPayload() *models.ClusterQueryResponse {
	return o.Payload
}

func (o *PostClustersQueryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ClusterQueryResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostClustersQueryBadRequest creates a PostClustersQueryBadRequest with default headers values
func NewPostClustersQueryBadRequest() *PostClustersQueryBadRequest {
	return &PostClustersQueryBadRequest{}
}

/*
PostClustersQueryBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PostClustersQueryBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this post clusters query bad request response has a 2xx status code
func (o *PostClustersQueryBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post clusters query bad request response has a 3xx status code
func (o *PostClustersQueryBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post clusters query bad request response has a 4xx status code
func (o *PostClustersQueryBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this post clusters query bad request response has a 5xx status code
func (o *PostClustersQueryBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this post clusters query bad request response a status code equal to that given
func (o *PostClustersQueryBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *PostClustersQueryBadRequest) Error() string {
	return fmt.Sprintf("[POST /v1/domains/{domainId}/clusters/queries][%d] postClustersQueryBadRequest  %+v", 400, o.Payload)
}

func (o *PostClustersQueryBadRequest) String() string {
	return fmt.Sprintf("[POST /v1/domains/{domainId}/clusters/queries][%d] postClustersQueryBadRequest  %+v", 400, o.Payload)
}

func (o *PostClustersQueryBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostClustersQueryBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostClustersQueryInternalServerError creates a PostClustersQueryInternalServerError with default headers values
func NewPostClustersQueryInternalServerError() *PostClustersQueryInternalServerError {
	return &PostClustersQueryInternalServerError{}
}

/*
PostClustersQueryInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PostClustersQueryInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this post clusters query internal server error response has a 2xx status code
func (o *PostClustersQueryInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post clusters query internal server error response has a 3xx status code
func (o *PostClustersQueryInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post clusters query internal server error response has a 4xx status code
func (o *PostClustersQueryInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this post clusters query internal server error response has a 5xx status code
func (o *PostClustersQueryInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this post clusters query internal server error response a status code equal to that given
func (o *PostClustersQueryInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *PostClustersQueryInternalServerError) Error() string {
	return fmt.Sprintf("[POST /v1/domains/{domainId}/clusters/queries][%d] postClustersQueryInternalServerError  %+v", 500, o.Payload)
}

func (o *PostClustersQueryInternalServerError) String() string {
	return fmt.Sprintf("[POST /v1/domains/{domainId}/clusters/queries][%d] postClustersQueryInternalServerError  %+v", 500, o.Payload)
}

func (o *PostClustersQueryInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostClustersQueryInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
