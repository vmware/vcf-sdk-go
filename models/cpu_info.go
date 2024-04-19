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

// CPUInfo Represents a cpu metric with used and total for a resource
//
// swagger:model CpuInfo
type CPUInfo struct {

	// Total value
	Total *FrequencyMetric `json:"total,omitempty"`

	// Used value
	Used *FrequencyMetric `json:"used,omitempty"`
}

// Validate validates this Cpu info
func (m *CPUInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTotal(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUsed(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CPUInfo) validateTotal(formats strfmt.Registry) error {
	if swag.IsZero(m.Total) { // not required
		return nil
	}

	if m.Total != nil {
		if err := m.Total.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("total")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("total")
			}
			return err
		}
	}

	return nil
}

func (m *CPUInfo) validateUsed(formats strfmt.Registry) error {
	if swag.IsZero(m.Used) { // not required
		return nil
	}

	if m.Used != nil {
		if err := m.Used.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("used")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("used")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this Cpu info based on the context it is used
func (m *CPUInfo) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateTotal(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUsed(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CPUInfo) contextValidateTotal(ctx context.Context, formats strfmt.Registry) error {

	if m.Total != nil {

		if swag.IsZero(m.Total) { // not required
			return nil
		}

		if err := m.Total.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("total")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("total")
			}
			return err
		}
	}

	return nil
}

func (m *CPUInfo) contextValidateUsed(ctx context.Context, formats strfmt.Registry) error {

	if m.Used != nil {

		if swag.IsZero(m.Used) { // not required
			return nil
		}

		if err := m.Used.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("used")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("used")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CPUInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CPUInfo) UnmarshalBinary(b []byte) error {
	var res CPUInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
