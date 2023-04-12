// Code generated by go-swagger; DO NOT EDIT.

package bundles

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"vcf-sdk-go/models"
)

// GetBundleReader is a Reader for the GetBundle structure.
type GetBundleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetBundleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetBundleOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetBundleNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetBundleInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetBundleOK creates a GetBundleOK with default headers values
func NewGetBundleOK() *GetBundleOK {
	return &GetBundleOK{}
}

/*
GetBundleOK describes a response with status code 200, with default header values.

Ok
*/
type GetBundleOK struct {
	Payload *models.Bundle
}

// IsSuccess returns true when this get bundle o k response has a 2xx status code
func (o *GetBundleOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get bundle o k response has a 3xx status code
func (o *GetBundleOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get bundle o k response has a 4xx status code
func (o *GetBundleOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get bundle o k response has a 5xx status code
func (o *GetBundleOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get bundle o k response a status code equal to that given
func (o *GetBundleOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetBundleOK) Error() string {
	return fmt.Sprintf("[GET /v1/bundles/{id}][%d] getBundleOK  %+v", 200, o.Payload)
}

func (o *GetBundleOK) String() string {
	return fmt.Sprintf("[GET /v1/bundles/{id}][%d] getBundleOK  %+v", 200, o.Payload)
}

func (o *GetBundleOK) GetPayload() *models.Bundle {
	return o.Payload
}

func (o *GetBundleOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Bundle)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetBundleNotFound creates a GetBundleNotFound with default headers values
func NewGetBundleNotFound() *GetBundleNotFound {
	return &GetBundleNotFound{}
}

/*
GetBundleNotFound describes a response with status code 404, with default header values.

Bundle Not Found
*/
type GetBundleNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this get bundle not found response has a 2xx status code
func (o *GetBundleNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get bundle not found response has a 3xx status code
func (o *GetBundleNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get bundle not found response has a 4xx status code
func (o *GetBundleNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get bundle not found response has a 5xx status code
func (o *GetBundleNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get bundle not found response a status code equal to that given
func (o *GetBundleNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetBundleNotFound) Error() string {
	return fmt.Sprintf("[GET /v1/bundles/{id}][%d] getBundleNotFound  %+v", 404, o.Payload)
}

func (o *GetBundleNotFound) String() string {
	return fmt.Sprintf("[GET /v1/bundles/{id}][%d] getBundleNotFound  %+v", 404, o.Payload)
}

func (o *GetBundleNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetBundleNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetBundleInternalServerError creates a GetBundleInternalServerError with default headers values
func NewGetBundleInternalServerError() *GetBundleInternalServerError {
	return &GetBundleInternalServerError{}
}

/*
GetBundleInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetBundleInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this get bundle internal server error response has a 2xx status code
func (o *GetBundleInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get bundle internal server error response has a 3xx status code
func (o *GetBundleInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get bundle internal server error response has a 4xx status code
func (o *GetBundleInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get bundle internal server error response has a 5xx status code
func (o *GetBundleInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get bundle internal server error response a status code equal to that given
func (o *GetBundleInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetBundleInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/bundles/{id}][%d] getBundleInternalServerError  %+v", 500, o.Payload)
}

func (o *GetBundleInternalServerError) String() string {
	return fmt.Sprintf("[GET /v1/bundles/{id}][%d] getBundleInternalServerError  %+v", 500, o.Payload)
}

func (o *GetBundleInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetBundleInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
