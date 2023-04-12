// Code generated by go-swagger; DO NOT EDIT.

package hosts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"vcf-sdk-go/models"
)

// DecommissionHostsReader is a Reader for the DecommissionHosts structure.
type DecommissionHostsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DecommissionHostsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDecommissionHostsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewDecommissionHostsAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDecommissionHostsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDecommissionHostsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDecommissionHostsOK creates a DecommissionHostsOK with default headers values
func NewDecommissionHostsOK() *DecommissionHostsOK {
	return &DecommissionHostsOK{}
}

/*
DecommissionHostsOK describes a response with status code 200, with default header values.

OK
*/
type DecommissionHostsOK struct {
	Payload *models.Task
}

// IsSuccess returns true when this decommission hosts o k response has a 2xx status code
func (o *DecommissionHostsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this decommission hosts o k response has a 3xx status code
func (o *DecommissionHostsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this decommission hosts o k response has a 4xx status code
func (o *DecommissionHostsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this decommission hosts o k response has a 5xx status code
func (o *DecommissionHostsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this decommission hosts o k response a status code equal to that given
func (o *DecommissionHostsOK) IsCode(code int) bool {
	return code == 200
}

func (o *DecommissionHostsOK) Error() string {
	return fmt.Sprintf("[DELETE /v1/hosts][%d] decommissionHostsOK  %+v", 200, o.Payload)
}

func (o *DecommissionHostsOK) String() string {
	return fmt.Sprintf("[DELETE /v1/hosts][%d] decommissionHostsOK  %+v", 200, o.Payload)
}

func (o *DecommissionHostsOK) GetPayload() *models.Task {
	return o.Payload
}

func (o *DecommissionHostsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Task)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDecommissionHostsAccepted creates a DecommissionHostsAccepted with default headers values
func NewDecommissionHostsAccepted() *DecommissionHostsAccepted {
	return &DecommissionHostsAccepted{}
}

/*
DecommissionHostsAccepted describes a response with status code 202, with default header values.

Accepted
*/
type DecommissionHostsAccepted struct {
	Payload *models.Task
}

// IsSuccess returns true when this decommission hosts accepted response has a 2xx status code
func (o *DecommissionHostsAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this decommission hosts accepted response has a 3xx status code
func (o *DecommissionHostsAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this decommission hosts accepted response has a 4xx status code
func (o *DecommissionHostsAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this decommission hosts accepted response has a 5xx status code
func (o *DecommissionHostsAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this decommission hosts accepted response a status code equal to that given
func (o *DecommissionHostsAccepted) IsCode(code int) bool {
	return code == 202
}

func (o *DecommissionHostsAccepted) Error() string {
	return fmt.Sprintf("[DELETE /v1/hosts][%d] decommissionHostsAccepted  %+v", 202, o.Payload)
}

func (o *DecommissionHostsAccepted) String() string {
	return fmt.Sprintf("[DELETE /v1/hosts][%d] decommissionHostsAccepted  %+v", 202, o.Payload)
}

func (o *DecommissionHostsAccepted) GetPayload() *models.Task {
	return o.Payload
}

func (o *DecommissionHostsAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Task)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDecommissionHostsBadRequest creates a DecommissionHostsBadRequest with default headers values
func NewDecommissionHostsBadRequest() *DecommissionHostsBadRequest {
	return &DecommissionHostsBadRequest{}
}

/*
DecommissionHostsBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type DecommissionHostsBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this decommission hosts bad request response has a 2xx status code
func (o *DecommissionHostsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this decommission hosts bad request response has a 3xx status code
func (o *DecommissionHostsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this decommission hosts bad request response has a 4xx status code
func (o *DecommissionHostsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this decommission hosts bad request response has a 5xx status code
func (o *DecommissionHostsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this decommission hosts bad request response a status code equal to that given
func (o *DecommissionHostsBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *DecommissionHostsBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /v1/hosts][%d] decommissionHostsBadRequest  %+v", 400, o.Payload)
}

func (o *DecommissionHostsBadRequest) String() string {
	return fmt.Sprintf("[DELETE /v1/hosts][%d] decommissionHostsBadRequest  %+v", 400, o.Payload)
}

func (o *DecommissionHostsBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *DecommissionHostsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDecommissionHostsInternalServerError creates a DecommissionHostsInternalServerError with default headers values
func NewDecommissionHostsInternalServerError() *DecommissionHostsInternalServerError {
	return &DecommissionHostsInternalServerError{}
}

/*
DecommissionHostsInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type DecommissionHostsInternalServerError struct {
}

// IsSuccess returns true when this decommission hosts internal server error response has a 2xx status code
func (o *DecommissionHostsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this decommission hosts internal server error response has a 3xx status code
func (o *DecommissionHostsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this decommission hosts internal server error response has a 4xx status code
func (o *DecommissionHostsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this decommission hosts internal server error response has a 5xx status code
func (o *DecommissionHostsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this decommission hosts internal server error response a status code equal to that given
func (o *DecommissionHostsInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *DecommissionHostsInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /v1/hosts][%d] decommissionHostsInternalServerError ", 500)
}

func (o *DecommissionHostsInternalServerError) String() string {
	return fmt.Sprintf("[DELETE /v1/hosts][%d] decommissionHostsInternalServerError ", 500)
}

func (o *DecommissionHostsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}