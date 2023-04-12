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

// CancelTaskReader is a Reader for the CancelTask structure.
type CancelTaskReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CancelTaskReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCancelTaskOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewCancelTaskNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewCancelTaskConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCancelTaskInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCancelTaskOK creates a CancelTaskOK with default headers values
func NewCancelTaskOK() *CancelTaskOK {
	return &CancelTaskOK{}
}

/*
CancelTaskOK describes a response with status code 200, with default header values.

Task was cancelled successfully.
*/
type CancelTaskOK struct {
}

// IsSuccess returns true when this cancel task o k response has a 2xx status code
func (o *CancelTaskOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this cancel task o k response has a 3xx status code
func (o *CancelTaskOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cancel task o k response has a 4xx status code
func (o *CancelTaskOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this cancel task o k response has a 5xx status code
func (o *CancelTaskOK) IsServerError() bool {
	return false
}

// IsCode returns true when this cancel task o k response a status code equal to that given
func (o *CancelTaskOK) IsCode(code int) bool {
	return code == 200
}

func (o *CancelTaskOK) Error() string {
	return fmt.Sprintf("[DELETE /v1/tasks/{id}][%d] cancelTaskOK ", 200)
}

func (o *CancelTaskOK) String() string {
	return fmt.Sprintf("[DELETE /v1/tasks/{id}][%d] cancelTaskOK ", 200)
}

func (o *CancelTaskOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCancelTaskNotFound creates a CancelTaskNotFound with default headers values
func NewCancelTaskNotFound() *CancelTaskNotFound {
	return &CancelTaskNotFound{}
}

/*
CancelTaskNotFound describes a response with status code 404, with default header values.

Task not found
*/
type CancelTaskNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this cancel task not found response has a 2xx status code
func (o *CancelTaskNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this cancel task not found response has a 3xx status code
func (o *CancelTaskNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cancel task not found response has a 4xx status code
func (o *CancelTaskNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this cancel task not found response has a 5xx status code
func (o *CancelTaskNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this cancel task not found response a status code equal to that given
func (o *CancelTaskNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *CancelTaskNotFound) Error() string {
	return fmt.Sprintf("[DELETE /v1/tasks/{id}][%d] cancelTaskNotFound  %+v", 404, o.Payload)
}

func (o *CancelTaskNotFound) String() string {
	return fmt.Sprintf("[DELETE /v1/tasks/{id}][%d] cancelTaskNotFound  %+v", 404, o.Payload)
}

func (o *CancelTaskNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *CancelTaskNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCancelTaskConflict creates a CancelTaskConflict with default headers values
func NewCancelTaskConflict() *CancelTaskConflict {
	return &CancelTaskConflict{}
}

/*
CancelTaskConflict describes a response with status code 409, with default header values.

Task can not be cancelled. Only a IN_PROGRESS task can be cancelled.
*/
type CancelTaskConflict struct {
	Payload *models.Error
}

// IsSuccess returns true when this cancel task conflict response has a 2xx status code
func (o *CancelTaskConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this cancel task conflict response has a 3xx status code
func (o *CancelTaskConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cancel task conflict response has a 4xx status code
func (o *CancelTaskConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this cancel task conflict response has a 5xx status code
func (o *CancelTaskConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this cancel task conflict response a status code equal to that given
func (o *CancelTaskConflict) IsCode(code int) bool {
	return code == 409
}

func (o *CancelTaskConflict) Error() string {
	return fmt.Sprintf("[DELETE /v1/tasks/{id}][%d] cancelTaskConflict  %+v", 409, o.Payload)
}

func (o *CancelTaskConflict) String() string {
	return fmt.Sprintf("[DELETE /v1/tasks/{id}][%d] cancelTaskConflict  %+v", 409, o.Payload)
}

func (o *CancelTaskConflict) GetPayload() *models.Error {
	return o.Payload
}

func (o *CancelTaskConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCancelTaskInternalServerError creates a CancelTaskInternalServerError with default headers values
func NewCancelTaskInternalServerError() *CancelTaskInternalServerError {
	return &CancelTaskInternalServerError{}
}

/*
CancelTaskInternalServerError describes a response with status code 500, with default header values.

Unexpected error
*/
type CancelTaskInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this cancel task internal server error response has a 2xx status code
func (o *CancelTaskInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this cancel task internal server error response has a 3xx status code
func (o *CancelTaskInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cancel task internal server error response has a 4xx status code
func (o *CancelTaskInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this cancel task internal server error response has a 5xx status code
func (o *CancelTaskInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this cancel task internal server error response a status code equal to that given
func (o *CancelTaskInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *CancelTaskInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /v1/tasks/{id}][%d] cancelTaskInternalServerError  %+v", 500, o.Payload)
}

func (o *CancelTaskInternalServerError) String() string {
	return fmt.Sprintf("[DELETE /v1/tasks/{id}][%d] cancelTaskInternalServerError  %+v", 500, o.Payload)
}

func (o *CancelTaskInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *CancelTaskInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}