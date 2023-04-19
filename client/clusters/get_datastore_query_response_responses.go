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

// GETDatastoreQueryResponseReader is a Reader for the GETDatastoreQueryResponse structure.
type GETDatastoreQueryResponseReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GETDatastoreQueryResponseReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGETDatastoreQueryResponseOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGETDatastoreQueryResponseBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGETDatastoreQueryResponseNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGETDatastoreQueryResponseInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGETDatastoreQueryResponseOK creates a GETDatastoreQueryResponseOK with default headers values
func NewGETDatastoreQueryResponseOK() *GETDatastoreQueryResponseOK {
	return &GETDatastoreQueryResponseOK{}
}

/*
GETDatastoreQueryResponseOK describes a response with status code 200, with default header values.

Ok
*/
type GETDatastoreQueryResponseOK struct {
	Payload *models.DatastoreQueryResponse
}

// IsSuccess returns true when this get datastore query response o k response has a 2xx status code
func (o *GETDatastoreQueryResponseOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get datastore query response o k response has a 3xx status code
func (o *GETDatastoreQueryResponseOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get datastore query response o k response has a 4xx status code
func (o *GETDatastoreQueryResponseOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get datastore query response o k response has a 5xx status code
func (o *GETDatastoreQueryResponseOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get datastore query response o k response a status code equal to that given
func (o *GETDatastoreQueryResponseOK) IsCode(code int) bool {
	return code == 200
}

func (o *GETDatastoreQueryResponseOK) Error() string {
	return fmt.Sprintf("[GET /v1/clusters/{clusterId}/datastores/queries/{queryId}][%d] getDatastoreQueryResponseOK  %+v", 200, o.Payload)
}

func (o *GETDatastoreQueryResponseOK) String() string {
	return fmt.Sprintf("[GET /v1/clusters/{clusterId}/datastores/queries/{queryId}][%d] getDatastoreQueryResponseOK  %+v", 200, o.Payload)
}

func (o *GETDatastoreQueryResponseOK) GetPayload() *models.DatastoreQueryResponse {
	return o.Payload
}

func (o *GETDatastoreQueryResponseOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DatastoreQueryResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGETDatastoreQueryResponseBadRequest creates a GETDatastoreQueryResponseBadRequest with default headers values
func NewGETDatastoreQueryResponseBadRequest() *GETDatastoreQueryResponseBadRequest {
	return &GETDatastoreQueryResponseBadRequest{}
}

/*
GETDatastoreQueryResponseBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GETDatastoreQueryResponseBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this get datastore query response bad request response has a 2xx status code
func (o *GETDatastoreQueryResponseBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get datastore query response bad request response has a 3xx status code
func (o *GETDatastoreQueryResponseBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get datastore query response bad request response has a 4xx status code
func (o *GETDatastoreQueryResponseBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get datastore query response bad request response has a 5xx status code
func (o *GETDatastoreQueryResponseBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get datastore query response bad request response a status code equal to that given
func (o *GETDatastoreQueryResponseBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GETDatastoreQueryResponseBadRequest) Error() string {
	return fmt.Sprintf("[GET /v1/clusters/{clusterId}/datastores/queries/{queryId}][%d] getDatastoreQueryResponseBadRequest  %+v", 400, o.Payload)
}

func (o *GETDatastoreQueryResponseBadRequest) String() string {
	return fmt.Sprintf("[GET /v1/clusters/{clusterId}/datastores/queries/{queryId}][%d] getDatastoreQueryResponseBadRequest  %+v", 400, o.Payload)
}

func (o *GETDatastoreQueryResponseBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *GETDatastoreQueryResponseBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGETDatastoreQueryResponseNotFound creates a GETDatastoreQueryResponseNotFound with default headers values
func NewGETDatastoreQueryResponseNotFound() *GETDatastoreQueryResponseNotFound {
	return &GETDatastoreQueryResponseNotFound{}
}

/*
GETDatastoreQueryResponseNotFound describes a response with status code 404, with default header values.

Query Not Found
*/
type GETDatastoreQueryResponseNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this get datastore query response not found response has a 2xx status code
func (o *GETDatastoreQueryResponseNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get datastore query response not found response has a 3xx status code
func (o *GETDatastoreQueryResponseNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get datastore query response not found response has a 4xx status code
func (o *GETDatastoreQueryResponseNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get datastore query response not found response has a 5xx status code
func (o *GETDatastoreQueryResponseNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get datastore query response not found response a status code equal to that given
func (o *GETDatastoreQueryResponseNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GETDatastoreQueryResponseNotFound) Error() string {
	return fmt.Sprintf("[GET /v1/clusters/{clusterId}/datastores/queries/{queryId}][%d] getDatastoreQueryResponseNotFound  %+v", 404, o.Payload)
}

func (o *GETDatastoreQueryResponseNotFound) String() string {
	return fmt.Sprintf("[GET /v1/clusters/{clusterId}/datastores/queries/{queryId}][%d] getDatastoreQueryResponseNotFound  %+v", 404, o.Payload)
}

func (o *GETDatastoreQueryResponseNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GETDatastoreQueryResponseNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGETDatastoreQueryResponseInternalServerError creates a GETDatastoreQueryResponseInternalServerError with default headers values
func NewGETDatastoreQueryResponseInternalServerError() *GETDatastoreQueryResponseInternalServerError {
	return &GETDatastoreQueryResponseInternalServerError{}
}

/*
GETDatastoreQueryResponseInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GETDatastoreQueryResponseInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this get datastore query response internal server error response has a 2xx status code
func (o *GETDatastoreQueryResponseInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get datastore query response internal server error response has a 3xx status code
func (o *GETDatastoreQueryResponseInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get datastore query response internal server error response has a 4xx status code
func (o *GETDatastoreQueryResponseInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get datastore query response internal server error response has a 5xx status code
func (o *GETDatastoreQueryResponseInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get datastore query response internal server error response a status code equal to that given
func (o *GETDatastoreQueryResponseInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GETDatastoreQueryResponseInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/clusters/{clusterId}/datastores/queries/{queryId}][%d] getDatastoreQueryResponseInternalServerError  %+v", 500, o.Payload)
}

func (o *GETDatastoreQueryResponseInternalServerError) String() string {
	return fmt.Sprintf("[GET /v1/clusters/{clusterId}/datastores/queries/{queryId}][%d] getDatastoreQueryResponseInternalServerError  %+v", 500, o.Payload)
}

func (o *GETDatastoreQueryResponseInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *GETDatastoreQueryResponseInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
