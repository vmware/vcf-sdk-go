// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2023 VMware, Inc.
// SPDX-License-Identifier: BSD-2-Clause

package nsxt_clusters

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/vmware/vcf-sdk-go/models"
)

// StartNsxCriteriaQueryReader is a Reader for the StartNsxCriteriaQuery structure.
type StartNsxCriteriaQueryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StartNsxCriteriaQueryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStartNsxCriteriaQueryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewStartNsxCriteriaQueryAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewStartNsxCriteriaQueryBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewStartNsxCriteriaQueryInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /v1/nsxt-clusters/queries] startNsxCriteriaQuery", response, response.Code())
	}
}

// NewStartNsxCriteriaQueryOK creates a StartNsxCriteriaQueryOK with default headers values
func NewStartNsxCriteriaQueryOK() *StartNsxCriteriaQueryOK {
	return &StartNsxCriteriaQueryOK{}
}

/*
StartNsxCriteriaQueryOK describes a response with status code 200, with default header values.

OK
*/
type StartNsxCriteriaQueryOK struct {
	Payload *models.NsxTQueryResponse
}

// IsSuccess returns true when this start nsx criteria query o k response has a 2xx status code
func (o *StartNsxCriteriaQueryOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this start nsx criteria query o k response has a 3xx status code
func (o *StartNsxCriteriaQueryOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this start nsx criteria query o k response has a 4xx status code
func (o *StartNsxCriteriaQueryOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this start nsx criteria query o k response has a 5xx status code
func (o *StartNsxCriteriaQueryOK) IsServerError() bool {
	return false
}

// IsCode returns true when this start nsx criteria query o k response a status code equal to that given
func (o *StartNsxCriteriaQueryOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the start nsx criteria query o k response
func (o *StartNsxCriteriaQueryOK) Code() int {
	return 200
}

func (o *StartNsxCriteriaQueryOK) Error() string {
	return fmt.Sprintf("[POST /v1/nsxt-clusters/queries][%d] startNsxCriteriaQueryOK  %+v", 200, o.Payload)
}

func (o *StartNsxCriteriaQueryOK) String() string {
	return fmt.Sprintf("[POST /v1/nsxt-clusters/queries][%d] startNsxCriteriaQueryOK  %+v", 200, o.Payload)
}

func (o *StartNsxCriteriaQueryOK) GetPayload() *models.NsxTQueryResponse {
	return o.Payload
}

func (o *StartNsxCriteriaQueryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NsxTQueryResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStartNsxCriteriaQueryAccepted creates a StartNsxCriteriaQueryAccepted with default headers values
func NewStartNsxCriteriaQueryAccepted() *StartNsxCriteriaQueryAccepted {
	return &StartNsxCriteriaQueryAccepted{}
}

/*
StartNsxCriteriaQueryAccepted describes a response with status code 202, with default header values.

Accepted
*/
type StartNsxCriteriaQueryAccepted struct {
	Payload *models.NsxTQueryResponse
}

// IsSuccess returns true when this start nsx criteria query accepted response has a 2xx status code
func (o *StartNsxCriteriaQueryAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this start nsx criteria query accepted response has a 3xx status code
func (o *StartNsxCriteriaQueryAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this start nsx criteria query accepted response has a 4xx status code
func (o *StartNsxCriteriaQueryAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this start nsx criteria query accepted response has a 5xx status code
func (o *StartNsxCriteriaQueryAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this start nsx criteria query accepted response a status code equal to that given
func (o *StartNsxCriteriaQueryAccepted) IsCode(code int) bool {
	return code == 202
}

// Code gets the status code for the start nsx criteria query accepted response
func (o *StartNsxCriteriaQueryAccepted) Code() int {
	return 202
}

func (o *StartNsxCriteriaQueryAccepted) Error() string {
	return fmt.Sprintf("[POST /v1/nsxt-clusters/queries][%d] startNsxCriteriaQueryAccepted  %+v", 202, o.Payload)
}

func (o *StartNsxCriteriaQueryAccepted) String() string {
	return fmt.Sprintf("[POST /v1/nsxt-clusters/queries][%d] startNsxCriteriaQueryAccepted  %+v", 202, o.Payload)
}

func (o *StartNsxCriteriaQueryAccepted) GetPayload() *models.NsxTQueryResponse {
	return o.Payload
}

func (o *StartNsxCriteriaQueryAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NsxTQueryResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStartNsxCriteriaQueryBadRequest creates a StartNsxCriteriaQueryBadRequest with default headers values
func NewStartNsxCriteriaQueryBadRequest() *StartNsxCriteriaQueryBadRequest {
	return &StartNsxCriteriaQueryBadRequest{}
}

/*
StartNsxCriteriaQueryBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type StartNsxCriteriaQueryBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this start nsx criteria query bad request response has a 2xx status code
func (o *StartNsxCriteriaQueryBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this start nsx criteria query bad request response has a 3xx status code
func (o *StartNsxCriteriaQueryBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this start nsx criteria query bad request response has a 4xx status code
func (o *StartNsxCriteriaQueryBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this start nsx criteria query bad request response has a 5xx status code
func (o *StartNsxCriteriaQueryBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this start nsx criteria query bad request response a status code equal to that given
func (o *StartNsxCriteriaQueryBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the start nsx criteria query bad request response
func (o *StartNsxCriteriaQueryBadRequest) Code() int {
	return 400
}

func (o *StartNsxCriteriaQueryBadRequest) Error() string {
	return fmt.Sprintf("[POST /v1/nsxt-clusters/queries][%d] startNsxCriteriaQueryBadRequest  %+v", 400, o.Payload)
}

func (o *StartNsxCriteriaQueryBadRequest) String() string {
	return fmt.Sprintf("[POST /v1/nsxt-clusters/queries][%d] startNsxCriteriaQueryBadRequest  %+v", 400, o.Payload)
}

func (o *StartNsxCriteriaQueryBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *StartNsxCriteriaQueryBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStartNsxCriteriaQueryInternalServerError creates a StartNsxCriteriaQueryInternalServerError with default headers values
func NewStartNsxCriteriaQueryInternalServerError() *StartNsxCriteriaQueryInternalServerError {
	return &StartNsxCriteriaQueryInternalServerError{}
}

/*
StartNsxCriteriaQueryInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type StartNsxCriteriaQueryInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this start nsx criteria query internal server error response has a 2xx status code
func (o *StartNsxCriteriaQueryInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this start nsx criteria query internal server error response has a 3xx status code
func (o *StartNsxCriteriaQueryInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this start nsx criteria query internal server error response has a 4xx status code
func (o *StartNsxCriteriaQueryInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this start nsx criteria query internal server error response has a 5xx status code
func (o *StartNsxCriteriaQueryInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this start nsx criteria query internal server error response a status code equal to that given
func (o *StartNsxCriteriaQueryInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the start nsx criteria query internal server error response
func (o *StartNsxCriteriaQueryInternalServerError) Code() int {
	return 500
}

func (o *StartNsxCriteriaQueryInternalServerError) Error() string {
	return fmt.Sprintf("[POST /v1/nsxt-clusters/queries][%d] startNsxCriteriaQueryInternalServerError  %+v", 500, o.Payload)
}

func (o *StartNsxCriteriaQueryInternalServerError) String() string {
	return fmt.Sprintf("[POST /v1/nsxt-clusters/queries][%d] startNsxCriteriaQueryInternalServerError  %+v", 500, o.Payload)
}

func (o *StartNsxCriteriaQueryInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *StartNsxCriteriaQueryInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
