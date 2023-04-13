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

	"vcf-sdk-go/models"
)

// ValidateVSANRemoteDatastoreReader is a Reader for the ValidateVSANRemoteDatastore structure.
type ValidateVSANRemoteDatastoreReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ValidateVSANRemoteDatastoreReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewValidateVSANRemoteDatastoreOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewValidateVSANRemoteDatastoreBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewValidateVSANRemoteDatastoreInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewValidateVSANRemoteDatastoreOK creates a ValidateVSANRemoteDatastoreOK with default headers values
func NewValidateVSANRemoteDatastoreOK() *ValidateVSANRemoteDatastoreOK {
	return &ValidateVSANRemoteDatastoreOK{}
}

/*
ValidateVSANRemoteDatastoreOK describes a response with status code 200, with default header values.

Ok
*/
type ValidateVSANRemoteDatastoreOK struct {
	Payload *models.Validation
}

// IsSuccess returns true when this validate Vsan remote datastore o k response has a 2xx status code
func (o *ValidateVSANRemoteDatastoreOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this validate Vsan remote datastore o k response has a 3xx status code
func (o *ValidateVSANRemoteDatastoreOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this validate Vsan remote datastore o k response has a 4xx status code
func (o *ValidateVSANRemoteDatastoreOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this validate Vsan remote datastore o k response has a 5xx status code
func (o *ValidateVSANRemoteDatastoreOK) IsServerError() bool {
	return false
}

// IsCode returns true when this validate Vsan remote datastore o k response a status code equal to that given
func (o *ValidateVSANRemoteDatastoreOK) IsCode(code int) bool {
	return code == 200
}

func (o *ValidateVSANRemoteDatastoreOK) Error() string {
	return fmt.Sprintf("[POST /v1/clusters/{clusterId}/datastores/validation][%d] validateVsanRemoteDatastoreOK  %+v", 200, o.Payload)
}

func (o *ValidateVSANRemoteDatastoreOK) String() string {
	return fmt.Sprintf("[POST /v1/clusters/{clusterId}/datastores/validation][%d] validateVsanRemoteDatastoreOK  %+v", 200, o.Payload)
}

func (o *ValidateVSANRemoteDatastoreOK) GetPayload() *models.Validation {
	return o.Payload
}

func (o *ValidateVSANRemoteDatastoreOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Validation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewValidateVSANRemoteDatastoreBadRequest creates a ValidateVSANRemoteDatastoreBadRequest with default headers values
func NewValidateVSANRemoteDatastoreBadRequest() *ValidateVSANRemoteDatastoreBadRequest {
	return &ValidateVSANRemoteDatastoreBadRequest{}
}

/*
ValidateVSANRemoteDatastoreBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type ValidateVSANRemoteDatastoreBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this validate Vsan remote datastore bad request response has a 2xx status code
func (o *ValidateVSANRemoteDatastoreBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this validate Vsan remote datastore bad request response has a 3xx status code
func (o *ValidateVSANRemoteDatastoreBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this validate Vsan remote datastore bad request response has a 4xx status code
func (o *ValidateVSANRemoteDatastoreBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this validate Vsan remote datastore bad request response has a 5xx status code
func (o *ValidateVSANRemoteDatastoreBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this validate Vsan remote datastore bad request response a status code equal to that given
func (o *ValidateVSANRemoteDatastoreBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *ValidateVSANRemoteDatastoreBadRequest) Error() string {
	return fmt.Sprintf("[POST /v1/clusters/{clusterId}/datastores/validation][%d] validateVsanRemoteDatastoreBadRequest  %+v", 400, o.Payload)
}

func (o *ValidateVSANRemoteDatastoreBadRequest) String() string {
	return fmt.Sprintf("[POST /v1/clusters/{clusterId}/datastores/validation][%d] validateVsanRemoteDatastoreBadRequest  %+v", 400, o.Payload)
}

func (o *ValidateVSANRemoteDatastoreBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *ValidateVSANRemoteDatastoreBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewValidateVSANRemoteDatastoreInternalServerError creates a ValidateVSANRemoteDatastoreInternalServerError with default headers values
func NewValidateVSANRemoteDatastoreInternalServerError() *ValidateVSANRemoteDatastoreInternalServerError {
	return &ValidateVSANRemoteDatastoreInternalServerError{}
}

/*
ValidateVSANRemoteDatastoreInternalServerError describes a response with status code 500, with default header values.

InternalServerError
*/
type ValidateVSANRemoteDatastoreInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this validate Vsan remote datastore internal server error response has a 2xx status code
func (o *ValidateVSANRemoteDatastoreInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this validate Vsan remote datastore internal server error response has a 3xx status code
func (o *ValidateVSANRemoteDatastoreInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this validate Vsan remote datastore internal server error response has a 4xx status code
func (o *ValidateVSANRemoteDatastoreInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this validate Vsan remote datastore internal server error response has a 5xx status code
func (o *ValidateVSANRemoteDatastoreInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this validate Vsan remote datastore internal server error response a status code equal to that given
func (o *ValidateVSANRemoteDatastoreInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *ValidateVSANRemoteDatastoreInternalServerError) Error() string {
	return fmt.Sprintf("[POST /v1/clusters/{clusterId}/datastores/validation][%d] validateVsanRemoteDatastoreInternalServerError  %+v", 500, o.Payload)
}

func (o *ValidateVSANRemoteDatastoreInternalServerError) String() string {
	return fmt.Sprintf("[POST /v1/clusters/{clusterId}/datastores/validation][%d] validateVsanRemoteDatastoreInternalServerError  %+v", 500, o.Payload)
}

func (o *ValidateVSANRemoteDatastoreInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *ValidateVSANRemoteDatastoreInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
