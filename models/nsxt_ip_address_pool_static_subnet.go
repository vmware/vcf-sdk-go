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

// NSXTIPAddressPoolStaticSubnet NSX IP address pool static subnet representation
//
// swagger:model NsxtIpAddressPoolStaticSubnet
type NSXTIPAddressPoolStaticSubnet struct {

	// The CIDR of the IP address subnet
	Cidr string `json:"cidr,omitempty"`

	// The gateway IP address
	Gateway string `json:"gateway,omitempty"`

	// The list of IP address ranges
	IPAddressPoolRanges []*NSXTIPAddressPoolRange `json:"ipAddressPoolRanges"`
}

// Validate validates this Nsxt Ip address pool static subnet
func (m *NSXTIPAddressPoolStaticSubnet) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateIPAddressPoolRanges(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NSXTIPAddressPoolStaticSubnet) validateIPAddressPoolRanges(formats strfmt.Registry) error {
	if swag.IsZero(m.IPAddressPoolRanges) { // not required
		return nil
	}

	for i := 0; i < len(m.IPAddressPoolRanges); i++ {
		if swag.IsZero(m.IPAddressPoolRanges[i]) { // not required
			continue
		}

		if m.IPAddressPoolRanges[i] != nil {
			if err := m.IPAddressPoolRanges[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("ipAddressPoolRanges" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("ipAddressPoolRanges" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this Nsxt Ip address pool static subnet based on the context it is used
func (m *NSXTIPAddressPoolStaticSubnet) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateIPAddressPoolRanges(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NSXTIPAddressPoolStaticSubnet) contextValidateIPAddressPoolRanges(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.IPAddressPoolRanges); i++ {

		if m.IPAddressPoolRanges[i] != nil {

			if swag.IsZero(m.IPAddressPoolRanges[i]) { // not required
				return nil
			}

			if err := m.IPAddressPoolRanges[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("ipAddressPoolRanges" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("ipAddressPoolRanges" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *NSXTIPAddressPoolStaticSubnet) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NSXTIPAddressPoolStaticSubnet) UnmarshalBinary(b []byte) error {
	var res NSXTIPAddressPoolStaticSubnet
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
