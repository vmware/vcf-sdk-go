// Code generated by go-swagger; DO NOT EDIT.

// Copyright © 2023 VMware, Inc. All Rights Reserved.
// Code licenced under: BSD-2

package v_r_s_l_c_m

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/vmware/vcf-sdk-go/models"
)

// DeployVRSLCMReader is a Reader for the DeployVRSLCM structure.
type DeployVRSLCMReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeployVRSLCMReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewDeployVRSLCMAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeployVRSLCMBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 405:
		result := NewDeployVRSLCMMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeployVRSLCMInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeployVRSLCMAccepted creates a DeployVRSLCMAccepted with default headers values
func NewDeployVRSLCMAccepted() *DeployVRSLCMAccepted {
	return &DeployVRSLCMAccepted{}
}

/*
DeployVRSLCMAccepted describes a response with status code 202, with default header values.

Accepted
*/
type DeployVRSLCMAccepted struct {
	Payload *models.Task
}

// IsSuccess returns true when this deploy Vrslcm accepted response has a 2xx status code
func (o *DeployVRSLCMAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this deploy Vrslcm accepted response has a 3xx status code
func (o *DeployVRSLCMAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this deploy Vrslcm accepted response has a 4xx status code
func (o *DeployVRSLCMAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this deploy Vrslcm accepted response has a 5xx status code
func (o *DeployVRSLCMAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this deploy Vrslcm accepted response a status code equal to that given
func (o *DeployVRSLCMAccepted) IsCode(code int) bool {
	return code == 202
}

func (o *DeployVRSLCMAccepted) Error() string {
	return fmt.Sprintf("[POST /v1/vrslcms][%d] deployVrslcmAccepted  %+v", 202, o.Payload)
}

func (o *DeployVRSLCMAccepted) String() string {
	return fmt.Sprintf("[POST /v1/vrslcms][%d] deployVrslcmAccepted  %+v", 202, o.Payload)
}

func (o *DeployVRSLCMAccepted) GetPayload() *models.Task {
	return o.Payload
}

func (o *DeployVRSLCMAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Task)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeployVRSLCMBadRequest creates a DeployVRSLCMBadRequest with default headers values
func NewDeployVRSLCMBadRequest() *DeployVRSLCMBadRequest {
	return &DeployVRSLCMBadRequest{}
}

/*
DeployVRSLCMBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type DeployVRSLCMBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this deploy Vrslcm bad request response has a 2xx status code
func (o *DeployVRSLCMBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this deploy Vrslcm bad request response has a 3xx status code
func (o *DeployVRSLCMBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this deploy Vrslcm bad request response has a 4xx status code
func (o *DeployVRSLCMBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this deploy Vrslcm bad request response has a 5xx status code
func (o *DeployVRSLCMBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this deploy Vrslcm bad request response a status code equal to that given
func (o *DeployVRSLCMBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *DeployVRSLCMBadRequest) Error() string {
	return fmt.Sprintf("[POST /v1/vrslcms][%d] deployVrslcmBadRequest  %+v", 400, o.Payload)
}

func (o *DeployVRSLCMBadRequest) String() string {
	return fmt.Sprintf("[POST /v1/vrslcms][%d] deployVrslcmBadRequest  %+v", 400, o.Payload)
}

func (o *DeployVRSLCMBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeployVRSLCMBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeployVRSLCMMethodNotAllowed creates a DeployVRSLCMMethodNotAllowed with default headers values
func NewDeployVRSLCMMethodNotAllowed() *DeployVRSLCMMethodNotAllowed {
	return &DeployVRSLCMMethodNotAllowed{}
}

/*
DeployVRSLCMMethodNotAllowed describes a response with status code 405, with default header values.

Method Not Allowed
*/
type DeployVRSLCMMethodNotAllowed struct {
	Payload *models.Error
}

// IsSuccess returns true when this deploy Vrslcm method not allowed response has a 2xx status code
func (o *DeployVRSLCMMethodNotAllowed) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this deploy Vrslcm method not allowed response has a 3xx status code
func (o *DeployVRSLCMMethodNotAllowed) IsRedirect() bool {
	return false
}

// IsClientError returns true when this deploy Vrslcm method not allowed response has a 4xx status code
func (o *DeployVRSLCMMethodNotAllowed) IsClientError() bool {
	return true
}

// IsServerError returns true when this deploy Vrslcm method not allowed response has a 5xx status code
func (o *DeployVRSLCMMethodNotAllowed) IsServerError() bool {
	return false
}

// IsCode returns true when this deploy Vrslcm method not allowed response a status code equal to that given
func (o *DeployVRSLCMMethodNotAllowed) IsCode(code int) bool {
	return code == 405
}

func (o *DeployVRSLCMMethodNotAllowed) Error() string {
	return fmt.Sprintf("[POST /v1/vrslcms][%d] deployVrslcmMethodNotAllowed  %+v", 405, o.Payload)
}

func (o *DeployVRSLCMMethodNotAllowed) String() string {
	return fmt.Sprintf("[POST /v1/vrslcms][%d] deployVrslcmMethodNotAllowed  %+v", 405, o.Payload)
}

func (o *DeployVRSLCMMethodNotAllowed) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeployVRSLCMMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeployVRSLCMInternalServerError creates a DeployVRSLCMInternalServerError with default headers values
func NewDeployVRSLCMInternalServerError() *DeployVRSLCMInternalServerError {
	return &DeployVRSLCMInternalServerError{}
}

/*
DeployVRSLCMInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type DeployVRSLCMInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this deploy Vrslcm internal server error response has a 2xx status code
func (o *DeployVRSLCMInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this deploy Vrslcm internal server error response has a 3xx status code
func (o *DeployVRSLCMInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this deploy Vrslcm internal server error response has a 4xx status code
func (o *DeployVRSLCMInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this deploy Vrslcm internal server error response has a 5xx status code
func (o *DeployVRSLCMInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this deploy Vrslcm internal server error response a status code equal to that given
func (o *DeployVRSLCMInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *DeployVRSLCMInternalServerError) Error() string {
	return fmt.Sprintf("[POST /v1/vrslcms][%d] deployVrslcmInternalServerError  %+v", 500, o.Payload)
}

func (o *DeployVRSLCMInternalServerError) String() string {
	return fmt.Sprintf("[POST /v1/vrslcms][%d] deployVrslcmInternalServerError  %+v", 500, o.Payload)
}

func (o *DeployVRSLCMInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeployVRSLCMInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
