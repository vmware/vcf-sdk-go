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

// ConfigureBackupSettingsReader is a Reader for the ConfigureBackupSettings structure.
type ConfigureBackupSettingsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ConfigureBackupSettingsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewConfigureBackupSettingsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewConfigureBackupSettingsAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewConfigureBackupSettingsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewConfigureBackupSettingsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewConfigureBackupSettingsOK creates a ConfigureBackupSettingsOK with default headers values
func NewConfigureBackupSettingsOK() *ConfigureBackupSettingsOK {
	return &ConfigureBackupSettingsOK{}
}

/*
ConfigureBackupSettingsOK describes a response with status code 200, with default header values.

OK
*/
type ConfigureBackupSettingsOK struct {
	Payload *models.Task
}

// IsSuccess returns true when this configure backup settings o k response has a 2xx status code
func (o *ConfigureBackupSettingsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this configure backup settings o k response has a 3xx status code
func (o *ConfigureBackupSettingsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this configure backup settings o k response has a 4xx status code
func (o *ConfigureBackupSettingsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this configure backup settings o k response has a 5xx status code
func (o *ConfigureBackupSettingsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this configure backup settings o k response a status code equal to that given
func (o *ConfigureBackupSettingsOK) IsCode(code int) bool {
	return code == 200
}

func (o *ConfigureBackupSettingsOK) Error() string {
	return fmt.Sprintf("[PUT /v1/system/backup-configuration][%d] configureBackupSettingsOK  %+v", 200, o.Payload)
}

func (o *ConfigureBackupSettingsOK) String() string {
	return fmt.Sprintf("[PUT /v1/system/backup-configuration][%d] configureBackupSettingsOK  %+v", 200, o.Payload)
}

func (o *ConfigureBackupSettingsOK) GetPayload() *models.Task {
	return o.Payload
}

func (o *ConfigureBackupSettingsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Task)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewConfigureBackupSettingsAccepted creates a ConfigureBackupSettingsAccepted with default headers values
func NewConfigureBackupSettingsAccepted() *ConfigureBackupSettingsAccepted {
	return &ConfigureBackupSettingsAccepted{}
}

/*
ConfigureBackupSettingsAccepted describes a response with status code 202, with default header values.

Accepted
*/
type ConfigureBackupSettingsAccepted struct {
	Payload *models.Task
}

// IsSuccess returns true when this configure backup settings accepted response has a 2xx status code
func (o *ConfigureBackupSettingsAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this configure backup settings accepted response has a 3xx status code
func (o *ConfigureBackupSettingsAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this configure backup settings accepted response has a 4xx status code
func (o *ConfigureBackupSettingsAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this configure backup settings accepted response has a 5xx status code
func (o *ConfigureBackupSettingsAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this configure backup settings accepted response a status code equal to that given
func (o *ConfigureBackupSettingsAccepted) IsCode(code int) bool {
	return code == 202
}

func (o *ConfigureBackupSettingsAccepted) Error() string {
	return fmt.Sprintf("[PUT /v1/system/backup-configuration][%d] configureBackupSettingsAccepted  %+v", 202, o.Payload)
}

func (o *ConfigureBackupSettingsAccepted) String() string {
	return fmt.Sprintf("[PUT /v1/system/backup-configuration][%d] configureBackupSettingsAccepted  %+v", 202, o.Payload)
}

func (o *ConfigureBackupSettingsAccepted) GetPayload() *models.Task {
	return o.Payload
}

func (o *ConfigureBackupSettingsAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Task)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewConfigureBackupSettingsBadRequest creates a ConfigureBackupSettingsBadRequest with default headers values
func NewConfigureBackupSettingsBadRequest() *ConfigureBackupSettingsBadRequest {
	return &ConfigureBackupSettingsBadRequest{}
}

/*
ConfigureBackupSettingsBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type ConfigureBackupSettingsBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this configure backup settings bad request response has a 2xx status code
func (o *ConfigureBackupSettingsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this configure backup settings bad request response has a 3xx status code
func (o *ConfigureBackupSettingsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this configure backup settings bad request response has a 4xx status code
func (o *ConfigureBackupSettingsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this configure backup settings bad request response has a 5xx status code
func (o *ConfigureBackupSettingsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this configure backup settings bad request response a status code equal to that given
func (o *ConfigureBackupSettingsBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *ConfigureBackupSettingsBadRequest) Error() string {
	return fmt.Sprintf("[PUT /v1/system/backup-configuration][%d] configureBackupSettingsBadRequest  %+v", 400, o.Payload)
}

func (o *ConfigureBackupSettingsBadRequest) String() string {
	return fmt.Sprintf("[PUT /v1/system/backup-configuration][%d] configureBackupSettingsBadRequest  %+v", 400, o.Payload)
}

func (o *ConfigureBackupSettingsBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *ConfigureBackupSettingsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewConfigureBackupSettingsInternalServerError creates a ConfigureBackupSettingsInternalServerError with default headers values
func NewConfigureBackupSettingsInternalServerError() *ConfigureBackupSettingsInternalServerError {
	return &ConfigureBackupSettingsInternalServerError{}
}

/*
ConfigureBackupSettingsInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type ConfigureBackupSettingsInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this configure backup settings internal server error response has a 2xx status code
func (o *ConfigureBackupSettingsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this configure backup settings internal server error response has a 3xx status code
func (o *ConfigureBackupSettingsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this configure backup settings internal server error response has a 4xx status code
func (o *ConfigureBackupSettingsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this configure backup settings internal server error response has a 5xx status code
func (o *ConfigureBackupSettingsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this configure backup settings internal server error response a status code equal to that given
func (o *ConfigureBackupSettingsInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *ConfigureBackupSettingsInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /v1/system/backup-configuration][%d] configureBackupSettingsInternalServerError  %+v", 500, o.Payload)
}

func (o *ConfigureBackupSettingsInternalServerError) String() string {
	return fmt.Sprintf("[PUT /v1/system/backup-configuration][%d] configureBackupSettingsInternalServerError  %+v", 500, o.Payload)
}

func (o *ConfigureBackupSettingsInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *ConfigureBackupSettingsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
