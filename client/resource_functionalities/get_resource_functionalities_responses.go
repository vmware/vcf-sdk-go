// Code generated by go-swagger; DO NOT EDIT.

// Copyright © 2023 VMware, Inc. All Rights Reserved.
// Code licenced under: BSD-2

package resource_functionalities

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/vmware/vcf-sdk-go/models"
)

// GETResourceFunctionalitiesReader is a Reader for the GETResourceFunctionalities structure.
type GETResourceFunctionalitiesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GETResourceFunctionalitiesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGETResourceFunctionalitiesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGETResourceFunctionalitiesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGETResourceFunctionalitiesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGETResourceFunctionalitiesOK creates a GETResourceFunctionalitiesOK with default headers values
func NewGETResourceFunctionalitiesOK() *GETResourceFunctionalitiesOK {
	return &GETResourceFunctionalitiesOK{}
}

/*
GETResourceFunctionalitiesOK describes a response with status code 200, with default header values.

OK
*/
type GETResourceFunctionalitiesOK struct {
	Payload *models.PageOfResourceFunctionalities
}

// IsSuccess returns true when this get resource functionalities o k response has a 2xx status code
func (o *GETResourceFunctionalitiesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get resource functionalities o k response has a 3xx status code
func (o *GETResourceFunctionalitiesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get resource functionalities o k response has a 4xx status code
func (o *GETResourceFunctionalitiesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get resource functionalities o k response has a 5xx status code
func (o *GETResourceFunctionalitiesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get resource functionalities o k response a status code equal to that given
func (o *GETResourceFunctionalitiesOK) IsCode(code int) bool {
	return code == 200
}

func (o *GETResourceFunctionalitiesOK) Error() string {
	return fmt.Sprintf("[GET /v1/resource-functionalities][%d] getResourceFunctionalitiesOK  %+v", 200, o.Payload)
}

func (o *GETResourceFunctionalitiesOK) String() string {
	return fmt.Sprintf("[GET /v1/resource-functionalities][%d] getResourceFunctionalitiesOK  %+v", 200, o.Payload)
}

func (o *GETResourceFunctionalitiesOK) GetPayload() *models.PageOfResourceFunctionalities {
	return o.Payload
}

func (o *GETResourceFunctionalitiesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PageOfResourceFunctionalities)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGETResourceFunctionalitiesBadRequest creates a GETResourceFunctionalitiesBadRequest with default headers values
func NewGETResourceFunctionalitiesBadRequest() *GETResourceFunctionalitiesBadRequest {
	return &GETResourceFunctionalitiesBadRequest{}
}

/*
GETResourceFunctionalitiesBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GETResourceFunctionalitiesBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this get resource functionalities bad request response has a 2xx status code
func (o *GETResourceFunctionalitiesBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get resource functionalities bad request response has a 3xx status code
func (o *GETResourceFunctionalitiesBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get resource functionalities bad request response has a 4xx status code
func (o *GETResourceFunctionalitiesBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get resource functionalities bad request response has a 5xx status code
func (o *GETResourceFunctionalitiesBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get resource functionalities bad request response a status code equal to that given
func (o *GETResourceFunctionalitiesBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GETResourceFunctionalitiesBadRequest) Error() string {
	return fmt.Sprintf("[GET /v1/resource-functionalities][%d] getResourceFunctionalitiesBadRequest  %+v", 400, o.Payload)
}

func (o *GETResourceFunctionalitiesBadRequest) String() string {
	return fmt.Sprintf("[GET /v1/resource-functionalities][%d] getResourceFunctionalitiesBadRequest  %+v", 400, o.Payload)
}

func (o *GETResourceFunctionalitiesBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *GETResourceFunctionalitiesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGETResourceFunctionalitiesInternalServerError creates a GETResourceFunctionalitiesInternalServerError with default headers values
func NewGETResourceFunctionalitiesInternalServerError() *GETResourceFunctionalitiesInternalServerError {
	return &GETResourceFunctionalitiesInternalServerError{}
}

/*
GETResourceFunctionalitiesInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GETResourceFunctionalitiesInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this get resource functionalities internal server error response has a 2xx status code
func (o *GETResourceFunctionalitiesInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get resource functionalities internal server error response has a 3xx status code
func (o *GETResourceFunctionalitiesInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get resource functionalities internal server error response has a 4xx status code
func (o *GETResourceFunctionalitiesInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get resource functionalities internal server error response has a 5xx status code
func (o *GETResourceFunctionalitiesInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get resource functionalities internal server error response a status code equal to that given
func (o *GETResourceFunctionalitiesInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GETResourceFunctionalitiesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/resource-functionalities][%d] getResourceFunctionalitiesInternalServerError  %+v", 500, o.Payload)
}

func (o *GETResourceFunctionalitiesInternalServerError) String() string {
	return fmt.Sprintf("[GET /v1/resource-functionalities][%d] getResourceFunctionalitiesInternalServerError  %+v", 500, o.Payload)
}

func (o *GETResourceFunctionalitiesInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *GETResourceFunctionalitiesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
