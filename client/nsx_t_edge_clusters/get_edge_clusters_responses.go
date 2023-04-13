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

// GETEdgeClustersReader is a Reader for the GETEdgeClusters structure.
type GETEdgeClustersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GETEdgeClustersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGETEdgeClustersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGETEdgeClustersBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGETEdgeClustersInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGETEdgeClustersOK creates a GETEdgeClustersOK with default headers values
func NewGETEdgeClustersOK() *GETEdgeClustersOK {
	return &GETEdgeClustersOK{}
}

/*
GETEdgeClustersOK describes a response with status code 200, with default header values.

Ok
*/
type GETEdgeClustersOK struct {
	Payload *models.PageOfEdgeCluster
}

// IsSuccess returns true when this get edge clusters o k response has a 2xx status code
func (o *GETEdgeClustersOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get edge clusters o k response has a 3xx status code
func (o *GETEdgeClustersOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get edge clusters o k response has a 4xx status code
func (o *GETEdgeClustersOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get edge clusters o k response has a 5xx status code
func (o *GETEdgeClustersOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get edge clusters o k response a status code equal to that given
func (o *GETEdgeClustersOK) IsCode(code int) bool {
	return code == 200
}

func (o *GETEdgeClustersOK) Error() string {
	return fmt.Sprintf("[GET /v1/edge-clusters][%d] getEdgeClustersOK  %+v", 200, o.Payload)
}

func (o *GETEdgeClustersOK) String() string {
	return fmt.Sprintf("[GET /v1/edge-clusters][%d] getEdgeClustersOK  %+v", 200, o.Payload)
}

func (o *GETEdgeClustersOK) GetPayload() *models.PageOfEdgeCluster {
	return o.Payload
}

func (o *GETEdgeClustersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PageOfEdgeCluster)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGETEdgeClustersBadRequest creates a GETEdgeClustersBadRequest with default headers values
func NewGETEdgeClustersBadRequest() *GETEdgeClustersBadRequest {
	return &GETEdgeClustersBadRequest{}
}

/*
GETEdgeClustersBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GETEdgeClustersBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this get edge clusters bad request response has a 2xx status code
func (o *GETEdgeClustersBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get edge clusters bad request response has a 3xx status code
func (o *GETEdgeClustersBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get edge clusters bad request response has a 4xx status code
func (o *GETEdgeClustersBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get edge clusters bad request response has a 5xx status code
func (o *GETEdgeClustersBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get edge clusters bad request response a status code equal to that given
func (o *GETEdgeClustersBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GETEdgeClustersBadRequest) Error() string {
	return fmt.Sprintf("[GET /v1/edge-clusters][%d] getEdgeClustersBadRequest  %+v", 400, o.Payload)
}

func (o *GETEdgeClustersBadRequest) String() string {
	return fmt.Sprintf("[GET /v1/edge-clusters][%d] getEdgeClustersBadRequest  %+v", 400, o.Payload)
}

func (o *GETEdgeClustersBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *GETEdgeClustersBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGETEdgeClustersInternalServerError creates a GETEdgeClustersInternalServerError with default headers values
func NewGETEdgeClustersInternalServerError() *GETEdgeClustersInternalServerError {
	return &GETEdgeClustersInternalServerError{}
}

/*
GETEdgeClustersInternalServerError describes a response with status code 500, with default header values.

InternalServerError
*/
type GETEdgeClustersInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this get edge clusters internal server error response has a 2xx status code
func (o *GETEdgeClustersInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get edge clusters internal server error response has a 3xx status code
func (o *GETEdgeClustersInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get edge clusters internal server error response has a 4xx status code
func (o *GETEdgeClustersInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get edge clusters internal server error response has a 5xx status code
func (o *GETEdgeClustersInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get edge clusters internal server error response a status code equal to that given
func (o *GETEdgeClustersInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GETEdgeClustersInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/edge-clusters][%d] getEdgeClustersInternalServerError  %+v", 500, o.Payload)
}

func (o *GETEdgeClustersInternalServerError) String() string {
	return fmt.Sprintf("[GET /v1/edge-clusters][%d] getEdgeClustersInternalServerError  %+v", 500, o.Payload)
}

func (o *GETEdgeClustersInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *GETEdgeClustersInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
