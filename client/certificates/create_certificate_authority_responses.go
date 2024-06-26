// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2023 VMware, Inc.
// SPDX-License-Identifier: BSD-2-Clause

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

// CreateCertificateAuthorityReader is a Reader for the CreateCertificateAuthority structure.
type CreateCertificateAuthorityReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateCertificateAuthorityReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateCertificateAuthorityOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateCertificateAuthorityBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCreateCertificateAuthorityInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[PUT /v1/certificate-authorities] createCertificateAuthority", response, response.Code())
	}
}

// NewCreateCertificateAuthorityOK creates a CreateCertificateAuthorityOK with default headers values
func NewCreateCertificateAuthorityOK() *CreateCertificateAuthorityOK {
	return &CreateCertificateAuthorityOK{}
}

/*
CreateCertificateAuthorityOK describes a response with status code 200, with default header values.

OK
*/
type CreateCertificateAuthorityOK struct {
	Payload interface{}
}

// IsSuccess returns true when this create certificate authority o k response has a 2xx status code
func (o *CreateCertificateAuthorityOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create certificate authority o k response has a 3xx status code
func (o *CreateCertificateAuthorityOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create certificate authority o k response has a 4xx status code
func (o *CreateCertificateAuthorityOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this create certificate authority o k response has a 5xx status code
func (o *CreateCertificateAuthorityOK) IsServerError() bool {
	return false
}

// IsCode returns true when this create certificate authority o k response a status code equal to that given
func (o *CreateCertificateAuthorityOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the create certificate authority o k response
func (o *CreateCertificateAuthorityOK) Code() int {
	return 200
}

func (o *CreateCertificateAuthorityOK) Error() string {
	return fmt.Sprintf("[PUT /v1/certificate-authorities][%d] createCertificateAuthorityOK  %+v", 200, o.Payload)
}

func (o *CreateCertificateAuthorityOK) String() string {
	return fmt.Sprintf("[PUT /v1/certificate-authorities][%d] createCertificateAuthorityOK  %+v", 200, o.Payload)
}

func (o *CreateCertificateAuthorityOK) GetPayload() interface{} {
	return o.Payload
}

func (o *CreateCertificateAuthorityOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateCertificateAuthorityBadRequest creates a CreateCertificateAuthorityBadRequest with default headers values
func NewCreateCertificateAuthorityBadRequest() *CreateCertificateAuthorityBadRequest {
	return &CreateCertificateAuthorityBadRequest{}
}

/*
CreateCertificateAuthorityBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type CreateCertificateAuthorityBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this create certificate authority bad request response has a 2xx status code
func (o *CreateCertificateAuthorityBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create certificate authority bad request response has a 3xx status code
func (o *CreateCertificateAuthorityBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create certificate authority bad request response has a 4xx status code
func (o *CreateCertificateAuthorityBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this create certificate authority bad request response has a 5xx status code
func (o *CreateCertificateAuthorityBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this create certificate authority bad request response a status code equal to that given
func (o *CreateCertificateAuthorityBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the create certificate authority bad request response
func (o *CreateCertificateAuthorityBadRequest) Code() int {
	return 400
}

func (o *CreateCertificateAuthorityBadRequest) Error() string {
	return fmt.Sprintf("[PUT /v1/certificate-authorities][%d] createCertificateAuthorityBadRequest  %+v", 400, o.Payload)
}

func (o *CreateCertificateAuthorityBadRequest) String() string {
	return fmt.Sprintf("[PUT /v1/certificate-authorities][%d] createCertificateAuthorityBadRequest  %+v", 400, o.Payload)
}

func (o *CreateCertificateAuthorityBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateCertificateAuthorityBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateCertificateAuthorityInternalServerError creates a CreateCertificateAuthorityInternalServerError with default headers values
func NewCreateCertificateAuthorityInternalServerError() *CreateCertificateAuthorityInternalServerError {
	return &CreateCertificateAuthorityInternalServerError{}
}

/*
CreateCertificateAuthorityInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type CreateCertificateAuthorityInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this create certificate authority internal server error response has a 2xx status code
func (o *CreateCertificateAuthorityInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create certificate authority internal server error response has a 3xx status code
func (o *CreateCertificateAuthorityInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create certificate authority internal server error response has a 4xx status code
func (o *CreateCertificateAuthorityInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this create certificate authority internal server error response has a 5xx status code
func (o *CreateCertificateAuthorityInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this create certificate authority internal server error response a status code equal to that given
func (o *CreateCertificateAuthorityInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the create certificate authority internal server error response
func (o *CreateCertificateAuthorityInternalServerError) Code() int {
	return 500
}

func (o *CreateCertificateAuthorityInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /v1/certificate-authorities][%d] createCertificateAuthorityInternalServerError  %+v", 500, o.Payload)
}

func (o *CreateCertificateAuthorityInternalServerError) String() string {
	return fmt.Sprintf("[PUT /v1/certificate-authorities][%d] createCertificateAuthorityInternalServerError  %+v", 500, o.Payload)
}

func (o *CreateCertificateAuthorityInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateCertificateAuthorityInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
