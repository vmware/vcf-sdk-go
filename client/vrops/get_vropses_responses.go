// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2023 VMware, Inc.
// SPDX-License-Identifier: BSD-2-Clause

package vrops

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/vmware/vcf-sdk-go/models"
)

// GetVropsesReader is a Reader for the GetVropses structure.
type GetVropsesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetVropsesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetVropsesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetVropsesOK creates a GetVropsesOK with default headers values
func NewGetVropsesOK() *GetVropsesOK {
	return &GetVropsesOK{}
}

/*
GetVropsesOK describes a response with status code 200, with default header values.

OK
*/
type GetVropsesOK struct {
	Payload *models.PageOfVROPS
}

// IsSuccess returns true when this get vropses o k response has a 2xx status code
func (o *GetVropsesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get vropses o k response has a 3xx status code
func (o *GetVropsesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get vropses o k response has a 4xx status code
func (o *GetVropsesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get vropses o k response has a 5xx status code
func (o *GetVropsesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get vropses o k response a status code equal to that given
func (o *GetVropsesOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetVropsesOK) Error() string {
	return fmt.Sprintf("[GET /v1/vropses][%d] getVropsesOK  %+v", 200, o.Payload)
}

func (o *GetVropsesOK) String() string {
	return fmt.Sprintf("[GET /v1/vropses][%d] getVropsesOK  %+v", 200, o.Payload)
}

func (o *GetVropsesOK) GetPayload() *models.PageOfVROPS {
	return o.Payload
}

func (o *GetVropsesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PageOfVROPS)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
