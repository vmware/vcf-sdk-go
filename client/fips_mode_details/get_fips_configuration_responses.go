// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2023 VMware, Inc.
// SPDX-License-Identifier: BSD-2-Clause

package fips_mode_details

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/vmware/vcf-sdk-go/models"
)

// GetFIPSConfigurationReader is a Reader for the GetFIPSConfiguration structure.
type GetFIPSConfigurationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetFIPSConfigurationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetFIPSConfigurationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewGetFIPSConfigurationInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /v1/system/security/fips] getFIPSConfiguration", response, response.Code())
	}
}

// NewGetFIPSConfigurationOK creates a GetFIPSConfigurationOK with default headers values
func NewGetFIPSConfigurationOK() *GetFIPSConfigurationOK {
	return &GetFIPSConfigurationOK{}
}

/*
GetFIPSConfigurationOK describes a response with status code 200, with default header values.

OK
*/
type GetFIPSConfigurationOK struct {
	Payload *models.FIPS
}

// IsSuccess returns true when this get Fips configuration o k response has a 2xx status code
func (o *GetFIPSConfigurationOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get Fips configuration o k response has a 3xx status code
func (o *GetFIPSConfigurationOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get Fips configuration o k response has a 4xx status code
func (o *GetFIPSConfigurationOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get Fips configuration o k response has a 5xx status code
func (o *GetFIPSConfigurationOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get Fips configuration o k response a status code equal to that given
func (o *GetFIPSConfigurationOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get Fips configuration o k response
func (o *GetFIPSConfigurationOK) Code() int {
	return 200
}

func (o *GetFIPSConfigurationOK) Error() string {
	return fmt.Sprintf("[GET /v1/system/security/fips][%d] getFipsConfigurationOK  %+v", 200, o.Payload)
}

func (o *GetFIPSConfigurationOK) String() string {
	return fmt.Sprintf("[GET /v1/system/security/fips][%d] getFipsConfigurationOK  %+v", 200, o.Payload)
}

func (o *GetFIPSConfigurationOK) GetPayload() *models.FIPS {
	return o.Payload
}

func (o *GetFIPSConfigurationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FIPS)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFIPSConfigurationInternalServerError creates a GetFIPSConfigurationInternalServerError with default headers values
func NewGetFIPSConfigurationInternalServerError() *GetFIPSConfigurationInternalServerError {
	return &GetFIPSConfigurationInternalServerError{}
}

/*
GetFIPSConfigurationInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type GetFIPSConfigurationInternalServerError struct {
}

// IsSuccess returns true when this get Fips configuration internal server error response has a 2xx status code
func (o *GetFIPSConfigurationInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get Fips configuration internal server error response has a 3xx status code
func (o *GetFIPSConfigurationInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get Fips configuration internal server error response has a 4xx status code
func (o *GetFIPSConfigurationInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get Fips configuration internal server error response has a 5xx status code
func (o *GetFIPSConfigurationInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get Fips configuration internal server error response a status code equal to that given
func (o *GetFIPSConfigurationInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get Fips configuration internal server error response
func (o *GetFIPSConfigurationInternalServerError) Code() int {
	return 500
}

func (o *GetFIPSConfigurationInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/system/security/fips][%d] getFipsConfigurationInternalServerError ", 500)
}

func (o *GetFIPSConfigurationInternalServerError) String() string {
	return fmt.Sprintf("[GET /v1/system/security/fips][%d] getFipsConfigurationInternalServerError ", 500)
}

func (o *GetFIPSConfigurationInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}