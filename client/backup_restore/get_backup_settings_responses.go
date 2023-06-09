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

// GetBackupSettingsReader is a Reader for the GetBackupSettings structure.
type GetBackupSettingsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetBackupSettingsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetBackupSettingsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetBackupSettingsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetBackupSettingsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetBackupSettingsOK creates a GetBackupSettingsOK with default headers values
func NewGetBackupSettingsOK() *GetBackupSettingsOK {
	return &GetBackupSettingsOK{}
}

/*
GetBackupSettingsOK describes a response with status code 200, with default header values.

Ok
*/
type GetBackupSettingsOK struct {
	Payload *models.BackupConfiguration
}

// IsSuccess returns true when this get backup settings o k response has a 2xx status code
func (o *GetBackupSettingsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get backup settings o k response has a 3xx status code
func (o *GetBackupSettingsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get backup settings o k response has a 4xx status code
func (o *GetBackupSettingsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get backup settings o k response has a 5xx status code
func (o *GetBackupSettingsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get backup settings o k response a status code equal to that given
func (o *GetBackupSettingsOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetBackupSettingsOK) Error() string {
	return fmt.Sprintf("[GET /v1/system/backup-configuration][%d] getBackupSettingsOK  %+v", 200, o.Payload)
}

func (o *GetBackupSettingsOK) String() string {
	return fmt.Sprintf("[GET /v1/system/backup-configuration][%d] getBackupSettingsOK  %+v", 200, o.Payload)
}

func (o *GetBackupSettingsOK) GetPayload() *models.BackupConfiguration {
	return o.Payload
}

func (o *GetBackupSettingsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BackupConfiguration)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetBackupSettingsBadRequest creates a GetBackupSettingsBadRequest with default headers values
func NewGetBackupSettingsBadRequest() *GetBackupSettingsBadRequest {
	return &GetBackupSettingsBadRequest{}
}

/*
GetBackupSettingsBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetBackupSettingsBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this get backup settings bad request response has a 2xx status code
func (o *GetBackupSettingsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get backup settings bad request response has a 3xx status code
func (o *GetBackupSettingsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get backup settings bad request response has a 4xx status code
func (o *GetBackupSettingsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get backup settings bad request response has a 5xx status code
func (o *GetBackupSettingsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get backup settings bad request response a status code equal to that given
func (o *GetBackupSettingsBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetBackupSettingsBadRequest) Error() string {
	return fmt.Sprintf("[GET /v1/system/backup-configuration][%d] getBackupSettingsBadRequest  %+v", 400, o.Payload)
}

func (o *GetBackupSettingsBadRequest) String() string {
	return fmt.Sprintf("[GET /v1/system/backup-configuration][%d] getBackupSettingsBadRequest  %+v", 400, o.Payload)
}

func (o *GetBackupSettingsBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetBackupSettingsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetBackupSettingsInternalServerError creates a GetBackupSettingsInternalServerError with default headers values
func NewGetBackupSettingsInternalServerError() *GetBackupSettingsInternalServerError {
	return &GetBackupSettingsInternalServerError{}
}

/*
GetBackupSettingsInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type GetBackupSettingsInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this get backup settings internal server error response has a 2xx status code
func (o *GetBackupSettingsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get backup settings internal server error response has a 3xx status code
func (o *GetBackupSettingsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get backup settings internal server error response has a 4xx status code
func (o *GetBackupSettingsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get backup settings internal server error response has a 5xx status code
func (o *GetBackupSettingsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get backup settings internal server error response a status code equal to that given
func (o *GetBackupSettingsInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetBackupSettingsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/system/backup-configuration][%d] getBackupSettingsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetBackupSettingsInternalServerError) String() string {
	return fmt.Sprintf("[GET /v1/system/backup-configuration][%d] getBackupSettingsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetBackupSettingsInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetBackupSettingsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
