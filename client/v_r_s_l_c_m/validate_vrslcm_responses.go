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

// ValidateVrslcmReader is a Reader for the ValidateVrslcm structure.
type ValidateVrslcmReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ValidateVrslcmReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewValidateVrslcmAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewValidateVrslcmBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewValidateVrslcmAccepted creates a ValidateVrslcmAccepted with default headers values
func NewValidateVrslcmAccepted() *ValidateVrslcmAccepted {
	return &ValidateVrslcmAccepted{}
}

/*
ValidateVrslcmAccepted describes a response with status code 202, with default header values.

Accepted
*/
type ValidateVrslcmAccepted struct {
	Payload *models.Validation
}

// IsSuccess returns true when this validate vrslcm accepted response has a 2xx status code
func (o *ValidateVrslcmAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this validate vrslcm accepted response has a 3xx status code
func (o *ValidateVrslcmAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this validate vrslcm accepted response has a 4xx status code
func (o *ValidateVrslcmAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this validate vrslcm accepted response has a 5xx status code
func (o *ValidateVrslcmAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this validate vrslcm accepted response a status code equal to that given
func (o *ValidateVrslcmAccepted) IsCode(code int) bool {
	return code == 202
}

func (o *ValidateVrslcmAccepted) Error() string {
	return fmt.Sprintf("[POST /v1/vrslcms/validations][%d] validateVrslcmAccepted  %+v", 202, o.Payload)
}

func (o *ValidateVrslcmAccepted) String() string {
	return fmt.Sprintf("[POST /v1/vrslcms/validations][%d] validateVrslcmAccepted  %+v", 202, o.Payload)
}

func (o *ValidateVrslcmAccepted) GetPayload() *models.Validation {
	return o.Payload
}

func (o *ValidateVrslcmAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Validation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewValidateVrslcmBadRequest creates a ValidateVrslcmBadRequest with default headers values
func NewValidateVrslcmBadRequest() *ValidateVrslcmBadRequest {
	return &ValidateVrslcmBadRequest{}
}

/*
ValidateVrslcmBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type ValidateVrslcmBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this validate vrslcm bad request response has a 2xx status code
func (o *ValidateVrslcmBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this validate vrslcm bad request response has a 3xx status code
func (o *ValidateVrslcmBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this validate vrslcm bad request response has a 4xx status code
func (o *ValidateVrslcmBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this validate vrslcm bad request response has a 5xx status code
func (o *ValidateVrslcmBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this validate vrslcm bad request response a status code equal to that given
func (o *ValidateVrslcmBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *ValidateVrslcmBadRequest) Error() string {
	return fmt.Sprintf("[POST /v1/vrslcms/validations][%d] validateVrslcmBadRequest  %+v", 400, o.Payload)
}

func (o *ValidateVrslcmBadRequest) String() string {
	return fmt.Sprintf("[POST /v1/vrslcms/validations][%d] validateVrslcmBadRequest  %+v", 400, o.Payload)
}

func (o *ValidateVrslcmBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *ValidateVrslcmBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}