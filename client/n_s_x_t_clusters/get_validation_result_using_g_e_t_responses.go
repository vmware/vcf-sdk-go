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

// GetValidationResultUsingGETReader is a Reader for the GetValidationResultUsingGET structure.
type GetValidationResultUsingGETReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetValidationResultUsingGETReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetValidationResultUsingGETOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetValidationResultUsingGETOK creates a GetValidationResultUsingGETOK with default headers values
func NewGetValidationResultUsingGETOK() *GetValidationResultUsingGETOK {
	return &GetValidationResultUsingGETOK{}
}

/*
GetValidationResultUsingGETOK describes a response with status code 200, with default header values.

OK
*/
type GetValidationResultUsingGETOK struct {
	Payload *models.Validation
}

// IsSuccess returns true when this get validation result using g e t o k response has a 2xx status code
func (o *GetValidationResultUsingGETOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get validation result using g e t o k response has a 3xx status code
func (o *GetValidationResultUsingGETOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get validation result using g e t o k response has a 4xx status code
func (o *GetValidationResultUsingGETOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get validation result using g e t o k response has a 5xx status code
func (o *GetValidationResultUsingGETOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get validation result using g e t o k response a status code equal to that given
func (o *GetValidationResultUsingGETOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetValidationResultUsingGETOK) Error() string {
	return fmt.Sprintf("[GET /v1/nsxt-clusters/ip-address-pools/validations/{id}][%d] getValidationResultUsingGETOK  %+v", 200, o.Payload)
}

func (o *GetValidationResultUsingGETOK) String() string {
	return fmt.Sprintf("[GET /v1/nsxt-clusters/ip-address-pools/validations/{id}][%d] getValidationResultUsingGETOK  %+v", 200, o.Payload)
}

func (o *GetValidationResultUsingGETOK) GetPayload() *models.Validation {
	return o.Payload
}

func (o *GetValidationResultUsingGETOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Validation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
