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

// NSXTIPAddressPoolBlockSubnet NSX-T IP address pool block subnet representation
//
// swagger:model NsxtIpAddressPoolBlockSubnet
type NSXTIPAddressPoolBlockSubnet struct {

	// The CIDR of the IP address subnet
	Cidr string `json:"cidr,omitempty"`

	// The boolean representing the state of the gateway assignment
	GatewayAssigned bool `json:"gatewayAssigned,omitempty"`

	// The size of the IP address block
	Size int32 `json:"size,omitempty"`
}

// Validate validates this Nsxt Ip address pool block subnet
func (m *NSXTIPAddressPoolBlockSubnet) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this Nsxt Ip address pool block subnet based on context it is used
func (m *NSXTIPAddressPoolBlockSubnet) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *NSXTIPAddressPoolBlockSubnet) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NSXTIPAddressPoolBlockSubnet) UnmarshalBinary(b []byte) error {
	var res NSXTIPAddressPoolBlockSubnet
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
