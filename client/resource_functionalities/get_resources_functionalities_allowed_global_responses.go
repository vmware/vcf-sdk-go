// Code generated by go-swagger; DO NOT EDIT.

package resource_functionalities

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"vcf-sdk-go/models"
)

// GetResourcesFunctionalitiesAllowedGlobalReader is a Reader for the GetResourcesFunctionalitiesAllowedGlobal structure.
type GetResourcesFunctionalitiesAllowedGlobalReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetResourcesFunctionalitiesAllowedGlobalReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetResourcesFunctionalitiesAllowedGlobalOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetResourcesFunctionalitiesAllowedGlobalBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetResourcesFunctionalitiesAllowedGlobalInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetResourcesFunctionalitiesAllowedGlobalOK creates a GetResourcesFunctionalitiesAllowedGlobalOK with default headers values
func NewGetResourcesFunctionalitiesAllowedGlobalOK() *GetResourcesFunctionalitiesAllowedGlobalOK {
	return &GetResourcesFunctionalitiesAllowedGlobalOK{}
}

/*
GetResourcesFunctionalitiesAllowedGlobalOK describes a response with status code 200, with default header values.

OK
*/
type GetResourcesFunctionalitiesAllowedGlobalOK struct {
	Payload *models.ResourceFunctionalitiesGlobalConfiguration
}

// IsSuccess returns true when this get resources functionalities allowed global o k response has a 2xx status code
func (o *GetResourcesFunctionalitiesAllowedGlobalOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get resources functionalities allowed global o k response has a 3xx status code
func (o *GetResourcesFunctionalitiesAllowedGlobalOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get resources functionalities allowed global o k response has a 4xx status code
func (o *GetResourcesFunctionalitiesAllowedGlobalOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get resources functionalities allowed global o k response has a 5xx status code
func (o *GetResourcesFunctionalitiesAllowedGlobalOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get resources functionalities allowed global o k response a status code equal to that given
func (o *GetResourcesFunctionalitiesAllowedGlobalOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetResourcesFunctionalitiesAllowedGlobalOK) Error() string {
	return fmt.Sprintf("[GET /v1/resource-functionalities/global][%d] getResourcesFunctionalitiesAllowedGlobalOK  %+v", 200, o.Payload)
}

func (o *GetResourcesFunctionalitiesAllowedGlobalOK) String() string {
	return fmt.Sprintf("[GET /v1/resource-functionalities/global][%d] getResourcesFunctionalitiesAllowedGlobalOK  %+v", 200, o.Payload)
}

func (o *GetResourcesFunctionalitiesAllowedGlobalOK) GetPayload() *models.ResourceFunctionalitiesGlobalConfiguration {
	return o.Payload
}

func (o *GetResourcesFunctionalitiesAllowedGlobalOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResourceFunctionalitiesGlobalConfiguration)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetResourcesFunctionalitiesAllowedGlobalBadRequest creates a GetResourcesFunctionalitiesAllowedGlobalBadRequest with default headers values
func NewGetResourcesFunctionalitiesAllowedGlobalBadRequest() *GetResourcesFunctionalitiesAllowedGlobalBadRequest {
	return &GetResourcesFunctionalitiesAllowedGlobalBadRequest{}
}

/*
GetResourcesFunctionalitiesAllowedGlobalBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetResourcesFunctionalitiesAllowedGlobalBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this get resources functionalities allowed global bad request response has a 2xx status code
func (o *GetResourcesFunctionalitiesAllowedGlobalBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get resources functionalities allowed global bad request response has a 3xx status code
func (o *GetResourcesFunctionalitiesAllowedGlobalBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get resources functionalities allowed global bad request response has a 4xx status code
func (o *GetResourcesFunctionalitiesAllowedGlobalBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get resources functionalities allowed global bad request response has a 5xx status code
func (o *GetResourcesFunctionalitiesAllowedGlobalBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get resources functionalities allowed global bad request response a status code equal to that given
func (o *GetResourcesFunctionalitiesAllowedGlobalBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetResourcesFunctionalitiesAllowedGlobalBadRequest) Error() string {
	return fmt.Sprintf("[GET /v1/resource-functionalities/global][%d] getResourcesFunctionalitiesAllowedGlobalBadRequest  %+v", 400, o.Payload)
}

func (o *GetResourcesFunctionalitiesAllowedGlobalBadRequest) String() string {
	return fmt.Sprintf("[GET /v1/resource-functionalities/global][%d] getResourcesFunctionalitiesAllowedGlobalBadRequest  %+v", 400, o.Payload)
}

func (o *GetResourcesFunctionalitiesAllowedGlobalBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetResourcesFunctionalitiesAllowedGlobalBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetResourcesFunctionalitiesAllowedGlobalInternalServerError creates a GetResourcesFunctionalitiesAllowedGlobalInternalServerError with default headers values
func NewGetResourcesFunctionalitiesAllowedGlobalInternalServerError() *GetResourcesFunctionalitiesAllowedGlobalInternalServerError {
	return &GetResourcesFunctionalitiesAllowedGlobalInternalServerError{}
}

/*
GetResourcesFunctionalitiesAllowedGlobalInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetResourcesFunctionalitiesAllowedGlobalInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this get resources functionalities allowed global internal server error response has a 2xx status code
func (o *GetResourcesFunctionalitiesAllowedGlobalInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get resources functionalities allowed global internal server error response has a 3xx status code
func (o *GetResourcesFunctionalitiesAllowedGlobalInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get resources functionalities allowed global internal server error response has a 4xx status code
func (o *GetResourcesFunctionalitiesAllowedGlobalInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get resources functionalities allowed global internal server error response has a 5xx status code
func (o *GetResourcesFunctionalitiesAllowedGlobalInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get resources functionalities allowed global internal server error response a status code equal to that given
func (o *GetResourcesFunctionalitiesAllowedGlobalInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetResourcesFunctionalitiesAllowedGlobalInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/resource-functionalities/global][%d] getResourcesFunctionalitiesAllowedGlobalInternalServerError  %+v", 500, o.Payload)
}

func (o *GetResourcesFunctionalitiesAllowedGlobalInternalServerError) String() string {
	return fmt.Sprintf("[GET /v1/resource-functionalities/global][%d] getResourcesFunctionalitiesAllowedGlobalInternalServerError  %+v", 500, o.Payload)
}

func (o *GetResourcesFunctionalitiesAllowedGlobalInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetResourcesFunctionalitiesAllowedGlobalInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
