// Code generated by go-swagger; DO NOT EDIT.

// Copyright © 2023 VMware, Inc. All Rights Reserved.
// Code licenced under: BSD-2

package manifests

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/vmware/vcf-sdk-go/models"
)

// SaveManifestReader is a Reader for the SaveManifest structure.
type SaveManifestReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SaveManifestReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewSaveManifestAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewSaveManifestBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewSaveManifestInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewSaveManifestAccepted creates a SaveManifestAccepted with default headers values
func NewSaveManifestAccepted() *SaveManifestAccepted {
	return &SaveManifestAccepted{}
}

/*
SaveManifestAccepted describes a response with status code 202, with default header values.

Accepted
*/
type SaveManifestAccepted struct {
}

// IsSuccess returns true when this save manifest accepted response has a 2xx status code
func (o *SaveManifestAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this save manifest accepted response has a 3xx status code
func (o *SaveManifestAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this save manifest accepted response has a 4xx status code
func (o *SaveManifestAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this save manifest accepted response has a 5xx status code
func (o *SaveManifestAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this save manifest accepted response a status code equal to that given
func (o *SaveManifestAccepted) IsCode(code int) bool {
	return code == 202
}

func (o *SaveManifestAccepted) Error() string {
	return fmt.Sprintf("[POST /v1/manifests][%d] saveManifestAccepted ", 202)
}

func (o *SaveManifestAccepted) String() string {
	return fmt.Sprintf("[POST /v1/manifests][%d] saveManifestAccepted ", 202)
}

func (o *SaveManifestAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewSaveManifestBadRequest creates a SaveManifestBadRequest with default headers values
func NewSaveManifestBadRequest() *SaveManifestBadRequest {
	return &SaveManifestBadRequest{}
}

/*
SaveManifestBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type SaveManifestBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this save manifest bad request response has a 2xx status code
func (o *SaveManifestBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this save manifest bad request response has a 3xx status code
func (o *SaveManifestBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this save manifest bad request response has a 4xx status code
func (o *SaveManifestBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this save manifest bad request response has a 5xx status code
func (o *SaveManifestBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this save manifest bad request response a status code equal to that given
func (o *SaveManifestBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *SaveManifestBadRequest) Error() string {
	return fmt.Sprintf("[POST /v1/manifests][%d] saveManifestBadRequest  %+v", 400, o.Payload)
}

func (o *SaveManifestBadRequest) String() string {
	return fmt.Sprintf("[POST /v1/manifests][%d] saveManifestBadRequest  %+v", 400, o.Payload)
}

func (o *SaveManifestBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *SaveManifestBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSaveManifestInternalServerError creates a SaveManifestInternalServerError with default headers values
func NewSaveManifestInternalServerError() *SaveManifestInternalServerError {
	return &SaveManifestInternalServerError{}
}

/*
SaveManifestInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type SaveManifestInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this save manifest internal server error response has a 2xx status code
func (o *SaveManifestInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this save manifest internal server error response has a 3xx status code
func (o *SaveManifestInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this save manifest internal server error response has a 4xx status code
func (o *SaveManifestInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this save manifest internal server error response has a 5xx status code
func (o *SaveManifestInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this save manifest internal server error response a status code equal to that given
func (o *SaveManifestInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *SaveManifestInternalServerError) Error() string {
	return fmt.Sprintf("[POST /v1/manifests][%d] saveManifestInternalServerError  %+v", 500, o.Payload)
}

func (o *SaveManifestInternalServerError) String() string {
	return fmt.Sprintf("[POST /v1/manifests][%d] saveManifestInternalServerError  %+v", 500, o.Payload)
}

func (o *SaveManifestInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *SaveManifestInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
