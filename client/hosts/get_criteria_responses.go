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

// GetCriteriaReader is a Reader for the GetCriteria structure.
type GetCriteriaReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCriteriaReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetCriteriaOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetCriteriaOK creates a GetCriteriaOK with default headers values
func NewGetCriteriaOK() *GetCriteriaOK {
	return &GetCriteriaOK{}
}

/*
GetCriteriaOK describes a response with status code 200, with default header values.

Ok
*/
type GetCriteriaOK struct {
	Payload *models.PageOfHostCriterion
}

// IsSuccess returns true when this get criteria o k response has a 2xx status code
func (o *GetCriteriaOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get criteria o k response has a 3xx status code
func (o *GetCriteriaOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get criteria o k response has a 4xx status code
func (o *GetCriteriaOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get criteria o k response has a 5xx status code
func (o *GetCriteriaOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get criteria o k response a status code equal to that given
func (o *GetCriteriaOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetCriteriaOK) Error() string {
	return fmt.Sprintf("[GET /v1/hosts/criteria][%d] getCriteriaOK  %+v", 200, o.Payload)
}

func (o *GetCriteriaOK) String() string {
	return fmt.Sprintf("[GET /v1/hosts/criteria][%d] getCriteriaOK  %+v", 200, o.Payload)
}

func (o *GetCriteriaOK) GetPayload() *models.PageOfHostCriterion {
	return o.Payload
}

func (o *GetCriteriaOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PageOfHostCriterion)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
