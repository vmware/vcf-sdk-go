// Code generated by go-swagger; DO NOT EDIT.

package v_r_s_l_c_m

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"vcf-sdk-go/models"
)

// GetVrslcmValidationReader is a Reader for the GetVrslcmValidation structure.
type GetVrslcmValidationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetVrslcmValidationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetVrslcmValidationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetVrslcmValidationBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetVrslcmValidationNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetVrslcmValidationOK creates a GetVrslcmValidationOK with default headers values
func NewGetVrslcmValidationOK() *GetVrslcmValidationOK {
	return &GetVrslcmValidationOK{}
}

/*
GetVrslcmValidationOK describes a response with status code 200, with default header values.

OK
*/
type GetVrslcmValidationOK struct {
	Payload *models.Validation
}

// IsSuccess returns true when this get vrslcm validation o k response has a 2xx status code
func (o *GetVrslcmValidationOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get vrslcm validation o k response has a 3xx status code
func (o *GetVrslcmValidationOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get vrslcm validation o k response has a 4xx status code
func (o *GetVrslcmValidationOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get vrslcm validation o k response has a 5xx status code
func (o *GetVrslcmValidationOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get vrslcm validation o k response a status code equal to that given
func (o *GetVrslcmValidationOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetVrslcmValidationOK) Error() string {
	return fmt.Sprintf("[GET /v1/vrslcms/validations/{id}][%d] getVrslcmValidationOK  %+v", 200, o.Payload)
}

func (o *GetVrslcmValidationOK) String() string {
	return fmt.Sprintf("[GET /v1/vrslcms/validations/{id}][%d] getVrslcmValidationOK  %+v", 200, o.Payload)
}

func (o *GetVrslcmValidationOK) GetPayload() *models.Validation {
	return o.Payload
}

func (o *GetVrslcmValidationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Validation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetVrslcmValidationBadRequest creates a GetVrslcmValidationBadRequest with default headers values
func NewGetVrslcmValidationBadRequest() *GetVrslcmValidationBadRequest {
	return &GetVrslcmValidationBadRequest{}
}

/*
GetVrslcmValidationBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetVrslcmValidationBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this get vrslcm validation bad request response has a 2xx status code
func (o *GetVrslcmValidationBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get vrslcm validation bad request response has a 3xx status code
func (o *GetVrslcmValidationBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get vrslcm validation bad request response has a 4xx status code
func (o *GetVrslcmValidationBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get vrslcm validation bad request response has a 5xx status code
func (o *GetVrslcmValidationBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get vrslcm validation bad request response a status code equal to that given
func (o *GetVrslcmValidationBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetVrslcmValidationBadRequest) Error() string {
	return fmt.Sprintf("[GET /v1/vrslcms/validations/{id}][%d] getVrslcmValidationBadRequest  %+v", 400, o.Payload)
}

func (o *GetVrslcmValidationBadRequest) String() string {
	return fmt.Sprintf("[GET /v1/vrslcms/validations/{id}][%d] getVrslcmValidationBadRequest  %+v", 400, o.Payload)
}

func (o *GetVrslcmValidationBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetVrslcmValidationBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetVrslcmValidationNotFound creates a GetVrslcmValidationNotFound with default headers values
func NewGetVrslcmValidationNotFound() *GetVrslcmValidationNotFound {
	return &GetVrslcmValidationNotFound{}
}

/*
GetVrslcmValidationNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetVrslcmValidationNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this get vrslcm validation not found response has a 2xx status code
func (o *GetVrslcmValidationNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get vrslcm validation not found response has a 3xx status code
func (o *GetVrslcmValidationNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get vrslcm validation not found response has a 4xx status code
func (o *GetVrslcmValidationNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get vrslcm validation not found response has a 5xx status code
func (o *GetVrslcmValidationNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get vrslcm validation not found response a status code equal to that given
func (o *GetVrslcmValidationNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetVrslcmValidationNotFound) Error() string {
	return fmt.Sprintf("[GET /v1/vrslcms/validations/{id}][%d] getVrslcmValidationNotFound  %+v", 404, o.Payload)
}

func (o *GetVrslcmValidationNotFound) String() string {
	return fmt.Sprintf("[GET /v1/vrslcms/validations/{id}][%d] getVrslcmValidationNotFound  %+v", 404, o.Payload)
}

func (o *GetVrslcmValidationNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetVrslcmValidationNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}