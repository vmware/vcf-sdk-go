// Code generated by go-swagger; DO NOT EDIT.

package upgrades

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"vcf-sdk-go/models"
)

// GetPrecheckUsingGETReader is a Reader for the GetPrecheckUsingGET structure.
type GetPrecheckUsingGETReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPrecheckUsingGETReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetPrecheckUsingGETOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetPrecheckUsingGETBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetPrecheckUsingGETForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetPrecheckUsingGETInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetPrecheckUsingGETOK creates a GetPrecheckUsingGETOK with default headers values
func NewGetPrecheckUsingGETOK() *GetPrecheckUsingGETOK {
	return &GetPrecheckUsingGETOK{}
}

/*
GetPrecheckUsingGETOK describes a response with status code 200, with default header values.

Ok
*/
type GetPrecheckUsingGETOK struct {
	Payload *models.Task
}

// IsSuccess returns true when this get precheck using g e t o k response has a 2xx status code
func (o *GetPrecheckUsingGETOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get precheck using g e t o k response has a 3xx status code
func (o *GetPrecheckUsingGETOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get precheck using g e t o k response has a 4xx status code
func (o *GetPrecheckUsingGETOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get precheck using g e t o k response has a 5xx status code
func (o *GetPrecheckUsingGETOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get precheck using g e t o k response a status code equal to that given
func (o *GetPrecheckUsingGETOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetPrecheckUsingGETOK) Error() string {
	return fmt.Sprintf("[GET /v1/upgrades/{upgradeId}/prechecks/{precheckId}][%d] getPrecheckUsingGETOK  %+v", 200, o.Payload)
}

func (o *GetPrecheckUsingGETOK) String() string {
	return fmt.Sprintf("[GET /v1/upgrades/{upgradeId}/prechecks/{precheckId}][%d] getPrecheckUsingGETOK  %+v", 200, o.Payload)
}

func (o *GetPrecheckUsingGETOK) GetPayload() *models.Task {
	return o.Payload
}

func (o *GetPrecheckUsingGETOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Task)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPrecheckUsingGETBadRequest creates a GetPrecheckUsingGETBadRequest with default headers values
func NewGetPrecheckUsingGETBadRequest() *GetPrecheckUsingGETBadRequest {
	return &GetPrecheckUsingGETBadRequest{}
}

/*
GetPrecheckUsingGETBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetPrecheckUsingGETBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this get precheck using g e t bad request response has a 2xx status code
func (o *GetPrecheckUsingGETBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get precheck using g e t bad request response has a 3xx status code
func (o *GetPrecheckUsingGETBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get precheck using g e t bad request response has a 4xx status code
func (o *GetPrecheckUsingGETBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get precheck using g e t bad request response has a 5xx status code
func (o *GetPrecheckUsingGETBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get precheck using g e t bad request response a status code equal to that given
func (o *GetPrecheckUsingGETBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetPrecheckUsingGETBadRequest) Error() string {
	return fmt.Sprintf("[GET /v1/upgrades/{upgradeId}/prechecks/{precheckId}][%d] getPrecheckUsingGETBadRequest  %+v", 400, o.Payload)
}

func (o *GetPrecheckUsingGETBadRequest) String() string {
	return fmt.Sprintf("[GET /v1/upgrades/{upgradeId}/prechecks/{precheckId}][%d] getPrecheckUsingGETBadRequest  %+v", 400, o.Payload)
}

func (o *GetPrecheckUsingGETBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetPrecheckUsingGETBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPrecheckUsingGETForbidden creates a GetPrecheckUsingGETForbidden with default headers values
func NewGetPrecheckUsingGETForbidden() *GetPrecheckUsingGETForbidden {
	return &GetPrecheckUsingGETForbidden{}
}

/*
GetPrecheckUsingGETForbidden describes a response with status code 403, with default header values.

Operation not allowed
*/
type GetPrecheckUsingGETForbidden struct {
	Payload *models.Error
}

// IsSuccess returns true when this get precheck using g e t forbidden response has a 2xx status code
func (o *GetPrecheckUsingGETForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get precheck using g e t forbidden response has a 3xx status code
func (o *GetPrecheckUsingGETForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get precheck using g e t forbidden response has a 4xx status code
func (o *GetPrecheckUsingGETForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get precheck using g e t forbidden response has a 5xx status code
func (o *GetPrecheckUsingGETForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get precheck using g e t forbidden response a status code equal to that given
func (o *GetPrecheckUsingGETForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetPrecheckUsingGETForbidden) Error() string {
	return fmt.Sprintf("[GET /v1/upgrades/{upgradeId}/prechecks/{precheckId}][%d] getPrecheckUsingGETForbidden  %+v", 403, o.Payload)
}

func (o *GetPrecheckUsingGETForbidden) String() string {
	return fmt.Sprintf("[GET /v1/upgrades/{upgradeId}/prechecks/{precheckId}][%d] getPrecheckUsingGETForbidden  %+v", 403, o.Payload)
}

func (o *GetPrecheckUsingGETForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetPrecheckUsingGETForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPrecheckUsingGETInternalServerError creates a GetPrecheckUsingGETInternalServerError with default headers values
func NewGetPrecheckUsingGETInternalServerError() *GetPrecheckUsingGETInternalServerError {
	return &GetPrecheckUsingGETInternalServerError{}
}

/*
GetPrecheckUsingGETInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetPrecheckUsingGETInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this get precheck using g e t internal server error response has a 2xx status code
func (o *GetPrecheckUsingGETInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get precheck using g e t internal server error response has a 3xx status code
func (o *GetPrecheckUsingGETInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get precheck using g e t internal server error response has a 4xx status code
func (o *GetPrecheckUsingGETInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get precheck using g e t internal server error response has a 5xx status code
func (o *GetPrecheckUsingGETInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get precheck using g e t internal server error response a status code equal to that given
func (o *GetPrecheckUsingGETInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetPrecheckUsingGETInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/upgrades/{upgradeId}/prechecks/{precheckId}][%d] getPrecheckUsingGETInternalServerError  %+v", 500, o.Payload)
}

func (o *GetPrecheckUsingGETInternalServerError) String() string {
	return fmt.Sprintf("[GET /v1/upgrades/{upgradeId}/prechecks/{precheckId}][%d] getPrecheckUsingGETInternalServerError  %+v", 500, o.Payload)
}

func (o *GetPrecheckUsingGETInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetPrecheckUsingGETInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
