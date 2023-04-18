// Code generated by go-swagger; DO NOT EDIT.

// Copyright © 2023 VMware, Inc. All Rights Reserved.
// Code licenced under: BSD-2

package tasks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/vmware/vcf-sdk-go/models"
)

// GETTaskReader is a Reader for the GETTask structure.
type GETTaskReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GETTaskReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGETTaskOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGETTaskNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGETTaskInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGETTaskOK creates a GETTaskOK with default headers values
func NewGETTaskOK() *GETTaskOK {
	return &GETTaskOK{}
}

/*
GETTaskOK describes a response with status code 200, with default header values.

A task object.
*/
type GETTaskOK struct {
	Payload *models.Task
}

// IsSuccess returns true when this get task o k response has a 2xx status code
func (o *GETTaskOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get task o k response has a 3xx status code
func (o *GETTaskOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get task o k response has a 4xx status code
func (o *GETTaskOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get task o k response has a 5xx status code
func (o *GETTaskOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get task o k response a status code equal to that given
func (o *GETTaskOK) IsCode(code int) bool {
	return code == 200
}

func (o *GETTaskOK) Error() string {
	return fmt.Sprintf("[GET /v1/tasks/{id}][%d] getTaskOK  %+v", 200, o.Payload)
}

func (o *GETTaskOK) String() string {
	return fmt.Sprintf("[GET /v1/tasks/{id}][%d] getTaskOK  %+v", 200, o.Payload)
}

func (o *GETTaskOK) GetPayload() *models.Task {
	return o.Payload
}

func (o *GETTaskOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Task)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGETTaskNotFound creates a GETTaskNotFound with default headers values
func NewGETTaskNotFound() *GETTaskNotFound {
	return &GETTaskNotFound{}
}

/*
GETTaskNotFound describes a response with status code 404, with default header values.

Task not found
*/
type GETTaskNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this get task not found response has a 2xx status code
func (o *GETTaskNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get task not found response has a 3xx status code
func (o *GETTaskNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get task not found response has a 4xx status code
func (o *GETTaskNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get task not found response has a 5xx status code
func (o *GETTaskNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get task not found response a status code equal to that given
func (o *GETTaskNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GETTaskNotFound) Error() string {
	return fmt.Sprintf("[GET /v1/tasks/{id}][%d] getTaskNotFound  %+v", 404, o.Payload)
}

func (o *GETTaskNotFound) String() string {
	return fmt.Sprintf("[GET /v1/tasks/{id}][%d] getTaskNotFound  %+v", 404, o.Payload)
}

func (o *GETTaskNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GETTaskNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGETTaskInternalServerError creates a GETTaskInternalServerError with default headers values
func NewGETTaskInternalServerError() *GETTaskInternalServerError {
	return &GETTaskInternalServerError{}
}

/*
GETTaskInternalServerError describes a response with status code 500, with default header values.

Unexpected error
*/
type GETTaskInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this get task internal server error response has a 2xx status code
func (o *GETTaskInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get task internal server error response has a 3xx status code
func (o *GETTaskInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get task internal server error response has a 4xx status code
func (o *GETTaskInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get task internal server error response has a 5xx status code
func (o *GETTaskInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get task internal server error response a status code equal to that given
func (o *GETTaskInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GETTaskInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/tasks/{id}][%d] getTaskInternalServerError  %+v", 500, o.Payload)
}

func (o *GETTaskInternalServerError) String() string {
	return fmt.Sprintf("[GET /v1/tasks/{id}][%d] getTaskInternalServerError  %+v", 500, o.Payload)
}

func (o *GETTaskInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *GETTaskInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
