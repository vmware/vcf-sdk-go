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

// PostDatastoreQueryReader is a Reader for the PostDatastoreQuery structure.
type PostDatastoreQueryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostDatastoreQueryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostDatastoreQueryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostDatastoreQueryBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostDatastoreQueryInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostDatastoreQueryOK creates a PostDatastoreQueryOK with default headers values
func NewPostDatastoreQueryOK() *PostDatastoreQueryOK {
	return &PostDatastoreQueryOK{}
}

/*
PostDatastoreQueryOK describes a response with status code 200, with default header values.

Ok
*/
type PostDatastoreQueryOK struct {
	Payload *models.DatastoreQueryResponse
}

// IsSuccess returns true when this post datastore query o k response has a 2xx status code
func (o *PostDatastoreQueryOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post datastore query o k response has a 3xx status code
func (o *PostDatastoreQueryOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post datastore query o k response has a 4xx status code
func (o *PostDatastoreQueryOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this post datastore query o k response has a 5xx status code
func (o *PostDatastoreQueryOK) IsServerError() bool {
	return false
}

// IsCode returns true when this post datastore query o k response a status code equal to that given
func (o *PostDatastoreQueryOK) IsCode(code int) bool {
	return code == 200
}

func (o *PostDatastoreQueryOK) Error() string {
	return fmt.Sprintf("[POST /v1/clusters/{id}/datastores/queries][%d] postDatastoreQueryOK  %+v", 200, o.Payload)
}

func (o *PostDatastoreQueryOK) String() string {
	return fmt.Sprintf("[POST /v1/clusters/{id}/datastores/queries][%d] postDatastoreQueryOK  %+v", 200, o.Payload)
}

func (o *PostDatastoreQueryOK) GetPayload() *models.DatastoreQueryResponse {
	return o.Payload
}

func (o *PostDatastoreQueryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DatastoreQueryResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostDatastoreQueryBadRequest creates a PostDatastoreQueryBadRequest with default headers values
func NewPostDatastoreQueryBadRequest() *PostDatastoreQueryBadRequest {
	return &PostDatastoreQueryBadRequest{}
}

/*
PostDatastoreQueryBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PostDatastoreQueryBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this post datastore query bad request response has a 2xx status code
func (o *PostDatastoreQueryBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post datastore query bad request response has a 3xx status code
func (o *PostDatastoreQueryBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post datastore query bad request response has a 4xx status code
func (o *PostDatastoreQueryBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this post datastore query bad request response has a 5xx status code
func (o *PostDatastoreQueryBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this post datastore query bad request response a status code equal to that given
func (o *PostDatastoreQueryBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *PostDatastoreQueryBadRequest) Error() string {
	return fmt.Sprintf("[POST /v1/clusters/{id}/datastores/queries][%d] postDatastoreQueryBadRequest  %+v", 400, o.Payload)
}

func (o *PostDatastoreQueryBadRequest) String() string {
	return fmt.Sprintf("[POST /v1/clusters/{id}/datastores/queries][%d] postDatastoreQueryBadRequest  %+v", 400, o.Payload)
}

func (o *PostDatastoreQueryBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostDatastoreQueryBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostDatastoreQueryInternalServerError creates a PostDatastoreQueryInternalServerError with default headers values
func NewPostDatastoreQueryInternalServerError() *PostDatastoreQueryInternalServerError {
	return &PostDatastoreQueryInternalServerError{}
}

/*
PostDatastoreQueryInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PostDatastoreQueryInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this post datastore query internal server error response has a 2xx status code
func (o *PostDatastoreQueryInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post datastore query internal server error response has a 3xx status code
func (o *PostDatastoreQueryInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post datastore query internal server error response has a 4xx status code
func (o *PostDatastoreQueryInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this post datastore query internal server error response has a 5xx status code
func (o *PostDatastoreQueryInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this post datastore query internal server error response a status code equal to that given
func (o *PostDatastoreQueryInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *PostDatastoreQueryInternalServerError) Error() string {
	return fmt.Sprintf("[POST /v1/clusters/{id}/datastores/queries][%d] postDatastoreQueryInternalServerError  %+v", 500, o.Payload)
}

func (o *PostDatastoreQueryInternalServerError) String() string {
	return fmt.Sprintf("[POST /v1/clusters/{id}/datastores/queries][%d] postDatastoreQueryInternalServerError  %+v", 500, o.Payload)
}

func (o *PostDatastoreQueryInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostDatastoreQueryInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
