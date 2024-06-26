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

// IdentitySourceSpec Represents Identity Source specification
//
// swagger:model IdentitySourceSpec
type IdentitySourceSpec struct {

	// The LDAP specification when the protocol is LDAP
	// Required: true
	Ldap *LdapSpec `json:"ldap"`

	// The user-friendly name for the identity Source
	// Required: true
	Name *string `json:"name"`
}

// Validate validates this identity source spec
func (m *IdentitySourceSpec) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLdap(formats); err != nil {
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

func (m *IdentitySourceSpec) validateLdap(formats strfmt.Registry) error {

	if err := validate.Required("ldap", "body", m.Ldap); err != nil {
		return err
	}

	if m.Ldap != nil {
		if err := m.Ldap.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ldap")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("ldap")
			}
			return err
		}
	}

	return nil
}

func (m *IdentitySourceSpec) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this identity source spec based on the context it is used
func (m *IdentitySourceSpec) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLdap(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IdentitySourceSpec) contextValidateLdap(ctx context.Context, formats strfmt.Registry) error {

	if m.Ldap != nil {

		if err := m.Ldap.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ldap")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("ldap")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IdentitySourceSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IdentitySourceSpec) UnmarshalBinary(b []byte) error {
	var res IdentitySourceSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
