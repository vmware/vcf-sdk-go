// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"vcf-sdk-go/models"
)

// GetSSODomainsReader is a Reader for the GetSSODomains structure.
type GetSSODomainsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSSODomainsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetSSODomainsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetSSODomainsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetSSODomainsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetSSODomainsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetSSODomainsOK creates a GetSSODomainsOK with default headers values
func NewGetSSODomainsOK() *GetSSODomainsOK {
	return &GetSSODomainsOK{}
}

/*
GetSSODomainsOK describes a response with status code 200, with default header values.

OK
*/
type GetSSODomainsOK struct {
	Payload *models.PageOfstring
}

// IsSuccess returns true when this get s s o domains o k response has a 2xx status code
func (o *GetSSODomainsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get s s o domains o k response has a 3xx status code
func (o *GetSSODomainsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get s s o domains o k response has a 4xx status code
func (o *GetSSODomainsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get s s o domains o k response has a 5xx status code
func (o *GetSSODomainsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get s s o domains o k response a status code equal to that given
func (o *GetSSODomainsOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetSSODomainsOK) Error() string {
	return fmt.Sprintf("[GET /v1/sso-domains][%d] getSSODomainsOK  %+v", 200, o.Payload)
}

func (o *GetSSODomainsOK) String() string {
	return fmt.Sprintf("[GET /v1/sso-domains][%d] getSSODomainsOK  %+v", 200, o.Payload)
}

func (o *GetSSODomainsOK) GetPayload() *models.PageOfstring {
	return o.Payload
}

func (o *GetSSODomainsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PageOfstring)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSSODomainsUnauthorized creates a GetSSODomainsUnauthorized with default headers values
func NewGetSSODomainsUnauthorized() *GetSSODomainsUnauthorized {
	return &GetSSODomainsUnauthorized{}
}

/*
GetSSODomainsUnauthorized describes a response with status code 401, with default header values.

Unauthorized Request
*/
type GetSSODomainsUnauthorized struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this get s s o domains unauthorized response has a 2xx status code
func (o *GetSSODomainsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get s s o domains unauthorized response has a 3xx status code
func (o *GetSSODomainsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get s s o domains unauthorized response has a 4xx status code
func (o *GetSSODomainsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get s s o domains unauthorized response has a 5xx status code
func (o *GetSSODomainsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get s s o domains unauthorized response a status code equal to that given
func (o *GetSSODomainsUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetSSODomainsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v1/sso-domains][%d] getSSODomainsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetSSODomainsUnauthorized) String() string {
	return fmt.Sprintf("[GET /v1/sso-domains][%d] getSSODomainsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetSSODomainsUnauthorized) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetSSODomainsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSSODomainsForbidden creates a GetSSODomainsForbidden with default headers values
func NewGetSSODomainsForbidden() *GetSSODomainsForbidden {
	return &GetSSODomainsForbidden{}
}

/*
GetSSODomainsForbidden describes a response with status code 403, with default header values.

Forbidden request
*/
type GetSSODomainsForbidden struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this get s s o domains forbidden response has a 2xx status code
func (o *GetSSODomainsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get s s o domains forbidden response has a 3xx status code
func (o *GetSSODomainsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get s s o domains forbidden response has a 4xx status code
func (o *GetSSODomainsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get s s o domains forbidden response has a 5xx status code
func (o *GetSSODomainsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get s s o domains forbidden response a status code equal to that given
func (o *GetSSODomainsForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetSSODomainsForbidden) Error() string {
	return fmt.Sprintf("[GET /v1/sso-domains][%d] getSSODomainsForbidden  %+v", 403, o.Payload)
}

func (o *GetSSODomainsForbidden) String() string {
	return fmt.Sprintf("[GET /v1/sso-domains][%d] getSSODomainsForbidden  %+v", 403, o.Payload)
}

func (o *GetSSODomainsForbidden) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetSSODomainsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSSODomainsInternalServerError creates a GetSSODomainsInternalServerError with default headers values
func NewGetSSODomainsInternalServerError() *GetSSODomainsInternalServerError {
	return &GetSSODomainsInternalServerError{}
}

/*
GetSSODomainsInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetSSODomainsInternalServerError struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this get s s o domains internal server error response has a 2xx status code
func (o *GetSSODomainsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get s s o domains internal server error response has a 3xx status code
func (o *GetSSODomainsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get s s o domains internal server error response has a 4xx status code
func (o *GetSSODomainsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get s s o domains internal server error response has a 5xx status code
func (o *GetSSODomainsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get s s o domains internal server error response a status code equal to that given
func (o *GetSSODomainsInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetSSODomainsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/sso-domains][%d] getSSODomainsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetSSODomainsInternalServerError) String() string {
	return fmt.Sprintf("[GET /v1/sso-domains][%d] getSSODomainsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetSSODomainsInternalServerError) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetSSODomainsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}