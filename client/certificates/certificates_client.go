// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2023 VMware, Inc.
// SPDX-License-Identifier: BSD-2-Clause

package certificates

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new certificates API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for certificates API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	ConfigureCertificateAuthority(params *ConfigureCertificateAuthorityParams, opts ...ClientOption) (*ConfigureCertificateAuthorityOK, error)

	CreateCertificateAuthority(params *CreateCertificateAuthorityParams, opts ...ClientOption) (*CreateCertificateAuthorityOK, error)

	DownloadCSR(params *DownloadCSRParams, opts ...ClientOption) (*DownloadCSROK, error)

	GenerateCertificates(params *GenerateCertificatesParams, opts ...ClientOption) (*GenerateCertificatesOK, *GenerateCertificatesAccepted, error)

	GeneratesCSRs(params *GeneratesCSRsParams, opts ...ClientOption) (*GeneratesCSRsOK, *GeneratesCSRsAccepted, error)

	GetCSRs(params *GetCSRsParams, opts ...ClientOption) (*GetCSRsOK, error)

	GetCertificateAuthorities(params *GetCertificateAuthoritiesParams, opts ...ClientOption) (*GetCertificateAuthoritiesOK, error)

	GetCertificateAuthorityByID(params *GetCertificateAuthorityByIDParams, opts ...ClientOption) (*GetCertificateAuthorityByIDOK, error)

	GetCertificatesByDomain(params *GetCertificatesByDomainParams, opts ...ClientOption) (*GetCertificatesByDomainOK, error)

	GetDomainCertificates(params *GetDomainCertificatesParams, opts ...ClientOption) (*GetDomainCertificatesOK, error)

	GetResourceCertificatesValidationByID(params *GetResourceCertificatesValidationByIDParams, opts ...ClientOption) (*GetResourceCertificatesValidationByIDOK, error)

	RemoveCertificateAuthority(params *RemoveCertificateAuthorityParams, opts ...ClientOption) (*RemoveCertificateAuthorityOK, *RemoveCertificateAuthorityNoContent, error)

	ReplaceCertificates(params *ReplaceCertificatesParams, opts ...ClientOption) (*ReplaceCertificatesOK, *ReplaceCertificatesAccepted, error)

	ReplaceResourceCertificates(params *ReplaceResourceCertificatesParams, opts ...ClientOption) (*ReplaceResourceCertificatesOK, error)

	UploadCertificates(params *UploadCertificatesParams, opts ...ClientOption) (*UploadCertificatesOK, error)

	ValidateResourceCertificates(params *ValidateResourceCertificatesParams, opts ...ClientOption) (*ValidateResourceCertificatesOK, *ValidateResourceCertificatesCreated, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
ConfigureCertificateAuthority updates the configuration of a certificate authority

Update the configuration of a Certificate Authority
*/
func (a *Client) ConfigureCertificateAuthority(params *ConfigureCertificateAuthorityParams, opts ...ClientOption) (*ConfigureCertificateAuthorityOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewConfigureCertificateAuthorityParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "configureCertificateAuthority",
		Method:             "PATCH",
		PathPattern:        "/v1/certificate-authorities",
		ProducesMediaTypes: []string{"*/*"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ConfigureCertificateAuthorityReader{formats: a.formats},
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
	success, ok := result.(*ConfigureCertificateAuthorityOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for configureCertificateAuthority: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
CreateCertificateAuthority configures integration with a certificate authority

Creates a certificate authority. This is required to generate signed certificates by supporting CAs.
*/
func (a *Client) CreateCertificateAuthority(params *CreateCertificateAuthorityParams, opts ...ClientOption) (*CreateCertificateAuthorityOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateCertificateAuthorityParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "createCertificateAuthority",
		Method:             "PUT",
		PathPattern:        "/v1/certificate-authorities",
		ProducesMediaTypes: []string{"*/*"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateCertificateAuthorityReader{formats: a.formats},
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
	success, ok := result.(*CreateCertificateAuthorityOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for createCertificateAuthority: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
DownloadCSR requests the download of CSR s for a domain in tar gz format

Download available CSR(s) in tar.gz format
*/
func (a *Client) DownloadCSR(params *DownloadCSRParams, opts ...ClientOption) (*DownloadCSROK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDownloadCSRParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "downloadCSR",
		Method:             "GET",
		PathPattern:        "/v1/domains/{id}/csrs/downloads",
		ProducesMediaTypes: []string{"application/octet-stream"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DownloadCSRReader{formats: a.formats},
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
	success, ok := result.(*DownloadCSROK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for downloadCSR: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GenerateCertificates requests the creation of signed certificate for resources of a domain

Generate certificate(s) for the selected resource(s) in a domain. CA must be configured and CSR must be generated beforehand.
*/
func (a *Client) GenerateCertificates(params *GenerateCertificatesParams, opts ...ClientOption) (*GenerateCertificatesOK, *GenerateCertificatesAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGenerateCertificatesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "generateCertificates",
		Method:             "PUT",
		PathPattern:        "/v1/domains/{id}/certificates",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GenerateCertificatesReader{formats: a.formats},
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
	case *GenerateCertificatesOK:
		return value, nil, nil
	case *GenerateCertificatesAccepted:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for certificates: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
	GeneratesCSRs requests the creation of certificate signing request CSR files for resources of a domain

	Generate CSR(s) for the selected resource(s) in the domain.

*Warning:*
_Avoid using wildcard certificates. Instead, use subdomain-specific certificates that are rotated often. A compromised wildcard certificate can lead to security repercussions_
*/
func (a *Client) GeneratesCSRs(params *GeneratesCSRsParams, opts ...ClientOption) (*GeneratesCSRsOK, *GeneratesCSRsAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGeneratesCSRsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "generatesCSRs",
		Method:             "PUT",
		PathPattern:        "/v1/domains/{id}/csrs",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GeneratesCSRsReader{formats: a.formats},
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
	case *GeneratesCSRsOK:
		return value, nil, nil
	case *GeneratesCSRsAccepted:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for certificates: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetCSRs requests available CSR s in JSON format for a domain

Get available CSR(s) in json format
*/
func (a *Client) GetCSRs(params *GetCSRsParams, opts ...ClientOption) (*GetCSRsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCSRsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getCSRs",
		Method:             "GET",
		PathPattern:        "/v1/domains/{id}/csrs",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetCSRsReader{formats: a.formats},
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
	success, ok := result.(*GetCSRsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getCSRs: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetCertificateAuthorities retrieves a list of certificate authorities

Get certificate authorities information
*/
func (a *Client) GetCertificateAuthorities(params *GetCertificateAuthoritiesParams, opts ...ClientOption) (*GetCertificateAuthoritiesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCertificateAuthoritiesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getCertificateAuthorities",
		Method:             "GET",
		PathPattern:        "/v1/certificate-authorities",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetCertificateAuthoritiesReader{formats: a.formats},
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
	success, ok := result.(*GetCertificateAuthoritiesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getCertificateAuthorities: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetCertificateAuthorityByID retrives the details of a certificate authority by ID

Get certificate authority information
*/
func (a *Client) GetCertificateAuthorityByID(params *GetCertificateAuthorityByIDParams, opts ...ClientOption) (*GetCertificateAuthorityByIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCertificateAuthorityByIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getCertificateAuthorityById",
		Method:             "GET",
		PathPattern:        "/v1/certificate-authorities/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetCertificateAuthorityByIDReader{formats: a.formats},
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
	success, ok := result.(*GetCertificateAuthorityByIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getCertificateAuthorityById: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetCertificatesByDomain retrieves the certificate details for all resources in a domain

View detailed metadata about the certificate(s) of all the resources in a domain
*/
func (a *Client) GetCertificatesByDomain(params *GetCertificatesByDomainParams, opts ...ClientOption) (*GetCertificatesByDomainOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCertificatesByDomainParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getCertificatesByDomain",
		Method:             "GET",
		PathPattern:        "/v1/domains/{id}/resource-certificates",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetCertificatesByDomainReader{formats: a.formats},
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
	success, ok := result.(*GetCertificatesByDomainOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getCertificatesByDomain: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetDomainCertificates retrieves the latest generated certificates for a domain by its ID

Get latest generated certificate(s) in a domain.
*/
func (a *Client) GetDomainCertificates(params *GetDomainCertificatesParams, opts ...ClientOption) (*GetDomainCertificatesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetDomainCertificatesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getDomainCertificates",
		Method:             "GET",
		PathPattern:        "/v1/domains/{id}/certificates",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetDomainCertificatesReader{formats: a.formats},
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
	success, ok := result.(*GetDomainCertificatesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getDomainCertificates: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetResourceCertificatesValidationByID retrieves the results of a certificate validation by its ID

Get the resource certificate validation result
*/
func (a *Client) GetResourceCertificatesValidationByID(params *GetResourceCertificatesValidationByIDParams, opts ...ClientOption) (*GetResourceCertificatesValidationByIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetResourceCertificatesValidationByIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getResourceCertificatesValidationByID",
		Method:             "GET",
		PathPattern:        "/v1/domains/{id}/resource-certificates/validations/{validationId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetResourceCertificatesValidationByIDReader{formats: a.formats},
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
	success, ok := result.(*GetResourceCertificatesValidationByIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getResourceCertificatesValidationByID: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
RemoveCertificateAuthority removes the configuration of a certificate authority

Deletes CA configuration file
*/
func (a *Client) RemoveCertificateAuthority(params *RemoveCertificateAuthorityParams, opts ...ClientOption) (*RemoveCertificateAuthorityOK, *RemoveCertificateAuthorityNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRemoveCertificateAuthorityParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "removeCertificateAuthority",
		Method:             "DELETE",
		PathPattern:        "/v1/certificate-authorities/{id}",
		ProducesMediaTypes: []string{"*/*"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &RemoveCertificateAuthorityReader{formats: a.formats},
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
	case *RemoveCertificateAuthorityOK:
		return value, nil, nil
	case *RemoveCertificateAuthorityNoContent:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for certificates: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
ReplaceCertificates replaces certificate s for the selected resource s in a domain

Replace certificate(s) for the selected resource(s) in a domain
*/
func (a *Client) ReplaceCertificates(params *ReplaceCertificatesParams, opts ...ClientOption) (*ReplaceCertificatesOK, *ReplaceCertificatesAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewReplaceCertificatesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "replaceCertificates",
		Method:             "PATCH",
		PathPattern:        "/v1/domains/{id}/certificates",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ReplaceCertificatesReader{formats: a.formats},
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
	case *ReplaceCertificatesOK:
		return value, nil, nil
	case *ReplaceCertificatesAccepted:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for certificates: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
ReplaceResourceCertificates replaces the certificate s for selected resources of a domain

Replace resource certificates
*/
func (a *Client) ReplaceResourceCertificates(params *ReplaceResourceCertificatesParams, opts ...ClientOption) (*ReplaceResourceCertificatesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewReplaceResourceCertificatesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "replaceResourceCertificates",
		Method:             "PUT",
		PathPattern:        "/v1/domains/{id}/resource-certificates",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ReplaceResourceCertificatesReader{formats: a.formats},
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
	success, ok := result.(*ReplaceResourceCertificatesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for replaceResourceCertificates: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
UploadCertificates imports certificate s to the certificate store for a domain

Upload certificates to the certificate store
*/
func (a *Client) UploadCertificates(params *UploadCertificatesParams, opts ...ClientOption) (*UploadCertificatesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUploadCertificatesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "uploadCertificates",
		Method:             "PUT",
		PathPattern:        "/v1/domains/{id}/certificates/uploads",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"multipart/form-data"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UploadCertificatesReader{formats: a.formats},
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
	success, ok := result.(*UploadCertificatesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for uploadCertificates: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
ValidateResourceCertificates performs validation of the resource certificate spec specification

Validate resource certificates
*/
func (a *Client) ValidateResourceCertificates(params *ValidateResourceCertificatesParams, opts ...ClientOption) (*ValidateResourceCertificatesOK, *ValidateResourceCertificatesCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewValidateResourceCertificatesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "validateResourceCertificates",
		Method:             "PUT",
		PathPattern:        "/v1/domains/{id}/resource-certificates/validations",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ValidateResourceCertificatesReader{formats: a.formats},
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
	case *ValidateResourceCertificatesOK:
		return value, nil, nil
	case *ValidateResourceCertificatesCreated:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for certificates: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
