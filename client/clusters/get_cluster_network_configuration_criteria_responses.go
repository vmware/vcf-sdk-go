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

// GetClusterNetworkConfigurationCriteriaReader is a Reader for the GetClusterNetworkConfigurationCriteria structure.
type GetClusterNetworkConfigurationCriteriaReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetClusterNetworkConfigurationCriteriaReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetClusterNetworkConfigurationCriteriaOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetClusterNetworkConfigurationCriteriaBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetClusterNetworkConfigurationCriteriaInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /v1/clusters/{id}/network/criteria] getClusterNetworkConfigurationCriteria", response, response.Code())
	}
}

// NewGetClusterNetworkConfigurationCriteriaOK creates a GetClusterNetworkConfigurationCriteriaOK with default headers values
func NewGetClusterNetworkConfigurationCriteriaOK() *GetClusterNetworkConfigurationCriteriaOK {
	return &GetClusterNetworkConfigurationCriteriaOK{}
}

/*
GetClusterNetworkConfigurationCriteriaOK describes a response with status code 200, with default header values.

Ok
*/
type GetClusterNetworkConfigurationCriteriaOK struct {
	Payload *models.PageOfClusterNetworkConfigurationCriterion
}

// IsSuccess returns true when this get cluster network configuration criteria o k response has a 2xx status code
func (o *GetClusterNetworkConfigurationCriteriaOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get cluster network configuration criteria o k response has a 3xx status code
func (o *GetClusterNetworkConfigurationCriteriaOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get cluster network configuration criteria o k response has a 4xx status code
func (o *GetClusterNetworkConfigurationCriteriaOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get cluster network configuration criteria o k response has a 5xx status code
func (o *GetClusterNetworkConfigurationCriteriaOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get cluster network configuration criteria o k response a status code equal to that given
func (o *GetClusterNetworkConfigurationCriteriaOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get cluster network configuration criteria o k response
func (o *GetClusterNetworkConfigurationCriteriaOK) Code() int {
	return 200
}

func (o *GetClusterNetworkConfigurationCriteriaOK) Error() string {
	return fmt.Sprintf("[GET /v1/clusters/{id}/network/criteria][%d] getClusterNetworkConfigurationCriteriaOK  %+v", 200, o.Payload)
}

func (o *GetClusterNetworkConfigurationCriteriaOK) String() string {
	return fmt.Sprintf("[GET /v1/clusters/{id}/network/criteria][%d] getClusterNetworkConfigurationCriteriaOK  %+v", 200, o.Payload)
}

func (o *GetClusterNetworkConfigurationCriteriaOK) GetPayload() *models.PageOfClusterNetworkConfigurationCriterion {
	return o.Payload
}

func (o *GetClusterNetworkConfigurationCriteriaOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PageOfClusterNetworkConfigurationCriterion)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetClusterNetworkConfigurationCriteriaBadRequest creates a GetClusterNetworkConfigurationCriteriaBadRequest with default headers values
func NewGetClusterNetworkConfigurationCriteriaBadRequest() *GetClusterNetworkConfigurationCriteriaBadRequest {
	return &GetClusterNetworkConfigurationCriteriaBadRequest{}
}

/*
GetClusterNetworkConfigurationCriteriaBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetClusterNetworkConfigurationCriteriaBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this get cluster network configuration criteria bad request response has a 2xx status code
func (o *GetClusterNetworkConfigurationCriteriaBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get cluster network configuration criteria bad request response has a 3xx status code
func (o *GetClusterNetworkConfigurationCriteriaBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get cluster network configuration criteria bad request response has a 4xx status code
func (o *GetClusterNetworkConfigurationCriteriaBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get cluster network configuration criteria bad request response has a 5xx status code
func (o *GetClusterNetworkConfigurationCriteriaBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get cluster network configuration criteria bad request response a status code equal to that given
func (o *GetClusterNetworkConfigurationCriteriaBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get cluster network configuration criteria bad request response
func (o *GetClusterNetworkConfigurationCriteriaBadRequest) Code() int {
	return 400
}

func (o *GetClusterNetworkConfigurationCriteriaBadRequest) Error() string {
	return fmt.Sprintf("[GET /v1/clusters/{id}/network/criteria][%d] getClusterNetworkConfigurationCriteriaBadRequest  %+v", 400, o.Payload)
}

func (o *GetClusterNetworkConfigurationCriteriaBadRequest) String() string {
	return fmt.Sprintf("[GET /v1/clusters/{id}/network/criteria][%d] getClusterNetworkConfigurationCriteriaBadRequest  %+v", 400, o.Payload)
}

func (o *GetClusterNetworkConfigurationCriteriaBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetClusterNetworkConfigurationCriteriaBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetClusterNetworkConfigurationCriteriaInternalServerError creates a GetClusterNetworkConfigurationCriteriaInternalServerError with default headers values
func NewGetClusterNetworkConfigurationCriteriaInternalServerError() *GetClusterNetworkConfigurationCriteriaInternalServerError {
	return &GetClusterNetworkConfigurationCriteriaInternalServerError{}
}

/*
GetClusterNetworkConfigurationCriteriaInternalServerError describes a response with status code 500, with default header values.

InternalServerError
*/
type GetClusterNetworkConfigurationCriteriaInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this get cluster network configuration criteria internal server error response has a 2xx status code
func (o *GetClusterNetworkConfigurationCriteriaInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get cluster network configuration criteria internal server error response has a 3xx status code
func (o *GetClusterNetworkConfigurationCriteriaInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get cluster network configuration criteria internal server error response has a 4xx status code
func (o *GetClusterNetworkConfigurationCriteriaInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get cluster network configuration criteria internal server error response has a 5xx status code
func (o *GetClusterNetworkConfigurationCriteriaInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get cluster network configuration criteria internal server error response a status code equal to that given
func (o *GetClusterNetworkConfigurationCriteriaInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get cluster network configuration criteria internal server error response
func (o *GetClusterNetworkConfigurationCriteriaInternalServerError) Code() int {
	return 500
}

func (o *GetClusterNetworkConfigurationCriteriaInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/clusters/{id}/network/criteria][%d] getClusterNetworkConfigurationCriteriaInternalServerError  %+v", 500, o.Payload)
}

func (o *GetClusterNetworkConfigurationCriteriaInternalServerError) String() string {
	return fmt.Sprintf("[GET /v1/clusters/{id}/network/criteria][%d] getClusterNetworkConfigurationCriteriaInternalServerError  %+v", 500, o.Payload)
}

func (o *GetClusterNetworkConfigurationCriteriaInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetClusterNetworkConfigurationCriteriaInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
