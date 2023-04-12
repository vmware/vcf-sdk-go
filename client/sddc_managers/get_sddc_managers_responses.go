// Code generated by go-swagger; DO NOT EDIT.

package sddc_managers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"vcf-sdk-go/models"
)

// GetSddcManagersReader is a Reader for the GetSddcManagers structure.
type GetSddcManagersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSddcManagersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetSddcManagersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetSddcManagersBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetSddcManagersInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetSddcManagersOK creates a GetSddcManagersOK with default headers values
func NewGetSddcManagersOK() *GetSddcManagersOK {
	return &GetSddcManagersOK{}
}

/*
GetSddcManagersOK describes a response with status code 200, with default header values.

Ok
*/
type GetSddcManagersOK struct {
	Payload *models.PageOfSddcManager
}

// IsSuccess returns true when this get sddc managers o k response has a 2xx status code
func (o *GetSddcManagersOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get sddc managers o k response has a 3xx status code
func (o *GetSddcManagersOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get sddc managers o k response has a 4xx status code
func (o *GetSddcManagersOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get sddc managers o k response has a 5xx status code
func (o *GetSddcManagersOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get sddc managers o k response a status code equal to that given
func (o *GetSddcManagersOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetSddcManagersOK) Error() string {
	return fmt.Sprintf("[GET /v1/sddc-managers][%d] getSddcManagersOK  %+v", 200, o.Payload)
}

func (o *GetSddcManagersOK) String() string {
	return fmt.Sprintf("[GET /v1/sddc-managers][%d] getSddcManagersOK  %+v", 200, o.Payload)
}

func (o *GetSddcManagersOK) GetPayload() *models.PageOfSddcManager {
	return o.Payload
}

func (o *GetSddcManagersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PageOfSddcManager)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSddcManagersBadRequest creates a GetSddcManagersBadRequest with default headers values
func NewGetSddcManagersBadRequest() *GetSddcManagersBadRequest {
	return &GetSddcManagersBadRequest{}
}

/*
GetSddcManagersBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetSddcManagersBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this get sddc managers bad request response has a 2xx status code
func (o *GetSddcManagersBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get sddc managers bad request response has a 3xx status code
func (o *GetSddcManagersBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get sddc managers bad request response has a 4xx status code
func (o *GetSddcManagersBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get sddc managers bad request response has a 5xx status code
func (o *GetSddcManagersBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get sddc managers bad request response a status code equal to that given
func (o *GetSddcManagersBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetSddcManagersBadRequest) Error() string {
	return fmt.Sprintf("[GET /v1/sddc-managers][%d] getSddcManagersBadRequest  %+v", 400, o.Payload)
}

func (o *GetSddcManagersBadRequest) String() string {
	return fmt.Sprintf("[GET /v1/sddc-managers][%d] getSddcManagersBadRequest  %+v", 400, o.Payload)
}

func (o *GetSddcManagersBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetSddcManagersBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSddcManagersInternalServerError creates a GetSddcManagersInternalServerError with default headers values
func NewGetSddcManagersInternalServerError() *GetSddcManagersInternalServerError {
	return &GetSddcManagersInternalServerError{}
}

/*
GetSddcManagersInternalServerError describes a response with status code 500, with default header values.

InternalServerError
*/
type GetSddcManagersInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this get sddc managers internal server error response has a 2xx status code
func (o *GetSddcManagersInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get sddc managers internal server error response has a 3xx status code
func (o *GetSddcManagersInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get sddc managers internal server error response has a 4xx status code
func (o *GetSddcManagersInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get sddc managers internal server error response has a 5xx status code
func (o *GetSddcManagersInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get sddc managers internal server error response a status code equal to that given
func (o *GetSddcManagersInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetSddcManagersInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/sddc-managers][%d] getSddcManagersInternalServerError  %+v", 500, o.Payload)
}

func (o *GetSddcManagersInternalServerError) String() string {
	return fmt.Sprintf("[GET /v1/sddc-managers][%d] getSddcManagersInternalServerError  %+v", 500, o.Payload)
}

func (o *GetSddcManagersInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetSddcManagersInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
