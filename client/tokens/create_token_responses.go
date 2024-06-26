// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2023 VMware, Inc.
// SPDX-License-Identifier: BSD-2-Clause

package tokens

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/vmware/vcf-sdk-go/models"
)

// CreateTokenReader is a Reader for the CreateToken structure.
type CreateTokenReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateTokenReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateTokenOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 201:
		result := NewCreateTokenCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateTokenBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCreateTokenInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /v1/tokens] createToken", response, response.Code())
	}
}

// NewCreateTokenOK creates a CreateTokenOK with default headers values
func NewCreateTokenOK() *CreateTokenOK {
	return &CreateTokenOK{}
}

/*
CreateTokenOK describes a response with status code 200, with default header values.

OK
*/
type CreateTokenOK struct {
	Payload *models.TokenPair
}

// IsSuccess returns true when this create token o k response has a 2xx status code
func (o *CreateTokenOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create token o k response has a 3xx status code
func (o *CreateTokenOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create token o k response has a 4xx status code
func (o *CreateTokenOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this create token o k response has a 5xx status code
func (o *CreateTokenOK) IsServerError() bool {
	return false
}

// IsCode returns true when this create token o k response a status code equal to that given
func (o *CreateTokenOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the create token o k response
func (o *CreateTokenOK) Code() int {
	return 200
}

func (o *CreateTokenOK) Error() string {
	return fmt.Sprintf("[POST /v1/tokens][%d] createTokenOK  %+v", 200, o.Payload)
}

func (o *CreateTokenOK) String() string {
	return fmt.Sprintf("[POST /v1/tokens][%d] createTokenOK  %+v", 200, o.Payload)
}

func (o *CreateTokenOK) GetPayload() *models.TokenPair {
	return o.Payload
}

func (o *CreateTokenOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TokenPair)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateTokenCreated creates a CreateTokenCreated with default headers values
func NewCreateTokenCreated() *CreateTokenCreated {
	return &CreateTokenCreated{}
}

/*
CreateTokenCreated describes a response with status code 201, with default header values.

Created
*/
type CreateTokenCreated struct {
	Payload *models.TokenPair
}

// IsSuccess returns true when this create token created response has a 2xx status code
func (o *CreateTokenCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create token created response has a 3xx status code
func (o *CreateTokenCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create token created response has a 4xx status code
func (o *CreateTokenCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this create token created response has a 5xx status code
func (o *CreateTokenCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this create token created response a status code equal to that given
func (o *CreateTokenCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the create token created response
func (o *CreateTokenCreated) Code() int {
	return 201
}

func (o *CreateTokenCreated) Error() string {
	return fmt.Sprintf("[POST /v1/tokens][%d] createTokenCreated  %+v", 201, o.Payload)
}

func (o *CreateTokenCreated) String() string {
	return fmt.Sprintf("[POST /v1/tokens][%d] createTokenCreated  %+v", 201, o.Payload)
}

func (o *CreateTokenCreated) GetPayload() *models.TokenPair {
	return o.Payload
}

func (o *CreateTokenCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TokenPair)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateTokenBadRequest creates a CreateTokenBadRequest with default headers values
func NewCreateTokenBadRequest() *CreateTokenBadRequest {
	return &CreateTokenBadRequest{}
}

/*
CreateTokenBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type CreateTokenBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this create token bad request response has a 2xx status code
func (o *CreateTokenBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create token bad request response has a 3xx status code
func (o *CreateTokenBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create token bad request response has a 4xx status code
func (o *CreateTokenBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this create token bad request response has a 5xx status code
func (o *CreateTokenBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this create token bad request response a status code equal to that given
func (o *CreateTokenBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the create token bad request response
func (o *CreateTokenBadRequest) Code() int {
	return 400
}

func (o *CreateTokenBadRequest) Error() string {
	return fmt.Sprintf("[POST /v1/tokens][%d] createTokenBadRequest  %+v", 400, o.Payload)
}

func (o *CreateTokenBadRequest) String() string {
	return fmt.Sprintf("[POST /v1/tokens][%d] createTokenBadRequest  %+v", 400, o.Payload)
}

func (o *CreateTokenBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateTokenBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateTokenInternalServerError creates a CreateTokenInternalServerError with default headers values
func NewCreateTokenInternalServerError() *CreateTokenInternalServerError {
	return &CreateTokenInternalServerError{}
}

/*
CreateTokenInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type CreateTokenInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this create token internal server error response has a 2xx status code
func (o *CreateTokenInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create token internal server error response has a 3xx status code
func (o *CreateTokenInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create token internal server error response has a 4xx status code
func (o *CreateTokenInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this create token internal server error response has a 5xx status code
func (o *CreateTokenInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this create token internal server error response a status code equal to that given
func (o *CreateTokenInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the create token internal server error response
func (o *CreateTokenInternalServerError) Code() int {
	return 500
}

func (o *CreateTokenInternalServerError) Error() string {
	return fmt.Sprintf("[POST /v1/tokens][%d] createTokenInternalServerError  %+v", 500, o.Payload)
}

func (o *CreateTokenInternalServerError) String() string {
	return fmt.Sprintf("[POST /v1/tokens][%d] createTokenInternalServerError  %+v", 500, o.Payload)
}

func (o *CreateTokenInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateTokenInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
