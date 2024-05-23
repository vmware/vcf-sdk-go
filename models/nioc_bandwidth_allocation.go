// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2023 VMware, Inc.
// SPDX-License-Identifier: BSD-2-Clause

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NiocBandwidthAllocation Network traffic per resource type
//
// swagger:model NiocBandwidthAllocation
type NiocBandwidthAllocation struct {

	// Specify network traffic allocation for a resource
	NiocTrafficResourceAllocation *NiocTrafficResourceAllocation `json:"niocTrafficResourceAllocation,omitempty"`

	// Host infrastructure traffic type
	// Example: Example: management, faultTolerance, vmotion, virtualMachine, iSCSI, nfs, hbr, vsan, vdp etc.
	Type string `json:"type,omitempty"`
}

// Validate validates this nioc bandwidth allocation
func (m *NiocBandwidthAllocation) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateNiocTrafficResourceAllocation(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NiocBandwidthAllocation) validateNiocTrafficResourceAllocation(formats strfmt.Registry) error {
	if swag.IsZero(m.NiocTrafficResourceAllocation) { // not required
		return nil
	}

	if m.NiocTrafficResourceAllocation != nil {
		if err := m.NiocTrafficResourceAllocation.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("niocTrafficResourceAllocation")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("niocTrafficResourceAllocation")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this nioc bandwidth allocation based on the context it is used
func (m *NiocBandwidthAllocation) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateNiocTrafficResourceAllocation(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NiocBandwidthAllocation) contextValidateNiocTrafficResourceAllocation(ctx context.Context, formats strfmt.Registry) error {

	if m.NiocTrafficResourceAllocation != nil {

		if swag.IsZero(m.NiocTrafficResourceAllocation) { // not required
			return nil
		}

		if err := m.NiocTrafficResourceAllocation.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("niocTrafficResourceAllocation")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("niocTrafficResourceAllocation")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *NiocBandwidthAllocation) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NiocBandwidthAllocation) UnmarshalBinary(b []byte) error {
	var res NiocBandwidthAllocation
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
