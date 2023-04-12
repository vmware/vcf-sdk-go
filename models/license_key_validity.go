// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// LicenseKeyValidity Defines license key validity details
//
// swagger:model LicenseKeyValidity
type LicenseKeyValidity struct {

	// The license key expiry date
	ExpiryDate string `json:"expiryDate,omitempty"`

	// The validity status of the license key
	// Example: One among: EXPIRED, ACTIVE, NEVER_EXPIRES
	LicenseKeyStatus string `json:"licenseKeyStatus,omitempty"`
}

// Validate validates this license key validity
func (m *LicenseKeyValidity) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this license key validity based on context it is used
func (m *LicenseKeyValidity) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LicenseKeyValidity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LicenseKeyValidity) UnmarshalBinary(b []byte) error {
	var res LicenseKeyValidity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}