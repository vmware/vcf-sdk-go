// Code generated by go-swagger; DO NOT EDIT.

// Copyright © 2023 VMware, Inc. All Rights Reserved.
// Code licenced under: BSD-2

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

// GETClusterTagManagerURLReader is a Reader for the GETClusterTagManagerURL structure.
type GETClusterTagManagerURLReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GETClusterTagManagerURLReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGETClusterTagManagerURLOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGETClusterTagManagerURLBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGETClusterTagManagerURLInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGETClusterTagManagerURLOK creates a GETClusterTagManagerURLOK with default headers values
func NewGETClusterTagManagerURLOK() *GETClusterTagManagerURLOK {
	return &GETClusterTagManagerURLOK{}
}

/*
GETClusterTagManagerURLOK describes a response with status code 200, with default header values.

Ok
*/
type GETClusterTagManagerURLOK struct {
	Payload *models.TagManagerModel
}

// IsSuccess returns true when this get cluster tag manager Url o k response has a 2xx status code
func (o *GETClusterTagManagerURLOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get cluster tag manager Url o k response has a 3xx status code
func (o *GETClusterTagManagerURLOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get cluster tag manager Url o k response has a 4xx status code
func (o *GETClusterTagManagerURLOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get cluster tag manager Url o k response has a 5xx status code
func (o *GETClusterTagManagerURLOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get cluster tag manager Url o k response a status code equal to that given
func (o *GETClusterTagManagerURLOK) IsCode(code int) bool {
	return code == 200
}

func (o *GETClusterTagManagerURLOK) Error() string {
	return fmt.Sprintf("[GET /v1/clusters/{id}/tags/tag-manager][%d] getClusterTagManagerUrlOK  %+v", 200, o.Payload)
}

func (o *GETClusterTagManagerURLOK) String() string {
	return fmt.Sprintf("[GET /v1/clusters/{id}/tags/tag-manager][%d] getClusterTagManagerUrlOK  %+v", 200, o.Payload)
}

func (o *GETClusterTagManagerURLOK) GetPayload() *models.TagManagerModel {
	return o.Payload
}

func (o *GETClusterTagManagerURLOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TagManagerModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGETClusterTagManagerURLBadRequest creates a GETClusterTagManagerURLBadRequest with default headers values
func NewGETClusterTagManagerURLBadRequest() *GETClusterTagManagerURLBadRequest {
	return &GETClusterTagManagerURLBadRequest{}
}

/*
GETClusterTagManagerURLBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GETClusterTagManagerURLBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this get cluster tag manager Url bad request response has a 2xx status code
func (o *GETClusterTagManagerURLBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get cluster tag manager Url bad request response has a 3xx status code
func (o *GETClusterTagManagerURLBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get cluster tag manager Url bad request response has a 4xx status code
func (o *GETClusterTagManagerURLBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get cluster tag manager Url bad request response has a 5xx status code
func (o *GETClusterTagManagerURLBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get cluster tag manager Url bad request response a status code equal to that given
func (o *GETClusterTagManagerURLBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GETClusterTagManagerURLBadRequest) Error() string {
	return fmt.Sprintf("[GET /v1/clusters/{id}/tags/tag-manager][%d] getClusterTagManagerUrlBadRequest  %+v", 400, o.Payload)
}

func (o *GETClusterTagManagerURLBadRequest) String() string {
	return fmt.Sprintf("[GET /v1/clusters/{id}/tags/tag-manager][%d] getClusterTagManagerUrlBadRequest  %+v", 400, o.Payload)
}

func (o *GETClusterTagManagerURLBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *GETClusterTagManagerURLBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGETClusterTagManagerURLInternalServerError creates a GETClusterTagManagerURLInternalServerError with default headers values
func NewGETClusterTagManagerURLInternalServerError() *GETClusterTagManagerURLInternalServerError {
	return &GETClusterTagManagerURLInternalServerError{}
}

/*
GETClusterTagManagerURLInternalServerError describes a response with status code 500, with default header values.

InternalServerError
*/
type GETClusterTagManagerURLInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this get cluster tag manager Url internal server error response has a 2xx status code
func (o *GETClusterTagManagerURLInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get cluster tag manager Url internal server error response has a 3xx status code
func (o *GETClusterTagManagerURLInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get cluster tag manager Url internal server error response has a 4xx status code
func (o *GETClusterTagManagerURLInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get cluster tag manager Url internal server error response has a 5xx status code
func (o *GETClusterTagManagerURLInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get cluster tag manager Url internal server error response a status code equal to that given
func (o *GETClusterTagManagerURLInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GETClusterTagManagerURLInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/clusters/{id}/tags/tag-manager][%d] getClusterTagManagerUrlInternalServerError  %+v", 500, o.Payload)
}

func (o *GETClusterTagManagerURLInternalServerError) String() string {
	return fmt.Sprintf("[GET /v1/clusters/{id}/tags/tag-manager][%d] getClusterTagManagerUrlInternalServerError  %+v", 500, o.Payload)
}

func (o *GETClusterTagManagerURLInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *GETClusterTagManagerURLInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
