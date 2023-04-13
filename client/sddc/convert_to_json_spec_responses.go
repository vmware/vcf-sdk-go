// Code generated by go-swagger; DO NOT EDIT.

// Copyright © 2023 VMware, Inc. All Rights Reserved.
// Code licenced under: BSD-2

package sddc

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"vcf-sdk-go/models"
)

// ConvertToJSONSpecReader is a Reader for the ConvertToJSONSpec structure.
type ConvertToJSONSpecReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ConvertToJSONSpecReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewConvertToJSONSpecOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewConvertToJSONSpecBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewConvertToJSONSpecNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewConvertToJSONSpecInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 501:
		result := NewConvertToJSONSpecNotImplemented()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewConvertToJSONSpecOK creates a ConvertToJSONSpecOK with default headers values
func NewConvertToJSONSpecOK() *ConvertToJSONSpecOK {
	return &ConvertToJSONSpecOK{}
}

/*
ConvertToJSONSpecOK describes a response with status code 200, with default header values.

OK
*/
type ConvertToJSONSpecOK struct {
	Payload *models.SDDCSpec
}

// IsSuccess returns true when this convert to Json spec o k response has a 2xx status code
func (o *ConvertToJSONSpecOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this convert to Json spec o k response has a 3xx status code
func (o *ConvertToJSONSpecOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this convert to Json spec o k response has a 4xx status code
func (o *ConvertToJSONSpecOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this convert to Json spec o k response has a 5xx status code
func (o *ConvertToJSONSpecOK) IsServerError() bool {
	return false
}

// IsCode returns true when this convert to Json spec o k response a status code equal to that given
func (o *ConvertToJSONSpecOK) IsCode(code int) bool {
	return code == 200
}

func (o *ConvertToJSONSpecOK) Error() string {
	return fmt.Sprintf("[POST /v1/system/sddc-spec-converter][%d] convertToJsonSpecOK  %+v", 200, o.Payload)
}

func (o *ConvertToJSONSpecOK) String() string {
	return fmt.Sprintf("[POST /v1/system/sddc-spec-converter][%d] convertToJsonSpecOK  %+v", 200, o.Payload)
}

func (o *ConvertToJSONSpecOK) GetPayload() *models.SDDCSpec {
	return o.Payload
}

func (o *ConvertToJSONSpecOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SDDCSpec)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewConvertToJSONSpecBadRequest creates a ConvertToJSONSpecBadRequest with default headers values
func NewConvertToJSONSpecBadRequest() *ConvertToJSONSpecBadRequest {
	return &ConvertToJSONSpecBadRequest{}
}

/*
ConvertToJSONSpecBadRequest describes a response with status code 400, with default header values.

Converting SDDC specification failed
*/
type ConvertToJSONSpecBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this convert to Json spec bad request response has a 2xx status code
func (o *ConvertToJSONSpecBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this convert to Json spec bad request response has a 3xx status code
func (o *ConvertToJSONSpecBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this convert to Json spec bad request response has a 4xx status code
func (o *ConvertToJSONSpecBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this convert to Json spec bad request response has a 5xx status code
func (o *ConvertToJSONSpecBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this convert to Json spec bad request response a status code equal to that given
func (o *ConvertToJSONSpecBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *ConvertToJSONSpecBadRequest) Error() string {
	return fmt.Sprintf("[POST /v1/system/sddc-spec-converter][%d] convertToJsonSpecBadRequest  %+v", 400, o.Payload)
}

func (o *ConvertToJSONSpecBadRequest) String() string {
	return fmt.Sprintf("[POST /v1/system/sddc-spec-converter][%d] convertToJsonSpecBadRequest  %+v", 400, o.Payload)
}

func (o *ConvertToJSONSpecBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *ConvertToJSONSpecBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewConvertToJSONSpecNotFound creates a ConvertToJSONSpecNotFound with default headers values
func NewConvertToJSONSpecNotFound() *ConvertToJSONSpecNotFound {
	return &ConvertToJSONSpecNotFound{}
}

/*
ConvertToJSONSpecNotFound describes a response with status code 404, with default header values.

Not Found
*/
type ConvertToJSONSpecNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this convert to Json spec not found response has a 2xx status code
func (o *ConvertToJSONSpecNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this convert to Json spec not found response has a 3xx status code
func (o *ConvertToJSONSpecNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this convert to Json spec not found response has a 4xx status code
func (o *ConvertToJSONSpecNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this convert to Json spec not found response has a 5xx status code
func (o *ConvertToJSONSpecNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this convert to Json spec not found response a status code equal to that given
func (o *ConvertToJSONSpecNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *ConvertToJSONSpecNotFound) Error() string {
	return fmt.Sprintf("[POST /v1/system/sddc-spec-converter][%d] convertToJsonSpecNotFound  %+v", 404, o.Payload)
}

func (o *ConvertToJSONSpecNotFound) String() string {
	return fmt.Sprintf("[POST /v1/system/sddc-spec-converter][%d] convertToJsonSpecNotFound  %+v", 404, o.Payload)
}

func (o *ConvertToJSONSpecNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *ConvertToJSONSpecNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewConvertToJSONSpecInternalServerError creates a ConvertToJSONSpecInternalServerError with default headers values
func NewConvertToJSONSpecInternalServerError() *ConvertToJSONSpecInternalServerError {
	return &ConvertToJSONSpecInternalServerError{}
}

/*
ConvertToJSONSpecInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type ConvertToJSONSpecInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this convert to Json spec internal server error response has a 2xx status code
func (o *ConvertToJSONSpecInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this convert to Json spec internal server error response has a 3xx status code
func (o *ConvertToJSONSpecInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this convert to Json spec internal server error response has a 4xx status code
func (o *ConvertToJSONSpecInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this convert to Json spec internal server error response has a 5xx status code
func (o *ConvertToJSONSpecInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this convert to Json spec internal server error response a status code equal to that given
func (o *ConvertToJSONSpecInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *ConvertToJSONSpecInternalServerError) Error() string {
	return fmt.Sprintf("[POST /v1/system/sddc-spec-converter][%d] convertToJsonSpecInternalServerError  %+v", 500, o.Payload)
}

func (o *ConvertToJSONSpecInternalServerError) String() string {
	return fmt.Sprintf("[POST /v1/system/sddc-spec-converter][%d] convertToJsonSpecInternalServerError  %+v", 500, o.Payload)
}

func (o *ConvertToJSONSpecInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *ConvertToJSONSpecInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewConvertToJSONSpecNotImplemented creates a ConvertToJSONSpecNotImplemented with default headers values
func NewConvertToJSONSpecNotImplemented() *ConvertToJSONSpecNotImplemented {
	return &ConvertToJSONSpecNotImplemented{}
}

/*
ConvertToJSONSpecNotImplemented describes a response with status code 501, with default header values.

Not Implemented
*/
type ConvertToJSONSpecNotImplemented struct {
	Payload *models.Error
}

// IsSuccess returns true when this convert to Json spec not implemented response has a 2xx status code
func (o *ConvertToJSONSpecNotImplemented) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this convert to Json spec not implemented response has a 3xx status code
func (o *ConvertToJSONSpecNotImplemented) IsRedirect() bool {
	return false
}

// IsClientError returns true when this convert to Json spec not implemented response has a 4xx status code
func (o *ConvertToJSONSpecNotImplemented) IsClientError() bool {
	return false
}

// IsServerError returns true when this convert to Json spec not implemented response has a 5xx status code
func (o *ConvertToJSONSpecNotImplemented) IsServerError() bool {
	return true
}

// IsCode returns true when this convert to Json spec not implemented response a status code equal to that given
func (o *ConvertToJSONSpecNotImplemented) IsCode(code int) bool {
	return code == 501
}

func (o *ConvertToJSONSpecNotImplemented) Error() string {
	return fmt.Sprintf("[POST /v1/system/sddc-spec-converter][%d] convertToJsonSpecNotImplemented  %+v", 501, o.Payload)
}

func (o *ConvertToJSONSpecNotImplemented) String() string {
	return fmt.Sprintf("[POST /v1/system/sddc-spec-converter][%d] convertToJsonSpecNotImplemented  %+v", 501, o.Payload)
}

func (o *ConvertToJSONSpecNotImplemented) GetPayload() *models.Error {
	return o.Payload
}

func (o *ConvertToJSONSpecNotImplemented) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
