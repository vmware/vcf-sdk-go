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

// BaseCredential Represents credentials of a resource in the system
//
// swagger:model BaseCredential
type BaseCredential struct {

	// Account type
	// Example: One among: USER, SYSTEM, SERVICE
	AccountType string `json:"accountType,omitempty"`

	// Credential type
	// Example: One among: SSO, SSH, API, FTP, AUDIT
	CredentialType string `json:"credentialType,omitempty"`

	// Password
	Password string `json:"password,omitempty"`

	// Username
	// Required: true
	Username *string `json:"username"`
}

// Validate validates this base credential
func (m *BaseCredential) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateUsername(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BaseCredential) validateUsername(formats strfmt.Registry) error {

	if err := validate.Required("username", "body", m.Username); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this base credential based on context it is used
func (m *BaseCredential) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *BaseCredential) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BaseCredential) UnmarshalBinary(b []byte) error {
	var res BaseCredential
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
