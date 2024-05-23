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
	"github.com/go-openapi/validate"
)

// SDDCNetworkSpec Defines a network spec
//
// swagger:model SddcNetworkSpec
type SDDCNetworkSpec struct {

	// Active Uplinks for teaming policy, specify uplink1 for failover_explicit VSAN Teaming Policy
	ActiveUplinks []string `json:"activeUplinks"`

	// IP Addresse ranges to be excluded
	ExcludeIPAddressRanges []string `json:"excludeIpAddressRanges"`

	// IP Addresses to be excluded
	ExcludeIpaddresses []string `json:"excludeIpaddresses"`

	// Gateway
	// Max Length: 15
	// Min Length: 7
	Gateway string `json:"gateway,omitempty"`

	// IP Addresses to be included
	IncludeIPAddress []string `json:"includeIpAddress"`

	// IP Addresse ranges to be included
	IncludeIPAddressRanges []*IPRange `json:"includeIpAddressRanges"`

	// MTU size
	// Max Length: 4
	// Min Length: 4
	Mtu string `json:"mtu,omitempty"`

	// Network Type
	// Example: One among: VSAN, VMOTION, MANAGEMENT, VM_MANAGEMENT or any custom network type
	// Required: true
	NetworkType *string `json:"networkType"`

	// Portgroup key name
	PortGroupKey string `json:"portGroupKey,omitempty"`

	// Standby Uplinks for teaming policy, specify uplink2 for failover_explicit VSAN Teaming Policy
	StandbyUplinks []string `json:"standbyUplinks"`

	// Subnet
	// Max Length: 15
	// Min Length: 7
	Subnet string `json:"subnet,omitempty"`

	// Subnet Mask
	// Max Length: 15
	// Min Length: 7
	SubnetMask string `json:"subnetMask,omitempty"`

	// Teaming Policy for VSAN and VMOTION network types, Default is loadbalance_loadbased
	// Example: One among:loadbalance_ip, loadbalance_srcmac, loadbalance_srcid, failover_explicit, loadbalance_loadbased
	TeamingPolicy string `json:"teamingPolicy,omitempty"`

	// VLAN Id
	// Required: true
	// Max Length: 4
	// Min Length: 1
	VlanID *string `json:"vlanId"`
}

// Validate validates this Sddc network spec
func (m *SDDCNetworkSpec) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateGateway(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIncludeIPAddressRanges(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMtu(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNetworkType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSubnet(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSubnetMask(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVlanID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SDDCNetworkSpec) validateGateway(formats strfmt.Registry) error {
	if swag.IsZero(m.Gateway) { // not required
		return nil
	}

	if err := validate.MinLength("gateway", "body", m.Gateway, 7); err != nil {
		return err
	}

	if err := validate.MaxLength("gateway", "body", m.Gateway, 15); err != nil {
		return err
	}

	return nil
}

func (m *SDDCNetworkSpec) validateIncludeIPAddressRanges(formats strfmt.Registry) error {
	if swag.IsZero(m.IncludeIPAddressRanges) { // not required
		return nil
	}

	for i := 0; i < len(m.IncludeIPAddressRanges); i++ {
		if swag.IsZero(m.IncludeIPAddressRanges[i]) { // not required
			continue
		}

		if m.IncludeIPAddressRanges[i] != nil {
			if err := m.IncludeIPAddressRanges[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("includeIpAddressRanges" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("includeIpAddressRanges" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *SDDCNetworkSpec) validateMtu(formats strfmt.Registry) error {
	if swag.IsZero(m.Mtu) { // not required
		return nil
	}

	if err := validate.MinLength("mtu", "body", m.Mtu, 4); err != nil {
		return err
	}

	if err := validate.MaxLength("mtu", "body", m.Mtu, 4); err != nil {
		return err
	}

	return nil
}

func (m *SDDCNetworkSpec) validateNetworkType(formats strfmt.Registry) error {

	if err := validate.Required("networkType", "body", m.NetworkType); err != nil {
		return err
	}

	return nil
}

func (m *SDDCNetworkSpec) validateSubnet(formats strfmt.Registry) error {
	if swag.IsZero(m.Subnet) { // not required
		return nil
	}

	if err := validate.MinLength("subnet", "body", m.Subnet, 7); err != nil {
		return err
	}

	if err := validate.MaxLength("subnet", "body", m.Subnet, 15); err != nil {
		return err
	}

	return nil
}

func (m *SDDCNetworkSpec) validateSubnetMask(formats strfmt.Registry) error {
	if swag.IsZero(m.SubnetMask) { // not required
		return nil
	}

	if err := validate.MinLength("subnetMask", "body", m.SubnetMask, 7); err != nil {
		return err
	}

	if err := validate.MaxLength("subnetMask", "body", m.SubnetMask, 15); err != nil {
		return err
	}

	return nil
}

func (m *SDDCNetworkSpec) validateVlanID(formats strfmt.Registry) error {

	if err := validate.Required("vlanId", "body", m.VlanID); err != nil {
		return err
	}

	if err := validate.MinLength("vlanId", "body", *m.VlanID, 1); err != nil {
		return err
	}

	if err := validate.MaxLength("vlanId", "body", *m.VlanID, 4); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this Sddc network spec based on the context it is used
func (m *SDDCNetworkSpec) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateIncludeIPAddressRanges(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SDDCNetworkSpec) contextValidateIncludeIPAddressRanges(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.IncludeIPAddressRanges); i++ {

		if m.IncludeIPAddressRanges[i] != nil {

			if swag.IsZero(m.IncludeIPAddressRanges[i]) { // not required
				return nil
			}

			if err := m.IncludeIPAddressRanges[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("includeIpAddressRanges" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("includeIpAddressRanges" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *SDDCNetworkSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SDDCNetworkSpec) UnmarshalBinary(b []byte) error {
	var res SDDCNetworkSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
