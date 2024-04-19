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

// GetEdgeClusterReader is a Reader for the GetEdgeCluster structure.
type GetEdgeClusterReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetEdgeClusterReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetEdgeClusterOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetEdgeClusterNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetEdgeClusterInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /v1/edge-clusters/{id}] getEdgeCluster", response, response.Code())
	}
}

// NewGetEdgeClusterOK creates a GetEdgeClusterOK with default headers values
func NewGetEdgeClusterOK() *GetEdgeClusterOK {
	return &GetEdgeClusterOK{}
}

/*
GetEdgeClusterOK describes a response with status code 200, with default header values.

Ok
*/
type GetEdgeClusterOK struct {
	Payload *models.EdgeCluster
}

// IsSuccess returns true when this get edge cluster o k response has a 2xx status code
func (o *GetEdgeClusterOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get edge cluster o k response has a 3xx status code
func (o *GetEdgeClusterOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get edge cluster o k response has a 4xx status code
func (o *GetEdgeClusterOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get edge cluster o k response has a 5xx status code
func (o *GetEdgeClusterOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get edge cluster o k response a status code equal to that given
func (o *GetEdgeClusterOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get edge cluster o k response
func (o *GetEdgeClusterOK) Code() int {
	return 200
}

func (o *GetEdgeClusterOK) Error() string {
	return fmt.Sprintf("[GET /v1/edge-clusters/{id}][%d] getEdgeClusterOK  %+v", 200, o.Payload)
}

func (o *GetEdgeClusterOK) String() string {
	return fmt.Sprintf("[GET /v1/edge-clusters/{id}][%d] getEdgeClusterOK  %+v", 200, o.Payload)
}

func (o *GetEdgeClusterOK) GetPayload() *models.EdgeCluster {
	return o.Payload
}

func (o *GetEdgeClusterOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EdgeCluster)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetEdgeClusterNotFound creates a GetEdgeClusterNotFound with default headers values
func NewGetEdgeClusterNotFound() *GetEdgeClusterNotFound {
	return &GetEdgeClusterNotFound{}
}

/*
GetEdgeClusterNotFound describes a response with status code 404, with default header values.

Edge Cluster not found
*/
type GetEdgeClusterNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this get edge cluster not found response has a 2xx status code
func (o *GetEdgeClusterNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get edge cluster not found response has a 3xx status code
func (o *GetEdgeClusterNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get edge cluster not found response has a 4xx status code
func (o *GetEdgeClusterNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get edge cluster not found response has a 5xx status code
func (o *GetEdgeClusterNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get edge cluster not found response a status code equal to that given
func (o *GetEdgeClusterNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get edge cluster not found response
func (o *GetEdgeClusterNotFound) Code() int {
	return 404
}

func (o *GetEdgeClusterNotFound) Error() string {
	return fmt.Sprintf("[GET /v1/edge-clusters/{id}][%d] getEdgeClusterNotFound  %+v", 404, o.Payload)
}

func (o *GetEdgeClusterNotFound) String() string {
	return fmt.Sprintf("[GET /v1/edge-clusters/{id}][%d] getEdgeClusterNotFound  %+v", 404, o.Payload)
}

func (o *GetEdgeClusterNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetEdgeClusterNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetEdgeClusterInternalServerError creates a GetEdgeClusterInternalServerError with default headers values
func NewGetEdgeClusterInternalServerError() *GetEdgeClusterInternalServerError {
	return &GetEdgeClusterInternalServerError{}
}

/*
GetEdgeClusterInternalServerError describes a response with status code 500, with default header values.

InternalServerError
*/
type GetEdgeClusterInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this get edge cluster internal server error response has a 2xx status code
func (o *GetEdgeClusterInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get edge cluster internal server error response has a 3xx status code
func (o *GetEdgeClusterInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get edge cluster internal server error response has a 4xx status code
func (o *GetEdgeClusterInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get edge cluster internal server error response has a 5xx status code
func (o *GetEdgeClusterInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get edge cluster internal server error response a status code equal to that given
func (o *GetEdgeClusterInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get edge cluster internal server error response
func (o *GetEdgeClusterInternalServerError) Code() int {
	return 500
}

func (o *GetEdgeClusterInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/edge-clusters/{id}][%d] getEdgeClusterInternalServerError  %+v", 500, o.Payload)
}

func (o *GetEdgeClusterInternalServerError) String() string {
	return fmt.Sprintf("[GET /v1/edge-clusters/{id}][%d] getEdgeClusterInternalServerError  %+v", 500, o.Payload)
}

func (o *GetEdgeClusterInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetEdgeClusterInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
