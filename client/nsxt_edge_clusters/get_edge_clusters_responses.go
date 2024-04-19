// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2023 VMware, Inc.
// SPDX-License-Identifier: BSD-2-Clause

package nsxt_edge_clusters

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/vmware/vcf-sdk-go/models"
)

// GetEdgeClustersReader is a Reader for the GetEdgeClusters structure.
type GetEdgeClustersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetEdgeClustersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetEdgeClustersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetEdgeClustersBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetEdgeClustersInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /v1/edge-clusters] getEdgeClusters", response, response.Code())
	}
}

// NewGetEdgeClustersOK creates a GetEdgeClustersOK with default headers values
func NewGetEdgeClustersOK() *GetEdgeClustersOK {
	return &GetEdgeClustersOK{}
}

/*
GetEdgeClustersOK describes a response with status code 200, with default header values.

Ok
*/
type GetEdgeClustersOK struct {
	Payload *models.PageOfEdgeCluster
}

// IsSuccess returns true when this get edge clusters o k response has a 2xx status code
func (o *GetEdgeClustersOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get edge clusters o k response has a 3xx status code
func (o *GetEdgeClustersOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get edge clusters o k response has a 4xx status code
func (o *GetEdgeClustersOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get edge clusters o k response has a 5xx status code
func (o *GetEdgeClustersOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get edge clusters o k response a status code equal to that given
func (o *GetEdgeClustersOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get edge clusters o k response
func (o *GetEdgeClustersOK) Code() int {
	return 200
}

func (o *GetEdgeClustersOK) Error() string {
	return fmt.Sprintf("[GET /v1/edge-clusters][%d] getEdgeClustersOK  %+v", 200, o.Payload)
}

func (o *GetEdgeClustersOK) String() string {
	return fmt.Sprintf("[GET /v1/edge-clusters][%d] getEdgeClustersOK  %+v", 200, o.Payload)
}

func (o *GetEdgeClustersOK) GetPayload() *models.PageOfEdgeCluster {
	return o.Payload
}

func (o *GetEdgeClustersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PageOfEdgeCluster)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetEdgeClustersBadRequest creates a GetEdgeClustersBadRequest with default headers values
func NewGetEdgeClustersBadRequest() *GetEdgeClustersBadRequest {
	return &GetEdgeClustersBadRequest{}
}

/*
GetEdgeClustersBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetEdgeClustersBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this get edge clusters bad request response has a 2xx status code
func (o *GetEdgeClustersBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get edge clusters bad request response has a 3xx status code
func (o *GetEdgeClustersBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get edge clusters bad request response has a 4xx status code
func (o *GetEdgeClustersBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get edge clusters bad request response has a 5xx status code
func (o *GetEdgeClustersBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get edge clusters bad request response a status code equal to that given
func (o *GetEdgeClustersBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get edge clusters bad request response
func (o *GetEdgeClustersBadRequest) Code() int {
	return 400
}

func (o *GetEdgeClustersBadRequest) Error() string {
	return fmt.Sprintf("[GET /v1/edge-clusters][%d] getEdgeClustersBadRequest  %+v", 400, o.Payload)
}

func (o *GetEdgeClustersBadRequest) String() string {
	return fmt.Sprintf("[GET /v1/edge-clusters][%d] getEdgeClustersBadRequest  %+v", 400, o.Payload)
}

func (o *GetEdgeClustersBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetEdgeClustersBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetEdgeClustersInternalServerError creates a GetEdgeClustersInternalServerError with default headers values
func NewGetEdgeClustersInternalServerError() *GetEdgeClustersInternalServerError {
	return &GetEdgeClustersInternalServerError{}
}

/*
GetEdgeClustersInternalServerError describes a response with status code 500, with default header values.

InternalServerError
*/
type GetEdgeClustersInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this get edge clusters internal server error response has a 2xx status code
func (o *GetEdgeClustersInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get edge clusters internal server error response has a 3xx status code
func (o *GetEdgeClustersInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get edge clusters internal server error response has a 4xx status code
func (o *GetEdgeClustersInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get edge clusters internal server error response has a 5xx status code
func (o *GetEdgeClustersInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get edge clusters internal server error response a status code equal to that given
func (o *GetEdgeClustersInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get edge clusters internal server error response
func (o *GetEdgeClustersInternalServerError) Code() int {
	return 500
}

func (o *GetEdgeClustersInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/edge-clusters][%d] getEdgeClustersInternalServerError  %+v", 500, o.Payload)
}

func (o *GetEdgeClustersInternalServerError) String() string {
	return fmt.Sprintf("[GET /v1/edge-clusters][%d] getEdgeClustersInternalServerError  %+v", 500, o.Payload)
}

func (o *GetEdgeClustersInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetEdgeClustersInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
