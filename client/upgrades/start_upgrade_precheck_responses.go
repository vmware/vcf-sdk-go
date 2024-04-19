// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2023 VMware, Inc.
// SPDX-License-Identifier: BSD-2-Clause

package upgrades

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/vmware/vcf-sdk-go/models"
)

// StartUpgradePrecheckReader is a Reader for the StartUpgradePrecheck structure.
type StartUpgradePrecheckReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StartUpgradePrecheckReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStartUpgradePrecheckOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewStartUpgradePrecheckAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewStartUpgradePrecheckBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewStartUpgradePrecheckForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewStartUpgradePrecheckInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /v1/upgrades/{upgradeId}/prechecks] startUpgradePrecheck", response, response.Code())
	}
}

// NewStartUpgradePrecheckOK creates a StartUpgradePrecheckOK with default headers values
func NewStartUpgradePrecheckOK() *StartUpgradePrecheckOK {
	return &StartUpgradePrecheckOK{}
}

/*
StartUpgradePrecheckOK describes a response with status code 200, with default header values.

OK
*/
type StartUpgradePrecheckOK struct {
	Payload *models.Task
}

// IsSuccess returns true when this start upgrade precheck o k response has a 2xx status code
func (o *StartUpgradePrecheckOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this start upgrade precheck o k response has a 3xx status code
func (o *StartUpgradePrecheckOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this start upgrade precheck o k response has a 4xx status code
func (o *StartUpgradePrecheckOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this start upgrade precheck o k response has a 5xx status code
func (o *StartUpgradePrecheckOK) IsServerError() bool {
	return false
}

// IsCode returns true when this start upgrade precheck o k response a status code equal to that given
func (o *StartUpgradePrecheckOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the start upgrade precheck o k response
func (o *StartUpgradePrecheckOK) Code() int {
	return 200
}

func (o *StartUpgradePrecheckOK) Error() string {
	return fmt.Sprintf("[POST /v1/upgrades/{upgradeId}/prechecks][%d] startUpgradePrecheckOK  %+v", 200, o.Payload)
}

func (o *StartUpgradePrecheckOK) String() string {
	return fmt.Sprintf("[POST /v1/upgrades/{upgradeId}/prechecks][%d] startUpgradePrecheckOK  %+v", 200, o.Payload)
}

func (o *StartUpgradePrecheckOK) GetPayload() *models.Task {
	return o.Payload
}

func (o *StartUpgradePrecheckOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Task)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStartUpgradePrecheckAccepted creates a StartUpgradePrecheckAccepted with default headers values
func NewStartUpgradePrecheckAccepted() *StartUpgradePrecheckAccepted {
	return &StartUpgradePrecheckAccepted{}
}

/*
StartUpgradePrecheckAccepted describes a response with status code 202, with default header values.

Accepted
*/
type StartUpgradePrecheckAccepted struct {
	Payload *models.Task
}

// IsSuccess returns true when this start upgrade precheck accepted response has a 2xx status code
func (o *StartUpgradePrecheckAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this start upgrade precheck accepted response has a 3xx status code
func (o *StartUpgradePrecheckAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this start upgrade precheck accepted response has a 4xx status code
func (o *StartUpgradePrecheckAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this start upgrade precheck accepted response has a 5xx status code
func (o *StartUpgradePrecheckAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this start upgrade precheck accepted response a status code equal to that given
func (o *StartUpgradePrecheckAccepted) IsCode(code int) bool {
	return code == 202
}

// Code gets the status code for the start upgrade precheck accepted response
func (o *StartUpgradePrecheckAccepted) Code() int {
	return 202
}

func (o *StartUpgradePrecheckAccepted) Error() string {
	return fmt.Sprintf("[POST /v1/upgrades/{upgradeId}/prechecks][%d] startUpgradePrecheckAccepted  %+v", 202, o.Payload)
}

func (o *StartUpgradePrecheckAccepted) String() string {
	return fmt.Sprintf("[POST /v1/upgrades/{upgradeId}/prechecks][%d] startUpgradePrecheckAccepted  %+v", 202, o.Payload)
}

func (o *StartUpgradePrecheckAccepted) GetPayload() *models.Task {
	return o.Payload
}

func (o *StartUpgradePrecheckAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Task)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStartUpgradePrecheckBadRequest creates a StartUpgradePrecheckBadRequest with default headers values
func NewStartUpgradePrecheckBadRequest() *StartUpgradePrecheckBadRequest {
	return &StartUpgradePrecheckBadRequest{}
}

/*
StartUpgradePrecheckBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type StartUpgradePrecheckBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this start upgrade precheck bad request response has a 2xx status code
func (o *StartUpgradePrecheckBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this start upgrade precheck bad request response has a 3xx status code
func (o *StartUpgradePrecheckBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this start upgrade precheck bad request response has a 4xx status code
func (o *StartUpgradePrecheckBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this start upgrade precheck bad request response has a 5xx status code
func (o *StartUpgradePrecheckBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this start upgrade precheck bad request response a status code equal to that given
func (o *StartUpgradePrecheckBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the start upgrade precheck bad request response
func (o *StartUpgradePrecheckBadRequest) Code() int {
	return 400
}

func (o *StartUpgradePrecheckBadRequest) Error() string {
	return fmt.Sprintf("[POST /v1/upgrades/{upgradeId}/prechecks][%d] startUpgradePrecheckBadRequest  %+v", 400, o.Payload)
}

func (o *StartUpgradePrecheckBadRequest) String() string {
	return fmt.Sprintf("[POST /v1/upgrades/{upgradeId}/prechecks][%d] startUpgradePrecheckBadRequest  %+v", 400, o.Payload)
}

func (o *StartUpgradePrecheckBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *StartUpgradePrecheckBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStartUpgradePrecheckForbidden creates a StartUpgradePrecheckForbidden with default headers values
func NewStartUpgradePrecheckForbidden() *StartUpgradePrecheckForbidden {
	return &StartUpgradePrecheckForbidden{}
}

/*
StartUpgradePrecheckForbidden describes a response with status code 403, with default header values.

Operation not allowed
*/
type StartUpgradePrecheckForbidden struct {
	Payload *models.Error
}

// IsSuccess returns true when this start upgrade precheck forbidden response has a 2xx status code
func (o *StartUpgradePrecheckForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this start upgrade precheck forbidden response has a 3xx status code
func (o *StartUpgradePrecheckForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this start upgrade precheck forbidden response has a 4xx status code
func (o *StartUpgradePrecheckForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this start upgrade precheck forbidden response has a 5xx status code
func (o *StartUpgradePrecheckForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this start upgrade precheck forbidden response a status code equal to that given
func (o *StartUpgradePrecheckForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the start upgrade precheck forbidden response
func (o *StartUpgradePrecheckForbidden) Code() int {
	return 403
}

func (o *StartUpgradePrecheckForbidden) Error() string {
	return fmt.Sprintf("[POST /v1/upgrades/{upgradeId}/prechecks][%d] startUpgradePrecheckForbidden  %+v", 403, o.Payload)
}

func (o *StartUpgradePrecheckForbidden) String() string {
	return fmt.Sprintf("[POST /v1/upgrades/{upgradeId}/prechecks][%d] startUpgradePrecheckForbidden  %+v", 403, o.Payload)
}

func (o *StartUpgradePrecheckForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *StartUpgradePrecheckForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStartUpgradePrecheckInternalServerError creates a StartUpgradePrecheckInternalServerError with default headers values
func NewStartUpgradePrecheckInternalServerError() *StartUpgradePrecheckInternalServerError {
	return &StartUpgradePrecheckInternalServerError{}
}

/*
StartUpgradePrecheckInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type StartUpgradePrecheckInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this start upgrade precheck internal server error response has a 2xx status code
func (o *StartUpgradePrecheckInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this start upgrade precheck internal server error response has a 3xx status code
func (o *StartUpgradePrecheckInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this start upgrade precheck internal server error response has a 4xx status code
func (o *StartUpgradePrecheckInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this start upgrade precheck internal server error response has a 5xx status code
func (o *StartUpgradePrecheckInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this start upgrade precheck internal server error response a status code equal to that given
func (o *StartUpgradePrecheckInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the start upgrade precheck internal server error response
func (o *StartUpgradePrecheckInternalServerError) Code() int {
	return 500
}

func (o *StartUpgradePrecheckInternalServerError) Error() string {
	return fmt.Sprintf("[POST /v1/upgrades/{upgradeId}/prechecks][%d] startUpgradePrecheckInternalServerError  %+v", 500, o.Payload)
}

func (o *StartUpgradePrecheckInternalServerError) String() string {
	return fmt.Sprintf("[POST /v1/upgrades/{upgradeId}/prechecks][%d] startUpgradePrecheckInternalServerError  %+v", 500, o.Payload)
}

func (o *StartUpgradePrecheckInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *StartUpgradePrecheckInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}