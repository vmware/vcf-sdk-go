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

// IdentityProviderDirectory The directory configuration when the provider is via broker federation
//
// swagger:model IdentityProviderDirectory
type IdentityProviderDirectory struct {

	// The trusted default domain of the directory
	// Required: true
	DefaultDomain *string `json:"defaultDomain"`

	// The id of the directory
	DirectoryID string `json:"directoryId,omitempty"`

	// The set of trusted domains of the directory
	// Required: true
	Domains []string `json:"domains"`

	// The user-friendly name for the directory
	// Required: true
	Name *string `json:"name"`
}

// Validate validates this identity provider directory
func (m *IdentityProviderDirectory) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDefaultDomain(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDomains(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IdentityProviderDirectory) validateDefaultDomain(formats strfmt.Registry) error {

	if err := validate.Required("defaultDomain", "body", m.DefaultDomain); err != nil {
		return err
	}

	return nil
}

func (m *IdentityProviderDirectory) validateDomains(formats strfmt.Registry) error {

	if err := validate.Required("domains", "body", m.Domains); err != nil {
		return err
	}

	return nil
}

func (m *IdentityProviderDirectory) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this identity provider directory based on context it is used
func (m *IdentityProviderDirectory) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *IdentityProviderDirectory) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IdentityProviderDirectory) UnmarshalBinary(b []byte) error {
	var res IdentityProviderDirectory
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
