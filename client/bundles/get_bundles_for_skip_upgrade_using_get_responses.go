// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2023 VMware, Inc.
// SPDX-License-Identifier: BSD-2-Clause

package bundles

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/vmware/vcf-sdk-go/models"
)

// GETBundlesForSkipUpgradeUsingGETReader is a Reader for the GETBundlesForSkipUpgradeUsingGET structure.
type GETBundlesForSkipUpgradeUsingGETReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GETBundlesForSkipUpgradeUsingGETReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGETBundlesForSkipUpgradeUsingGETOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGETBundlesForSkipUpgradeUsingGETNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGETBundlesForSkipUpgradeUsingGETInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGETBundlesForSkipUpgradeUsingGETOK creates a GETBundlesForSkipUpgradeUsingGETOK with default headers values
func NewGETBundlesForSkipUpgradeUsingGETOK() *GETBundlesForSkipUpgradeUsingGETOK {
	return &GETBundlesForSkipUpgradeUsingGETOK{}
}

/*
GETBundlesForSkipUpgradeUsingGETOK describes a response with status code 200, with default header values.

Ok
*/
type GETBundlesForSkipUpgradeUsingGETOK struct {
	Payload *models.PageOfBundle
}

// IsSuccess returns true when this get bundles for skip upgrade using Get o k response has a 2xx status code
func (o *GETBundlesForSkipUpgradeUsingGETOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get bundles for skip upgrade using Get o k response has a 3xx status code
func (o *GETBundlesForSkipUpgradeUsingGETOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get bundles for skip upgrade using Get o k response has a 4xx status code
func (o *GETBundlesForSkipUpgradeUsingGETOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get bundles for skip upgrade using Get o k response has a 5xx status code
func (o *GETBundlesForSkipUpgradeUsingGETOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get bundles for skip upgrade using Get o k response a status code equal to that given
func (o *GETBundlesForSkipUpgradeUsingGETOK) IsCode(code int) bool {
	return code == 200
}

func (o *GETBundlesForSkipUpgradeUsingGETOK) Error() string {
	return fmt.Sprintf("[GET /v1/bundles/domains/{id}][%d] getBundlesForSkipUpgradeUsingGetOK  %+v", 200, o.Payload)
}

func (o *GETBundlesForSkipUpgradeUsingGETOK) String() string {
	return fmt.Sprintf("[GET /v1/bundles/domains/{id}][%d] getBundlesForSkipUpgradeUsingGetOK  %+v", 200, o.Payload)
}

func (o *GETBundlesForSkipUpgradeUsingGETOK) GetPayload() *models.PageOfBundle {
	return o.Payload
}

func (o *GETBundlesForSkipUpgradeUsingGETOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PageOfBundle)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGETBundlesForSkipUpgradeUsingGETNotFound creates a GETBundlesForSkipUpgradeUsingGETNotFound with default headers values
func NewGETBundlesForSkipUpgradeUsingGETNotFound() *GETBundlesForSkipUpgradeUsingGETNotFound {
	return &GETBundlesForSkipUpgradeUsingGETNotFound{}
}

/*
GETBundlesForSkipUpgradeUsingGETNotFound describes a response with status code 404, with default header values.

Domain not found with given ID.
*/
type GETBundlesForSkipUpgradeUsingGETNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this get bundles for skip upgrade using Get not found response has a 2xx status code
func (o *GETBundlesForSkipUpgradeUsingGETNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get bundles for skip upgrade using Get not found response has a 3xx status code
func (o *GETBundlesForSkipUpgradeUsingGETNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get bundles for skip upgrade using Get not found response has a 4xx status code
func (o *GETBundlesForSkipUpgradeUsingGETNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get bundles for skip upgrade using Get not found response has a 5xx status code
func (o *GETBundlesForSkipUpgradeUsingGETNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get bundles for skip upgrade using Get not found response a status code equal to that given
func (o *GETBundlesForSkipUpgradeUsingGETNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GETBundlesForSkipUpgradeUsingGETNotFound) Error() string {
	return fmt.Sprintf("[GET /v1/bundles/domains/{id}][%d] getBundlesForSkipUpgradeUsingGetNotFound  %+v", 404, o.Payload)
}

func (o *GETBundlesForSkipUpgradeUsingGETNotFound) String() string {
	return fmt.Sprintf("[GET /v1/bundles/domains/{id}][%d] getBundlesForSkipUpgradeUsingGetNotFound  %+v", 404, o.Payload)
}

func (o *GETBundlesForSkipUpgradeUsingGETNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GETBundlesForSkipUpgradeUsingGETNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGETBundlesForSkipUpgradeUsingGETInternalServerError creates a GETBundlesForSkipUpgradeUsingGETInternalServerError with default headers values
func NewGETBundlesForSkipUpgradeUsingGETInternalServerError() *GETBundlesForSkipUpgradeUsingGETInternalServerError {
	return &GETBundlesForSkipUpgradeUsingGETInternalServerError{}
}

/*
GETBundlesForSkipUpgradeUsingGETInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GETBundlesForSkipUpgradeUsingGETInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this get bundles for skip upgrade using Get internal server error response has a 2xx status code
func (o *GETBundlesForSkipUpgradeUsingGETInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get bundles for skip upgrade using Get internal server error response has a 3xx status code
func (o *GETBundlesForSkipUpgradeUsingGETInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get bundles for skip upgrade using Get internal server error response has a 4xx status code
func (o *GETBundlesForSkipUpgradeUsingGETInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get bundles for skip upgrade using Get internal server error response has a 5xx status code
func (o *GETBundlesForSkipUpgradeUsingGETInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get bundles for skip upgrade using Get internal server error response a status code equal to that given
func (o *GETBundlesForSkipUpgradeUsingGETInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GETBundlesForSkipUpgradeUsingGETInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/bundles/domains/{id}][%d] getBundlesForSkipUpgradeUsingGetInternalServerError  %+v", 500, o.Payload)
}

func (o *GETBundlesForSkipUpgradeUsingGETInternalServerError) String() string {
	return fmt.Sprintf("[GET /v1/bundles/domains/{id}][%d] getBundlesForSkipUpgradeUsingGetInternalServerError  %+v", 500, o.Payload)
}

func (o *GETBundlesForSkipUpgradeUsingGETInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *GETBundlesForSkipUpgradeUsingGETInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
