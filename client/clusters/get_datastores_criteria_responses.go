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

// GetDatastoresCriteriaReader is a Reader for the GetDatastoresCriteria structure.
type GetDatastoresCriteriaReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDatastoresCriteriaReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetDatastoresCriteriaOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetDatastoresCriteriaOK creates a GetDatastoresCriteriaOK with default headers values
func NewGetDatastoresCriteriaOK() *GetDatastoresCriteriaOK {
	return &GetDatastoresCriteriaOK{}
}

/*
GetDatastoresCriteriaOK describes a response with status code 200, with default header values.

Ok
*/
type GetDatastoresCriteriaOK struct {
	Payload *models.PageOfDatastoreCriterion
}

// IsSuccess returns true when this get datastores criteria o k response has a 2xx status code
func (o *GetDatastoresCriteriaOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get datastores criteria o k response has a 3xx status code
func (o *GetDatastoresCriteriaOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get datastores criteria o k response has a 4xx status code
func (o *GetDatastoresCriteriaOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get datastores criteria o k response has a 5xx status code
func (o *GetDatastoresCriteriaOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get datastores criteria o k response a status code equal to that given
func (o *GetDatastoresCriteriaOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetDatastoresCriteriaOK) Error() string {
	return fmt.Sprintf("[GET /v1/clusters/{id}/datastores/criteria][%d] getDatastoresCriteriaOK  %+v", 200, o.Payload)
}

func (o *GetDatastoresCriteriaOK) String() string {
	return fmt.Sprintf("[GET /v1/clusters/{id}/datastores/criteria][%d] getDatastoresCriteriaOK  %+v", 200, o.Payload)
}

func (o *GetDatastoresCriteriaOK) GetPayload() *models.PageOfDatastoreCriterion {
	return o.Payload
}

func (o *GetDatastoresCriteriaOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PageOfDatastoreCriterion)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
