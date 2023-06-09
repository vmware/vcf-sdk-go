// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2023 VMware, Inc.
// SPDX-License-Identifier: BSD-2-Clause

package sddc

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/vmware/vcf-sdk-go/models"
)

// GetSDDCValidationReader is a Reader for the GetSDDCValidation structure.
type GetSDDCValidationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSDDCValidationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetSDDCValidationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetSDDCValidationBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetSDDCValidationNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetSDDCValidationInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 501:
		result := NewGetSDDCValidationNotImplemented()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetSDDCValidationOK creates a GetSDDCValidationOK with default headers values
func NewGetSDDCValidationOK() *GetSDDCValidationOK {
	return &GetSDDCValidationOK{}
}

/*
GetSDDCValidationOK describes a response with status code 200, with default header values.

Accepted
*/
type GetSDDCValidationOK struct {
	Payload *models.Validation
}

// IsSuccess returns true when this get Sddc validation o k response has a 2xx status code
func (o *GetSDDCValidationOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get Sddc validation o k response has a 3xx status code
func (o *GetSDDCValidationOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get Sddc validation o k response has a 4xx status code
func (o *GetSDDCValidationOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get Sddc validation o k response has a 5xx status code
func (o *GetSDDCValidationOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get Sddc validation o k response a status code equal to that given
func (o *GetSDDCValidationOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetSDDCValidationOK) Error() string {
	return fmt.Sprintf("[GET /v1/sddcs/validations/{id}][%d] getSddcValidationOK  %+v", 200, o.Payload)
}

func (o *GetSDDCValidationOK) String() string {
	return fmt.Sprintf("[GET /v1/sddcs/validations/{id}][%d] getSddcValidationOK  %+v", 200, o.Payload)
}

func (o *GetSDDCValidationOK) GetPayload() *models.Validation {
	return o.Payload
}

func (o *GetSDDCValidationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Validation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSDDCValidationBadRequest creates a GetSDDCValidationBadRequest with default headers values
func NewGetSDDCValidationBadRequest() *GetSDDCValidationBadRequest {
	return &GetSDDCValidationBadRequest{}
}

/*
GetSDDCValidationBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetSDDCValidationBadRequest struct {
	Payload *models.Validation
}

// IsSuccess returns true when this get Sddc validation bad request response has a 2xx status code
func (o *GetSDDCValidationBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get Sddc validation bad request response has a 3xx status code
func (o *GetSDDCValidationBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get Sddc validation bad request response has a 4xx status code
func (o *GetSDDCValidationBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get Sddc validation bad request response has a 5xx status code
func (o *GetSDDCValidationBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get Sddc validation bad request response a status code equal to that given
func (o *GetSDDCValidationBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetSDDCValidationBadRequest) Error() string {
	return fmt.Sprintf("[GET /v1/sddcs/validations/{id}][%d] getSddcValidationBadRequest  %+v", 400, o.Payload)
}

func (o *GetSDDCValidationBadRequest) String() string {
	return fmt.Sprintf("[GET /v1/sddcs/validations/{id}][%d] getSddcValidationBadRequest  %+v", 400, o.Payload)
}

func (o *GetSDDCValidationBadRequest) GetPayload() *models.Validation {
	return o.Payload
}

func (o *GetSDDCValidationBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Validation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSDDCValidationNotFound creates a GetSDDCValidationNotFound with default headers values
func NewGetSDDCValidationNotFound() *GetSDDCValidationNotFound {
	return &GetSDDCValidationNotFound{}
}

/*
GetSDDCValidationNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetSDDCValidationNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this get Sddc validation not found response has a 2xx status code
func (o *GetSDDCValidationNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get Sddc validation not found response has a 3xx status code
func (o *GetSDDCValidationNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get Sddc validation not found response has a 4xx status code
func (o *GetSDDCValidationNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get Sddc validation not found response has a 5xx status code
func (o *GetSDDCValidationNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get Sddc validation not found response a status code equal to that given
func (o *GetSDDCValidationNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetSDDCValidationNotFound) Error() string {
	return fmt.Sprintf("[GET /v1/sddcs/validations/{id}][%d] getSddcValidationNotFound  %+v", 404, o.Payload)
}

func (o *GetSDDCValidationNotFound) String() string {
	return fmt.Sprintf("[GET /v1/sddcs/validations/{id}][%d] getSddcValidationNotFound  %+v", 404, o.Payload)
}

func (o *GetSDDCValidationNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetSDDCValidationNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSDDCValidationInternalServerError creates a GetSDDCValidationInternalServerError with default headers values
func NewGetSDDCValidationInternalServerError() *GetSDDCValidationInternalServerError {
	return &GetSDDCValidationInternalServerError{}
}

/*
GetSDDCValidationInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetSDDCValidationInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this get Sddc validation internal server error response has a 2xx status code
func (o *GetSDDCValidationInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get Sddc validation internal server error response has a 3xx status code
func (o *GetSDDCValidationInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get Sddc validation internal server error response has a 4xx status code
func (o *GetSDDCValidationInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get Sddc validation internal server error response has a 5xx status code
func (o *GetSDDCValidationInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get Sddc validation internal server error response a status code equal to that given
func (o *GetSDDCValidationInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetSDDCValidationInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/sddcs/validations/{id}][%d] getSddcValidationInternalServerError  %+v", 500, o.Payload)
}

func (o *GetSDDCValidationInternalServerError) String() string {
	return fmt.Sprintf("[GET /v1/sddcs/validations/{id}][%d] getSddcValidationInternalServerError  %+v", 500, o.Payload)
}

func (o *GetSDDCValidationInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetSDDCValidationInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSDDCValidationNotImplemented creates a GetSDDCValidationNotImplemented with default headers values
func NewGetSDDCValidationNotImplemented() *GetSDDCValidationNotImplemented {
	return &GetSDDCValidationNotImplemented{}
}

/*
GetSDDCValidationNotImplemented describes a response with status code 501, with default header values.

Not Implemented
*/
type GetSDDCValidationNotImplemented struct {
	Payload *models.Error
}

// IsSuccess returns true when this get Sddc validation not implemented response has a 2xx status code
func (o *GetSDDCValidationNotImplemented) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get Sddc validation not implemented response has a 3xx status code
func (o *GetSDDCValidationNotImplemented) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get Sddc validation not implemented response has a 4xx status code
func (o *GetSDDCValidationNotImplemented) IsClientError() bool {
	return false
}

// IsServerError returns true when this get Sddc validation not implemented response has a 5xx status code
func (o *GetSDDCValidationNotImplemented) IsServerError() bool {
	return true
}

// IsCode returns true when this get Sddc validation not implemented response a status code equal to that given
func (o *GetSDDCValidationNotImplemented) IsCode(code int) bool {
	return code == 501
}

func (o *GetSDDCValidationNotImplemented) Error() string {
	return fmt.Sprintf("[GET /v1/sddcs/validations/{id}][%d] getSddcValidationNotImplemented  %+v", 501, o.Payload)
}

func (o *GetSDDCValidationNotImplemented) String() string {
	return fmt.Sprintf("[GET /v1/sddcs/validations/{id}][%d] getSddcValidationNotImplemented  %+v", 501, o.Payload)
}

func (o *GetSDDCValidationNotImplemented) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetSDDCValidationNotImplemented) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
