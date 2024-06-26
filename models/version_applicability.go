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

// VersionApplicability Represents the version applicability of a Configuration Drift for a product
//
// swagger:model VersionApplicability
type VersionApplicability struct {

	// The maximum version for current product to apply current Configuration Drift
	MaxVersion string `json:"maxVersion,omitempty"`

	// The minimum version for current product to apply current Configuration Drift
	MinVersion string `json:"minVersion,omitempty"`

	// Product type
	// Example: SDDC_MANAGER|VCENTER|ESXI|NSXT_MANAGER|VXRAIL
	ProductType string `json:"productType,omitempty"`
}

// Validate validates this version applicability
func (m *VersionApplicability) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this version applicability based on context it is used
func (m *VersionApplicability) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *VersionApplicability) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VersionApplicability) UnmarshalBinary(b []byte) error {
	var res VersionApplicability
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
