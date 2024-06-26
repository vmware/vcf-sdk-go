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

// GetBringupAppInfoReader is a Reader for the GetBringupAppInfo structure.
type GetBringupAppInfoReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetBringupAppInfoReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetBringupAppInfoOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetBringupAppInfoNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetBringupAppInfoInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 501:
		result := NewGetBringupAppInfoNotImplemented()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /v1/sddcs/about] getBringupAppInfo", response, response.Code())
	}
}

// NewGetBringupAppInfoOK creates a GetBringupAppInfoOK with default headers values
func NewGetBringupAppInfoOK() *GetBringupAppInfoOK {
	return &GetBringupAppInfoOK{}
}

/*
GetBringupAppInfoOK describes a response with status code 200, with default header values.

OK
*/
type GetBringupAppInfoOK struct {
	Payload *models.VcfService
}

// IsSuccess returns true when this get bringup app info o k response has a 2xx status code
func (o *GetBringupAppInfoOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get bringup app info o k response has a 3xx status code
func (o *GetBringupAppInfoOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get bringup app info o k response has a 4xx status code
func (o *GetBringupAppInfoOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get bringup app info o k response has a 5xx status code
func (o *GetBringupAppInfoOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get bringup app info o k response a status code equal to that given
func (o *GetBringupAppInfoOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get bringup app info o k response
func (o *GetBringupAppInfoOK) Code() int {
	return 200
}

func (o *GetBringupAppInfoOK) Error() string {
	return fmt.Sprintf("[GET /v1/sddcs/about][%d] getBringupAppInfoOK  %+v", 200, o.Payload)
}

func (o *GetBringupAppInfoOK) String() string {
	return fmt.Sprintf("[GET /v1/sddcs/about][%d] getBringupAppInfoOK  %+v", 200, o.Payload)
}

func (o *GetBringupAppInfoOK) GetPayload() *models.VcfService {
	return o.Payload
}

func (o *GetBringupAppInfoOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VcfService)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetBringupAppInfoNotFound creates a GetBringupAppInfoNotFound with default headers values
func NewGetBringupAppInfoNotFound() *GetBringupAppInfoNotFound {
	return &GetBringupAppInfoNotFound{}
}

/*
GetBringupAppInfoNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetBringupAppInfoNotFound struct {
}

// IsSuccess returns true when this get bringup app info not found response has a 2xx status code
func (o *GetBringupAppInfoNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get bringup app info not found response has a 3xx status code
func (o *GetBringupAppInfoNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get bringup app info not found response has a 4xx status code
func (o *GetBringupAppInfoNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get bringup app info not found response has a 5xx status code
func (o *GetBringupAppInfoNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get bringup app info not found response a status code equal to that given
func (o *GetBringupAppInfoNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get bringup app info not found response
func (o *GetBringupAppInfoNotFound) Code() int {
	return 404
}

func (o *GetBringupAppInfoNotFound) Error() string {
	return fmt.Sprintf("[GET /v1/sddcs/about][%d] getBringupAppInfoNotFound ", 404)
}

func (o *GetBringupAppInfoNotFound) String() string {
	return fmt.Sprintf("[GET /v1/sddcs/about][%d] getBringupAppInfoNotFound ", 404)
}

func (o *GetBringupAppInfoNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetBringupAppInfoInternalServerError creates a GetBringupAppInfoInternalServerError with default headers values
func NewGetBringupAppInfoInternalServerError() *GetBringupAppInfoInternalServerError {
	return &GetBringupAppInfoInternalServerError{}
}

/*
GetBringupAppInfoInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetBringupAppInfoInternalServerError struct {
}

// IsSuccess returns true when this get bringup app info internal server error response has a 2xx status code
func (o *GetBringupAppInfoInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get bringup app info internal server error response has a 3xx status code
func (o *GetBringupAppInfoInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get bringup app info internal server error response has a 4xx status code
func (o *GetBringupAppInfoInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get bringup app info internal server error response has a 5xx status code
func (o *GetBringupAppInfoInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get bringup app info internal server error response a status code equal to that given
func (o *GetBringupAppInfoInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get bringup app info internal server error response
func (o *GetBringupAppInfoInternalServerError) Code() int {
	return 500
}

func (o *GetBringupAppInfoInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/sddcs/about][%d] getBringupAppInfoInternalServerError ", 500)
}

func (o *GetBringupAppInfoInternalServerError) String() string {
	return fmt.Sprintf("[GET /v1/sddcs/about][%d] getBringupAppInfoInternalServerError ", 500)
}

func (o *GetBringupAppInfoInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetBringupAppInfoNotImplemented creates a GetBringupAppInfoNotImplemented with default headers values
func NewGetBringupAppInfoNotImplemented() *GetBringupAppInfoNotImplemented {
	return &GetBringupAppInfoNotImplemented{}
}

/*
GetBringupAppInfoNotImplemented describes a response with status code 501, with default header values.

Not Implemented
*/
type GetBringupAppInfoNotImplemented struct {
}

// IsSuccess returns true when this get bringup app info not implemented response has a 2xx status code
func (o *GetBringupAppInfoNotImplemented) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get bringup app info not implemented response has a 3xx status code
func (o *GetBringupAppInfoNotImplemented) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get bringup app info not implemented response has a 4xx status code
func (o *GetBringupAppInfoNotImplemented) IsClientError() bool {
	return false
}

// IsServerError returns true when this get bringup app info not implemented response has a 5xx status code
func (o *GetBringupAppInfoNotImplemented) IsServerError() bool {
	return true
}

// IsCode returns true when this get bringup app info not implemented response a status code equal to that given
func (o *GetBringupAppInfoNotImplemented) IsCode(code int) bool {
	return code == 501
}

// Code gets the status code for the get bringup app info not implemented response
func (o *GetBringupAppInfoNotImplemented) Code() int {
	return 501
}

func (o *GetBringupAppInfoNotImplemented) Error() string {
	return fmt.Sprintf("[GET /v1/sddcs/about][%d] getBringupAppInfoNotImplemented ", 501)
}

func (o *GetBringupAppInfoNotImplemented) String() string {
	return fmt.Sprintf("[GET /v1/sddcs/about][%d] getBringupAppInfoNotImplemented ", 501)
}

func (o *GetBringupAppInfoNotImplemented) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
