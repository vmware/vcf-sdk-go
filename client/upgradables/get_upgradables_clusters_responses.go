// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2023 VMware, Inc.
// SPDX-License-Identifier: BSD-2-Clause

package upgradables

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/vmware/vcf-sdk-go/models"
)

// GetUpgradablesClustersReader is a Reader for the GetUpgradablesClusters structure.
type GetUpgradablesClustersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetUpgradablesClustersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetUpgradablesClustersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetUpgradablesClustersNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetUpgradablesClustersInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /v1/upgradables/domains/{domainId}/clusters] getUpgradablesClusters", response, response.Code())
	}
}

// NewGetUpgradablesClustersOK creates a GetUpgradablesClustersOK with default headers values
func NewGetUpgradablesClustersOK() *GetUpgradablesClustersOK {
	return &GetUpgradablesClustersOK{}
}

/*
GetUpgradablesClustersOK describes a response with status code 200, with default header values.

Ok
*/
type GetUpgradablesClustersOK struct {
	Payload *models.PageOfUpgradablesClusterResource
}

// IsSuccess returns true when this get upgradables clusters o k response has a 2xx status code
func (o *GetUpgradablesClustersOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get upgradables clusters o k response has a 3xx status code
func (o *GetUpgradablesClustersOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get upgradables clusters o k response has a 4xx status code
func (o *GetUpgradablesClustersOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get upgradables clusters o k response has a 5xx status code
func (o *GetUpgradablesClustersOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get upgradables clusters o k response a status code equal to that given
func (o *GetUpgradablesClustersOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get upgradables clusters o k response
func (o *GetUpgradablesClustersOK) Code() int {
	return 200
}

func (o *GetUpgradablesClustersOK) Error() string {
	return fmt.Sprintf("[GET /v1/upgradables/domains/{domainId}/clusters][%d] getUpgradablesClustersOK  %+v", 200, o.Payload)
}

func (o *GetUpgradablesClustersOK) String() string {
	return fmt.Sprintf("[GET /v1/upgradables/domains/{domainId}/clusters][%d] getUpgradablesClustersOK  %+v", 200, o.Payload)
}

func (o *GetUpgradablesClustersOK) GetPayload() *models.PageOfUpgradablesClusterResource {
	return o.Payload
}

func (o *GetUpgradablesClustersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PageOfUpgradablesClusterResource)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUpgradablesClustersNotFound creates a GetUpgradablesClustersNotFound with default headers values
func NewGetUpgradablesClustersNotFound() *GetUpgradablesClustersNotFound {
	return &GetUpgradablesClustersNotFound{}
}

/*
GetUpgradablesClustersNotFound describes a response with status code 404, with default header values.

Domain Not Found
*/
type GetUpgradablesClustersNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this get upgradables clusters not found response has a 2xx status code
func (o *GetUpgradablesClustersNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get upgradables clusters not found response has a 3xx status code
func (o *GetUpgradablesClustersNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get upgradables clusters not found response has a 4xx status code
func (o *GetUpgradablesClustersNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get upgradables clusters not found response has a 5xx status code
func (o *GetUpgradablesClustersNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get upgradables clusters not found response a status code equal to that given
func (o *GetUpgradablesClustersNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get upgradables clusters not found response
func (o *GetUpgradablesClustersNotFound) Code() int {
	return 404
}

func (o *GetUpgradablesClustersNotFound) Error() string {
	return fmt.Sprintf("[GET /v1/upgradables/domains/{domainId}/clusters][%d] getUpgradablesClustersNotFound  %+v", 404, o.Payload)
}

func (o *GetUpgradablesClustersNotFound) String() string {
	return fmt.Sprintf("[GET /v1/upgradables/domains/{domainId}/clusters][%d] getUpgradablesClustersNotFound  %+v", 404, o.Payload)
}

func (o *GetUpgradablesClustersNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetUpgradablesClustersNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUpgradablesClustersInternalServerError creates a GetUpgradablesClustersInternalServerError with default headers values
func NewGetUpgradablesClustersInternalServerError() *GetUpgradablesClustersInternalServerError {
	return &GetUpgradablesClustersInternalServerError{}
}

/*
GetUpgradablesClustersInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetUpgradablesClustersInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this get upgradables clusters internal server error response has a 2xx status code
func (o *GetUpgradablesClustersInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get upgradables clusters internal server error response has a 3xx status code
func (o *GetUpgradablesClustersInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get upgradables clusters internal server error response has a 4xx status code
func (o *GetUpgradablesClustersInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get upgradables clusters internal server error response has a 5xx status code
func (o *GetUpgradablesClustersInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get upgradables clusters internal server error response a status code equal to that given
func (o *GetUpgradablesClustersInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get upgradables clusters internal server error response
func (o *GetUpgradablesClustersInternalServerError) Code() int {
	return 500
}

func (o *GetUpgradablesClustersInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/upgradables/domains/{domainId}/clusters][%d] getUpgradablesClustersInternalServerError  %+v", 500, o.Payload)
}

func (o *GetUpgradablesClustersInternalServerError) String() string {
	return fmt.Sprintf("[GET /v1/upgradables/domains/{domainId}/clusters][%d] getUpgradablesClustersInternalServerError  %+v", 500, o.Payload)
}

func (o *GetUpgradablesClustersInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetUpgradablesClustersInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
