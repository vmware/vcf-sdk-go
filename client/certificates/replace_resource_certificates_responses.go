// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2023 VMware, Inc.
// SPDX-License-Identifier: BSD-2-Clause

package certificates

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/vmware/vcf-sdk-go/models"
)

// ReplaceResourceCertificatesReader is a Reader for the ReplaceResourceCertificates structure.
type ReplaceResourceCertificatesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ReplaceResourceCertificatesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewReplaceResourceCertificatesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewReplaceResourceCertificatesAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewReplaceResourceCertificatesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewReplaceResourceCertificatesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewReplaceResourceCertificatesOK creates a ReplaceResourceCertificatesOK with default headers values
func NewReplaceResourceCertificatesOK() *ReplaceResourceCertificatesOK {
	return &ReplaceResourceCertificatesOK{}
}

/*
ReplaceResourceCertificatesOK describes a response with status code 200, with default header values.

OK
*/
type ReplaceResourceCertificatesOK struct {
	Payload *models.Task
}

// IsSuccess returns true when this replace resource certificates o k response has a 2xx status code
func (o *ReplaceResourceCertificatesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this replace resource certificates o k response has a 3xx status code
func (o *ReplaceResourceCertificatesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this replace resource certificates o k response has a 4xx status code
func (o *ReplaceResourceCertificatesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this replace resource certificates o k response has a 5xx status code
func (o *ReplaceResourceCertificatesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this replace resource certificates o k response a status code equal to that given
func (o *ReplaceResourceCertificatesOK) IsCode(code int) bool {
	return code == 200
}

func (o *ReplaceResourceCertificatesOK) Error() string {
	return fmt.Sprintf("[PUT /v1/domains/{id}/resource-certificates][%d] replaceResourceCertificatesOK  %+v", 200, o.Payload)
}

func (o *ReplaceResourceCertificatesOK) String() string {
	return fmt.Sprintf("[PUT /v1/domains/{id}/resource-certificates][%d] replaceResourceCertificatesOK  %+v", 200, o.Payload)
}

func (o *ReplaceResourceCertificatesOK) GetPayload() *models.Task {
	return o.Payload
}

func (o *ReplaceResourceCertificatesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Task)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReplaceResourceCertificatesAccepted creates a ReplaceResourceCertificatesAccepted with default headers values
func NewReplaceResourceCertificatesAccepted() *ReplaceResourceCertificatesAccepted {
	return &ReplaceResourceCertificatesAccepted{}
}

/*
ReplaceResourceCertificatesAccepted describes a response with status code 202, with default header values.

Accepted
*/
type ReplaceResourceCertificatesAccepted struct {
	Payload *models.Task
}

// IsSuccess returns true when this replace resource certificates accepted response has a 2xx status code
func (o *ReplaceResourceCertificatesAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this replace resource certificates accepted response has a 3xx status code
func (o *ReplaceResourceCertificatesAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this replace resource certificates accepted response has a 4xx status code
func (o *ReplaceResourceCertificatesAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this replace resource certificates accepted response has a 5xx status code
func (o *ReplaceResourceCertificatesAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this replace resource certificates accepted response a status code equal to that given
func (o *ReplaceResourceCertificatesAccepted) IsCode(code int) bool {
	return code == 202
}

func (o *ReplaceResourceCertificatesAccepted) Error() string {
	return fmt.Sprintf("[PUT /v1/domains/{id}/resource-certificates][%d] replaceResourceCertificatesAccepted  %+v", 202, o.Payload)
}

func (o *ReplaceResourceCertificatesAccepted) String() string {
	return fmt.Sprintf("[PUT /v1/domains/{id}/resource-certificates][%d] replaceResourceCertificatesAccepted  %+v", 202, o.Payload)
}

func (o *ReplaceResourceCertificatesAccepted) GetPayload() *models.Task {
	return o.Payload
}

func (o *ReplaceResourceCertificatesAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Task)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReplaceResourceCertificatesNotFound creates a ReplaceResourceCertificatesNotFound with default headers values
func NewReplaceResourceCertificatesNotFound() *ReplaceResourceCertificatesNotFound {
	return &ReplaceResourceCertificatesNotFound{}
}

/*
ReplaceResourceCertificatesNotFound describes a response with status code 404, with default header values.

Not Found
*/
type ReplaceResourceCertificatesNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this replace resource certificates not found response has a 2xx status code
func (o *ReplaceResourceCertificatesNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this replace resource certificates not found response has a 3xx status code
func (o *ReplaceResourceCertificatesNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this replace resource certificates not found response has a 4xx status code
func (o *ReplaceResourceCertificatesNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this replace resource certificates not found response has a 5xx status code
func (o *ReplaceResourceCertificatesNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this replace resource certificates not found response a status code equal to that given
func (o *ReplaceResourceCertificatesNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *ReplaceResourceCertificatesNotFound) Error() string {
	return fmt.Sprintf("[PUT /v1/domains/{id}/resource-certificates][%d] replaceResourceCertificatesNotFound  %+v", 404, o.Payload)
}

func (o *ReplaceResourceCertificatesNotFound) String() string {
	return fmt.Sprintf("[PUT /v1/domains/{id}/resource-certificates][%d] replaceResourceCertificatesNotFound  %+v", 404, o.Payload)
}

func (o *ReplaceResourceCertificatesNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *ReplaceResourceCertificatesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReplaceResourceCertificatesInternalServerError creates a ReplaceResourceCertificatesInternalServerError with default headers values
func NewReplaceResourceCertificatesInternalServerError() *ReplaceResourceCertificatesInternalServerError {
	return &ReplaceResourceCertificatesInternalServerError{}
}

/*
ReplaceResourceCertificatesInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type ReplaceResourceCertificatesInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this replace resource certificates internal server error response has a 2xx status code
func (o *ReplaceResourceCertificatesInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this replace resource certificates internal server error response has a 3xx status code
func (o *ReplaceResourceCertificatesInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this replace resource certificates internal server error response has a 4xx status code
func (o *ReplaceResourceCertificatesInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this replace resource certificates internal server error response has a 5xx status code
func (o *ReplaceResourceCertificatesInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this replace resource certificates internal server error response a status code equal to that given
func (o *ReplaceResourceCertificatesInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *ReplaceResourceCertificatesInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /v1/domains/{id}/resource-certificates][%d] replaceResourceCertificatesInternalServerError  %+v", 500, o.Payload)
}

func (o *ReplaceResourceCertificatesInternalServerError) String() string {
	return fmt.Sprintf("[PUT /v1/domains/{id}/resource-certificates][%d] replaceResourceCertificatesInternalServerError  %+v", 500, o.Payload)
}

func (o *ReplaceResourceCertificatesInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *ReplaceResourceCertificatesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
