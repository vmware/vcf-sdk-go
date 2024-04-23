// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2023 VMware, Inc.
// SPDX-License-Identifier: BSD-2-Clause

package network_pools

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/vmware/vcf-sdk-go/models"
)

// UpdateNetworkPoolReader is a Reader for the UpdateNetworkPool structure.
type UpdateNetworkPoolReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateNetworkPoolReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateNetworkPoolOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateNetworkPoolBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateNetworkPoolNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUpdateNetworkPoolInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[PATCH /v1/network-pools/{id}] updateNetworkPool", response, response.Code())
	}
}

// NewUpdateNetworkPoolOK creates a UpdateNetworkPoolOK with default headers values
func NewUpdateNetworkPoolOK() *UpdateNetworkPoolOK {
	return &UpdateNetworkPoolOK{}
}

/*
UpdateNetworkPoolOK describes a response with status code 200, with default header values.

Network Pool update completed
*/
type UpdateNetworkPoolOK struct {
	Payload *models.NetworkPool
}

// IsSuccess returns true when this update network pool o k response has a 2xx status code
func (o *UpdateNetworkPoolOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update network pool o k response has a 3xx status code
func (o *UpdateNetworkPoolOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update network pool o k response has a 4xx status code
func (o *UpdateNetworkPoolOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update network pool o k response has a 5xx status code
func (o *UpdateNetworkPoolOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update network pool o k response a status code equal to that given
func (o *UpdateNetworkPoolOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the update network pool o k response
func (o *UpdateNetworkPoolOK) Code() int {
	return 200
}

func (o *UpdateNetworkPoolOK) Error() string {
	return fmt.Sprintf("[PATCH /v1/network-pools/{id}][%d] updateNetworkPoolOK  %+v", 200, o.Payload)
}

func (o *UpdateNetworkPoolOK) String() string {
	return fmt.Sprintf("[PATCH /v1/network-pools/{id}][%d] updateNetworkPoolOK  %+v", 200, o.Payload)
}

func (o *UpdateNetworkPoolOK) GetPayload() *models.NetworkPool {
	return o.Payload
}

func (o *UpdateNetworkPoolOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NetworkPool)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateNetworkPoolBadRequest creates a UpdateNetworkPoolBadRequest with default headers values
func NewUpdateNetworkPoolBadRequest() *UpdateNetworkPoolBadRequest {
	return &UpdateNetworkPoolBadRequest{}
}

/*
UpdateNetworkPoolBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type UpdateNetworkPoolBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this update network pool bad request response has a 2xx status code
func (o *UpdateNetworkPoolBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update network pool bad request response has a 3xx status code
func (o *UpdateNetworkPoolBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update network pool bad request response has a 4xx status code
func (o *UpdateNetworkPoolBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this update network pool bad request response has a 5xx status code
func (o *UpdateNetworkPoolBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this update network pool bad request response a status code equal to that given
func (o *UpdateNetworkPoolBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the update network pool bad request response
func (o *UpdateNetworkPoolBadRequest) Code() int {
	return 400
}

func (o *UpdateNetworkPoolBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /v1/network-pools/{id}][%d] updateNetworkPoolBadRequest  %+v", 400, o.Payload)
}

func (o *UpdateNetworkPoolBadRequest) String() string {
	return fmt.Sprintf("[PATCH /v1/network-pools/{id}][%d] updateNetworkPoolBadRequest  %+v", 400, o.Payload)
}

func (o *UpdateNetworkPoolBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateNetworkPoolBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateNetworkPoolNotFound creates a UpdateNetworkPoolNotFound with default headers values
func NewUpdateNetworkPoolNotFound() *UpdateNetworkPoolNotFound {
	return &UpdateNetworkPoolNotFound{}
}

/*
UpdateNetworkPoolNotFound describes a response with status code 404, with default header values.

Network Pool not found
*/
type UpdateNetworkPoolNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this update network pool not found response has a 2xx status code
func (o *UpdateNetworkPoolNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update network pool not found response has a 3xx status code
func (o *UpdateNetworkPoolNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update network pool not found response has a 4xx status code
func (o *UpdateNetworkPoolNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this update network pool not found response has a 5xx status code
func (o *UpdateNetworkPoolNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this update network pool not found response a status code equal to that given
func (o *UpdateNetworkPoolNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the update network pool not found response
func (o *UpdateNetworkPoolNotFound) Code() int {
	return 404
}

func (o *UpdateNetworkPoolNotFound) Error() string {
	return fmt.Sprintf("[PATCH /v1/network-pools/{id}][%d] updateNetworkPoolNotFound  %+v", 404, o.Payload)
}

func (o *UpdateNetworkPoolNotFound) String() string {
	return fmt.Sprintf("[PATCH /v1/network-pools/{id}][%d] updateNetworkPoolNotFound  %+v", 404, o.Payload)
}

func (o *UpdateNetworkPoolNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateNetworkPoolNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateNetworkPoolInternalServerError creates a UpdateNetworkPoolInternalServerError with default headers values
func NewUpdateNetworkPoolInternalServerError() *UpdateNetworkPoolInternalServerError {
	return &UpdateNetworkPoolInternalServerError{}
}

/*
UpdateNetworkPoolInternalServerError describes a response with status code 500, with default header values.

Unexpected error
*/
type UpdateNetworkPoolInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this update network pool internal server error response has a 2xx status code
func (o *UpdateNetworkPoolInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update network pool internal server error response has a 3xx status code
func (o *UpdateNetworkPoolInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update network pool internal server error response has a 4xx status code
func (o *UpdateNetworkPoolInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this update network pool internal server error response has a 5xx status code
func (o *UpdateNetworkPoolInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this update network pool internal server error response a status code equal to that given
func (o *UpdateNetworkPoolInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the update network pool internal server error response
func (o *UpdateNetworkPoolInternalServerError) Code() int {
	return 500
}

func (o *UpdateNetworkPoolInternalServerError) Error() string {
	return fmt.Sprintf("[PATCH /v1/network-pools/{id}][%d] updateNetworkPoolInternalServerError  %+v", 500, o.Payload)
}

func (o *UpdateNetworkPoolInternalServerError) String() string {
	return fmt.Sprintf("[PATCH /v1/network-pools/{id}][%d] updateNetworkPoolInternalServerError  %+v", 500, o.Payload)
}

func (o *UpdateNetworkPoolInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateNetworkPoolInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
