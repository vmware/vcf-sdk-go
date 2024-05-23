// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2023 VMware, Inc.
// SPDX-License-Identifier: BSD-2-Clause

package network_pools

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/vmware/vcf-sdk-go/models"
)

// DeleteIPPoolFromNetworkOfNetworkPoolReader is a Reader for the DeleteIPPoolFromNetworkOfNetworkPool structure.
type DeleteIPPoolFromNetworkOfNetworkPoolReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteIPPoolFromNetworkOfNetworkPoolReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteIPPoolFromNetworkOfNetworkPoolNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteIPPoolFromNetworkOfNetworkPoolBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteIPPoolFromNetworkOfNetworkPoolNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteIPPoolFromNetworkOfNetworkPoolInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[DELETE /v1/network-pools/{id}/networks/{networkId}/ip-pools] deleteIpPoolFromNetworkOfNetworkPool", response, response.Code())
	}
}

// NewDeleteIPPoolFromNetworkOfNetworkPoolNoContent creates a DeleteIPPoolFromNetworkOfNetworkPoolNoContent with default headers values
func NewDeleteIPPoolFromNetworkOfNetworkPoolNoContent() *DeleteIPPoolFromNetworkOfNetworkPoolNoContent {
	return &DeleteIPPoolFromNetworkOfNetworkPoolNoContent{}
}

/*
DeleteIPPoolFromNetworkOfNetworkPoolNoContent describes a response with status code 204, with default header values.

No Content
*/
type DeleteIPPoolFromNetworkOfNetworkPoolNoContent struct {
}

