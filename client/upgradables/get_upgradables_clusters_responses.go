// Code generated by go-swagger; DO NOT EDIT.

// Copyright © 2023 VMware, Inc. All Rights Reserved.
// Code licenced under: BSD-2

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

// GETUpgradablesClustersReader is a Reader for the GETUpgradablesClusters structure.
type GETUpgradablesClustersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GETUpgradablesClustersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGETUpgradablesClustersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGETUpgradablesClustersNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGETUpgradablesClustersInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGETUpgradablesClustersOK creates a GETUpgradablesClustersOK with default headers values
func NewGETUpgradablesClustersOK() *GETUpgradablesClustersOK {
	return &GETUpgradablesClustersOK{}
}

/*
GETUpgradablesClustersOK describes a response with status code 200, with default header values.

Ok
*/
type GETUpgradablesClustersOK struct {
	Payload *models.PageOfUpgradablesClusterResource
}

// IsSuccess returns true when this get upgradables clusters o k response has a 2xx status code
func (o *GETUpgradablesClustersOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get upgradables clusters o k response has a 3xx status code
func (o *GETUpgradablesClustersOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get upgradables clusters o k response has a 4xx status code
func (o *GETUpgradablesClustersOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get upgradables clusters o k response has a 5xx status code
func (o *GETUpgradablesClustersOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get upgradables clusters o k response a status code equal to that given
func (o *GETUpgradablesClustersOK) IsCode(code int) bool {
	return code == 200
}

func (o *GETUpgradablesClustersOK) Error() string {
	return fmt.Sprintf("[GET /v1/upgradables/domains/{domainId}/clusters][%d] getUpgradablesClustersOK  %+v", 200, o.Payload)
}

func (o *GETUpgradablesClustersOK) String() string {
	return fmt.Sprintf("[GET /v1/upgradables/domains/{domainId}/clusters][%d] getUpgradablesClustersOK  %+v", 200, o.Payload)
}

func (o *GETUpgradablesClustersOK) GetPayload() *models.PageOfUpgradablesClusterResource {
	return o.Payload
}

func (o *GETUpgradablesClustersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PageOfUpgradablesClusterResource)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGETUpgradablesClustersNotFound creates a GETUpgradablesClustersNotFound with default headers values
func NewGETUpgradablesClustersNotFound() *GETUpgradablesClustersNotFound {
	return &GETUpgradablesClustersNotFound{}
}

/*
GETUpgradablesClustersNotFound describes a response with status code 404, with default header values.

Domain Not Found
*/
type GETUpgradablesClustersNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this get upgradables clusters not found response has a 2xx status code
func (o *GETUpgradablesClustersNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get upgradables clusters not found response has a 3xx status code
func (o *GETUpgradablesClustersNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get upgradables clusters not found response has a 4xx status code
func (o *GETUpgradablesClustersNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get upgradables clusters not found response has a 5xx status code
func (o *GETUpgradablesClustersNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get upgradables clusters not found response a status code equal to that given
func (o *GETUpgradablesClustersNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GETUpgradablesClustersNotFound) Error() string {
	return fmt.Sprintf("[GET /v1/upgradables/domains/{domainId}/clusters][%d] getUpgradablesClustersNotFound  %+v", 404, o.Payload)
}

func (o *GETUpgradablesClustersNotFound) String() string {
	return fmt.Sprintf("[GET /v1/upgradables/domains/{domainId}/clusters][%d] getUpgradablesClustersNotFound  %+v", 404, o.Payload)
}

func (o *GETUpgradablesClustersNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GETUpgradablesClustersNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGETUpgradablesClustersInternalServerError creates a GETUpgradablesClustersInternalServerError with default headers values
func NewGETUpgradablesClustersInternalServerError() *GETUpgradablesClustersInternalServerError {
	return &GETUpgradablesClustersInternalServerError{}
}

/*
GETUpgradablesClustersInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GETUpgradablesClustersInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this get upgradables clusters internal server error response has a 2xx status code
func (o *GETUpgradablesClustersInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get upgradables clusters internal server error response has a 3xx status code
func (o *GETUpgradablesClustersInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get upgradables clusters internal server error response has a 4xx status code
func (o *GETUpgradablesClustersInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get upgradables clusters internal server error response has a 5xx status code
func (o *GETUpgradablesClustersInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get upgradables clusters internal server error response a status code equal to that given
func (o *GETUpgradablesClustersInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GETUpgradablesClustersInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/upgradables/domains/{domainId}/clusters][%d] getUpgradablesClustersInternalServerError  %+v", 500, o.Payload)
}

func (o *GETUpgradablesClustersInternalServerError) String() string {
	return fmt.Sprintf("[GET /v1/upgradables/domains/{domainId}/clusters][%d] getUpgradablesClustersInternalServerError  %+v", 500, o.Payload)
}

func (o *GETUpgradablesClustersInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *GETUpgradablesClustersInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
