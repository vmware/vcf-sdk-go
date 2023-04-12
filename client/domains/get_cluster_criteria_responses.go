// Code generated by go-swagger; DO NOT EDIT.

package domains

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"vcf-sdk-go/models"
)

// GetClusterCriteriaReader is a Reader for the GetClusterCriteria structure.
type GetClusterCriteriaReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetClusterCriteriaReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetClusterCriteriaOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetClusterCriteriaOK creates a GetClusterCriteriaOK with default headers values
func NewGetClusterCriteriaOK() *GetClusterCriteriaOK {
	return &GetClusterCriteriaOK{}
}

/*
GetClusterCriteriaOK describes a response with status code 200, with default header values.

Ok
*/
type GetClusterCriteriaOK struct {
	Payload *models.PageOfClusterCriterion
}

// IsSuccess returns true when this get cluster criteria o k response has a 2xx status code
func (o *GetClusterCriteriaOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get cluster criteria o k response has a 3xx status code
func (o *GetClusterCriteriaOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get cluster criteria o k response has a 4xx status code
func (o *GetClusterCriteriaOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get cluster criteria o k response has a 5xx status code
func (o *GetClusterCriteriaOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get cluster criteria o k response a status code equal to that given
func (o *GetClusterCriteriaOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetClusterCriteriaOK) Error() string {
	return fmt.Sprintf("[GET /v1/domains/{domainId}/clusters/criteria][%d] getClusterCriteriaOK  %+v", 200, o.Payload)
}

func (o *GetClusterCriteriaOK) String() string {
	return fmt.Sprintf("[GET /v1/domains/{domainId}/clusters/criteria][%d] getClusterCriteriaOK  %+v", 200, o.Payload)
}

func (o *GetClusterCriteriaOK) GetPayload() *models.PageOfClusterCriterion {
	return o.Payload
}

func (o *GetClusterCriteriaOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PageOfClusterCriterion)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
