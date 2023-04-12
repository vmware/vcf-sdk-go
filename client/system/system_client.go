// Code generated by go-swagger; DO NOT EDIT.

package system

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new system API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for system API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	ConfigureDNS(params *ConfigureDNSParams, opts ...ClientOption) (*ConfigureDNSOK, *ConfigureDNSAccepted, error)

	ConfigureNtp(params *ConfigureNtpParams, opts ...ClientOption) (*ConfigureNtpOK, *ConfigureNtpAccepted, error)

	GetDNSConfiguration(params *GetDNSConfigurationParams, opts ...ClientOption) (*GetDNSConfigurationOK, error)

	GetNtpConfiguration(params *GetNtpConfigurationParams, opts ...ClientOption) (*GetNtpConfigurationOK, error)

	GetSystem(params *GetSystemParams, opts ...ClientOption) (*GetSystemOK, error)

	GetValidationOfDNSConfiguration(params *GetValidationOfDNSConfigurationParams, opts ...ClientOption) (*GetValidationOfDNSConfigurationOK, error)

	GetValidationOfNtpConfiguration(params *GetValidationOfNtpConfigurationParams, opts ...ClientOption) (*GetValidationOfNtpConfigurationOK, error)

	GetValidationsOfDNSConfiguration(params *GetValidationsOfDNSConfigurationParams, opts ...ClientOption) (*GetValidationsOfDNSConfigurationOK, error)

	GetValidationsOfNtpConfiguration(params *GetValidationsOfNtpConfigurationParams, opts ...ClientOption) (*GetValidationsOfNtpConfigurationOK, error)

	ValidateDNSConfiguration(params *ValidateDNSConfigurationParams, opts ...ClientOption) (*ValidateDNSConfigurationOK, *ValidateDNSConfigurationAccepted, error)

	ValidateNtpConfiguration(params *ValidateNtpConfigurationParams, opts ...ClientOption) (*ValidateNtpConfigurationOK, *ValidateNtpConfigurationAccepted, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
ConfigureDNS configures the Dns server
*/
func (a *Client) ConfigureDNS(params *ConfigureDNSParams, opts ...ClientOption) (*ConfigureDNSOK, *ConfigureDNSAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewConfigureDNSParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "configureDns",
		Method:             "PUT",
		PathPattern:        "/v1/system/dns-configuration",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ConfigureDNSReader{formats: a.formats},
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
	case *ConfigureDNSOK:
		return value, nil, nil
	case *ConfigureDNSAccepted:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for system: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
ConfigureNtp configures the ntp server
*/
func (a *Client) ConfigureNtp(params *ConfigureNtpParams, opts ...ClientOption) (*ConfigureNtpOK, *ConfigureNtpAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewConfigureNtpParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "configureNtp",
		Method:             "PUT",
		PathPattern:        "/v1/system/ntp-configuration",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ConfigureNtpReader{formats: a.formats},
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
	case *ConfigureNtpOK:
		return value, nil, nil
	case *ConfigureNtpAccepted:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for system: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetDNSConfiguration gets the current Dns configuration
*/
func (a *Client) GetDNSConfiguration(params *GetDNSConfigurationParams, opts ...ClientOption) (*GetDNSConfigurationOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetDNSConfigurationParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getDnsConfiguration",
		Method:             "GET",
		PathPattern:        "/v1/system/dns-configuration",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetDNSConfigurationReader{formats: a.formats},
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
	success, ok := result.(*GetDNSConfigurationOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getDnsConfiguration: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetNtpConfiguration gets the current ntp configuration
*/
func (a *Client) GetNtpConfiguration(params *GetNtpConfigurationParams, opts ...ClientOption) (*GetNtpConfigurationOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetNtpConfigurationParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getNtpConfiguration",
		Method:             "GET",
		PathPattern:        "/v1/system/ntp-configuration",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetNtpConfigurationReader{formats: a.formats},
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
	success, ok := result.(*GetNtpConfigurationOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getNtpConfiguration: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetSystem gets the system

Get the system
*/
func (a *Client) GetSystem(params *GetSystemParams, opts ...ClientOption) (*GetSystemOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetSystemParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getSystem",
		Method:             "GET",
		PathPattern:        "/v1/system",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetSystemReader{formats: a.formats},
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
	success, ok := result.(*GetSystemOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getSystem: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetValidationOfDNSConfiguration gets the status of the validation of the input Dns configuration
*/
func (a *Client) GetValidationOfDNSConfiguration(params *GetValidationOfDNSConfigurationParams, opts ...ClientOption) (*GetValidationOfDNSConfigurationOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetValidationOfDNSConfigurationParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getValidationOfDnsConfiguration",
		Method:             "GET",
		PathPattern:        "/v1/system/dns-configuration/validations/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetValidationOfDNSConfigurationReader{formats: a.formats},
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
	success, ok := result.(*GetValidationOfDNSConfigurationOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getValidationOfDnsConfiguration: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetValidationOfNtpConfiguration gets the status of the validation of the input ntp servers to configure new ntp server
*/
func (a *Client) GetValidationOfNtpConfiguration(params *GetValidationOfNtpConfigurationParams, opts ...ClientOption) (*GetValidationOfNtpConfigurationOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetValidationOfNtpConfigurationParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getValidationOfNtpConfiguration",
		Method:             "GET",
		PathPattern:        "/v1/system/ntp-configuration/validations/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetValidationOfNtpConfigurationReader{formats: a.formats},
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
	success, ok := result.(*GetValidationOfNtpConfigurationOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getValidationOfNtpConfiguration: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetValidationsOfDNSConfiguration gets the validations of the input dns servers to configure new DNS server
*/
func (a *Client) GetValidationsOfDNSConfiguration(params *GetValidationsOfDNSConfigurationParams, opts ...ClientOption) (*GetValidationsOfDNSConfigurationOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetValidationsOfDNSConfigurationParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getValidationsOfDNSConfiguration",
		Method:             "GET",
		PathPattern:        "/v1/system/dns-configuration/validations",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetValidationsOfDNSConfigurationReader{formats: a.formats},
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
	success, ok := result.(*GetValidationsOfDNSConfigurationOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getValidationsOfDNSConfiguration: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetValidationsOfNtpConfiguration gets the validations of the input ntp servers to configure new ntp server
*/
func (a *Client) GetValidationsOfNtpConfiguration(params *GetValidationsOfNtpConfigurationParams, opts ...ClientOption) (*GetValidationsOfNtpConfigurationOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetValidationsOfNtpConfigurationParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getValidationsOfNtpConfiguration",
		Method:             "GET",
		PathPattern:        "/v1/system/ntp-configuration/validations",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetValidationsOfNtpConfigurationReader{formats: a.formats},
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
	success, ok := result.(*GetValidationsOfNtpConfigurationOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getValidationsOfNtpConfiguration: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
ValidateDNSConfiguration validates Dns configuration input spec and system health before DNS configuration
*/
func (a *Client) ValidateDNSConfiguration(params *ValidateDNSConfigurationParams, opts ...ClientOption) (*ValidateDNSConfigurationOK, *ValidateDNSConfigurationAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewValidateDNSConfigurationParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "validateDnsConfiguration",
		Method:             "POST",
		PathPattern:        "/v1/system/dns-configuration/validations",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ValidateDNSConfigurationReader{formats: a.formats},
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
	case *ValidateDNSConfigurationOK:
		return value, nil, nil
	case *ValidateDNSConfigurationAccepted:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for system: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
ValidateNtpConfiguration validates ntp configuration input spec and system health before ntp configuration
*/
func (a *Client) ValidateNtpConfiguration(params *ValidateNtpConfigurationParams, opts ...ClientOption) (*ValidateNtpConfigurationOK, *ValidateNtpConfigurationAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewValidateNtpConfigurationParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "validateNtpConfiguration",
		Method:             "POST",
		PathPattern:        "/v1/system/ntp-configuration/validations",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ValidateNtpConfigurationReader{formats: a.formats},
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
	case *ValidateNtpConfigurationOK:
		return value, nil, nil
	case *ValidateNtpConfigurationAccepted:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for system: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}