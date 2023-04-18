// Code generated by go-swagger; DO NOT EDIT.

// Copyright © 2023 VMware, Inc. All Rights Reserved.
// Code licenced under: BSD-2

package credentials

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/vmware/vcf-sdk-go/models"
)

// RetryCredentialsTaskReader is a Reader for the RetryCredentialsTask structure.
type RetryCredentialsTaskReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RetryCredentialsTaskReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRetryCredentialsTaskOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewRetryCredentialsTaskAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewRetryCredentialsTaskBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewRetryCredentialsTaskUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewRetryCredentialsTaskForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewRetryCredentialsTaskInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewRetryCredentialsTaskOK creates a RetryCredentialsTaskOK with default headers values
func NewRetryCredentialsTaskOK() *RetryCredentialsTaskOK {
	return &RetryCredentialsTaskOK{}
}

/*
RetryCredentialsTaskOK describes a response with status code 200, with default header values.

OK
*/
type RetryCredentialsTaskOK struct {
	Payload *models.Task
}

// IsSuccess returns true when this retry credentials task o k response has a 2xx status code
func (o *RetryCredentialsTaskOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this retry credentials task o k response has a 3xx status code
func (o *RetryCredentialsTaskOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this retry credentials task o k response has a 4xx status code
func (o *RetryCredentialsTaskOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this retry credentials task o k response has a 5xx status code
func (o *RetryCredentialsTaskOK) IsServerError() bool {
	return false
}

// IsCode returns true when this retry credentials task o k response a status code equal to that given
func (o *RetryCredentialsTaskOK) IsCode(code int) bool {
	return code == 200
}

func (o *RetryCredentialsTaskOK) Error() string {
	return fmt.Sprintf("[PATCH /v1/credentials/tasks/{id}][%d] retryCredentialsTaskOK  %+v", 200, o.Payload)
}

func (o *RetryCredentialsTaskOK) String() string {
	return fmt.Sprintf("[PATCH /v1/credentials/tasks/{id}][%d] retryCredentialsTaskOK  %+v", 200, o.Payload)
}

func (o *RetryCredentialsTaskOK) GetPayload() *models.Task {
	return o.Payload
}

func (o *RetryCredentialsTaskOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Task)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRetryCredentialsTaskAccepted creates a RetryCredentialsTaskAccepted with default headers values
func NewRetryCredentialsTaskAccepted() *RetryCredentialsTaskAccepted {
	return &RetryCredentialsTaskAccepted{}
}

/*
RetryCredentialsTaskAccepted describes a response with status code 202, with default header values.

Accepted
*/
type RetryCredentialsTaskAccepted struct {
	Payload *models.Task
}

// IsSuccess returns true when this retry credentials task accepted response has a 2xx status code
func (o *RetryCredentialsTaskAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this retry credentials task accepted response has a 3xx status code
func (o *RetryCredentialsTaskAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this retry credentials task accepted response has a 4xx status code
func (o *RetryCredentialsTaskAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this retry credentials task accepted response has a 5xx status code
func (o *RetryCredentialsTaskAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this retry credentials task accepted response a status code equal to that given
func (o *RetryCredentialsTaskAccepted) IsCode(code int) bool {
	return code == 202
}

func (o *RetryCredentialsTaskAccepted) Error() string {
	return fmt.Sprintf("[PATCH /v1/credentials/tasks/{id}][%d] retryCredentialsTaskAccepted  %+v", 202, o.Payload)
}

func (o *RetryCredentialsTaskAccepted) String() string {
	return fmt.Sprintf("[PATCH /v1/credentials/tasks/{id}][%d] retryCredentialsTaskAccepted  %+v", 202, o.Payload)
}

func (o *RetryCredentialsTaskAccepted) GetPayload() *models.Task {
	return o.Payload
}

func (o *RetryCredentialsTaskAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Task)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRetryCredentialsTaskBadRequest creates a RetryCredentialsTaskBadRequest with default headers values
func NewRetryCredentialsTaskBadRequest() *RetryCredentialsTaskBadRequest {
	return &RetryCredentialsTaskBadRequest{}
}

/*
RetryCredentialsTaskBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type RetryCredentialsTaskBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this retry credentials task bad request response has a 2xx status code
func (o *RetryCredentialsTaskBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this retry credentials task bad request response has a 3xx status code
func (o *RetryCredentialsTaskBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this retry credentials task bad request response has a 4xx status code
func (o *RetryCredentialsTaskBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this retry credentials task bad request response has a 5xx status code
func (o *RetryCredentialsTaskBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this retry credentials task bad request response a status code equal to that given
func (o *RetryCredentialsTaskBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *RetryCredentialsTaskBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /v1/credentials/tasks/{id}][%d] retryCredentialsTaskBadRequest  %+v", 400, o.Payload)
}

func (o *RetryCredentialsTaskBadRequest) String() string {
	return fmt.Sprintf("[PATCH /v1/credentials/tasks/{id}][%d] retryCredentialsTaskBadRequest  %+v", 400, o.Payload)
}

func (o *RetryCredentialsTaskBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *RetryCredentialsTaskBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRetryCredentialsTaskUnauthorized creates a RetryCredentialsTaskUnauthorized with default headers values
func NewRetryCredentialsTaskUnauthorized() *RetryCredentialsTaskUnauthorized {
	return &RetryCredentialsTaskUnauthorized{}
}

/*
RetryCredentialsTaskUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type RetryCredentialsTaskUnauthorized struct {
	Payload *models.Error
}

// IsSuccess returns true when this retry credentials task unauthorized response has a 2xx status code
func (o *RetryCredentialsTaskUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this retry credentials task unauthorized response has a 3xx status code
func (o *RetryCredentialsTaskUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this retry credentials task unauthorized response has a 4xx status code
func (o *RetryCredentialsTaskUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this retry credentials task unauthorized response has a 5xx status code
func (o *RetryCredentialsTaskUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this retry credentials task unauthorized response a status code equal to that given
func (o *RetryCredentialsTaskUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *RetryCredentialsTaskUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /v1/credentials/tasks/{id}][%d] retryCredentialsTaskUnauthorized  %+v", 401, o.Payload)
}

func (o *RetryCredentialsTaskUnauthorized) String() string {
	return fmt.Sprintf("[PATCH /v1/credentials/tasks/{id}][%d] retryCredentialsTaskUnauthorized  %+v", 401, o.Payload)
}

func (o *RetryCredentialsTaskUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *RetryCredentialsTaskUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRetryCredentialsTaskForbidden creates a RetryCredentialsTaskForbidden with default headers values
func NewRetryCredentialsTaskForbidden() *RetryCredentialsTaskForbidden {
	return &RetryCredentialsTaskForbidden{}
}

/*
RetryCredentialsTaskForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type RetryCredentialsTaskForbidden struct {
	Payload *models.Error
}

// IsSuccess returns true when this retry credentials task forbidden response has a 2xx status code
func (o *RetryCredentialsTaskForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this retry credentials task forbidden response has a 3xx status code
func (o *RetryCredentialsTaskForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this retry credentials task forbidden response has a 4xx status code
func (o *RetryCredentialsTaskForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this retry credentials task forbidden response has a 5xx status code
func (o *RetryCredentialsTaskForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this retry credentials task forbidden response a status code equal to that given
func (o *RetryCredentialsTaskForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *RetryCredentialsTaskForbidden) Error() string {
	return fmt.Sprintf("[PATCH /v1/credentials/tasks/{id}][%d] retryCredentialsTaskForbidden  %+v", 403, o.Payload)
}

func (o *RetryCredentialsTaskForbidden) String() string {
	return fmt.Sprintf("[PATCH /v1/credentials/tasks/{id}][%d] retryCredentialsTaskForbidden  %+v", 403, o.Payload)
}

func (o *RetryCredentialsTaskForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *RetryCredentialsTaskForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRetryCredentialsTaskInternalServerError creates a RetryCredentialsTaskInternalServerError with default headers values
func NewRetryCredentialsTaskInternalServerError() *RetryCredentialsTaskInternalServerError {
	return &RetryCredentialsTaskInternalServerError{}
}

/*
RetryCredentialsTaskInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type RetryCredentialsTaskInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this retry credentials task internal server error response has a 2xx status code
func (o *RetryCredentialsTaskInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this retry credentials task internal server error response has a 3xx status code
func (o *RetryCredentialsTaskInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this retry credentials task internal server error response has a 4xx status code
func (o *RetryCredentialsTaskInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this retry credentials task internal server error response has a 5xx status code
func (o *RetryCredentialsTaskInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this retry credentials task internal server error response a status code equal to that given
func (o *RetryCredentialsTaskInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *RetryCredentialsTaskInternalServerError) Error() string {
	return fmt.Sprintf("[PATCH /v1/credentials/tasks/{id}][%d] retryCredentialsTaskInternalServerError  %+v", 500, o.Payload)
}

func (o *RetryCredentialsTaskInternalServerError) String() string {
	return fmt.Sprintf("[PATCH /v1/credentials/tasks/{id}][%d] retryCredentialsTaskInternalServerError  %+v", 500, o.Payload)
}

func (o *RetryCredentialsTaskInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *RetryCredentialsTaskInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
