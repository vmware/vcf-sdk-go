// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2023 VMware, Inc.
// SPDX-License-Identifier: BSD-2-Clause

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// RefreshToken This contains refresh token id for the user API access.
//
// swagger:model RefreshToken
type RefreshToken struct {

	// Refresh token id that can be used to request new access token
	ID string `json:"id,omitempty"`
}

// Validate validates this refresh token
func (m *RefreshToken) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this refresh token based on context it is used
func (m *RefreshToken) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *RefreshToken) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RefreshToken) UnmarshalBinary(b []byte) error {
	var res RefreshToken
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
