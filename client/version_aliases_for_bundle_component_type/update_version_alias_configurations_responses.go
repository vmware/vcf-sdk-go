// Code generated by go-swagger; DO NOT EDIT.

// Copyright © 2023 VMware, Inc. All Rights Reserved.
// Code licenced under: BSD-2

package version_aliases_for_bundle_component_type

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/vmware/vcf-sdk-go/models"
)

// UpdateVersionAliasConfigurationsReader is a Reader for the UpdateVersionAliasConfigurations structure.
type UpdateVersionAliasConfigurationsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateVersionAliasConfigurationsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateVersionAliasConfigurationsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateVersionAliasConfigurationsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUpdateVersionAliasConfigurationsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateVersionAliasConfigurationsOK creates a UpdateVersionAliasConfigurationsOK with default headers values
func NewUpdateVersionAliasConfigurationsOK() *UpdateVersionAliasConfigurationsOK {
	return &UpdateVersionAliasConfigurationsOK{}
}

/*
UpdateVersionAliasConfigurationsOK describes a response with status code 200, with default header values.

Ok
*/
type UpdateVersionAliasConfigurationsOK struct {
	Payload *models.PageOfVersionAliasesForBundleComponentType
}

// IsSuccess returns true when this update version alias configurations o k response has a 2xx status code
func (o *UpdateVersionAliasConfigurationsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update version alias configurations o k response has a 3xx status code
func (o *UpdateVersionAliasConfigurationsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update version alias configurations o k response has a 4xx status code
func (o *UpdateVersionAliasConfigurationsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update version alias configurations o k response has a 5xx status code
func (o *UpdateVersionAliasConfigurationsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update version alias configurations o k response a status code equal to that given
func (o *UpdateVersionAliasConfigurationsOK) IsCode(code int) bool {
	return code == 200
}

func (o *UpdateVersionAliasConfigurationsOK) Error() string {
	return fmt.Sprintf("[PUT /v1/system/settings/version-aliases][%d] updateVersionAliasConfigurationsOK  %+v", 200, o.Payload)
}

func (o *UpdateVersionAliasConfigurationsOK) String() string {
	return fmt.Sprintf("[PUT /v1/system/settings/version-aliases][%d] updateVersionAliasConfigurationsOK  %+v", 200, o.Payload)
}

func (o *UpdateVersionAliasConfigurationsOK) GetPayload() *models.PageOfVersionAliasesForBundleComponentType {
	return o.Payload
}

func (o *UpdateVersionAliasConfigurationsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PageOfVersionAliasesForBundleComponentType)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateVersionAliasConfigurationsBadRequest creates a UpdateVersionAliasConfigurationsBadRequest with default headers values
func NewUpdateVersionAliasConfigurationsBadRequest() *UpdateVersionAliasConfigurationsBadRequest {
	return &UpdateVersionAliasConfigurationsBadRequest{}
}

/*
UpdateVersionAliasConfigurationsBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type UpdateVersionAliasConfigurationsBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this update version alias configurations bad request response has a 2xx status code
func (o *UpdateVersionAliasConfigurationsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update version alias configurations bad request response has a 3xx status code
func (o *UpdateVersionAliasConfigurationsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update version alias configurations bad request response has a 4xx status code
func (o *UpdateVersionAliasConfigurationsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this update version alias configurations bad request response has a 5xx status code
func (o *UpdateVersionAliasConfigurationsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this update version alias configurations bad request response a status code equal to that given
func (o *UpdateVersionAliasConfigurationsBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *UpdateVersionAliasConfigurationsBadRequest) Error() string {
	return fmt.Sprintf("[PUT /v1/system/settings/version-aliases][%d] updateVersionAliasConfigurationsBadRequest  %+v", 400, o.Payload)
}

func (o *UpdateVersionAliasConfigurationsBadRequest) String() string {
	return fmt.Sprintf("[PUT /v1/system/settings/version-aliases][%d] updateVersionAliasConfigurationsBadRequest  %+v", 400, o.Payload)
}

func (o *UpdateVersionAliasConfigurationsBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateVersionAliasConfigurationsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateVersionAliasConfigurationsInternalServerError creates a UpdateVersionAliasConfigurationsInternalServerError with default headers values
func NewUpdateVersionAliasConfigurationsInternalServerError() *UpdateVersionAliasConfigurationsInternalServerError {
	return &UpdateVersionAliasConfigurationsInternalServerError{}
}

/*
UpdateVersionAliasConfigurationsInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type UpdateVersionAliasConfigurationsInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this update version alias configurations internal server error response has a 2xx status code
func (o *UpdateVersionAliasConfigurationsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update version alias configurations internal server error response has a 3xx status code
func (o *UpdateVersionAliasConfigurationsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update version alias configurations internal server error response has a 4xx status code
func (o *UpdateVersionAliasConfigurationsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this update version alias configurations internal server error response has a 5xx status code
func (o *UpdateVersionAliasConfigurationsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this update version alias configurations internal server error response a status code equal to that given
func (o *UpdateVersionAliasConfigurationsInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *UpdateVersionAliasConfigurationsInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /v1/system/settings/version-aliases][%d] updateVersionAliasConfigurationsInternalServerError  %+v", 500, o.Payload)
}

func (o *UpdateVersionAliasConfigurationsInternalServerError) String() string {
	return fmt.Sprintf("[PUT /v1/system/settings/version-aliases][%d] updateVersionAliasConfigurationsInternalServerError  %+v", 500, o.Payload)
}

func (o *UpdateVersionAliasConfigurationsInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateVersionAliasConfigurationsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
