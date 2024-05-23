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

// NsxTEdgeUplinkNetwork This specification contains input  parameters required to configure  uplink network for NSX  edge node VM
//
// swagger:model NsxTEdgeUplinkNetwork
type NsxTEdgeUplinkNetwork struct {

	// [Deprecated] ASN of Peer (please use bgpPeers instead)
	AsnPeer int64 `json:"asnPeer,omitempty"`

	// [Deprecated] BGP Peer Password (please use bgpPeers instead)
	BgpPeerPassword string `json:"bgpPeerPassword,omitempty"`

	// List of BGP Peer configuration
	BgpPeers []*BgpPeerSpec `json:"bgpPeers"`

	// [Deprecated] BGP Peer IP (please use bgpPeers instead)
	PeerIP string `json:"peerIP,omitempty"`

	// Uplink IP
	// Required: true
	UplinkInterfaceIP *string `json:"uplinkInterfaceIP"`

	// Uplink Vlan
	// Required: true
	UplinkVlan *int32 `json:"uplinkVlan"`
}

// Validate validates this nsx t edge uplink network
func (m *NsxTEdgeUplinkNetwork) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBgpPeers(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUplinkInterfaceIP(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUplinkVlan(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NsxTEdgeUplinkNetwork) validateBgpPeers(formats strfmt.Registry) error {
	if swag.IsZero(m.BgpPeers) { // not required
		return nil
	}

	for i := 0; i < len(m.BgpPeers); i++ {
		if swag.IsZero(m.BgpPeers[i]) { // not required
			continue
		}

		if m.BgpPeers[i] != nil {
			if err := m.BgpPeers[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("bgpPeers" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("bgpPeers" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *NsxTEdgeUplinkNetwork) validateUplinkInterfaceIP(formats strfmt.Registry) error {

	if err := validate.Required("uplinkInterfaceIP", "body", m.UplinkInterfaceIP); err != nil {
		return err
	}

	return nil
}

func (m *NsxTEdgeUplinkNetwork) validateUplinkVlan(formats strfmt.Registry) error {

	if err := validate.Required("uplinkVlan", "body", m.UplinkVlan); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this nsx t edge uplink network based on the context it is used
func (m *NsxTEdgeUplinkNetwork) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateBgpPeers(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NsxTEdgeUplinkNetwork) contextValidateBgpPeers(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.BgpPeers); i++ {

		if m.BgpPeers[i] != nil {

			if swag.IsZero(m.BgpPeers[i]) { // not required
				return nil
			}

			if err := m.BgpPeers[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("bgpPeers" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("bgpPeers" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *NsxTEdgeUplinkNetwork) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NsxTEdgeUplinkNetwork) UnmarshalBinary(b []byte) error {
	var res NsxTEdgeUplinkNetwork
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
