// Code generated by go-swagger; DO NOT EDIT.

package nsx_t_edge_clusters

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"vcf-sdk-go/models"
)

// GetValidationForCreateEdgeClusterReader is a Reader for the GetValidationForCreateEdgeCluster structure.
type GetValidationForCreateEdgeClusterReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetValidationForCreateEdgeClusterReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetValidationForCreateEdgeClusterOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetValidationForCreateEdgeClusterBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetValidationForCreateEdgeClusterInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetValidationForCreateEdgeClusterOK creates a GetValidationForCreateEdgeClusterOK with default headers values
func NewGetValidationForCreateEdgeClusterOK() *GetValidationForCreateEdgeClusterOK {
	return &GetValidationForCreateEdgeClusterOK{}
}

/*
GetValidationForCreateEdgeClusterOK describes a response with status code 200, with default header values.

OK
*/
type GetValidationForCreateEdgeClusterOK struct {
	Payload *models.Validation
}

// IsSuccess returns true when this get validation for create edge cluster o k response has a 2xx status code
func (o *GetValidationForCreateEdgeClusterOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get validation for create edge cluster o k response has a 3xx status code
func (o *GetValidationForCreateEdgeClusterOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get validation for create edge cluster o k response has a 4xx status code
func (o *GetValidationForCreateEdgeClusterOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get validation for create edge cluster o k response has a 5xx status code
func (o *GetValidationForCreateEdgeClusterOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get validation for create edge cluster o k response a status code equal to that given
func (o *GetValidationForCreateEdgeClusterOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetValidationForCreateEdgeClusterOK) Error() string {
	return fmt.Sprintf("[GET /v1/edge-clusters/validations/{id}][%d] getValidationForCreateEdgeClusterOK  %+v", 200, o.Payload)
}

func (o *GetValidationForCreateEdgeClusterOK) String() string {
	return fmt.Sprintf("[GET /v1/edge-clusters/validations/{id}][%d] getValidationForCreateEdgeClusterOK  %+v", 200, o.Payload)
}

func (o *GetValidationForCreateEdgeClusterOK) GetPayload() *models.Validation {
	return o.Payload
}

func (o *GetValidationForCreateEdgeClusterOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Validation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetValidationForCreateEdgeClusterBadRequest creates a GetValidationForCreateEdgeClusterBadRequest with default headers values
func NewGetValidationForCreateEdgeClusterBadRequest() *GetValidationForCreateEdgeClusterBadRequest {
	return &GetValidationForCreateEdgeClusterBadRequest{}
}

/*
GetValidationForCreateEdgeClusterBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetValidationForCreateEdgeClusterBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this get validation for create edge cluster bad request response has a 2xx status code
func (o *GetValidationForCreateEdgeClusterBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get validation for create edge cluster bad request response has a 3xx status code
func (o *GetValidationForCreateEdgeClusterBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get validation for create edge cluster bad request response has a 4xx status code
func (o *GetValidationForCreateEdgeClusterBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get validation for create edge cluster bad request response has a 5xx status code
func (o *GetValidationForCreateEdgeClusterBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get validation for create edge cluster bad request response a status code equal to that given
func (o *GetValidationForCreateEdgeClusterBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetValidationForCreateEdgeClusterBadRequest) Error() string {
	return fmt.Sprintf("[GET /v1/edge-clusters/validations/{id}][%d] getValidationForCreateEdgeClusterBadRequest  %+v", 400, o.Payload)
}

func (o *GetValidationForCreateEdgeClusterBadRequest) String() string {
	return fmt.Sprintf("[GET /v1/edge-clusters/validations/{id}][%d] getValidationForCreateEdgeClusterBadRequest  %+v", 400, o.Payload)
}

func (o *GetValidationForCreateEdgeClusterBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetValidationForCreateEdgeClusterBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetValidationForCreateEdgeClusterInternalServerError creates a GetValidationForCreateEdgeClusterInternalServerError with default headers values
func NewGetValidationForCreateEdgeClusterInternalServerError() *GetValidationForCreateEdgeClusterInternalServerError {
	return &GetValidationForCreateEdgeClusterInternalServerError{}
}

/*
GetValidationForCreateEdgeClusterInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type GetValidationForCreateEdgeClusterInternalServerError struct {
}

// IsSuccess returns true when this get validation for create edge cluster internal server error response has a 2xx status code
func (o *GetValidationForCreateEdgeClusterInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get validation for create edge cluster internal server error response has a 3xx status code
func (o *GetValidationForCreateEdgeClusterInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get validation for create edge cluster internal server error response has a 4xx status code
func (o *GetValidationForCreateEdgeClusterInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get validation for create edge cluster internal server error response has a 5xx status code
func (o *GetValidationForCreateEdgeClusterInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get validation for create edge cluster internal server error response a status code equal to that given
func (o *GetValidationForCreateEdgeClusterInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetValidationForCreateEdgeClusterInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/edge-clusters/validations/{id}][%d] getValidationForCreateEdgeClusterInternalServerError ", 500)
}

func (o *GetValidationForCreateEdgeClusterInternalServerError) String() string {
	return fmt.Sprintf("[GET /v1/edge-clusters/validations/{id}][%d] getValidationForCreateEdgeClusterInternalServerError ", 500)
}

func (o *GetValidationForCreateEdgeClusterInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}