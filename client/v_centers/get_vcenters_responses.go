// Code generated by go-swagger; DO NOT EDIT.

package v_centers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"vcf-sdk-go/models"
)

// GetVcentersReader is a Reader for the GetVcenters structure.
type GetVcentersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetVcentersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetVcentersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetVcentersBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetVcentersInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetVcentersOK creates a GetVcentersOK with default headers values
func NewGetVcentersOK() *GetVcentersOK {
	return &GetVcentersOK{}
}

/*
GetVcentersOK describes a response with status code 200, with default header values.

Ok
*/
type GetVcentersOK struct {
	Payload *models.PageOfVcenter
}

// IsSuccess returns true when this get vcenters o k response has a 2xx status code
func (o *GetVcentersOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get vcenters o k response has a 3xx status code
func (o *GetVcentersOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get vcenters o k response has a 4xx status code
func (o *GetVcentersOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get vcenters o k response has a 5xx status code
func (o *GetVcentersOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get vcenters o k response a status code equal to that given
func (o *GetVcentersOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetVcentersOK) Error() string {
	return fmt.Sprintf("[GET /v1/vcenters][%d] getVcentersOK  %+v", 200, o.Payload)
}

func (o *GetVcentersOK) String() string {
	return fmt.Sprintf("[GET /v1/vcenters][%d] getVcentersOK  %+v", 200, o.Payload)
}

func (o *GetVcentersOK) GetPayload() *models.PageOfVcenter {
	return o.Payload
}

func (o *GetVcentersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PageOfVcenter)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetVcentersBadRequest creates a GetVcentersBadRequest with default headers values
func NewGetVcentersBadRequest() *GetVcentersBadRequest {
	return &GetVcentersBadRequest{}
}

/*
GetVcentersBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetVcentersBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this get vcenters bad request response has a 2xx status code
func (o *GetVcentersBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get vcenters bad request response has a 3xx status code
func (o *GetVcentersBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get vcenters bad request response has a 4xx status code
func (o *GetVcentersBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get vcenters bad request response has a 5xx status code
func (o *GetVcentersBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get vcenters bad request response a status code equal to that given
func (o *GetVcentersBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetVcentersBadRequest) Error() string {
	return fmt.Sprintf("[GET /v1/vcenters][%d] getVcentersBadRequest  %+v", 400, o.Payload)
}

func (o *GetVcentersBadRequest) String() string {
	return fmt.Sprintf("[GET /v1/vcenters][%d] getVcentersBadRequest  %+v", 400, o.Payload)
}

func (o *GetVcentersBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetVcentersBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetVcentersInternalServerError creates a GetVcentersInternalServerError with default headers values
func NewGetVcentersInternalServerError() *GetVcentersInternalServerError {
	return &GetVcentersInternalServerError{}
}

/*
GetVcentersInternalServerError describes a response with status code 500, with default header values.

InternalServerError
*/
type GetVcentersInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this get vcenters internal server error response has a 2xx status code
func (o *GetVcentersInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get vcenters internal server error response has a 3xx status code
func (o *GetVcentersInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get vcenters internal server error response has a 4xx status code
func (o *GetVcentersInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get vcenters internal server error response has a 5xx status code
func (o *GetVcentersInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get vcenters internal server error response a status code equal to that given
func (o *GetVcentersInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetVcentersInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/vcenters][%d] getVcentersInternalServerError  %+v", 500, o.Payload)
}

func (o *GetVcentersInternalServerError) String() string {
	return fmt.Sprintf("[GET /v1/vcenters][%d] getVcentersInternalServerError  %+v", 500, o.Payload)
}

func (o *GetVcentersInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetVcentersInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
