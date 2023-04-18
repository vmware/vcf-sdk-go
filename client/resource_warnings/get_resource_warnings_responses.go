// Code generated by go-swagger; DO NOT EDIT.

// Copyright © 2023 VMware, Inc. All Rights Reserved.
// Code licenced under: BSD-2

package resource_warnings

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/vmware/vcf-sdk-go/models"
)

// GETResourceWarningsReader is a Reader for the GETResourceWarnings structure.
type GETResourceWarningsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GETResourceWarningsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGETResourceWarningsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGETResourceWarningsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGETResourceWarningsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGETResourceWarningsOK creates a GETResourceWarningsOK with default headers values
func NewGETResourceWarningsOK() *GETResourceWarningsOK {
	return &GETResourceWarningsOK{}
}

/*
GETResourceWarningsOK describes a response with status code 200, with default header values.

OK
*/
type GETResourceWarningsOK struct {
	Payload *models.PageOfResourceWarning
}

// IsSuccess returns true when this get resource warnings o k response has a 2xx status code
func (o *GETResourceWarningsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get resource warnings o k response has a 3xx status code
func (o *GETResourceWarningsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get resource warnings o k response has a 4xx status code
func (o *GETResourceWarningsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get resource warnings o k response has a 5xx status code
func (o *GETResourceWarningsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get resource warnings o k response a status code equal to that given
func (o *GETResourceWarningsOK) IsCode(code int) bool {
	return code == 200
}

func (o *GETResourceWarningsOK) Error() string {
	return fmt.Sprintf("[GET /v1/resource-warnings][%d] getResourceWarningsOK  %+v", 200, o.Payload)
}

func (o *GETResourceWarningsOK) String() string {
	return fmt.Sprintf("[GET /v1/resource-warnings][%d] getResourceWarningsOK  %+v", 200, o.Payload)
}

func (o *GETResourceWarningsOK) GetPayload() *models.PageOfResourceWarning {
	return o.Payload
}

func (o *GETResourceWarningsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PageOfResourceWarning)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGETResourceWarningsBadRequest creates a GETResourceWarningsBadRequest with default headers values
func NewGETResourceWarningsBadRequest() *GETResourceWarningsBadRequest {
	return &GETResourceWarningsBadRequest{}
}

/*
GETResourceWarningsBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GETResourceWarningsBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this get resource warnings bad request response has a 2xx status code
func (o *GETResourceWarningsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get resource warnings bad request response has a 3xx status code
func (o *GETResourceWarningsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get resource warnings bad request response has a 4xx status code
func (o *GETResourceWarningsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get resource warnings bad request response has a 5xx status code
func (o *GETResourceWarningsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get resource warnings bad request response a status code equal to that given
func (o *GETResourceWarningsBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GETResourceWarningsBadRequest) Error() string {
	return fmt.Sprintf("[GET /v1/resource-warnings][%d] getResourceWarningsBadRequest  %+v", 400, o.Payload)
}

func (o *GETResourceWarningsBadRequest) String() string {
	return fmt.Sprintf("[GET /v1/resource-warnings][%d] getResourceWarningsBadRequest  %+v", 400, o.Payload)
}

func (o *GETResourceWarningsBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *GETResourceWarningsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGETResourceWarningsInternalServerError creates a GETResourceWarningsInternalServerError with default headers values
func NewGETResourceWarningsInternalServerError() *GETResourceWarningsInternalServerError {
	return &GETResourceWarningsInternalServerError{}
}

/*
GETResourceWarningsInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GETResourceWarningsInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this get resource warnings internal server error response has a 2xx status code
func (o *GETResourceWarningsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get resource warnings internal server error response has a 3xx status code
func (o *GETResourceWarningsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get resource warnings internal server error response has a 4xx status code
func (o *GETResourceWarningsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get resource warnings internal server error response has a 5xx status code
func (o *GETResourceWarningsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get resource warnings internal server error response a status code equal to that given
func (o *GETResourceWarningsInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GETResourceWarningsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/resource-warnings][%d] getResourceWarningsInternalServerError  %+v", 500, o.Payload)
}

func (o *GETResourceWarningsInternalServerError) String() string {
	return fmt.Sprintf("[GET /v1/resource-warnings][%d] getResourceWarningsInternalServerError  %+v", 500, o.Payload)
}

func (o *GETResourceWarningsInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *GETResourceWarningsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
