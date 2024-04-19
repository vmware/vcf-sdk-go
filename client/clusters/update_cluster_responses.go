// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2023 VMware, Inc.
// SPDX-License-Identifier: BSD-2-Clause

package clusters

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/vmware/vcf-sdk-go/models"
)

// UpdateClusterReader is a Reader for the UpdateCluster structure.
type UpdateClusterReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateClusterReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateClusterOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewUpdateClusterAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateClusterBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateClusterNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUpdateClusterInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[PATCH /v1/clusters/{id}] updateCluster", response, response.Code())
	}
}

// NewUpdateClusterOK creates a UpdateClusterOK with default headers values
func NewUpdateClusterOK() *UpdateClusterOK {
	return &UpdateClusterOK{}
}

/*
UpdateClusterOK describes a response with status code 200, with default header values.

Ok
*/
type UpdateClusterOK struct {
	Payload *models.Task
}

// IsSuccess returns true when this update cluster o k response has a 2xx status code
func (o *UpdateClusterOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update cluster o k response has a 3xx status code
func (o *UpdateClusterOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update cluster o k response has a 4xx status code
func (o *UpdateClusterOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update cluster o k response has a 5xx status code
func (o *UpdateClusterOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update cluster o k response a status code equal to that given
func (o *UpdateClusterOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the update cluster o k response
func (o *UpdateClusterOK) Code() int {
	return 200
}

func (o *UpdateClusterOK) Error() string {
	return fmt.Sprintf("[PATCH /v1/clusters/{id}][%d] updateClusterOK  %+v", 200, o.Payload)
}

func (o *UpdateClusterOK) String() string {
	return fmt.Sprintf("[PATCH /v1/clusters/{id}][%d] updateClusterOK  %+v", 200, o.Payload)
}

func (o *UpdateClusterOK) GetPayload() *models.Task {
	return o.Payload
}

func (o *UpdateClusterOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Task)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateClusterAccepted creates a UpdateClusterAccepted with default headers values
func NewUpdateClusterAccepted() *UpdateClusterAccepted {
	return &UpdateClusterAccepted{}
}

/*
UpdateClusterAccepted describes a response with status code 202, with default header values.

Accepted
*/
type UpdateClusterAccepted struct {
	Payload *models.Task
}

// IsSuccess returns true when this update cluster accepted response has a 2xx status code
func (o *UpdateClusterAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update cluster accepted response has a 3xx status code
func (o *UpdateClusterAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update cluster accepted response has a 4xx status code
func (o *UpdateClusterAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this update cluster accepted response has a 5xx status code
func (o *UpdateClusterAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this update cluster accepted response a status code equal to that given
func (o *UpdateClusterAccepted) IsCode(code int) bool {
	return code == 202
}

// Code gets the status code for the update cluster accepted response
func (o *UpdateClusterAccepted) Code() int {
	return 202
}

func (o *UpdateClusterAccepted) Error() string {
	return fmt.Sprintf("[PATCH /v1/clusters/{id}][%d] updateClusterAccepted  %+v", 202, o.Payload)
}

func (o *UpdateClusterAccepted) String() string {
	return fmt.Sprintf("[PATCH /v1/clusters/{id}][%d] updateClusterAccepted  %+v", 202, o.Payload)
}

func (o *UpdateClusterAccepted) GetPayload() *models.Task {
	return o.Payload
}

func (o *UpdateClusterAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Task)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateClusterBadRequest creates a UpdateClusterBadRequest with default headers values
func NewUpdateClusterBadRequest() *UpdateClusterBadRequest {
	return &UpdateClusterBadRequest{}
}

/*
UpdateClusterBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type UpdateClusterBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this update cluster bad request response has a 2xx status code
func (o *UpdateClusterBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update cluster bad request response has a 3xx status code
func (o *UpdateClusterBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update cluster bad request response has a 4xx status code
func (o *UpdateClusterBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this update cluster bad request response has a 5xx status code
func (o *UpdateClusterBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this update cluster bad request response a status code equal to that given
func (o *UpdateClusterBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the update cluster bad request response
func (o *UpdateClusterBadRequest) Code() int {
	return 400
}

func (o *UpdateClusterBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /v1/clusters/{id}][%d] updateClusterBadRequest  %+v", 400, o.Payload)
}

func (o *UpdateClusterBadRequest) String() string {
	return fmt.Sprintf("[PATCH /v1/clusters/{id}][%d] updateClusterBadRequest  %+v", 400, o.Payload)
}

func (o *UpdateClusterBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateClusterBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateClusterNotFound creates a UpdateClusterNotFound with default headers values
func NewUpdateClusterNotFound() *UpdateClusterNotFound {
	return &UpdateClusterNotFound{}
}

/*
UpdateClusterNotFound describes a response with status code 404, with default header values.

Not Found
*/
type UpdateClusterNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this update cluster not found response has a 2xx status code
func (o *UpdateClusterNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update cluster not found response has a 3xx status code
func (o *UpdateClusterNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update cluster not found response has a 4xx status code
func (o *UpdateClusterNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this update cluster not found response has a 5xx status code
func (o *UpdateClusterNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this update cluster not found response a status code equal to that given
func (o *UpdateClusterNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the update cluster not found response
func (o *UpdateClusterNotFound) Code() int {
	return 404
}

func (o *UpdateClusterNotFound) Error() string {
	return fmt.Sprintf("[PATCH /v1/clusters/{id}][%d] updateClusterNotFound  %+v", 404, o.Payload)
}

func (o *UpdateClusterNotFound) String() string {
	return fmt.Sprintf("[PATCH /v1/clusters/{id}][%d] updateClusterNotFound  %+v", 404, o.Payload)
}

func (o *UpdateClusterNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateClusterNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateClusterInternalServerError creates a UpdateClusterInternalServerError with default headers values
func NewUpdateClusterInternalServerError() *UpdateClusterInternalServerError {
	return &UpdateClusterInternalServerError{}
}

/*
UpdateClusterInternalServerError describes a response with status code 500, with default header values.

InternalServerError
*/
type UpdateClusterInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this update cluster internal server error response has a 2xx status code
func (o *UpdateClusterInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update cluster internal server error response has a 3xx status code
func (o *UpdateClusterInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update cluster internal server error response has a 4xx status code
func (o *UpdateClusterInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this update cluster internal server error response has a 5xx status code
func (o *UpdateClusterInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this update cluster internal server error response a status code equal to that given
func (o *UpdateClusterInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the update cluster internal server error response
func (o *UpdateClusterInternalServerError) Code() int {
	return 500
}

func (o *UpdateClusterInternalServerError) Error() string {
	return fmt.Sprintf("[PATCH /v1/clusters/{id}][%d] updateClusterInternalServerError  %+v", 500, o.Payload)
}

func (o *UpdateClusterInternalServerError) String() string {
	return fmt.Sprintf("[PATCH /v1/clusters/{id}][%d] updateClusterInternalServerError  %+v", 500, o.Payload)
}

func (o *UpdateClusterInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateClusterInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
