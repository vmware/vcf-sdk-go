// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2023 VMware, Inc.
// SPDX-License-Identifier: BSD-2-Clause

package sddc

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/vmware/vcf-sdk-go/models"
)

// GetBringupTaskByIDReader is a Reader for the GetBringupTaskByID structure.
type GetBringupTaskByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetBringupTaskByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetBringupTaskByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetBringupTaskByIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetBringupTaskByIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetBringupTaskByIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /v1/sddcs/{id}] getBringupTaskByID", response, response.Code())
	}
}

// NewGetBringupTaskByIDOK creates a GetBringupTaskByIDOK with default headers values
func NewGetBringupTaskByIDOK() *GetBringupTaskByIDOK {
	return &GetBringupTaskByIDOK{}
}

/*
GetBringupTaskByIDOK describes a response with status code 200, with default header values.

OK
*/
type GetBringupTaskByIDOK struct {
	Payload *models.SDDCTask
}

// IsSuccess returns true when this get bringup task by Id o k response has a 2xx status code
func (o *GetBringupTaskByIDOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get bringup task by Id o k response has a 3xx status code
func (o *GetBringupTaskByIDOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get bringup task by Id o k response has a 4xx status code
func (o *GetBringupTaskByIDOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get bringup task by Id o k response has a 5xx status code
func (o *GetBringupTaskByIDOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get bringup task by Id o k response a status code equal to that given
func (o *GetBringupTaskByIDOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get bringup task by Id o k response
func (o *GetBringupTaskByIDOK) Code() int {
	return 200
}

func (o *GetBringupTaskByIDOK) Error() string {
	return fmt.Sprintf("[GET /v1/sddcs/{id}][%d] getBringupTaskByIdOK  %+v", 200, o.Payload)
}

func (o *GetBringupTaskByIDOK) String() string {
	return fmt.Sprintf("[GET /v1/sddcs/{id}][%d] getBringupTaskByIdOK  %+v", 200, o.Payload)
}

func (o *GetBringupTaskByIDOK) GetPayload() *models.SDDCTask {
	return o.Payload
}

func (o *GetBringupTaskByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SDDCTask)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetBringupTaskByIDBadRequest creates a GetBringupTaskByIDBadRequest with default headers values
func NewGetBringupTaskByIDBadRequest() *GetBringupTaskByIDBadRequest {
	return &GetBringupTaskByIDBadRequest{}
}

/*
GetBringupTaskByIDBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetBringupTaskByIDBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this get bringup task by Id bad request response has a 2xx status code
func (o *GetBringupTaskByIDBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get bringup task by Id bad request response has a 3xx status code
func (o *GetBringupTaskByIDBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get bringup task by Id bad request response has a 4xx status code
func (o *GetBringupTaskByIDBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get bringup task by Id bad request response has a 5xx status code
func (o *GetBringupTaskByIDBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get bringup task by Id bad request response a status code equal to that given
func (o *GetBringupTaskByIDBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get bringup task by Id bad request response
func (o *GetBringupTaskByIDBadRequest) Code() int {
	return 400
}

func (o *GetBringupTaskByIDBadRequest) Error() string {
	return fmt.Sprintf("[GET /v1/sddcs/{id}][%d] getBringupTaskByIdBadRequest  %+v", 400, o.Payload)
}

func (o *GetBringupTaskByIDBadRequest) String() string {
	return fmt.Sprintf("[GET /v1/sddcs/{id}][%d] getBringupTaskByIdBadRequest  %+v", 400, o.Payload)
}

func (o *GetBringupTaskByIDBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetBringupTaskByIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetBringupTaskByIDNotFound creates a GetBringupTaskByIDNotFound with default headers values
func NewGetBringupTaskByIDNotFound() *GetBringupTaskByIDNotFound {
	return &GetBringupTaskByIDNotFound{}
}

/*
GetBringupTaskByIDNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetBringupTaskByIDNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this get bringup task by Id not found response has a 2xx status code
func (o *GetBringupTaskByIDNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get bringup task by Id not found response has a 3xx status code
func (o *GetBringupTaskByIDNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get bringup task by Id not found response has a 4xx status code
func (o *GetBringupTaskByIDNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get bringup task by Id not found response has a 5xx status code
func (o *GetBringupTaskByIDNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get bringup task by Id not found response a status code equal to that given
func (o *GetBringupTaskByIDNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get bringup task by Id not found response
func (o *GetBringupTaskByIDNotFound) Code() int {
	return 404
}

func (o *GetBringupTaskByIDNotFound) Error() string {
	return fmt.Sprintf("[GET /v1/sddcs/{id}][%d] getBringupTaskByIdNotFound  %+v", 404, o.Payload)
}

func (o *GetBringupTaskByIDNotFound) String() string {
	return fmt.Sprintf("[GET /v1/sddcs/{id}][%d] getBringupTaskByIdNotFound  %+v", 404, o.Payload)
}

func (o *GetBringupTaskByIDNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetBringupTaskByIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetBringupTaskByIDInternalServerError creates a GetBringupTaskByIDInternalServerError with default headers values
func NewGetBringupTaskByIDInternalServerError() *GetBringupTaskByIDInternalServerError {
	return &GetBringupTaskByIDInternalServerError{}
}

/*
GetBringupTaskByIDInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetBringupTaskByIDInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this get bringup task by Id internal server error response has a 2xx status code
func (o *GetBringupTaskByIDInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get bringup task by Id internal server error response has a 3xx status code
func (o *GetBringupTaskByIDInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get bringup task by Id internal server error response has a 4xx status code
func (o *GetBringupTaskByIDInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get bringup task by Id internal server error response has a 5xx status code
func (o *GetBringupTaskByIDInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get bringup task by Id internal server error response a status code equal to that given
func (o *GetBringupTaskByIDInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get bringup task by Id internal server error response
func (o *GetBringupTaskByIDInternalServerError) Code() int {
	return 500
}

func (o *GetBringupTaskByIDInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/sddcs/{id}][%d] getBringupTaskByIdInternalServerError  %+v", 500, o.Payload)
}

func (o *GetBringupTaskByIDInternalServerError) String() string {
	return fmt.Sprintf("[GET /v1/sddcs/{id}][%d] getBringupTaskByIdInternalServerError  %+v", 500, o.Payload)
}

func (o *GetBringupTaskByIDInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetBringupTaskByIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
