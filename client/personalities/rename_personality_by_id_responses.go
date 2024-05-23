// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2023 VMware, Inc.
// SPDX-License-Identifier: BSD-2-Clause

package personalities

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/vmware/vcf-sdk-go/models"
)

// RenamePersonalityByIDReader is a Reader for the RenamePersonalityByID structure.
type RenamePersonalityByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RenamePersonalityByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRenamePersonalityByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewRenamePersonalityByIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewRenamePersonalityByIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewRenamePersonalityByIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[PATCH /v1/personalities/{personalityId}] renamePersonalityById", response, response.Code())
	}
}

// NewRenamePersonalityByIDOK creates a RenamePersonalityByIDOK with default headers values
func NewRenamePersonalityByIDOK() *RenamePersonalityByIDOK {
	return &RenamePersonalityByIDOK{}
}

/*
RenamePersonalityByIDOK describes a response with status code 200, with default header values.

Ok
*/
type RenamePersonalityByIDOK struct {
	Payload *models.Personality
}

// IsSuccess returns true when this rename personality by Id o k response has a 2xx status code
func (o *RenamePersonalityByIDOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this rename personality by Id o k response has a 3xx status code
func (o *RenamePersonalityByIDOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this rename personality by Id o k response has a 4xx status code
func (o *RenamePersonalityByIDOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this rename personality by Id o k response has a 5xx status code
func (o *RenamePersonalityByIDOK) IsServerError() bool {
	return false
}

// IsCode returns true when this rename personality by Id o k response a status code equal to that given
func (o *RenamePersonalityByIDOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the rename personality by Id o k response
func (o *RenamePersonalityByIDOK) Code() int {
	return 200
}

func (o *RenamePersonalityByIDOK) Error() string {
	return fmt.Sprintf("[PATCH /v1/personalities/{personalityId}][%d] renamePersonalityByIdOK  %+v", 200, o.Payload)
}

func (o *RenamePersonalityByIDOK) String() string {
	return fmt.Sprintf("[PATCH /v1/personalities/{personalityId}][%d] renamePersonalityByIdOK  %+v", 200, o.Payload)
}

func (o *RenamePersonalityByIDOK) GetPayload() *models.Personality {
	return o.Payload
}

func (o *RenamePersonalityByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Personality)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRenamePersonalityByIDBadRequest creates a RenamePersonalityByIDBadRequest with default headers values
func NewRenamePersonalityByIDBadRequest() *RenamePersonalityByIDBadRequest {
	return &RenamePersonalityByIDBadRequest{}
}

/*
RenamePersonalityByIDBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type RenamePersonalityByIDBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this rename personality by Id bad request response has a 2xx status code
func (o *RenamePersonalityByIDBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this rename personality by Id bad request response has a 3xx status code
func (o *RenamePersonalityByIDBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this rename personality by Id bad request response has a 4xx status code
func (o *RenamePersonalityByIDBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this rename personality by Id bad request response has a 5xx status code
func (o *RenamePersonalityByIDBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this rename personality by Id bad request response a status code equal to that given
func (o *RenamePersonalityByIDBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the rename personality by Id bad request response
func (o *RenamePersonalityByIDBadRequest) Code() int {
	return 400
}

func (o *RenamePersonalityByIDBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /v1/personalities/{personalityId}][%d] renamePersonalityByIdBadRequest  %+v", 400, o.Payload)
}

func (o *RenamePersonalityByIDBadRequest) String() string {
	return fmt.Sprintf("[PATCH /v1/personalities/{personalityId}][%d] renamePersonalityByIdBadRequest  %+v", 400, o.Payload)
}

func (o *RenamePersonalityByIDBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *RenamePersonalityByIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRenamePersonalityByIDNotFound creates a RenamePersonalityByIDNotFound with default headers values
func NewRenamePersonalityByIDNotFound() *RenamePersonalityByIDNotFound {
	return &RenamePersonalityByIDNotFound{}
}

/*
RenamePersonalityByIDNotFound describes a response with status code 404, with default header values.

Not Found
*/
type RenamePersonalityByIDNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this rename personality by Id not found response has a 2xx status code
func (o *RenamePersonalityByIDNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this rename personality by Id not found response has a 3xx status code
func (o *RenamePersonalityByIDNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this rename personality by Id not found response has a 4xx status code
func (o *RenamePersonalityByIDNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this rename personality by Id not found response has a 5xx status code
func (o *RenamePersonalityByIDNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this rename personality by Id not found response a status code equal to that given
func (o *RenamePersonalityByIDNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the rename personality by Id not found response
func (o *RenamePersonalityByIDNotFound) Code() int {
	return 404
}

func (o *RenamePersonalityByIDNotFound) Error() string {
	return fmt.Sprintf("[PATCH /v1/personalities/{personalityId}][%d] renamePersonalityByIdNotFound  %+v", 404, o.Payload)
}

func (o *RenamePersonalityByIDNotFound) String() string {
	return fmt.Sprintf("[PATCH /v1/personalities/{personalityId}][%d] renamePersonalityByIdNotFound  %+v", 404, o.Payload)
}

func (o *RenamePersonalityByIDNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *RenamePersonalityByIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRenamePersonalityByIDInternalServerError creates a RenamePersonalityByIDInternalServerError with default headers values
func NewRenamePersonalityByIDInternalServerError() *RenamePersonalityByIDInternalServerError {
	return &RenamePersonalityByIDInternalServerError{}
}

/*
RenamePersonalityByIDInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type RenamePersonalityByIDInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this rename personality by Id internal server error response has a 2xx status code
func (o *RenamePersonalityByIDInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this rename personality by Id internal server error response has a 3xx status code
func (o *RenamePersonalityByIDInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this rename personality by Id internal server error response has a 4xx status code
func (o *RenamePersonalityByIDInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this rename personality by Id internal server error response has a 5xx status code
func (o *RenamePersonalityByIDInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this rename personality by Id internal server error response a status code equal to that given
func (o *RenamePersonalityByIDInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the rename personality by Id internal server error response
func (o *RenamePersonalityByIDInternalServerError) Code() int {
	return 500
}

func (o *RenamePersonalityByIDInternalServerError) Error() string {
	return fmt.Sprintf("[PATCH /v1/personalities/{personalityId}][%d] renamePersonalityByIdInternalServerError  %+v", 500, o.Payload)
}

func (o *RenamePersonalityByIDInternalServerError) String() string {
	return fmt.Sprintf("[PATCH /v1/personalities/{personalityId}][%d] renamePersonalityByIdInternalServerError  %+v", 500, o.Payload)
}

func (o *RenamePersonalityByIDInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *RenamePersonalityByIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
