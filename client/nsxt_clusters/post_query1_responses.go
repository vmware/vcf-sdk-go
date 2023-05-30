// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2023 VMware, Inc.
// SPDX-License-Identifier: BSD-2-Clause

package nsxt_clusters

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/vmware/vcf-sdk-go/models"
)

// PostQuery1Reader is a Reader for the PostQuery1 structure.
type PostQuery1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostQuery1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostQuery1OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewPostQuery1Accepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostQuery1BadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostQuery1InternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostQuery1OK creates a PostQuery1OK with default headers values
func NewPostQuery1OK() *PostQuery1OK {
	return &PostQuery1OK{}
}

/*
PostQuery1OK describes a response with status code 200, with default header values.

OK
*/
type PostQuery1OK struct {
	Payload *models.NsxTQueryResponse
}

// IsSuccess returns true when this post query1 o k response has a 2xx status code
func (o *PostQuery1OK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post query1 o k response has a 3xx status code
func (o *PostQuery1OK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post query1 o k response has a 4xx status code
func (o *PostQuery1OK) IsClientError() bool {
	return false
}

// IsServerError returns true when this post query1 o k response has a 5xx status code
func (o *PostQuery1OK) IsServerError() bool {
	return false
}

// IsCode returns true when this post query1 o k response a status code equal to that given
func (o *PostQuery1OK) IsCode(code int) bool {
	return code == 200
}

func (o *PostQuery1OK) Error() string {
	return fmt.Sprintf("[POST /v1/nsxt-clusters/queries][%d] postQuery1OK  %+v", 200, o.Payload)
}

func (o *PostQuery1OK) String() string {
	return fmt.Sprintf("[POST /v1/nsxt-clusters/queries][%d] postQuery1OK  %+v", 200, o.Payload)
}

func (o *PostQuery1OK) GetPayload() *models.NsxTQueryResponse {
	return o.Payload
}

func (o *PostQuery1OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NsxTQueryResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostQuery1Accepted creates a PostQuery1Accepted with default headers values
func NewPostQuery1Accepted() *PostQuery1Accepted {
	return &PostQuery1Accepted{}
}

/*
PostQuery1Accepted describes a response with status code 202, with default header values.

Accepted
*/
type PostQuery1Accepted struct {
	Payload *models.NsxTQueryResponse
}

// IsSuccess returns true when this post query1 accepted response has a 2xx status code
func (o *PostQuery1Accepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post query1 accepted response has a 3xx status code
func (o *PostQuery1Accepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post query1 accepted response has a 4xx status code
func (o *PostQuery1Accepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this post query1 accepted response has a 5xx status code
func (o *PostQuery1Accepted) IsServerError() bool {
	return false
}

// IsCode returns true when this post query1 accepted response a status code equal to that given
func (o *PostQuery1Accepted) IsCode(code int) bool {
	return code == 202
}

func (o *PostQuery1Accepted) Error() string {
	return fmt.Sprintf("[POST /v1/nsxt-clusters/queries][%d] postQuery1Accepted  %+v", 202, o.Payload)
}

func (o *PostQuery1Accepted) String() string {
	return fmt.Sprintf("[POST /v1/nsxt-clusters/queries][%d] postQuery1Accepted  %+v", 202, o.Payload)
}

func (o *PostQuery1Accepted) GetPayload() *models.NsxTQueryResponse {
	return o.Payload
}

func (o *PostQuery1Accepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NsxTQueryResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostQuery1BadRequest creates a PostQuery1BadRequest with default headers values
func NewPostQuery1BadRequest() *PostQuery1BadRequest {
	return &PostQuery1BadRequest{}
}

/*
PostQuery1BadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PostQuery1BadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this post query1 bad request response has a 2xx status code
func (o *PostQuery1BadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post query1 bad request response has a 3xx status code
func (o *PostQuery1BadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post query1 bad request response has a 4xx status code
func (o *PostQuery1BadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this post query1 bad request response has a 5xx status code
func (o *PostQuery1BadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this post query1 bad request response a status code equal to that given
func (o *PostQuery1BadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *PostQuery1BadRequest) Error() string {
	return fmt.Sprintf("[POST /v1/nsxt-clusters/queries][%d] postQuery1BadRequest  %+v", 400, o.Payload)
}

func (o *PostQuery1BadRequest) String() string {
	return fmt.Sprintf("[POST /v1/nsxt-clusters/queries][%d] postQuery1BadRequest  %+v", 400, o.Payload)
}

func (o *PostQuery1BadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostQuery1BadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostQuery1InternalServerError creates a PostQuery1InternalServerError with default headers values
func NewPostQuery1InternalServerError() *PostQuery1InternalServerError {
	return &PostQuery1InternalServerError{}
}

/*
PostQuery1InternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PostQuery1InternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this post query1 internal server error response has a 2xx status code
func (o *PostQuery1InternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post query1 internal server error response has a 3xx status code
func (o *PostQuery1InternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post query1 internal server error response has a 4xx status code
func (o *PostQuery1InternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this post query1 internal server error response has a 5xx status code
func (o *PostQuery1InternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this post query1 internal server error response a status code equal to that given
func (o *PostQuery1InternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *PostQuery1InternalServerError) Error() string {
	return fmt.Sprintf("[POST /v1/nsxt-clusters/queries][%d] postQuery1InternalServerError  %+v", 500, o.Payload)
}

func (o *PostQuery1InternalServerError) String() string {
	return fmt.Sprintf("[POST /v1/nsxt-clusters/queries][%d] postQuery1InternalServerError  %+v", 500, o.Payload)
}

func (o *PostQuery1InternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostQuery1InternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
