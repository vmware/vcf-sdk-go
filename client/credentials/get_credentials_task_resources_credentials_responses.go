// Code generated by go-swagger; DO NOT EDIT.

// Copyright © 2023 VMware, Inc. All Rights Reserved.
// Code licenced under: BSD-2

package credentials

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"vcf-sdk-go/models"
)

// GETCredentialsTaskResourcesCredentialsReader is a Reader for the GETCredentialsTaskResourcesCredentials structure.
type GETCredentialsTaskResourcesCredentialsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GETCredentialsTaskResourcesCredentialsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGETCredentialsTaskResourcesCredentialsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGETCredentialsTaskResourcesCredentialsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGETCredentialsTaskResourcesCredentialsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGETCredentialsTaskResourcesCredentialsOK creates a GETCredentialsTaskResourcesCredentialsOK with default headers values
func NewGETCredentialsTaskResourcesCredentialsOK() *GETCredentialsTaskResourcesCredentialsOK {
	return &GETCredentialsTaskResourcesCredentialsOK{}
}

/*
GETCredentialsTaskResourcesCredentialsOK describes a response with status code 200, with default header values.

OK
*/
type GETCredentialsTaskResourcesCredentialsOK struct {
	Payload []*models.ResourceCredentials
}

// IsSuccess returns true when this get credentials task resources credentials o k response has a 2xx status code
func (o *GETCredentialsTaskResourcesCredentialsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get credentials task resources credentials o k response has a 3xx status code
func (o *GETCredentialsTaskResourcesCredentialsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get credentials task resources credentials o k response has a 4xx status code
func (o *GETCredentialsTaskResourcesCredentialsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get credentials task resources credentials o k response has a 5xx status code
func (o *GETCredentialsTaskResourcesCredentialsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get credentials task resources credentials o k response a status code equal to that given
func (o *GETCredentialsTaskResourcesCredentialsOK) IsCode(code int) bool {
	return code == 200
}

func (o *GETCredentialsTaskResourcesCredentialsOK) Error() string {
	return fmt.Sprintf("[GET /v1/credentials/tasks/{id}/resource-credentials][%d] getCredentialsTaskResourcesCredentialsOK  %+v", 200, o.Payload)
}

func (o *GETCredentialsTaskResourcesCredentialsOK) String() string {
	return fmt.Sprintf("[GET /v1/credentials/tasks/{id}/resource-credentials][%d] getCredentialsTaskResourcesCredentialsOK  %+v", 200, o.Payload)
}

func (o *GETCredentialsTaskResourcesCredentialsOK) GetPayload() []*models.ResourceCredentials {
	return o.Payload
}

func (o *GETCredentialsTaskResourcesCredentialsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGETCredentialsTaskResourcesCredentialsBadRequest creates a GETCredentialsTaskResourcesCredentialsBadRequest with default headers values
func NewGETCredentialsTaskResourcesCredentialsBadRequest() *GETCredentialsTaskResourcesCredentialsBadRequest {
	return &GETCredentialsTaskResourcesCredentialsBadRequest{}
}

/*
GETCredentialsTaskResourcesCredentialsBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GETCredentialsTaskResourcesCredentialsBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this get credentials task resources credentials bad request response has a 2xx status code
func (o *GETCredentialsTaskResourcesCredentialsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get credentials task resources credentials bad request response has a 3xx status code
func (o *GETCredentialsTaskResourcesCredentialsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get credentials task resources credentials bad request response has a 4xx status code
func (o *GETCredentialsTaskResourcesCredentialsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get credentials task resources credentials bad request response has a 5xx status code
func (o *GETCredentialsTaskResourcesCredentialsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get credentials task resources credentials bad request response a status code equal to that given
func (o *GETCredentialsTaskResourcesCredentialsBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GETCredentialsTaskResourcesCredentialsBadRequest) Error() string {
	return fmt.Sprintf("[GET /v1/credentials/tasks/{id}/resource-credentials][%d] getCredentialsTaskResourcesCredentialsBadRequest  %+v", 400, o.Payload)
}

func (o *GETCredentialsTaskResourcesCredentialsBadRequest) String() string {
	return fmt.Sprintf("[GET /v1/credentials/tasks/{id}/resource-credentials][%d] getCredentialsTaskResourcesCredentialsBadRequest  %+v", 400, o.Payload)
}

func (o *GETCredentialsTaskResourcesCredentialsBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *GETCredentialsTaskResourcesCredentialsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGETCredentialsTaskResourcesCredentialsInternalServerError creates a GETCredentialsTaskResourcesCredentialsInternalServerError with default headers values
func NewGETCredentialsTaskResourcesCredentialsInternalServerError() *GETCredentialsTaskResourcesCredentialsInternalServerError {
	return &GETCredentialsTaskResourcesCredentialsInternalServerError{}
}

/*
GETCredentialsTaskResourcesCredentialsInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GETCredentialsTaskResourcesCredentialsInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this get credentials task resources credentials internal server error response has a 2xx status code
func (o *GETCredentialsTaskResourcesCredentialsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get credentials task resources credentials internal server error response has a 3xx status code
func (o *GETCredentialsTaskResourcesCredentialsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get credentials task resources credentials internal server error response has a 4xx status code
func (o *GETCredentialsTaskResourcesCredentialsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get credentials task resources credentials internal server error response has a 5xx status code
func (o *GETCredentialsTaskResourcesCredentialsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get credentials task resources credentials internal server error response a status code equal to that given
func (o *GETCredentialsTaskResourcesCredentialsInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GETCredentialsTaskResourcesCredentialsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/credentials/tasks/{id}/resource-credentials][%d] getCredentialsTaskResourcesCredentialsInternalServerError  %+v", 500, o.Payload)
}

func (o *GETCredentialsTaskResourcesCredentialsInternalServerError) String() string {
	return fmt.Sprintf("[GET /v1/credentials/tasks/{id}/resource-credentials][%d] getCredentialsTaskResourcesCredentialsInternalServerError  %+v", 500, o.Payload)
}

func (o *GETCredentialsTaskResourcesCredentialsInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *GETCredentialsTaskResourcesCredentialsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
