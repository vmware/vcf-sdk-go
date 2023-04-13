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

// POSTDatastoreQueryReader is a Reader for the POSTDatastoreQuery structure.
type POSTDatastoreQueryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *POSTDatastoreQueryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPOSTDatastoreQueryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPOSTDatastoreQueryBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPOSTDatastoreQueryInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPOSTDatastoreQueryOK creates a POSTDatastoreQueryOK with default headers values
func NewPOSTDatastoreQueryOK() *POSTDatastoreQueryOK {
	return &POSTDatastoreQueryOK{}
}

/*
POSTDatastoreQueryOK describes a response with status code 200, with default header values.

Ok
*/
type POSTDatastoreQueryOK struct {
	Payload *models.DatastoreQueryResponse
}

// IsSuccess returns true when this post datastore query o k response has a 2xx status code
func (o *POSTDatastoreQueryOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post datastore query o k response has a 3xx status code
func (o *POSTDatastoreQueryOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post datastore query o k response has a 4xx status code
func (o *POSTDatastoreQueryOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this post datastore query o k response has a 5xx status code
func (o *POSTDatastoreQueryOK) IsServerError() bool {
	return false
}

// IsCode returns true when this post datastore query o k response a status code equal to that given
func (o *POSTDatastoreQueryOK) IsCode(code int) bool {
	return code == 200
}

func (o *POSTDatastoreQueryOK) Error() string {
	return fmt.Sprintf("[POST /v1/clusters/{id}/datastores/queries][%d] postDatastoreQueryOK  %+v", 200, o.Payload)
}

func (o *POSTDatastoreQueryOK) String() string {
	return fmt.Sprintf("[POST /v1/clusters/{id}/datastores/queries][%d] postDatastoreQueryOK  %+v", 200, o.Payload)
}

func (o *POSTDatastoreQueryOK) GetPayload() *models.DatastoreQueryResponse {
	return o.Payload
}

func (o *POSTDatastoreQueryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DatastoreQueryResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPOSTDatastoreQueryBadRequest creates a POSTDatastoreQueryBadRequest with default headers values
func NewPOSTDatastoreQueryBadRequest() *POSTDatastoreQueryBadRequest {
	return &POSTDatastoreQueryBadRequest{}
}

/*
POSTDatastoreQueryBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type POSTDatastoreQueryBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this post datastore query bad request response has a 2xx status code
func (o *POSTDatastoreQueryBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post datastore query bad request response has a 3xx status code
func (o *POSTDatastoreQueryBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post datastore query bad request response has a 4xx status code
func (o *POSTDatastoreQueryBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this post datastore query bad request response has a 5xx status code
func (o *POSTDatastoreQueryBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this post datastore query bad request response a status code equal to that given
func (o *POSTDatastoreQueryBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *POSTDatastoreQueryBadRequest) Error() string {
	return fmt.Sprintf("[POST /v1/clusters/{id}/datastores/queries][%d] postDatastoreQueryBadRequest  %+v", 400, o.Payload)
}

func (o *POSTDatastoreQueryBadRequest) String() string {
	return fmt.Sprintf("[POST /v1/clusters/{id}/datastores/queries][%d] postDatastoreQueryBadRequest  %+v", 400, o.Payload)
}

func (o *POSTDatastoreQueryBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *POSTDatastoreQueryBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPOSTDatastoreQueryInternalServerError creates a POSTDatastoreQueryInternalServerError with default headers values
func NewPOSTDatastoreQueryInternalServerError() *POSTDatastoreQueryInternalServerError {
	return &POSTDatastoreQueryInternalServerError{}
}

/*
POSTDatastoreQueryInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type POSTDatastoreQueryInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this post datastore query internal server error response has a 2xx status code
func (o *POSTDatastoreQueryInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post datastore query internal server error response has a 3xx status code
func (o *POSTDatastoreQueryInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post datastore query internal server error response has a 4xx status code
func (o *POSTDatastoreQueryInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this post datastore query internal server error response has a 5xx status code
func (o *POSTDatastoreQueryInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this post datastore query internal server error response a status code equal to that given
func (o *POSTDatastoreQueryInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *POSTDatastoreQueryInternalServerError) Error() string {
	return fmt.Sprintf("[POST /v1/clusters/{id}/datastores/queries][%d] postDatastoreQueryInternalServerError  %+v", 500, o.Payload)
}

func (o *POSTDatastoreQueryInternalServerError) String() string {
	return fmt.Sprintf("[POST /v1/clusters/{id}/datastores/queries][%d] postDatastoreQueryInternalServerError  %+v", 500, o.Payload)
}

func (o *POSTDatastoreQueryInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *POSTDatastoreQueryInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
