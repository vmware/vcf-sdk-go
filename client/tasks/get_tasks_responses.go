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

// GetTasksReader is a Reader for the GetTasks structure.
type GetTasksReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetTasksReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetTasksOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewGetTasksInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetTasksOK creates a GetTasksOK with default headers values
func NewGetTasksOK() *GetTasksOK {
	return &GetTasksOK{}
}

/*
GetTasksOK describes a response with status code 200, with default header values.

Returns the list of tasks.
*/
type GetTasksOK struct {
	Payload *models.PageOfTask
}

// IsSuccess returns true when this get tasks o k response has a 2xx status code
func (o *GetTasksOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get tasks o k response has a 3xx status code
func (o *GetTasksOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get tasks o k response has a 4xx status code
func (o *GetTasksOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get tasks o k response has a 5xx status code
func (o *GetTasksOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get tasks o k response a status code equal to that given
func (o *GetTasksOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetTasksOK) Error() string {
	return fmt.Sprintf("[GET /v1/tasks][%d] getTasksOK  %+v", 200, o.Payload)
}

func (o *GetTasksOK) String() string {
	return fmt.Sprintf("[GET /v1/tasks][%d] getTasksOK  %+v", 200, o.Payload)
}

func (o *GetTasksOK) GetPayload() *models.PageOfTask {
	return o.Payload
}

func (o *GetTasksOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PageOfTask)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTasksInternalServerError creates a GetTasksInternalServerError with default headers values
func NewGetTasksInternalServerError() *GetTasksInternalServerError {
	return &GetTasksInternalServerError{}
}

/*
GetTasksInternalServerError describes a response with status code 500, with default header values.

Unexpected error
*/
type GetTasksInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this get tasks internal server error response has a 2xx status code
func (o *GetTasksInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get tasks internal server error response has a 3xx status code
func (o *GetTasksInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get tasks internal server error response has a 4xx status code
func (o *GetTasksInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get tasks internal server error response has a 5xx status code
func (o *GetTasksInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get tasks internal server error response a status code equal to that given
func (o *GetTasksInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetTasksInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/tasks][%d] getTasksInternalServerError  %+v", 500, o.Payload)
}

func (o *GetTasksInternalServerError) String() string {
	return fmt.Sprintf("[GET /v1/tasks][%d] getTasksInternalServerError  %+v", 500, o.Payload)
}

func (o *GetTasksInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetTasksInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}