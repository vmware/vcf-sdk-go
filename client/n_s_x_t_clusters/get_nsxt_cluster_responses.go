// Code generated by go-swagger; DO NOT EDIT.

package n_s_x_t_clusters

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"vcf-sdk-go/models"
)

// GetNsxtClusterReader is a Reader for the GetNsxtCluster structure.
type GetNsxtClusterReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetNsxtClusterReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetNsxtClusterOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetNsxtClusterNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetNsxtClusterInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetNsxtClusterOK creates a GetNsxtClusterOK with default headers values
func NewGetNsxtClusterOK() *GetNsxtClusterOK {
	return &GetNsxtClusterOK{}
}

/*
GetNsxtClusterOK describes a response with status code 200, with default header values.

Ok
*/
type GetNsxtClusterOK struct {
	Payload *models.NsxTCluster
}

// IsSuccess returns true when this get nsxt cluster o k response has a 2xx status code
func (o *GetNsxtClusterOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get nsxt cluster o k response has a 3xx status code
func (o *GetNsxtClusterOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get nsxt cluster o k response has a 4xx status code
func (o *GetNsxtClusterOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get nsxt cluster o k response has a 5xx status code
func (o *GetNsxtClusterOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get nsxt cluster o k response a status code equal to that given
func (o *GetNsxtClusterOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetNsxtClusterOK) Error() string {
	return fmt.Sprintf("[GET /v1/nsxt-clusters/{id}][%d] getNsxtClusterOK  %+v", 200, o.Payload)
}

func (o *GetNsxtClusterOK) String() string {
	return fmt.Sprintf("[GET /v1/nsxt-clusters/{id}][%d] getNsxtClusterOK  %+v", 200, o.Payload)
}

func (o *GetNsxtClusterOK) GetPayload() *models.NsxTCluster {
	return o.Payload
}

func (o *GetNsxtClusterOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NsxTCluster)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetNsxtClusterNotFound creates a GetNsxtClusterNotFound with default headers values
func NewGetNsxtClusterNotFound() *GetNsxtClusterNotFound {
	return &GetNsxtClusterNotFound{}
}

/*
GetNsxtClusterNotFound describes a response with status code 404, with default header values.

NSX-T cluster not found
*/
type GetNsxtClusterNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this get nsxt cluster not found response has a 2xx status code
func (o *GetNsxtClusterNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get nsxt cluster not found response has a 3xx status code
func (o *GetNsxtClusterNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get nsxt cluster not found response has a 4xx status code
func (o *GetNsxtClusterNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get nsxt cluster not found response has a 5xx status code
func (o *GetNsxtClusterNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get nsxt cluster not found response a status code equal to that given
func (o *GetNsxtClusterNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetNsxtClusterNotFound) Error() string {
	return fmt.Sprintf("[GET /v1/nsxt-clusters/{id}][%d] getNsxtClusterNotFound  %+v", 404, o.Payload)
}

func (o *GetNsxtClusterNotFound) String() string {
	return fmt.Sprintf("[GET /v1/nsxt-clusters/{id}][%d] getNsxtClusterNotFound  %+v", 404, o.Payload)
}

func (o *GetNsxtClusterNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetNsxtClusterNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetNsxtClusterInternalServerError creates a GetNsxtClusterInternalServerError with default headers values
func NewGetNsxtClusterInternalServerError() *GetNsxtClusterInternalServerError {
	return &GetNsxtClusterInternalServerError{}
}

/*
GetNsxtClusterInternalServerError describes a response with status code 500, with default header values.

InternalServerError
*/
type GetNsxtClusterInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this get nsxt cluster internal server error response has a 2xx status code
func (o *GetNsxtClusterInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get nsxt cluster internal server error response has a 3xx status code
func (o *GetNsxtClusterInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get nsxt cluster internal server error response has a 4xx status code
func (o *GetNsxtClusterInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get nsxt cluster internal server error response has a 5xx status code
func (o *GetNsxtClusterInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get nsxt cluster internal server error response a status code equal to that given
func (o *GetNsxtClusterInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetNsxtClusterInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/nsxt-clusters/{id}][%d] getNsxtClusterInternalServerError  %+v", 500, o.Payload)
}

func (o *GetNsxtClusterInternalServerError) String() string {
	return fmt.Sprintf("[GET /v1/nsxt-clusters/{id}][%d] getNsxtClusterInternalServerError  %+v", 500, o.Payload)
}

func (o *GetNsxtClusterInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetNsxtClusterInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
