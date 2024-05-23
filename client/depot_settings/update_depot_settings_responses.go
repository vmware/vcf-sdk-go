// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2023 VMware, Inc.
// SPDX-License-Identifier: BSD-2-Clause

package depot_settings

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/vmware/vcf-sdk-go/models"
)

// UpdateDepotSettingsReader is a Reader for the UpdateDepotSettings structure.
type UpdateDepotSettingsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateDepotSettingsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateDepotSettingsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewUpdateDepotSettingsAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateDepotSettingsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUpdateDepotSettingsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[PUT /v1/system/settings/depot] updateDepotSettings", response, response.Code())
	}
}

// NewUpdateDepotSettingsOK creates a UpdateDepotSettingsOK with default headers values
func NewUpdateDepotSettingsOK() *UpdateDepotSettingsOK {
	return &UpdateDepotSettingsOK{}
}

/*
UpdateDepotSettingsOK describes a response with status code 200, with default header values.

OK
*/
type UpdateDepotSettingsOK struct {
	Payload *models.DepotSettings
}

// IsSuccess returns true when this update depot settings o k response has a 2xx status code
func (o *UpdateDepotSettingsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update depot settings o k response has a 3xx status code
func (o *UpdateDepotSettingsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update depot settings o k response has a 4xx status code
func (o *UpdateDepotSettingsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update depot settings o k response has a 5xx status code
func (o *UpdateDepotSettingsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update depot settings o k response a status code equal to that given
func (o *UpdateDepotSettingsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the update depot settings o k response
func (o *UpdateDepotSettingsOK) Code() int {
	return 200
}

func (o *UpdateDepotSettingsOK) Error() string {
	return fmt.Sprintf("[PUT /v1/system/settings/depot][%d] updateDepotSettingsOK  %+v", 200, o.Payload)
}

func (o *UpdateDepotSettingsOK) String() string {
	return fmt.Sprintf("[PUT /v1/system/settings/depot][%d] updateDepotSettingsOK  %+v", 200, o.Payload)
}

func (o *UpdateDepotSettingsOK) GetPayload() *models.DepotSettings {
	return o.Payload
}

func (o *UpdateDepotSettingsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DepotSettings)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateDepotSettingsAccepted creates a UpdateDepotSettingsAccepted with default headers values
func NewUpdateDepotSettingsAccepted() *UpdateDepotSettingsAccepted {
	return &UpdateDepotSettingsAccepted{}
}

/*
UpdateDepotSettingsAccepted describes a response with status code 202, with default header values.

Ok
*/
type UpdateDepotSettingsAccepted struct {
	Payload *models.DepotSettings
}

// IsSuccess returns true when this update depot settings accepted response has a 2xx status code
func (o *UpdateDepotSettingsAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update depot settings accepted response has a 3xx status code
func (o *UpdateDepotSettingsAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update depot settings accepted response has a 4xx status code
func (o *UpdateDepotSettingsAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this update depot settings accepted response has a 5xx status code
func (o *UpdateDepotSettingsAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this update depot settings accepted response a status code equal to that given
func (o *UpdateDepotSettingsAccepted) IsCode(code int) bool {
	return code == 202
}

// Code gets the status code for the update depot settings accepted response
func (o *UpdateDepotSettingsAccepted) Code() int {
	return 202
}

func (o *UpdateDepotSettingsAccepted) Error() string {
	return fmt.Sprintf("[PUT /v1/system/settings/depot][%d] updateDepotSettingsAccepted  %+v", 202, o.Payload)
}

func (o *UpdateDepotSettingsAccepted) String() string {
	return fmt.Sprintf("[PUT /v1/system/settings/depot][%d] updateDepotSettingsAccepted  %+v", 202, o.Payload)
}

func (o *UpdateDepotSettingsAccepted) GetPayload() *models.DepotSettings {
	return o.Payload
}

func (o *UpdateDepotSettingsAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DepotSettings)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateDepotSettingsBadRequest creates a UpdateDepotSettingsBadRequest with default headers values
func NewUpdateDepotSettingsBadRequest() *UpdateDepotSettingsBadRequest {
	return &UpdateDepotSettingsBadRequest{}
}

/*
UpdateDepotSettingsBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type UpdateDepotSettingsBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this update depot settings bad request response has a 2xx status code
func (o *UpdateDepotSettingsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update depot settings bad request response has a 3xx status code
func (o *UpdateDepotSettingsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update depot settings bad request response has a 4xx status code
func (o *UpdateDepotSettingsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this update depot settings bad request response has a 5xx status code
func (o *UpdateDepotSettingsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this update depot settings bad request response a status code equal to that given
func (o *UpdateDepotSettingsBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the update depot settings bad request response
func (o *UpdateDepotSettingsBadRequest) Code() int {
	return 400
}

func (o *UpdateDepotSettingsBadRequest) Error() string {
	return fmt.Sprintf("[PUT /v1/system/settings/depot][%d] updateDepotSettingsBadRequest  %+v", 400, o.Payload)
}

func (o *UpdateDepotSettingsBadRequest) String() string {
	return fmt.Sprintf("[PUT /v1/system/settings/depot][%d] updateDepotSettingsBadRequest  %+v", 400, o.Payload)
}

func (o *UpdateDepotSettingsBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateDepotSettingsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateDepotSettingsInternalServerError creates a UpdateDepotSettingsInternalServerError with default headers values
func NewUpdateDepotSettingsInternalServerError() *UpdateDepotSettingsInternalServerError {
	return &UpdateDepotSettingsInternalServerError{}
}

/*
UpdateDepotSettingsInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type UpdateDepotSettingsInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this update depot settings internal server error response has a 2xx status code
func (o *UpdateDepotSettingsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update depot settings internal server error response has a 3xx status code
func (o *UpdateDepotSettingsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update depot settings internal server error response has a 4xx status code
func (o *UpdateDepotSettingsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this update depot settings internal server error response has a 5xx status code
func (o *UpdateDepotSettingsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this update depot settings internal server error response a status code equal to that given
func (o *UpdateDepotSettingsInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the update depot settings internal server error response
func (o *UpdateDepotSettingsInternalServerError) Code() int {
	return 500
}

func (o *UpdateDepotSettingsInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /v1/system/settings/depot][%d] updateDepotSettingsInternalServerError  %+v", 500, o.Payload)
}

func (o *UpdateDepotSettingsInternalServerError) String() string {
	return fmt.Sprintf("[PUT /v1/system/settings/depot][%d] updateDepotSettingsInternalServerError  %+v", 500, o.Payload)
}

func (o *UpdateDepotSettingsInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateDepotSettingsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
