// Code generated by go-swagger; DO NOT EDIT.

// Copyright © 2023 VMware, Inc. All Rights Reserved.
// Code licenced under: BSD-2

package certificates

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/vmware/vcf-sdk-go/models"
)

// GeneratesCSRsReader is a Reader for the GeneratesCSRs structure.
type GeneratesCSRsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GeneratesCSRsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGeneratesCSRsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewGeneratesCSRsAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGeneratesCSRsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGeneratesCSRsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewGeneratesCSRsConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGeneratesCSRsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGeneratesCSRsOK creates a GeneratesCSRsOK with default headers values
func NewGeneratesCSRsOK() *GeneratesCSRsOK {
	return &GeneratesCSRsOK{}
}

/*
GeneratesCSRsOK describes a response with status code 200, with default header values.

OK
*/
type GeneratesCSRsOK struct {
	Payload *models.Task
}

// IsSuccess returns true when this generates c s rs o k response has a 2xx status code
func (o *GeneratesCSRsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this generates c s rs o k response has a 3xx status code
func (o *GeneratesCSRsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this generates c s rs o k response has a 4xx status code
func (o *GeneratesCSRsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this generates c s rs o k response has a 5xx status code
func (o *GeneratesCSRsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this generates c s rs o k response a status code equal to that given
func (o *GeneratesCSRsOK) IsCode(code int) bool {
	return code == 200
}

func (o *GeneratesCSRsOK) Error() string {
	return fmt.Sprintf("[PUT /v1/domains/{domainName}/csrs][%d] generatesCSRsOK  %+v", 200, o.Payload)
}

func (o *GeneratesCSRsOK) String() string {
	return fmt.Sprintf("[PUT /v1/domains/{domainName}/csrs][%d] generatesCSRsOK  %+v", 200, o.Payload)
}

func (o *GeneratesCSRsOK) GetPayload() *models.Task {
	return o.Payload
}

func (o *GeneratesCSRsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Task)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGeneratesCSRsAccepted creates a GeneratesCSRsAccepted with default headers values
func NewGeneratesCSRsAccepted() *GeneratesCSRsAccepted {
	return &GeneratesCSRsAccepted{}
}

/*
GeneratesCSRsAccepted describes a response with status code 202, with default header values.

Accepted
*/
type GeneratesCSRsAccepted struct {
	Payload *models.Task
}

// IsSuccess returns true when this generates c s rs accepted response has a 2xx status code
func (o *GeneratesCSRsAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this generates c s rs accepted response has a 3xx status code
func (o *GeneratesCSRsAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this generates c s rs accepted response has a 4xx status code
func (o *GeneratesCSRsAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this generates c s rs accepted response has a 5xx status code
func (o *GeneratesCSRsAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this generates c s rs accepted response a status code equal to that given
func (o *GeneratesCSRsAccepted) IsCode(code int) bool {
	return code == 202
}

func (o *GeneratesCSRsAccepted) Error() string {
	return fmt.Sprintf("[PUT /v1/domains/{domainName}/csrs][%d] generatesCSRsAccepted  %+v", 202, o.Payload)
}

func (o *GeneratesCSRsAccepted) String() string {
	return fmt.Sprintf("[PUT /v1/domains/{domainName}/csrs][%d] generatesCSRsAccepted  %+v", 202, o.Payload)
}

func (o *GeneratesCSRsAccepted) GetPayload() *models.Task {
	return o.Payload
}

func (o *GeneratesCSRsAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Task)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGeneratesCSRsBadRequest creates a GeneratesCSRsBadRequest with default headers values
func NewGeneratesCSRsBadRequest() *GeneratesCSRsBadRequest {
	return &GeneratesCSRsBadRequest{}
}

/*
GeneratesCSRsBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type GeneratesCSRsBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this generates c s rs bad request response has a 2xx status code
func (o *GeneratesCSRsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this generates c s rs bad request response has a 3xx status code
func (o *GeneratesCSRsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this generates c s rs bad request response has a 4xx status code
func (o *GeneratesCSRsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this generates c s rs bad request response has a 5xx status code
func (o *GeneratesCSRsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this generates c s rs bad request response a status code equal to that given
func (o *GeneratesCSRsBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GeneratesCSRsBadRequest) Error() string {
	return fmt.Sprintf("[PUT /v1/domains/{domainName}/csrs][%d] generatesCSRsBadRequest  %+v", 400, o.Payload)
}

func (o *GeneratesCSRsBadRequest) String() string {
	return fmt.Sprintf("[PUT /v1/domains/{domainName}/csrs][%d] generatesCSRsBadRequest  %+v", 400, o.Payload)
}

func (o *GeneratesCSRsBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *GeneratesCSRsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGeneratesCSRsNotFound creates a GeneratesCSRsNotFound with default headers values
func NewGeneratesCSRsNotFound() *GeneratesCSRsNotFound {
	return &GeneratesCSRsNotFound{}
}

/*
GeneratesCSRsNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GeneratesCSRsNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this generates c s rs not found response has a 2xx status code
func (o *GeneratesCSRsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this generates c s rs not found response has a 3xx status code
func (o *GeneratesCSRsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this generates c s rs not found response has a 4xx status code
func (o *GeneratesCSRsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this generates c s rs not found response has a 5xx status code
func (o *GeneratesCSRsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this generates c s rs not found response a status code equal to that given
func (o *GeneratesCSRsNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GeneratesCSRsNotFound) Error() string {
	return fmt.Sprintf("[PUT /v1/domains/{domainName}/csrs][%d] generatesCSRsNotFound  %+v", 404, o.Payload)
}

func (o *GeneratesCSRsNotFound) String() string {
	return fmt.Sprintf("[PUT /v1/domains/{domainName}/csrs][%d] generatesCSRsNotFound  %+v", 404, o.Payload)
}

func (o *GeneratesCSRsNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GeneratesCSRsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGeneratesCSRsConflict creates a GeneratesCSRsConflict with default headers values
func NewGeneratesCSRsConflict() *GeneratesCSRsConflict {
	return &GeneratesCSRsConflict{}
}

/*
GeneratesCSRsConflict describes a response with status code 409, with default header values.

Conflict
*/
type GeneratesCSRsConflict struct {
	Payload *models.Error
}

// IsSuccess returns true when this generates c s rs conflict response has a 2xx status code
func (o *GeneratesCSRsConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this generates c s rs conflict response has a 3xx status code
func (o *GeneratesCSRsConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this generates c s rs conflict response has a 4xx status code
func (o *GeneratesCSRsConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this generates c s rs conflict response has a 5xx status code
func (o *GeneratesCSRsConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this generates c s rs conflict response a status code equal to that given
func (o *GeneratesCSRsConflict) IsCode(code int) bool {
	return code == 409
}

func (o *GeneratesCSRsConflict) Error() string {
	return fmt.Sprintf("[PUT /v1/domains/{domainName}/csrs][%d] generatesCSRsConflict  %+v", 409, o.Payload)
}

func (o *GeneratesCSRsConflict) String() string {
	return fmt.Sprintf("[PUT /v1/domains/{domainName}/csrs][%d] generatesCSRsConflict  %+v", 409, o.Payload)
}

func (o *GeneratesCSRsConflict) GetPayload() *models.Error {
	return o.Payload
}

func (o *GeneratesCSRsConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGeneratesCSRsInternalServerError creates a GeneratesCSRsInternalServerError with default headers values
func NewGeneratesCSRsInternalServerError() *GeneratesCSRsInternalServerError {
	return &GeneratesCSRsInternalServerError{}
}

/*
GeneratesCSRsInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type GeneratesCSRsInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this generates c s rs internal server error response has a 2xx status code
func (o *GeneratesCSRsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this generates c s rs internal server error response has a 3xx status code
func (o *GeneratesCSRsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this generates c s rs internal server error response has a 4xx status code
func (o *GeneratesCSRsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this generates c s rs internal server error response has a 5xx status code
func (o *GeneratesCSRsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this generates c s rs internal server error response a status code equal to that given
func (o *GeneratesCSRsInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GeneratesCSRsInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /v1/domains/{domainName}/csrs][%d] generatesCSRsInternalServerError  %+v", 500, o.Payload)
}

func (o *GeneratesCSRsInternalServerError) String() string {
	return fmt.Sprintf("[PUT /v1/domains/{domainName}/csrs][%d] generatesCSRsInternalServerError  %+v", 500, o.Payload)
}

func (o *GeneratesCSRsInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *GeneratesCSRsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
