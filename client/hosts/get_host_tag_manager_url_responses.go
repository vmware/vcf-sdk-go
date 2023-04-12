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

// GetHostTagManagerURLReader is a Reader for the GetHostTagManagerURL structure.
type GetHostTagManagerURLReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetHostTagManagerURLReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetHostTagManagerURLOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetHostTagManagerURLBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetHostTagManagerURLInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetHostTagManagerURLOK creates a GetHostTagManagerURLOK with default headers values
func NewGetHostTagManagerURLOK() *GetHostTagManagerURLOK {
	return &GetHostTagManagerURLOK{}
}

/*
GetHostTagManagerURLOK describes a response with status code 200, with default header values.

Ok
*/
type GetHostTagManagerURLOK struct {
	Payload *models.TagManagerModel
}

// IsSuccess returns true when this get host tag manager Url o k response has a 2xx status code
func (o *GetHostTagManagerURLOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get host tag manager Url o k response has a 3xx status code
func (o *GetHostTagManagerURLOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get host tag manager Url o k response has a 4xx status code
func (o *GetHostTagManagerURLOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get host tag manager Url o k response has a 5xx status code
func (o *GetHostTagManagerURLOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get host tag manager Url o k response a status code equal to that given
func (o *GetHostTagManagerURLOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetHostTagManagerURLOK) Error() string {
	return fmt.Sprintf("[GET /v1/hosts/{id}/tags/tag-manager][%d] getHostTagManagerUrlOK  %+v", 200, o.Payload)
}

func (o *GetHostTagManagerURLOK) String() string {
	return fmt.Sprintf("[GET /v1/hosts/{id}/tags/tag-manager][%d] getHostTagManagerUrlOK  %+v", 200, o.Payload)
}

func (o *GetHostTagManagerURLOK) GetPayload() *models.TagManagerModel {
	return o.Payload
}

func (o *GetHostTagManagerURLOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TagManagerModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetHostTagManagerURLBadRequest creates a GetHostTagManagerURLBadRequest with default headers values
func NewGetHostTagManagerURLBadRequest() *GetHostTagManagerURLBadRequest {
	return &GetHostTagManagerURLBadRequest{}
}

/*
GetHostTagManagerURLBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetHostTagManagerURLBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this get host tag manager Url bad request response has a 2xx status code
func (o *GetHostTagManagerURLBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get host tag manager Url bad request response has a 3xx status code
func (o *GetHostTagManagerURLBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get host tag manager Url bad request response has a 4xx status code
func (o *GetHostTagManagerURLBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get host tag manager Url bad request response has a 5xx status code
func (o *GetHostTagManagerURLBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get host tag manager Url bad request response a status code equal to that given
func (o *GetHostTagManagerURLBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetHostTagManagerURLBadRequest) Error() string {
	return fmt.Sprintf("[GET /v1/hosts/{id}/tags/tag-manager][%d] getHostTagManagerUrlBadRequest  %+v", 400, o.Payload)
}

func (o *GetHostTagManagerURLBadRequest) String() string {
	return fmt.Sprintf("[GET /v1/hosts/{id}/tags/tag-manager][%d] getHostTagManagerUrlBadRequest  %+v", 400, o.Payload)
}

func (o *GetHostTagManagerURLBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetHostTagManagerURLBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetHostTagManagerURLInternalServerError creates a GetHostTagManagerURLInternalServerError with default headers values
func NewGetHostTagManagerURLInternalServerError() *GetHostTagManagerURLInternalServerError {
	return &GetHostTagManagerURLInternalServerError{}
}

/*
GetHostTagManagerURLInternalServerError describes a response with status code 500, with default header values.

InternalServerError
*/
type GetHostTagManagerURLInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this get host tag manager Url internal server error response has a 2xx status code
func (o *GetHostTagManagerURLInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get host tag manager Url internal server error response has a 3xx status code
func (o *GetHostTagManagerURLInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get host tag manager Url internal server error response has a 4xx status code
func (o *GetHostTagManagerURLInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get host tag manager Url internal server error response has a 5xx status code
func (o *GetHostTagManagerURLInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get host tag manager Url internal server error response a status code equal to that given
func (o *GetHostTagManagerURLInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetHostTagManagerURLInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/hosts/{id}/tags/tag-manager][%d] getHostTagManagerUrlInternalServerError  %+v", 500, o.Payload)
}

func (o *GetHostTagManagerURLInternalServerError) String() string {
	return fmt.Sprintf("[GET /v1/hosts/{id}/tags/tag-manager][%d] getHostTagManagerUrlInternalServerError  %+v", 500, o.Payload)
}

func (o *GetHostTagManagerURLInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetHostTagManagerURLInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
