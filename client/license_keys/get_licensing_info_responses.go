// Code generated by go-swagger; DO NOT EDIT.

// Copyright © 2023 VMware, Inc. All Rights Reserved.
// Code licenced under: BSD-2

package license_keys

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"vcf-sdk-go/models"
)

// GETLicensingInfoReader is a Reader for the GETLicensingInfo structure.
type GETLicensingInfoReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GETLicensingInfoReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGETLicensingInfoOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGETLicensingInfoBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGETLicensingInfoInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGETLicensingInfoOK creates a GETLicensingInfoOK with default headers values
func NewGETLicensingInfoOK() *GETLicensingInfoOK {
	return &GETLicensingInfoOK{}
}

/*
GETLicensingInfoOK describes a response with status code 200, with default header values.

Successful
*/
type GETLicensingInfoOK struct {
	Payload []*models.LicensingInfo
}

// IsSuccess returns true when this get licensing info o k response has a 2xx status code
func (o *GETLicensingInfoOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get licensing info o k response has a 3xx status code
func (o *GETLicensingInfoOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get licensing info o k response has a 4xx status code
func (o *GETLicensingInfoOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get licensing info o k response has a 5xx status code
func (o *GETLicensingInfoOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get licensing info o k response a status code equal to that given
func (o *GETLicensingInfoOK) IsCode(code int) bool {
	return code == 200
}

func (o *GETLicensingInfoOK) Error() string {
	return fmt.Sprintf("[GET /v1/licensing-info][%d] getLicensingInfoOK  %+v", 200, o.Payload)
}

func (o *GETLicensingInfoOK) String() string {
	return fmt.Sprintf("[GET /v1/licensing-info][%d] getLicensingInfoOK  %+v", 200, o.Payload)
}

func (o *GETLicensingInfoOK) GetPayload() []*models.LicensingInfo {
	return o.Payload
}

func (o *GETLicensingInfoOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGETLicensingInfoBadRequest creates a GETLicensingInfoBadRequest with default headers values
func NewGETLicensingInfoBadRequest() *GETLicensingInfoBadRequest {
	return &GETLicensingInfoBadRequest{}
}

/*
GETLicensingInfoBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GETLicensingInfoBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this get licensing info bad request response has a 2xx status code
func (o *GETLicensingInfoBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get licensing info bad request response has a 3xx status code
func (o *GETLicensingInfoBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get licensing info bad request response has a 4xx status code
func (o *GETLicensingInfoBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get licensing info bad request response has a 5xx status code
func (o *GETLicensingInfoBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get licensing info bad request response a status code equal to that given
func (o *GETLicensingInfoBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GETLicensingInfoBadRequest) Error() string {
	return fmt.Sprintf("[GET /v1/licensing-info][%d] getLicensingInfoBadRequest  %+v", 400, o.Payload)
}

func (o *GETLicensingInfoBadRequest) String() string {
	return fmt.Sprintf("[GET /v1/licensing-info][%d] getLicensingInfoBadRequest  %+v", 400, o.Payload)
}

func (o *GETLicensingInfoBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *GETLicensingInfoBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGETLicensingInfoInternalServerError creates a GETLicensingInfoInternalServerError with default headers values
func NewGETLicensingInfoInternalServerError() *GETLicensingInfoInternalServerError {
	return &GETLicensingInfoInternalServerError{}
}

/*
GETLicensingInfoInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type GETLicensingInfoInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this get licensing info internal server error response has a 2xx status code
func (o *GETLicensingInfoInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get licensing info internal server error response has a 3xx status code
func (o *GETLicensingInfoInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get licensing info internal server error response has a 4xx status code
func (o *GETLicensingInfoInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get licensing info internal server error response has a 5xx status code
func (o *GETLicensingInfoInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get licensing info internal server error response a status code equal to that given
func (o *GETLicensingInfoInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GETLicensingInfoInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/licensing-info][%d] getLicensingInfoInternalServerError  %+v", 500, o.Payload)
}

func (o *GETLicensingInfoInternalServerError) String() string {
	return fmt.Sprintf("[GET /v1/licensing-info][%d] getLicensingInfoInternalServerError  %+v", 500, o.Payload)
}

func (o *GETLicensingInfoInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *GETLicensingInfoInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
