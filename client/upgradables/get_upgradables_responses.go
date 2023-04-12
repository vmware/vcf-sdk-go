// Code generated by go-swagger; DO NOT EDIT.

package upgradables

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"vcf-sdk-go/models"
)

// GetUpgradablesReader is a Reader for the GetUpgradables structure.
type GetUpgradablesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetUpgradablesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetUpgradablesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewGetUpgradablesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetUpgradablesOK creates a GetUpgradablesOK with default headers values
func NewGetUpgradablesOK() *GetUpgradablesOK {
	return &GetUpgradablesOK{}
}

/*
GetUpgradablesOK describes a response with status code 200, with default header values.

Ok
*/
type GetUpgradablesOK struct {
	Payload *models.PageOfUpgradable
}

// IsSuccess returns true when this get upgradables o k response has a 2xx status code
func (o *GetUpgradablesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get upgradables o k response has a 3xx status code
func (o *GetUpgradablesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get upgradables o k response has a 4xx status code
func (o *GetUpgradablesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get upgradables o k response has a 5xx status code
func (o *GetUpgradablesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get upgradables o k response a status code equal to that given
func (o *GetUpgradablesOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetUpgradablesOK) Error() string {
	return fmt.Sprintf("[GET /v1/system/upgradables][%d] getUpgradablesOK  %+v", 200, o.Payload)
}

func (o *GetUpgradablesOK) String() string {
	return fmt.Sprintf("[GET /v1/system/upgradables][%d] getUpgradablesOK  %+v", 200, o.Payload)
}

func (o *GetUpgradablesOK) GetPayload() *models.PageOfUpgradable {
	return o.Payload
}

func (o *GetUpgradablesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PageOfUpgradable)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUpgradablesInternalServerError creates a GetUpgradablesInternalServerError with default headers values
func NewGetUpgradablesInternalServerError() *GetUpgradablesInternalServerError {
	return &GetUpgradablesInternalServerError{}
}

/*
GetUpgradablesInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetUpgradablesInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this get upgradables internal server error response has a 2xx status code
func (o *GetUpgradablesInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get upgradables internal server error response has a 3xx status code
func (o *GetUpgradablesInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get upgradables internal server error response has a 4xx status code
func (o *GetUpgradablesInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get upgradables internal server error response has a 5xx status code
func (o *GetUpgradablesInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get upgradables internal server error response a status code equal to that given
func (o *GetUpgradablesInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetUpgradablesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/system/upgradables][%d] getUpgradablesInternalServerError  %+v", 500, o.Payload)
}

func (o *GetUpgradablesInternalServerError) String() string {
	return fmt.Sprintf("[GET /v1/system/upgradables][%d] getUpgradablesInternalServerError  %+v", 500, o.Payload)
}

func (o *GetUpgradablesInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetUpgradablesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
