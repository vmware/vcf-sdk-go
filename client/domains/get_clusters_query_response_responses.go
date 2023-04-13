// Code generated by go-swagger; DO NOT EDIT.

// Copyright © 2023 VMware, Inc. All Rights Reserved.
// Code licenced under: BSD-2

package domains

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"vcf-sdk-go/models"
)

// GETClustersQueryResponseReader is a Reader for the GETClustersQueryResponse structure.
type GETClustersQueryResponseReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GETClustersQueryResponseReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGETClustersQueryResponseOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGETClustersQueryResponseBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGETClustersQueryResponseNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGETClustersQueryResponseInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGETClustersQueryResponseOK creates a GETClustersQueryResponseOK with default headers values
func NewGETClustersQueryResponseOK() *GETClustersQueryResponseOK {
	return &GETClustersQueryResponseOK{}
}

/*
GETClustersQueryResponseOK describes a response with status code 200, with default header values.

Ok
*/
type GETClustersQueryResponseOK struct {
	Payload *models.ClusterQueryResponse
}

// IsSuccess returns true when this get clusters query response o k response has a 2xx status code
func (o *GETClustersQueryResponseOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get clusters query response o k response has a 3xx status code
func (o *GETClustersQueryResponseOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get clusters query response o k response has a 4xx status code
func (o *GETClustersQueryResponseOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get clusters query response o k response has a 5xx status code
func (o *GETClustersQueryResponseOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get clusters query response o k response a status code equal to that given
func (o *GETClustersQueryResponseOK) IsCode(code int) bool {
	return code == 200
}

func (o *GETClustersQueryResponseOK) Error() string {
	return fmt.Sprintf("[GET /v1/domains/{domainId}/clusters/queries/{queryId}][%d] getClustersQueryResponseOK  %+v", 200, o.Payload)
}

func (o *GETClustersQueryResponseOK) String() string {
	return fmt.Sprintf("[GET /v1/domains/{domainId}/clusters/queries/{queryId}][%d] getClustersQueryResponseOK  %+v", 200, o.Payload)
}

func (o *GETClustersQueryResponseOK) GetPayload() *models.ClusterQueryResponse {
	return o.Payload
}

func (o *GETClustersQueryResponseOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ClusterQueryResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGETClustersQueryResponseBadRequest creates a GETClustersQueryResponseBadRequest with default headers values
func NewGETClustersQueryResponseBadRequest() *GETClustersQueryResponseBadRequest {
	return &GETClustersQueryResponseBadRequest{}
}

/*
GETClustersQueryResponseBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GETClustersQueryResponseBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this get clusters query response bad request response has a 2xx status code
func (o *GETClustersQueryResponseBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get clusters query response bad request response has a 3xx status code
func (o *GETClustersQueryResponseBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get clusters query response bad request response has a 4xx status code
func (o *GETClustersQueryResponseBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get clusters query response bad request response has a 5xx status code
func (o *GETClustersQueryResponseBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get clusters query response bad request response a status code equal to that given
func (o *GETClustersQueryResponseBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GETClustersQueryResponseBadRequest) Error() string {
	return fmt.Sprintf("[GET /v1/domains/{domainId}/clusters/queries/{queryId}][%d] getClustersQueryResponseBadRequest  %+v", 400, o.Payload)
}

func (o *GETClustersQueryResponseBadRequest) String() string {
	return fmt.Sprintf("[GET /v1/domains/{domainId}/clusters/queries/{queryId}][%d] getClustersQueryResponseBadRequest  %+v", 400, o.Payload)
}

func (o *GETClustersQueryResponseBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *GETClustersQueryResponseBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGETClustersQueryResponseNotFound creates a GETClustersQueryResponseNotFound with default headers values
func NewGETClustersQueryResponseNotFound() *GETClustersQueryResponseNotFound {
	return &GETClustersQueryResponseNotFound{}
}

/*
GETClustersQueryResponseNotFound describes a response with status code 404, with default header values.

Query Not Found
*/
type GETClustersQueryResponseNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this get clusters query response not found response has a 2xx status code
func (o *GETClustersQueryResponseNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get clusters query response not found response has a 3xx status code
func (o *GETClustersQueryResponseNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get clusters query response not found response has a 4xx status code
func (o *GETClustersQueryResponseNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get clusters query response not found response has a 5xx status code
func (o *GETClustersQueryResponseNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get clusters query response not found response a status code equal to that given
func (o *GETClustersQueryResponseNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GETClustersQueryResponseNotFound) Error() string {
	return fmt.Sprintf("[GET /v1/domains/{domainId}/clusters/queries/{queryId}][%d] getClustersQueryResponseNotFound  %+v", 404, o.Payload)
}

func (o *GETClustersQueryResponseNotFound) String() string {
	return fmt.Sprintf("[GET /v1/domains/{domainId}/clusters/queries/{queryId}][%d] getClustersQueryResponseNotFound  %+v", 404, o.Payload)
}

func (o *GETClustersQueryResponseNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GETClustersQueryResponseNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGETClustersQueryResponseInternalServerError creates a GETClustersQueryResponseInternalServerError with default headers values
func NewGETClustersQueryResponseInternalServerError() *GETClustersQueryResponseInternalServerError {
	return &GETClustersQueryResponseInternalServerError{}
}

/*
GETClustersQueryResponseInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GETClustersQueryResponseInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this get clusters query response internal server error response has a 2xx status code
func (o *GETClustersQueryResponseInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get clusters query response internal server error response has a 3xx status code
func (o *GETClustersQueryResponseInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get clusters query response internal server error response has a 4xx status code
func (o *GETClustersQueryResponseInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get clusters query response internal server error response has a 5xx status code
func (o *GETClustersQueryResponseInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get clusters query response internal server error response a status code equal to that given
func (o *GETClustersQueryResponseInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GETClustersQueryResponseInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/domains/{domainId}/clusters/queries/{queryId}][%d] getClustersQueryResponseInternalServerError  %+v", 500, o.Payload)
}

func (o *GETClustersQueryResponseInternalServerError) String() string {
	return fmt.Sprintf("[GET /v1/domains/{domainId}/clusters/queries/{queryId}][%d] getClustersQueryResponseInternalServerError  %+v", 500, o.Payload)
}

func (o *GETClustersQueryResponseInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *GETClustersQueryResponseInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
