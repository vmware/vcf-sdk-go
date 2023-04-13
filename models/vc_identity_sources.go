// Code generated by go-swagger; DO NOT EDIT.

// Copyright © 2023 VMware, Inc. All Rights Reserved.
// Code licenced under: BSD-2

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// VcIdentitySources Represents embedded Identity Sources and its attributes.
//
// swagger:model VcIdentitySources
type VcIdentitySources struct {

	// The Domains of the Identity Source
	DomainNames []string `json:"domainNames"`

	// LDAP configuration details of the Identity Source
	Ldap *LdapInfo `json:"ldap,omitempty"`

	// Name of the Identity Source.
	Name string `json:"name,omitempty"`

	// Type of the Identity Source.
	Type string `json:"type,omitempty"`
}

// Validate validates this vc identity sources
func (m *VcIdentitySources) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLdap(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VcIdentitySources) validateLdap(formats strfmt.Registry) error {
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

// ContextValidate validate this vc identity sources based on the context it is used
func (m *VcIdentitySources) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLdap(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VcIdentitySources) contextValidateLdap(ctx context.Context, formats strfmt.Registry) error {

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
func (m *VcIdentitySources) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VcIdentitySources) UnmarshalBinary(b []byte) error {
	var res VcIdentitySources
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
