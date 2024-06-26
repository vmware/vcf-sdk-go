// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2023 VMware, Inc.
// SPDX-License-Identifier: BSD-2-Clause

package system_prechecks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/vmware/vcf-sdk-go/models"
)

// StartPrecheckReader is a Reader for the StartPrecheck structure.
type StartPrecheckReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StartPrecheckReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStartPrecheckOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewStartPrecheckAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewStartPrecheckBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewStartPrecheckInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /v1/system/prechecks] startPrecheck", response, response.Code())
	}
}

// NewStartPrecheckOK creates a StartPrecheckOK with default headers values
func NewStartPrecheckOK() *StartPrecheckOK {
	return &StartPrecheckOK{}
}

/*
StartPrecheckOK describes a response with status code 200, with default header values.

OK
*/
type StartPrecheckOK struct {
	Payload *models.Task
}

// IsSuccess returns true when this start precheck o k response has a 2xx status code
func (o *StartPrecheckOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this start precheck o k response has a 3xx status code
func (o *StartPrecheckOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this start precheck o k response has a 4xx status code
func (o *StartPrecheckOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this start precheck o k response has a 5xx status code
func (o *StartPrecheckOK) IsServerError() bool {
	return false
}

// IsCode returns true when this start precheck o k response a status code equal to that given
func (o *StartPrecheckOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the start precheck o k response
func (o *StartPrecheckOK) Code() int {
	return 200
}

func (o *StartPrecheckOK) Error() string {
	return fmt.Sprintf("[POST /v1/system/prechecks][%d] startPrecheckOK  %+v", 200, o.Payload)
}

func (o *StartPrecheckOK) String() string {
	return fmt.Sprintf("[POST /v1/system/prechecks][%d] startPrecheckOK  %+v", 200, o.Payload)
}

func (o *StartPrecheckOK) GetPayload() *models.Task {
	return o.Payload
}

func (o *StartPrecheckOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Task)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStartPrecheckAccepted creates a StartPrecheckAccepted with default headers values
func NewStartPrecheckAccepted() *StartPrecheckAccepted {
	return &StartPrecheckAccepted{}
}

/*
StartPrecheckAccepted describes a response with status code 202, with default header values.

Accepted
*/
type StartPrecheckAccepted struct {
	Payload *models.Task
}

// IsSuccess returns true when this start precheck accepted response has a 2xx status code
func (o *StartPrecheckAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this start precheck accepted response has a 3xx status code
func (o *StartPrecheckAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this start precheck accepted response has a 4xx status code
func (o *StartPrecheckAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this start precheck accepted response has a 5xx status code
func (o *StartPrecheckAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this start precheck accepted response a status code equal to that given
func (o *StartPrecheckAccepted) IsCode(code int) bool {
	return code == 202
}

// Code gets the status code for the start precheck accepted response
func (o *StartPrecheckAccepted) Code() int {
	return 202
}

func (o *StartPrecheckAccepted) Error() string {
	return fmt.Sprintf("[POST /v1/system/prechecks][%d] startPrecheckAccepted  %+v", 202, o.Payload)
}

func (o *StartPrecheckAccepted) String() string {
	return fmt.Sprintf("[POST /v1/system/prechecks][%d] startPrecheckAccepted  %+v", 202, o.Payload)
}

func (o *StartPrecheckAccepted) GetPayload() *models.Task {
	return o.Payload
}

func (o *StartPrecheckAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Task)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStartPrecheckBadRequest creates a StartPrecheckBadRequest with default headers values
func NewStartPrecheckBadRequest() *StartPrecheckBadRequest {
	return &StartPrecheckBadRequest{}
}

/*
StartPrecheckBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type StartPrecheckBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this start precheck bad request response has a 2xx status code
func (o *StartPrecheckBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this start precheck bad request response has a 3xx status code
func (o *StartPrecheckBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this start precheck bad request response has a 4xx status code
func (o *StartPrecheckBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this start precheck bad request response has a 5xx status code
func (o *StartPrecheckBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this start precheck bad request response a status code equal to that given
func (o *StartPrecheckBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the start precheck bad request response
func (o *StartPrecheckBadRequest) Code() int {
	return 400
}

func (o *StartPrecheckBadRequest) Error() string {
	return fmt.Sprintf("[POST /v1/system/prechecks][%d] startPrecheckBadRequest  %+v", 400, o.Payload)
}

func (o *StartPrecheckBadRequest) String() string {
	return fmt.Sprintf("[POST /v1/system/prechecks][%d] startPrecheckBadRequest  %+v", 400, o.Payload)
}

func (o *StartPrecheckBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *StartPrecheckBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStartPrecheckInternalServerError creates a StartPrecheckInternalServerError with default headers values
func NewStartPrecheckInternalServerError() *StartPrecheckInternalServerError {
	return &StartPrecheckInternalServerError{}
}

/*
StartPrecheckInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type StartPrecheckInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this start precheck internal server error response has a 2xx status code
func (o *StartPrecheckInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this start precheck internal server error response has a 3xx status code
func (o *StartPrecheckInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this start precheck internal server error response has a 4xx status code
func (o *StartPrecheckInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this start precheck internal server error response has a 5xx status code
func (o *StartPrecheckInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this start precheck internal server error response a status code equal to that given
func (o *StartPrecheckInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the start precheck internal server error response
func (o *StartPrecheckInternalServerError) Code() int {
	return 500
}

func (o *StartPrecheckInternalServerError) Error() string {
	return fmt.Sprintf("[POST /v1/system/prechecks][%d] startPrecheckInternalServerError  %+v", 500, o.Payload)
}

func (o *StartPrecheckInternalServerError) String() string {
	return fmt.Sprintf("[POST /v1/system/prechecks][%d] startPrecheckInternalServerError  %+v", 500, o.Payload)
}

func (o *StartPrecheckInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *StartPrecheckInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
