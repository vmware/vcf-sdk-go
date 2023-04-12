// Code generated by go-swagger; DO NOT EDIT.

package tasks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"vcf-sdk-go/models"
)

// GetTaskReader is a Reader for the GetTask structure.
type GetTaskReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetTaskReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetTaskOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetTaskNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetTaskInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetTaskOK creates a GetTaskOK with default headers values
func NewGetTaskOK() *GetTaskOK {
	return &GetTaskOK{}
}

/*
GetTaskOK describes a response with status code 200, with default header values.

A task object.
*/
type GetTaskOK struct {
	Payload *models.Task
}

// IsSuccess returns true when this get task o k response has a 2xx status code
func (o *GetTaskOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get task o k response has a 3xx status code
func (o *GetTaskOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get task o k response has a 4xx status code
func (o *GetTaskOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get task o k response has a 5xx status code
func (o *GetTaskOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get task o k response a status code equal to that given
func (o *GetTaskOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetTaskOK) Error() string {
	return fmt.Sprintf("[GET /v1/tasks/{id}][%d] getTaskOK  %+v", 200, o.Payload)
}

func (o *GetTaskOK) String() string {
	return fmt.Sprintf("[GET /v1/tasks/{id}][%d] getTaskOK  %+v", 200, o.Payload)
}

func (o *GetTaskOK) GetPayload() *models.Task {
	return o.Payload
}

func (o *GetTaskOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Task)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTaskNotFound creates a GetTaskNotFound with default headers values
func NewGetTaskNotFound() *GetTaskNotFound {
	return &GetTaskNotFound{}
}

/*
GetTaskNotFound describes a response with status code 404, with default header values.

Task not found
*/
type GetTaskNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this get task not found response has a 2xx status code
func (o *GetTaskNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get task not found response has a 3xx status code
func (o *GetTaskNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get task not found response has a 4xx status code
func (o *GetTaskNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get task not found response has a 5xx status code
func (o *GetTaskNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get task not found response a status code equal to that given
func (o *GetTaskNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetTaskNotFound) Error() string {
	return fmt.Sprintf("[GET /v1/tasks/{id}][%d] getTaskNotFound  %+v", 404, o.Payload)
}

func (o *GetTaskNotFound) String() string {
	return fmt.Sprintf("[GET /v1/tasks/{id}][%d] getTaskNotFound  %+v", 404, o.Payload)
}

func (o *GetTaskNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetTaskNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTaskInternalServerError creates a GetTaskInternalServerError with default headers values
func NewGetTaskInternalServerError() *GetTaskInternalServerError {
	return &GetTaskInternalServerError{}
}

/*
GetTaskInternalServerError describes a response with status code 500, with default header values.

Unexpected error
*/
type GetTaskInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this get task internal server error response has a 2xx status code
func (o *GetTaskInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get task internal server error response has a 3xx status code
func (o *GetTaskInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get task internal server error response has a 4xx status code
func (o *GetTaskInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get task internal server error response has a 5xx status code
func (o *GetTaskInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get task internal server error response a status code equal to that given
func (o *GetTaskInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetTaskInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/tasks/{id}][%d] getTaskInternalServerError  %+v", 500, o.Payload)
}

func (o *GetTaskInternalServerError) String() string {
	return fmt.Sprintf("[GET /v1/tasks/{id}][%d] getTaskInternalServerError  %+v", 500, o.Payload)
}

func (o *GetTaskInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetTaskInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
