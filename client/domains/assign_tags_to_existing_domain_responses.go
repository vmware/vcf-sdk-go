// Code generated by go-swagger; DO NOT EDIT.

package domains

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"vcf-sdk-go/models"
)

// AssignTagsToExistingDomainReader is a Reader for the AssignTagsToExistingDomain structure.
type AssignTagsToExistingDomainReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AssignTagsToExistingDomainReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAssignTagsToExistingDomainOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewAssignTagsToExistingDomainBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewAssignTagsToExistingDomainInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewAssignTagsToExistingDomainOK creates a AssignTagsToExistingDomainOK with default headers values
func NewAssignTagsToExistingDomainOK() *AssignTagsToExistingDomainOK {
	return &AssignTagsToExistingDomainOK{}
}

/*
AssignTagsToExistingDomainOK describes a response with status code 200, with default header values.

Ok
*/
type AssignTagsToExistingDomainOK struct {
	Payload *models.TagAssignmentResult
}

// IsSuccess returns true when this assign tags to existing domain o k response has a 2xx status code
func (o *AssignTagsToExistingDomainOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this assign tags to existing domain o k response has a 3xx status code
func (o *AssignTagsToExistingDomainOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this assign tags to existing domain o k response has a 4xx status code
func (o *AssignTagsToExistingDomainOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this assign tags to existing domain o k response has a 5xx status code
func (o *AssignTagsToExistingDomainOK) IsServerError() bool {
	return false
}

// IsCode returns true when this assign tags to existing domain o k response a status code equal to that given
func (o *AssignTagsToExistingDomainOK) IsCode(code int) bool {
	return code == 200
}

func (o *AssignTagsToExistingDomainOK) Error() string {
	return fmt.Sprintf("[PUT /v1/domains/{id}/tags][%d] assignTagsToExistingDomainOK  %+v", 200, o.Payload)
}

func (o *AssignTagsToExistingDomainOK) String() string {
	return fmt.Sprintf("[PUT /v1/domains/{id}/tags][%d] assignTagsToExistingDomainOK  %+v", 200, o.Payload)
}

func (o *AssignTagsToExistingDomainOK) GetPayload() *models.TagAssignmentResult {
	return o.Payload
}

func (o *AssignTagsToExistingDomainOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TagAssignmentResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAssignTagsToExistingDomainBadRequest creates a AssignTagsToExistingDomainBadRequest with default headers values
func NewAssignTagsToExistingDomainBadRequest() *AssignTagsToExistingDomainBadRequest {
	return &AssignTagsToExistingDomainBadRequest{}
}

/*
AssignTagsToExistingDomainBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type AssignTagsToExistingDomainBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this assign tags to existing domain bad request response has a 2xx status code
func (o *AssignTagsToExistingDomainBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this assign tags to existing domain bad request response has a 3xx status code
func (o *AssignTagsToExistingDomainBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this assign tags to existing domain bad request response has a 4xx status code
func (o *AssignTagsToExistingDomainBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this assign tags to existing domain bad request response has a 5xx status code
func (o *AssignTagsToExistingDomainBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this assign tags to existing domain bad request response a status code equal to that given
func (o *AssignTagsToExistingDomainBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *AssignTagsToExistingDomainBadRequest) Error() string {
	return fmt.Sprintf("[PUT /v1/domains/{id}/tags][%d] assignTagsToExistingDomainBadRequest  %+v", 400, o.Payload)
}

func (o *AssignTagsToExistingDomainBadRequest) String() string {
	return fmt.Sprintf("[PUT /v1/domains/{id}/tags][%d] assignTagsToExistingDomainBadRequest  %+v", 400, o.Payload)
}

func (o *AssignTagsToExistingDomainBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *AssignTagsToExistingDomainBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAssignTagsToExistingDomainInternalServerError creates a AssignTagsToExistingDomainInternalServerError with default headers values
func NewAssignTagsToExistingDomainInternalServerError() *AssignTagsToExistingDomainInternalServerError {
	return &AssignTagsToExistingDomainInternalServerError{}
}

/*
AssignTagsToExistingDomainInternalServerError describes a response with status code 500, with default header values.

InternalServerError
*/
type AssignTagsToExistingDomainInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this assign tags to existing domain internal server error response has a 2xx status code
func (o *AssignTagsToExistingDomainInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this assign tags to existing domain internal server error response has a 3xx status code
func (o *AssignTagsToExistingDomainInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this assign tags to existing domain internal server error response has a 4xx status code
func (o *AssignTagsToExistingDomainInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this assign tags to existing domain internal server error response has a 5xx status code
func (o *AssignTagsToExistingDomainInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this assign tags to existing domain internal server error response a status code equal to that given
func (o *AssignTagsToExistingDomainInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *AssignTagsToExistingDomainInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /v1/domains/{id}/tags][%d] assignTagsToExistingDomainInternalServerError  %+v", 500, o.Payload)
}

func (o *AssignTagsToExistingDomainInternalServerError) String() string {
	return fmt.Sprintf("[PUT /v1/domains/{id}/tags][%d] assignTagsToExistingDomainInternalServerError  %+v", 500, o.Payload)
}

func (o *AssignTagsToExistingDomainInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *AssignTagsToExistingDomainInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
