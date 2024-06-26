// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2023 VMware, Inc.
// SPDX-License-Identifier: BSD-2-Clause

package nsxt_clusters

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/vmware/vcf-sdk-go/models"
)

// GetNsxClustersReader is a Reader for the GetNsxClusters structure.
type GetNsxClustersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetNsxClustersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetNsxClustersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetNsxClustersBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetNsxClustersInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /v1/nsxt-clusters] getNsxClusters", response, response.Code())
	}
}

// NewGetNsxClustersOK creates a GetNsxClustersOK with default headers values
func NewGetNsxClustersOK() *GetNsxClustersOK {
	return &GetNsxClustersOK{}
}

/*
GetNsxClustersOK describes a response with status code 200, with default header values.

Ok
*/
type GetNsxClustersOK struct {
	Payload *models.PageOfNsxTCluster
}

// IsSuccess returns true when this get nsx clusters o k response has a 2xx status code
func (o *GetNsxClustersOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get nsx clusters o k response has a 3xx status code
func (o *GetNsxClustersOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get nsx clusters o k response has a 4xx status code
func (o *GetNsxClustersOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get nsx clusters o k response has a 5xx status code
func (o *GetNsxClustersOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get nsx clusters o k response a status code equal to that given
func (o *GetNsxClustersOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get nsx clusters o k response
func (o *GetNsxClustersOK) Code() int {
	return 200
}

func (o *GetNsxClustersOK) Error() string {
	return fmt.Sprintf("[GET /v1/nsxt-clusters][%d] getNsxClustersOK  %+v", 200, o.Payload)
}

func (o *GetNsxClustersOK) String() string {
	return fmt.Sprintf("[GET /v1/nsxt-clusters][%d] getNsxClustersOK  %+v", 200, o.Payload)
}

func (o *GetNsxClustersOK) GetPayload() *models.PageOfNsxTCluster {
	return o.Payload
}

func (o *GetNsxClustersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PageOfNsxTCluster)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetNsxClustersBadRequest creates a GetNsxClustersBadRequest with default headers values
func NewGetNsxClustersBadRequest() *GetNsxClustersBadRequest {
	return &GetNsxClustersBadRequest{}
}

/*
GetNsxClustersBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetNsxClustersBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this get nsx clusters bad request response has a 2xx status code
func (o *GetNsxClustersBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get nsx clusters bad request response has a 3xx status code
func (o *GetNsxClustersBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get nsx clusters bad request response has a 4xx status code
func (o *GetNsxClustersBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get nsx clusters bad request response has a 5xx status code
func (o *GetNsxClustersBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get nsx clusters bad request response a status code equal to that given
func (o *GetNsxClustersBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get nsx clusters bad request response
func (o *GetNsxClustersBadRequest) Code() int {
	return 400
}

func (o *GetNsxClustersBadRequest) Error() string {
	return fmt.Sprintf("[GET /v1/nsxt-clusters][%d] getNsxClustersBadRequest  %+v", 400, o.Payload)
}

func (o *GetNsxClustersBadRequest) String() string {
	return fmt.Sprintf("[GET /v1/nsxt-clusters][%d] getNsxClustersBadRequest  %+v", 400, o.Payload)
}

func (o *GetNsxClustersBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetNsxClustersBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetNsxClustersInternalServerError creates a GetNsxClustersInternalServerError with default headers values
func NewGetNsxClustersInternalServerError() *GetNsxClustersInternalServerError {
	return &GetNsxClustersInternalServerError{}
}

/*
GetNsxClustersInternalServerError describes a response with status code 500, with default header values.

InternalServerError
*/
type GetNsxClustersInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this get nsx clusters internal server error response has a 2xx status code
func (o *GetNsxClustersInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get nsx clusters internal server error response has a 3xx status code
func (o *GetNsxClustersInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get nsx clusters internal server error response has a 4xx status code
func (o *GetNsxClustersInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get nsx clusters internal server error response has a 5xx status code
func (o *GetNsxClustersInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get nsx clusters internal server error response a status code equal to that given
func (o *GetNsxClustersInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get nsx clusters internal server error response
func (o *GetNsxClustersInternalServerError) Code() int {
	return 500
}

func (o *GetNsxClustersInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/nsxt-clusters][%d] getNsxClustersInternalServerError  %+v", 500, o.Payload)
}

func (o *GetNsxClustersInternalServerError) String() string {
	return fmt.Sprintf("[GET /v1/nsxt-clusters][%d] getNsxClustersInternalServerError  %+v", 500, o.Payload)
}

func (o *GetNsxClustersInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetNsxClustersInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
