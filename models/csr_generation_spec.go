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

// CSRGenerationSpec This specification contains Certificate Signing Request (CSR) details
//
// swagger:model CsrGenerationSpec
type CSRGenerationSpec struct {

	// ISO 3166 country code where company is legally registered
	// Example: IN
	// Required: true
	Country *string `json:"country"`

	// Contact email address
	// Example: admin@vmware.com
	Email string `json:"email,omitempty"`

	// The public key algorithm of the certificate
	// Example: One among: RSA
	// Required: true
	KeyAlgorithm *string `json:"keyAlgorithm"`

	// Certificate public key size
	// Example: One among: 2048, 3072, 4096
	// Required: true
	KeySize *string `json:"keySize"`

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

// Validate validates this Csr generation spec
func (m *CSRGenerationSpec) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCountry(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateKeyAlgorithm(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateKeySize(formats); err != nil {
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

func (m *CSRGenerationSpec) validateCountry(formats strfmt.Registry) error {

	if err := validate.Required("country", "body", m.Country); err != nil {
		return err
	}

	return nil
}

func (m *CSRGenerationSpec) validateKeyAlgorithm(formats strfmt.Registry) error {

	if err := validate.Required("keyAlgorithm", "body", m.KeyAlgorithm); err != nil {
		return err
	}

	return nil
}

func (m *CSRGenerationSpec) validateKeySize(formats strfmt.Registry) error {

	if err := validate.Required("keySize", "body", m.KeySize); err != nil {
		return err
	}

	return nil
}

func (m *CSRGenerationSpec) validateLocality(formats strfmt.Registry) error {

	if err := validate.Required("locality", "body", m.Locality); err != nil {
		return err
	}

	return nil
}

func (m *CSRGenerationSpec) validateOrganization(formats strfmt.Registry) error {

	if err := validate.Required("organization", "body", m.Organization); err != nil {
		return err
	}

	return nil
}

func (m *CSRGenerationSpec) validateOrganizationUnit(formats strfmt.Registry) error {

	if err := validate.Required("organizationUnit", "body", m.OrganizationUnit); err != nil {
		return err
	}

	return nil
}

func (m *CSRGenerationSpec) validateState(formats strfmt.Registry) error {

	if err := validate.Required("state", "body", m.State); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this Csr generation spec based on context it is used
func (m *CSRGenerationSpec) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CSRGenerationSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CSRGenerationSpec) UnmarshalBinary(b []byte) error {
	var res CSRGenerationSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
