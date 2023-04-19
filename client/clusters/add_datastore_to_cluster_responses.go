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

// AddDatastoreToClusterReader is a Reader for the AddDatastoreToCluster structure.
type AddDatastoreToClusterReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddDatastoreToClusterReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAddDatastoreToClusterOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewAddDatastoreToClusterAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewAddDatastoreToClusterBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewAddDatastoreToClusterNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewAddDatastoreToClusterInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewAddDatastoreToClusterOK creates a AddDatastoreToClusterOK with default headers values
func NewAddDatastoreToClusterOK() *AddDatastoreToClusterOK {
	return &AddDatastoreToClusterOK{}
}

/*
AddDatastoreToClusterOK describes a response with status code 200, with default header values.

OK
*/
type AddDatastoreToClusterOK struct {
	Payload *models.Task
}

// IsSuccess returns true when this add datastore to cluster o k response has a 2xx status code
func (o *AddDatastoreToClusterOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this add datastore to cluster o k response has a 3xx status code
func (o *AddDatastoreToClusterOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this add datastore to cluster o k response has a 4xx status code
func (o *AddDatastoreToClusterOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this add datastore to cluster o k response has a 5xx status code
func (o *AddDatastoreToClusterOK) IsServerError() bool {
	return false
}

// IsCode returns true when this add datastore to cluster o k response a status code equal to that given
func (o *AddDatastoreToClusterOK) IsCode(code int) bool {
	return code == 200
}

func (o *AddDatastoreToClusterOK) Error() string {
	return fmt.Sprintf("[POST /v1/clusters/{id}/datastores][%d] addDatastoreToClusterOK  %+v", 200, o.Payload)
}

func (o *AddDatastoreToClusterOK) String() string {
	return fmt.Sprintf("[POST /v1/clusters/{id}/datastores][%d] addDatastoreToClusterOK  %+v", 200, o.Payload)
}

func (o *AddDatastoreToClusterOK) GetPayload() *models.Task {
	return o.Payload
}

func (o *AddDatastoreToClusterOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Task)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddDatastoreToClusterAccepted creates a AddDatastoreToClusterAccepted with default headers values
func NewAddDatastoreToClusterAccepted() *AddDatastoreToClusterAccepted {
	return &AddDatastoreToClusterAccepted{}
}

/*
AddDatastoreToClusterAccepted describes a response with status code 202, with default header values.

Accepted
*/
type AddDatastoreToClusterAccepted struct {
	Payload *models.Task
}

// IsSuccess returns true when this add datastore to cluster accepted response has a 2xx status code
func (o *AddDatastoreToClusterAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this add datastore to cluster accepted response has a 3xx status code
func (o *AddDatastoreToClusterAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this add datastore to cluster accepted response has a 4xx status code
func (o *AddDatastoreToClusterAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this add datastore to cluster accepted response has a 5xx status code
func (o *AddDatastoreToClusterAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this add datastore to cluster accepted response a status code equal to that given
func (o *AddDatastoreToClusterAccepted) IsCode(code int) bool {
	return code == 202
}

func (o *AddDatastoreToClusterAccepted) Error() string {
	return fmt.Sprintf("[POST /v1/clusters/{id}/datastores][%d] addDatastoreToClusterAccepted  %+v", 202, o.Payload)
}

func (o *AddDatastoreToClusterAccepted) String() string {
	return fmt.Sprintf("[POST /v1/clusters/{id}/datastores][%d] addDatastoreToClusterAccepted  %+v", 202, o.Payload)
}

func (o *AddDatastoreToClusterAccepted) GetPayload() *models.Task {
	return o.Payload
}

func (o *AddDatastoreToClusterAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Task)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddDatastoreToClusterBadRequest creates a AddDatastoreToClusterBadRequest with default headers values
func NewAddDatastoreToClusterBadRequest() *AddDatastoreToClusterBadRequest {
	return &AddDatastoreToClusterBadRequest{}
}

/*
AddDatastoreToClusterBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type AddDatastoreToClusterBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this add datastore to cluster bad request response has a 2xx status code
func (o *AddDatastoreToClusterBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this add datastore to cluster bad request response has a 3xx status code
func (o *AddDatastoreToClusterBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this add datastore to cluster bad request response has a 4xx status code
func (o *AddDatastoreToClusterBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this add datastore to cluster bad request response has a 5xx status code
func (o *AddDatastoreToClusterBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this add datastore to cluster bad request response a status code equal to that given
func (o *AddDatastoreToClusterBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *AddDatastoreToClusterBadRequest) Error() string {
	return fmt.Sprintf("[POST /v1/clusters/{id}/datastores][%d] addDatastoreToClusterBadRequest  %+v", 400, o.Payload)
}

func (o *AddDatastoreToClusterBadRequest) String() string {
	return fmt.Sprintf("[POST /v1/clusters/{id}/datastores][%d] addDatastoreToClusterBadRequest  %+v", 400, o.Payload)
}

func (o *AddDatastoreToClusterBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *AddDatastoreToClusterBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddDatastoreToClusterNotFound creates a AddDatastoreToClusterNotFound with default headers values
func NewAddDatastoreToClusterNotFound() *AddDatastoreToClusterNotFound {
	return &AddDatastoreToClusterNotFound{}
}

/*
AddDatastoreToClusterNotFound describes a response with status code 404, with default header values.

Not Found
*/
type AddDatastoreToClusterNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this add datastore to cluster not found response has a 2xx status code
func (o *AddDatastoreToClusterNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this add datastore to cluster not found response has a 3xx status code
func (o *AddDatastoreToClusterNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this add datastore to cluster not found response has a 4xx status code
func (o *AddDatastoreToClusterNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this add datastore to cluster not found response has a 5xx status code
func (o *AddDatastoreToClusterNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this add datastore to cluster not found response a status code equal to that given
func (o *AddDatastoreToClusterNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *AddDatastoreToClusterNotFound) Error() string {
	return fmt.Sprintf("[POST /v1/clusters/{id}/datastores][%d] addDatastoreToClusterNotFound  %+v", 404, o.Payload)
}

func (o *AddDatastoreToClusterNotFound) String() string {
	return fmt.Sprintf("[POST /v1/clusters/{id}/datastores][%d] addDatastoreToClusterNotFound  %+v", 404, o.Payload)
}

func (o *AddDatastoreToClusterNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *AddDatastoreToClusterNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddDatastoreToClusterInternalServerError creates a AddDatastoreToClusterInternalServerError with default headers values
func NewAddDatastoreToClusterInternalServerError() *AddDatastoreToClusterInternalServerError {
	return &AddDatastoreToClusterInternalServerError{}
}

/*
AddDatastoreToClusterInternalServerError describes a response with status code 500, with default header values.

InternalServerError
*/
type AddDatastoreToClusterInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this add datastore to cluster internal server error response has a 2xx status code
func (o *AddDatastoreToClusterInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this add datastore to cluster internal server error response has a 3xx status code
func (o *AddDatastoreToClusterInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this add datastore to cluster internal server error response has a 4xx status code
func (o *AddDatastoreToClusterInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this add datastore to cluster internal server error response has a 5xx status code
func (o *AddDatastoreToClusterInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this add datastore to cluster internal server error response a status code equal to that given
func (o *AddDatastoreToClusterInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *AddDatastoreToClusterInternalServerError) Error() string {
	return fmt.Sprintf("[POST /v1/clusters/{id}/datastores][%d] addDatastoreToClusterInternalServerError  %+v", 500, o.Payload)
}

func (o *AddDatastoreToClusterInternalServerError) String() string {
	return fmt.Sprintf("[POST /v1/clusters/{id}/datastores][%d] addDatastoreToClusterInternalServerError  %+v", 500, o.Payload)
}

func (o *AddDatastoreToClusterInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *AddDatastoreToClusterInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
