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

	"vcf-sdk-go/models"
)

// GETCertificateAuthorityByIDReader is a Reader for the GETCertificateAuthorityByID structure.
type GETCertificateAuthorityByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GETCertificateAuthorityByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGETCertificateAuthorityByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGETCertificateAuthorityByIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGETCertificateAuthorityByIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGETCertificateAuthorityByIDOK creates a GETCertificateAuthorityByIDOK with default headers values
func NewGETCertificateAuthorityByIDOK() *GETCertificateAuthorityByIDOK {
	return &GETCertificateAuthorityByIDOK{}
}

/*
GETCertificateAuthorityByIDOK describes a response with status code 200, with default header values.

OK
*/
type GETCertificateAuthorityByIDOK struct {
	Payload *models.CertificateAuthority
}

// IsSuccess returns true when this get certificate authority by Id o k response has a 2xx status code
func (o *GETCertificateAuthorityByIDOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get certificate authority by Id o k response has a 3xx status code
func (o *GETCertificateAuthorityByIDOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get certificate authority by Id o k response has a 4xx status code
func (o *GETCertificateAuthorityByIDOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get certificate authority by Id o k response has a 5xx status code
func (o *GETCertificateAuthorityByIDOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get certificate authority by Id o k response a status code equal to that given
func (o *GETCertificateAuthorityByIDOK) IsCode(code int) bool {
	return code == 200
}

func (o *GETCertificateAuthorityByIDOK) Error() string {
	return fmt.Sprintf("[GET /v1/certificate-authorities/{id}][%d] getCertificateAuthorityByIdOK  %+v", 200, o.Payload)
}

func (o *GETCertificateAuthorityByIDOK) String() string {
	return fmt.Sprintf("[GET /v1/certificate-authorities/{id}][%d] getCertificateAuthorityByIdOK  %+v", 200, o.Payload)
}

func (o *GETCertificateAuthorityByIDOK) GetPayload() *models.CertificateAuthority {
	return o.Payload
}

func (o *GETCertificateAuthorityByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CertificateAuthority)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGETCertificateAuthorityByIDNotFound creates a GETCertificateAuthorityByIDNotFound with default headers values
func NewGETCertificateAuthorityByIDNotFound() *GETCertificateAuthorityByIDNotFound {
	return &GETCertificateAuthorityByIDNotFound{}
}

/*
GETCertificateAuthorityByIDNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GETCertificateAuthorityByIDNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this get certificate authority by Id not found response has a 2xx status code
func (o *GETCertificateAuthorityByIDNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get certificate authority by Id not found response has a 3xx status code
func (o *GETCertificateAuthorityByIDNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get certificate authority by Id not found response has a 4xx status code
func (o *GETCertificateAuthorityByIDNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get certificate authority by Id not found response has a 5xx status code
func (o *GETCertificateAuthorityByIDNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get certificate authority by Id not found response a status code equal to that given
func (o *GETCertificateAuthorityByIDNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GETCertificateAuthorityByIDNotFound) Error() string {
	return fmt.Sprintf("[GET /v1/certificate-authorities/{id}][%d] getCertificateAuthorityByIdNotFound  %+v", 404, o.Payload)
}

func (o *GETCertificateAuthorityByIDNotFound) String() string {
	return fmt.Sprintf("[GET /v1/certificate-authorities/{id}][%d] getCertificateAuthorityByIdNotFound  %+v", 404, o.Payload)
}

func (o *GETCertificateAuthorityByIDNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GETCertificateAuthorityByIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGETCertificateAuthorityByIDInternalServerError creates a GETCertificateAuthorityByIDInternalServerError with default headers values
func NewGETCertificateAuthorityByIDInternalServerError() *GETCertificateAuthorityByIDInternalServerError {
	return &GETCertificateAuthorityByIDInternalServerError{}
}

/*
GETCertificateAuthorityByIDInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GETCertificateAuthorityByIDInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this get certificate authority by Id internal server error response has a 2xx status code
func (o *GETCertificateAuthorityByIDInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get certificate authority by Id internal server error response has a 3xx status code
func (o *GETCertificateAuthorityByIDInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get certificate authority by Id internal server error response has a 4xx status code
func (o *GETCertificateAuthorityByIDInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get certificate authority by Id internal server error response has a 5xx status code
func (o *GETCertificateAuthorityByIDInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get certificate authority by Id internal server error response a status code equal to that given
func (o *GETCertificateAuthorityByIDInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GETCertificateAuthorityByIDInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/certificate-authorities/{id}][%d] getCertificateAuthorityByIdInternalServerError  %+v", 500, o.Payload)
}

func (o *GETCertificateAuthorityByIDInternalServerError) String() string {
	return fmt.Sprintf("[GET /v1/certificate-authorities/{id}][%d] getCertificateAuthorityByIdInternalServerError  %+v", 500, o.Payload)
}

func (o *GETCertificateAuthorityByIDInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *GETCertificateAuthorityByIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
