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

// IdentityProviderSpec Represents Identity provider configuration
//
// swagger:model IdentityProviderSpec
type IdentityProviderSpec struct {

	// The root certificate chain required to connect to the external server
	CertChain []string `json:"certChain"`

	// The identity management configuration when the provider is via broker federation
	FedIdpSpec *FederatedIdentityProviderSpec `json:"fedIdpSpec,omitempty"`

	// The LDAP specification when the protocol is LDAP
	Ldap *LdapSpec `json:"ldap,omitempty"`

	// The user-friendly name for the Identity Provider
	// Required: true
	Name *string `json:"name"`

	// The identity management configuration when the provider is based on oidc
	Oidc *OidcSpec `json:"oidc,omitempty"`

	// The type of Identity Identity Provider
	// Required: true
	Type *string `json:"type"`
}

// Validate validates this identity provider spec
func (m *IdentityProviderSpec) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFedIdpSpec(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLdap(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOidc(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IdentityProviderSpec) validateFedIdpSpec(formats strfmt.Registry) error {
	if swag.IsZero(m.FedIdpSpec) { // not required
		return nil
	}

	if m.FedIdpSpec != nil {
		if err := m.FedIdpSpec.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("fedIdpSpec")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("fedIdpSpec")
			}
			return err
		}
	}

	return nil
}

func (m *IdentityProviderSpec) validateLdap(formats strfmt.Registry) error {
	if swag.IsZero(m.Ldap) { // not required
		return nil
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

func (m *IdentityProviderSpec) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *IdentityProviderSpec) validateOidc(formats strfmt.Registry) error {
	if swag.IsZero(m.Oidc) { // not required
		return nil
	}

	if m.Oidc != nil {
		if err := m.Oidc.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("oidc")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("oidc")
			}
			return err
		}
	}

	return nil
}

func (m *IdentityProviderSpec) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this identity provider spec based on the context it is used
func (m *IdentityProviderSpec) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateFedIdpSpec(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLdap(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateOidc(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IdentityProviderSpec) contextValidateFedIdpSpec(ctx context.Context, formats strfmt.Registry) error {

	if m.FedIdpSpec != nil {

		if swag.IsZero(m.FedIdpSpec) { // not required
			return nil
		}

		if err := m.FedIdpSpec.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("fedIdpSpec")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("fedIdpSpec")
			}
			return err
		}
	}

	return nil
}

func (m *IdentityProviderSpec) contextValidateLdap(ctx context.Context, formats strfmt.Registry) error {

	if m.Ldap != nil {

		if swag.IsZero(m.Ldap) { // not required
			return nil
		}

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

func (m *IdentityProviderSpec) contextValidateOidc(ctx context.Context, formats strfmt.Registry) error {

	if m.Oidc != nil {

		if swag.IsZero(m.Oidc) { // not required
			return nil
		}

		if err := m.Oidc.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("oidc")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("oidc")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IdentityProviderSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IdentityProviderSpec) UnmarshalBinary(b []byte) error {
	var res IdentityProviderSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
