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

// GetUsersReader is a Reader for the GetUsers structure.
type GetUsersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetUsersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetUsersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetUsersBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetUsersUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetUsersInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /v1/users] getUsers", response, response.Code())
	}
}

// NewGetUsersOK creates a GetUsersOK with default headers values
func NewGetUsersOK() *GetUsersOK {
	return &GetUsersOK{}
}

/*
GetUsersOK describes a response with status code 200, with default header values.

OK
*/
type GetUsersOK struct {
	Payload *models.PageOfUser
}

// IsSuccess returns true when this get users o k response has a 2xx status code
func (o *GetUsersOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get users o k response has a 3xx status code
func (o *GetUsersOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get users o k response has a 4xx status code
func (o *GetUsersOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get users o k response has a 5xx status code
func (o *GetUsersOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get users o k response a status code equal to that given
func (o *GetUsersOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get users o k response
func (o *GetUsersOK) Code() int {
	return 200
}

func (o *GetUsersOK) Error() string {
	return fmt.Sprintf("[GET /v1/users][%d] getUsersOK  %+v", 200, o.Payload)
}

func (o *GetUsersOK) String() string {
	return fmt.Sprintf("[GET /v1/users][%d] getUsersOK  %+v", 200, o.Payload)
}

func (o *GetUsersOK) GetPayload() *models.PageOfUser {
	return o.Payload
}

func (o *GetUsersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PageOfUser)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUsersBadRequest creates a GetUsersBadRequest with default headers values
func NewGetUsersBadRequest() *GetUsersBadRequest {
	return &GetUsersBadRequest{}
}

/*
GetUsersBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type GetUsersBadRequest struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this get users bad request response has a 2xx status code
func (o *GetUsersBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get users bad request response has a 3xx status code
func (o *GetUsersBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get users bad request response has a 4xx status code
func (o *GetUsersBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get users bad request response has a 5xx status code
func (o *GetUsersBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get users bad request response a status code equal to that given
func (o *GetUsersBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get users bad request response
func (o *GetUsersBadRequest) Code() int {
	return 400
}

func (o *GetUsersBadRequest) Error() string {
	return fmt.Sprintf("[GET /v1/users][%d] getUsersBadRequest  %+v", 400, o.Payload)
}

func (o *GetUsersBadRequest) String() string {
	return fmt.Sprintf("[GET /v1/users][%d] getUsersBadRequest  %+v", 400, o.Payload)
}

func (o *GetUsersBadRequest) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetUsersBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUsersUnauthorized creates a GetUsersUnauthorized with default headers values
func NewGetUsersUnauthorized() *GetUsersUnauthorized {
	return &GetUsersUnauthorized{}
}

/*
GetUsersUnauthorized describes a response with status code 401, with default header values.

Unauthorized Request
*/
type GetUsersUnauthorized struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this get users unauthorized response has a 2xx status code
func (o *GetUsersUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get users unauthorized response has a 3xx status code
func (o *GetUsersUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get users unauthorized response has a 4xx status code
func (o *GetUsersUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get users unauthorized response has a 5xx status code
func (o *GetUsersUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get users unauthorized response a status code equal to that given
func (o *GetUsersUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get users unauthorized response
func (o *GetUsersUnauthorized) Code() int {
	return 401
}

func (o *GetUsersUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v1/users][%d] getUsersUnauthorized  %+v", 401, o.Payload)
}

func (o *GetUsersUnauthorized) String() string {
	return fmt.Sprintf("[GET /v1/users][%d] getUsersUnauthorized  %+v", 401, o.Payload)
}

func (o *GetUsersUnauthorized) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetUsersUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUsersInternalServerError creates a GetUsersInternalServerError with default headers values
func NewGetUsersInternalServerError() *GetUsersInternalServerError {
	return &GetUsersInternalServerError{}
}

/*
GetUsersInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetUsersInternalServerError struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this get users internal server error response has a 2xx status code
func (o *GetUsersInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get users internal server error response has a 3xx status code
func (o *GetUsersInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get users internal server error response has a 4xx status code
func (o *GetUsersInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get users internal server error response has a 5xx status code
func (o *GetUsersInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get users internal server error response a status code equal to that given
func (o *GetUsersInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get users internal server error response
func (o *GetUsersInternalServerError) Code() int {
	return 500
}

func (o *GetUsersInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/users][%d] getUsersInternalServerError  %+v", 500, o.Payload)
}

func (o *GetUsersInternalServerError) String() string {
	return fmt.Sprintf("[GET /v1/users][%d] getUsersInternalServerError  %+v", 500, o.Payload)
}

func (o *GetUsersInternalServerError) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetUsersInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
