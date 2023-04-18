// Code generated by go-swagger; DO NOT EDIT.

// Copyright © 2023 VMware, Inc. All Rights Reserved.
// Code licenced under: BSD-2

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

// GetsupportBundleReader is a Reader for the GetsupportBundle structure.
type GetsupportBundleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetsupportBundleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetsupportBundleOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetsupportBundleBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetsupportBundleUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetsupportBundleInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetsupportBundleOK creates a GetsupportBundleOK with default headers values
func NewGetsupportBundleOK() *GetsupportBundleOK {
	return &GetsupportBundleOK{}
}

/*
GetsupportBundleOK describes a response with status code 200, with default header values.

Ok
*/
type GetsupportBundleOK struct {
	Payload *models.SupportBundle
}

// IsSuccess returns true when this getsupport bundle o k response has a 2xx status code
func (o *GetsupportBundleOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this getsupport bundle o k response has a 3xx status code
func (o *GetsupportBundleOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this getsupport bundle o k response has a 4xx status code
func (o *GetsupportBundleOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this getsupport bundle o k response has a 5xx status code
func (o *GetsupportBundleOK) IsServerError() bool {
	return false
}

// IsCode returns true when this getsupport bundle o k response a status code equal to that given
func (o *GetsupportBundleOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetsupportBundleOK) Error() string {
	return fmt.Sprintf("[GET /v1/system/support-bundles/{id}][%d] getsupportBundleOK  %+v", 200, o.Payload)
}

func (o *GetsupportBundleOK) String() string {
	return fmt.Sprintf("[GET /v1/system/support-bundles/{id}][%d] getsupportBundleOK  %+v", 200, o.Payload)
}

func (o *GetsupportBundleOK) GetPayload() *models.SupportBundle {
	return o.Payload
}

func (o *GetsupportBundleOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SupportBundle)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetsupportBundleBadRequest creates a GetsupportBundleBadRequest with default headers values
func NewGetsupportBundleBadRequest() *GetsupportBundleBadRequest {
	return &GetsupportBundleBadRequest{}
}

/*
GetsupportBundleBadRequest describes a response with status code 400, with default header values.

Bad request! Invalid Headers or Data. Error: {error}
*/
type GetsupportBundleBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this getsupport bundle bad request response has a 2xx status code
func (o *GetsupportBundleBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this getsupport bundle bad request response has a 3xx status code
func (o *GetsupportBundleBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this getsupport bundle bad request response has a 4xx status code
func (o *GetsupportBundleBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this getsupport bundle bad request response has a 5xx status code
func (o *GetsupportBundleBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this getsupport bundle bad request response a status code equal to that given
func (o *GetsupportBundleBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetsupportBundleBadRequest) Error() string {
	return fmt.Sprintf("[GET /v1/system/support-bundles/{id}][%d] getsupportBundleBadRequest  %+v", 400, o.Payload)
}

func (o *GetsupportBundleBadRequest) String() string {
	return fmt.Sprintf("[GET /v1/system/support-bundles/{id}][%d] getsupportBundleBadRequest  %+v", 400, o.Payload)
}

func (o *GetsupportBundleBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetsupportBundleBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetsupportBundleUnauthorized creates a GetsupportBundleUnauthorized with default headers values
func NewGetsupportBundleUnauthorized() *GetsupportBundleUnauthorized {
	return &GetsupportBundleUnauthorized{}
}

/*
GetsupportBundleUnauthorized describes a response with status code 401, with default header values.

Bad request! Authorization Header is missing or not in correct format.
*/
type GetsupportBundleUnauthorized struct {
	Payload *models.Error
}

// IsSuccess returns true when this getsupport bundle unauthorized response has a 2xx status code
func (o *GetsupportBundleUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this getsupport bundle unauthorized response has a 3xx status code
func (o *GetsupportBundleUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this getsupport bundle unauthorized response has a 4xx status code
func (o *GetsupportBundleUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this getsupport bundle unauthorized response has a 5xx status code
func (o *GetsupportBundleUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this getsupport bundle unauthorized response a status code equal to that given
func (o *GetsupportBundleUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetsupportBundleUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v1/system/support-bundles/{id}][%d] getsupportBundleUnauthorized  %+v", 401, o.Payload)
}

func (o *GetsupportBundleUnauthorized) String() string {
	return fmt.Sprintf("[GET /v1/system/support-bundles/{id}][%d] getsupportBundleUnauthorized  %+v", 401, o.Payload)
}

func (o *GetsupportBundleUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetsupportBundleUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetsupportBundleInternalServerError creates a GetsupportBundleInternalServerError with default headers values
func NewGetsupportBundleInternalServerError() *GetsupportBundleInternalServerError {
	return &GetsupportBundleInternalServerError{}
}

/*
GetsupportBundleInternalServerError describes a response with status code 500, with default header values.

Something went wrong. Internal server error occurred. Error {error}
*/
type GetsupportBundleInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this getsupport bundle internal server error response has a 2xx status code
func (o *GetsupportBundleInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this getsupport bundle internal server error response has a 3xx status code
func (o *GetsupportBundleInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this getsupport bundle internal server error response has a 4xx status code
func (o *GetsupportBundleInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this getsupport bundle internal server error response has a 5xx status code
func (o *GetsupportBundleInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this getsupport bundle internal server error response a status code equal to that given
func (o *GetsupportBundleInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetsupportBundleInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/system/support-bundles/{id}][%d] getsupportBundleInternalServerError  %+v", 500, o.Payload)
}

func (o *GetsupportBundleInternalServerError) String() string {
	return fmt.Sprintf("[GET /v1/system/support-bundles/{id}][%d] getsupportBundleInternalServerError  %+v", 500, o.Payload)
}

func (o *GetsupportBundleInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetsupportBundleInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
