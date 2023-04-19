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

// OidcSpec Identity management configuration when the Identity Provider is based on oidc
//
// swagger:model OidcSpec
type OidcSpec struct {

	// Client identifier to connect to the provider
	// Required: true
	ClientID *string `json:"clientId"`

	// The secret shared between the client and the provider
	// Required: true
	ClientSecret *string `json:"clientSecret"`

	// Endpoint to retrieve the provider metadata
	// Required: true
	DiscoveryEndpoint *string `json:"discoveryEndpoint"`
}

// Validate validates this oidc spec
func (m *OidcSpec) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateClientID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateClientSecret(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDiscoveryEndpoint(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OidcSpec) validateClientID(formats strfmt.Registry) error {

	if err := validate.Required("clientId", "body", m.ClientID); err != nil {
		return err
	}

	return nil
}

func (m *OidcSpec) validateClientSecret(formats strfmt.Registry) error {

	if err := validate.Required("clientSecret", "body", m.ClientSecret); err != nil {
		return err
	}

	return nil
}

func (m *OidcSpec) validateDiscoveryEndpoint(formats strfmt.Registry) error {

	if err := validate.Required("discoveryEndpoint", "body", m.DiscoveryEndpoint); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this oidc spec based on context it is used
func (m *OidcSpec) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *OidcSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OidcSpec) UnmarshalBinary(b []byte) error {
	var res OidcSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
