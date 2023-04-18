// Code generated by go-swagger; DO NOT EDIT.

// Copyright © 2023 VMware, Inc. All Rights Reserved.
// Code licenced under: BSD-2

package v_r_o_ps

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/vmware/vcf-sdk-go/models"
)

// GETVropsesReader is a Reader for the GETVropses structure.
type GETVropsesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GETVropsesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGETVropsesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGETVropsesOK creates a GETVropsesOK with default headers values
func NewGETVropsesOK() *GETVropsesOK {
	return &GETVropsesOK{}
}

/*
GETVropsesOK describes a response with status code 200, with default header values.

OK
*/
type GETVropsesOK struct {
	Payload *models.PageOfVROPS
}

// IsSuccess returns true when this get vropses o k response has a 2xx status code
func (o *GETVropsesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get vropses o k response has a 3xx status code
func (o *GETVropsesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get vropses o k response has a 4xx status code
func (o *GETVropsesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get vropses o k response has a 5xx status code
func (o *GETVropsesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get vropses o k response a status code equal to that given
func (o *GETVropsesOK) IsCode(code int) bool {
	return code == 200
}

func (o *GETVropsesOK) Error() string {
	return fmt.Sprintf("[GET /v1/vropses][%d] getVropsesOK  %+v", 200, o.Payload)
}

func (o *GETVropsesOK) String() string {
	return fmt.Sprintf("[GET /v1/vropses][%d] getVropsesOK  %+v", 200, o.Payload)
}

func (o *GETVropsesOK) GetPayload() *models.PageOfVROPS {
	return o.Payload
}

func (o *GETVropsesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PageOfVROPS)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
