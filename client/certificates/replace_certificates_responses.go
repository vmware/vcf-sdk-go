// Code generated by go-swagger; DO NOT EDIT.

package certificates

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"vcf-sdk-go/models"
)

// ReplaceCertificatesReader is a Reader for the ReplaceCertificates structure.
type ReplaceCertificatesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ReplaceCertificatesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewReplaceCertificatesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewReplaceCertificatesAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewReplaceCertificatesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewReplaceCertificatesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewReplaceCertificatesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewReplaceCertificatesConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewReplaceCertificatesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewReplaceCertificatesOK creates a ReplaceCertificatesOK with default headers values
func NewReplaceCertificatesOK() *ReplaceCertificatesOK {
	return &ReplaceCertificatesOK{}
}

/*
ReplaceCertificatesOK describes a response with status code 200, with default header values.

OK
*/
type ReplaceCertificatesOK struct {
	Payload *models.Task
}

// IsSuccess returns true when this replace certificates o k response has a 2xx status code
func (o *ReplaceCertificatesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this replace certificates o k response has a 3xx status code
func (o *ReplaceCertificatesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this replace certificates o k response has a 4xx status code
func (o *ReplaceCertificatesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this replace certificates o k response has a 5xx status code
func (o *ReplaceCertificatesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this replace certificates o k response a status code equal to that given
func (o *ReplaceCertificatesOK) IsCode(code int) bool {
	return code == 200
}

func (o *ReplaceCertificatesOK) Error() string {
	return fmt.Sprintf("[PATCH /v1/domains/{domainName}/certificates][%d] replaceCertificatesOK  %+v", 200, o.Payload)
}

func (o *ReplaceCertificatesOK) String() string {
	return fmt.Sprintf("[PATCH /v1/domains/{domainName}/certificates][%d] replaceCertificatesOK  %+v", 200, o.Payload)
}

func (o *ReplaceCertificatesOK) GetPayload() *models.Task {
	return o.Payload
}

func (o *ReplaceCertificatesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Task)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReplaceCertificatesAccepted creates a ReplaceCertificatesAccepted with default headers values
func NewReplaceCertificatesAccepted() *ReplaceCertificatesAccepted {
	return &ReplaceCertificatesAccepted{}
}

/*
ReplaceCertificatesAccepted describes a response with status code 202, with default header values.

Accepted
*/
type ReplaceCertificatesAccepted struct {
	Payload *models.Task
}

// IsSuccess returns true when this replace certificates accepted response has a 2xx status code
func (o *ReplaceCertificatesAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this replace certificates accepted response has a 3xx status code
func (o *ReplaceCertificatesAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this replace certificates accepted response has a 4xx status code
func (o *ReplaceCertificatesAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this replace certificates accepted response has a 5xx status code
func (o *ReplaceCertificatesAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this replace certificates accepted response a status code equal to that given
func (o *ReplaceCertificatesAccepted) IsCode(code int) bool {
	return code == 202
}

func (o *ReplaceCertificatesAccepted) Error() string {
	return fmt.Sprintf("[PATCH /v1/domains/{domainName}/certificates][%d] replaceCertificatesAccepted  %+v", 202, o.Payload)
}

func (o *ReplaceCertificatesAccepted) String() string {
	return fmt.Sprintf("[PATCH /v1/domains/{domainName}/certificates][%d] replaceCertificatesAccepted  %+v", 202, o.Payload)
}

func (o *ReplaceCertificatesAccepted) GetPayload() *models.Task {
	return o.Payload
}

func (o *ReplaceCertificatesAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Task)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReplaceCertificatesBadRequest creates a ReplaceCertificatesBadRequest with default headers values
func NewReplaceCertificatesBadRequest() *ReplaceCertificatesBadRequest {
	return &ReplaceCertificatesBadRequest{}
}

/*
ReplaceCertificatesBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type ReplaceCertificatesBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this replace certificates bad request response has a 2xx status code
func (o *ReplaceCertificatesBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this replace certificates bad request response has a 3xx status code
func (o *ReplaceCertificatesBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this replace certificates bad request response has a 4xx status code
func (o *ReplaceCertificatesBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this replace certificates bad request response has a 5xx status code
func (o *ReplaceCertificatesBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this replace certificates bad request response a status code equal to that given
func (o *ReplaceCertificatesBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *ReplaceCertificatesBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /v1/domains/{domainName}/certificates][%d] replaceCertificatesBadRequest  %+v", 400, o.Payload)
}

func (o *ReplaceCertificatesBadRequest) String() string {
	return fmt.Sprintf("[PATCH /v1/domains/{domainName}/certificates][%d] replaceCertificatesBadRequest  %+v", 400, o.Payload)
}

func (o *ReplaceCertificatesBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *ReplaceCertificatesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReplaceCertificatesForbidden creates a ReplaceCertificatesForbidden with default headers values
func NewReplaceCertificatesForbidden() *ReplaceCertificatesForbidden {
	return &ReplaceCertificatesForbidden{}
}

/*
ReplaceCertificatesForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type ReplaceCertificatesForbidden struct {
	Payload *models.Error
}

// IsSuccess returns true when this replace certificates forbidden response has a 2xx status code
func (o *ReplaceCertificatesForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this replace certificates forbidden response has a 3xx status code
func (o *ReplaceCertificatesForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this replace certificates forbidden response has a 4xx status code
func (o *ReplaceCertificatesForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this replace certificates forbidden response has a 5xx status code
func (o *ReplaceCertificatesForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this replace certificates forbidden response a status code equal to that given
func (o *ReplaceCertificatesForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *ReplaceCertificatesForbidden) Error() string {
	return fmt.Sprintf("[PATCH /v1/domains/{domainName}/certificates][%d] replaceCertificatesForbidden  %+v", 403, o.Payload)
}

func (o *ReplaceCertificatesForbidden) String() string {
	return fmt.Sprintf("[PATCH /v1/domains/{domainName}/certificates][%d] replaceCertificatesForbidden  %+v", 403, o.Payload)
}

func (o *ReplaceCertificatesForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *ReplaceCertificatesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReplaceCertificatesNotFound creates a ReplaceCertificatesNotFound with default headers values
func NewReplaceCertificatesNotFound() *ReplaceCertificatesNotFound {
	return &ReplaceCertificatesNotFound{}
}

/*
ReplaceCertificatesNotFound describes a response with status code 404, with default header values.

Not Found
*/
type ReplaceCertificatesNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this replace certificates not found response has a 2xx status code
func (o *ReplaceCertificatesNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this replace certificates not found response has a 3xx status code
func (o *ReplaceCertificatesNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this replace certificates not found response has a 4xx status code
func (o *ReplaceCertificatesNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this replace certificates not found response has a 5xx status code
func (o *ReplaceCertificatesNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this replace certificates not found response a status code equal to that given
func (o *ReplaceCertificatesNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *ReplaceCertificatesNotFound) Error() string {
	return fmt.Sprintf("[PATCH /v1/domains/{domainName}/certificates][%d] replaceCertificatesNotFound  %+v", 404, o.Payload)
}

func (o *ReplaceCertificatesNotFound) String() string {
	return fmt.Sprintf("[PATCH /v1/domains/{domainName}/certificates][%d] replaceCertificatesNotFound  %+v", 404, o.Payload)
}

func (o *ReplaceCertificatesNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *ReplaceCertificatesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReplaceCertificatesConflict creates a ReplaceCertificatesConflict with default headers values
func NewReplaceCertificatesConflict() *ReplaceCertificatesConflict {
	return &ReplaceCertificatesConflict{}
}

/*
ReplaceCertificatesConflict describes a response with status code 409, with default header values.

Conflict
*/
type ReplaceCertificatesConflict struct {
	Payload *models.Error
}

// IsSuccess returns true when this replace certificates conflict response has a 2xx status code
func (o *ReplaceCertificatesConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this replace certificates conflict response has a 3xx status code
func (o *ReplaceCertificatesConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this replace certificates conflict response has a 4xx status code
func (o *ReplaceCertificatesConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this replace certificates conflict response has a 5xx status code
func (o *ReplaceCertificatesConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this replace certificates conflict response a status code equal to that given
func (o *ReplaceCertificatesConflict) IsCode(code int) bool {
	return code == 409
}

func (o *ReplaceCertificatesConflict) Error() string {
	return fmt.Sprintf("[PATCH /v1/domains/{domainName}/certificates][%d] replaceCertificatesConflict  %+v", 409, o.Payload)
}

func (o *ReplaceCertificatesConflict) String() string {
	return fmt.Sprintf("[PATCH /v1/domains/{domainName}/certificates][%d] replaceCertificatesConflict  %+v", 409, o.Payload)
}

func (o *ReplaceCertificatesConflict) GetPayload() *models.Error {
	return o.Payload
}

func (o *ReplaceCertificatesConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReplaceCertificatesInternalServerError creates a ReplaceCertificatesInternalServerError with default headers values
func NewReplaceCertificatesInternalServerError() *ReplaceCertificatesInternalServerError {
	return &ReplaceCertificatesInternalServerError{}
}

/*
ReplaceCertificatesInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type ReplaceCertificatesInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this replace certificates internal server error response has a 2xx status code
func (o *ReplaceCertificatesInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this replace certificates internal server error response has a 3xx status code
func (o *ReplaceCertificatesInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this replace certificates internal server error response has a 4xx status code
func (o *ReplaceCertificatesInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this replace certificates internal server error response has a 5xx status code
func (o *ReplaceCertificatesInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this replace certificates internal server error response a status code equal to that given
func (o *ReplaceCertificatesInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *ReplaceCertificatesInternalServerError) Error() string {
	return fmt.Sprintf("[PATCH /v1/domains/{domainName}/certificates][%d] replaceCertificatesInternalServerError  %+v", 500, o.Payload)
}

func (o *ReplaceCertificatesInternalServerError) String() string {
	return fmt.Sprintf("[PATCH /v1/domains/{domainName}/certificates][%d] replaceCertificatesInternalServerError  %+v", 500, o.Payload)
}

func (o *ReplaceCertificatesInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *ReplaceCertificatesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
