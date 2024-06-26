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

// GetDatastoreQueryResponse1Reader is a Reader for the GetDatastoreQueryResponse1 structure.
type GetDatastoreQueryResponse1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDatastoreQueryResponse1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetDatastoreQueryResponse1OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetDatastoreQueryResponse1BadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetDatastoreQueryResponse1NotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetDatastoreQueryResponse1InternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /v1/domains/{domainId}/datastores/queries/{queryId}] getDatastoreQueryResponse_1", response, response.Code())
	}
}

// NewGetDatastoreQueryResponse1OK creates a GetDatastoreQueryResponse1OK with default headers values
func NewGetDatastoreQueryResponse1OK() *GetDatastoreQueryResponse1OK {
	return &GetDatastoreQueryResponse1OK{}
}

/*
GetDatastoreQueryResponse1OK describes a response with status code 200, with default header values.

Ok
*/
type GetDatastoreQueryResponse1OK struct {
	Payload *models.DatastoreQueryResponse
}

// IsSuccess returns true when this get datastore query response1 o k response has a 2xx status code
func (o *GetDatastoreQueryResponse1OK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get datastore query response1 o k response has a 3xx status code
func (o *GetDatastoreQueryResponse1OK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get datastore query response1 o k response has a 4xx status code
func (o *GetDatastoreQueryResponse1OK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get datastore query response1 o k response has a 5xx status code
func (o *GetDatastoreQueryResponse1OK) IsServerError() bool {
	return false
}

// IsCode returns true when this get datastore query response1 o k response a status code equal to that given
func (o *GetDatastoreQueryResponse1OK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get datastore query response1 o k response
func (o *GetDatastoreQueryResponse1OK) Code() int {
	return 200
}

func (o *GetDatastoreQueryResponse1OK) Error() string {
	return fmt.Sprintf("[GET /v1/domains/{domainId}/datastores/queries/{queryId}][%d] getDatastoreQueryResponse1OK  %+v", 200, o.Payload)
}

func (o *GetDatastoreQueryResponse1OK) String() string {
	return fmt.Sprintf("[GET /v1/domains/{domainId}/datastores/queries/{queryId}][%d] getDatastoreQueryResponse1OK  %+v", 200, o.Payload)
}

func (o *GetDatastoreQueryResponse1OK) GetPayload() *models.DatastoreQueryResponse {
	return o.Payload
}

func (o *GetDatastoreQueryResponse1OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DatastoreQueryResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDatastoreQueryResponse1BadRequest creates a GetDatastoreQueryResponse1BadRequest with default headers values
func NewGetDatastoreQueryResponse1BadRequest() *GetDatastoreQueryResponse1BadRequest {
	return &GetDatastoreQueryResponse1BadRequest{}
}

/*
GetDatastoreQueryResponse1BadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetDatastoreQueryResponse1BadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this get datastore query response1 bad request response has a 2xx status code
func (o *GetDatastoreQueryResponse1BadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get datastore query response1 bad request response has a 3xx status code
func (o *GetDatastoreQueryResponse1BadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get datastore query response1 bad request response has a 4xx status code
func (o *GetDatastoreQueryResponse1BadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get datastore query response1 bad request response has a 5xx status code
func (o *GetDatastoreQueryResponse1BadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get datastore query response1 bad request response a status code equal to that given
func (o *GetDatastoreQueryResponse1BadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get datastore query response1 bad request response
func (o *GetDatastoreQueryResponse1BadRequest) Code() int {
	return 400
}

func (o *GetDatastoreQueryResponse1BadRequest) Error() string {
	return fmt.Sprintf("[GET /v1/domains/{domainId}/datastores/queries/{queryId}][%d] getDatastoreQueryResponse1BadRequest  %+v", 400, o.Payload)
}

func (o *GetDatastoreQueryResponse1BadRequest) String() string {
	return fmt.Sprintf("[GET /v1/domains/{domainId}/datastores/queries/{queryId}][%d] getDatastoreQueryResponse1BadRequest  %+v", 400, o.Payload)
}

func (o *GetDatastoreQueryResponse1BadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetDatastoreQueryResponse1BadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDatastoreQueryResponse1NotFound creates a GetDatastoreQueryResponse1NotFound with default headers values
func NewGetDatastoreQueryResponse1NotFound() *GetDatastoreQueryResponse1NotFound {
	return &GetDatastoreQueryResponse1NotFound{}
}

/*
GetDatastoreQueryResponse1NotFound describes a response with status code 404, with default header values.

Query Not Found
*/
type GetDatastoreQueryResponse1NotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this get datastore query response1 not found response has a 2xx status code
func (o *GetDatastoreQueryResponse1NotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get datastore query response1 not found response has a 3xx status code
func (o *GetDatastoreQueryResponse1NotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get datastore query response1 not found response has a 4xx status code
func (o *GetDatastoreQueryResponse1NotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get datastore query response1 not found response has a 5xx status code
func (o *GetDatastoreQueryResponse1NotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get datastore query response1 not found response a status code equal to that given
func (o *GetDatastoreQueryResponse1NotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get datastore query response1 not found response
func (o *GetDatastoreQueryResponse1NotFound) Code() int {
	return 404
}

func (o *GetDatastoreQueryResponse1NotFound) Error() string {
	return fmt.Sprintf("[GET /v1/domains/{domainId}/datastores/queries/{queryId}][%d] getDatastoreQueryResponse1NotFound  %+v", 404, o.Payload)
}

func (o *GetDatastoreQueryResponse1NotFound) String() string {
	return fmt.Sprintf("[GET /v1/domains/{domainId}/datastores/queries/{queryId}][%d] getDatastoreQueryResponse1NotFound  %+v", 404, o.Payload)
}

func (o *GetDatastoreQueryResponse1NotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetDatastoreQueryResponse1NotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDatastoreQueryResponse1InternalServerError creates a GetDatastoreQueryResponse1InternalServerError with default headers values
func NewGetDatastoreQueryResponse1InternalServerError() *GetDatastoreQueryResponse1InternalServerError {
	return &GetDatastoreQueryResponse1InternalServerError{}
}

/*
GetDatastoreQueryResponse1InternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetDatastoreQueryResponse1InternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this get datastore query response1 internal server error response has a 2xx status code
func (o *GetDatastoreQueryResponse1InternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get datastore query response1 internal server error response has a 3xx status code
func (o *GetDatastoreQueryResponse1InternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get datastore query response1 internal server error response has a 4xx status code
func (o *GetDatastoreQueryResponse1InternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get datastore query response1 internal server error response has a 5xx status code
func (o *GetDatastoreQueryResponse1InternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get datastore query response1 internal server error response a status code equal to that given
func (o *GetDatastoreQueryResponse1InternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get datastore query response1 internal server error response
func (o *GetDatastoreQueryResponse1InternalServerError) Code() int {
	return 500
}

func (o *GetDatastoreQueryResponse1InternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/domains/{domainId}/datastores/queries/{queryId}][%d] getDatastoreQueryResponse1InternalServerError  %+v", 500, o.Payload)
}

func (o *GetDatastoreQueryResponse1InternalServerError) String() string {
	return fmt.Sprintf("[GET /v1/domains/{domainId}/datastores/queries/{queryId}][%d] getDatastoreQueryResponse1InternalServerError  %+v", 500, o.Payload)
}

func (o *GetDatastoreQueryResponse1InternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetDatastoreQueryResponse1InternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
