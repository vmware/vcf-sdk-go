// Code generated by go-swagger; DO NOT EDIT.

package v_r_l_i

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"vcf-sdk-go/models"
)

// GetIntegratedDomains1Reader is a Reader for the GetIntegratedDomains1 structure.
type GetIntegratedDomains1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetIntegratedDomains1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetIntegratedDomains1OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetIntegratedDomains1OK creates a GetIntegratedDomains1OK with default headers values
func NewGetIntegratedDomains1OK() *GetIntegratedDomains1OK {
	return &GetIntegratedDomains1OK{}
}

/*
GetIntegratedDomains1OK describes a response with status code 200, with default header values.

OK
*/
type GetIntegratedDomains1OK struct {
	Payload *models.PageOfDomainIntegration
}

// IsSuccess returns true when this get integrated domains1 o k response has a 2xx status code
func (o *GetIntegratedDomains1OK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get integrated domains1 o k response has a 3xx status code
func (o *GetIntegratedDomains1OK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get integrated domains1 o k response has a 4xx status code
func (o *GetIntegratedDomains1OK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get integrated domains1 o k response has a 5xx status code
func (o *GetIntegratedDomains1OK) IsServerError() bool {
	return false
}

// IsCode returns true when this get integrated domains1 o k response a status code equal to that given
func (o *GetIntegratedDomains1OK) IsCode(code int) bool {
	return code == 200
}

func (o *GetIntegratedDomains1OK) Error() string {
	return fmt.Sprintf("[GET /v1/vrli/domains][%d] getIntegratedDomains1OK  %+v", 200, o.Payload)
}

func (o *GetIntegratedDomains1OK) String() string {
	return fmt.Sprintf("[GET /v1/vrli/domains][%d] getIntegratedDomains1OK  %+v", 200, o.Payload)
}

func (o *GetIntegratedDomains1OK) GetPayload() *models.PageOfDomainIntegration {
	return o.Payload
}

func (o *GetIntegratedDomains1OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PageOfDomainIntegration)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
