// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2023 VMware, Inc.
// SPDX-License-Identifier: BSD-2-Clause

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Network Represents a network in a network pool
//
// swagger:model Network
type Network struct {

	// List of free IPs to use
	FreeIps []string `json:"freeIps"`

	// Gateway for the network
	Gateway string `json:"gateway,omitempty"`

	// The ID of the network
	ID string `json:"id,omitempty"`

	// List of IP pool ranges to use
	IPPools []*IPPool `json:"ipPools"`

	// Subnet mask for the subnet of the network
	Mask string `json:"mask,omitempty"`

	// MTU of the network
	Mtu int32 `json:"mtu,omitempty"`

	// Subnet associated with the network
	Subnet string `json:"subnet,omitempty"`

	// Network Type of the network
	Type string `json:"type,omitempty"`

	// List of used IPs
	UsedIps []string `json:"usedIps"`

	// VLAN ID associated with the network
	VlanID int32 `json:"vlanId"`
}

// Validate validates this network
func (m *Network) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateIPPools(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Network) validateIPPools(formats strfmt.Registry) error {
	if swag.IsZero(m.IPPools) { // not required
		return nil
	}

	for i := 0; i < len(m.IPPools); i++ {
		if swag.IsZero(m.IPPools[i]) { // not required
			continue
		}

		if m.IPPools[i] != nil {
			if err := m.IPPools[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("ipPools" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("ipPools" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this network based on the context it is used
func (m *Network) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateIPPools(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Network) contextValidateIPPools(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.IPPools); i++ {

		if m.IPPools[i] != nil {

			if swag.IsZero(m.IPPools[i]) { // not required
				return nil
			}

			if err := m.IPPools[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("ipPools" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("ipPools" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *Network) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Network) UnmarshalBinary(b []byte) error {
	var res Network
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
