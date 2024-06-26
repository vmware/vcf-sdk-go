// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2023 VMware, Inc.
// SPDX-License-Identifier: BSD-2-Clause

package sddc

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// ExportBringupDetailReportReader is a Reader for the ExportBringupDetailReport structure.
type ExportBringupDetailReportReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExportBringupDetailReportReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewExportBringupDetailReportOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewExportBringupDetailReportNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewExportBringupDetailReportInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 501:
		result := NewExportBringupDetailReportNotImplemented()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /v1/sddcs/{id}/detail-report] exportBringupDetailReport", response, response.Code())
	}
}

// NewExportBringupDetailReportOK creates a ExportBringupDetailReportOK with default headers values
func NewExportBringupDetailReportOK() *ExportBringupDetailReportOK {
	return &ExportBringupDetailReportOK{}
}

/*
ExportBringupDetailReportOK describes a response with status code 200, with default header values.

OK
*/
type ExportBringupDetailReportOK struct {
	Payload string
}

// IsSuccess returns true when this export bringup detail report o k response has a 2xx status code
func (o *ExportBringupDetailReportOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this export bringup detail report o k response has a 3xx status code
func (o *ExportBringupDetailReportOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this export bringup detail report o k response has a 4xx status code
func (o *ExportBringupDetailReportOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this export bringup detail report o k response has a 5xx status code
func (o *ExportBringupDetailReportOK) IsServerError() bool {
	return false
}

// IsCode returns true when this export bringup detail report o k response a status code equal to that given
func (o *ExportBringupDetailReportOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the export bringup detail report o k response
func (o *ExportBringupDetailReportOK) Code() int {
	return 200
}

func (o *ExportBringupDetailReportOK) Error() string {
	return fmt.Sprintf("[GET /v1/sddcs/{id}/detail-report][%d] exportBringupDetailReportOK  %+v", 200, o.Payload)
}

func (o *ExportBringupDetailReportOK) String() string {
	return fmt.Sprintf("[GET /v1/sddcs/{id}/detail-report][%d] exportBringupDetailReportOK  %+v", 200, o.Payload)
}

func (o *ExportBringupDetailReportOK) GetPayload() string {
	return o.Payload
}

func (o *ExportBringupDetailReportOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewExportBringupDetailReportNoContent creates a ExportBringupDetailReportNoContent with default headers values
func NewExportBringupDetailReportNoContent() *ExportBringupDetailReportNoContent {
	return &ExportBringupDetailReportNoContent{}
}

/*
ExportBringupDetailReportNoContent describes a response with status code 204, with default header values.

No content
*/
type ExportBringupDetailReportNoContent struct {
	Payload string
}

// IsSuccess returns true when this export bringup detail report no content response has a 2xx status code
func (o *ExportBringupDetailReportNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this export bringup detail report no content response has a 3xx status code
func (o *ExportBringupDetailReportNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this export bringup detail report no content response has a 4xx status code
func (o *ExportBringupDetailReportNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this export bringup detail report no content response has a 5xx status code
func (o *ExportBringupDetailReportNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this export bringup detail report no content response a status code equal to that given
func (o *ExportBringupDetailReportNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the export bringup detail report no content response
func (o *ExportBringupDetailReportNoContent) Code() int {
	return 204
}

func (o *ExportBringupDetailReportNoContent) Error() string {
	return fmt.Sprintf("[GET /v1/sddcs/{id}/detail-report][%d] exportBringupDetailReportNoContent  %+v", 204, o.Payload)
}

func (o *ExportBringupDetailReportNoContent) String() string {
	return fmt.Sprintf("[GET /v1/sddcs/{id}/detail-report][%d] exportBringupDetailReportNoContent  %+v", 204, o.Payload)
}

func (o *ExportBringupDetailReportNoContent) GetPayload() string {
	return o.Payload
}

func (o *ExportBringupDetailReportNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewExportBringupDetailReportInternalServerError creates a ExportBringupDetailReportInternalServerError with default headers values
func NewExportBringupDetailReportInternalServerError() *ExportBringupDetailReportInternalServerError {
	return &ExportBringupDetailReportInternalServerError{}
}

/*
ExportBringupDetailReportInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type ExportBringupDetailReportInternalServerError struct {
}

// IsSuccess returns true when this export bringup detail report internal server error response has a 2xx status code
func (o *ExportBringupDetailReportInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this export bringup detail report internal server error response has a 3xx status code
func (o *ExportBringupDetailReportInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this export bringup detail report internal server error response has a 4xx status code
func (o *ExportBringupDetailReportInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this export bringup detail report internal server error response has a 5xx status code
func (o *ExportBringupDetailReportInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this export bringup detail report internal server error response a status code equal to that given
func (o *ExportBringupDetailReportInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the export bringup detail report internal server error response
func (o *ExportBringupDetailReportInternalServerError) Code() int {
	return 500
}

func (o *ExportBringupDetailReportInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/sddcs/{id}/detail-report][%d] exportBringupDetailReportInternalServerError ", 500)
}

func (o *ExportBringupDetailReportInternalServerError) String() string {
	return fmt.Sprintf("[GET /v1/sddcs/{id}/detail-report][%d] exportBringupDetailReportInternalServerError ", 500)
}

func (o *ExportBringupDetailReportInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewExportBringupDetailReportNotImplemented creates a ExportBringupDetailReportNotImplemented with default headers values
func NewExportBringupDetailReportNotImplemented() *ExportBringupDetailReportNotImplemented {
	return &ExportBringupDetailReportNotImplemented{}
}

/*
ExportBringupDetailReportNotImplemented describes a response with status code 501, with default header values.

Not Implemented
*/
type ExportBringupDetailReportNotImplemented struct {
}

// IsSuccess returns true when this export bringup detail report not implemented response has a 2xx status code
func (o *ExportBringupDetailReportNotImplemented) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this export bringup detail report not implemented response has a 3xx status code
func (o *ExportBringupDetailReportNotImplemented) IsRedirect() bool {
	return false
}

// IsClientError returns true when this export bringup detail report not implemented response has a 4xx status code
func (o *ExportBringupDetailReportNotImplemented) IsClientError() bool {
	return false
}

// IsServerError returns true when this export bringup detail report not implemented response has a 5xx status code
func (o *ExportBringupDetailReportNotImplemented) IsServerError() bool {
	return true
}

// IsCode returns true when this export bringup detail report not implemented response a status code equal to that given
func (o *ExportBringupDetailReportNotImplemented) IsCode(code int) bool {
	return code == 501
}

// Code gets the status code for the export bringup detail report not implemented response
func (o *ExportBringupDetailReportNotImplemented) Code() int {
	return 501
}

func (o *ExportBringupDetailReportNotImplemented) Error() string {
	return fmt.Sprintf("[GET /v1/sddcs/{id}/detail-report][%d] exportBringupDetailReportNotImplemented ", 501)
}

func (o *ExportBringupDetailReportNotImplemented) String() string {
	return fmt.Sprintf("[GET /v1/sddcs/{id}/detail-report][%d] exportBringupDetailReportNotImplemented ", 501)
}

func (o *ExportBringupDetailReportNotImplemented) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
