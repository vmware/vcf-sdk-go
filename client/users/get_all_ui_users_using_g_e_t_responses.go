// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"vcf-sdk-go/models"
)

// GetAllUIUsersUsingGETReader is a Reader for the GetAllUIUsersUsingGET structure.
type GetAllUIUsersUsingGETReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAllUIUsersUsingGETReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAllUIUsersUsingGETOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetAllUIUsersUsingGETOK creates a GetAllUIUsersUsingGETOK with default headers values
func NewGetAllUIUsersUsingGETOK() *GetAllUIUsersUsingGETOK {
	return &GetAllUIUsersUsingGETOK{}
}

/*
GetAllUIUsersUsingGETOK describes a response with status code 200, with default header values.

OK
*/
type GetAllUIUsersUsingGETOK struct {
	Payload *models.PageOfUser
}

// IsSuccess returns true when this get all Ui users using g e t o k response has a 2xx status code
func (o *GetAllUIUsersUsingGETOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get all Ui users using g e t o k response has a 3xx status code
func (o *GetAllUIUsersUsingGETOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get all Ui users using g e t o k response has a 4xx status code
func (o *GetAllUIUsersUsingGETOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get all Ui users using g e t o k response has a 5xx status code
func (o *GetAllUIUsersUsingGETOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get all Ui users using g e t o k response a status code equal to that given
func (o *GetAllUIUsersUsingGETOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetAllUIUsersUsingGETOK) Error() string {
	return fmt.Sprintf("[GET /v1/users/ui][%d] getAllUiUsersUsingGETOK  %+v", 200, o.Payload)
}

func (o *GetAllUIUsersUsingGETOK) String() string {
	return fmt.Sprintf("[GET /v1/users/ui][%d] getAllUiUsersUsingGETOK  %+v", 200, o.Payload)
}

func (o *GetAllUIUsersUsingGETOK) GetPayload() *models.PageOfUser {
	return o.Payload
}

func (o *GetAllUIUsersUsingGETOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PageOfUser)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}