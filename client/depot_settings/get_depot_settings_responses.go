// Code generated by go-swagger; DO NOT EDIT.

// Copyright © 2023 VMware, Inc. All Rights Reserved.
// Code licenced under: BSD-2

package depot_settings

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/vmware/vcf-sdk-go/models"
)

// GETDepotSettingsReader is a Reader for the GETDepotSettings structure.
type GETDepotSettingsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GETDepotSettingsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGETDepotSettingsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewGETDepotSettingsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGETDepotSettingsOK creates a GETDepotSettingsOK with default headers values
func NewGETDepotSettingsOK() *GETDepotSettingsOK {
	return &GETDepotSettingsOK{}
}

/*
GETDepotSettingsOK describes a response with status code 200, with default header values.

Ok
*/
type GETDepotSettingsOK struct {
	Payload []*models.DepotSettings
}

// IsSuccess returns true when this get depot settings o k response has a 2xx status code
func (o *GETDepotSettingsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get depot settings o k response has a 3xx status code
func (o *GETDepotSettingsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get depot settings o k response has a 4xx status code
func (o *GETDepotSettingsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get depot settings o k response has a 5xx status code
func (o *GETDepotSettingsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get depot settings o k response a status code equal to that given
func (o *GETDepotSettingsOK) IsCode(code int) bool {
	return code == 200
}

func (o *GETDepotSettingsOK) Error() string {
	return fmt.Sprintf("[GET /v1/system/settings/depot][%d] getDepotSettingsOK  %+v", 200, o.Payload)
}

func (o *GETDepotSettingsOK) String() string {
	return fmt.Sprintf("[GET /v1/system/settings/depot][%d] getDepotSettingsOK  %+v", 200, o.Payload)
}

func (o *GETDepotSettingsOK) GetPayload() []*models.DepotSettings {
	return o.Payload
}

func (o *GETDepotSettingsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGETDepotSettingsInternalServerError creates a GETDepotSettingsInternalServerError with default headers values
func NewGETDepotSettingsInternalServerError() *GETDepotSettingsInternalServerError {
	return &GETDepotSettingsInternalServerError{}
}

/*
GETDepotSettingsInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GETDepotSettingsInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this get depot settings internal server error response has a 2xx status code
func (o *GETDepotSettingsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get depot settings internal server error response has a 3xx status code
func (o *GETDepotSettingsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get depot settings internal server error response has a 4xx status code
func (o *GETDepotSettingsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get depot settings internal server error response has a 5xx status code
func (o *GETDepotSettingsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get depot settings internal server error response a status code equal to that given
func (o *GETDepotSettingsInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GETDepotSettingsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/system/settings/depot][%d] getDepotSettingsInternalServerError  %+v", 500, o.Payload)
}

func (o *GETDepotSettingsInternalServerError) String() string {
	return fmt.Sprintf("[GET /v1/system/settings/depot][%d] getDepotSettingsInternalServerError  %+v", 500, o.Payload)
}

func (o *GETDepotSettingsInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *GETDepotSettingsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
