// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2023 VMware, Inc.
// SPDX-License-Identifier: BSD-2-Clause

package nsx_t_edge_clusters

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/vmware/vcf-sdk-go/models"
)

// GetEdgeClusterValidationByIDReader is a Reader for the GetEdgeClusterValidationByID structure.
type GetEdgeClusterValidationByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetEdgeClusterValidationByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetEdgeClusterValidationByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetEdgeClusterValidationByIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetEdgeClusterValidationByIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /v1/edge-clusters/validations/{id}] getEdgeClusterValidationByID", response, response.Code())
	}
}

// NewGetEdgeClusterValidationByIDOK creates a GetEdgeClusterValidationByIDOK with default headers values
func NewGetEdgeClusterValidationByIDOK() *GetEdgeClusterValidationByIDOK {
	return &GetEdgeClusterValidationByIDOK{}
}

/*
GetEdgeClusterValidationByIDOK describes a response with status code 200, with default header values.

OK
*/
type GetEdgeClusterValidationByIDOK struct {
	Payload *models.Validation
}

// IsSuccess returns true when this get edge cluster validation by Id o k response has a 2xx status code
func (o *GetEdgeClusterValidationByIDOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get edge cluster validation by Id o k response has a 3xx status code
func (o *GetEdgeClusterValidationByIDOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get edge cluster validation by Id o k response has a 4xx status code
func (o *GetEdgeClusterValidationByIDOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get edge cluster validation by Id o k response has a 5xx status code
func (o *GetEdgeClusterValidationByIDOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get edge cluster validation by Id o k response a status code equal to that given
func (o *GetEdgeClusterValidationByIDOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get edge cluster validation by Id o k response
func (o *GetEdgeClusterValidationByIDOK) Code() int {
	return 200
}

func (o *GetEdgeClusterValidationByIDOK) Error() string {
	return fmt.Sprintf("[GET /v1/edge-clusters/validations/{id}][%d] getEdgeClusterValidationByIdOK  %+v", 200, o.Payload)
}

func (o *GetEdgeClusterValidationByIDOK) String() string {
	return fmt.Sprintf("[GET /v1/edge-clusters/validations/{id}][%d] getEdgeClusterValidationByIdOK  %+v", 200, o.Payload)
}

func (o *GetEdgeClusterValidationByIDOK) GetPayload() *models.Validation {
	return o.Payload
}

func (o *GetEdgeClusterValidationByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Validation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetEdgeClusterValidationByIDBadRequest creates a GetEdgeClusterValidationByIDBadRequest with default headers values
func NewGetEdgeClusterValidationByIDBadRequest() *GetEdgeClusterValidationByIDBadRequest {
	return &GetEdgeClusterValidationByIDBadRequest{}
}

/*
GetEdgeClusterValidationByIDBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetEdgeClusterValidationByIDBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this get edge cluster validation by Id bad request response has a 2xx status code
func (o *GetEdgeClusterValidationByIDBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get edge cluster validation by Id bad request response has a 3xx status code
func (o *GetEdgeClusterValidationByIDBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get edge cluster validation by Id bad request response has a 4xx status code
func (o *GetEdgeClusterValidationByIDBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get edge cluster validation by Id bad request response has a 5xx status code
func (o *GetEdgeClusterValidationByIDBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get edge cluster validation by Id bad request response a status code equal to that given
func (o *GetEdgeClusterValidationByIDBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get edge cluster validation by Id bad request response
func (o *GetEdgeClusterValidationByIDBadRequest) Code() int {
	return 400
}

func (o *GetEdgeClusterValidationByIDBadRequest) Error() string {
	return fmt.Sprintf("[GET /v1/edge-clusters/validations/{id}][%d] getEdgeClusterValidationByIdBadRequest  %+v", 400, o.Payload)
}

func (o *GetEdgeClusterValidationByIDBadRequest) String() string {
	return fmt.Sprintf("[GET /v1/edge-clusters/validations/{id}][%d] getEdgeClusterValidationByIdBadRequest  %+v", 400, o.Payload)
}

func (o *GetEdgeClusterValidationByIDBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetEdgeClusterValidationByIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetEdgeClusterValidationByIDInternalServerError creates a GetEdgeClusterValidationByIDInternalServerError with default headers values
func NewGetEdgeClusterValidationByIDInternalServerError() *GetEdgeClusterValidationByIDInternalServerError {
	return &GetEdgeClusterValidationByIDInternalServerError{}
}

/*
GetEdgeClusterValidationByIDInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type GetEdgeClusterValidationByIDInternalServerError struct {
}

// IsSuccess returns true when this get edge cluster validation by Id internal server error response has a 2xx status code
func (o *GetEdgeClusterValidationByIDInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get edge cluster validation by Id internal server error response has a 3xx status code
func (o *GetEdgeClusterValidationByIDInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get edge cluster validation by Id internal server error response has a 4xx status code
func (o *GetEdgeClusterValidationByIDInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get edge cluster validation by Id internal server error response has a 5xx status code
func (o *GetEdgeClusterValidationByIDInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get edge cluster validation by Id internal server error response a status code equal to that given
func (o *GetEdgeClusterValidationByIDInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get edge cluster validation by Id internal server error response
func (o *GetEdgeClusterValidationByIDInternalServerError) Code() int {
	return 500
}

func (o *GetEdgeClusterValidationByIDInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/edge-clusters/validations/{id}][%d] getEdgeClusterValidationByIdInternalServerError ", 500)
}

func (o *GetEdgeClusterValidationByIDInternalServerError) String() string {
	return fmt.Sprintf("[GET /v1/edge-clusters/validations/{id}][%d] getEdgeClusterValidationByIdInternalServerError ", 500)
}

func (o *GetEdgeClusterValidationByIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
