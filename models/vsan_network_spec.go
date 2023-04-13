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

// VSANNetworkSpec Contains the vSAN Network details
//
// swagger:model VSANNetworkSpec
type VSANNetworkSpec struct {

	// vSAN subnet cidr of the ESXi host
	VSANCidr string `json:"vsanCidr,omitempty"`

	// vSAN Gateway IP of the ESXi host
	VSANGatewayIP string `json:"vsanGatewayIP,omitempty"`
}

// Validate validates this VSAN network spec
func (m *VSANNetworkSpec) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this VSAN network spec based on context it is used
func (m *VSANNetworkSpec) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *VSANNetworkSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VSANNetworkSpec) UnmarshalBinary(b []byte) error {
	var res VSANNetworkSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
