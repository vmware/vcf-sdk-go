// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2023 VMware, Inc.
// SPDX-License-Identifier: BSD-2-Clause

package system

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/vmware/vcf-sdk-go/models"
)

// ConfigureNtpReader is a Reader for the ConfigureNtp structure.
type ConfigureNtpReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ConfigureNtpReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewConfigureNtpOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewConfigureNtpAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewConfigureNtpBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewConfigureNtpInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewConfigureNtpOK creates a ConfigureNtpOK with default headers values
func NewConfigureNtpOK() *ConfigureNtpOK {
	return &ConfigureNtpOK{}
}

/*
ConfigureNtpOK describes a response with status code 200, with default header values.

OK
*/
type ConfigureNtpOK struct {
	Payload *models.Task
}

// IsSuccess returns true when this configure ntp o k response has a 2xx status code
func (o *ConfigureNtpOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this configure ntp o k response has a 3xx status code
func (o *ConfigureNtpOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this configure ntp o k response has a 4xx status code
func (o *ConfigureNtpOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this configure ntp o k response has a 5xx status code
func (o *ConfigureNtpOK) IsServerError() bool {
	return false
}

// IsCode returns true when this configure ntp o k response a status code equal to that given
func (o *ConfigureNtpOK) IsCode(code int) bool {
	return code == 200
}

func (o *ConfigureNtpOK) Error() string {
	return fmt.Sprintf("[PUT /v1/system/ntp-configuration][%d] configureNtpOK  %+v", 200, o.Payload)
}

func (o *ConfigureNtpOK) String() string {
	return fmt.Sprintf("[PUT /v1/system/ntp-configuration][%d] configureNtpOK  %+v", 200, o.Payload)
}

func (o *ConfigureNtpOK) GetPayload() *models.Task {
	return o.Payload
}

func (o *ConfigureNtpOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Task)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewConfigureNtpAccepted creates a ConfigureNtpAccepted with default headers values
func NewConfigureNtpAccepted() *ConfigureNtpAccepted {
	return &ConfigureNtpAccepted{}
}

/*
ConfigureNtpAccepted describes a response with status code 202, with default header values.

Accepted
*/
type ConfigureNtpAccepted struct {
	Payload *models.Task
}

// IsSuccess returns true when this configure ntp accepted response has a 2xx status code
func (o *ConfigureNtpAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this configure ntp accepted response has a 3xx status code
func (o *ConfigureNtpAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this configure ntp accepted response has a 4xx status code
func (o *ConfigureNtpAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this configure ntp accepted response has a 5xx status code
func (o *ConfigureNtpAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this configure ntp accepted response a status code equal to that given
func (o *ConfigureNtpAccepted) IsCode(code int) bool {
	return code == 202
}

func (o *ConfigureNtpAccepted) Error() string {
	return fmt.Sprintf("[PUT /v1/system/ntp-configuration][%d] configureNtpAccepted  %+v", 202, o.Payload)
}

func (o *ConfigureNtpAccepted) String() string {
	return fmt.Sprintf("[PUT /v1/system/ntp-configuration][%d] configureNtpAccepted  %+v", 202, o.Payload)
}

func (o *ConfigureNtpAccepted) GetPayload() *models.Task {
	return o.Payload
}

func (o *ConfigureNtpAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Task)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewConfigureNtpBadRequest creates a ConfigureNtpBadRequest with default headers values
func NewConfigureNtpBadRequest() *ConfigureNtpBadRequest {
	return &ConfigureNtpBadRequest{}
}

/*
ConfigureNtpBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type ConfigureNtpBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this configure ntp bad request response has a 2xx status code
func (o *ConfigureNtpBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this configure ntp bad request response has a 3xx status code
func (o *ConfigureNtpBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this configure ntp bad request response has a 4xx status code
func (o *ConfigureNtpBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this configure ntp bad request response has a 5xx status code
func (o *ConfigureNtpBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this configure ntp bad request response a status code equal to that given
func (o *ConfigureNtpBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *ConfigureNtpBadRequest) Error() string {
	return fmt.Sprintf("[PUT /v1/system/ntp-configuration][%d] configureNtpBadRequest  %+v", 400, o.Payload)
}

func (o *ConfigureNtpBadRequest) String() string {
	return fmt.Sprintf("[PUT /v1/system/ntp-configuration][%d] configureNtpBadRequest  %+v", 400, o.Payload)
}

func (o *ConfigureNtpBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *ConfigureNtpBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewConfigureNtpInternalServerError creates a ConfigureNtpInternalServerError with default headers values
func NewConfigureNtpInternalServerError() *ConfigureNtpInternalServerError {
	return &ConfigureNtpInternalServerError{}
}

/*
ConfigureNtpInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type ConfigureNtpInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this configure ntp internal server error response has a 2xx status code
func (o *ConfigureNtpInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this configure ntp internal server error response has a 3xx status code
func (o *ConfigureNtpInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this configure ntp internal server error response has a 4xx status code
func (o *ConfigureNtpInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this configure ntp internal server error response has a 5xx status code
func (o *ConfigureNtpInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this configure ntp internal server error response a status code equal to that given
func (o *ConfigureNtpInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *ConfigureNtpInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /v1/system/ntp-configuration][%d] configureNtpInternalServerError  %+v", 500, o.Payload)
}

func (o *ConfigureNtpInternalServerError) String() string {
	return fmt.Sprintf("[PUT /v1/system/ntp-configuration][%d] configureNtpInternalServerError  %+v", 500, o.Payload)
}

func (o *ConfigureNtpInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *ConfigureNtpInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
