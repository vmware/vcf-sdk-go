// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2023 VMware, Inc.
// SPDX-License-Identifier: BSD-2-Clause

package backup_restore

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/vmware/vcf-sdk-go/models"
)

// SetBackupConfigurationReader is a Reader for the SetBackupConfiguration structure.
type SetBackupConfigurationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SetBackupConfigurationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSetBackupConfigurationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewSetBackupConfigurationAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewSetBackupConfigurationBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewSetBackupConfigurationInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[PUT /v1/system/backup-configuration] setBackupConfiguration", response, response.Code())
	}
}

// NewSetBackupConfigurationOK creates a SetBackupConfigurationOK with default headers values
func NewSetBackupConfigurationOK() *SetBackupConfigurationOK {
	return &SetBackupConfigurationOK{}
}

/*
SetBackupConfigurationOK describes a response with status code 200, with default header values.

OK
*/
type SetBackupConfigurationOK struct {
	Payload *models.Task
}

// IsSuccess returns true when this set backup configuration o k response has a 2xx status code
func (o *SetBackupConfigurationOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this set backup configuration o k response has a 3xx status code
func (o *SetBackupConfigurationOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this set backup configuration o k response has a 4xx status code
func (o *SetBackupConfigurationOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this set backup configuration o k response has a 5xx status code
func (o *SetBackupConfigurationOK) IsServerError() bool {
	return false
}

// IsCode returns true when this set backup configuration o k response a status code equal to that given
func (o *SetBackupConfigurationOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the set backup configuration o k response
func (o *SetBackupConfigurationOK) Code() int {
	return 200
}

func (o *SetBackupConfigurationOK) Error() string {
	return fmt.Sprintf("[PUT /v1/system/backup-configuration][%d] setBackupConfigurationOK  %+v", 200, o.Payload)
}

func (o *SetBackupConfigurationOK) String() string {
	return fmt.Sprintf("[PUT /v1/system/backup-configuration][%d] setBackupConfigurationOK  %+v", 200, o.Payload)
}

func (o *SetBackupConfigurationOK) GetPayload() *models.Task {
	return o.Payload
}

func (o *SetBackupConfigurationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Task)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSetBackupConfigurationAccepted creates a SetBackupConfigurationAccepted with default headers values
func NewSetBackupConfigurationAccepted() *SetBackupConfigurationAccepted {
	return &SetBackupConfigurationAccepted{}
}

/*
SetBackupConfigurationAccepted describes a response with status code 202, with default header values.

Accepted
*/
type SetBackupConfigurationAccepted struct {
	Payload *models.Task
}

// IsSuccess returns true when this set backup configuration accepted response has a 2xx status code
func (o *SetBackupConfigurationAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this set backup configuration accepted response has a 3xx status code
func (o *SetBackupConfigurationAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this set backup configuration accepted response has a 4xx status code
func (o *SetBackupConfigurationAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this set backup configuration accepted response has a 5xx status code
func (o *SetBackupConfigurationAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this set backup configuration accepted response a status code equal to that given
func (o *SetBackupConfigurationAccepted) IsCode(code int) bool {
	return code == 202
}

// Code gets the status code for the set backup configuration accepted response
func (o *SetBackupConfigurationAccepted) Code() int {
	return 202
}

func (o *SetBackupConfigurationAccepted) Error() string {
	return fmt.Sprintf("[PUT /v1/system/backup-configuration][%d] setBackupConfigurationAccepted  %+v", 202, o.Payload)
}

func (o *SetBackupConfigurationAccepted) String() string {
	return fmt.Sprintf("[PUT /v1/system/backup-configuration][%d] setBackupConfigurationAccepted  %+v", 202, o.Payload)
}

func (o *SetBackupConfigurationAccepted) GetPayload() *models.Task {
	return o.Payload
}

func (o *SetBackupConfigurationAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Task)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSetBackupConfigurationBadRequest creates a SetBackupConfigurationBadRequest with default headers values
func NewSetBackupConfigurationBadRequest() *SetBackupConfigurationBadRequest {
	return &SetBackupConfigurationBadRequest{}
}

/*
SetBackupConfigurationBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type SetBackupConfigurationBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this set backup configuration bad request response has a 2xx status code
func (o *SetBackupConfigurationBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this set backup configuration bad request response has a 3xx status code
func (o *SetBackupConfigurationBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this set backup configuration bad request response has a 4xx status code
func (o *SetBackupConfigurationBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this set backup configuration bad request response has a 5xx status code
func (o *SetBackupConfigurationBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this set backup configuration bad request response a status code equal to that given
func (o *SetBackupConfigurationBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the set backup configuration bad request response
func (o *SetBackupConfigurationBadRequest) Code() int {
	return 400
}

func (o *SetBackupConfigurationBadRequest) Error() string {
	return fmt.Sprintf("[PUT /v1/system/backup-configuration][%d] setBackupConfigurationBadRequest  %+v", 400, o.Payload)
}

func (o *SetBackupConfigurationBadRequest) String() string {
	return fmt.Sprintf("[PUT /v1/system/backup-configuration][%d] setBackupConfigurationBadRequest  %+v", 400, o.Payload)
}

func (o *SetBackupConfigurationBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *SetBackupConfigurationBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSetBackupConfigurationInternalServerError creates a SetBackupConfigurationInternalServerError with default headers values
func NewSetBackupConfigurationInternalServerError() *SetBackupConfigurationInternalServerError {
	return &SetBackupConfigurationInternalServerError{}
}

/*
SetBackupConfigurationInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type SetBackupConfigurationInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this set backup configuration internal server error response has a 2xx status code
func (o *SetBackupConfigurationInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this set backup configuration internal server error response has a 3xx status code
func (o *SetBackupConfigurationInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this set backup configuration internal server error response has a 4xx status code
func (o *SetBackupConfigurationInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this set backup configuration internal server error response has a 5xx status code
func (o *SetBackupConfigurationInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this set backup configuration internal server error response a status code equal to that given
func (o *SetBackupConfigurationInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the set backup configuration internal server error response
func (o *SetBackupConfigurationInternalServerError) Code() int {
	return 500
}

func (o *SetBackupConfigurationInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /v1/system/backup-configuration][%d] setBackupConfigurationInternalServerError  %+v", 500, o.Payload)
}

func (o *SetBackupConfigurationInternalServerError) String() string {
	return fmt.Sprintf("[PUT /v1/system/backup-configuration][%d] setBackupConfigurationInternalServerError  %+v", 500, o.Payload)
}

func (o *SetBackupConfigurationInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *SetBackupConfigurationInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
