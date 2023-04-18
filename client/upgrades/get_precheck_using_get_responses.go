// Code generated by go-swagger; DO NOT EDIT.

// Copyright © 2023 VMware, Inc. All Rights Reserved.
// Code licenced under: BSD-2

package upgrades

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/vmware/vcf-sdk-go/models"
)

// GETPrecheckUsingGETReader is a Reader for the GETPrecheckUsingGET structure.
type GETPrecheckUsingGETReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GETPrecheckUsingGETReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGETPrecheckUsingGETOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGETPrecheckUsingGETBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGETPrecheckUsingGETForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGETPrecheckUsingGETInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGETPrecheckUsingGETOK creates a GETPrecheckUsingGETOK with default headers values
func NewGETPrecheckUsingGETOK() *GETPrecheckUsingGETOK {
	return &GETPrecheckUsingGETOK{}
}

/*
GETPrecheckUsingGETOK describes a response with status code 200, with default header values.

Ok
*/
type GETPrecheckUsingGETOK struct {
	Payload *models.Task
}

// IsSuccess returns true when this get precheck using Get o k response has a 2xx status code
func (o *GETPrecheckUsingGETOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get precheck using Get o k response has a 3xx status code
func (o *GETPrecheckUsingGETOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get precheck using Get o k response has a 4xx status code
func (o *GETPrecheckUsingGETOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get precheck using Get o k response has a 5xx status code
func (o *GETPrecheckUsingGETOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get precheck using Get o k response a status code equal to that given
func (o *GETPrecheckUsingGETOK) IsCode(code int) bool {
	return code == 200
}

func (o *GETPrecheckUsingGETOK) Error() string {
	return fmt.Sprintf("[GET /v1/upgrades/{upgradeId}/prechecks/{precheckId}][%d] getPrecheckUsingGetOK  %+v", 200, o.Payload)
}

func (o *GETPrecheckUsingGETOK) String() string {
	return fmt.Sprintf("[GET /v1/upgrades/{upgradeId}/prechecks/{precheckId}][%d] getPrecheckUsingGetOK  %+v", 200, o.Payload)
}

func (o *GETPrecheckUsingGETOK) GetPayload() *models.Task {
	return o.Payload
}

func (o *GETPrecheckUsingGETOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Task)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGETPrecheckUsingGETBadRequest creates a GETPrecheckUsingGETBadRequest with default headers values
func NewGETPrecheckUsingGETBadRequest() *GETPrecheckUsingGETBadRequest {
	return &GETPrecheckUsingGETBadRequest{}
}

/*
GETPrecheckUsingGETBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GETPrecheckUsingGETBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this get precheck using Get bad request response has a 2xx status code
func (o *GETPrecheckUsingGETBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get precheck using Get bad request response has a 3xx status code
func (o *GETPrecheckUsingGETBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get precheck using Get bad request response has a 4xx status code
func (o *GETPrecheckUsingGETBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get precheck using Get bad request response has a 5xx status code
func (o *GETPrecheckUsingGETBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get precheck using Get bad request response a status code equal to that given
func (o *GETPrecheckUsingGETBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GETPrecheckUsingGETBadRequest) Error() string {
	return fmt.Sprintf("[GET /v1/upgrades/{upgradeId}/prechecks/{precheckId}][%d] getPrecheckUsingGetBadRequest  %+v", 400, o.Payload)
}

func (o *GETPrecheckUsingGETBadRequest) String() string {
	return fmt.Sprintf("[GET /v1/upgrades/{upgradeId}/prechecks/{precheckId}][%d] getPrecheckUsingGetBadRequest  %+v", 400, o.Payload)
}

func (o *GETPrecheckUsingGETBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *GETPrecheckUsingGETBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGETPrecheckUsingGETForbidden creates a GETPrecheckUsingGETForbidden with default headers values
func NewGETPrecheckUsingGETForbidden() *GETPrecheckUsingGETForbidden {
	return &GETPrecheckUsingGETForbidden{}
}

/*
GETPrecheckUsingGETForbidden describes a response with status code 403, with default header values.

Operation not allowed
*/
type GETPrecheckUsingGETForbidden struct {
	Payload *models.Error
}

// IsSuccess returns true when this get precheck using Get forbidden response has a 2xx status code
func (o *GETPrecheckUsingGETForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get precheck using Get forbidden response has a 3xx status code
func (o *GETPrecheckUsingGETForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get precheck using Get forbidden response has a 4xx status code
func (o *GETPrecheckUsingGETForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get precheck using Get forbidden response has a 5xx status code
func (o *GETPrecheckUsingGETForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get precheck using Get forbidden response a status code equal to that given
func (o *GETPrecheckUsingGETForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GETPrecheckUsingGETForbidden) Error() string {
	return fmt.Sprintf("[GET /v1/upgrades/{upgradeId}/prechecks/{precheckId}][%d] getPrecheckUsingGetForbidden  %+v", 403, o.Payload)
}

func (o *GETPrecheckUsingGETForbidden) String() string {
	return fmt.Sprintf("[GET /v1/upgrades/{upgradeId}/prechecks/{precheckId}][%d] getPrecheckUsingGetForbidden  %+v", 403, o.Payload)
}

func (o *GETPrecheckUsingGETForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *GETPrecheckUsingGETForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGETPrecheckUsingGETInternalServerError creates a GETPrecheckUsingGETInternalServerError with default headers values
func NewGETPrecheckUsingGETInternalServerError() *GETPrecheckUsingGETInternalServerError {
	return &GETPrecheckUsingGETInternalServerError{}
}

/*
GETPrecheckUsingGETInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GETPrecheckUsingGETInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this get precheck using Get internal server error response has a 2xx status code
func (o *GETPrecheckUsingGETInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get precheck using Get internal server error response has a 3xx status code
func (o *GETPrecheckUsingGETInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get precheck using Get internal server error response has a 4xx status code
func (o *GETPrecheckUsingGETInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get precheck using Get internal server error response has a 5xx status code
func (o *GETPrecheckUsingGETInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get precheck using Get internal server error response a status code equal to that given
func (o *GETPrecheckUsingGETInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GETPrecheckUsingGETInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/upgrades/{upgradeId}/prechecks/{precheckId}][%d] getPrecheckUsingGetInternalServerError  %+v", 500, o.Payload)
}

func (o *GETPrecheckUsingGETInternalServerError) String() string {
	return fmt.Sprintf("[GET /v1/upgrades/{upgradeId}/prechecks/{precheckId}][%d] getPrecheckUsingGetInternalServerError  %+v", 500, o.Payload)
}

func (o *GETPrecheckUsingGETInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *GETPrecheckUsingGETInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
