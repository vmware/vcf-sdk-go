// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2023 VMware, Inc.
// SPDX-License-Identifier: BSD-2-Clause

package personalities

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/vmware/vcf-sdk-go/models"
)

// UploadPersonalityReader is a Reader for the UploadPersonality structure.
type UploadPersonalityReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UploadPersonalityReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUploadPersonalityOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewUploadPersonalityAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUploadPersonalityBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUploadPersonalityInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /v1/personalities] uploadPersonality", response, response.Code())
	}
}

// NewUploadPersonalityOK creates a UploadPersonalityOK with default headers values
func NewUploadPersonalityOK() *UploadPersonalityOK {
	return &UploadPersonalityOK{}
}

/*
UploadPersonalityOK describes a response with status code 200, with default header values.

OK
*/
type UploadPersonalityOK struct {
	Payload *models.Task
}

// IsSuccess returns true when this upload personality o k response has a 2xx status code
func (o *UploadPersonalityOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this upload personality o k response has a 3xx status code
func (o *UploadPersonalityOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this upload personality o k response has a 4xx status code
func (o *UploadPersonalityOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this upload personality o k response has a 5xx status code
func (o *UploadPersonalityOK) IsServerError() bool {
	return false
}

// IsCode returns true when this upload personality o k response a status code equal to that given
func (o *UploadPersonalityOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the upload personality o k response
func (o *UploadPersonalityOK) Code() int {
	return 200
}

func (o *UploadPersonalityOK) Error() string {
	return fmt.Sprintf("[POST /v1/personalities][%d] uploadPersonalityOK  %+v", 200, o.Payload)
}

func (o *UploadPersonalityOK) String() string {
	return fmt.Sprintf("[POST /v1/personalities][%d] uploadPersonalityOK  %+v", 200, o.Payload)
}

func (o *UploadPersonalityOK) GetPayload() *models.Task {
	return o.Payload
}

func (o *UploadPersonalityOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Task)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUploadPersonalityAccepted creates a UploadPersonalityAccepted with default headers values
func NewUploadPersonalityAccepted() *UploadPersonalityAccepted {
	return &UploadPersonalityAccepted{}
}

/*
UploadPersonalityAccepted describes a response with status code 202, with default header values.

Accepted
*/
type UploadPersonalityAccepted struct {
	Payload *models.Task
}

// IsSuccess returns true when this upload personality accepted response has a 2xx status code
func (o *UploadPersonalityAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this upload personality accepted response has a 3xx status code
func (o *UploadPersonalityAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this upload personality accepted response has a 4xx status code
func (o *UploadPersonalityAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this upload personality accepted response has a 5xx status code
func (o *UploadPersonalityAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this upload personality accepted response a status code equal to that given
func (o *UploadPersonalityAccepted) IsCode(code int) bool {
	return code == 202
}

// Code gets the status code for the upload personality accepted response
func (o *UploadPersonalityAccepted) Code() int {
	return 202
}

func (o *UploadPersonalityAccepted) Error() string {
	return fmt.Sprintf("[POST /v1/personalities][%d] uploadPersonalityAccepted  %+v", 202, o.Payload)
}

func (o *UploadPersonalityAccepted) String() string {
	return fmt.Sprintf("[POST /v1/personalities][%d] uploadPersonalityAccepted  %+v", 202, o.Payload)
}

func (o *UploadPersonalityAccepted) GetPayload() *models.Task {
	return o.Payload
}

func (o *UploadPersonalityAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Task)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUploadPersonalityBadRequest creates a UploadPersonalityBadRequest with default headers values
func NewUploadPersonalityBadRequest() *UploadPersonalityBadRequest {
	return &UploadPersonalityBadRequest{}
}

/*
UploadPersonalityBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type UploadPersonalityBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this upload personality bad request response has a 2xx status code
func (o *UploadPersonalityBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this upload personality bad request response has a 3xx status code
func (o *UploadPersonalityBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this upload personality bad request response has a 4xx status code
func (o *UploadPersonalityBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this upload personality bad request response has a 5xx status code
func (o *UploadPersonalityBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this upload personality bad request response a status code equal to that given
func (o *UploadPersonalityBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the upload personality bad request response
func (o *UploadPersonalityBadRequest) Code() int {
	return 400
}

func (o *UploadPersonalityBadRequest) Error() string {
	return fmt.Sprintf("[POST /v1/personalities][%d] uploadPersonalityBadRequest  %+v", 400, o.Payload)
}

func (o *UploadPersonalityBadRequest) String() string {
	return fmt.Sprintf("[POST /v1/personalities][%d] uploadPersonalityBadRequest  %+v", 400, o.Payload)
}

func (o *UploadPersonalityBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *UploadPersonalityBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUploadPersonalityInternalServerError creates a UploadPersonalityInternalServerError with default headers values
func NewUploadPersonalityInternalServerError() *UploadPersonalityInternalServerError {
	return &UploadPersonalityInternalServerError{}
}

/*
UploadPersonalityInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type UploadPersonalityInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this upload personality internal server error response has a 2xx status code
func (o *UploadPersonalityInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this upload personality internal server error response has a 3xx status code
func (o *UploadPersonalityInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this upload personality internal server error response has a 4xx status code
func (o *UploadPersonalityInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this upload personality internal server error response has a 5xx status code
func (o *UploadPersonalityInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this upload personality internal server error response a status code equal to that given
func (o *UploadPersonalityInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the upload personality internal server error response
func (o *UploadPersonalityInternalServerError) Code() int {
	return 500
}

func (o *UploadPersonalityInternalServerError) Error() string {
	return fmt.Sprintf("[POST /v1/personalities][%d] uploadPersonalityInternalServerError  %+v", 500, o.Payload)
}

func (o *UploadPersonalityInternalServerError) String() string {
	return fmt.Sprintf("[POST /v1/personalities][%d] uploadPersonalityInternalServerError  %+v", 500, o.Payload)
}

func (o *UploadPersonalityInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *UploadPersonalityInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
