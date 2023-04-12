// Code generated by go-swagger; DO NOT EDIT.

package identity_providers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"vcf-sdk-go/models"
)

// GetAllIdpsReader is a Reader for the GetAllIdps structure.
type GetAllIdpsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAllIdpsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAllIdpsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetAllIdpsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetAllIdpsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetAllIdpsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetAllIdpsOK creates a GetAllIdpsOK with default headers values
func NewGetAllIdpsOK() *GetAllIdpsOK {
	return &GetAllIdpsOK{}
}

/*
GetAllIdpsOK describes a response with status code 200, with default header values.

OK
*/
type GetAllIdpsOK struct {
	Payload *models.PageOfIdentityProvider
}

// IsSuccess returns true when this get all idps o k response has a 2xx status code
func (o *GetAllIdpsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get all idps o k response has a 3xx status code
func (o *GetAllIdpsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get all idps o k response has a 4xx status code
func (o *GetAllIdpsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get all idps o k response has a 5xx status code
func (o *GetAllIdpsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get all idps o k response a status code equal to that given
func (o *GetAllIdpsOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetAllIdpsOK) Error() string {
	return fmt.Sprintf("[GET /v1/identity-providers][%d] getAllIdpsOK  %+v", 200, o.Payload)
}

func (o *GetAllIdpsOK) String() string {
	return fmt.Sprintf("[GET /v1/identity-providers][%d] getAllIdpsOK  %+v", 200, o.Payload)
}

func (o *GetAllIdpsOK) GetPayload() *models.PageOfIdentityProvider {
	return o.Payload
}

func (o *GetAllIdpsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PageOfIdentityProvider)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAllIdpsBadRequest creates a GetAllIdpsBadRequest with default headers values
func NewGetAllIdpsBadRequest() *GetAllIdpsBadRequest {
	return &GetAllIdpsBadRequest{}
}

/*
GetAllIdpsBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type GetAllIdpsBadRequest struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this get all idps bad request response has a 2xx status code
func (o *GetAllIdpsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get all idps bad request response has a 3xx status code
func (o *GetAllIdpsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get all idps bad request response has a 4xx status code
func (o *GetAllIdpsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get all idps bad request response has a 5xx status code
func (o *GetAllIdpsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get all idps bad request response a status code equal to that given
func (o *GetAllIdpsBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetAllIdpsBadRequest) Error() string {
	return fmt.Sprintf("[GET /v1/identity-providers][%d] getAllIdpsBadRequest  %+v", 400, o.Payload)
}

func (o *GetAllIdpsBadRequest) String() string {
	return fmt.Sprintf("[GET /v1/identity-providers][%d] getAllIdpsBadRequest  %+v", 400, o.Payload)
}

func (o *GetAllIdpsBadRequest) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetAllIdpsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAllIdpsUnauthorized creates a GetAllIdpsUnauthorized with default headers values
func NewGetAllIdpsUnauthorized() *GetAllIdpsUnauthorized {
	return &GetAllIdpsUnauthorized{}
}

/*
GetAllIdpsUnauthorized describes a response with status code 401, with default header values.

Unauthorized Request
*/
type GetAllIdpsUnauthorized struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this get all idps unauthorized response has a 2xx status code
func (o *GetAllIdpsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get all idps unauthorized response has a 3xx status code
func (o *GetAllIdpsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get all idps unauthorized response has a 4xx status code
func (o *GetAllIdpsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get all idps unauthorized response has a 5xx status code
func (o *GetAllIdpsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get all idps unauthorized response a status code equal to that given
func (o *GetAllIdpsUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetAllIdpsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v1/identity-providers][%d] getAllIdpsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetAllIdpsUnauthorized) String() string {
	return fmt.Sprintf("[GET /v1/identity-providers][%d] getAllIdpsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetAllIdpsUnauthorized) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetAllIdpsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAllIdpsInternalServerError creates a GetAllIdpsInternalServerError with default headers values
func NewGetAllIdpsInternalServerError() *GetAllIdpsInternalServerError {
	return &GetAllIdpsInternalServerError{}
}

/*
GetAllIdpsInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetAllIdpsInternalServerError struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this get all idps internal server error response has a 2xx status code
func (o *GetAllIdpsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get all idps internal server error response has a 3xx status code
func (o *GetAllIdpsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get all idps internal server error response has a 4xx status code
func (o *GetAllIdpsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get all idps internal server error response has a 5xx status code
func (o *GetAllIdpsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get all idps internal server error response a status code equal to that given
func (o *GetAllIdpsInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetAllIdpsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/identity-providers][%d] getAllIdpsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetAllIdpsInternalServerError) String() string {
	return fmt.Sprintf("[GET /v1/identity-providers][%d] getAllIdpsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetAllIdpsInternalServerError) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetAllIdpsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
