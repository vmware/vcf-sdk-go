// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2023 VMware, Inc.
// SPDX-License-Identifier: BSD-2-Clause

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

// UpdateOrRotatePasswordsReader is a Reader for the UpdateOrRotatePasswords structure.
type UpdateOrRotatePasswordsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateOrRotatePasswordsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateOrRotatePasswordsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewUpdateOrRotatePasswordsAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateOrRotatePasswordsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewUpdateOrRotatePasswordsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateOrRotatePasswordsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUpdateOrRotatePasswordsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateOrRotatePasswordsOK creates a UpdateOrRotatePasswordsOK with default headers values
func NewUpdateOrRotatePasswordsOK() *UpdateOrRotatePasswordsOK {
	return &UpdateOrRotatePasswordsOK{}
}

/*
UpdateOrRotatePasswordsOK describes a response with status code 200, with default header values.

OK
*/
type UpdateOrRotatePasswordsOK struct {
	Payload *models.Task
}

// IsSuccess returns true when this update or rotate passwords o k response has a 2xx status code
func (o *UpdateOrRotatePasswordsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update or rotate passwords o k response has a 3xx status code
func (o *UpdateOrRotatePasswordsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update or rotate passwords o k response has a 4xx status code
func (o *UpdateOrRotatePasswordsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update or rotate passwords o k response has a 5xx status code
func (o *UpdateOrRotatePasswordsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update or rotate passwords o k response a status code equal to that given
func (o *UpdateOrRotatePasswordsOK) IsCode(code int) bool {
	return code == 200
}

func (o *UpdateOrRotatePasswordsOK) Error() string {
	return fmt.Sprintf("[PATCH /v1/credentials][%d] updateOrRotatePasswordsOK  %+v", 200, o.Payload)
}

func (o *UpdateOrRotatePasswordsOK) String() string {
	return fmt.Sprintf("[PATCH /v1/credentials][%d] updateOrRotatePasswordsOK  %+v", 200, o.Payload)
}

func (o *UpdateOrRotatePasswordsOK) GetPayload() *models.Task {
	return o.Payload
}

func (o *UpdateOrRotatePasswordsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Task)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateOrRotatePasswordsAccepted creates a UpdateOrRotatePasswordsAccepted with default headers values
func NewUpdateOrRotatePasswordsAccepted() *UpdateOrRotatePasswordsAccepted {
	return &UpdateOrRotatePasswordsAccepted{}
}

/*
UpdateOrRotatePasswordsAccepted describes a response with status code 202, with default header values.

Accepted
*/
type UpdateOrRotatePasswordsAccepted struct {
	Payload *models.Task
}

// IsSuccess returns true when this update or rotate passwords accepted response has a 2xx status code
func (o *UpdateOrRotatePasswordsAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update or rotate passwords accepted response has a 3xx status code
func (o *UpdateOrRotatePasswordsAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update or rotate passwords accepted response has a 4xx status code
func (o *UpdateOrRotatePasswordsAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this update or rotate passwords accepted response has a 5xx status code
func (o *UpdateOrRotatePasswordsAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this update or rotate passwords accepted response a status code equal to that given
func (o *UpdateOrRotatePasswordsAccepted) IsCode(code int) bool {
	return code == 202
}

func (o *UpdateOrRotatePasswordsAccepted) Error() string {
	return fmt.Sprintf("[PATCH /v1/credentials][%d] updateOrRotatePasswordsAccepted  %+v", 202, o.Payload)
}

func (o *UpdateOrRotatePasswordsAccepted) String() string {
	return fmt.Sprintf("[PATCH /v1/credentials][%d] updateOrRotatePasswordsAccepted  %+v", 202, o.Payload)
}

func (o *UpdateOrRotatePasswordsAccepted) GetPayload() *models.Task {
	return o.Payload
}

func (o *UpdateOrRotatePasswordsAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Task)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateOrRotatePasswordsBadRequest creates a UpdateOrRotatePasswordsBadRequest with default headers values
func NewUpdateOrRotatePasswordsBadRequest() *UpdateOrRotatePasswordsBadRequest {
	return &UpdateOrRotatePasswordsBadRequest{}
}

/*
UpdateOrRotatePasswordsBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type UpdateOrRotatePasswordsBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this update or rotate passwords bad request response has a 2xx status code
func (o *UpdateOrRotatePasswordsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update or rotate passwords bad request response has a 3xx status code
func (o *UpdateOrRotatePasswordsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update or rotate passwords bad request response has a 4xx status code
func (o *UpdateOrRotatePasswordsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this update or rotate passwords bad request response has a 5xx status code
func (o *UpdateOrRotatePasswordsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this update or rotate passwords bad request response a status code equal to that given
func (o *UpdateOrRotatePasswordsBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *UpdateOrRotatePasswordsBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /v1/credentials][%d] updateOrRotatePasswordsBadRequest  %+v", 400, o.Payload)
}

func (o *UpdateOrRotatePasswordsBadRequest) String() string {
	return fmt.Sprintf("[PATCH /v1/credentials][%d] updateOrRotatePasswordsBadRequest  %+v", 400, o.Payload)
}

func (o *UpdateOrRotatePasswordsBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateOrRotatePasswordsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateOrRotatePasswordsUnauthorized creates a UpdateOrRotatePasswordsUnauthorized with default headers values
func NewUpdateOrRotatePasswordsUnauthorized() *UpdateOrRotatePasswordsUnauthorized {
	return &UpdateOrRotatePasswordsUnauthorized{}
}

/*
UpdateOrRotatePasswordsUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type UpdateOrRotatePasswordsUnauthorized struct {
	Payload *models.Error
}

// IsSuccess returns true when this update or rotate passwords unauthorized response has a 2xx status code
func (o *UpdateOrRotatePasswordsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update or rotate passwords unauthorized response has a 3xx status code
func (o *UpdateOrRotatePasswordsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update or rotate passwords unauthorized response has a 4xx status code
func (o *UpdateOrRotatePasswordsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this update or rotate passwords unauthorized response has a 5xx status code
func (o *UpdateOrRotatePasswordsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this update or rotate passwords unauthorized response a status code equal to that given
func (o *UpdateOrRotatePasswordsUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *UpdateOrRotatePasswordsUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /v1/credentials][%d] updateOrRotatePasswordsUnauthorized  %+v", 401, o.Payload)
}

func (o *UpdateOrRotatePasswordsUnauthorized) String() string {
	return fmt.Sprintf("[PATCH /v1/credentials][%d] updateOrRotatePasswordsUnauthorized  %+v", 401, o.Payload)
}

func (o *UpdateOrRotatePasswordsUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateOrRotatePasswordsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateOrRotatePasswordsForbidden creates a UpdateOrRotatePasswordsForbidden with default headers values
func NewUpdateOrRotatePasswordsForbidden() *UpdateOrRotatePasswordsForbidden {
	return &UpdateOrRotatePasswordsForbidden{}
}

/*
UpdateOrRotatePasswordsForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type UpdateOrRotatePasswordsForbidden struct {
	Payload *models.Error
}

// IsSuccess returns true when this update or rotate passwords forbidden response has a 2xx status code
func (o *UpdateOrRotatePasswordsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update or rotate passwords forbidden response has a 3xx status code
func (o *UpdateOrRotatePasswordsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update or rotate passwords forbidden response has a 4xx status code
func (o *UpdateOrRotatePasswordsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this update or rotate passwords forbidden response has a 5xx status code
func (o *UpdateOrRotatePasswordsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this update or rotate passwords forbidden response a status code equal to that given
func (o *UpdateOrRotatePasswordsForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *UpdateOrRotatePasswordsForbidden) Error() string {
	return fmt.Sprintf("[PATCH /v1/credentials][%d] updateOrRotatePasswordsForbidden  %+v", 403, o.Payload)
}

func (o *UpdateOrRotatePasswordsForbidden) String() string {
	return fmt.Sprintf("[PATCH /v1/credentials][%d] updateOrRotatePasswordsForbidden  %+v", 403, o.Payload)
}

func (o *UpdateOrRotatePasswordsForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateOrRotatePasswordsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateOrRotatePasswordsInternalServerError creates a UpdateOrRotatePasswordsInternalServerError with default headers values
func NewUpdateOrRotatePasswordsInternalServerError() *UpdateOrRotatePasswordsInternalServerError {
	return &UpdateOrRotatePasswordsInternalServerError{}
}

/*
UpdateOrRotatePasswordsInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type UpdateOrRotatePasswordsInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this update or rotate passwords internal server error response has a 2xx status code
func (o *UpdateOrRotatePasswordsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update or rotate passwords internal server error response has a 3xx status code
func (o *UpdateOrRotatePasswordsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update or rotate passwords internal server error response has a 4xx status code
func (o *UpdateOrRotatePasswordsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this update or rotate passwords internal server error response has a 5xx status code
func (o *UpdateOrRotatePasswordsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this update or rotate passwords internal server error response a status code equal to that given
func (o *UpdateOrRotatePasswordsInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *UpdateOrRotatePasswordsInternalServerError) Error() string {
	return fmt.Sprintf("[PATCH /v1/credentials][%d] updateOrRotatePasswordsInternalServerError  %+v", 500, o.Payload)
}

func (o *UpdateOrRotatePasswordsInternalServerError) String() string {
	return fmt.Sprintf("[PATCH /v1/credentials][%d] updateOrRotatePasswordsInternalServerError  %+v", 500, o.Payload)
}

func (o *UpdateOrRotatePasswordsInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateOrRotatePasswordsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
