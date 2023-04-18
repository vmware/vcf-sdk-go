// Code generated by go-swagger; DO NOT EDIT.

// Copyright © 2023 VMware, Inc. All Rights Reserved.
// Code licenced under: BSD-2

package vra

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/vmware/vcf-sdk-go/models"
)

// GETVrasReader is a Reader for the GETVras structure.
type GETVrasReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GETVrasReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGETVrasOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGETVrasOK creates a GETVrasOK with default headers values
func NewGETVrasOK() *GETVrasOK {
	return &GETVrasOK{}
}

/*
GETVrasOK describes a response with status code 200, with default header values.

OK
*/
type GETVrasOK struct {
	Payload *models.PageOfVRA
}

// IsSuccess returns true when this get vras o k response has a 2xx status code
func (o *GETVrasOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get vras o k response has a 3xx status code
func (o *GETVrasOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get vras o k response has a 4xx status code
func (o *GETVrasOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get vras o k response has a 5xx status code
func (o *GETVrasOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get vras o k response a status code equal to that given
func (o *GETVrasOK) IsCode(code int) bool {
	return code == 200
}

func (o *GETVrasOK) Error() string {
	return fmt.Sprintf("[GET /v1/vras][%d] getVrasOK  %+v", 200, o.Payload)
}

func (o *GETVrasOK) String() string {
	return fmt.Sprintf("[GET /v1/vras][%d] getVrasOK  %+v", 200, o.Payload)
}

func (o *GETVrasOK) GetPayload() *models.PageOfVRA {
	return o.Payload
}

func (o *GETVrasOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PageOfVRA)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
