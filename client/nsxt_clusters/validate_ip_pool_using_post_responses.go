// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2023 VMware, Inc.
// SPDX-License-Identifier: BSD-2-Clause

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

// ValidateIPPoolUsingPOSTReader is a Reader for the ValidateIPPoolUsingPOST structure.
type ValidateIPPoolUsingPOSTReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ValidateIPPoolUsingPOSTReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewValidateIPPoolUsingPOSTOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[POST /v1/nsxt-clusters/ip-address-pools/validations] validateIpPoolUsingPOST", response, response.Code())
	}
}

// NewValidateIPPoolUsingPOSTOK creates a ValidateIPPoolUsingPOSTOK with default headers values
func NewValidateIPPoolUsingPOSTOK() *ValidateIPPoolUsingPOSTOK {
	return &ValidateIPPoolUsingPOSTOK{}
}

/*
ValidateIPPoolUsingPOSTOK describes a response with status code 200, with default header values.

OK
*/
type ValidateIPPoolUsingPOSTOK struct {
	Payload *models.Validation
}

// IsSuccess returns true when this validate Ip pool using p o s t o k response has a 2xx status code
func (o *ValidateIPPoolUsingPOSTOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this validate Ip pool using p o s t o k response has a 3xx status code
func (o *ValidateIPPoolUsingPOSTOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this validate Ip pool using p o s t o k response has a 4xx status code
func (o *ValidateIPPoolUsingPOSTOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this validate Ip pool using p o s t o k response has a 5xx status code
func (o *ValidateIPPoolUsingPOSTOK) IsServerError() bool {
	return false
}

// IsCode returns true when this validate Ip pool using p o s t o k response a status code equal to that given
func (o *ValidateIPPoolUsingPOSTOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the validate Ip pool using p o s t o k response
func (o *ValidateIPPoolUsingPOSTOK) Code() int {
	return 200
}

func (o *ValidateIPPoolUsingPOSTOK) Error() string {
	return fmt.Sprintf("[POST /v1/nsxt-clusters/ip-address-pools/validations][%d] validateIpPoolUsingPOSTOK  %+v", 200, o.Payload)
}

func (o *ValidateIPPoolUsingPOSTOK) String() string {
	return fmt.Sprintf("[POST /v1/nsxt-clusters/ip-address-pools/validations][%d] validateIpPoolUsingPOSTOK  %+v", 200, o.Payload)
}

func (o *ValidateIPPoolUsingPOSTOK) GetPayload() *models.Validation {
	return o.Payload
}

func (o *ValidateIPPoolUsingPOSTOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Validation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
