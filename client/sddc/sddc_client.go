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
	ConvertToJSONSpec(params *ConvertToJSONSpecParams, opts ...ClientOption) (*ConvertToJSONSpecOK, error)

	CreateSDDC(params *CreateSDDCParams, opts ...ClientOption) (*CreateSDDCOK, *CreateSDDCAccepted, error)

	GetAllSDDCValidations(params *GetAllSDDCValidationsParams, opts ...ClientOption) (*GetAllSDDCValidationsOK, error)

	GetBringupDetailReport(params *GetBringupDetailReportParams, opts ...ClientOption) (*GetBringupDetailReportOK, *GetBringupDetailReportNoContent, error)

	GetBringupInfo(params *GetBringupInfoParams, opts ...ClientOption) (*GetBringupInfoOK, error)

	GetBringupValidationReport(params *GetBringupValidationReportParams, opts ...ClientOption) (*GetBringupValidationReportOK, error)

	GetSDDCManagerInfo(params *GetSDDCManagerInfoParams, opts ...ClientOption) (*GetSDDCManagerInfoOK, error)

	GetSDDCValidation(params *GetSDDCValidationParams, opts ...ClientOption) (*GetSDDCValidationOK, error)

	RetrieveAllSddcs(params *RetrieveAllSddcsParams, opts ...ClientOption) (*RetrieveAllSddcsOK, error)

	RetrieveSDDC(params *RetrieveSDDCParams, opts ...ClientOption) (*RetrieveSDDCOK, error)

	RetrySDDC(params *RetrySDDCParams, opts ...ClientOption) (*RetrySDDCOK, *RetrySDDCAccepted, error)

	RetrySDDCValidation(params *RetrySDDCValidationParams, opts ...ClientOption) (*RetrySDDCValidationOK, error)

	ValidateSDDCSpec(params *ValidateSDDCSpecParams, opts ...ClientOption) (*ValidateSDDCSpecOK, *ValidateSDDCSpecAccepted, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
ConvertToJSONSpec converts SDDC specification Json excel file

SDDC specification incorporates all the client inputs regarding VMW component parameters constituting the SDDC: NTP, DNS spec, ESXi, VC, VSAN, NSX spec et al.
*/
func (a *Client) ConvertToJSONSpec(params *ConvertToJSONSpecParams, opts ...ClientOption) (*ConvertToJSONSpecOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewConvertToJSONSpecParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "convertToJsonSpec",
		Method:             "POST",
		PathPattern:        "/v1/system/sddc-spec-converter",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"multipart/form-data"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ConvertToJSONSpecReader{formats: a.formats},
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
	success, ok := result.(*ConvertToJSONSpecOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for convertToJsonSpec: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
CreateSDDC creates SDDC
*/
func (a *Client) CreateSDDC(params *CreateSDDCParams, opts ...ClientOption) (*CreateSDDCOK, *CreateSDDCAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateSDDCParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "createSddc",
		Method:             "POST",
		PathPattern:        "/v1/sddcs",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateSDDCReader{formats: a.formats},
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
	case *CreateSDDCOK:
		return value, nil, nil
	case *CreateSDDCAccepted:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for sddc: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetAllSDDCValidations gets all SDDC specification validations
*/
func (a *Client) GetAllSDDCValidations(params *GetAllSDDCValidationsParams, opts ...ClientOption) (*GetAllSDDCValidationsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAllSDDCValidationsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getAllSddcValidations",
		Method:             "GET",
		PathPattern:        "/v1/sddcs/validations",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetAllSDDCValidationsReader{formats: a.formats},
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
	success, ok := result.(*GetAllSDDCValidationsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getAllSddcValidations: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetBringupDetailReport gets SDDC report by ID

Returns the bringup report. Reports are generated in PDF and CSV formats.
*/
func (a *Client) GetBringupDetailReport(params *GetBringupDetailReportParams, opts ...ClientOption) (*GetBringupDetailReportOK, *GetBringupDetailReportNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetBringupDetailReportParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getBringupDetailReport",
		Method:             "GET",
		PathPattern:        "/v1/sddcs/{id}/detail-report",
		ProducesMediaTypes: []string{"application/pdf", "text/csv"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetBringupDetailReportReader{formats: a.formats},
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
	case *GetBringupDetailReportOK:
		return value, nil, nil
	case *GetBringupDetailReportNoContent:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for sddc: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetBringupInfo gets bringup info

GET Method to retrieve information about Bringup app
*/
func (a *Client) GetBringupInfo(params *GetBringupInfoParams, opts ...ClientOption) (*GetBringupInfoOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetBringupInfoParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getBringupInfo",
		Method:             "GET",
		PathPattern:        "/v1/sddcs/about",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetBringupInfoReader{formats: a.formats},
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
	success, ok := result.(*GetBringupInfoOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getBringupInfo: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetBringupValidationReport gets validation report by ID

Returns the bringup report for a validation. Reports are generated in PDF format.
*/
func (a *Client) GetBringupValidationReport(params *GetBringupValidationReportParams, opts ...ClientOption) (*GetBringupValidationReportOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetBringupValidationReportParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getBringupValidationReport",
		Method:             "GET",
		PathPattern:        "/v1/sddcs/validations/{validationId}/report",
		ProducesMediaTypes: []string{"application/pdf"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetBringupValidationReportReader{formats: a.formats},
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
	success, ok := result.(*GetBringupValidationReportOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getBringupValidationReport: API contract not enforced by server. Client expected to get an error, but got: %T", result)
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
GetSDDCValidation gets SDDC specification validation status by ID
*/
func (a *Client) GetSDDCValidation(params *GetSDDCValidationParams, opts ...ClientOption) (*GetSDDCValidationOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetSDDCValidationParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getSddcValidation",
		Method:             "GET",
		PathPattern:        "/v1/sddcs/validations/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetSDDCValidationReader{formats: a.formats},
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
	success, ok := result.(*GetSDDCValidationOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getSddcValidation: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
RetrieveAllSddcs retrieves all s d d cs
*/
func (a *Client) RetrieveAllSddcs(params *RetrieveAllSddcsParams, opts ...ClientOption) (*RetrieveAllSddcsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRetrieveAllSddcsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "retrieveAllSddcs",
		Method:             "GET",
		PathPattern:        "/v1/sddcs",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &RetrieveAllSddcsReader{formats: a.formats},
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
	success, ok := result.(*RetrieveAllSddcsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for retrieveAllSddcs: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
RetrieveSDDC retrieves a SDDC
*/
func (a *Client) RetrieveSDDC(params *RetrieveSDDCParams, opts ...ClientOption) (*RetrieveSDDCOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRetrieveSDDCParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "retrieveSDDC",
		Method:             "GET",
		PathPattern:        "/v1/sddcs/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &RetrieveSDDCReader{formats: a.formats},
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
	success, ok := result.(*RetrieveSDDCOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for retrieveSDDC: API contract not enforced by server. Client expected to get an error, but got: %T", result)
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
RetrySDDCValidation retries SDDC validation

Retry a completed SDDC validation
*/
func (a *Client) RetrySDDCValidation(params *RetrySDDCValidationParams, opts ...ClientOption) (*RetrySDDCValidationOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRetrySDDCValidationParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "retrySddcValidation",
		Method:             "PATCH",
		PathPattern:        "/v1/sddcs/validations/{id}",
		ProducesMediaTypes: []string{"*/*"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &RetrySDDCValidationReader{formats: a.formats},
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
	success, ok := result.(*RetrySDDCValidationOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for retrySddcValidation: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
ValidateSDDCSpec validates SDDC specification before creation

SDDC specification incorporates all the client inputs regarding VMW component parameters constituting the SDDC: NTP, DNS spec, ESXi, VC, VSAN, NSX spec et al.
*/
func (a *Client) ValidateSDDCSpec(params *ValidateSDDCSpecParams, opts ...ClientOption) (*ValidateSDDCSpecOK, *ValidateSDDCSpecAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewValidateSDDCSpecParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "validateSddcSpec",
		Method:             "POST",
		PathPattern:        "/v1/sddcs/validations",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ValidateSDDCSpecReader{formats: a.formats},
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
	case *ValidateSDDCSpecOK:
		return value, nil, nil
	case *ValidateSDDCSpecAccepted:
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
