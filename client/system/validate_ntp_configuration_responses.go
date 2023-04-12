// Code generated by go-swagger; DO NOT EDIT.

package system

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"vcf-sdk-go/models"
)

// ValidateNtpConfigurationReader is a Reader for the ValidateNtpConfiguration structure.
type ValidateNtpConfigurationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ValidateNtpConfigurationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewValidateNtpConfigurationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewValidateNtpConfigurationAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewValidateNtpConfigurationBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewValidateNtpConfigurationInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewValidateNtpConfigurationOK creates a ValidateNtpConfigurationOK with default headers values
func NewValidateNtpConfigurationOK() *ValidateNtpConfigurationOK {
	return &ValidateNtpConfigurationOK{}
}

/*
ValidateNtpConfigurationOK describes a response with status code 200, with default header values.

OK
*/
type ValidateNtpConfigurationOK struct {
	Payload *models.Validation
}

// IsSuccess returns true when this validate ntp configuration o k response has a 2xx status code
func (o *ValidateNtpConfigurationOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this validate ntp configuration o k response has a 3xx status code
func (o *ValidateNtpConfigurationOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this validate ntp configuration o k response has a 4xx status code
func (o *ValidateNtpConfigurationOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this validate ntp configuration o k response has a 5xx status code
func (o *ValidateNtpConfigurationOK) IsServerError() bool {
	return false
}

// IsCode returns true when this validate ntp configuration o k response a status code equal to that given
func (o *ValidateNtpConfigurationOK) IsCode(code int) bool {
	return code == 200
}

func (o *ValidateNtpConfigurationOK) Error() string {
	return fmt.Sprintf("[POST /v1/system/ntp-configuration/validations][%d] validateNtpConfigurationOK  %+v", 200, o.Payload)
}

func (o *ValidateNtpConfigurationOK) String() string {
	return fmt.Sprintf("[POST /v1/system/ntp-configuration/validations][%d] validateNtpConfigurationOK  %+v", 200, o.Payload)
}

func (o *ValidateNtpConfigurationOK) GetPayload() *models.Validation {
	return o.Payload
}

func (o *ValidateNtpConfigurationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Validation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewValidateNtpConfigurationAccepted creates a ValidateNtpConfigurationAccepted with default headers values
func NewValidateNtpConfigurationAccepted() *ValidateNtpConfigurationAccepted {
	return &ValidateNtpConfigurationAccepted{}
}

/*
ValidateNtpConfigurationAccepted describes a response with status code 202, with default header values.

Accepted
*/
type ValidateNtpConfigurationAccepted struct {
	Payload *models.Validation
}

// IsSuccess returns true when this validate ntp configuration accepted response has a 2xx status code
func (o *ValidateNtpConfigurationAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this validate ntp configuration accepted response has a 3xx status code
func (o *ValidateNtpConfigurationAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this validate ntp configuration accepted response has a 4xx status code
func (o *ValidateNtpConfigurationAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this validate ntp configuration accepted response has a 5xx status code
func (o *ValidateNtpConfigurationAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this validate ntp configuration accepted response a status code equal to that given
func (o *ValidateNtpConfigurationAccepted) IsCode(code int) bool {
	return code == 202
}

func (o *ValidateNtpConfigurationAccepted) Error() string {
	return fmt.Sprintf("[POST /v1/system/ntp-configuration/validations][%d] validateNtpConfigurationAccepted  %+v", 202, o.Payload)
}

func (o *ValidateNtpConfigurationAccepted) String() string {
	return fmt.Sprintf("[POST /v1/system/ntp-configuration/validations][%d] validateNtpConfigurationAccepted  %+v", 202, o.Payload)
}

func (o *ValidateNtpConfigurationAccepted) GetPayload() *models.Validation {
	return o.Payload
}

func (o *ValidateNtpConfigurationAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Validation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewValidateNtpConfigurationBadRequest creates a ValidateNtpConfigurationBadRequest with default headers values
func NewValidateNtpConfigurationBadRequest() *ValidateNtpConfigurationBadRequest {
	return &ValidateNtpConfigurationBadRequest{}
}

/*
ValidateNtpConfigurationBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type ValidateNtpConfigurationBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this validate ntp configuration bad request response has a 2xx status code
func (o *ValidateNtpConfigurationBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this validate ntp configuration bad request response has a 3xx status code
func (o *ValidateNtpConfigurationBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this validate ntp configuration bad request response has a 4xx status code
func (o *ValidateNtpConfigurationBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this validate ntp configuration bad request response has a 5xx status code
func (o *ValidateNtpConfigurationBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this validate ntp configuration bad request response a status code equal to that given
func (o *ValidateNtpConfigurationBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *ValidateNtpConfigurationBadRequest) Error() string {
	return fmt.Sprintf("[POST /v1/system/ntp-configuration/validations][%d] validateNtpConfigurationBadRequest  %+v", 400, o.Payload)
}

func (o *ValidateNtpConfigurationBadRequest) String() string {
	return fmt.Sprintf("[POST /v1/system/ntp-configuration/validations][%d] validateNtpConfigurationBadRequest  %+v", 400, o.Payload)
}

func (o *ValidateNtpConfigurationBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *ValidateNtpConfigurationBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewValidateNtpConfigurationInternalServerError creates a ValidateNtpConfigurationInternalServerError with default headers values
func NewValidateNtpConfigurationInternalServerError() *ValidateNtpConfigurationInternalServerError {
	return &ValidateNtpConfigurationInternalServerError{}
}

/*
ValidateNtpConfigurationInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type ValidateNtpConfigurationInternalServerError struct {
}

// IsSuccess returns true when this validate ntp configuration internal server error response has a 2xx status code
func (o *ValidateNtpConfigurationInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this validate ntp configuration internal server error response has a 3xx status code
func (o *ValidateNtpConfigurationInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this validate ntp configuration internal server error response has a 4xx status code
func (o *ValidateNtpConfigurationInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this validate ntp configuration internal server error response has a 5xx status code
func (o *ValidateNtpConfigurationInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this validate ntp configuration internal server error response a status code equal to that given
func (o *ValidateNtpConfigurationInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *ValidateNtpConfigurationInternalServerError) Error() string {
	return fmt.Sprintf("[POST /v1/system/ntp-configuration/validations][%d] validateNtpConfigurationInternalServerError ", 500)
}

func (o *ValidateNtpConfigurationInternalServerError) String() string {
	return fmt.Sprintf("[POST /v1/system/ntp-configuration/validations][%d] validateNtpConfigurationInternalServerError ", 500)
}

func (o *ValidateNtpConfigurationInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
