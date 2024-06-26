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

// GetRolesReader is a Reader for the GetRoles structure.
type GetRolesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRolesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetRolesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetRolesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetRolesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetRolesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /v1/roles] getRoles", response, response.Code())
	}
}

// NewGetRolesOK creates a GetRolesOK with default headers values
func NewGetRolesOK() *GetRolesOK {
	return &GetRolesOK{}
}

/*
GetRolesOK describes a response with status code 200, with default header values.

OK
*/
type GetRolesOK struct {
	Payload *models.PageOfRole
}

// IsSuccess returns true when this get roles o k response has a 2xx status code
func (o *GetRolesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get roles o k response has a 3xx status code
func (o *GetRolesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get roles o k response has a 4xx status code
func (o *GetRolesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get roles o k response has a 5xx status code
func (o *GetRolesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get roles o k response a status code equal to that given
func (o *GetRolesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get roles o k response
func (o *GetRolesOK) Code() int {
	return 200
}

func (o *GetRolesOK) Error() string {
	return fmt.Sprintf("[GET /v1/roles][%d] getRolesOK  %+v", 200, o.Payload)
}

func (o *GetRolesOK) String() string {
	return fmt.Sprintf("[GET /v1/roles][%d] getRolesOK  %+v", 200, o.Payload)
}

func (o *GetRolesOK) GetPayload() *models.PageOfRole {
	return o.Payload
}

func (o *GetRolesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PageOfRole)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRolesBadRequest creates a GetRolesBadRequest with default headers values
func NewGetRolesBadRequest() *GetRolesBadRequest {
	return &GetRolesBadRequest{}
}

/*
GetRolesBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type GetRolesBadRequest struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this get roles bad request response has a 2xx status code
func (o *GetRolesBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get roles bad request response has a 3xx status code
func (o *GetRolesBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get roles bad request response has a 4xx status code
func (o *GetRolesBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get roles bad request response has a 5xx status code
func (o *GetRolesBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get roles bad request response a status code equal to that given
func (o *GetRolesBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get roles bad request response
func (o *GetRolesBadRequest) Code() int {
	return 400
}

func (o *GetRolesBadRequest) Error() string {
	return fmt.Sprintf("[GET /v1/roles][%d] getRolesBadRequest  %+v", 400, o.Payload)
}

func (o *GetRolesBadRequest) String() string {
	return fmt.Sprintf("[GET /v1/roles][%d] getRolesBadRequest  %+v", 400, o.Payload)
}

func (o *GetRolesBadRequest) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetRolesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRolesUnauthorized creates a GetRolesUnauthorized with default headers values
func NewGetRolesUnauthorized() *GetRolesUnauthorized {
	return &GetRolesUnauthorized{}
}

/*
GetRolesUnauthorized describes a response with status code 401, with default header values.

Unauthorized Request
*/
type GetRolesUnauthorized struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this get roles unauthorized response has a 2xx status code
func (o *GetRolesUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get roles unauthorized response has a 3xx status code
func (o *GetRolesUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get roles unauthorized response has a 4xx status code
func (o *GetRolesUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get roles unauthorized response has a 5xx status code
func (o *GetRolesUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get roles unauthorized response a status code equal to that given
func (o *GetRolesUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get roles unauthorized response
func (o *GetRolesUnauthorized) Code() int {
	return 401
}

func (o *GetRolesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v1/roles][%d] getRolesUnauthorized  %+v", 401, o.Payload)
}

func (o *GetRolesUnauthorized) String() string {
	return fmt.Sprintf("[GET /v1/roles][%d] getRolesUnauthorized  %+v", 401, o.Payload)
}

func (o *GetRolesUnauthorized) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetRolesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRolesInternalServerError creates a GetRolesInternalServerError with default headers values
func NewGetRolesInternalServerError() *GetRolesInternalServerError {
	return &GetRolesInternalServerError{}
}

/*
GetRolesInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetRolesInternalServerError struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this get roles internal server error response has a 2xx status code
func (o *GetRolesInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get roles internal server error response has a 3xx status code
func (o *GetRolesInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get roles internal server error response has a 4xx status code
func (o *GetRolesInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get roles internal server error response has a 5xx status code
func (o *GetRolesInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get roles internal server error response a status code equal to that given
func (o *GetRolesInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get roles internal server error response
func (o *GetRolesInternalServerError) Code() int {
	return 500
}

func (o *GetRolesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/roles][%d] getRolesInternalServerError  %+v", 500, o.Payload)
}

func (o *GetRolesInternalServerError) String() string {
	return fmt.Sprintf("[GET /v1/roles][%d] getRolesInternalServerError  %+v", 500, o.Payload)
}

func (o *GetRolesInternalServerError) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetRolesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
