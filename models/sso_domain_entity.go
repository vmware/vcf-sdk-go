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

// SSODomainEntity Represents SSO domain entity
//
// swagger:model SsoDomainEntity
type SSODomainEntity struct {

	// The ID of the SSO domain entity
	ID string `json:"id,omitempty"`

	// The name of the SSO domain entity
	Name string `json:"name,omitempty"`

	// The type of the SSO domain entity
	// Example: One among: USER, GROUP
	// Required: true
	Type *string `json:"type"`
}

// Validate validates this Sso domain entity
func (m *SSODomainEntity) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SSODomainEntity) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this Sso domain entity based on context it is used
func (m *SSODomainEntity) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SSODomainEntity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SSODomainEntity) UnmarshalBinary(b []byte) error {
	var res SSODomainEntity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
