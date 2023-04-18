// Code generated by go-swagger; DO NOT EDIT.

// Copyright © 2023 VMware, Inc. All Rights Reserved.
// Code licenced under: BSD-2

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

// DeleteNetworkPoolReader is a Reader for the DeleteNetworkPool structure.
type DeleteNetworkPoolReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteNetworkPoolReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteNetworkPoolNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteNetworkPoolBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteNetworkPoolNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteNetworkPoolInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteNetworkPoolNoContent creates a DeleteNetworkPoolNoContent with default headers values
func NewDeleteNetworkPoolNoContent() *DeleteNetworkPoolNoContent {
	return &DeleteNetworkPoolNoContent{}
}

/*
DeleteNetworkPoolNoContent describes a response with status code 204, with default header values.

The specification of the deleted network pool
*/
type DeleteNetworkPoolNoContent struct {
}

// IsSuccess returns true when this delete network pool no content response has a 2xx status code
func (o *DeleteNetworkPoolNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete network pool no content response has a 3xx status code
func (o *DeleteNetworkPoolNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete network pool no content response has a 4xx status code
func (o *DeleteNetworkPoolNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete network pool no content response has a 5xx status code
func (o *DeleteNetworkPoolNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this delete network pool no content response a status code equal to that given
func (o *DeleteNetworkPoolNoContent) IsCode(code int) bool {
	return code == 204
}

func (o *DeleteNetworkPoolNoContent) Error() string {
	return fmt.Sprintf("[DELETE /v1/network-pools/{id}][%d] deleteNetworkPoolNoContent ", 204)
}

func (o *DeleteNetworkPoolNoContent) String() string {
	return fmt.Sprintf("[DELETE /v1/network-pools/{id}][%d] deleteNetworkPoolNoContent ", 204)
}

func (o *DeleteNetworkPoolNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteNetworkPoolBadRequest creates a DeleteNetworkPoolBadRequest with default headers values
func NewDeleteNetworkPoolBadRequest() *DeleteNetworkPoolBadRequest {
	return &DeleteNetworkPoolBadRequest{}
}

/*
DeleteNetworkPoolBadRequest describes a response with status code 400, with default header values.

Hosts are still associated with NetworkPool
*/
type DeleteNetworkPoolBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this delete network pool bad request response has a 2xx status code
func (o *DeleteNetworkPoolBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete network pool bad request response has a 3xx status code
func (o *DeleteNetworkPoolBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete network pool bad request response has a 4xx status code
func (o *DeleteNetworkPoolBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete network pool bad request response has a 5xx status code
func (o *DeleteNetworkPoolBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this delete network pool bad request response a status code equal to that given
func (o *DeleteNetworkPoolBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *DeleteNetworkPoolBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /v1/network-pools/{id}][%d] deleteNetworkPoolBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteNetworkPoolBadRequest) String() string {
	return fmt.Sprintf("[DELETE /v1/network-pools/{id}][%d] deleteNetworkPoolBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteNetworkPoolBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteNetworkPoolBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteNetworkPoolNotFound creates a DeleteNetworkPoolNotFound with default headers values
func NewDeleteNetworkPoolNotFound() *DeleteNetworkPoolNotFound {
	return &DeleteNetworkPoolNotFound{}
}

/*
DeleteNetworkPoolNotFound describes a response with status code 404, with default header values.

Referenced network pool not found
*/
type DeleteNetworkPoolNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this delete network pool not found response has a 2xx status code
func (o *DeleteNetworkPoolNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete network pool not found response has a 3xx status code
func (o *DeleteNetworkPoolNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete network pool not found response has a 4xx status code
func (o *DeleteNetworkPoolNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete network pool not found response has a 5xx status code
func (o *DeleteNetworkPoolNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this delete network pool not found response a status code equal to that given
func (o *DeleteNetworkPoolNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *DeleteNetworkPoolNotFound) Error() string {
	return fmt.Sprintf("[DELETE /v1/network-pools/{id}][%d] deleteNetworkPoolNotFound  %+v", 404, o.Payload)
}

func (o *DeleteNetworkPoolNotFound) String() string {
	return fmt.Sprintf("[DELETE /v1/network-pools/{id}][%d] deleteNetworkPoolNotFound  %+v", 404, o.Payload)
}

func (o *DeleteNetworkPoolNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteNetworkPoolNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteNetworkPoolInternalServerError creates a DeleteNetworkPoolInternalServerError with default headers values
func NewDeleteNetworkPoolInternalServerError() *DeleteNetworkPoolInternalServerError {
	return &DeleteNetworkPoolInternalServerError{}
}

/*
DeleteNetworkPoolInternalServerError describes a response with status code 500, with default header values.

Unexpected error
*/
type DeleteNetworkPoolInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this delete network pool internal server error response has a 2xx status code
func (o *DeleteNetworkPoolInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete network pool internal server error response has a 3xx status code
func (o *DeleteNetworkPoolInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete network pool internal server error response has a 4xx status code
func (o *DeleteNetworkPoolInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete network pool internal server error response has a 5xx status code
func (o *DeleteNetworkPoolInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this delete network pool internal server error response a status code equal to that given
func (o *DeleteNetworkPoolInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *DeleteNetworkPoolInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /v1/network-pools/{id}][%d] deleteNetworkPoolInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteNetworkPoolInternalServerError) String() string {
	return fmt.Sprintf("[DELETE /v1/network-pools/{id}][%d] deleteNetworkPoolInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteNetworkPoolInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteNetworkPoolInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
