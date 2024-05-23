// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2023 VMware, Inc.
// SPDX-License-Identifier: BSD-2-Clause

package sddc

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new sddc API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for sddc API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	ExportBringupDetailReport(params *ExportBringupDetailReportParams, opts ...ClientOption) (*ExportBringupDetailReportOK, *ExportBringupDetailReportNoContent, error)

	ExportBringupValidationReport(params *ExportBringupValidationReportParams, opts ...ClientOption) (*ExportBringupValidationReportOK, error)

	GetBringupAppInfo(params *GetBringupAppInfoParams, opts ...ClientOption) (*GetBringupAppInfoOK, error)

	GetBringupTaskByID(params *GetBringupTaskByIDParams, opts ...ClientOption) (*GetBringupTaskByIDOK, error)

	GetBringupTasks(params *GetBringupTasksParams, opts ...ClientOption) (*GetBringupTasksOK, error)

	GetBringupValidation(params *GetBringupValidationParams, opts ...ClientOption) (*GetBringupValidationOK, error)

	GetBringupValidations(params *GetBringupValidationsParams, opts ...ClientOption) (*GetBringupValidationsOK, error)

	GetSDDCManagerInfo(params *GetSDDCManagerInfoParams, opts ...ClientOption) (*GetSDDCManagerInfoOK, error)

	RetryBringupValidation(params *RetryBringupValidationParams, opts ...ClientOption) (*RetryBringupValidationOK, error)

	RetrySDDC(params *RetrySDDCParams, opts ...ClientOption) (*RetrySDDCOK, *RetrySDDCAccepted, error)

	StartBringup(params *StartBringupParams, opts ...ClientOption) (*StartBringupOK, *StartBringupAccepted, error)

	StartBringupSpecConversion(params *StartBringupSpecConversionParams, opts ...ClientOption) (*StartBringupSpecConversionOK, error)

	ValidateBringupSpec(params *ValidateBringupSpecParams, opts ...ClientOption) (*ValidateBringupSpecOK, *ValidateBringupSpecAccepted, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
ExportBringupDetailReport gets bringup report by ID

Returns the bringup report. Reports are generated in PDF and CSV formats.
*/
func (a *Client) ExportBringupDetailReport(params *ExportBringupDetailReportParams, opts ...ClientOption) (*ExportBringupDetailReportOK, *ExportBringupDetailReportNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewExportBringupDetailReportParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "exportBringupDetailReport",
		Method:             "GET",
		PathPattern:        "/v1/sddcs/{id}/detail-report",
		ProducesMediaTypes: []string{"application/pdf", "text/csv"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ExportBringupDetailReportReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *ExportBringupDetailReportOK:
		return value, nil, nil
	case *ExportBringupDetailReportNoContent:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for sddc: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
ExportBringupValidationReport gets validation report by ID

Returns the bringup report for a validation. Reports are generated in PDF format.
*/
func (a *Client) ExportBringupValidationReport(params *ExportBringupValidationReportParams, opts ...ClientOption) (*ExportBringupValidationReportOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewExportBringupValidationReportParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "exportBringupValidationReport",
		Method:             "GET",
		PathPattern:        "/v1/sddcs/validations/{validationId}/report",
		ProducesMediaTypes: []string{"application/pdf"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ExportBringupValidationReportReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ExportBringupValidationReportOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for exportBringupValidationReport: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetBringupAppInfo gets information about the bringup application

GET Method to retrieve information about Bringup app
*/
func (a *Client) GetBringupAppInfo(params *GetBringupAppInfoParams, opts ...ClientOption) (*GetBringupAppInfoOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetBringupAppInfoParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getBringupAppInfo",
		Method:             "GET",
		PathPattern:        "/v1/sddcs/about",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetBringupAppInfoReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetBringupAppInfoOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getBringupAppInfo: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetBringupTaskByID gets a bringup task by its id
*/
func (a *Client) GetBringupTaskByID(params *GetBringupTaskByIDParams, opts ...ClientOption) (*GetBringupTaskByIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetBringupTaskByIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getBringupTaskByID",
		Method:             "GET",
		PathPattern:        "/v1/sddcs/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetBringupTaskByIDReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetBringupTaskByIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getBringupTaskByID: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetBringupTasks retrieves all bringup tasks
*/
func (a *Client) GetBringupTasks(params *GetBringupTasksParams, opts ...ClientOption) (*GetBringupTasksOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetBringupTasksParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getBringupTasks",
		Method:             "GET",
		PathPattern:        "/v1/sddcs",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetBringupTasksReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetBringupTasksOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getBringupTasks: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetBringupValidation retrieves the results of a bringup validation by its ID
*/
func (a *Client) GetBringupValidation(params *GetBringupValidationParams, opts ...ClientOption) (*GetBringupValidationOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetBringupValidationParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getBringupValidation",
		Method:             "GET",
		PathPattern:        "/v1/sddcs/validations/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetBringupValidationReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetBringupValidationOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getBringupValidation: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetBringupValidations retrieves a list of bringup validations
*/
func (a *Client) GetBringupValidations(params *GetBringupValidationsParams, opts ...ClientOption) (*GetBringupValidationsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetBringupValidationsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getBringupValidations",
		Method:             "GET",
		PathPattern:        "/v1/sddcs/validations",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetBringupValidationsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetBringupValidationsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getBringupValidations: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetSDDCManagerInfo retrieves SDDC manager VM details

Retrieves the details of SDDC Manager VM
*/
func (a *Client) GetSDDCManagerInfo(params *GetSDDCManagerInfoParams, opts ...ClientOption) (*GetSDDCManagerInfoOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetSDDCManagerInfoParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getSddcManagerInfo",
		Method:             "GET",
		PathPattern:        "/v1/sddcs/{id}/sddc-manager",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetSDDCManagerInfoReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetSDDCManagerInfoOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getSddcManagerInfo: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
RetryBringupValidation retries bringup validation

Retry a completed SDDC validation
*/
func (a *Client) RetryBringupValidation(params *RetryBringupValidationParams, opts ...ClientOption) (*RetryBringupValidationOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRetryBringupValidationParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "retryBringupValidation",
		Method:             "PATCH",
		PathPattern:        "/v1/sddcs/validations/{id}",
		ProducesMediaTypes: []string{"*/*"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &RetryBringupValidationReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*RetryBringupValidationOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for retryBringupValidation: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
RetrySDDC retries failed SDDC creation
*/
func (a *Client) RetrySDDC(params *RetrySDDCParams, opts ...ClientOption) (*RetrySDDCOK, *RetrySDDCAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRetrySDDCParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "retrySddc",
		Method:             "PATCH",
		PathPattern:        "/v1/sddcs/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &RetrySDDCReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *RetrySDDCOK:
		return value, nil, nil
	case *RetrySDDCAccepted:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for sddc: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
StartBringup deploys a management domain
*/
func (a *Client) StartBringup(params *StartBringupParams, opts ...ClientOption) (*StartBringupOK, *StartBringupAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewStartBringupParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "startBringup",
		Method:             "POST",
		PathPattern:        "/v1/sddcs",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &StartBringupReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *StartBringupOK:
		return value, nil, nil
	case *StartBringupAccepted:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for sddc: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
StartBringupSpecConversion converts SDDC specification Json excel file

SDDC specification incorporates all the client inputs regarding VMW component parameters constituting the SDDC: NTP, DNS spec, ESXi, VC, VSAN, NSX spec et al.
*/
func (a *Client) StartBringupSpecConversion(params *StartBringupSpecConversionParams, opts ...ClientOption) (*StartBringupSpecConversionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewStartBringupSpecConversionParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "startBringupSpecConversion",
		Method:             "POST",
		PathPattern:        "/v1/system/sddc-spec-converter",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"multipart/form-data"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &StartBringupSpecConversionReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*StartBringupSpecConversionOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for startBringupSpecConversion: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
ValidateBringupSpec performs validation of the Sddc spec specification

SDDC specification incorporates all the client inputs regarding VMW component parameters constituting the SDDC: NTP, DNS spec, ESXi, VC, VSAN, NSX spec et al.
*/
func (a *Client) ValidateBringupSpec(params *ValidateBringupSpecParams, opts ...ClientOption) (*ValidateBringupSpecOK, *ValidateBringupSpecAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewValidateBringupSpecParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "validateBringupSpec",
		Method:             "POST",
		PathPattern:        "/v1/sddcs/validations",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ValidateBringupSpecReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *ValidateBringupSpecOK:
		return value, nil, nil
	case *ValidateBringupSpecAccepted:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for sddc: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
