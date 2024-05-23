// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2023 VMware, Inc.
// SPDX-License-Identifier: BSD-2-Clause

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/vmware/vcf-sdk-go/models"
)

// UpdateLocalUserPasswordReader is a Reader for the UpdateLocalUserPassword structure.
type UpdateLocalUserPasswordReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateLocalUserPasswordReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewUpdateLocalUserPasswordNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewUpdateLocalUserPasswordInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[PATCH /v1/users/local/admin] updateLocalUserPassword", response, response.Code())
	}
}

// NewUpdateLocalUserPasswordNoContent creates a UpdateLocalUserPasswordNoContent with default headers values
func NewUpdateLocalUserPasswordNoContent() *UpdateLocalUserPasswordNoContent {
	return &UpdateLocalUserPasswordNoContent{}
}

/*
UpdateLocalUserPasswordNoContent describes a response with status code 204, with default header values.

No content
*/
type UpdateLocalUserPasswordNoContent struct {
}

// IsSuccess returns true when this update local user password no content response has a 2xx status code
func (o *UpdateLocalUserPasswordNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update local user password no content response has a 3xx status code
func (o *UpdateLocalUserPasswordNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update local user password no content response has a 4xx status code
func (o *UpdateLocalUserPasswordNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this update local user password no content response has a 5xx status code
func (o *UpdateLocalUserPasswordNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this update local user password no content response a status code equal to that given
func (o *UpdateLocalUserPasswordNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the update local user password no content response
func (o *UpdateLocalUserPasswordNoContent) Code() int {
	return 204
}

func (o *UpdateLocalUserPasswordNoContent) Error() string {
	return fmt.Sprintf("[PATCH /v1/users/local/admin][%d] updateLocalUserPasswordNoContent ", 204)
}

func (o *UpdateLocalUserPasswordNoContent) String() string {
	return fmt.Sprintf("[PATCH /v1/users/local/admin][%d] updateLocalUserPasswordNoContent ", 204)
}

func (o *UpdateLocalUserPasswordNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateLocalUserPasswordInternalServerError creates a UpdateLocalUserPasswordInternalServerError with default headers values
func NewUpdateLocalUserPasswordInternalServerError() *UpdateLocalUserPasswordInternalServerError {
	return &UpdateLocalUserPasswordInternalServerError{}
}

/*
UpdateLocalUserPasswordInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type UpdateLocalUserPasswordInternalServerError struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this update local user password internal server error response has a 2xx status code
func (o *UpdateLocalUserPasswordInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update local user password internal server error response has a 3xx status code
func (o *UpdateLocalUserPasswordInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update local user password internal server error response has a 4xx status code
func (o *UpdateLocalUserPasswordInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this update local user password internal server error response has a 5xx status code
func (o *UpdateLocalUserPasswordInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this update local user password internal server error response a status code equal to that given
func (o *UpdateLocalUserPasswordInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the update local user password internal server error response
func (o *UpdateLocalUserPasswordInternalServerError) Code() int {
	return 500
}

func (o *UpdateLocalUserPasswordInternalServerError) Error() string {
	return fmt.Sprintf("[PATCH /v1/users/local/admin][%d] updateLocalUserPasswordInternalServerError  %+v", 500, o.Payload)
}

func (o *UpdateLocalUserPasswordInternalServerError) String() string {
	return fmt.Sprintf("[PATCH /v1/users/local/admin][%d] updateLocalUserPasswordInternalServerError  %+v", 500, o.Payload)
}

func (o *UpdateLocalUserPasswordInternalServerError) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *UpdateLocalUserPasswordInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
