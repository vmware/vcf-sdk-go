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

// ConfigureCertificateAuthorityReader is a Reader for the ConfigureCertificateAuthority structure.
type ConfigureCertificateAuthorityReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ConfigureCertificateAuthorityReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewConfigureCertificateAuthorityOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewConfigureCertificateAuthorityBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewConfigureCertificateAuthorityNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewConfigureCertificateAuthorityInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewConfigureCertificateAuthorityOK creates a ConfigureCertificateAuthorityOK with default headers values
func NewConfigureCertificateAuthorityOK() *ConfigureCertificateAuthorityOK {
	return &ConfigureCertificateAuthorityOK{}
}

/*
ConfigureCertificateAuthorityOK describes a response with status code 200, with default header values.

OK
*/
type ConfigureCertificateAuthorityOK struct {
	Payload interface{}
}

// IsSuccess returns true when this configure certificate authority o k response has a 2xx status code
func (o *ConfigureCertificateAuthorityOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this configure certificate authority o k response has a 3xx status code
func (o *ConfigureCertificateAuthorityOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this configure certificate authority o k response has a 4xx status code
func (o *ConfigureCertificateAuthorityOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this configure certificate authority o k response has a 5xx status code
func (o *ConfigureCertificateAuthorityOK) IsServerError() bool {
	return false
}

// IsCode returns true when this configure certificate authority o k response a status code equal to that given
func (o *ConfigureCertificateAuthorityOK) IsCode(code int) bool {
	return code == 200
}

func (o *ConfigureCertificateAuthorityOK) Error() string {
	return fmt.Sprintf("[PATCH /v1/certificate-authorities][%d] configureCertificateAuthorityOK  %+v", 200, o.Payload)
}

func (o *ConfigureCertificateAuthorityOK) String() string {
	return fmt.Sprintf("[PATCH /v1/certificate-authorities][%d] configureCertificateAuthorityOK  %+v", 200, o.Payload)
}

func (o *ConfigureCertificateAuthorityOK) GetPayload() interface{} {
	return o.Payload
}

func (o *ConfigureCertificateAuthorityOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewConfigureCertificateAuthorityBadRequest creates a ConfigureCertificateAuthorityBadRequest with default headers values
func NewConfigureCertificateAuthorityBadRequest() *ConfigureCertificateAuthorityBadRequest {
	return &ConfigureCertificateAuthorityBadRequest{}
}

/*
ConfigureCertificateAuthorityBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type ConfigureCertificateAuthorityBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this configure certificate authority bad request response has a 2xx status code
func (o *ConfigureCertificateAuthorityBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this configure certificate authority bad request response has a 3xx status code
func (o *ConfigureCertificateAuthorityBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this configure certificate authority bad request response has a 4xx status code
func (o *ConfigureCertificateAuthorityBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this configure certificate authority bad request response has a 5xx status code
func (o *ConfigureCertificateAuthorityBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this configure certificate authority bad request response a status code equal to that given
func (o *ConfigureCertificateAuthorityBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *ConfigureCertificateAuthorityBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /v1/certificate-authorities][%d] configureCertificateAuthorityBadRequest  %+v", 400, o.Payload)
}

func (o *ConfigureCertificateAuthorityBadRequest) String() string {
	return fmt.Sprintf("[PATCH /v1/certificate-authorities][%d] configureCertificateAuthorityBadRequest  %+v", 400, o.Payload)
}

func (o *ConfigureCertificateAuthorityBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *ConfigureCertificateAuthorityBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewConfigureCertificateAuthorityNotFound creates a ConfigureCertificateAuthorityNotFound with default headers values
func NewConfigureCertificateAuthorityNotFound() *ConfigureCertificateAuthorityNotFound {
	return &ConfigureCertificateAuthorityNotFound{}
}

/*
ConfigureCertificateAuthorityNotFound describes a response with status code 404, with default header values.

Not Found
*/
type ConfigureCertificateAuthorityNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this configure certificate authority not found response has a 2xx status code
func (o *ConfigureCertificateAuthorityNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this configure certificate authority not found response has a 3xx status code
func (o *ConfigureCertificateAuthorityNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this configure certificate authority not found response has a 4xx status code
func (o *ConfigureCertificateAuthorityNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this configure certificate authority not found response has a 5xx status code
func (o *ConfigureCertificateAuthorityNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this configure certificate authority not found response a status code equal to that given
func (o *ConfigureCertificateAuthorityNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *ConfigureCertificateAuthorityNotFound) Error() string {
	return fmt.Sprintf("[PATCH /v1/certificate-authorities][%d] configureCertificateAuthorityNotFound  %+v", 404, o.Payload)
}

func (o *ConfigureCertificateAuthorityNotFound) String() string {
	return fmt.Sprintf("[PATCH /v1/certificate-authorities][%d] configureCertificateAuthorityNotFound  %+v", 404, o.Payload)
}

func (o *ConfigureCertificateAuthorityNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *ConfigureCertificateAuthorityNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewConfigureCertificateAuthorityInternalServerError creates a ConfigureCertificateAuthorityInternalServerError with default headers values
func NewConfigureCertificateAuthorityInternalServerError() *ConfigureCertificateAuthorityInternalServerError {
	return &ConfigureCertificateAuthorityInternalServerError{}
}

/*
ConfigureCertificateAuthorityInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type ConfigureCertificateAuthorityInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this configure certificate authority internal server error response has a 2xx status code
func (o *ConfigureCertificateAuthorityInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this configure certificate authority internal server error response has a 3xx status code
func (o *ConfigureCertificateAuthorityInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this configure certificate authority internal server error response has a 4xx status code
func (o *ConfigureCertificateAuthorityInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this configure certificate authority internal server error response has a 5xx status code
func (o *ConfigureCertificateAuthorityInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this configure certificate authority internal server error response a status code equal to that given
func (o *ConfigureCertificateAuthorityInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *ConfigureCertificateAuthorityInternalServerError) Error() string {
	return fmt.Sprintf("[PATCH /v1/certificate-authorities][%d] configureCertificateAuthorityInternalServerError  %+v", 500, o.Payload)
}

func (o *ConfigureCertificateAuthorityInternalServerError) String() string {
	return fmt.Sprintf("[PATCH /v1/certificate-authorities][%d] configureCertificateAuthorityInternalServerError  %+v", 500, o.Payload)
}

func (o *ConfigureCertificateAuthorityInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *ConfigureCertificateAuthorityInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
