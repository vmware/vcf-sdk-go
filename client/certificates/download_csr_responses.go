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

// DownloadCSRReader is a Reader for the DownloadCSR structure.
type DownloadCSRReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DownloadCSRReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDownloadCSROK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewDownloadCSRNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDownloadCSRInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /v1/domains/{id}/csrs/downloads] downloadCSR", response, response.Code())
	}
}

// NewDownloadCSROK creates a DownloadCSROK with default headers values
func NewDownloadCSROK() *DownloadCSROK {
	return &DownloadCSROK{}
}

/*
DownloadCSROK describes a response with status code 200, with default header values.

OK
*/
type DownloadCSROK struct {
	Payload strfmt.Base64
}

// IsSuccess returns true when this download Csr o k response has a 2xx status code
func (o *DownloadCSROK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this download Csr o k response has a 3xx status code
func (o *DownloadCSROK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this download Csr o k response has a 4xx status code
func (o *DownloadCSROK) IsClientError() bool {
	return false
}

// IsServerError returns true when this download Csr o k response has a 5xx status code
func (o *DownloadCSROK) IsServerError() bool {
	return false
}

// IsCode returns true when this download Csr o k response a status code equal to that given
func (o *DownloadCSROK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the download Csr o k response
func (o *DownloadCSROK) Code() int {
	return 200
}

func (o *DownloadCSROK) Error() string {
	return fmt.Sprintf("[GET /v1/domains/{id}/csrs/downloads][%d] downloadCsrOK  %+v", 200, o.Payload)
}

func (o *DownloadCSROK) String() string {
	return fmt.Sprintf("[GET /v1/domains/{id}/csrs/downloads][%d] downloadCsrOK  %+v", 200, o.Payload)
}

func (o *DownloadCSROK) GetPayload() strfmt.Base64 {
	return o.Payload
}

func (o *DownloadCSROK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDownloadCSRNotFound creates a DownloadCSRNotFound with default headers values
func NewDownloadCSRNotFound() *DownloadCSRNotFound {
	return &DownloadCSRNotFound{}
}

/*
DownloadCSRNotFound describes a response with status code 404, with default header values.

Not Found
*/
type DownloadCSRNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this download Csr not found response has a 2xx status code
func (o *DownloadCSRNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this download Csr not found response has a 3xx status code
func (o *DownloadCSRNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this download Csr not found response has a 4xx status code
func (o *DownloadCSRNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this download Csr not found response has a 5xx status code
func (o *DownloadCSRNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this download Csr not found response a status code equal to that given
func (o *DownloadCSRNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the download Csr not found response
func (o *DownloadCSRNotFound) Code() int {
	return 404
}

func (o *DownloadCSRNotFound) Error() string {
	return fmt.Sprintf("[GET /v1/domains/{id}/csrs/downloads][%d] downloadCsrNotFound  %+v", 404, o.Payload)
}

func (o *DownloadCSRNotFound) String() string {
	return fmt.Sprintf("[GET /v1/domains/{id}/csrs/downloads][%d] downloadCsrNotFound  %+v", 404, o.Payload)
}

func (o *DownloadCSRNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *DownloadCSRNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDownloadCSRInternalServerError creates a DownloadCSRInternalServerError with default headers values
func NewDownloadCSRInternalServerError() *DownloadCSRInternalServerError {
	return &DownloadCSRInternalServerError{}
}

/*
DownloadCSRInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type DownloadCSRInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this download Csr internal server error response has a 2xx status code
func (o *DownloadCSRInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this download Csr internal server error response has a 3xx status code
func (o *DownloadCSRInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this download Csr internal server error response has a 4xx status code
func (o *DownloadCSRInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this download Csr internal server error response has a 5xx status code
func (o *DownloadCSRInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this download Csr internal server error response a status code equal to that given
func (o *DownloadCSRInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the download Csr internal server error response
func (o *DownloadCSRInternalServerError) Code() int {
	return 500
}

func (o *DownloadCSRInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/domains/{id}/csrs/downloads][%d] downloadCsrInternalServerError  %+v", 500, o.Payload)
}

func (o *DownloadCSRInternalServerError) String() string {
	return fmt.Sprintf("[GET /v1/domains/{id}/csrs/downloads][%d] downloadCsrInternalServerError  %+v", 500, o.Payload)
}

func (o *DownloadCSRInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *DownloadCSRInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
