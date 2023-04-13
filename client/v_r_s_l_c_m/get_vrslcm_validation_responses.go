// Code generated by go-swagger; DO NOT EDIT.

// Copyright © 2023 VMware, Inc. All Rights Reserved.
// Code licenced under: BSD-2

package v_r_s_l_c_m

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"vcf-sdk-go/models"
)

// GETVRSLCMValidationReader is a Reader for the GETVRSLCMValidation structure.
type GETVRSLCMValidationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GETVRSLCMValidationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGETVRSLCMValidationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGETVRSLCMValidationBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGETVRSLCMValidationNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGETVRSLCMValidationOK creates a GETVRSLCMValidationOK with default headers values
func NewGETVRSLCMValidationOK() *GETVRSLCMValidationOK {
	return &GETVRSLCMValidationOK{}
}

/*
GETVRSLCMValidationOK describes a response with status code 200, with default header values.

OK
*/
type GETVRSLCMValidationOK struct {
	Payload *models.Validation
}

// IsSuccess returns true when this get Vrslcm validation o k response has a 2xx status code
func (o *GETVRSLCMValidationOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get Vrslcm validation o k response has a 3xx status code
func (o *GETVRSLCMValidationOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get Vrslcm validation o k response has a 4xx status code
func (o *GETVRSLCMValidationOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get Vrslcm validation o k response has a 5xx status code
func (o *GETVRSLCMValidationOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get Vrslcm validation o k response a status code equal to that given
func (o *GETVRSLCMValidationOK) IsCode(code int) bool {
	return code == 200
}

func (o *GETVRSLCMValidationOK) Error() string {
	return fmt.Sprintf("[GET /v1/vrslcms/validations/{id}][%d] getVrslcmValidationOK  %+v", 200, o.Payload)
}

func (o *GETVRSLCMValidationOK) String() string {
	return fmt.Sprintf("[GET /v1/vrslcms/validations/{id}][%d] getVrslcmValidationOK  %+v", 200, o.Payload)
}

func (o *GETVRSLCMValidationOK) GetPayload() *models.Validation {
	return o.Payload
}

func (o *GETVRSLCMValidationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Validation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGETVRSLCMValidationBadRequest creates a GETVRSLCMValidationBadRequest with default headers values
func NewGETVRSLCMValidationBadRequest() *GETVRSLCMValidationBadRequest {
	return &GETVRSLCMValidationBadRequest{}
}

/*
GETVRSLCMValidationBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GETVRSLCMValidationBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this get Vrslcm validation bad request response has a 2xx status code
func (o *GETVRSLCMValidationBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get Vrslcm validation bad request response has a 3xx status code
func (o *GETVRSLCMValidationBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get Vrslcm validation bad request response has a 4xx status code
func (o *GETVRSLCMValidationBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get Vrslcm validation bad request response has a 5xx status code
func (o *GETVRSLCMValidationBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get Vrslcm validation bad request response a status code equal to that given
func (o *GETVRSLCMValidationBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GETVRSLCMValidationBadRequest) Error() string {
	return fmt.Sprintf("[GET /v1/vrslcms/validations/{id}][%d] getVrslcmValidationBadRequest  %+v", 400, o.Payload)
}

func (o *GETVRSLCMValidationBadRequest) String() string {
	return fmt.Sprintf("[GET /v1/vrslcms/validations/{id}][%d] getVrslcmValidationBadRequest  %+v", 400, o.Payload)
}

func (o *GETVRSLCMValidationBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *GETVRSLCMValidationBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGETVRSLCMValidationNotFound creates a GETVRSLCMValidationNotFound with default headers values
func NewGETVRSLCMValidationNotFound() *GETVRSLCMValidationNotFound {
	return &GETVRSLCMValidationNotFound{}
}

/*
GETVRSLCMValidationNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GETVRSLCMValidationNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this get Vrslcm validation not found response has a 2xx status code
func (o *GETVRSLCMValidationNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get Vrslcm validation not found response has a 3xx status code
func (o *GETVRSLCMValidationNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get Vrslcm validation not found response has a 4xx status code
func (o *GETVRSLCMValidationNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get Vrslcm validation not found response has a 5xx status code
func (o *GETVRSLCMValidationNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get Vrslcm validation not found response a status code equal to that given
func (o *GETVRSLCMValidationNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GETVRSLCMValidationNotFound) Error() string {
	return fmt.Sprintf("[GET /v1/vrslcms/validations/{id}][%d] getVrslcmValidationNotFound  %+v", 404, o.Payload)
}

func (o *GETVRSLCMValidationNotFound) String() string {
	return fmt.Sprintf("[GET /v1/vrslcms/validations/{id}][%d] getVrslcmValidationNotFound  %+v", 404, o.Payload)
}

func (o *GETVRSLCMValidationNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GETVRSLCMValidationNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
