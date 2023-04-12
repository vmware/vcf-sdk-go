// Code generated by go-swagger; DO NOT EDIT.

package clusters

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"vcf-sdk-go/models"
)

// GetTagsAssignedToClustersReader is a Reader for the GetTagsAssignedToClusters structure.
type GetTagsAssignedToClustersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetTagsAssignedToClustersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetTagsAssignedToClustersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetTagsAssignedToClustersBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetTagsAssignedToClustersInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetTagsAssignedToClustersOK creates a GetTagsAssignedToClustersOK with default headers values
func NewGetTagsAssignedToClustersOK() *GetTagsAssignedToClustersOK {
	return &GetTagsAssignedToClustersOK{}
}

/*
GetTagsAssignedToClustersOK describes a response with status code 200, with default header values.

Ok
*/
type GetTagsAssignedToClustersOK struct {
	Payload *models.PageOfTagsForResource
}

// IsSuccess returns true when this get tags assigned to clusters o k response has a 2xx status code
func (o *GetTagsAssignedToClustersOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get tags assigned to clusters o k response has a 3xx status code
func (o *GetTagsAssignedToClustersOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get tags assigned to clusters o k response has a 4xx status code
func (o *GetTagsAssignedToClustersOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get tags assigned to clusters o k response has a 5xx status code
func (o *GetTagsAssignedToClustersOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get tags assigned to clusters o k response a status code equal to that given
func (o *GetTagsAssignedToClustersOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetTagsAssignedToClustersOK) Error() string {
	return fmt.Sprintf("[GET /v1/clusters/tags][%d] getTagsAssignedToClustersOK  %+v", 200, o.Payload)
}

func (o *GetTagsAssignedToClustersOK) String() string {
	return fmt.Sprintf("[GET /v1/clusters/tags][%d] getTagsAssignedToClustersOK  %+v", 200, o.Payload)
}

func (o *GetTagsAssignedToClustersOK) GetPayload() *models.PageOfTagsForResource {
	return o.Payload
}

func (o *GetTagsAssignedToClustersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PageOfTagsForResource)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTagsAssignedToClustersBadRequest creates a GetTagsAssignedToClustersBadRequest with default headers values
func NewGetTagsAssignedToClustersBadRequest() *GetTagsAssignedToClustersBadRequest {
	return &GetTagsAssignedToClustersBadRequest{}
}

/*
GetTagsAssignedToClustersBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetTagsAssignedToClustersBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this get tags assigned to clusters bad request response has a 2xx status code
func (o *GetTagsAssignedToClustersBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get tags assigned to clusters bad request response has a 3xx status code
func (o *GetTagsAssignedToClustersBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get tags assigned to clusters bad request response has a 4xx status code
func (o *GetTagsAssignedToClustersBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get tags assigned to clusters bad request response has a 5xx status code
func (o *GetTagsAssignedToClustersBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get tags assigned to clusters bad request response a status code equal to that given
func (o *GetTagsAssignedToClustersBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetTagsAssignedToClustersBadRequest) Error() string {
	return fmt.Sprintf("[GET /v1/clusters/tags][%d] getTagsAssignedToClustersBadRequest  %+v", 400, o.Payload)
}

func (o *GetTagsAssignedToClustersBadRequest) String() string {
	return fmt.Sprintf("[GET /v1/clusters/tags][%d] getTagsAssignedToClustersBadRequest  %+v", 400, o.Payload)
}

func (o *GetTagsAssignedToClustersBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetTagsAssignedToClustersBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTagsAssignedToClustersInternalServerError creates a GetTagsAssignedToClustersInternalServerError with default headers values
func NewGetTagsAssignedToClustersInternalServerError() *GetTagsAssignedToClustersInternalServerError {
	return &GetTagsAssignedToClustersInternalServerError{}
}

/*
GetTagsAssignedToClustersInternalServerError describes a response with status code 500, with default header values.

InternalServerError
*/
type GetTagsAssignedToClustersInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this get tags assigned to clusters internal server error response has a 2xx status code
func (o *GetTagsAssignedToClustersInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get tags assigned to clusters internal server error response has a 3xx status code
func (o *GetTagsAssignedToClustersInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get tags assigned to clusters internal server error response has a 4xx status code
func (o *GetTagsAssignedToClustersInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get tags assigned to clusters internal server error response has a 5xx status code
func (o *GetTagsAssignedToClustersInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get tags assigned to clusters internal server error response a status code equal to that given
func (o *GetTagsAssignedToClustersInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetTagsAssignedToClustersInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/clusters/tags][%d] getTagsAssignedToClustersInternalServerError  %+v", 500, o.Payload)
}

func (o *GetTagsAssignedToClustersInternalServerError) String() string {
	return fmt.Sprintf("[GET /v1/clusters/tags][%d] getTagsAssignedToClustersInternalServerError  %+v", 500, o.Payload)
}

func (o *GetTagsAssignedToClustersInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetTagsAssignedToClustersInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
