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

// NSXTSwitchConfig This specification contains the configurations to be associated with the vSphere Distributed Switch managed by NSX
//
// swagger:model NsxtSwitchConfig
type NSXTSwitchConfig struct {

	// vSphere Distributed Switch name
	// Example: One among: STANDARD, ENS, ENS_INTERRUPT
	HostSwitchOperationalMode string `json:"hostSwitchOperationalMode,omitempty"`

	// The list of transport zones to be associated with the vSphere Distributed Switch managed by NSX
	TransportZones []*TransportZone `json:"transportZones"`
}

// Validate validates this Nsxt switch config
func (m *NSXTSwitchConfig) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTransportZones(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NSXTSwitchConfig) validateTransportZones(formats strfmt.Registry) error {
	if swag.IsZero(m.TransportZones) { // not required
		return nil
	}

	for i := 0; i < len(m.TransportZones); i++ {
		if swag.IsZero(m.TransportZones[i]) { // not required
			continue
		}

		if m.TransportZones[i] != nil {
			if err := m.TransportZones[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("transportZones" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("transportZones" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this Nsxt switch config based on the context it is used
func (m *NSXTSwitchConfig) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateTransportZones(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NSXTSwitchConfig) contextValidateTransportZones(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.TransportZones); i++ {

		if m.TransportZones[i] != nil {

			if swag.IsZero(m.TransportZones[i]) { // not required
				return nil
			}

			if err := m.TransportZones[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("transportZones" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("transportZones" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *NSXTSwitchConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NSXTSwitchConfig) UnmarshalBinary(b []byte) error {
	var res NSXTSwitchConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}