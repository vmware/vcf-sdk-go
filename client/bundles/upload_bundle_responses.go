// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2023 VMware, Inc.
// SPDX-License-Identifier: BSD-2-Clause

package bundles

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/vmware/vcf-sdk-go/models"
)

// UploadBundleReader is a Reader for the UploadBundle structure.
type UploadBundleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UploadBundleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUploadBundleOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewUploadBundleAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUploadBundleBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUploadBundleInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUploadBundleOK creates a UploadBundleOK with default headers values
func NewUploadBundleOK() *UploadBundleOK {
	return &UploadBundleOK{}
}

/*
UploadBundleOK describes a response with status code 200, with default header values.

OK
*/
type UploadBundleOK struct {
	Payload *models.Task
}

// IsSuccess returns true when this upload bundle o k response has a 2xx status code
func (o *UploadBundleOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this upload bundle o k response has a 3xx status code
func (o *UploadBundleOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this upload bundle o k response has a 4xx status code
func (o *UploadBundleOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this upload bundle o k response has a 5xx status code
func (o *UploadBundleOK) IsServerError() bool {
	return false
}

// IsCode returns true when this upload bundle o k response a status code equal to that given
func (o *UploadBundleOK) IsCode(code int) bool {
	return code == 200
}

func (o *UploadBundleOK) Error() string {
	return fmt.Sprintf("[POST /v1/bundles][%d] uploadBundleOK  %+v", 200, o.Payload)
}

func (o *UploadBundleOK) String() string {
	return fmt.Sprintf("[POST /v1/bundles][%d] uploadBundleOK  %+v", 200, o.Payload)
}

func (o *UploadBundleOK) GetPayload() *models.Task {
	return o.Payload
}

func (o *UploadBundleOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Task)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUploadBundleAccepted creates a UploadBundleAccepted with default headers values
func NewUploadBundleAccepted() *UploadBundleAccepted {
	return &UploadBundleAccepted{}
}

/*
UploadBundleAccepted describes a response with status code 202, with default header values.

Accepted
*/
type UploadBundleAccepted struct {
	Payload *models.Task
}

// IsSuccess returns true when this upload bundle accepted response has a 2xx status code
func (o *UploadBundleAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this upload bundle accepted response has a 3xx status code
func (o *UploadBundleAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this upload bundle accepted response has a 4xx status code
func (o *UploadBundleAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this upload bundle accepted response has a 5xx status code
func (o *UploadBundleAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this upload bundle accepted response a status code equal to that given
func (o *UploadBundleAccepted) IsCode(code int) bool {
	return code == 202
}

func (o *UploadBundleAccepted) Error() string {
	return fmt.Sprintf("[POST /v1/bundles][%d] uploadBundleAccepted  %+v", 202, o.Payload)
}

func (o *UploadBundleAccepted) String() string {
	return fmt.Sprintf("[POST /v1/bundles][%d] uploadBundleAccepted  %+v", 202, o.Payload)
}

func (o *UploadBundleAccepted) GetPayload() *models.Task {
	return o.Payload
}

func (o *UploadBundleAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Task)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUploadBundleBadRequest creates a UploadBundleBadRequest with default headers values
func NewUploadBundleBadRequest() *UploadBundleBadRequest {
	return &UploadBundleBadRequest{}
}

/*
UploadBundleBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type UploadBundleBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this upload bundle bad request response has a 2xx status code
func (o *UploadBundleBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this upload bundle bad request response has a 3xx status code
func (o *UploadBundleBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this upload bundle bad request response has a 4xx status code
func (o *UploadBundleBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this upload bundle bad request response has a 5xx status code
func (o *UploadBundleBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this upload bundle bad request response a status code equal to that given
func (o *UploadBundleBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *UploadBundleBadRequest) Error() string {
	return fmt.Sprintf("[POST /v1/bundles][%d] uploadBundleBadRequest  %+v", 400, o.Payload)
}

func (o *UploadBundleBadRequest) String() string {
	return fmt.Sprintf("[POST /v1/bundles][%d] uploadBundleBadRequest  %+v", 400, o.Payload)
}

func (o *UploadBundleBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *UploadBundleBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUploadBundleInternalServerError creates a UploadBundleInternalServerError with default headers values
func NewUploadBundleInternalServerError() *UploadBundleInternalServerError {
	return &UploadBundleInternalServerError{}
}

/*
UploadBundleInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type UploadBundleInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this upload bundle internal server error response has a 2xx status code
func (o *UploadBundleInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this upload bundle internal server error response has a 3xx status code
func (o *UploadBundleInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this upload bundle internal server error response has a 4xx status code
func (o *UploadBundleInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this upload bundle internal server error response has a 5xx status code
func (o *UploadBundleInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this upload bundle internal server error response a status code equal to that given
func (o *UploadBundleInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *UploadBundleInternalServerError) Error() string {
	return fmt.Sprintf("[POST /v1/bundles][%d] uploadBundleInternalServerError  %+v", 500, o.Payload)
}

func (o *UploadBundleInternalServerError) String() string {
	return fmt.Sprintf("[POST /v1/bundles][%d] uploadBundleInternalServerError  %+v", 500, o.Payload)
}

func (o *UploadBundleInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *UploadBundleInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
