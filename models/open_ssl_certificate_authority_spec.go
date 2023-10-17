// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2023 VMware, Inc.
// SPDX-License-Identifier: BSD-2-Clause

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// OpenSSLCertificateAuthoritySpec This specification contains OpenSSL CA configuration details
//
// swagger:model OpenSSLCertificateAuthoritySpec
type OpenSSLCertificateAuthoritySpec struct {

	// OpenSSL CA domain name
	// Example: OpenSSL CA
	// Required: true
	CommonName *string `json:"commonName"`

	// ISO 3166 country code where company is legally registered
	// Example: IN
	// Required: true
	Country *string `json:"country"`

	// The city or locality where company is legally registered
	// Example: Bengaluru
	// Required: true
	Locality *string `json:"locality"`

	// The name under which company is legally registered
	// Example: VMware Inc.
	// Required: true
	Organization *string `json:"organization"`

	// Organization with which the certificate is associated
	// Example: VCF
	// Required: true
	OrganizationUnit *string `json:"organizationUnit"`

	// The full name of the state where company is legally registered
	// Example: Karnataka
	// Required: true
	State *string `json:"state"`
}

// Validate validates this open SSL certificate authority spec
func (m *OpenSSLCertificateAuthoritySpec) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCommonName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCountry(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLocality(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOrganization(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOrganizationUnit(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateState(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OpenSSLCertificateAuthoritySpec) validateCommonName(formats strfmt.Registry) error {

	if err := validate.Required("commonName", "body", m.CommonName); err != nil {
		return err
	}

	return nil
}

func (m *OpenSSLCertificateAuthoritySpec) validateCountry(formats strfmt.Registry) error {

	if err := validate.Required("country", "body", m.Country); err != nil {
		return err
	}

	return nil
}

func (m *OpenSSLCertificateAuthoritySpec) validateLocality(formats strfmt.Registry) error {

	if err := validate.Required("locality", "body", m.Locality); err != nil {
		return err
	}

	return nil
}

func (m *OpenSSLCertificateAuthoritySpec) validateOrganization(formats strfmt.Registry) error {

	if err := validate.Required("organization", "body", m.Organization); err != nil {
		return err
	}

	return nil
}

func (m *OpenSSLCertificateAuthoritySpec) validateOrganizationUnit(formats strfmt.Registry) error {

	if err := validate.Required("organizationUnit", "body", m.OrganizationUnit); err != nil {
		return err
	}

	return nil
}

func (m *OpenSSLCertificateAuthoritySpec) validateState(formats strfmt.Registry) error {

	if err := validate.Required("state", "body", m.State); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this open SSL certificate authority spec based on context it is used
func (m *OpenSSLCertificateAuthoritySpec) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *OpenSSLCertificateAuthoritySpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OpenSSLCertificateAuthoritySpec) UnmarshalBinary(b []byte) error {
	var res OpenSSLCertificateAuthoritySpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
