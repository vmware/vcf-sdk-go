// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2023 VMware, Inc.
// SPDX-License-Identifier: BSD-2-Clause

package vcenters

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/vmware/vcf-sdk-go/models"
)

// GetVCENTERSReader is a Reader for the GetVCENTERS structure.
type GetVCENTERSReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetVCENTERSReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetVCENTERSOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetVCENTERSBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetVCENTERSInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /v1/vcenters] getVcenters", response, response.Code())
	}
}

// NewGetVCENTERSOK creates a GetVCENTERSOK with default headers values
func NewGetVCENTERSOK() *GetVCENTERSOK {
	return &GetVCENTERSOK{}
}

/*
GetVCENTERSOK describes a response with status code 200, with default header values.

Ok
*/
type GetVCENTERSOK struct {
	Payload *models.PageOfVcenter
}

// IsSuccess returns true when this get Vcenters o k response has a 2xx status code
func (o *GetVCENTERSOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get Vcenters o k response has a 3xx status code
func (o *GetVCENTERSOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get Vcenters o k response has a 4xx status code
func (o *GetVCENTERSOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get Vcenters o k response has a 5xx status code
func (o *GetVCENTERSOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get Vcenters o k response a status code equal to that given
func (o *GetVCENTERSOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get Vcenters o k response
func (o *GetVCENTERSOK) Code() int {
	return 200
}

func (o *GetVCENTERSOK) Error() string {
	return fmt.Sprintf("[GET /v1/vcenters][%d] getVcentersOK  %+v", 200, o.Payload)
}

func (o *GetVCENTERSOK) String() string {
	return fmt.Sprintf("[GET /v1/vcenters][%d] getVcentersOK  %+v", 200, o.Payload)
}

func (o *GetVCENTERSOK) GetPayload() *models.PageOfVcenter {
	return o.Payload
}

func (o *GetVCENTERSOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PageOfVcenter)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetVCENTERSBadRequest creates a GetVCENTERSBadRequest with default headers values
func NewGetVCENTERSBadRequest() *GetVCENTERSBadRequest {
	return &GetVCENTERSBadRequest{}
}

/*
GetVCENTERSBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetVCENTERSBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this get Vcenters bad request response has a 2xx status code
func (o *GetVCENTERSBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get Vcenters bad request response has a 3xx status code
func (o *GetVCENTERSBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get Vcenters bad request response has a 4xx status code
func (o *GetVCENTERSBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get Vcenters bad request response has a 5xx status code
func (o *GetVCENTERSBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get Vcenters bad request response a status code equal to that given
func (o *GetVCENTERSBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get Vcenters bad request response
func (o *GetVCENTERSBadRequest) Code() int {
	return 400
}

func (o *GetVCENTERSBadRequest) Error() string {
	return fmt.Sprintf("[GET /v1/vcenters][%d] getVcentersBadRequest  %+v", 400, o.Payload)
}

func (o *GetVCENTERSBadRequest) String() string {
	return fmt.Sprintf("[GET /v1/vcenters][%d] getVcentersBadRequest  %+v", 400, o.Payload)
}

func (o *GetVCENTERSBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetVCENTERSBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetVCENTERSInternalServerError creates a GetVCENTERSInternalServerError with default headers values
func NewGetVCENTERSInternalServerError() *GetVCENTERSInternalServerError {
	return &GetVCENTERSInternalServerError{}
}

/*
GetVCENTERSInternalServerError describes a response with status code 500, with default header values.

InternalServerError
*/
type GetVCENTERSInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this get Vcenters internal server error response has a 2xx status code
func (o *GetVCENTERSInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get Vcenters internal server error response has a 3xx status code
func (o *GetVCENTERSInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get Vcenters internal server error response has a 4xx status code
func (o *GetVCENTERSInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get Vcenters internal server error response has a 5xx status code
func (o *GetVCENTERSInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get Vcenters internal server error response a status code equal to that given
func (o *GetVCENTERSInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get Vcenters internal server error response
func (o *GetVCENTERSInternalServerError) Code() int {
	return 500
}

func (o *GetVCENTERSInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/vcenters][%d] getVcentersInternalServerError  %+v", 500, o.Payload)
}

func (o *GetVCENTERSInternalServerError) String() string {
	return fmt.Sprintf("[GET /v1/vcenters][%d] getVcentersInternalServerError  %+v", 500, o.Payload)
}

func (o *GetVCENTERSInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetVCENTERSInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
