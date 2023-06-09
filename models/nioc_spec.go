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

// NiocSpec Spec contains parameters for Network IO Control
//
// swagger:model NiocSpec
type NiocSpec struct {

	// Traffic Type
	// Example: One among:VSAN, VMOTION, VIRTUALMACHINE, MANAGEMENT, NFS, VDP, HBR, FAULTTOLERANCE, ISCSI
	// Required: true
	TrafficType *string `json:"trafficType"`

	// NIOC Value
	// Required: true
	Value *string `json:"value"`
}

// Validate validates this nioc spec
func (m *NiocSpec) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTrafficType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValue(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NiocSpec) validateTrafficType(formats strfmt.Registry) error {

	if err := validate.Required("trafficType", "body", m.TrafficType); err != nil {
		return err
	}

	return nil
}

func (m *NiocSpec) validateValue(formats strfmt.Registry) error {

	if err := validate.Required("value", "body", m.Value); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this nioc spec based on context it is used
func (m *NiocSpec) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *NiocSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NiocSpec) UnmarshalBinary(b []byte) error {
	var res NiocSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
