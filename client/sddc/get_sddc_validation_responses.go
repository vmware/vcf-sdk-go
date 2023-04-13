// Code generated by go-swagger; DO NOT EDIT.

// Copyright © 2023 VMware, Inc. All Rights Reserved.
// Code licenced under: BSD-2

package sddc

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"vcf-sdk-go/models"
)

// GETSDDCValidationReader is a Reader for the GETSDDCValidation structure.
type GETSDDCValidationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GETSDDCValidationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGETSDDCValidationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGETSDDCValidationBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGETSDDCValidationNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGETSDDCValidationInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 501:
		result := NewGETSDDCValidationNotImplemented()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGETSDDCValidationOK creates a GETSDDCValidationOK with default headers values
func NewGETSDDCValidationOK() *GETSDDCValidationOK {
	return &GETSDDCValidationOK{}
}

/*
GETSDDCValidationOK describes a response with status code 200, with default header values.

Accepted
*/
type GETSDDCValidationOK struct {
	Payload *models.Validation
}

// IsSuccess returns true when this get Sddc validation o k response has a 2xx status code
func (o *GETSDDCValidationOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get Sddc validation o k response has a 3xx status code
func (o *GETSDDCValidationOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get Sddc validation o k response has a 4xx status code
func (o *GETSDDCValidationOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get Sddc validation o k response has a 5xx status code
func (o *GETSDDCValidationOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get Sddc validation o k response a status code equal to that given
func (o *GETSDDCValidationOK) IsCode(code int) bool {
	return code == 200
}

func (o *GETSDDCValidationOK) Error() string {
	return fmt.Sprintf("[GET /v1/sddcs/validations/{id}][%d] getSddcValidationOK  %+v", 200, o.Payload)
}

func (o *GETSDDCValidationOK) String() string {
	return fmt.Sprintf("[GET /v1/sddcs/validations/{id}][%d] getSddcValidationOK  %+v", 200, o.Payload)
}

func (o *GETSDDCValidationOK) GetPayload() *models.Validation {
	return o.Payload
}

func (o *GETSDDCValidationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Validation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGETSDDCValidationBadRequest creates a GETSDDCValidationBadRequest with default headers values
func NewGETSDDCValidationBadRequest() *GETSDDCValidationBadRequest {
	return &GETSDDCValidationBadRequest{}
}

/*
GETSDDCValidationBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GETSDDCValidationBadRequest struct {
	Payload *models.Validation
}

// IsSuccess returns true when this get Sddc validation bad request response has a 2xx status code
func (o *GETSDDCValidationBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get Sddc validation bad request response has a 3xx status code
func (o *GETSDDCValidationBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get Sddc validation bad request response has a 4xx status code
func (o *GETSDDCValidationBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get Sddc validation bad request response has a 5xx status code
func (o *GETSDDCValidationBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get Sddc validation bad request response a status code equal to that given
func (o *GETSDDCValidationBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GETSDDCValidationBadRequest) Error() string {
	return fmt.Sprintf("[GET /v1/sddcs/validations/{id}][%d] getSddcValidationBadRequest  %+v", 400, o.Payload)
}

func (o *GETSDDCValidationBadRequest) String() string {
	return fmt.Sprintf("[GET /v1/sddcs/validations/{id}][%d] getSddcValidationBadRequest  %+v", 400, o.Payload)
}

func (o *GETSDDCValidationBadRequest) GetPayload() *models.Validation {
	return o.Payload
}

func (o *GETSDDCValidationBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Validation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGETSDDCValidationNotFound creates a GETSDDCValidationNotFound with default headers values
func NewGETSDDCValidationNotFound() *GETSDDCValidationNotFound {
	return &GETSDDCValidationNotFound{}
}

/*
GETSDDCValidationNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GETSDDCValidationNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this get Sddc validation not found response has a 2xx status code
func (o *GETSDDCValidationNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get Sddc validation not found response has a 3xx status code
func (o *GETSDDCValidationNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get Sddc validation not found response has a 4xx status code
func (o *GETSDDCValidationNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get Sddc validation not found response has a 5xx status code
func (o *GETSDDCValidationNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get Sddc validation not found response a status code equal to that given
func (o *GETSDDCValidationNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GETSDDCValidationNotFound) Error() string {
	return fmt.Sprintf("[GET /v1/sddcs/validations/{id}][%d] getSddcValidationNotFound  %+v", 404, o.Payload)
}

func (o *GETSDDCValidationNotFound) String() string {
	return fmt.Sprintf("[GET /v1/sddcs/validations/{id}][%d] getSddcValidationNotFound  %+v", 404, o.Payload)
}

func (o *GETSDDCValidationNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GETSDDCValidationNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGETSDDCValidationInternalServerError creates a GETSDDCValidationInternalServerError with default headers values
func NewGETSDDCValidationInternalServerError() *GETSDDCValidationInternalServerError {
	return &GETSDDCValidationInternalServerError{}
}

/*
GETSDDCValidationInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GETSDDCValidationInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this get Sddc validation internal server error response has a 2xx status code
func (o *GETSDDCValidationInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get Sddc validation internal server error response has a 3xx status code
func (o *GETSDDCValidationInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get Sddc validation internal server error response has a 4xx status code
func (o *GETSDDCValidationInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get Sddc validation internal server error response has a 5xx status code
func (o *GETSDDCValidationInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get Sddc validation internal server error response a status code equal to that given
func (o *GETSDDCValidationInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GETSDDCValidationInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/sddcs/validations/{id}][%d] getSddcValidationInternalServerError  %+v", 500, o.Payload)
}

func (o *GETSDDCValidationInternalServerError) String() string {
	return fmt.Sprintf("[GET /v1/sddcs/validations/{id}][%d] getSddcValidationInternalServerError  %+v", 500, o.Payload)
}

func (o *GETSDDCValidationInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *GETSDDCValidationInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGETSDDCValidationNotImplemented creates a GETSDDCValidationNotImplemented with default headers values
func NewGETSDDCValidationNotImplemented() *GETSDDCValidationNotImplemented {
	return &GETSDDCValidationNotImplemented{}
}

/*
GETSDDCValidationNotImplemented describes a response with status code 501, with default header values.

Not Implemented
*/
type GETSDDCValidationNotImplemented struct {
	Payload *models.Error
}

// IsSuccess returns true when this get Sddc validation not implemented response has a 2xx status code
func (o *GETSDDCValidationNotImplemented) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get Sddc validation not implemented response has a 3xx status code
func (o *GETSDDCValidationNotImplemented) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get Sddc validation not implemented response has a 4xx status code
func (o *GETSDDCValidationNotImplemented) IsClientError() bool {
	return false
}

// IsServerError returns true when this get Sddc validation not implemented response has a 5xx status code
func (o *GETSDDCValidationNotImplemented) IsServerError() bool {
	return true
}

// IsCode returns true when this get Sddc validation not implemented response a status code equal to that given
func (o *GETSDDCValidationNotImplemented) IsCode(code int) bool {
	return code == 501
}

func (o *GETSDDCValidationNotImplemented) Error() string {
	return fmt.Sprintf("[GET /v1/sddcs/validations/{id}][%d] getSddcValidationNotImplemented  %+v", 501, o.Payload)
}

func (o *GETSDDCValidationNotImplemented) String() string {
	return fmt.Sprintf("[GET /v1/sddcs/validations/{id}][%d] getSddcValidationNotImplemented  %+v", 501, o.Payload)
}

func (o *GETSDDCValidationNotImplemented) GetPayload() *models.Error {
	return o.Payload
}

func (o *GETSDDCValidationNotImplemented) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