// IsSuccess returns true when this delete Ip pool from network of network pool no content response has a 2xx status code
func (o *DeleteIPPoolFromNetworkOfNetworkPoolNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete Ip pool from network of network pool no content response has a 3xx status code
func (o *DeleteIPPoolFromNetworkOfNetworkPoolNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete Ip pool from network of network pool no content response has a 4xx status code
func (o *DeleteIPPoolFromNetworkOfNetworkPoolNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete Ip pool from network of network pool no content response has a 5xx status code
func (o *DeleteIPPoolFromNetworkOfNetworkPoolNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this delete Ip pool from network of network pool no content response a status code equal to that given
func (o *DeleteIPPoolFromNetworkOfNetworkPoolNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the delete Ip pool from network of network pool no content response
func (o *DeleteIPPoolFromNetworkOfNetworkPoolNoContent) Code() int {
	return 204
}

func (o *DeleteIPPoolFromNetworkOfNetworkPoolNoContent) Error() string {
	return fmt.Sprintf("[DELETE /v1/network-pools/{id}/networks/{networkId}/ip-pools][%d] deleteIpPoolFromNetworkOfNetworkPoolNoContent ", 204)
}

func (o *DeleteIPPoolFromNetworkOfNetworkPoolNoContent) String() string {
	return fmt.Sprintf("[DELETE /v1/network-pools/{id}/networks/{networkId}/ip-pools][%d] deleteIpPoolFromNetworkOfNetworkPoolNoContent ", 204)
}

func (o *DeleteIPPoolFromNetworkOfNetworkPoolNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteIPPoolFromNetworkOfNetworkPoolBadRequest creates a DeleteIPPoolFromNetworkOfNetworkPoolBadRequest with default headers values
func NewDeleteIPPoolFromNetworkOfNetworkPoolBadRequest() *DeleteIPPoolFromNetworkOfNetworkPoolBadRequest {
	return &DeleteIPPoolFromNetworkOfNetworkPoolBadRequest{}
}

/*
DeleteIPPoolFromNetworkOfNetworkPoolBadRequest describes a response with status code 400, with default header values.

Errors due to network/networkpool validations failures
*/
type DeleteIPPoolFromNetworkOfNetworkPoolBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this delete Ip pool from network of network pool bad request response has a 2xx status code
func (o *DeleteIPPoolFromNetworkOfNetworkPoolBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete Ip pool from network of network pool bad request response has a 3xx status code
func (o *DeleteIPPoolFromNetworkOfNetworkPoolBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete Ip pool from network of network pool bad request response has a 4xx status code
func (o *DeleteIPPoolFromNetworkOfNetworkPoolBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete Ip pool from network of network pool bad request response has a 5xx status code
func (o *DeleteIPPoolFromNetworkOfNetworkPoolBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this delete Ip pool from network of network pool bad request response a status code equal to that given
func (o *DeleteIPPoolFromNetworkOfNetworkPoolBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the delete Ip pool from network of network pool bad request response
func (o *DeleteIPPoolFromNetworkOfNetworkPoolBadRequest) Code() int {
	return 400
}

func (o *DeleteIPPoolFromNetworkOfNetworkPoolBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /v1/network-pools/{id}/networks/{networkId}/ip-pools][%d] deleteIpPoolFromNetworkOfNetworkPoolBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteIPPoolFromNetworkOfNetworkPoolBadRequest) String() string {
	return fmt.Sprintf("[DELETE /v1/network-pools/{id}/networks/{networkId}/ip-pools][%d] deleteIpPoolFromNetworkOfNetworkPoolBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteIPPoolFromNetworkOfNetworkPoolBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteIPPoolFromNetworkOfNetworkPoolBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteIPPoolFromNetworkOfNetworkPoolNotFound creates a DeleteIPPoolFromNetworkOfNetworkPoolNotFound with default headers values
func NewDeleteIPPoolFromNetworkOfNetworkPoolNotFound() *DeleteIPPoolFromNetworkOfNetworkPoolNotFound {
	return &DeleteIPPoolFromNetworkOfNetworkPoolNotFound{}
}

/*
DeleteIPPoolFromNetworkOfNetworkPoolNotFound describes a response with status code 404, with default header values.

Either network or Network pool not found
*/
type DeleteIPPoolFromNetworkOfNetworkPoolNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this delete Ip pool from network of network pool not found response has a 2xx status code
func (o *DeleteIPPoolFromNetworkOfNetworkPoolNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete Ip pool from network of network pool not found response has a 3xx status code
func (o *DeleteIPPoolFromNetworkOfNetworkPoolNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete Ip pool from network of network pool not found response has a 4xx status code
func (o *DeleteIPPoolFromNetworkOfNetworkPoolNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete Ip pool from network of network pool not found response has a 5xx status code
func (o *DeleteIPPoolFromNetworkOfNetworkPoolNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this delete Ip pool from network of network pool not found response a status code equal to that given
func (o *DeleteIPPoolFromNetworkOfNetworkPoolNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the delete Ip pool from network of network pool not found response
func (o *DeleteIPPoolFromNetworkOfNetworkPoolNotFound) Code() int {
	return 404
}

func (o *DeleteIPPoolFromNetworkOfNetworkPoolNotFound) Error() string {
	return fmt.Sprintf("[DELETE /v1/network-pools/{id}/networks/{networkId}/ip-pools][%d] deleteIpPoolFromNetworkOfNetworkPoolNotFound  %+v", 404, o.Payload)
}

func (o *DeleteIPPoolFromNetworkOfNetworkPoolNotFound) String() string {
	return fmt.Sprintf("[DELETE /v1/network-pools/{id}/networks/{networkId}/ip-pools][%d] deleteIpPoolFromNetworkOfNetworkPoolNotFound  %+v", 404, o.Payload)
}

func (o *DeleteIPPoolFromNetworkOfNetworkPoolNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteIPPoolFromNetworkOfNetworkPoolNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteIPPoolFromNetworkOfNetworkPoolInternalServerError creates a DeleteIPPoolFromNetworkOfNetworkPoolInternalServerError with default headers values
func NewDeleteIPPoolFromNetworkOfNetworkPoolInternalServerError() *DeleteIPPoolFromNetworkOfNetworkPoolInternalServerError {
	return &DeleteIPPoolFromNetworkOfNetworkPoolInternalServerError{}
}

/*
DeleteIPPoolFromNetworkOfNetworkPoolInternalServerError describes a response with status code 500, with default header values.

Unexpected error
*/
type DeleteIPPoolFromNetworkOfNetworkPoolInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this delete Ip pool from network of network pool internal server error response has a 2xx status code
func (o *DeleteIPPoolFromNetworkOfNetworkPoolInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete Ip pool from network of network pool internal server error response has a 3xx status code
func (o *DeleteIPPoolFromNetworkOfNetworkPoolInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete Ip pool from network of network pool internal server error response has a 4xx status code
func (o *DeleteIPPoolFromNetworkOfNetworkPoolInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete Ip pool from network of network pool internal server error response has a 5xx status code
func (o *DeleteIPPoolFromNetworkOfNetworkPoolInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this delete Ip pool from network of network pool internal server error response a status code equal to that given
func (o *DeleteIPPoolFromNetworkOfNetworkPoolInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the delete Ip pool from network of network pool internal server error response
func (o *DeleteIPPoolFromNetworkOfNetworkPoolInternalServerError) Code() int {
	return 500
}

func (o *DeleteIPPoolFromNetworkOfNetworkPoolInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /v1/network-pools/{id}/networks/{networkId}/ip-pools][%d] deleteIpPoolFromNetworkOfNetworkPoolInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteIPPoolFromNetworkOfNetworkPoolInternalServerError) String() string {
	return fmt.Sprintf("[DELETE /v1/network-pools/{id}/networks/{networkId}/ip-pools][%d] deleteIpPoolFromNetworkOfNetworkPoolInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteIPPoolFromNetworkOfNetworkPoolInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteIPPoolFromNetworkOfNetworkPoolInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
