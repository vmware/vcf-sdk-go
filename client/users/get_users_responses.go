// Code generated by go-swagger; DO NOT EDIT.

// Copyright © 2023 VMware, Inc. All Rights Reserved.
// Code licenced under: BSD-2

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"vcf-sdk-go/models"
)

// GETUsersReader is a Reader for the GETUsers structure.
type GETUsersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GETUsersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGETUsersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGETUsersBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGETUsersUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGETUsersInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGETUsersOK creates a GETUsersOK with default headers values
func NewGETUsersOK() *GETUsersOK {
	return &GETUsersOK{}
}

/*
GETUsersOK describes a response with status code 200, with default header values.

OK
*/
type GETUsersOK struct {
	Payload *models.PageOfUser
}

// IsSuccess returns true when this get users o k response has a 2xx status code
func (o *GETUsersOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get users o k response has a 3xx status code
func (o *GETUsersOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get users o k response has a 4xx status code
func (o *GETUsersOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get users o k response has a 5xx status code
func (o *GETUsersOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get users o k response a status code equal to that given
func (o *GETUsersOK) IsCode(code int) bool {
	return code == 200
}

func (o *GETUsersOK) Error() string {
	return fmt.Sprintf("[GET /v1/users][%d] getUsersOK  %+v", 200, o.Payload)
}

func (o *GETUsersOK) String() string {
	return fmt.Sprintf("[GET /v1/users][%d] getUsersOK  %+v", 200, o.Payload)
}

func (o *GETUsersOK) GetPayload() *models.PageOfUser {
	return o.Payload
}

func (o *GETUsersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PageOfUser)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGETUsersBadRequest creates a GETUsersBadRequest with default headers values
func NewGETUsersBadRequest() *GETUsersBadRequest {
	return &GETUsersBadRequest{}
}

/*
GETUsersBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type GETUsersBadRequest struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this get users bad request response has a 2xx status code
func (o *GETUsersBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get users bad request response has a 3xx status code
func (o *GETUsersBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get users bad request response has a 4xx status code
func (o *GETUsersBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get users bad request response has a 5xx status code
func (o *GETUsersBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get users bad request response a status code equal to that given
func (o *GETUsersBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GETUsersBadRequest) Error() string {
	return fmt.Sprintf("[GET /v1/users][%d] getUsersBadRequest  %+v", 400, o.Payload)
}

func (o *GETUsersBadRequest) String() string {
	return fmt.Sprintf("[GET /v1/users][%d] getUsersBadRequest  %+v", 400, o.Payload)
}

func (o *GETUsersBadRequest) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GETUsersBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGETUsersUnauthorized creates a GETUsersUnauthorized with default headers values
func NewGETUsersUnauthorized() *GETUsersUnauthorized {
	return &GETUsersUnauthorized{}
}

/*
GETUsersUnauthorized describes a response with status code 401, with default header values.

Unauthorized Request
*/
type GETUsersUnauthorized struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this get users unauthorized response has a 2xx status code
func (o *GETUsersUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get users unauthorized response has a 3xx status code
func (o *GETUsersUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get users unauthorized response has a 4xx status code
func (o *GETUsersUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get users unauthorized response has a 5xx status code
func (o *GETUsersUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get users unauthorized response a status code equal to that given
func (o *GETUsersUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GETUsersUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v1/users][%d] getUsersUnauthorized  %+v", 401, o.Payload)
}

func (o *GETUsersUnauthorized) String() string {
	return fmt.Sprintf("[GET /v1/users][%d] getUsersUnauthorized  %+v", 401, o.Payload)
}

func (o *GETUsersUnauthorized) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GETUsersUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGETUsersInternalServerError creates a GETUsersInternalServerError with default headers values
func NewGETUsersInternalServerError() *GETUsersInternalServerError {
	return &GETUsersInternalServerError{}
}

/*
GETUsersInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GETUsersInternalServerError struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this get users internal server error response has a 2xx status code
func (o *GETUsersInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get users internal server error response has a 3xx status code
func (o *GETUsersInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get users internal server error response has a 4xx status code
func (o *GETUsersInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get users internal server error response has a 5xx status code
func (o *GETUsersInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get users internal server error response a status code equal to that given
func (o *GETUsersInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GETUsersInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/users][%d] getUsersInternalServerError  %+v", 500, o.Payload)
}

func (o *GETUsersInternalServerError) String() string {
	return fmt.Sprintf("[GET /v1/users][%d] getUsersInternalServerError  %+v", 500, o.Payload)
}

func (o *GETUsersInternalServerError) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GETUsersInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
