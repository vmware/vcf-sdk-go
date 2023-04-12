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

// GetValidationsOfDNSConfigurationReader is a Reader for the GetValidationsOfDNSConfiguration structure.
type GetValidationsOfDNSConfigurationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetValidationsOfDNSConfigurationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetValidationsOfDNSConfigurationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetValidationsOfDNSConfigurationBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetValidationsOfDNSConfigurationInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetValidationsOfDNSConfigurationOK creates a GetValidationsOfDNSConfigurationOK with default headers values
func NewGetValidationsOfDNSConfigurationOK() *GetValidationsOfDNSConfigurationOK {
	return &GetValidationsOfDNSConfigurationOK{}
}

/*
GetValidationsOfDNSConfigurationOK describes a response with status code 200, with default header values.

OK
*/
type GetValidationsOfDNSConfigurationOK struct {
	Payload *models.Validation
}

// IsSuccess returns true when this get validations of Dns configuration o k response has a 2xx status code
func (o *GetValidationsOfDNSConfigurationOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get validations of Dns configuration o k response has a 3xx status code
func (o *GetValidationsOfDNSConfigurationOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get validations of Dns configuration o k response has a 4xx status code
func (o *GetValidationsOfDNSConfigurationOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get validations of Dns configuration o k response has a 5xx status code
func (o *GetValidationsOfDNSConfigurationOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get validations of Dns configuration o k response a status code equal to that given
func (o *GetValidationsOfDNSConfigurationOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetValidationsOfDNSConfigurationOK) Error() string {
	return fmt.Sprintf("[GET /v1/system/dns-configuration/validations][%d] getValidationsOfDnsConfigurationOK  %+v", 200, o.Payload)
}

func (o *GetValidationsOfDNSConfigurationOK) String() string {
	return fmt.Sprintf("[GET /v1/system/dns-configuration/validations][%d] getValidationsOfDnsConfigurationOK  %+v", 200, o.Payload)
}

func (o *GetValidationsOfDNSConfigurationOK) GetPayload() *models.Validation {
	return o.Payload
}

func (o *GetValidationsOfDNSConfigurationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Validation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetValidationsOfDNSConfigurationBadRequest creates a GetValidationsOfDNSConfigurationBadRequest with default headers values
func NewGetValidationsOfDNSConfigurationBadRequest() *GetValidationsOfDNSConfigurationBadRequest {
	return &GetValidationsOfDNSConfigurationBadRequest{}
}

/*
GetValidationsOfDNSConfigurationBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetValidationsOfDNSConfigurationBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this get validations of Dns configuration bad request response has a 2xx status code
func (o *GetValidationsOfDNSConfigurationBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get validations of Dns configuration bad request response has a 3xx status code
func (o *GetValidationsOfDNSConfigurationBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get validations of Dns configuration bad request response has a 4xx status code
func (o *GetValidationsOfDNSConfigurationBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get validations of Dns configuration bad request response has a 5xx status code
func (o *GetValidationsOfDNSConfigurationBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get validations of Dns configuration bad request response a status code equal to that given
func (o *GetValidationsOfDNSConfigurationBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetValidationsOfDNSConfigurationBadRequest) Error() string {
	return fmt.Sprintf("[GET /v1/system/dns-configuration/validations][%d] getValidationsOfDnsConfigurationBadRequest  %+v", 400, o.Payload)
}

func (o *GetValidationsOfDNSConfigurationBadRequest) String() string {
	return fmt.Sprintf("[GET /v1/system/dns-configuration/validations][%d] getValidationsOfDnsConfigurationBadRequest  %+v", 400, o.Payload)
}

func (o *GetValidationsOfDNSConfigurationBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetValidationsOfDNSConfigurationBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetValidationsOfDNSConfigurationInternalServerError creates a GetValidationsOfDNSConfigurationInternalServerError with default headers values
func NewGetValidationsOfDNSConfigurationInternalServerError() *GetValidationsOfDNSConfigurationInternalServerError {
	return &GetValidationsOfDNSConfigurationInternalServerError{}
}

/*
GetValidationsOfDNSConfigurationInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type GetValidationsOfDNSConfigurationInternalServerError struct {
}

// IsSuccess returns true when this get validations of Dns configuration internal server error response has a 2xx status code
func (o *GetValidationsOfDNSConfigurationInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get validations of Dns configuration internal server error response has a 3xx status code
func (o *GetValidationsOfDNSConfigurationInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get validations of Dns configuration internal server error response has a 4xx status code
func (o *GetValidationsOfDNSConfigurationInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get validations of Dns configuration internal server error response has a 5xx status code
func (o *GetValidationsOfDNSConfigurationInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get validations of Dns configuration internal server error response a status code equal to that given
func (o *GetValidationsOfDNSConfigurationInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetValidationsOfDNSConfigurationInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/system/dns-configuration/validations][%d] getValidationsOfDnsConfigurationInternalServerError ", 500)
}

func (o *GetValidationsOfDNSConfigurationInternalServerError) String() string {
	return fmt.Sprintf("[GET /v1/system/dns-configuration/validations][%d] getValidationsOfDnsConfigurationInternalServerError ", 500)
}

func (o *GetValidationsOfDNSConfigurationInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
