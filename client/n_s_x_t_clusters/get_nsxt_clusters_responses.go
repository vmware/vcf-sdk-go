// Code generated by go-swagger; DO NOT EDIT.

// Copyright © 2023 VMware, Inc. All Rights Reserved.
// Code licenced under: BSD-2

package n_s_x_t_clusters

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/vmware/vcf-sdk-go/models"
)

// GETNSXTClustersReader is a Reader for the GETNSXTClusters structure.
type GETNSXTClustersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GETNSXTClustersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGETNSXTClustersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGETNSXTClustersBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGETNSXTClustersInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGETNSXTClustersOK creates a GETNSXTClustersOK with default headers values
func NewGETNSXTClustersOK() *GETNSXTClustersOK {
	return &GETNSXTClustersOK{}
}

/*
GETNSXTClustersOK describes a response with status code 200, with default header values.

Ok
*/
type GETNSXTClustersOK struct {
	Payload *models.PageOfNsxTCluster
}

// IsSuccess returns true when this get Nsxt clusters o k response has a 2xx status code
func (o *GETNSXTClustersOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get Nsxt clusters o k response has a 3xx status code
func (o *GETNSXTClustersOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get Nsxt clusters o k response has a 4xx status code
func (o *GETNSXTClustersOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get Nsxt clusters o k response has a 5xx status code
func (o *GETNSXTClustersOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get Nsxt clusters o k response a status code equal to that given
func (o *GETNSXTClustersOK) IsCode(code int) bool {
	return code == 200
}

func (o *GETNSXTClustersOK) Error() string {
	return fmt.Sprintf("[GET /v1/nsxt-clusters][%d] getNsxtClustersOK  %+v", 200, o.Payload)
}

func (o *GETNSXTClustersOK) String() string {
	return fmt.Sprintf("[GET /v1/nsxt-clusters][%d] getNsxtClustersOK  %+v", 200, o.Payload)
}

func (o *GETNSXTClustersOK) GetPayload() *models.PageOfNsxTCluster {
	return o.Payload
}

func (o *GETNSXTClustersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PageOfNsxTCluster)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGETNSXTClustersBadRequest creates a GETNSXTClustersBadRequest with default headers values
func NewGETNSXTClustersBadRequest() *GETNSXTClustersBadRequest {
	return &GETNSXTClustersBadRequest{}
}

/*
GETNSXTClustersBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GETNSXTClustersBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this get Nsxt clusters bad request response has a 2xx status code
func (o *GETNSXTClustersBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get Nsxt clusters bad request response has a 3xx status code
func (o *GETNSXTClustersBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get Nsxt clusters bad request response has a 4xx status code
func (o *GETNSXTClustersBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get Nsxt clusters bad request response has a 5xx status code
func (o *GETNSXTClustersBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get Nsxt clusters bad request response a status code equal to that given
func (o *GETNSXTClustersBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GETNSXTClustersBadRequest) Error() string {
	return fmt.Sprintf("[GET /v1/nsxt-clusters][%d] getNsxtClustersBadRequest  %+v", 400, o.Payload)
}

func (o *GETNSXTClustersBadRequest) String() string {
	return fmt.Sprintf("[GET /v1/nsxt-clusters][%d] getNsxtClustersBadRequest  %+v", 400, o.Payload)
}

func (o *GETNSXTClustersBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *GETNSXTClustersBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGETNSXTClustersInternalServerError creates a GETNSXTClustersInternalServerError with default headers values
func NewGETNSXTClustersInternalServerError() *GETNSXTClustersInternalServerError {
	return &GETNSXTClustersInternalServerError{}
}

/*
GETNSXTClustersInternalServerError describes a response with status code 500, with default header values.

InternalServerError
*/
type GETNSXTClustersInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this get Nsxt clusters internal server error response has a 2xx status code
func (o *GETNSXTClustersInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get Nsxt clusters internal server error response has a 3xx status code
func (o *GETNSXTClustersInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get Nsxt clusters internal server error response has a 4xx status code
func (o *GETNSXTClustersInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get Nsxt clusters internal server error response has a 5xx status code
func (o *GETNSXTClustersInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get Nsxt clusters internal server error response a status code equal to that given
func (o *GETNSXTClustersInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GETNSXTClustersInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/nsxt-clusters][%d] getNsxtClustersInternalServerError  %+v", 500, o.Payload)
}

func (o *GETNSXTClustersInternalServerError) String() string {
	return fmt.Sprintf("[GET /v1/nsxt-clusters][%d] getNsxtClustersInternalServerError  %+v", 500, o.Payload)
}

func (o *GETNSXTClustersInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *GETNSXTClustersInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
