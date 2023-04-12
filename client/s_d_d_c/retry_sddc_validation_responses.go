// Code generated by go-swagger; DO NOT EDIT.

package s_d_d_c

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"vcf-sdk-go/models"
)

// RetrySddcValidationReader is a Reader for the RetrySddcValidation structure.
type RetrySddcValidationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RetrySddcValidationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRetrySddcValidationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewRetrySddcValidationBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 405:
		result := NewRetrySddcValidationMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewRetrySddcValidationInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewRetrySddcValidationOK creates a RetrySddcValidationOK with default headers values
func NewRetrySddcValidationOK() *RetrySddcValidationOK {
	return &RetrySddcValidationOK{}
}

/*
RetrySddcValidationOK describes a response with status code 200, with default header values.

Completed
*/
type RetrySddcValidationOK struct {
	Payload *models.Validation
}

// IsSuccess returns true when this retry sddc validation o k response has a 2xx status code
func (o *RetrySddcValidationOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this retry sddc validation o k response has a 3xx status code
func (o *RetrySddcValidationOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this retry sddc validation o k response has a 4xx status code
func (o *RetrySddcValidationOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this retry sddc validation o k response has a 5xx status code
func (o *RetrySddcValidationOK) IsServerError() bool {
	return false
}

// IsCode returns true when this retry sddc validation o k response a status code equal to that given
func (o *RetrySddcValidationOK) IsCode(code int) bool {
	return code == 200
}

func (o *RetrySddcValidationOK) Error() string {
	return fmt.Sprintf("[PATCH /v1/sddcs/validations/{id}][%d] retrySddcValidationOK  %+v", 200, o.Payload)
}

func (o *RetrySddcValidationOK) String() string {
	return fmt.Sprintf("[PATCH /v1/sddcs/validations/{id}][%d] retrySddcValidationOK  %+v", 200, o.Payload)
}

func (o *RetrySddcValidationOK) GetPayload() *models.Validation {
	return o.Payload
}

func (o *RetrySddcValidationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Validation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRetrySddcValidationBadRequest creates a RetrySddcValidationBadRequest with default headers values
func NewRetrySddcValidationBadRequest() *RetrySddcValidationBadRequest {
	return &RetrySddcValidationBadRequest{}
}

/*
RetrySddcValidationBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type RetrySddcValidationBadRequest struct {
}

// IsSuccess returns true when this retry sddc validation bad request response has a 2xx status code
func (o *RetrySddcValidationBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this retry sddc validation bad request response has a 3xx status code
func (o *RetrySddcValidationBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this retry sddc validation bad request response has a 4xx status code
func (o *RetrySddcValidationBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this retry sddc validation bad request response has a 5xx status code
func (o *RetrySddcValidationBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this retry sddc validation bad request response a status code equal to that given
func (o *RetrySddcValidationBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *RetrySddcValidationBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /v1/sddcs/validations/{id}][%d] retrySddcValidationBadRequest ", 400)
}

func (o *RetrySddcValidationBadRequest) String() string {
	return fmt.Sprintf("[PATCH /v1/sddcs/validations/{id}][%d] retrySddcValidationBadRequest ", 400)
}

func (o *RetrySddcValidationBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRetrySddcValidationMethodNotAllowed creates a RetrySddcValidationMethodNotAllowed with default headers values
func NewRetrySddcValidationMethodNotAllowed() *RetrySddcValidationMethodNotAllowed {
	return &RetrySddcValidationMethodNotAllowed{}
}

/*
RetrySddcValidationMethodNotAllowed describes a response with status code 405, with default header values.

Method Not Allowed
*/
type RetrySddcValidationMethodNotAllowed struct {
}

// IsSuccess returns true when this retry sddc validation method not allowed response has a 2xx status code
func (o *RetrySddcValidationMethodNotAllowed) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this retry sddc validation method not allowed response has a 3xx status code
func (o *RetrySddcValidationMethodNotAllowed) IsRedirect() bool {
	return false
}

// IsClientError returns true when this retry sddc validation method not allowed response has a 4xx status code
func (o *RetrySddcValidationMethodNotAllowed) IsClientError() bool {
	return true
}

// IsServerError returns true when this retry sddc validation method not allowed response has a 5xx status code
func (o *RetrySddcValidationMethodNotAllowed) IsServerError() bool {
	return false
}

// IsCode returns true when this retry sddc validation method not allowed response a status code equal to that given
func (o *RetrySddcValidationMethodNotAllowed) IsCode(code int) bool {
	return code == 405
}

func (o *RetrySddcValidationMethodNotAllowed) Error() string {
	return fmt.Sprintf("[PATCH /v1/sddcs/validations/{id}][%d] retrySddcValidationMethodNotAllowed ", 405)
}

func (o *RetrySddcValidationMethodNotAllowed) String() string {
	return fmt.Sprintf("[PATCH /v1/sddcs/validations/{id}][%d] retrySddcValidationMethodNotAllowed ", 405)
}

func (o *RetrySddcValidationMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRetrySddcValidationInternalServerError creates a RetrySddcValidationInternalServerError with default headers values
func NewRetrySddcValidationInternalServerError() *RetrySddcValidationInternalServerError {
	return &RetrySddcValidationInternalServerError{}
}

/*
RetrySddcValidationInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type RetrySddcValidationInternalServerError struct {
}

// IsSuccess returns true when this retry sddc validation internal server error response has a 2xx status code
func (o *RetrySddcValidationInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this retry sddc validation internal server error response has a 3xx status code
func (o *RetrySddcValidationInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this retry sddc validation internal server error response has a 4xx status code
func (o *RetrySddcValidationInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this retry sddc validation internal server error response has a 5xx status code
func (o *RetrySddcValidationInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this retry sddc validation internal server error response a status code equal to that given
func (o *RetrySddcValidationInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *RetrySddcValidationInternalServerError) Error() string {
	return fmt.Sprintf("[PATCH /v1/sddcs/validations/{id}][%d] retrySddcValidationInternalServerError ", 500)
}

func (o *RetrySddcValidationInternalServerError) String() string {
	return fmt.Sprintf("[PATCH /v1/sddcs/validations/{id}][%d] retrySddcValidationInternalServerError ", 500)
}

func (o *RetrySddcValidationInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
