// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2023 VMware, Inc.
// SPDX-License-Identifier: BSD-2-Clause

package target_upgrade_version

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/vmware/vcf-sdk-go/models"
)

// GetReleaseByDomain1Reader is a Reader for the GetReleaseByDomain1 structure.
type GetReleaseByDomain1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetReleaseByDomain1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetReleaseByDomain1OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewGetReleaseByDomain1InternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /v1/releases/domains] getReleaseByDomain_1", response, response.Code())
	}
}

// NewGetReleaseByDomain1OK creates a GetReleaseByDomain1OK with default headers values
func NewGetReleaseByDomain1OK() *GetReleaseByDomain1OK {
	return &GetReleaseByDomain1OK{}
}

/*
GetReleaseByDomain1OK describes a response with status code 200, with default header values.

Ok
*/
type GetReleaseByDomain1OK struct {
	Payload *models.PageOfDomainReleaseView
}

// IsSuccess returns true when this get release by domain1 o k response has a 2xx status code
func (o *GetReleaseByDomain1OK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get release by domain1 o k response has a 3xx status code
func (o *GetReleaseByDomain1OK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get release by domain1 o k response has a 4xx status code
func (o *GetReleaseByDomain1OK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get release by domain1 o k response has a 5xx status code
func (o *GetReleaseByDomain1OK) IsServerError() bool {
	return false
}

// IsCode returns true when this get release by domain1 o k response a status code equal to that given
func (o *GetReleaseByDomain1OK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get release by domain1 o k response
func (o *GetReleaseByDomain1OK) Code() int {
	return 200
}

func (o *GetReleaseByDomain1OK) Error() string {
	return fmt.Sprintf("[GET /v1/releases/domains][%d] getReleaseByDomain1OK  %+v", 200, o.Payload)
}

func (o *GetReleaseByDomain1OK) String() string {
	return fmt.Sprintf("[GET /v1/releases/domains][%d] getReleaseByDomain1OK  %+v", 200, o.Payload)
}

func (o *GetReleaseByDomain1OK) GetPayload() *models.PageOfDomainReleaseView {
	return o.Payload
}

func (o *GetReleaseByDomain1OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PageOfDomainReleaseView)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetReleaseByDomain1InternalServerError creates a GetReleaseByDomain1InternalServerError with default headers values
func NewGetReleaseByDomain1InternalServerError() *GetReleaseByDomain1InternalServerError {
	return &GetReleaseByDomain1InternalServerError{}
}

/*
GetReleaseByDomain1InternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetReleaseByDomain1InternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this get release by domain1 internal server error response has a 2xx status code
func (o *GetReleaseByDomain1InternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get release by domain1 internal server error response has a 3xx status code
func (o *GetReleaseByDomain1InternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get release by domain1 internal server error response has a 4xx status code
func (o *GetReleaseByDomain1InternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get release by domain1 internal server error response has a 5xx status code
func (o *GetReleaseByDomain1InternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get release by domain1 internal server error response a status code equal to that given
func (o *GetReleaseByDomain1InternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get release by domain1 internal server error response
func (o *GetReleaseByDomain1InternalServerError) Code() int {
	return 500
}

func (o *GetReleaseByDomain1InternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/releases/domains][%d] getReleaseByDomain1InternalServerError  %+v", 500, o.Payload)
}

func (o *GetReleaseByDomain1InternalServerError) String() string {
	return fmt.Sprintf("[GET /v1/releases/domains][%d] getReleaseByDomain1InternalServerError  %+v", 500, o.Payload)
}

func (o *GetReleaseByDomain1InternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetReleaseByDomain1InternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
