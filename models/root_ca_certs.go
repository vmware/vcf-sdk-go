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

// RootCaCerts Spec contains Root Certificate Authority parameters
//
// swagger:model RootCaCerts
type RootCaCerts struct {

	// Certificate alias
	Alias string `json:"alias,omitempty"`

	// List of Base64 encoded certificates
	CertChain []string `json:"certChain"`
}

// Validate validates this root ca certs
func (m *RootCaCerts) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this root ca certs based on context it is used
func (m *RootCaCerts) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *RootCaCerts) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RootCaCerts) UnmarshalBinary(b []byte) error {
	var res RootCaCerts
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
