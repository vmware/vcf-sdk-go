// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2023 VMware, Inc.
// SPDX-License-Identifier: BSD-2-Clause

package sddc

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/vmware/vcf-sdk-go/models"
)

// StartBringupReader is a Reader for the StartBringup structure.
type StartBringupReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StartBringupReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStartBringupOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewStartBringupAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewStartBringupBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewStartBringupInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /v1/sddcs] startBringup", response, response.Code())
	}
}

// NewStartBringupOK creates a StartBringupOK with default headers values
func NewStartBringupOK() *StartBringupOK {
	return &StartBringupOK{}
}

/*
StartBringupOK describes a response with status code 200, with default header values.

OK
*/
type StartBringupOK struct {
	Payload *models.SDDCTask
}

// IsSuccess returns true when this start bringup o k response has a 2xx status code
func (o *StartBringupOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this start bringup o k response has a 3xx status code
func (o *StartBringupOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this start bringup o k response has a 4xx status code
func (o *StartBringupOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this start bringup o k response has a 5xx status code
func (o *StartBringupOK) IsServerError() bool {
	return false
}

// IsCode returns true when this start bringup o k response a status code equal to that given
func (o *StartBringupOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the start bringup o k response
func (o *StartBringupOK) Code() int {
	return 200
}

func (o *StartBringupOK) Error() string {
	return fmt.Sprintf("[POST /v1/sddcs][%d] startBringupOK  %+v", 200, o.Payload)
}

func (o *StartBringupOK) String() string {
	return fmt.Sprintf("[POST /v1/sddcs][%d] startBringupOK  %+v", 200, o.Payload)
}

func (o *StartBringupOK) GetPayload() *models.SDDCTask {
	return o.Payload
}

func (o *StartBringupOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SDDCTask)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStartBringupAccepted creates a StartBringupAccepted with default headers values
func NewStartBringupAccepted() *StartBringupAccepted {
	return &StartBringupAccepted{}
}

/*
StartBringupAccepted describes a response with status code 202, with default header values.

Success
*/
type StartBringupAccepted struct {
	Payload *models.SDDCTask
}

// IsSuccess returns true when this start bringup accepted response has a 2xx status code
func (o *StartBringupAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this start bringup accepted response has a 3xx status code
func (o *StartBringupAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this start bringup accepted response has a 4xx status code
func (o *StartBringupAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this start bringup accepted response has a 5xx status code
func (o *StartBringupAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this start bringup accepted response a status code equal to that given
func (o *StartBringupAccepted) IsCode(code int) bool {
	return code == 202
}

// Code gets the status code for the start bringup accepted response
func (o *StartBringupAccepted) Code() int {
	return 202
}

func (o *StartBringupAccepted) Error() string {
	return fmt.Sprintf("[POST /v1/sddcs][%d] startBringupAccepted  %+v", 202, o.Payload)
}

func (o *StartBringupAccepted) String() string {
	return fmt.Sprintf("[POST /v1/sddcs][%d] startBringupAccepted  %+v", 202, o.Payload)
}

func (o *StartBringupAccepted) GetPayload() *models.SDDCTask {
	return o.Payload
}

func (o *StartBringupAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SDDCTask)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStartBringupBadRequest creates a StartBringupBadRequest with default headers values
func NewStartBringupBadRequest() *StartBringupBadRequest {
	return &StartBringupBadRequest{}
}

/*
StartBringupBadRequest describes a response with status code 400, with default header values.

SDDC already exists, Bad Request
*/
type StartBringupBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this start bringup bad request response has a 2xx status code
func (o *StartBringupBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this start bringup bad request response has a 3xx status code
func (o *StartBringupBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this start bringup bad request response has a 4xx status code
func (o *StartBringupBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this start bringup bad request response has a 5xx status code
func (o *StartBringupBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this start bringup bad request response a status code equal to that given
func (o *StartBringupBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the start bringup bad request response
func (o *StartBringupBadRequest) Code() int {
	return 400
}

func (o *StartBringupBadRequest) Error() string {
	return fmt.Sprintf("[POST /v1/sddcs][%d] startBringupBadRequest  %+v", 400, o.Payload)
}

func (o *StartBringupBadRequest) String() string {
	return fmt.Sprintf("[POST /v1/sddcs][%d] startBringupBadRequest  %+v", 400, o.Payload)
}

func (o *StartBringupBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *StartBringupBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStartBringupInternalServerError creates a StartBringupInternalServerError with default headers values
func NewStartBringupInternalServerError() *StartBringupInternalServerError {
	return &StartBringupInternalServerError{}
}

/*
StartBringupInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type StartBringupInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this start bringup internal server error response has a 2xx status code
func (o *StartBringupInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this start bringup internal server error response has a 3xx status code
func (o *StartBringupInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this start bringup internal server error response has a 4xx status code
func (o *StartBringupInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this start bringup internal server error response has a 5xx status code
func (o *StartBringupInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this start bringup internal server error response a status code equal to that given
func (o *StartBringupInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the start bringup internal server error response
func (o *StartBringupInternalServerError) Code() int {
	return 500
}

func (o *StartBringupInternalServerError) Error() string {
	return fmt.Sprintf("[POST /v1/sddcs][%d] startBringupInternalServerError  %+v", 500, o.Payload)
}

func (o *StartBringupInternalServerError) String() string {
	return fmt.Sprintf("[POST /v1/sddcs][%d] startBringupInternalServerError  %+v", 500, o.Payload)
}

func (o *StartBringupInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *StartBringupInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}