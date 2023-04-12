// Code generated by go-swagger; DO NOT EDIT.

package v_s_a_n_health_check

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"vcf-sdk-go/models"
)

// GetHealthCheckStatusTaskReader is a Reader for the GetHealthCheckStatusTask structure.
type GetHealthCheckStatusTaskReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetHealthCheckStatusTaskReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetHealthCheckStatusTaskOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetHealthCheckStatusTaskBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetHealthCheckStatusTaskInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetHealthCheckStatusTaskOK creates a GetHealthCheckStatusTaskOK with default headers values
func NewGetHealthCheckStatusTaskOK() *GetHealthCheckStatusTaskOK {
	return &GetHealthCheckStatusTaskOK{}
}

/*
GetHealthCheckStatusTaskOK describes a response with status code 200, with default header values.

Ok
*/
type GetHealthCheckStatusTaskOK struct {
	Payload *models.HealthCheckTask
}

// IsSuccess returns true when this get health check status task o k response has a 2xx status code
func (o *GetHealthCheckStatusTaskOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get health check status task o k response has a 3xx status code
func (o *GetHealthCheckStatusTaskOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get health check status task o k response has a 4xx status code
func (o *GetHealthCheckStatusTaskOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get health check status task o k response has a 5xx status code
func (o *GetHealthCheckStatusTaskOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get health check status task o k response a status code equal to that given
func (o *GetHealthCheckStatusTaskOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetHealthCheckStatusTaskOK) Error() string {
	return fmt.Sprintf("[GET /v1/domains/{domainId}/health-checks/tasks/{taskId}][%d] getHealthCheckStatusTaskOK  %+v", 200, o.Payload)
}

func (o *GetHealthCheckStatusTaskOK) String() string {
	return fmt.Sprintf("[GET /v1/domains/{domainId}/health-checks/tasks/{taskId}][%d] getHealthCheckStatusTaskOK  %+v", 200, o.Payload)
}

func (o *GetHealthCheckStatusTaskOK) GetPayload() *models.HealthCheckTask {
	return o.Payload
}

func (o *GetHealthCheckStatusTaskOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HealthCheckTask)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetHealthCheckStatusTaskBadRequest creates a GetHealthCheckStatusTaskBadRequest with default headers values
func NewGetHealthCheckStatusTaskBadRequest() *GetHealthCheckStatusTaskBadRequest {
	return &GetHealthCheckStatusTaskBadRequest{}
}

/*
GetHealthCheckStatusTaskBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetHealthCheckStatusTaskBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this get health check status task bad request response has a 2xx status code
func (o *GetHealthCheckStatusTaskBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get health check status task bad request response has a 3xx status code
func (o *GetHealthCheckStatusTaskBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get health check status task bad request response has a 4xx status code
func (o *GetHealthCheckStatusTaskBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get health check status task bad request response has a 5xx status code
func (o *GetHealthCheckStatusTaskBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get health check status task bad request response a status code equal to that given
func (o *GetHealthCheckStatusTaskBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetHealthCheckStatusTaskBadRequest) Error() string {
	return fmt.Sprintf("[GET /v1/domains/{domainId}/health-checks/tasks/{taskId}][%d] getHealthCheckStatusTaskBadRequest  %+v", 400, o.Payload)
}

func (o *GetHealthCheckStatusTaskBadRequest) String() string {
	return fmt.Sprintf("[GET /v1/domains/{domainId}/health-checks/tasks/{taskId}][%d] getHealthCheckStatusTaskBadRequest  %+v", 400, o.Payload)
}

func (o *GetHealthCheckStatusTaskBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetHealthCheckStatusTaskBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetHealthCheckStatusTaskInternalServerError creates a GetHealthCheckStatusTaskInternalServerError with default headers values
func NewGetHealthCheckStatusTaskInternalServerError() *GetHealthCheckStatusTaskInternalServerError {
	return &GetHealthCheckStatusTaskInternalServerError{}
}

/*
GetHealthCheckStatusTaskInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetHealthCheckStatusTaskInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this get health check status task internal server error response has a 2xx status code
func (o *GetHealthCheckStatusTaskInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get health check status task internal server error response has a 3xx status code
func (o *GetHealthCheckStatusTaskInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get health check status task internal server error response has a 4xx status code
func (o *GetHealthCheckStatusTaskInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get health check status task internal server error response has a 5xx status code
func (o *GetHealthCheckStatusTaskInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get health check status task internal server error response a status code equal to that given
func (o *GetHealthCheckStatusTaskInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetHealthCheckStatusTaskInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/domains/{domainId}/health-checks/tasks/{taskId}][%d] getHealthCheckStatusTaskInternalServerError  %+v", 500, o.Payload)
}

func (o *GetHealthCheckStatusTaskInternalServerError) String() string {
	return fmt.Sprintf("[GET /v1/domains/{domainId}/health-checks/tasks/{taskId}][%d] getHealthCheckStatusTaskInternalServerError  %+v", 500, o.Payload)
}

func (o *GetHealthCheckStatusTaskInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetHealthCheckStatusTaskInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}