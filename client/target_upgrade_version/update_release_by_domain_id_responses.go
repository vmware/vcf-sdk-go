// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2023 VMware, Inc.
// SPDX-License-Identifier: BSD-2-Clause

package target_upgrade_version

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/vmware/vcf-sdk-go/models"
)

// UpdateReleaseByDomainIDReader is a Reader for the UpdateReleaseByDomainID structure.
type UpdateReleaseByDomainIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateReleaseByDomainIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateReleaseByDomainIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewUpdateReleaseByDomainIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUpdateReleaseByDomainIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[PATCH /v1/releases/domains/{domainId}] updateReleaseByDomainID", response, response.Code())
	}
}

// NewUpdateReleaseByDomainIDOK creates a UpdateReleaseByDomainIDOK with default headers values
func NewUpdateReleaseByDomainIDOK() *UpdateReleaseByDomainIDOK {
	return &UpdateReleaseByDomainIDOK{}
}

/*
UpdateReleaseByDomainIDOK describes a response with status code 200, with default header values.

Ok
*/
type UpdateReleaseByDomainIDOK struct {
}

// IsSuccess returns true when this update release by domain Id o k response has a 2xx status code
func (o *UpdateReleaseByDomainIDOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update release by domain Id o k response has a 3xx status code
func (o *UpdateReleaseByDomainIDOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update release by domain Id o k response has a 4xx status code
func (o *UpdateReleaseByDomainIDOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update release by domain Id o k response has a 5xx status code
func (o *UpdateReleaseByDomainIDOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update release by domain Id o k response a status code equal to that given
func (o *UpdateReleaseByDomainIDOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the update release by domain Id o k response
func (o *UpdateReleaseByDomainIDOK) Code() int {
	return 200
}

func (o *UpdateReleaseByDomainIDOK) Error() string {
	return fmt.Sprintf("[PATCH /v1/releases/domains/{domainId}][%d] updateReleaseByDomainIdOK ", 200)
}

func (o *UpdateReleaseByDomainIDOK) String() string {
	return fmt.Sprintf("[PATCH /v1/releases/domains/{domainId}][%d] updateReleaseByDomainIdOK ", 200)
}

func (o *UpdateReleaseByDomainIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateReleaseByDomainIDNotFound creates a UpdateReleaseByDomainIDNotFound with default headers values
func NewUpdateReleaseByDomainIDNotFound() *UpdateReleaseByDomainIDNotFound {
	return &UpdateReleaseByDomainIDNotFound{}
}

/*
UpdateReleaseByDomainIDNotFound describes a response with status code 404, with default header values.

Domain or VCF target release not found
*/
type UpdateReleaseByDomainIDNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this update release by domain Id not found response has a 2xx status code
func (o *UpdateReleaseByDomainIDNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update release by domain Id not found response has a 3xx status code
func (o *UpdateReleaseByDomainIDNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update release by domain Id not found response has a 4xx status code
func (o *UpdateReleaseByDomainIDNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this update release by domain Id not found response has a 5xx status code
func (o *UpdateReleaseByDomainIDNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this update release by domain Id not found response a status code equal to that given
func (o *UpdateReleaseByDomainIDNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the update release by domain Id not found response
func (o *UpdateReleaseByDomainIDNotFound) Code() int {
	return 404
}

func (o *UpdateReleaseByDomainIDNotFound) Error() string {
	return fmt.Sprintf("[PATCH /v1/releases/domains/{domainId}][%d] updateReleaseByDomainIdNotFound  %+v", 404, o.Payload)
}

func (o *UpdateReleaseByDomainIDNotFound) String() string {
	return fmt.Sprintf("[PATCH /v1/releases/domains/{domainId}][%d] updateReleaseByDomainIdNotFound  %+v", 404, o.Payload)
}

func (o *UpdateReleaseByDomainIDNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateReleaseByDomainIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateReleaseByDomainIDInternalServerError creates a UpdateReleaseByDomainIDInternalServerError with default headers values
func NewUpdateReleaseByDomainIDInternalServerError() *UpdateReleaseByDomainIDInternalServerError {
	return &UpdateReleaseByDomainIDInternalServerError{}
}

/*
UpdateReleaseByDomainIDInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type UpdateReleaseByDomainIDInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this update release by domain Id internal server error response has a 2xx status code
func (o *UpdateReleaseByDomainIDInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update release by domain Id internal server error response has a 3xx status code
func (o *UpdateReleaseByDomainIDInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update release by domain Id internal server error response has a 4xx status code
func (o *UpdateReleaseByDomainIDInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this update release by domain Id internal server error response has a 5xx status code
func (o *UpdateReleaseByDomainIDInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this update release by domain Id internal server error response a status code equal to that given
func (o *UpdateReleaseByDomainIDInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the update release by domain Id internal server error response
func (o *UpdateReleaseByDomainIDInternalServerError) Code() int {
	return 500
}

func (o *UpdateReleaseByDomainIDInternalServerError) Error() string {
	return fmt.Sprintf("[PATCH /v1/releases/domains/{domainId}][%d] updateReleaseByDomainIdInternalServerError  %+v", 500, o.Payload)
}

func (o *UpdateReleaseByDomainIDInternalServerError) String() string {
	return fmt.Sprintf("[PATCH /v1/releases/domains/{domainId}][%d] updateReleaseByDomainIdInternalServerError  %+v", 500, o.Payload)
}

func (o *UpdateReleaseByDomainIDInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateReleaseByDomainIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
