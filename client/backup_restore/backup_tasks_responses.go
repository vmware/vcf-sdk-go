// Code generated by go-swagger; DO NOT EDIT.

package backup_restore

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"vcf-sdk-go/models"
)

// BackupTasksReader is a Reader for the BackupTasks structure.
type BackupTasksReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *BackupTasksReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewBackupTasksOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewBackupTasksAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewBackupTasksBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewBackupTasksInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewBackupTasksOK creates a BackupTasksOK with default headers values
func NewBackupTasksOK() *BackupTasksOK {
	return &BackupTasksOK{}
}

/*
BackupTasksOK describes a response with status code 200, with default header values.

OK
*/
type BackupTasksOK struct {
	Payload *models.BackupTask
}

// IsSuccess returns true when this backup tasks o k response has a 2xx status code
func (o *BackupTasksOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this backup tasks o k response has a 3xx status code
func (o *BackupTasksOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this backup tasks o k response has a 4xx status code
func (o *BackupTasksOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this backup tasks o k response has a 5xx status code
func (o *BackupTasksOK) IsServerError() bool {
	return false
}

// IsCode returns true when this backup tasks o k response a status code equal to that given
func (o *BackupTasksOK) IsCode(code int) bool {
	return code == 200
}

func (o *BackupTasksOK) Error() string {
	return fmt.Sprintf("[POST /v1/backups/tasks][%d] backupTasksOK  %+v", 200, o.Payload)
}

func (o *BackupTasksOK) String() string {
	return fmt.Sprintf("[POST /v1/backups/tasks][%d] backupTasksOK  %+v", 200, o.Payload)
}

func (o *BackupTasksOK) GetPayload() *models.BackupTask {
	return o.Payload
}

func (o *BackupTasksOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BackupTask)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBackupTasksAccepted creates a BackupTasksAccepted with default headers values
func NewBackupTasksAccepted() *BackupTasksAccepted {
	return &BackupTasksAccepted{}
}

/*
BackupTasksAccepted describes a response with status code 202, with default header values.

Accepted
*/
type BackupTasksAccepted struct {
	Payload *models.BackupTask
}

// IsSuccess returns true when this backup tasks accepted response has a 2xx status code
func (o *BackupTasksAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this backup tasks accepted response has a 3xx status code
func (o *BackupTasksAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this backup tasks accepted response has a 4xx status code
func (o *BackupTasksAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this backup tasks accepted response has a 5xx status code
func (o *BackupTasksAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this backup tasks accepted response a status code equal to that given
func (o *BackupTasksAccepted) IsCode(code int) bool {
	return code == 202
}

func (o *BackupTasksAccepted) Error() string {
	return fmt.Sprintf("[POST /v1/backups/tasks][%d] backupTasksAccepted  %+v", 202, o.Payload)
}

func (o *BackupTasksAccepted) String() string {
	return fmt.Sprintf("[POST /v1/backups/tasks][%d] backupTasksAccepted  %+v", 202, o.Payload)
}

func (o *BackupTasksAccepted) GetPayload() *models.BackupTask {
	return o.Payload
}

func (o *BackupTasksAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BackupTask)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBackupTasksBadRequest creates a BackupTasksBadRequest with default headers values
func NewBackupTasksBadRequest() *BackupTasksBadRequest {
	return &BackupTasksBadRequest{}
}

/*
BackupTasksBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type BackupTasksBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this backup tasks bad request response has a 2xx status code
func (o *BackupTasksBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this backup tasks bad request response has a 3xx status code
func (o *BackupTasksBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this backup tasks bad request response has a 4xx status code
func (o *BackupTasksBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this backup tasks bad request response has a 5xx status code
func (o *BackupTasksBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this backup tasks bad request response a status code equal to that given
func (o *BackupTasksBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *BackupTasksBadRequest) Error() string {
	return fmt.Sprintf("[POST /v1/backups/tasks][%d] backupTasksBadRequest  %+v", 400, o.Payload)
}

func (o *BackupTasksBadRequest) String() string {
	return fmt.Sprintf("[POST /v1/backups/tasks][%d] backupTasksBadRequest  %+v", 400, o.Payload)
}

func (o *BackupTasksBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *BackupTasksBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBackupTasksInternalServerError creates a BackupTasksInternalServerError with default headers values
func NewBackupTasksInternalServerError() *BackupTasksInternalServerError {
	return &BackupTasksInternalServerError{}
}

/*
BackupTasksInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type BackupTasksInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this backup tasks internal server error response has a 2xx status code
func (o *BackupTasksInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this backup tasks internal server error response has a 3xx status code
func (o *BackupTasksInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this backup tasks internal server error response has a 4xx status code
func (o *BackupTasksInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this backup tasks internal server error response has a 5xx status code
func (o *BackupTasksInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this backup tasks internal server error response a status code equal to that given
func (o *BackupTasksInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *BackupTasksInternalServerError) Error() string {
	return fmt.Sprintf("[POST /v1/backups/tasks][%d] backupTasksInternalServerError  %+v", 500, o.Payload)
}

func (o *BackupTasksInternalServerError) String() string {
	return fmt.Sprintf("[POST /v1/backups/tasks][%d] backupTasksInternalServerError  %+v", 500, o.Payload)
}

func (o *BackupTasksInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *BackupTasksInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
