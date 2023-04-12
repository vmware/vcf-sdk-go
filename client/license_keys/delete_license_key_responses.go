// Code generated by go-swagger; DO NOT EDIT.

package license_keys

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"vcf-sdk-go/models"
)

// DeleteLicenseKeyReader is a Reader for the DeleteLicenseKey structure.
type DeleteLicenseKeyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteLicenseKeyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteLicenseKeyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewDeleteLicenseKeyNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewDeleteLicenseKeyNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteLicenseKeyInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteLicenseKeyOK creates a DeleteLicenseKeyOK with default headers values
func NewDeleteLicenseKeyOK() *DeleteLicenseKeyOK {
	return &DeleteLicenseKeyOK{}
}

/*
DeleteLicenseKeyOK describes a response with status code 200, with default header values.

OK
*/
type DeleteLicenseKeyOK struct {
}

// IsSuccess returns true when this delete license key o k response has a 2xx status code
func (o *DeleteLicenseKeyOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete license key o k response has a 3xx status code
func (o *DeleteLicenseKeyOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete license key o k response has a 4xx status code
func (o *DeleteLicenseKeyOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete license key o k response has a 5xx status code
func (o *DeleteLicenseKeyOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete license key o k response a status code equal to that given
func (o *DeleteLicenseKeyOK) IsCode(code int) bool {
	return code == 200
}

func (o *DeleteLicenseKeyOK) Error() string {
	return fmt.Sprintf("[DELETE /v1/license-keys/{key}][%d] deleteLicenseKeyOK ", 200)
}

func (o *DeleteLicenseKeyOK) String() string {
	return fmt.Sprintf("[DELETE /v1/license-keys/{key}][%d] deleteLicenseKeyOK ", 200)
}

func (o *DeleteLicenseKeyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteLicenseKeyNoContent creates a DeleteLicenseKeyNoContent with default headers values
func NewDeleteLicenseKeyNoContent() *DeleteLicenseKeyNoContent {
	return &DeleteLicenseKeyNoContent{}
}

/*
DeleteLicenseKeyNoContent describes a response with status code 204, with default header values.

No content
*/
type DeleteLicenseKeyNoContent struct {
}

// IsSuccess returns true when this delete license key no content response has a 2xx status code
func (o *DeleteLicenseKeyNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete license key no content response has a 3xx status code
func (o *DeleteLicenseKeyNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete license key no content response has a 4xx status code
func (o *DeleteLicenseKeyNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete license key no content response has a 5xx status code
func (o *DeleteLicenseKeyNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this delete license key no content response a status code equal to that given
func (o *DeleteLicenseKeyNoContent) IsCode(code int) bool {
	return code == 204
}

func (o *DeleteLicenseKeyNoContent) Error() string {
	return fmt.Sprintf("[DELETE /v1/license-keys/{key}][%d] deleteLicenseKeyNoContent ", 204)
}

func (o *DeleteLicenseKeyNoContent) String() string {
	return fmt.Sprintf("[DELETE /v1/license-keys/{key}][%d] deleteLicenseKeyNoContent ", 204)
}

func (o *DeleteLicenseKeyNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteLicenseKeyNotFound creates a DeleteLicenseKeyNotFound with default headers values
func NewDeleteLicenseKeyNotFound() *DeleteLicenseKeyNotFound {
	return &DeleteLicenseKeyNotFound{}
}

/*
DeleteLicenseKeyNotFound describes a response with status code 404, with default header values.

License key not found
*/
type DeleteLicenseKeyNotFound struct {
}

// IsSuccess returns true when this delete license key not found response has a 2xx status code
func (o *DeleteLicenseKeyNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete license key not found response has a 3xx status code
func (o *DeleteLicenseKeyNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete license key not found response has a 4xx status code
func (o *DeleteLicenseKeyNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete license key not found response has a 5xx status code
func (o *DeleteLicenseKeyNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this delete license key not found response a status code equal to that given
func (o *DeleteLicenseKeyNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *DeleteLicenseKeyNotFound) Error() string {
	return fmt.Sprintf("[DELETE /v1/license-keys/{key}][%d] deleteLicenseKeyNotFound ", 404)
}

func (o *DeleteLicenseKeyNotFound) String() string {
	return fmt.Sprintf("[DELETE /v1/license-keys/{key}][%d] deleteLicenseKeyNotFound ", 404)
}

func (o *DeleteLicenseKeyNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteLicenseKeyInternalServerError creates a DeleteLicenseKeyInternalServerError with default headers values
func NewDeleteLicenseKeyInternalServerError() *DeleteLicenseKeyInternalServerError {
	return &DeleteLicenseKeyInternalServerError{}
}

/*
DeleteLicenseKeyInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type DeleteLicenseKeyInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this delete license key internal server error response has a 2xx status code
func (o *DeleteLicenseKeyInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete license key internal server error response has a 3xx status code
func (o *DeleteLicenseKeyInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete license key internal server error response has a 4xx status code
func (o *DeleteLicenseKeyInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete license key internal server error response has a 5xx status code
func (o *DeleteLicenseKeyInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this delete license key internal server error response a status code equal to that given
func (o *DeleteLicenseKeyInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *DeleteLicenseKeyInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /v1/license-keys/{key}][%d] deleteLicenseKeyInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteLicenseKeyInternalServerError) String() string {
	return fmt.Sprintf("[DELETE /v1/license-keys/{key}][%d] deleteLicenseKeyInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteLicenseKeyInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteLicenseKeyInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
