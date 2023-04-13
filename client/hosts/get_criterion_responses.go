// Code generated by go-swagger; DO NOT EDIT.

// Copyright © 2023 VMware, Inc. All Rights Reserved.
// Code licenced under: BSD-2

package hosts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"vcf-sdk-go/models"
)

// GETCriterionReader is a Reader for the GETCriterion structure.
type GETCriterionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GETCriterionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGETCriterionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGETCriterionNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGETCriterionOK creates a GETCriterionOK with default headers values
func NewGETCriterionOK() *GETCriterionOK {
	return &GETCriterionOK{}
}

/*
GETCriterionOK describes a response with status code 200, with default header values.

Ok
*/
type GETCriterionOK struct {
	Payload *models.HostCriterion
}

// IsSuccess returns true when this get criterion o k response has a 2xx status code
func (o *GETCriterionOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get criterion o k response has a 3xx status code
func (o *GETCriterionOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get criterion o k response has a 4xx status code
func (o *GETCriterionOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get criterion o k response has a 5xx status code
func (o *GETCriterionOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get criterion o k response a status code equal to that given
func (o *GETCriterionOK) IsCode(code int) bool {
	return code == 200
}

func (o *GETCriterionOK) Error() string {
	return fmt.Sprintf("[GET /v1/hosts/criteria/{name}][%d] getCriterionOK  %+v", 200, o.Payload)
}

func (o *GETCriterionOK) String() string {
	return fmt.Sprintf("[GET /v1/hosts/criteria/{name}][%d] getCriterionOK  %+v", 200, o.Payload)
}

func (o *GETCriterionOK) GetPayload() *models.HostCriterion {
	return o.Payload
}

func (o *GETCriterionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HostCriterion)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGETCriterionNotFound creates a GETCriterionNotFound with default headers values
func NewGETCriterionNotFound() *GETCriterionNotFound {
	return &GETCriterionNotFound{}
}

/*
GETCriterionNotFound describes a response with status code 404, with default header values.

Criterion Not Found
*/
type GETCriterionNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this get criterion not found response has a 2xx status code
func (o *GETCriterionNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get criterion not found response has a 3xx status code
func (o *GETCriterionNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get criterion not found response has a 4xx status code
func (o *GETCriterionNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get criterion not found response has a 5xx status code
func (o *GETCriterionNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get criterion not found response a status code equal to that given
func (o *GETCriterionNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GETCriterionNotFound) Error() string {
	return fmt.Sprintf("[GET /v1/hosts/criteria/{name}][%d] getCriterionNotFound  %+v", 404, o.Payload)
}

func (o *GETCriterionNotFound) String() string {
	return fmt.Sprintf("[GET /v1/hosts/criteria/{name}][%d] getCriterionNotFound  %+v", 404, o.Payload)
}

func (o *GETCriterionNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GETCriterionNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
