// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2023 VMware, Inc.
// SPDX-License-Identifier: BSD-2-Clause

package domains

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/vmware/vcf-sdk-go/models"
)

// GetClusterCriteriaReader is a Reader for the GetClusterCriteria structure.
type GetClusterCriteriaReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetClusterCriteriaReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetClusterCriteriaOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetClusterCriteriaBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetClusterCriteriaNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetClusterCriteriaInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /v1/domains/{domainId}/clusters/criteria] getClusterCriteria", response, response.Code())
	}
}

// NewGetClusterCriteriaOK creates a GetClusterCriteriaOK with default headers values
func NewGetClusterCriteriaOK() *GetClusterCriteriaOK {
	return &GetClusterCriteriaOK{}
}

/*
GetClusterCriteriaOK describes a response with status code 200, with default header values.

Ok
*/
type GetClusterCriteriaOK struct {
	Payload *models.PageOfClusterCriterion
}

// IsSuccess returns true when this get cluster criteria o k response has a 2xx status code
func (o *GetClusterCriteriaOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get cluster criteria o k response has a 3xx status code
func (o *GetClusterCriteriaOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get cluster criteria o k response has a 4xx status code
func (o *GetClusterCriteriaOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get cluster criteria o k response has a 5xx status code
func (o *GetClusterCriteriaOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get cluster criteria o k response a status code equal to that given
func (o *GetClusterCriteriaOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get cluster criteria o k response
func (o *GetClusterCriteriaOK) Code() int {
	return 200
}

func (o *GetClusterCriteriaOK) Error() string {
	return fmt.Sprintf("[GET /v1/domains/{domainId}/clusters/criteria][%d] getClusterCriteriaOK  %+v", 200, o.Payload)
}

func (o *GetClusterCriteriaOK) String() string {
	return fmt.Sprintf("[GET /v1/domains/{domainId}/clusters/criteria][%d] getClusterCriteriaOK  %+v", 200, o.Payload)
}

func (o *GetClusterCriteriaOK) GetPayload() *models.PageOfClusterCriterion {
	return o.Payload
}

func (o *GetClusterCriteriaOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PageOfClusterCriterion)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetClusterCriteriaBadRequest creates a GetClusterCriteriaBadRequest with default headers values
func NewGetClusterCriteriaBadRequest() *GetClusterCriteriaBadRequest {
	return &GetClusterCriteriaBadRequest{}
}

/*
GetClusterCriteriaBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetClusterCriteriaBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this get cluster criteria bad request response has a 2xx status code
func (o *GetClusterCriteriaBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get cluster criteria bad request response has a 3xx status code
func (o *GetClusterCriteriaBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get cluster criteria bad request response has a 4xx status code
func (o *GetClusterCriteriaBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get cluster criteria bad request response has a 5xx status code
func (o *GetClusterCriteriaBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get cluster criteria bad request response a status code equal to that given
func (o *GetClusterCriteriaBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get cluster criteria bad request response
func (o *GetClusterCriteriaBadRequest) Code() int {
	return 400
}

func (o *GetClusterCriteriaBadRequest) Error() string {
	return fmt.Sprintf("[GET /v1/domains/{domainId}/clusters/criteria][%d] getClusterCriteriaBadRequest  %+v", 400, o.Payload)
}

func (o *GetClusterCriteriaBadRequest) String() string {
	return fmt.Sprintf("[GET /v1/domains/{domainId}/clusters/criteria][%d] getClusterCriteriaBadRequest  %+v", 400, o.Payload)
}

func (o *GetClusterCriteriaBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetClusterCriteriaBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetClusterCriteriaNotFound creates a GetClusterCriteriaNotFound with default headers values
func NewGetClusterCriteriaNotFound() *GetClusterCriteriaNotFound {
	return &GetClusterCriteriaNotFound{}
}

/*
GetClusterCriteriaNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetClusterCriteriaNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this get cluster criteria not found response has a 2xx status code
func (o *GetClusterCriteriaNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get cluster criteria not found response has a 3xx status code
func (o *GetClusterCriteriaNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get cluster criteria not found response has a 4xx status code
func (o *GetClusterCriteriaNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get cluster criteria not found response has a 5xx status code
func (o *GetClusterCriteriaNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get cluster criteria not found response a status code equal to that given
func (o *GetClusterCriteriaNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get cluster criteria not found response
func (o *GetClusterCriteriaNotFound) Code() int {
	return 404
}

func (o *GetClusterCriteriaNotFound) Error() string {
	return fmt.Sprintf("[GET /v1/domains/{domainId}/clusters/criteria][%d] getClusterCriteriaNotFound  %+v", 404, o.Payload)
}

func (o *GetClusterCriteriaNotFound) String() string {
	return fmt.Sprintf("[GET /v1/domains/{domainId}/clusters/criteria][%d] getClusterCriteriaNotFound  %+v", 404, o.Payload)
}

func (o *GetClusterCriteriaNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetClusterCriteriaNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetClusterCriteriaInternalServerError creates a GetClusterCriteriaInternalServerError with default headers values
func NewGetClusterCriteriaInternalServerError() *GetClusterCriteriaInternalServerError {
	return &GetClusterCriteriaInternalServerError{}
}

/*
GetClusterCriteriaInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetClusterCriteriaInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this get cluster criteria internal server error response has a 2xx status code
func (o *GetClusterCriteriaInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get cluster criteria internal server error response has a 3xx status code
func (o *GetClusterCriteriaInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get cluster criteria internal server error response has a 4xx status code
func (o *GetClusterCriteriaInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get cluster criteria internal server error response has a 5xx status code
func (o *GetClusterCriteriaInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get cluster criteria internal server error response a status code equal to that given
func (o *GetClusterCriteriaInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get cluster criteria internal server error response
func (o *GetClusterCriteriaInternalServerError) Code() int {
	return 500
}

func (o *GetClusterCriteriaInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/domains/{domainId}/clusters/criteria][%d] getClusterCriteriaInternalServerError  %+v", 500, o.Payload)
}

func (o *GetClusterCriteriaInternalServerError) String() string {
	return fmt.Sprintf("[GET /v1/domains/{domainId}/clusters/criteria][%d] getClusterCriteriaInternalServerError  %+v", 500, o.Payload)
}

func (o *GetClusterCriteriaInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetClusterCriteriaInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
