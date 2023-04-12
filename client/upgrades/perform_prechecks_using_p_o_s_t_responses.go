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

// PerformPrechecksUsingPOSTReader is a Reader for the PerformPrechecksUsingPOST structure.
type PerformPrechecksUsingPOSTReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PerformPrechecksUsingPOSTReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPerformPrechecksUsingPOSTOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewPerformPrechecksUsingPOSTAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPerformPrechecksUsingPOSTBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPerformPrechecksUsingPOSTForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPerformPrechecksUsingPOSTInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPerformPrechecksUsingPOSTOK creates a PerformPrechecksUsingPOSTOK with default headers values
func NewPerformPrechecksUsingPOSTOK() *PerformPrechecksUsingPOSTOK {
	return &PerformPrechecksUsingPOSTOK{}
}

/*
PerformPrechecksUsingPOSTOK describes a response with status code 200, with default header values.

OK
*/
type PerformPrechecksUsingPOSTOK struct {
	Payload *models.Task
}

// IsSuccess returns true when this perform prechecks using p o s t o k response has a 2xx status code
func (o *PerformPrechecksUsingPOSTOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this perform prechecks using p o s t o k response has a 3xx status code
func (o *PerformPrechecksUsingPOSTOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this perform prechecks using p o s t o k response has a 4xx status code
func (o *PerformPrechecksUsingPOSTOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this perform prechecks using p o s t o k response has a 5xx status code
func (o *PerformPrechecksUsingPOSTOK) IsServerError() bool {
	return false
}

// IsCode returns true when this perform prechecks using p o s t o k response a status code equal to that given
func (o *PerformPrechecksUsingPOSTOK) IsCode(code int) bool {
	return code == 200
}

func (o *PerformPrechecksUsingPOSTOK) Error() string {
	return fmt.Sprintf("[POST /v1/upgrades/{upgradeId}/prechecks][%d] performPrechecksUsingPOSTOK  %+v", 200, o.Payload)
}

func (o *PerformPrechecksUsingPOSTOK) String() string {
	return fmt.Sprintf("[POST /v1/upgrades/{upgradeId}/prechecks][%d] performPrechecksUsingPOSTOK  %+v", 200, o.Payload)
}

func (o *PerformPrechecksUsingPOSTOK) GetPayload() *models.Task {
	return o.Payload
}

func (o *PerformPrechecksUsingPOSTOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Task)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPerformPrechecksUsingPOSTAccepted creates a PerformPrechecksUsingPOSTAccepted with default headers values
func NewPerformPrechecksUsingPOSTAccepted() *PerformPrechecksUsingPOSTAccepted {
	return &PerformPrechecksUsingPOSTAccepted{}
}

/*
PerformPrechecksUsingPOSTAccepted describes a response with status code 202, with default header values.

Accepted
*/
type PerformPrechecksUsingPOSTAccepted struct {
	Payload *models.Task
}

// IsSuccess returns true when this perform prechecks using p o s t accepted response has a 2xx status code
func (o *PerformPrechecksUsingPOSTAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this perform prechecks using p o s t accepted response has a 3xx status code
func (o *PerformPrechecksUsingPOSTAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this perform prechecks using p o s t accepted response has a 4xx status code
func (o *PerformPrechecksUsingPOSTAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this perform prechecks using p o s t accepted response has a 5xx status code
func (o *PerformPrechecksUsingPOSTAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this perform prechecks using p o s t accepted response a status code equal to that given
func (o *PerformPrechecksUsingPOSTAccepted) IsCode(code int) bool {
	return code == 202
}

func (o *PerformPrechecksUsingPOSTAccepted) Error() string {
	return fmt.Sprintf("[POST /v1/upgrades/{upgradeId}/prechecks][%d] performPrechecksUsingPOSTAccepted  %+v", 202, o.Payload)
}

func (o *PerformPrechecksUsingPOSTAccepted) String() string {
	return fmt.Sprintf("[POST /v1/upgrades/{upgradeId}/prechecks][%d] performPrechecksUsingPOSTAccepted  %+v", 202, o.Payload)
}

func (o *PerformPrechecksUsingPOSTAccepted) GetPayload() *models.Task {
	return o.Payload
}

func (o *PerformPrechecksUsingPOSTAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Task)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPerformPrechecksUsingPOSTBadRequest creates a PerformPrechecksUsingPOSTBadRequest with default headers values
func NewPerformPrechecksUsingPOSTBadRequest() *PerformPrechecksUsingPOSTBadRequest {
	return &PerformPrechecksUsingPOSTBadRequest{}
}

/*
PerformPrechecksUsingPOSTBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PerformPrechecksUsingPOSTBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this perform prechecks using p o s t bad request response has a 2xx status code
func (o *PerformPrechecksUsingPOSTBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this perform prechecks using p o s t bad request response has a 3xx status code
func (o *PerformPrechecksUsingPOSTBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this perform prechecks using p o s t bad request response has a 4xx status code
func (o *PerformPrechecksUsingPOSTBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this perform prechecks using p o s t bad request response has a 5xx status code
func (o *PerformPrechecksUsingPOSTBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this perform prechecks using p o s t bad request response a status code equal to that given
func (o *PerformPrechecksUsingPOSTBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *PerformPrechecksUsingPOSTBadRequest) Error() string {
	return fmt.Sprintf("[POST /v1/upgrades/{upgradeId}/prechecks][%d] performPrechecksUsingPOSTBadRequest  %+v", 400, o.Payload)
}

func (o *PerformPrechecksUsingPOSTBadRequest) String() string {
	return fmt.Sprintf("[POST /v1/upgrades/{upgradeId}/prechecks][%d] performPrechecksUsingPOSTBadRequest  %+v", 400, o.Payload)
}

func (o *PerformPrechecksUsingPOSTBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *PerformPrechecksUsingPOSTBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPerformPrechecksUsingPOSTForbidden creates a PerformPrechecksUsingPOSTForbidden with default headers values
func NewPerformPrechecksUsingPOSTForbidden() *PerformPrechecksUsingPOSTForbidden {
	return &PerformPrechecksUsingPOSTForbidden{}
}

/*
PerformPrechecksUsingPOSTForbidden describes a response with status code 403, with default header values.

Operation not allowed
*/
type PerformPrechecksUsingPOSTForbidden struct {
	Payload *models.Error
}

// IsSuccess returns true when this perform prechecks using p o s t forbidden response has a 2xx status code
func (o *PerformPrechecksUsingPOSTForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this perform prechecks using p o s t forbidden response has a 3xx status code
func (o *PerformPrechecksUsingPOSTForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this perform prechecks using p o s t forbidden response has a 4xx status code
func (o *PerformPrechecksUsingPOSTForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this perform prechecks using p o s t forbidden response has a 5xx status code
func (o *PerformPrechecksUsingPOSTForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this perform prechecks using p o s t forbidden response a status code equal to that given
func (o *PerformPrechecksUsingPOSTForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *PerformPrechecksUsingPOSTForbidden) Error() string {
	return fmt.Sprintf("[POST /v1/upgrades/{upgradeId}/prechecks][%d] performPrechecksUsingPOSTForbidden  %+v", 403, o.Payload)
}

func (o *PerformPrechecksUsingPOSTForbidden) String() string {
	return fmt.Sprintf("[POST /v1/upgrades/{upgradeId}/prechecks][%d] performPrechecksUsingPOSTForbidden  %+v", 403, o.Payload)
}

func (o *PerformPrechecksUsingPOSTForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *PerformPrechecksUsingPOSTForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPerformPrechecksUsingPOSTInternalServerError creates a PerformPrechecksUsingPOSTInternalServerError with default headers values
func NewPerformPrechecksUsingPOSTInternalServerError() *PerformPrechecksUsingPOSTInternalServerError {
	return &PerformPrechecksUsingPOSTInternalServerError{}
}

/*
PerformPrechecksUsingPOSTInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PerformPrechecksUsingPOSTInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this perform prechecks using p o s t internal server error response has a 2xx status code
func (o *PerformPrechecksUsingPOSTInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this perform prechecks using p o s t internal server error response has a 3xx status code
func (o *PerformPrechecksUsingPOSTInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this perform prechecks using p o s t internal server error response has a 4xx status code
func (o *PerformPrechecksUsingPOSTInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this perform prechecks using p o s t internal server error response has a 5xx status code
func (o *PerformPrechecksUsingPOSTInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this perform prechecks using p o s t internal server error response a status code equal to that given
func (o *PerformPrechecksUsingPOSTInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *PerformPrechecksUsingPOSTInternalServerError) Error() string {
	return fmt.Sprintf("[POST /v1/upgrades/{upgradeId}/prechecks][%d] performPrechecksUsingPOSTInternalServerError  %+v", 500, o.Payload)
}

func (o *PerformPrechecksUsingPOSTInternalServerError) String() string {
	return fmt.Sprintf("[POST /v1/upgrades/{upgradeId}/prechecks][%d] performPrechecksUsingPOSTInternalServerError  %+v", 500, o.Payload)
}

func (o *PerformPrechecksUsingPOSTInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *PerformPrechecksUsingPOSTInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}