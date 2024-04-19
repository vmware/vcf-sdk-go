// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2023 VMware, Inc.
// SPDX-License-Identifier: BSD-2-Clause

package domains

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/vmware/vcf-sdk-go/models"
)

// CreateDomainReader is a Reader for the CreateDomain structure.
type CreateDomainReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateDomainReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateDomainOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewCreateDomainAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateDomainBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCreateDomainInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /v1/domains] createDomain", response, response.Code())
	}
}

// NewCreateDomainOK creates a CreateDomainOK with default headers values
func NewCreateDomainOK() *CreateDomainOK {
	return &CreateDomainOK{}
}

/*
CreateDomainOK describes a response with status code 200, with default header values.

OK
*/
type CreateDomainOK struct {
	Payload *models.Task
}

// IsSuccess returns true when this create domain o k response has a 2xx status code
func (o *CreateDomainOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create domain o k response has a 3xx status code
func (o *CreateDomainOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create domain o k response has a 4xx status code
func (o *CreateDomainOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this create domain o k response has a 5xx status code
func (o *CreateDomainOK) IsServerError() bool {
	return false
}

// IsCode returns true when this create domain o k response a status code equal to that given
func (o *CreateDomainOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the create domain o k response
func (o *CreateDomainOK) Code() int {
	return 200
}

func (o *CreateDomainOK) Error() string {
	return fmt.Sprintf("[POST /v1/domains][%d] createDomainOK  %+v", 200, o.Payload)
}

func (o *CreateDomainOK) String() string {
	return fmt.Sprintf("[POST /v1/domains][%d] createDomainOK  %+v", 200, o.Payload)
}

func (o *CreateDomainOK) GetPayload() *models.Task {
	return o.Payload
}

func (o *CreateDomainOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Task)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateDomainAccepted creates a CreateDomainAccepted with default headers values
func NewCreateDomainAccepted() *CreateDomainAccepted {
	return &CreateDomainAccepted{}
}

/*
CreateDomainAccepted describes a response with status code 202, with default header values.

Accepted
*/
type CreateDomainAccepted struct {
	Payload *models.Task
}

// IsSuccess returns true when this create domain accepted response has a 2xx status code
func (o *CreateDomainAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create domain accepted response has a 3xx status code
func (o *CreateDomainAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create domain accepted response has a 4xx status code
func (o *CreateDomainAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this create domain accepted response has a 5xx status code
func (o *CreateDomainAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this create domain accepted response a status code equal to that given
func (o *CreateDomainAccepted) IsCode(code int) bool {
	return code == 202
}

// Code gets the status code for the create domain accepted response
func (o *CreateDomainAccepted) Code() int {
	return 202
}

func (o *CreateDomainAccepted) Error() string {
	return fmt.Sprintf("[POST /v1/domains][%d] createDomainAccepted  %+v", 202, o.Payload)
}

func (o *CreateDomainAccepted) String() string {
	return fmt.Sprintf("[POST /v1/domains][%d] createDomainAccepted  %+v", 202, o.Payload)
}

func (o *CreateDomainAccepted) GetPayload() *models.Task {
	return o.Payload
}

func (o *CreateDomainAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Task)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateDomainBadRequest creates a CreateDomainBadRequest with default headers values
func NewCreateDomainBadRequest() *CreateDomainBadRequest {
	return &CreateDomainBadRequest{}
}

/*
CreateDomainBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type CreateDomainBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this create domain bad request response has a 2xx status code
func (o *CreateDomainBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create domain bad request response has a 3xx status code
func (o *CreateDomainBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create domain bad request response has a 4xx status code
func (o *CreateDomainBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this create domain bad request response has a 5xx status code
func (o *CreateDomainBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this create domain bad request response a status code equal to that given
func (o *CreateDomainBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the create domain bad request response
func (o *CreateDomainBadRequest) Code() int {
	return 400
}

func (o *CreateDomainBadRequest) Error() string {
	return fmt.Sprintf("[POST /v1/domains][%d] createDomainBadRequest  %+v", 400, o.Payload)
}

func (o *CreateDomainBadRequest) String() string {
	return fmt.Sprintf("[POST /v1/domains][%d] createDomainBadRequest  %+v", 400, o.Payload)
}

func (o *CreateDomainBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateDomainBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateDomainInternalServerError creates a CreateDomainInternalServerError with default headers values
func NewCreateDomainInternalServerError() *CreateDomainInternalServerError {
	return &CreateDomainInternalServerError{}
}

/*
CreateDomainInternalServerError describes a response with status code 500, with default header values.

InternalServerError
*/
type CreateDomainInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this create domain internal server error response has a 2xx status code
func (o *CreateDomainInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create domain internal server error response has a 3xx status code
func (o *CreateDomainInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create domain internal server error response has a 4xx status code
func (o *CreateDomainInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this create domain internal server error response has a 5xx status code
func (o *CreateDomainInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this create domain internal server error response a status code equal to that given
func (o *CreateDomainInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the create domain internal server error response
func (o *CreateDomainInternalServerError) Code() int {
	return 500
}

func (o *CreateDomainInternalServerError) Error() string {
	return fmt.Sprintf("[POST /v1/domains][%d] createDomainInternalServerError  %+v", 500, o.Payload)
}

func (o *CreateDomainInternalServerError) String() string {
	return fmt.Sprintf("[POST /v1/domains][%d] createDomainInternalServerError  %+v", 500, o.Payload)
}

func (o *CreateDomainInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateDomainInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
