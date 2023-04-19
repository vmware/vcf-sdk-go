// Code generated by go-swagger; DO NOT EDIT.

// Copyright © 2023 VMware, Inc. All Rights Reserved.
// Code licenced under: BSD-2

package nsxt_clusters

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/vmware/vcf-sdk-go/models"
)

// GETCriteria1Reader is a Reader for the GETCriteria1 structure.
type GETCriteria1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GETCriteria1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGETCriteria1OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGETCriteria1OK creates a GETCriteria1OK with default headers values
func NewGETCriteria1OK() *GETCriteria1OK {
	return &GETCriteria1OK{}
}

/*
GETCriteria1OK describes a response with status code 200, with default header values.

Ok
*/
type GETCriteria1OK struct {
	Payload *models.PageOfNsxTCriterion
}

// IsSuccess returns true when this get criteria1 o k response has a 2xx status code
func (o *GETCriteria1OK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get criteria1 o k response has a 3xx status code
func (o *GETCriteria1OK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get criteria1 o k response has a 4xx status code
func (o *GETCriteria1OK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get criteria1 o k response has a 5xx status code
func (o *GETCriteria1OK) IsServerError() bool {
	return false
}

// IsCode returns true when this get criteria1 o k response a status code equal to that given
func (o *GETCriteria1OK) IsCode(code int) bool {
	return code == 200
}

func (o *GETCriteria1OK) Error() string {
	return fmt.Sprintf("[GET /v1/nsxt-clusters/criteria][%d] getCriteria1OK  %+v", 200, o.Payload)
}

func (o *GETCriteria1OK) String() string {
	return fmt.Sprintf("[GET /v1/nsxt-clusters/criteria][%d] getCriteria1OK  %+v", 200, o.Payload)
}

func (o *GETCriteria1OK) GetPayload() *models.PageOfNsxTCriterion {
	return o.Payload
}

func (o *GETCriteria1OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PageOfNsxTCriterion)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}