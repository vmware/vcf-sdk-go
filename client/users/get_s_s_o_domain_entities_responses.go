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

// GETSSODomainEntitiesReader is a Reader for the GETSSODomainEntities structure.
type GETSSODomainEntitiesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GETSSODomainEntitiesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGETSSODomainEntitiesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGETSSODomainEntitiesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGETSSODomainEntitiesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGETSSODomainEntitiesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGETSSODomainEntitiesOK creates a GETSSODomainEntitiesOK with default headers values
func NewGETSSODomainEntitiesOK() *GETSSODomainEntitiesOK {
	return &GETSSODomainEntitiesOK{}
}

/*
GETSSODomainEntitiesOK describes a response with status code 200, with default header values.

OK
*/
type GETSSODomainEntitiesOK struct {
	Payload *models.PageOfSsoDomainEntity
}

// IsSuccess returns true when this get s s o domain entities o k response has a 2xx status code
func (o *GETSSODomainEntitiesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get s s o domain entities o k response has a 3xx status code
func (o *GETSSODomainEntitiesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get s s o domain entities o k response has a 4xx status code
func (o *GETSSODomainEntitiesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get s s o domain entities o k response has a 5xx status code
func (o *GETSSODomainEntitiesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get s s o domain entities o k response a status code equal to that given
func (o *GETSSODomainEntitiesOK) IsCode(code int) bool {
	return code == 200
}

func (o *GETSSODomainEntitiesOK) Error() string {
	return fmt.Sprintf("[GET /v1/sso-domains/{sso-domain}/entities][%d] getSSODomainEntitiesOK  %+v", 200, o.Payload)
}

func (o *GETSSODomainEntitiesOK) String() string {
	return fmt.Sprintf("[GET /v1/sso-domains/{sso-domain}/entities][%d] getSSODomainEntitiesOK  %+v", 200, o.Payload)
}

func (o *GETSSODomainEntitiesOK) GetPayload() *models.PageOfSsoDomainEntity {
	return o.Payload
}

func (o *GETSSODomainEntitiesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PageOfSsoDomainEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGETSSODomainEntitiesUnauthorized creates a GETSSODomainEntitiesUnauthorized with default headers values
func NewGETSSODomainEntitiesUnauthorized() *GETSSODomainEntitiesUnauthorized {
	return &GETSSODomainEntitiesUnauthorized{}
}

/*
GETSSODomainEntitiesUnauthorized describes a response with status code 401, with default header values.

Unauthorized Request
*/
type GETSSODomainEntitiesUnauthorized struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this get s s o domain entities unauthorized response has a 2xx status code
func (o *GETSSODomainEntitiesUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get s s o domain entities unauthorized response has a 3xx status code
func (o *GETSSODomainEntitiesUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get s s o domain entities unauthorized response has a 4xx status code
func (o *GETSSODomainEntitiesUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get s s o domain entities unauthorized response has a 5xx status code
func (o *GETSSODomainEntitiesUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get s s o domain entities unauthorized response a status code equal to that given
func (o *GETSSODomainEntitiesUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GETSSODomainEntitiesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v1/sso-domains/{sso-domain}/entities][%d] getSSODomainEntitiesUnauthorized  %+v", 401, o.Payload)
}

func (o *GETSSODomainEntitiesUnauthorized) String() string {
	return fmt.Sprintf("[GET /v1/sso-domains/{sso-domain}/entities][%d] getSSODomainEntitiesUnauthorized  %+v", 401, o.Payload)
}

func (o *GETSSODomainEntitiesUnauthorized) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GETSSODomainEntitiesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGETSSODomainEntitiesForbidden creates a GETSSODomainEntitiesForbidden with default headers values
func NewGETSSODomainEntitiesForbidden() *GETSSODomainEntitiesForbidden {
	return &GETSSODomainEntitiesForbidden{}
}

/*
GETSSODomainEntitiesForbidden describes a response with status code 403, with default header values.

Forbidden request
*/
type GETSSODomainEntitiesForbidden struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this get s s o domain entities forbidden response has a 2xx status code
func (o *GETSSODomainEntitiesForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get s s o domain entities forbidden response has a 3xx status code
func (o *GETSSODomainEntitiesForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get s s o domain entities forbidden response has a 4xx status code
func (o *GETSSODomainEntitiesForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get s s o domain entities forbidden response has a 5xx status code
func (o *GETSSODomainEntitiesForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get s s o domain entities forbidden response a status code equal to that given
func (o *GETSSODomainEntitiesForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GETSSODomainEntitiesForbidden) Error() string {
	return fmt.Sprintf("[GET /v1/sso-domains/{sso-domain}/entities][%d] getSSODomainEntitiesForbidden  %+v", 403, o.Payload)
}

func (o *GETSSODomainEntitiesForbidden) String() string {
	return fmt.Sprintf("[GET /v1/sso-domains/{sso-domain}/entities][%d] getSSODomainEntitiesForbidden  %+v", 403, o.Payload)
}

func (o *GETSSODomainEntitiesForbidden) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GETSSODomainEntitiesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGETSSODomainEntitiesInternalServerError creates a GETSSODomainEntitiesInternalServerError with default headers values
func NewGETSSODomainEntitiesInternalServerError() *GETSSODomainEntitiesInternalServerError {
	return &GETSSODomainEntitiesInternalServerError{}
}

/*
GETSSODomainEntitiesInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GETSSODomainEntitiesInternalServerError struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this get s s o domain entities internal server error response has a 2xx status code
func (o *GETSSODomainEntitiesInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get s s o domain entities internal server error response has a 3xx status code
func (o *GETSSODomainEntitiesInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get s s o domain entities internal server error response has a 4xx status code
func (o *GETSSODomainEntitiesInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get s s o domain entities internal server error response has a 5xx status code
func (o *GETSSODomainEntitiesInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get s s o domain entities internal server error response a status code equal to that given
func (o *GETSSODomainEntitiesInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GETSSODomainEntitiesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/sso-domains/{sso-domain}/entities][%d] getSSODomainEntitiesInternalServerError  %+v", 500, o.Payload)
}

func (o *GETSSODomainEntitiesInternalServerError) String() string {
	return fmt.Sprintf("[GET /v1/sso-domains/{sso-domain}/entities][%d] getSSODomainEntitiesInternalServerError  %+v", 500, o.Payload)
}

func (o *GETSSODomainEntitiesInternalServerError) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GETSSODomainEntitiesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
