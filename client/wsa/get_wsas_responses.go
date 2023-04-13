// Code generated by go-swagger; DO NOT EDIT.

// Copyright © 2023 VMware, Inc. All Rights Reserved.
// Code licenced under: BSD-2

package wsa

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"vcf-sdk-go/models"
)

// GETWsasReader is a Reader for the GETWsas structure.
type GETWsasReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GETWsasReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGETWsasOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGETWsasOK creates a GETWsasOK with default headers values
func NewGETWsasOK() *GETWsasOK {
	return &GETWsasOK{}
}

/*
GETWsasOK describes a response with status code 200, with default header values.

OK
*/
type GETWsasOK struct {
	Payload *models.PageOfWSA
}

// IsSuccess returns true when this get wsas o k response has a 2xx status code
func (o *GETWsasOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get wsas o k response has a 3xx status code
func (o *GETWsasOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get wsas o k response has a 4xx status code
func (o *GETWsasOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get wsas o k response has a 5xx status code
func (o *GETWsasOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get wsas o k response a status code equal to that given
func (o *GETWsasOK) IsCode(code int) bool {
	return code == 200
}

func (o *GETWsasOK) Error() string {
	return fmt.Sprintf("[GET /v1/wsas][%d] getWsasOK  %+v", 200, o.Payload)
}

func (o *GETWsasOK) String() string {
	return fmt.Sprintf("[GET /v1/wsas][%d] getWsasOK  %+v", 200, o.Payload)
}

func (o *GETWsasOK) GetPayload() *models.PageOfWSA {
	return o.Payload
}

func (o *GETWsasOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PageOfWSA)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
