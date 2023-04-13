// Code generated by go-swagger; DO NOT EDIT.

// Copyright © 2023 VMware, Inc. All Rights Reserved.
// Code licenced under: BSD-2

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// VasaProviderUpdateSpec Represents a VASA Provider update specification
//
// swagger:model VasaProviderUpdateSpec
type VasaProviderUpdateSpec struct {

	// Name of the VASA Provider
	Name string `json:"name,omitempty"`

	// URL of the VASA Provider
	URL string `json:"url,omitempty"`
}

// Validate validates this vasa provider update spec
func (m *VasaProviderUpdateSpec) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this vasa provider update spec based on context it is used
func (m *VasaProviderUpdateSpec) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *VasaProviderUpdateSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VasaProviderUpdateSpec) UnmarshalBinary(b []byte) error {
	var res VasaProviderUpdateSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
