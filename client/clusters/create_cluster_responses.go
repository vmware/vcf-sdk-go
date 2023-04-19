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

// CreateClusterReader is a Reader for the CreateCluster structure.
type CreateClusterReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateClusterReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateClusterOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewCreateClusterAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateClusterBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCreateClusterInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateClusterOK creates a CreateClusterOK with default headers values
func NewCreateClusterOK() *CreateClusterOK {
	return &CreateClusterOK{}
}

/*
CreateClusterOK describes a response with status code 200, with default header values.

OK
*/
type CreateClusterOK struct {
	Payload *models.Task
}

// IsSuccess returns true when this create cluster o k response has a 2xx status code
func (o *CreateClusterOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create cluster o k response has a 3xx status code
func (o *CreateClusterOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create cluster o k response has a 4xx status code
func (o *CreateClusterOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this create cluster o k response has a 5xx status code
func (o *CreateClusterOK) IsServerError() bool {
	return false
}

// IsCode returns true when this create cluster o k response a status code equal to that given
func (o *CreateClusterOK) IsCode(code int) bool {
	return code == 200
}

func (o *CreateClusterOK) Error() string {
	return fmt.Sprintf("[POST /v1/clusters][%d] createClusterOK  %+v", 200, o.Payload)
}

func (o *CreateClusterOK) String() string {
	return fmt.Sprintf("[POST /v1/clusters][%d] createClusterOK  %+v", 200, o.Payload)
}

func (o *CreateClusterOK) GetPayload() *models.Task {
	return o.Payload
}

func (o *CreateClusterOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Task)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateClusterAccepted creates a CreateClusterAccepted with default headers values
func NewCreateClusterAccepted() *CreateClusterAccepted {
	return &CreateClusterAccepted{}
}

/*
CreateClusterAccepted describes a response with status code 202, with default header values.

Accepted
*/
type CreateClusterAccepted struct {
	Payload *models.Task
}

// IsSuccess returns true when this create cluster accepted response has a 2xx status code
func (o *CreateClusterAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create cluster accepted response has a 3xx status code
func (o *CreateClusterAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create cluster accepted response has a 4xx status code
func (o *CreateClusterAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this create cluster accepted response has a 5xx status code
func (o *CreateClusterAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this create cluster accepted response a status code equal to that given
func (o *CreateClusterAccepted) IsCode(code int) bool {
	return code == 202
}

func (o *CreateClusterAccepted) Error() string {
	return fmt.Sprintf("[POST /v1/clusters][%d] createClusterAccepted  %+v", 202, o.Payload)
}

func (o *CreateClusterAccepted) String() string {
	return fmt.Sprintf("[POST /v1/clusters][%d] createClusterAccepted  %+v", 202, o.Payload)
}

func (o *CreateClusterAccepted) GetPayload() *models.Task {
	return o.Payload
}

func (o *CreateClusterAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Task)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateClusterBadRequest creates a CreateClusterBadRequest with default headers values
func NewCreateClusterBadRequest() *CreateClusterBadRequest {
	return &CreateClusterBadRequest{}
}

/*
CreateClusterBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type CreateClusterBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this create cluster bad request response has a 2xx status code
func (o *CreateClusterBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create cluster bad request response has a 3xx status code
func (o *CreateClusterBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create cluster bad request response has a 4xx status code
func (o *CreateClusterBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this create cluster bad request response has a 5xx status code
func (o *CreateClusterBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this create cluster bad request response a status code equal to that given
func (o *CreateClusterBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *CreateClusterBadRequest) Error() string {
	return fmt.Sprintf("[POST /v1/clusters][%d] createClusterBadRequest  %+v", 400, o.Payload)
}

func (o *CreateClusterBadRequest) String() string {
	return fmt.Sprintf("[POST /v1/clusters][%d] createClusterBadRequest  %+v", 400, o.Payload)
}

func (o *CreateClusterBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateClusterBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateClusterInternalServerError creates a CreateClusterInternalServerError with default headers values
func NewCreateClusterInternalServerError() *CreateClusterInternalServerError {
	return &CreateClusterInternalServerError{}
}

/*
CreateClusterInternalServerError describes a response with status code 500, with default header values.

InternalServerError
*/
type CreateClusterInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this create cluster internal server error response has a 2xx status code
func (o *CreateClusterInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create cluster internal server error response has a 3xx status code
func (o *CreateClusterInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create cluster internal server error response has a 4xx status code
func (o *CreateClusterInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this create cluster internal server error response has a 5xx status code
func (o *CreateClusterInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this create cluster internal server error response a status code equal to that given
func (o *CreateClusterInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *CreateClusterInternalServerError) Error() string {
	return fmt.Sprintf("[POST /v1/clusters][%d] createClusterInternalServerError  %+v", 500, o.Payload)
}

func (o *CreateClusterInternalServerError) String() string {
	return fmt.Sprintf("[POST /v1/clusters][%d] createClusterInternalServerError  %+v", 500, o.Payload)
}

func (o *CreateClusterInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateClusterInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
