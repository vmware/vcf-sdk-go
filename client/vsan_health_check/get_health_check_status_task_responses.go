// Code generated by go-swagger; DO NOT EDIT.

// Copyright © 2023 VMware, Inc. All Rights Reserved.
// Code licenced under: BSD-2

package vsan_health_check

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"vcf-sdk-go/models"
)

// GETHealthCheckStatusTaskReader is a Reader for the GETHealthCheckStatusTask structure.
type GETHealthCheckStatusTaskReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GETHealthCheckStatusTaskReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGETHealthCheckStatusTaskOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGETHealthCheckStatusTaskBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGETHealthCheckStatusTaskInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGETHealthCheckStatusTaskOK creates a GETHealthCheckStatusTaskOK with default headers values
func NewGETHealthCheckStatusTaskOK() *GETHealthCheckStatusTaskOK {
	return &GETHealthCheckStatusTaskOK{}
}

/*
GETHealthCheckStatusTaskOK describes a response with status code 200, with default header values.

Ok
*/
type GETHealthCheckStatusTaskOK struct {
	Payload *models.HealthCheckTask
}

// IsSuccess returns true when this get health check status task o k response has a 2xx status code
func (o *GETHealthCheckStatusTaskOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get health check status task o k response has a 3xx status code
func (o *GETHealthCheckStatusTaskOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get health check status task o k response has a 4xx status code
func (o *GETHealthCheckStatusTaskOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get health check status task o k response has a 5xx status code
func (o *GETHealthCheckStatusTaskOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get health check status task o k response a status code equal to that given
func (o *GETHealthCheckStatusTaskOK) IsCode(code int) bool {
	return code == 200
}

func (o *GETHealthCheckStatusTaskOK) Error() string {
	return fmt.Sprintf("[GET /v1/domains/{domainId}/health-checks/tasks/{taskId}][%d] getHealthCheckStatusTaskOK  %+v", 200, o.Payload)
}

func (o *GETHealthCheckStatusTaskOK) String() string {
	return fmt.Sprintf("[GET /v1/domains/{domainId}/health-checks/tasks/{taskId}][%d] getHealthCheckStatusTaskOK  %+v", 200, o.Payload)
}

func (o *GETHealthCheckStatusTaskOK) GetPayload() *models.HealthCheckTask {
	return o.Payload
}

func (o *GETHealthCheckStatusTaskOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HealthCheckTask)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGETHealthCheckStatusTaskBadRequest creates a GETHealthCheckStatusTaskBadRequest with default headers values
func NewGETHealthCheckStatusTaskBadRequest() *GETHealthCheckStatusTaskBadRequest {
	return &GETHealthCheckStatusTaskBadRequest{}
}

/*
GETHealthCheckStatusTaskBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GETHealthCheckStatusTaskBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this get health check status task bad request response has a 2xx status code
func (o *GETHealthCheckStatusTaskBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get health check status task bad request response has a 3xx status code
func (o *GETHealthCheckStatusTaskBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get health check status task bad request response has a 4xx status code
func (o *GETHealthCheckStatusTaskBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get health check status task bad request response has a 5xx status code
func (o *GETHealthCheckStatusTaskBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get health check status task bad request response a status code equal to that given
func (o *GETHealthCheckStatusTaskBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GETHealthCheckStatusTaskBadRequest) Error() string {
	return fmt.Sprintf("[GET /v1/domains/{domainId}/health-checks/tasks/{taskId}][%d] getHealthCheckStatusTaskBadRequest  %+v", 400, o.Payload)
}

func (o *GETHealthCheckStatusTaskBadRequest) String() string {
	return fmt.Sprintf("[GET /v1/domains/{domainId}/health-checks/tasks/{taskId}][%d] getHealthCheckStatusTaskBadRequest  %+v", 400, o.Payload)
}

func (o *GETHealthCheckStatusTaskBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *GETHealthCheckStatusTaskBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGETHealthCheckStatusTaskInternalServerError creates a GETHealthCheckStatusTaskInternalServerError with default headers values
func NewGETHealthCheckStatusTaskInternalServerError() *GETHealthCheckStatusTaskInternalServerError {
	return &GETHealthCheckStatusTaskInternalServerError{}
}

/*
GETHealthCheckStatusTaskInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GETHealthCheckStatusTaskInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this get health check status task internal server error response has a 2xx status code
func (o *GETHealthCheckStatusTaskInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get health check status task internal server error response has a 3xx status code
func (o *GETHealthCheckStatusTaskInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get health check status task internal server error response has a 4xx status code
func (o *GETHealthCheckStatusTaskInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get health check status task internal server error response has a 5xx status code
func (o *GETHealthCheckStatusTaskInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get health check status task internal server error response a status code equal to that given
func (o *GETHealthCheckStatusTaskInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GETHealthCheckStatusTaskInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/domains/{domainId}/health-checks/tasks/{taskId}][%d] getHealthCheckStatusTaskInternalServerError  %+v", 500, o.Payload)
}

func (o *GETHealthCheckStatusTaskInternalServerError) String() string {
	return fmt.Sprintf("[GET /v1/domains/{domainId}/health-checks/tasks/{taskId}][%d] getHealthCheckStatusTaskInternalServerError  %+v", 500, o.Payload)
}

func (o *GETHealthCheckStatusTaskInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *GETHealthCheckStatusTaskInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
