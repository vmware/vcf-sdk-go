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

// AddVasaProviderStorageContainerReader is a Reader for the AddVasaProviderStorageContainer structure.
type AddVasaProviderStorageContainerReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddVasaProviderStorageContainerReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewAddVasaProviderStorageContainerCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewAddVasaProviderStorageContainerBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewAddVasaProviderStorageContainerNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewAddVasaProviderStorageContainerInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /v1/vasa-providers/{id}/storage-containers] addVasaProviderStorageContainer", response, response.Code())
	}
}

// NewAddVasaProviderStorageContainerCreated creates a AddVasaProviderStorageContainerCreated with default headers values
func NewAddVasaProviderStorageContainerCreated() *AddVasaProviderStorageContainerCreated {
	return &AddVasaProviderStorageContainerCreated{}
}

/*
AddVasaProviderStorageContainerCreated describes a response with status code 201, with default header values.

Created
*/
type AddVasaProviderStorageContainerCreated struct {
	Payload *models.VasaProvider
}

// IsSuccess returns true when this add vasa provider storage container created response has a 2xx status code
func (o *AddVasaProviderStorageContainerCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this add vasa provider storage container created response has a 3xx status code
func (o *AddVasaProviderStorageContainerCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this add vasa provider storage container created response has a 4xx status code
func (o *AddVasaProviderStorageContainerCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this add vasa provider storage container created response has a 5xx status code
func (o *AddVasaProviderStorageContainerCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this add vasa provider storage container created response a status code equal to that given
func (o *AddVasaProviderStorageContainerCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the add vasa provider storage container created response
func (o *AddVasaProviderStorageContainerCreated) Code() int {
	return 201
}

func (o *AddVasaProviderStorageContainerCreated) Error() string {
	return fmt.Sprintf("[POST /v1/vasa-providers/{id}/storage-containers][%d] addVasaProviderStorageContainerCreated  %+v", 201, o.Payload)
}

func (o *AddVasaProviderStorageContainerCreated) String() string {
	return fmt.Sprintf("[POST /v1/vasa-providers/{id}/storage-containers][%d] addVasaProviderStorageContainerCreated  %+v", 201, o.Payload)
}

func (o *AddVasaProviderStorageContainerCreated) GetPayload() *models.VasaProvider {
	return o.Payload
}

func (o *AddVasaProviderStorageContainerCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VasaProvider)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddVasaProviderStorageContainerBadRequest creates a AddVasaProviderStorageContainerBadRequest with default headers values
func NewAddVasaProviderStorageContainerBadRequest() *AddVasaProviderStorageContainerBadRequest {
	return &AddVasaProviderStorageContainerBadRequest{}
}

/*
AddVasaProviderStorageContainerBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type AddVasaProviderStorageContainerBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this add vasa provider storage container bad request response has a 2xx status code
func (o *AddVasaProviderStorageContainerBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this add vasa provider storage container bad request response has a 3xx status code
func (o *AddVasaProviderStorageContainerBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this add vasa provider storage container bad request response has a 4xx status code
func (o *AddVasaProviderStorageContainerBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this add vasa provider storage container bad request response has a 5xx status code
func (o *AddVasaProviderStorageContainerBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this add vasa provider storage container bad request response a status code equal to that given
func (o *AddVasaProviderStorageContainerBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the add vasa provider storage container bad request response
func (o *AddVasaProviderStorageContainerBadRequest) Code() int {
	return 400
}

func (o *AddVasaProviderStorageContainerBadRequest) Error() string {
	return fmt.Sprintf("[POST /v1/vasa-providers/{id}/storage-containers][%d] addVasaProviderStorageContainerBadRequest  %+v", 400, o.Payload)
}

func (o *AddVasaProviderStorageContainerBadRequest) String() string {
	return fmt.Sprintf("[POST /v1/vasa-providers/{id}/storage-containers][%d] addVasaProviderStorageContainerBadRequest  %+v", 400, o.Payload)
}

func (o *AddVasaProviderStorageContainerBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *AddVasaProviderStorageContainerBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddVasaProviderStorageContainerNotFound creates a AddVasaProviderStorageContainerNotFound with default headers values
func NewAddVasaProviderStorageContainerNotFound() *AddVasaProviderStorageContainerNotFound {
	return &AddVasaProviderStorageContainerNotFound{}
}

/*
AddVasaProviderStorageContainerNotFound describes a response with status code 404, with default header values.

VASA Provider not found
*/
type AddVasaProviderStorageContainerNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this add vasa provider storage container not found response has a 2xx status code
func (o *AddVasaProviderStorageContainerNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this add vasa provider storage container not found response has a 3xx status code
func (o *AddVasaProviderStorageContainerNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this add vasa provider storage container not found response has a 4xx status code
func (o *AddVasaProviderStorageContainerNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this add vasa provider storage container not found response has a 5xx status code
func (o *AddVasaProviderStorageContainerNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this add vasa provider storage container not found response a status code equal to that given
func (o *AddVasaProviderStorageContainerNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the add vasa provider storage container not found response
func (o *AddVasaProviderStorageContainerNotFound) Code() int {
	return 404
}

func (o *AddVasaProviderStorageContainerNotFound) Error() string {
	return fmt.Sprintf("[POST /v1/vasa-providers/{id}/storage-containers][%d] addVasaProviderStorageContainerNotFound  %+v", 404, o.Payload)
}

func (o *AddVasaProviderStorageContainerNotFound) String() string {
	return fmt.Sprintf("[POST /v1/vasa-providers/{id}/storage-containers][%d] addVasaProviderStorageContainerNotFound  %+v", 404, o.Payload)
}

func (o *AddVasaProviderStorageContainerNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *AddVasaProviderStorageContainerNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddVasaProviderStorageContainerInternalServerError creates a AddVasaProviderStorageContainerInternalServerError with default headers values
func NewAddVasaProviderStorageContainerInternalServerError() *AddVasaProviderStorageContainerInternalServerError {
	return &AddVasaProviderStorageContainerInternalServerError{}
}

/*
AddVasaProviderStorageContainerInternalServerError describes a response with status code 500, with default header values.

InternalServerError
*/
type AddVasaProviderStorageContainerInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this add vasa provider storage container internal server error response has a 2xx status code
func (o *AddVasaProviderStorageContainerInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this add vasa provider storage container internal server error response has a 3xx status code
func (o *AddVasaProviderStorageContainerInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this add vasa provider storage container internal server error response has a 4xx status code
func (o *AddVasaProviderStorageContainerInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this add vasa provider storage container internal server error response has a 5xx status code
func (o *AddVasaProviderStorageContainerInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this add vasa provider storage container internal server error response a status code equal to that given
func (o *AddVasaProviderStorageContainerInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the add vasa provider storage container internal server error response
func (o *AddVasaProviderStorageContainerInternalServerError) Code() int {
	return 500
}

func (o *AddVasaProviderStorageContainerInternalServerError) Error() string {
	return fmt.Sprintf("[POST /v1/vasa-providers/{id}/storage-containers][%d] addVasaProviderStorageContainerInternalServerError  %+v", 500, o.Payload)
}

func (o *AddVasaProviderStorageContainerInternalServerError) String() string {
	return fmt.Sprintf("[POST /v1/vasa-providers/{id}/storage-containers][%d] addVasaProviderStorageContainerInternalServerError  %+v", 500, o.Payload)
}

func (o *AddVasaProviderStorageContainerInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *AddVasaProviderStorageContainerInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
