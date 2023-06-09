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

// DepotAccount VMware Depot Account Information
//
// swagger:model DepotAccount
type DepotAccount struct {

	// Message explaining depot status
	Message string `json:"message,omitempty"`

	// Depot Password for Access
	// Required: true
	Password *string `json:"password"`

	// Depot Status
	// Example: One among: DEPOT_UNKNOWN_HOST, DEPOT_NOT_AVAILABLE, DEPOT_USER_NOT_SET, DEPOT_INVALID_CREDENTIAL, UNKNOWN_FAILURE, DEPOT_CONNECTION_SUCCESSFUL
	Status string `json:"status,omitempty"`

	// Depot Username for Access
	// Required: true
	Username *string `json:"username"`
}

// Validate validates this depot account
func (m *DepotAccount) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePassword(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUsername(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DepotAccount) validatePassword(formats strfmt.Registry) error {

	if err := validate.Required("password", "body", m.Password); err != nil {
		return err
	}

	return nil
}

func (m *DepotAccount) validateUsername(formats strfmt.Registry) error {

	if err := validate.Required("username", "body", m.Username); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this depot account based on context it is used
func (m *DepotAccount) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DepotAccount) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DepotAccount) UnmarshalBinary(b []byte) error {
	var res DepotAccount
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
