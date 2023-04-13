// Code generated by go-swagger; DO NOT EDIT.

// Copyright © 2023 VMware, Inc. All Rights Reserved.
// Code licenced under: BSD-2

package nsx_t_edge_clusters

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"vcf-sdk-go/models"
)

// GETEdgeClusterReader is a Reader for the GETEdgeCluster structure.
type GETEdgeClusterReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GETEdgeClusterReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGETEdgeClusterOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGETEdgeClusterNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGETEdgeClusterInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGETEdgeClusterOK creates a GETEdgeClusterOK with default headers values
func NewGETEdgeClusterOK() *GETEdgeClusterOK {
	return &GETEdgeClusterOK{}
}

/*
GETEdgeClusterOK describes a response with status code 200, with default header values.

Ok
*/
type GETEdgeClusterOK struct {
	Payload *models.EdgeCluster
}

// IsSuccess returns true when this get edge cluster o k response has a 2xx status code
func (o *GETEdgeClusterOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get edge cluster o k response has a 3xx status code
func (o *GETEdgeClusterOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get edge cluster o k response has a 4xx status code
func (o *GETEdgeClusterOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get edge cluster o k response has a 5xx status code
func (o *GETEdgeClusterOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get edge cluster o k response a status code equal to that given
func (o *GETEdgeClusterOK) IsCode(code int) bool {
	return code == 200
}

func (o *GETEdgeClusterOK) Error() string {
	return fmt.Sprintf("[GET /v1/edge-clusters/{id}][%d] getEdgeClusterOK  %+v", 200, o.Payload)
}

func (o *GETEdgeClusterOK) String() string {
	return fmt.Sprintf("[GET /v1/edge-clusters/{id}][%d] getEdgeClusterOK  %+v", 200, o.Payload)
}

func (o *GETEdgeClusterOK) GetPayload() *models.EdgeCluster {
	return o.Payload
}

func (o *GETEdgeClusterOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EdgeCluster)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGETEdgeClusterNotFound creates a GETEdgeClusterNotFound with default headers values
func NewGETEdgeClusterNotFound() *GETEdgeClusterNotFound {
	return &GETEdgeClusterNotFound{}
}

/*
GETEdgeClusterNotFound describes a response with status code 404, with default header values.

Edge Cluster not found
*/
type GETEdgeClusterNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this get edge cluster not found response has a 2xx status code
func (o *GETEdgeClusterNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get edge cluster not found response has a 3xx status code
func (o *GETEdgeClusterNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get edge cluster not found response has a 4xx status code
func (o *GETEdgeClusterNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get edge cluster not found response has a 5xx status code
func (o *GETEdgeClusterNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get edge cluster not found response a status code equal to that given
func (o *GETEdgeClusterNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GETEdgeClusterNotFound) Error() string {
	return fmt.Sprintf("[GET /v1/edge-clusters/{id}][%d] getEdgeClusterNotFound  %+v", 404, o.Payload)
}

func (o *GETEdgeClusterNotFound) String() string {
	return fmt.Sprintf("[GET /v1/edge-clusters/{id}][%d] getEdgeClusterNotFound  %+v", 404, o.Payload)
}

func (o *GETEdgeClusterNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GETEdgeClusterNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGETEdgeClusterInternalServerError creates a GETEdgeClusterInternalServerError with default headers values
func NewGETEdgeClusterInternalServerError() *GETEdgeClusterInternalServerError {
	return &GETEdgeClusterInternalServerError{}
}

/*
GETEdgeClusterInternalServerError describes a response with status code 500, with default header values.

InternalServerError
*/
type GETEdgeClusterInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this get edge cluster internal server error response has a 2xx status code
func (o *GETEdgeClusterInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get edge cluster internal server error response has a 3xx status code
func (o *GETEdgeClusterInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get edge cluster internal server error response has a 4xx status code
func (o *GETEdgeClusterInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get edge cluster internal server error response has a 5xx status code
func (o *GETEdgeClusterInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get edge cluster internal server error response a status code equal to that given
func (o *GETEdgeClusterInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GETEdgeClusterInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/edge-clusters/{id}][%d] getEdgeClusterInternalServerError  %+v", 500, o.Payload)
}

func (o *GETEdgeClusterInternalServerError) String() string {
	return fmt.Sprintf("[GET /v1/edge-clusters/{id}][%d] getEdgeClusterInternalServerError  %+v", 500, o.Payload)
}

func (o *GETEdgeClusterInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *GETEdgeClusterInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
