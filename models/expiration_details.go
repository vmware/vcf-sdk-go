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

// ExpirationDetails Represents expiration details of the credential
//
// swagger:model ExpirationDetails
type ExpirationDetails struct {

	// Connectivity status
	// Example: One among: ACTIVE, ERROR, UNKNOWN
	ConnectivityStatus string `json:"connectivityStatus,omitempty"`

	// Expiration date
	ExpiryDate string `json:"expiryDate,omitempty"`

	// Last checked date
	LastCheckedDate string `json:"lastCheckedDate,omitempty"`

	// Expiration status
	// Example: One among: ACTIVE, EXPIRING, EXPIRED, UNKNOWN
	Status string `json:"status,omitempty"`
}

// Validate validates this expiration details
func (m *ExpirationDetails) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this expiration details based on context it is used
func (m *ExpirationDetails) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ExpirationDetails) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ExpirationDetails) UnmarshalBinary(b []byte) error {
	var res ExpirationDetails
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
