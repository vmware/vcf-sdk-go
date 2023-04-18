// Code generated by go-swagger; DO NOT EDIT.

// Copyright © 2023 VMware, Inc. All Rights Reserved.
// Code licenced under: BSD-2

package domains

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/vmware/vcf-sdk-go/models"
)

// DeleteDomainReader is a Reader for the DeleteDomain structure.
type DeleteDomainReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteDomainReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteDomainOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewDeleteDomainAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteDomainBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteDomainNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteDomainInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteDomainOK creates a DeleteDomainOK with default headers values
func NewDeleteDomainOK() *DeleteDomainOK {
	return &DeleteDomainOK{}
}

/*
DeleteDomainOK describes a response with status code 200, with default header values.

OK
*/
type DeleteDomainOK struct {
	Payload *models.Task
}

// IsSuccess returns true when this delete domain o k response has a 2xx status code
func (o *DeleteDomainOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete domain o k response has a 3xx status code
func (o *DeleteDomainOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete domain o k response has a 4xx status code
func (o *DeleteDomainOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete domain o k response has a 5xx status code
func (o *DeleteDomainOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete domain o k response a status code equal to that given
func (o *DeleteDomainOK) IsCode(code int) bool {
	return code == 200
}

func (o *DeleteDomainOK) Error() string {
	return fmt.Sprintf("[DELETE /v1/domains/{id}][%d] deleteDomainOK  %+v", 200, o.Payload)
}

func (o *DeleteDomainOK) String() string {
	return fmt.Sprintf("[DELETE /v1/domains/{id}][%d] deleteDomainOK  %+v", 200, o.Payload)
}

func (o *DeleteDomainOK) GetPayload() *models.Task {
	return o.Payload
}

func (o *DeleteDomainOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Task)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteDomainAccepted creates a DeleteDomainAccepted with default headers values
func NewDeleteDomainAccepted() *DeleteDomainAccepted {
	return &DeleteDomainAccepted{}
}

/*
DeleteDomainAccepted describes a response with status code 202, with default header values.

Accepted
*/
type DeleteDomainAccepted struct {
	Payload *models.Task
}

// IsSuccess returns true when this delete domain accepted response has a 2xx status code
func (o *DeleteDomainAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete domain accepted response has a 3xx status code
func (o *DeleteDomainAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete domain accepted response has a 4xx status code
func (o *DeleteDomainAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete domain accepted response has a 5xx status code
func (o *DeleteDomainAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this delete domain accepted response a status code equal to that given
func (o *DeleteDomainAccepted) IsCode(code int) bool {
	return code == 202
}

func (o *DeleteDomainAccepted) Error() string {
	return fmt.Sprintf("[DELETE /v1/domains/{id}][%d] deleteDomainAccepted  %+v", 202, o.Payload)
}

func (o *DeleteDomainAccepted) String() string {
	return fmt.Sprintf("[DELETE /v1/domains/{id}][%d] deleteDomainAccepted  %+v", 202, o.Payload)
}

func (o *DeleteDomainAccepted) GetPayload() *models.Task {
	return o.Payload
}

func (o *DeleteDomainAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Task)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteDomainBadRequest creates a DeleteDomainBadRequest with default headers values
func NewDeleteDomainBadRequest() *DeleteDomainBadRequest {
	return &DeleteDomainBadRequest{}
}

/*
DeleteDomainBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type DeleteDomainBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this delete domain bad request response has a 2xx status code
func (o *DeleteDomainBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete domain bad request response has a 3xx status code
func (o *DeleteDomainBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete domain bad request response has a 4xx status code
func (o *DeleteDomainBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete domain bad request response has a 5xx status code
func (o *DeleteDomainBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this delete domain bad request response a status code equal to that given
func (o *DeleteDomainBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *DeleteDomainBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /v1/domains/{id}][%d] deleteDomainBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteDomainBadRequest) String() string {
	return fmt.Sprintf("[DELETE /v1/domains/{id}][%d] deleteDomainBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteDomainBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteDomainBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteDomainNotFound creates a DeleteDomainNotFound with default headers values
func NewDeleteDomainNotFound() *DeleteDomainNotFound {
	return &DeleteDomainNotFound{}
}

/*
DeleteDomainNotFound describes a response with status code 404, with default header values.

Not Found
*/
type DeleteDomainNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this delete domain not found response has a 2xx status code
func (o *DeleteDomainNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete domain not found response has a 3xx status code
func (o *DeleteDomainNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete domain not found response has a 4xx status code
func (o *DeleteDomainNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete domain not found response has a 5xx status code
func (o *DeleteDomainNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this delete domain not found response a status code equal to that given
func (o *DeleteDomainNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *DeleteDomainNotFound) Error() string {
	return fmt.Sprintf("[DELETE /v1/domains/{id}][%d] deleteDomainNotFound  %+v", 404, o.Payload)
}

func (o *DeleteDomainNotFound) String() string {
	return fmt.Sprintf("[DELETE /v1/domains/{id}][%d] deleteDomainNotFound  %+v", 404, o.Payload)
}

func (o *DeleteDomainNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteDomainNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteDomainInternalServerError creates a DeleteDomainInternalServerError with default headers values
func NewDeleteDomainInternalServerError() *DeleteDomainInternalServerError {
	return &DeleteDomainInternalServerError{}
}

/*
DeleteDomainInternalServerError describes a response with status code 500, with default header values.

InternalServerError
*/
type DeleteDomainInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this delete domain internal server error response has a 2xx status code
func (o *DeleteDomainInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete domain internal server error response has a 3xx status code
func (o *DeleteDomainInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete domain internal server error response has a 4xx status code
func (o *DeleteDomainInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete domain internal server error response has a 5xx status code
func (o *DeleteDomainInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this delete domain internal server error response a status code equal to that given
func (o *DeleteDomainInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *DeleteDomainInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /v1/domains/{id}][%d] deleteDomainInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteDomainInternalServerError) String() string {
	return fmt.Sprintf("[DELETE /v1/domains/{id}][%d] deleteDomainInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteDomainInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteDomainInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
