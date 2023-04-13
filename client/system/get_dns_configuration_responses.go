// Code generated by go-swagger; DO NOT EDIT.

// Copyright © 2023 VMware, Inc. All Rights Reserved.
// Code licenced under: BSD-2

package system

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"vcf-sdk-go/models"
)

// GETDNSConfigurationReader is a Reader for the GETDNSConfiguration structure.
type GETDNSConfigurationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GETDNSConfigurationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGETDNSConfigurationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGETDNSConfigurationBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGETDNSConfigurationInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGETDNSConfigurationOK creates a GETDNSConfigurationOK with default headers values
func NewGETDNSConfigurationOK() *GETDNSConfigurationOK {
	return &GETDNSConfigurationOK{}
}

/*
GETDNSConfigurationOK describes a response with status code 200, with default header values.

Ok
*/
type GETDNSConfigurationOK struct {
	Payload *models.DNSConfiguration
}

// IsSuccess returns true when this get Dns configuration o k response has a 2xx status code
func (o *GETDNSConfigurationOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get Dns configuration o k response has a 3xx status code
func (o *GETDNSConfigurationOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get Dns configuration o k response has a 4xx status code
func (o *GETDNSConfigurationOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get Dns configuration o k response has a 5xx status code
func (o *GETDNSConfigurationOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get Dns configuration o k response a status code equal to that given
func (o *GETDNSConfigurationOK) IsCode(code int) bool {
	return code == 200
}

func (o *GETDNSConfigurationOK) Error() string {
	return fmt.Sprintf("[GET /v1/system/dns-configuration][%d] getDnsConfigurationOK  %+v", 200, o.Payload)
}

func (o *GETDNSConfigurationOK) String() string {
	return fmt.Sprintf("[GET /v1/system/dns-configuration][%d] getDnsConfigurationOK  %+v", 200, o.Payload)
}

func (o *GETDNSConfigurationOK) GetPayload() *models.DNSConfiguration {
	return o.Payload
}

func (o *GETDNSConfigurationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DNSConfiguration)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGETDNSConfigurationBadRequest creates a GETDNSConfigurationBadRequest with default headers values
func NewGETDNSConfigurationBadRequest() *GETDNSConfigurationBadRequest {
	return &GETDNSConfigurationBadRequest{}
}

/*
GETDNSConfigurationBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GETDNSConfigurationBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this get Dns configuration bad request response has a 2xx status code
func (o *GETDNSConfigurationBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get Dns configuration bad request response has a 3xx status code
func (o *GETDNSConfigurationBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get Dns configuration bad request response has a 4xx status code
func (o *GETDNSConfigurationBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get Dns configuration bad request response has a 5xx status code
func (o *GETDNSConfigurationBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get Dns configuration bad request response a status code equal to that given
func (o *GETDNSConfigurationBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GETDNSConfigurationBadRequest) Error() string {
	return fmt.Sprintf("[GET /v1/system/dns-configuration][%d] getDnsConfigurationBadRequest  %+v", 400, o.Payload)
}

func (o *GETDNSConfigurationBadRequest) String() string {
	return fmt.Sprintf("[GET /v1/system/dns-configuration][%d] getDnsConfigurationBadRequest  %+v", 400, o.Payload)
}

func (o *GETDNSConfigurationBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *GETDNSConfigurationBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGETDNSConfigurationInternalServerError creates a GETDNSConfigurationInternalServerError with default headers values
func NewGETDNSConfigurationInternalServerError() *GETDNSConfigurationInternalServerError {
	return &GETDNSConfigurationInternalServerError{}
}

/*
GETDNSConfigurationInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type GETDNSConfigurationInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this get Dns configuration internal server error response has a 2xx status code
func (o *GETDNSConfigurationInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get Dns configuration internal server error response has a 3xx status code
func (o *GETDNSConfigurationInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get Dns configuration internal server error response has a 4xx status code
func (o *GETDNSConfigurationInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get Dns configuration internal server error response has a 5xx status code
func (o *GETDNSConfigurationInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get Dns configuration internal server error response a status code equal to that given
func (o *GETDNSConfigurationInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GETDNSConfigurationInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/system/dns-configuration][%d] getDnsConfigurationInternalServerError  %+v", 500, o.Payload)
}

func (o *GETDNSConfigurationInternalServerError) String() string {
	return fmt.Sprintf("[GET /v1/system/dns-configuration][%d] getDnsConfigurationInternalServerError  %+v", 500, o.Payload)
}

func (o *GETDNSConfigurationInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *GETDNSConfigurationInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
