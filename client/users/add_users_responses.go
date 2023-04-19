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

// AddUsersReader is a Reader for the AddUsers structure.
type AddUsersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddUsersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAddUsersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 201:
		result := NewAddUsersCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewAddUsersBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewAddUsersUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewAddUsersInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewAddUsersOK creates a AddUsersOK with default headers values
func NewAddUsersOK() *AddUsersOK {
	return &AddUsersOK{}
}

/*
AddUsersOK describes a response with status code 200, with default header values.

OK
*/
type AddUsersOK struct {
	Payload *models.PageOfUser
}

// IsSuccess returns true when this add users o k response has a 2xx status code
func (o *AddUsersOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this add users o k response has a 3xx status code
func (o *AddUsersOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this add users o k response has a 4xx status code
func (o *AddUsersOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this add users o k response has a 5xx status code
func (o *AddUsersOK) IsServerError() bool {
	return false
}

// IsCode returns true when this add users o k response a status code equal to that given
func (o *AddUsersOK) IsCode(code int) bool {
	return code == 200
}

func (o *AddUsersOK) Error() string {
	return fmt.Sprintf("[POST /v1/users][%d] addUsersOK  %+v", 200, o.Payload)
}

func (o *AddUsersOK) String() string {
	return fmt.Sprintf("[POST /v1/users][%d] addUsersOK  %+v", 200, o.Payload)
}

func (o *AddUsersOK) GetPayload() *models.PageOfUser {
	return o.Payload
}

func (o *AddUsersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PageOfUser)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddUsersCreated creates a AddUsersCreated with default headers values
func NewAddUsersCreated() *AddUsersCreated {
	return &AddUsersCreated{}
}

/*
AddUsersCreated describes a response with status code 201, with default header values.

Created
*/
type AddUsersCreated struct {
	Payload *models.PageOfUser
}

// IsSuccess returns true when this add users created response has a 2xx status code
func (o *AddUsersCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this add users created response has a 3xx status code
func (o *AddUsersCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this add users created response has a 4xx status code
func (o *AddUsersCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this add users created response has a 5xx status code
func (o *AddUsersCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this add users created response a status code equal to that given
func (o *AddUsersCreated) IsCode(code int) bool {
	return code == 201
}

func (o *AddUsersCreated) Error() string {
	return fmt.Sprintf("[POST /v1/users][%d] addUsersCreated  %+v", 201, o.Payload)
}

func (o *AddUsersCreated) String() string {
	return fmt.Sprintf("[POST /v1/users][%d] addUsersCreated  %+v", 201, o.Payload)
}

func (o *AddUsersCreated) GetPayload() *models.PageOfUser {
	return o.Payload
}

func (o *AddUsersCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PageOfUser)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddUsersBadRequest creates a AddUsersBadRequest with default headers values
func NewAddUsersBadRequest() *AddUsersBadRequest {
	return &AddUsersBadRequest{}
}

/*
AddUsersBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type AddUsersBadRequest struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this add users bad request response has a 2xx status code
func (o *AddUsersBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this add users bad request response has a 3xx status code
func (o *AddUsersBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this add users bad request response has a 4xx status code
func (o *AddUsersBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this add users bad request response has a 5xx status code
func (o *AddUsersBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this add users bad request response a status code equal to that given
func (o *AddUsersBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *AddUsersBadRequest) Error() string {
	return fmt.Sprintf("[POST /v1/users][%d] addUsersBadRequest  %+v", 400, o.Payload)
}

func (o *AddUsersBadRequest) String() string {
	return fmt.Sprintf("[POST /v1/users][%d] addUsersBadRequest  %+v", 400, o.Payload)
}

func (o *AddUsersBadRequest) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *AddUsersBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddUsersUnauthorized creates a AddUsersUnauthorized with default headers values
func NewAddUsersUnauthorized() *AddUsersUnauthorized {
	return &AddUsersUnauthorized{}
}

/*
AddUsersUnauthorized describes a response with status code 401, with default header values.

Unauthorized Request
*/
type AddUsersUnauthorized struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this add users unauthorized response has a 2xx status code
func (o *AddUsersUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this add users unauthorized response has a 3xx status code
func (o *AddUsersUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this add users unauthorized response has a 4xx status code
func (o *AddUsersUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this add users unauthorized response has a 5xx status code
func (o *AddUsersUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this add users unauthorized response a status code equal to that given
func (o *AddUsersUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *AddUsersUnauthorized) Error() string {
	return fmt.Sprintf("[POST /v1/users][%d] addUsersUnauthorized  %+v", 401, o.Payload)
}

func (o *AddUsersUnauthorized) String() string {
	return fmt.Sprintf("[POST /v1/users][%d] addUsersUnauthorized  %+v", 401, o.Payload)
}

func (o *AddUsersUnauthorized) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *AddUsersUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddUsersInternalServerError creates a AddUsersInternalServerError with default headers values
func NewAddUsersInternalServerError() *AddUsersInternalServerError {
	return &AddUsersInternalServerError{}
}

/*
AddUsersInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type AddUsersInternalServerError struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this add users internal server error response has a 2xx status code
func (o *AddUsersInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this add users internal server error response has a 3xx status code
func (o *AddUsersInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this add users internal server error response has a 4xx status code
func (o *AddUsersInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this add users internal server error response has a 5xx status code
func (o *AddUsersInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this add users internal server error response a status code equal to that given
func (o *AddUsersInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *AddUsersInternalServerError) Error() string {
	return fmt.Sprintf("[POST /v1/users][%d] addUsersInternalServerError  %+v", 500, o.Payload)
}

func (o *AddUsersInternalServerError) String() string {
	return fmt.Sprintf("[POST /v1/users][%d] addUsersInternalServerError  %+v", 500, o.Payload)
}

func (o *AddUsersInternalServerError) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *AddUsersInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
