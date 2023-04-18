// Code generated by go-swagger; DO NOT EDIT.

// Copyright © 2023 VMware, Inc. All Rights Reserved.
// Code licenced under: BSD-2

package v_r_l_i

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/vmware/vcf-sdk-go/models"
)

// GETVrlisReader is a Reader for the GETVrlis structure.
type GETVrlisReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GETVrlisReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGETVrlisOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGETVrlisOK creates a GETVrlisOK with default headers values
func NewGETVrlisOK() *GETVrlisOK {
	return &GETVrlisOK{}
}

/*
GETVrlisOK describes a response with status code 200, with default header values.

OK
*/
type GETVrlisOK struct {
	Payload *models.PageOfVrli
}

// IsSuccess returns true when this get vrlis o k response has a 2xx status code
func (o *GETVrlisOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get vrlis o k response has a 3xx status code
func (o *GETVrlisOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get vrlis o k response has a 4xx status code
func (o *GETVrlisOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get vrlis o k response has a 5xx status code
func (o *GETVrlisOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get vrlis o k response a status code equal to that given
func (o *GETVrlisOK) IsCode(code int) bool {
	return code == 200
}

func (o *GETVrlisOK) Error() string {
	return fmt.Sprintf("[GET /v1/vrlis][%d] getVrlisOK  %+v", 200, o.Payload)
}

func (o *GETVrlisOK) String() string {
	return fmt.Sprintf("[GET /v1/vrlis][%d] getVrlisOK  %+v", 200, o.Payload)
}

func (o *GETVrlisOK) GetPayload() *models.PageOfVrli {
	return o.Payload
}

func (o *GETVrlisOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PageOfVrli)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
