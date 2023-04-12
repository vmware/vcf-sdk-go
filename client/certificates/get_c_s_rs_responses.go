// Code generated by go-swagger; DO NOT EDIT.

package certificates

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"vcf-sdk-go/models"
)

// GetCSRsReader is a Reader for the GetCSRs structure.
type GetCSRsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCSRsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetCSRsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetCSRsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetCSRsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetCSRsOK creates a GetCSRsOK with default headers values
func NewGetCSRsOK() *GetCSRsOK {
	return &GetCSRsOK{}
}

/*
GetCSRsOK describes a response with status code 200, with default header values.

OK
*/
type GetCSRsOK struct {
	Payload *models.PageOfCsr
}

// IsSuccess returns true when this get c s rs o k response has a 2xx status code
func (o *GetCSRsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get c s rs o k response has a 3xx status code
func (o *GetCSRsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get c s rs o k response has a 4xx status code
func (o *GetCSRsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get c s rs o k response has a 5xx status code
func (o *GetCSRsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get c s rs o k response a status code equal to that given
func (o *GetCSRsOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetCSRsOK) Error() string {
	return fmt.Sprintf("[GET /v1/domains/{domainName}/csrs][%d] getCSRsOK  %+v", 200, o.Payload)
}

func (o *GetCSRsOK) String() string {
	return fmt.Sprintf("[GET /v1/domains/{domainName}/csrs][%d] getCSRsOK  %+v", 200, o.Payload)
}

func (o *GetCSRsOK) GetPayload() *models.PageOfCsr {
	return o.Payload
}

func (o *GetCSRsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PageOfCsr)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCSRsNotFound creates a GetCSRsNotFound with default headers values
func NewGetCSRsNotFound() *GetCSRsNotFound {
	return &GetCSRsNotFound{}
}

/*
GetCSRsNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetCSRsNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this get c s rs not found response has a 2xx status code
func (o *GetCSRsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get c s rs not found response has a 3xx status code
func (o *GetCSRsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get c s rs not found response has a 4xx status code
func (o *GetCSRsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get c s rs not found response has a 5xx status code
func (o *GetCSRsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get c s rs not found response a status code equal to that given
func (o *GetCSRsNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetCSRsNotFound) Error() string {
	return fmt.Sprintf("[GET /v1/domains/{domainName}/csrs][%d] getCSRsNotFound  %+v", 404, o.Payload)
}

func (o *GetCSRsNotFound) String() string {
	return fmt.Sprintf("[GET /v1/domains/{domainName}/csrs][%d] getCSRsNotFound  %+v", 404, o.Payload)
}

func (o *GetCSRsNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetCSRsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCSRsInternalServerError creates a GetCSRsInternalServerError with default headers values
func NewGetCSRsInternalServerError() *GetCSRsInternalServerError {
	return &GetCSRsInternalServerError{}
}

/*
GetCSRsInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type GetCSRsInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this get c s rs internal server error response has a 2xx status code
func (o *GetCSRsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get c s rs internal server error response has a 3xx status code
func (o *GetCSRsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get c s rs internal server error response has a 4xx status code
func (o *GetCSRsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get c s rs internal server error response has a 5xx status code
func (o *GetCSRsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get c s rs internal server error response a status code equal to that given
func (o *GetCSRsInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetCSRsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/domains/{domainName}/csrs][%d] getCSRsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetCSRsInternalServerError) String() string {
	return fmt.Sprintf("[GET /v1/domains/{domainName}/csrs][%d] getCSRsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetCSRsInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetCSRsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
