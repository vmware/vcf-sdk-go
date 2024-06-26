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
	"github.com/go-openapi/validate"
)

// NetworkSeparationSpec This specification contains the parameters required to provide network separation for the Isolated WLD
//
// swagger:model NetworkSeparationSpec
type NetworkSeparationSpec struct {

	// enable/disable distributed firewall rules for the Isolated WLD
	EnableSecurity bool `json:"enableSecurity,omitempty"`

	// The NSX segment specification
	// Required: true
	SegmentSpec *SegmentSpec `json:"segmentSpec"`
}

// Validate validates this network separation spec
func (m *NetworkSeparationSpec) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSegmentSpec(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NetworkSeparationSpec) validateSegmentSpec(formats strfmt.Registry) error {

	if err := validate.Required("segmentSpec", "body", m.SegmentSpec); err != nil {
		return err
	}

	if m.SegmentSpec != nil {
		if err := m.SegmentSpec.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("segmentSpec")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("segmentSpec")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this network separation spec based on the context it is used
func (m *NetworkSeparationSpec) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSegmentSpec(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NetworkSeparationSpec) contextValidateSegmentSpec(ctx context.Context, formats strfmt.Registry) error {

	if m.SegmentSpec != nil {

		if err := m.SegmentSpec.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("segmentSpec")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("segmentSpec")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *NetworkSeparationSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NetworkSeparationSpec) UnmarshalBinary(b []byte) error {
	var res NetworkSeparationSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
