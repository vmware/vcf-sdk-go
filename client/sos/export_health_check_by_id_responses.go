// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2023 VMware, Inc.
// SPDX-License-Identifier: BSD-2-Clause

package sos

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/vmware/vcf-sdk-go/models"
)

// ExportHealthCheckByIDReader is a Reader for the ExportHealthCheckByID structure.
type ExportHealthCheckByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExportHealthCheckByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewExportHealthCheckByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewExportHealthCheckByIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewExportHealthCheckByIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewExportHealthCheckByIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewExportHealthCheckByIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /v1/system/health-summary/{id}/data] exportHealthCheckByID", response, response.Code())
	}
}

// NewExportHealthCheckByIDOK creates a ExportHealthCheckByIDOK with default headers values
func NewExportHealthCheckByIDOK() *ExportHealthCheckByIDOK {
	return &ExportHealthCheckByIDOK{}
}

/*
ExportHealthCheckByIDOK describes a response with status code 200, with default header values.

Ok
*/
type ExportHealthCheckByIDOK struct {
	Payload strfmt.Base64
}

// IsSuccess returns true when this export health check by Id o k response has a 2xx status code
func (o *ExportHealthCheckByIDOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this export health check by Id o k response has a 3xx status code
func (o *ExportHealthCheckByIDOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this export health check by Id o k response has a 4xx status code
func (o *ExportHealthCheckByIDOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this export health check by Id o k response has a 5xx status code
func (o *ExportHealthCheckByIDOK) IsServerError() bool {
	return false
}

// IsCode returns true when this export health check by Id o k response a status code equal to that given
func (o *ExportHealthCheckByIDOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the export health check by Id o k response
func (o *ExportHealthCheckByIDOK) Code() int {
	return 200
}

func (o *ExportHealthCheckByIDOK) Error() string {
	return fmt.Sprintf("[GET /v1/system/health-summary/{id}/data][%d] exportHealthCheckByIdOK  %+v", 200, o.Payload)
}

func (o *ExportHealthCheckByIDOK) String() string {
	return fmt.Sprintf("[GET /v1/system/health-summary/{id}/data][%d] exportHealthCheckByIdOK  %+v", 200, o.Payload)
}

func (o *ExportHealthCheckByIDOK) GetPayload() strfmt.Base64 {
	return o.Payload
}

func (o *ExportHealthCheckByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewExportHealthCheckByIDBadRequest creates a ExportHealthCheckByIDBadRequest with default headers values
func NewExportHealthCheckByIDBadRequest() *ExportHealthCheckByIDBadRequest {
	return &ExportHealthCheckByIDBadRequest{}
}

/*
ExportHealthCheckByIDBadRequest describes a response with status code 400, with default header values.

Bad request! Invalid Headers or Data. Error: {error}.
*/
type ExportHealthCheckByIDBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this export health check by Id bad request response has a 2xx status code
func (o *ExportHealthCheckByIDBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this export health check by Id bad request response has a 3xx status code
func (o *ExportHealthCheckByIDBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this export health check by Id bad request response has a 4xx status code
func (o *ExportHealthCheckByIDBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this export health check by Id bad request response has a 5xx status code
func (o *ExportHealthCheckByIDBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this export health check by Id bad request response a status code equal to that given
func (o *ExportHealthCheckByIDBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the export health check by Id bad request response
func (o *ExportHealthCheckByIDBadRequest) Code() int {
	return 400
}

func (o *ExportHealthCheckByIDBadRequest) Error() string {
	return fmt.Sprintf("[GET /v1/system/health-summary/{id}/data][%d] exportHealthCheckByIdBadRequest  %+v", 400, o.Payload)
}

func (o *ExportHealthCheckByIDBadRequest) String() string {
	return fmt.Sprintf("[GET /v1/system/health-summary/{id}/data][%d] exportHealthCheckByIdBadRequest  %+v", 400, o.Payload)
}

func (o *ExportHealthCheckByIDBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *ExportHealthCheckByIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewExportHealthCheckByIDUnauthorized creates a ExportHealthCheckByIDUnauthorized with default headers values
func NewExportHealthCheckByIDUnauthorized() *ExportHealthCheckByIDUnauthorized {
	return &ExportHealthCheckByIDUnauthorized{}
}

/*
ExportHealthCheckByIDUnauthorized describes a response with status code 401, with default header values.

Bad request! Authorization Header is missing or not in correct format.
*/
type ExportHealthCheckByIDUnauthorized struct {
	Payload *models.Error
}

// IsSuccess returns true when this export health check by Id unauthorized response has a 2xx status code
func (o *ExportHealthCheckByIDUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this export health check by Id unauthorized response has a 3xx status code
func (o *ExportHealthCheckByIDUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this export health check by Id unauthorized response has a 4xx status code
func (o *ExportHealthCheckByIDUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this export health check by Id unauthorized response has a 5xx status code
func (o *ExportHealthCheckByIDUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this export health check by Id unauthorized response a status code equal to that given
func (o *ExportHealthCheckByIDUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the export health check by Id unauthorized response
func (o *ExportHealthCheckByIDUnauthorized) Code() int {
	return 401
}

func (o *ExportHealthCheckByIDUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v1/system/health-summary/{id}/data][%d] exportHealthCheckByIdUnauthorized  %+v", 401, o.Payload)
}

func (o *ExportHealthCheckByIDUnauthorized) String() string {
	return fmt.Sprintf("[GET /v1/system/health-summary/{id}/data][%d] exportHealthCheckByIdUnauthorized  %+v", 401, o.Payload)
}

func (o *ExportHealthCheckByIDUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *ExportHealthCheckByIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewExportHealthCheckByIDNotFound creates a ExportHealthCheckByIDNotFound with default headers values
func NewExportHealthCheckByIDNotFound() *ExportHealthCheckByIDNotFound {
	return &ExportHealthCheckByIDNotFound{}
}

/*
ExportHealthCheckByIDNotFound describes a response with status code 404, with default header values.

Bundle not found or not available for download. Id:{id}.
*/
type ExportHealthCheckByIDNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this export health check by Id not found response has a 2xx status code
func (o *ExportHealthCheckByIDNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this export health check by Id not found response has a 3xx status code
func (o *ExportHealthCheckByIDNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this export health check by Id not found response has a 4xx status code
func (o *ExportHealthCheckByIDNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this export health check by Id not found response has a 5xx status code
func (o *ExportHealthCheckByIDNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this export health check by Id not found response a status code equal to that given
func (o *ExportHealthCheckByIDNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the export health check by Id not found response
func (o *ExportHealthCheckByIDNotFound) Code() int {
	return 404
}

func (o *ExportHealthCheckByIDNotFound) Error() string {
	return fmt.Sprintf("[GET /v1/system/health-summary/{id}/data][%d] exportHealthCheckByIdNotFound  %+v", 404, o.Payload)
}

func (o *ExportHealthCheckByIDNotFound) String() string {
	return fmt.Sprintf("[GET /v1/system/health-summary/{id}/data][%d] exportHealthCheckByIdNotFound  %+v", 404, o.Payload)
}

func (o *ExportHealthCheckByIDNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *ExportHealthCheckByIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewExportHealthCheckByIDInternalServerError creates a ExportHealthCheckByIDInternalServerError with default headers values
func NewExportHealthCheckByIDInternalServerError() *ExportHealthCheckByIDInternalServerError {
	return &ExportHealthCheckByIDInternalServerError{}
}

/*
ExportHealthCheckByIDInternalServerError describes a response with status code 500, with default header values.

Something went wrong. Internal server error occurred. Error {error}.
*/
type ExportHealthCheckByIDInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this export health check by Id internal server error response has a 2xx status code
func (o *ExportHealthCheckByIDInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this export health check by Id internal server error response has a 3xx status code
func (o *ExportHealthCheckByIDInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this export health check by Id internal server error response has a 4xx status code
func (o *ExportHealthCheckByIDInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this export health check by Id internal server error response has a 5xx status code
func (o *ExportHealthCheckByIDInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this export health check by Id internal server error response a status code equal to that given
func (o *ExportHealthCheckByIDInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the export health check by Id internal server error response
func (o *ExportHealthCheckByIDInternalServerError) Code() int {
	return 500
}

func (o *ExportHealthCheckByIDInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/system/health-summary/{id}/data][%d] exportHealthCheckByIdInternalServerError  %+v", 500, o.Payload)
}

func (o *ExportHealthCheckByIDInternalServerError) String() string {
	return fmt.Sprintf("[GET /v1/system/health-summary/{id}/data][%d] exportHealthCheckByIdInternalServerError  %+v", 500, o.Payload)
}

func (o *ExportHealthCheckByIDInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *ExportHealthCheckByIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
