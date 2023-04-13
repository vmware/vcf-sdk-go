// Code generated by go-swagger; DO NOT EDIT.

// Copyright © 2023 VMware, Inc. All Rights Reserved.
// Code licenced under: BSD-2

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"vcf-sdk-go/models"
)

// GETSSODomainsReader is a Reader for the GETSSODomains structure.
type GETSSODomainsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GETSSODomainsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGETSSODomainsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGETSSODomainsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGETSSODomainsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGETSSODomainsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGETSSODomainsOK creates a GETSSODomainsOK with default headers values
func NewGETSSODomainsOK() *GETSSODomainsOK {
	return &GETSSODomainsOK{}
}

/*
GETSSODomainsOK describes a response with status code 200, with default header values.

OK
*/
type GETSSODomainsOK struct {
	Payload *models.PageOfstring
}

// IsSuccess returns true when this get s s o domains o k response has a 2xx status code
func (o *GETSSODomainsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get s s o domains o k response has a 3xx status code
func (o *GETSSODomainsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get s s o domains o k response has a 4xx status code
func (o *GETSSODomainsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get s s o domains o k response has a 5xx status code
func (o *GETSSODomainsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get s s o domains o k response a status code equal to that given
func (o *GETSSODomainsOK) IsCode(code int) bool {
	return code == 200
}

func (o *GETSSODomainsOK) Error() string {
	return fmt.Sprintf("[GET /v1/sso-domains][%d] getSSODomainsOK  %+v", 200, o.Payload)
}

func (o *GETSSODomainsOK) String() string {
	return fmt.Sprintf("[GET /v1/sso-domains][%d] getSSODomainsOK  %+v", 200, o.Payload)
}

func (o *GETSSODomainsOK) GetPayload() *models.PageOfstring {
	return o.Payload
}

func (o *GETSSODomainsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PageOfstring)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGETSSODomainsUnauthorized creates a GETSSODomainsUnauthorized with default headers values
func NewGETSSODomainsUnauthorized() *GETSSODomainsUnauthorized {
	return &GETSSODomainsUnauthorized{}
}

/*
GETSSODomainsUnauthorized describes a response with status code 401, with default header values.

Unauthorized Request
*/
type GETSSODomainsUnauthorized struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this get s s o domains unauthorized response has a 2xx status code
func (o *GETSSODomainsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get s s o domains unauthorized response has a 3xx status code
func (o *GETSSODomainsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get s s o domains unauthorized response has a 4xx status code
func (o *GETSSODomainsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get s s o domains unauthorized response has a 5xx status code
func (o *GETSSODomainsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get s s o domains unauthorized response a status code equal to that given
func (o *GETSSODomainsUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GETSSODomainsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v1/sso-domains][%d] getSSODomainsUnauthorized  %+v", 401, o.Payload)
}

func (o *GETSSODomainsUnauthorized) String() string {
	return fmt.Sprintf("[GET /v1/sso-domains][%d] getSSODomainsUnauthorized  %+v", 401, o.Payload)
}

func (o *GETSSODomainsUnauthorized) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GETSSODomainsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGETSSODomainsForbidden creates a GETSSODomainsForbidden with default headers values
func NewGETSSODomainsForbidden() *GETSSODomainsForbidden {
	return &GETSSODomainsForbidden{}
}

/*
GETSSODomainsForbidden describes a response with status code 403, with default header values.

Forbidden request
*/
type GETSSODomainsForbidden struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this get s s o domains forbidden response has a 2xx status code
func (o *GETSSODomainsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get s s o domains forbidden response has a 3xx status code
func (o *GETSSODomainsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get s s o domains forbidden response has a 4xx status code
func (o *GETSSODomainsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get s s o domains forbidden response has a 5xx status code
func (o *GETSSODomainsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get s s o domains forbidden response a status code equal to that given
func (o *GETSSODomainsForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GETSSODomainsForbidden) Error() string {
	return fmt.Sprintf("[GET /v1/sso-domains][%d] getSSODomainsForbidden  %+v", 403, o.Payload)
}

func (o *GETSSODomainsForbidden) String() string {
	return fmt.Sprintf("[GET /v1/sso-domains][%d] getSSODomainsForbidden  %+v", 403, o.Payload)
}

func (o *GETSSODomainsForbidden) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GETSSODomainsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGETSSODomainsInternalServerError creates a GETSSODomainsInternalServerError with default headers values
func NewGETSSODomainsInternalServerError() *GETSSODomainsInternalServerError {
	return &GETSSODomainsInternalServerError{}
}

/*
GETSSODomainsInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GETSSODomainsInternalServerError struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this get s s o domains internal server error response has a 2xx status code
func (o *GETSSODomainsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get s s o domains internal server error response has a 3xx status code
func (o *GETSSODomainsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get s s o domains internal server error response has a 4xx status code
func (o *GETSSODomainsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get s s o domains internal server error response has a 5xx status code
func (o *GETSSODomainsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get s s o domains internal server error response a status code equal to that given
func (o *GETSSODomainsInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GETSSODomainsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/sso-domains][%d] getSSODomainsInternalServerError  %+v", 500, o.Payload)
}

func (o *GETSSODomainsInternalServerError) String() string {
	return fmt.Sprintf("[GET /v1/sso-domains][%d] getSSODomainsInternalServerError  %+v", 500, o.Payload)
}

func (o *GETSSODomainsInternalServerError) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GETSSODomainsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
