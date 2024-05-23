// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2023 VMware, Inc.
// SPDX-License-Identifier: BSD-2-Clause

package license_keys

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/vmware/vcf-sdk-go/models"
)

// GetLicenseKeyReader is a Reader for the GetLicenseKey structure.
type GetLicenseKeyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLicenseKeyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLicenseKeyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetLicenseKeyNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetLicenseKeyInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /v1/license-keys/{key}] getLicenseKey", response, response.Code())
	}
}

// NewGetLicenseKeyOK creates a GetLicenseKeyOK with default headers values
func NewGetLicenseKeyOK() *GetLicenseKeyOK {
	return &GetLicenseKeyOK{}
}

/*
GetLicenseKeyOK describes a response with status code 200, with default header values.

Successful
*/
type GetLicenseKeyOK struct {
	Payload *models.LicenseKey
}

// IsSuccess returns true when this get license key o k response has a 2xx status code
func (o *GetLicenseKeyOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get license key o k response has a 3xx status code
func (o *GetLicenseKeyOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get license key o k response has a 4xx status code
func (o *GetLicenseKeyOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get license key o k response has a 5xx status code
func (o *GetLicenseKeyOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get license key o k response a status code equal to that given
func (o *GetLicenseKeyOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get license key o k response
func (o *GetLicenseKeyOK) Code() int {
	return 200
}

func (o *GetLicenseKeyOK) Error() string {
	return fmt.Sprintf("[GET /v1/license-keys/{key}][%d] getLicenseKeyOK  %+v", 200, o.Payload)
}

func (o *GetLicenseKeyOK) String() string {
	return fmt.Sprintf("[GET /v1/license-keys/{key}][%d] getLicenseKeyOK  %+v", 200, o.Payload)
}

func (o *GetLicenseKeyOK) GetPayload() *models.LicenseKey {
	return o.Payload
}

func (o *GetLicenseKeyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.LicenseKey)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLicenseKeyNotFound creates a GetLicenseKeyNotFound with default headers values
func NewGetLicenseKeyNotFound() *GetLicenseKeyNotFound {
	return &GetLicenseKeyNotFound{}
}

/*
GetLicenseKeyNotFound describes a response with status code 404, with default header values.

License key not found
*/
type GetLicenseKeyNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this get license key not found response has a 2xx status code
func (o *GetLicenseKeyNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get license key not found response has a 3xx status code
func (o *GetLicenseKeyNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get license key not found response has a 4xx status code
func (o *GetLicenseKeyNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get license key not found response has a 5xx status code
func (o *GetLicenseKeyNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get license key not found response a status code equal to that given
func (o *GetLicenseKeyNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get license key not found response
func (o *GetLicenseKeyNotFound) Code() int {
	return 404
}

func (o *GetLicenseKeyNotFound) Error() string {
	return fmt.Sprintf("[GET /v1/license-keys/{key}][%d] getLicenseKeyNotFound  %+v", 404, o.Payload)
}

func (o *GetLicenseKeyNotFound) String() string {
	return fmt.Sprintf("[GET /v1/license-keys/{key}][%d] getLicenseKeyNotFound  %+v", 404, o.Payload)
}

func (o *GetLicenseKeyNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetLicenseKeyNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLicenseKeyInternalServerError creates a GetLicenseKeyInternalServerError with default headers values
func NewGetLicenseKeyInternalServerError() *GetLicenseKeyInternalServerError {
	return &GetLicenseKeyInternalServerError{}
}

/*
GetLicenseKeyInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type GetLicenseKeyInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this get license key internal server error response has a 2xx status code
func (o *GetLicenseKeyInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get license key internal server error response has a 3xx status code
func (o *GetLicenseKeyInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get license key internal server error response has a 4xx status code
func (o *GetLicenseKeyInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get license key internal server error response has a 5xx status code
func (o *GetLicenseKeyInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get license key internal server error response a status code equal to that given
func (o *GetLicenseKeyInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get license key internal server error response
func (o *GetLicenseKeyInternalServerError) Code() int {
	return 500
}

func (o *GetLicenseKeyInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/license-keys/{key}][%d] getLicenseKeyInternalServerError  %+v", 500, o.Payload)
}

func (o *GetLicenseKeyInternalServerError) String() string {
	return fmt.Sprintf("[GET /v1/license-keys/{key}][%d] getLicenseKeyInternalServerError  %+v", 500, o.Payload)
}

func (o *GetLicenseKeyInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetLicenseKeyInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
