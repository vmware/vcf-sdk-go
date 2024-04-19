// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2023 VMware, Inc.
// SPDX-License-Identifier: BSD-2-Clause

package clusters

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/vmware/vcf-sdk-go/models"
)

// GetClusterNetworkConfigurationQueryResponseReader is a Reader for the GetClusterNetworkConfigurationQueryResponse structure.
type GetClusterNetworkConfigurationQueryResponseReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetClusterNetworkConfigurationQueryResponseReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetClusterNetworkConfigurationQueryResponseOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetClusterNetworkConfigurationQueryResponseBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetClusterNetworkConfigurationQueryResponseInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /v1/clusters/{id}/network/queries/{queryId}] getClusterNetworkConfigurationQueryResponse", response, response.Code())
	}
}

// NewGetClusterNetworkConfigurationQueryResponseOK creates a GetClusterNetworkConfigurationQueryResponseOK with default headers values
func NewGetClusterNetworkConfigurationQueryResponseOK() *GetClusterNetworkConfigurationQueryResponseOK {
	return &GetClusterNetworkConfigurationQueryResponseOK{}
}

/*
GetClusterNetworkConfigurationQueryResponseOK describes a response with status code 200, with default header values.

Ok
*/
type GetClusterNetworkConfigurationQueryResponseOK struct {
	Payload *models.ClusterNetworkConfigurationQueryResponse
}

// IsSuccess returns true when this get cluster network configuration query response o k response has a 2xx status code
func (o *GetClusterNetworkConfigurationQueryResponseOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get cluster network configuration query response o k response has a 3xx status code
func (o *GetClusterNetworkConfigurationQueryResponseOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get cluster network configuration query response o k response has a 4xx status code
func (o *GetClusterNetworkConfigurationQueryResponseOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get cluster network configuration query response o k response has a 5xx status code
func (o *GetClusterNetworkConfigurationQueryResponseOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get cluster network configuration query response o k response a status code equal to that given
func (o *GetClusterNetworkConfigurationQueryResponseOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get cluster network configuration query response o k response
func (o *GetClusterNetworkConfigurationQueryResponseOK) Code() int {
	return 200
}

func (o *GetClusterNetworkConfigurationQueryResponseOK) Error() string {
	return fmt.Sprintf("[GET /v1/clusters/{id}/network/queries/{queryId}][%d] getClusterNetworkConfigurationQueryResponseOK  %+v", 200, o.Payload)
}

func (o *GetClusterNetworkConfigurationQueryResponseOK) String() string {
	return fmt.Sprintf("[GET /v1/clusters/{id}/network/queries/{queryId}][%d] getClusterNetworkConfigurationQueryResponseOK  %+v", 200, o.Payload)
}

func (o *GetClusterNetworkConfigurationQueryResponseOK) GetPayload() *models.ClusterNetworkConfigurationQueryResponse {
	return o.Payload
}

func (o *GetClusterNetworkConfigurationQueryResponseOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ClusterNetworkConfigurationQueryResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetClusterNetworkConfigurationQueryResponseBadRequest creates a GetClusterNetworkConfigurationQueryResponseBadRequest with default headers values
func NewGetClusterNetworkConfigurationQueryResponseBadRequest() *GetClusterNetworkConfigurationQueryResponseBadRequest {
	return &GetClusterNetworkConfigurationQueryResponseBadRequest{}
}

/*
GetClusterNetworkConfigurationQueryResponseBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetClusterNetworkConfigurationQueryResponseBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this get cluster network configuration query response bad request response has a 2xx status code
func (o *GetClusterNetworkConfigurationQueryResponseBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get cluster network configuration query response bad request response has a 3xx status code
func (o *GetClusterNetworkConfigurationQueryResponseBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get cluster network configuration query response bad request response has a 4xx status code
func (o *GetClusterNetworkConfigurationQueryResponseBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get cluster network configuration query response bad request response has a 5xx status code
func (o *GetClusterNetworkConfigurationQueryResponseBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get cluster network configuration query response bad request response a status code equal to that given
func (o *GetClusterNetworkConfigurationQueryResponseBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get cluster network configuration query response bad request response
func (o *GetClusterNetworkConfigurationQueryResponseBadRequest) Code() int {
	return 400
}

func (o *GetClusterNetworkConfigurationQueryResponseBadRequest) Error() string {
	return fmt.Sprintf("[GET /v1/clusters/{id}/network/queries/{queryId}][%d] getClusterNetworkConfigurationQueryResponseBadRequest  %+v", 400, o.Payload)
}

func (o *GetClusterNetworkConfigurationQueryResponseBadRequest) String() string {
	return fmt.Sprintf("[GET /v1/clusters/{id}/network/queries/{queryId}][%d] getClusterNetworkConfigurationQueryResponseBadRequest  %+v", 400, o.Payload)
}

func (o *GetClusterNetworkConfigurationQueryResponseBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetClusterNetworkConfigurationQueryResponseBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetClusterNetworkConfigurationQueryResponseInternalServerError creates a GetClusterNetworkConfigurationQueryResponseInternalServerError with default headers values
func NewGetClusterNetworkConfigurationQueryResponseInternalServerError() *GetClusterNetworkConfigurationQueryResponseInternalServerError {
	return &GetClusterNetworkConfigurationQueryResponseInternalServerError{}
}

/*
GetClusterNetworkConfigurationQueryResponseInternalServerError describes a response with status code 500, with default header values.

InternalServerError
*/
type GetClusterNetworkConfigurationQueryResponseInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this get cluster network configuration query response internal server error response has a 2xx status code
func (o *GetClusterNetworkConfigurationQueryResponseInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get cluster network configuration query response internal server error response has a 3xx status code
func (o *GetClusterNetworkConfigurationQueryResponseInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get cluster network configuration query response internal server error response has a 4xx status code
func (o *GetClusterNetworkConfigurationQueryResponseInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get cluster network configuration query response internal server error response has a 5xx status code
func (o *GetClusterNetworkConfigurationQueryResponseInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get cluster network configuration query response internal server error response a status code equal to that given
func (o *GetClusterNetworkConfigurationQueryResponseInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get cluster network configuration query response internal server error response
func (o *GetClusterNetworkConfigurationQueryResponseInternalServerError) Code() int {
	return 500
}

func (o *GetClusterNetworkConfigurationQueryResponseInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/clusters/{id}/network/queries/{queryId}][%d] getClusterNetworkConfigurationQueryResponseInternalServerError  %+v", 500, o.Payload)
}

func (o *GetClusterNetworkConfigurationQueryResponseInternalServerError) String() string {
	return fmt.Sprintf("[GET /v1/clusters/{id}/network/queries/{queryId}][%d] getClusterNetworkConfigurationQueryResponseInternalServerError  %+v", 500, o.Payload)
}

func (o *GetClusterNetworkConfigurationQueryResponseInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetClusterNetworkConfigurationQueryResponseInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
