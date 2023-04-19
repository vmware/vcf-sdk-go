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

// NsxTEdgeClusterProfileSpec This specification contains edge cluster profile configurations
//
// swagger:model NsxTEdgeClusterProfileSpec
type NsxTEdgeClusterProfileSpec struct {

	// BFD allowed Hop
	// Required: true
	BfdAllowedHop *int64 `json:"bfdAllowedHop"`

	// BFD Declare Dead Multiple
	// Required: true
	BfdDeclareDeadMultiple *int64 `json:"bfdDeclareDeadMultiple"`

	// BFD Probe
	// Required: true
	BfdProbeInterval *int64 `json:"bfdProbeInterval"`

	// Name for the edge cluster profile.
	// Required: true
	EdgeClusterProfileName *string `json:"edgeClusterProfileName"`

	// Standby Relocation Threshold
	// Required: true
	StandbyRelocationThreshold *int64 `json:"standbyRelocationThreshold"`
}

// Validate validates this nsx t edge cluster profile spec
func (m *NsxTEdgeClusterProfileSpec) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBfdAllowedHop(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBfdDeclareDeadMultiple(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBfdProbeInterval(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEdgeClusterProfileName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStandbyRelocationThreshold(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NsxTEdgeClusterProfileSpec) validateBfdAllowedHop(formats strfmt.Registry) error {

	if err := validate.Required("bfdAllowedHop", "body", m.BfdAllowedHop); err != nil {
		return err
	}

	return nil
}

func (m *NsxTEdgeClusterProfileSpec) validateBfdDeclareDeadMultiple(formats strfmt.Registry) error {

	if err := validate.Required("bfdDeclareDeadMultiple", "body", m.BfdDeclareDeadMultiple); err != nil {
		return err
	}

	return nil
}

func (m *NsxTEdgeClusterProfileSpec) validateBfdProbeInterval(formats strfmt.Registry) error {

	if err := validate.Required("bfdProbeInterval", "body", m.BfdProbeInterval); err != nil {
		return err
	}

	return nil
}

func (m *NsxTEdgeClusterProfileSpec) validateEdgeClusterProfileName(formats strfmt.Registry) error {

	if err := validate.Required("edgeClusterProfileName", "body", m.EdgeClusterProfileName); err != nil {
		return err
	}

	return nil
}

func (m *NsxTEdgeClusterProfileSpec) validateStandbyRelocationThreshold(formats strfmt.Registry) error {

	if err := validate.Required("standbyRelocationThreshold", "body", m.StandbyRelocationThreshold); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this nsx t edge cluster profile spec based on context it is used
func (m *NsxTEdgeClusterProfileSpec) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *NsxTEdgeClusterProfileSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NsxTEdgeClusterProfileSpec) UnmarshalBinary(b []byte) error {
	var res NsxTEdgeClusterProfileSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
