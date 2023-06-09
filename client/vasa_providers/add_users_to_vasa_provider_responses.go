// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2023 VMware, Inc.
// SPDX-License-Identifier: BSD-2-Clause

package vasa_providers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/vmware/vcf-sdk-go/models"
)

// AddUsersToVasaProviderReader is a Reader for the AddUsersToVasaProvider structure.
type AddUsersToVasaProviderReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddUsersToVasaProviderReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewAddUsersToVasaProviderCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewAddUsersToVasaProviderBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewAddUsersToVasaProviderNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewAddUsersToVasaProviderInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewAddUsersToVasaProviderCreated creates a AddUsersToVasaProviderCreated with default headers values
func NewAddUsersToVasaProviderCreated() *AddUsersToVasaProviderCreated {
	return &AddUsersToVasaProviderCreated{}
}

/*
AddUsersToVasaProviderCreated describes a response with status code 201, with default header values.

Created
*/
type AddUsersToVasaProviderCreated struct {
	Payload *models.VasaProvider
}

// IsSuccess returns true when this add users to vasa provider created response has a 2xx status code
func (o *AddUsersToVasaProviderCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this add users to vasa provider created response has a 3xx status code
func (o *AddUsersToVasaProviderCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this add users to vasa provider created response has a 4xx status code
func (o *AddUsersToVasaProviderCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this add users to vasa provider created response has a 5xx status code
func (o *AddUsersToVasaProviderCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this add users to vasa provider created response a status code equal to that given
func (o *AddUsersToVasaProviderCreated) IsCode(code int) bool {
	return code == 201
}

func (o *AddUsersToVasaProviderCreated) Error() string {
	return fmt.Sprintf("[POST /v1/vasa-providers/{id}/users][%d] addUsersToVasaProviderCreated  %+v", 201, o.Payload)
}

func (o *AddUsersToVasaProviderCreated) String() string {
	return fmt.Sprintf("[POST /v1/vasa-providers/{id}/users][%d] addUsersToVasaProviderCreated  %+v", 201, o.Payload)
}

func (o *AddUsersToVasaProviderCreated) GetPayload() *models.VasaProvider {
	return o.Payload
}

func (o *AddUsersToVasaProviderCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VasaProvider)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddUsersToVasaProviderBadRequest creates a AddUsersToVasaProviderBadRequest with default headers values
func NewAddUsersToVasaProviderBadRequest() *AddUsersToVasaProviderBadRequest {
	return &AddUsersToVasaProviderBadRequest{}
}

/*
AddUsersToVasaProviderBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type AddUsersToVasaProviderBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this add users to vasa provider bad request response has a 2xx status code
func (o *AddUsersToVasaProviderBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this add users to vasa provider bad request response has a 3xx status code
func (o *AddUsersToVasaProviderBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this add users to vasa provider bad request response has a 4xx status code
func (o *AddUsersToVasaProviderBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this add users to vasa provider bad request response has a 5xx status code
func (o *AddUsersToVasaProviderBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this add users to vasa provider bad request response a status code equal to that given
func (o *AddUsersToVasaProviderBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *AddUsersToVasaProviderBadRequest) Error() string {
	return fmt.Sprintf("[POST /v1/vasa-providers/{id}/users][%d] addUsersToVasaProviderBadRequest  %+v", 400, o.Payload)
}

func (o *AddUsersToVasaProviderBadRequest) String() string {
	return fmt.Sprintf("[POST /v1/vasa-providers/{id}/users][%d] addUsersToVasaProviderBadRequest  %+v", 400, o.Payload)
}

func (o *AddUsersToVasaProviderBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *AddUsersToVasaProviderBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddUsersToVasaProviderNotFound creates a AddUsersToVasaProviderNotFound with default headers values
func NewAddUsersToVasaProviderNotFound() *AddUsersToVasaProviderNotFound {
	return &AddUsersToVasaProviderNotFound{}
}

/*
AddUsersToVasaProviderNotFound describes a response with status code 404, with default header values.

VASA Provider not found
*/
type AddUsersToVasaProviderNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this add users to vasa provider not found response has a 2xx status code
func (o *AddUsersToVasaProviderNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this add users to vasa provider not found response has a 3xx status code
func (o *AddUsersToVasaProviderNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this add users to vasa provider not found response has a 4xx status code
func (o *AddUsersToVasaProviderNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this add users to vasa provider not found response has a 5xx status code
func (o *AddUsersToVasaProviderNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this add users to vasa provider not found response a status code equal to that given
func (o *AddUsersToVasaProviderNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *AddUsersToVasaProviderNotFound) Error() string {
	return fmt.Sprintf("[POST /v1/vasa-providers/{id}/users][%d] addUsersToVasaProviderNotFound  %+v", 404, o.Payload)
}

func (o *AddUsersToVasaProviderNotFound) String() string {
	return fmt.Sprintf("[POST /v1/vasa-providers/{id}/users][%d] addUsersToVasaProviderNotFound  %+v", 404, o.Payload)
}

func (o *AddUsersToVasaProviderNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *AddUsersToVasaProviderNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddUsersToVasaProviderInternalServerError creates a AddUsersToVasaProviderInternalServerError with default headers values
func NewAddUsersToVasaProviderInternalServerError() *AddUsersToVasaProviderInternalServerError {
	return &AddUsersToVasaProviderInternalServerError{}
}

/*
AddUsersToVasaProviderInternalServerError describes a response with status code 500, with default header values.

InternalServerError
*/
type AddUsersToVasaProviderInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this add users to vasa provider internal server error response has a 2xx status code
func (o *AddUsersToVasaProviderInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this add users to vasa provider internal server error response has a 3xx status code
func (o *AddUsersToVasaProviderInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this add users to vasa provider internal server error response has a 4xx status code
func (o *AddUsersToVasaProviderInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this add users to vasa provider internal server error response has a 5xx status code
func (o *AddUsersToVasaProviderInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this add users to vasa provider internal server error response a status code equal to that given
func (o *AddUsersToVasaProviderInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *AddUsersToVasaProviderInternalServerError) Error() string {
	return fmt.Sprintf("[POST /v1/vasa-providers/{id}/users][%d] addUsersToVasaProviderInternalServerError  %+v", 500, o.Payload)
}

func (o *AddUsersToVasaProviderInternalServerError) String() string {
	return fmt.Sprintf("[POST /v1/vasa-providers/{id}/users][%d] addUsersToVasaProviderInternalServerError  %+v", 500, o.Payload)
}

func (o *AddUsersToVasaProviderInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *AddUsersToVasaProviderInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
