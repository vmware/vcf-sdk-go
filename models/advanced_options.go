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

// AdvancedOptions Advanced Options used to add Cluster
//
// swagger:model AdvancedOptions
type AdvancedOptions struct {

	// EVC mode for new cluster, if needed
	// Example: One among: INTEL_MEROM, INTEL_PENRYN, INTEL_NEALEM, INTEL_WESTMERE, INTEL_SANDYBRIDGE, INTEL_IVYBRIDGE, INTEL_HASWELL, INTEL_BROADWELL, INTEL_SKYLAKE, INTEL_CASCADELAKE, AMD_REV_E, AMD_REV_F, AMD_GREYHOUND_NO3DNOW, AMD_GREYHOUND, AMD_BULLDOZER, AMD_PILEDRIVER, AMD_STREAMROLLER, AMD_ZEN
	EvcMode string `json:"evcMode,omitempty"`

	// High availability settings for the cluster
	HighAvailability *HighAvailability `json:"highAvailability,omitempty"`
}

// Validate validates this advanced options
func (m *AdvancedOptions) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateHighAvailability(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AdvancedOptions) validateHighAvailability(formats strfmt.Registry) error {
	if swag.IsZero(m.HighAvailability) { // not required
		return nil
	}

	if m.HighAvailability != nil {
		if err := m.HighAvailability.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("highAvailability")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("highAvailability")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this advanced options based on the context it is used
func (m *AdvancedOptions) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateHighAvailability(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AdvancedOptions) contextValidateHighAvailability(ctx context.Context, formats strfmt.Registry) error {

	if m.HighAvailability != nil {

		if swag.IsZero(m.HighAvailability) { // not required
			return nil
		}

		if err := m.HighAvailability.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("highAvailability")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("highAvailability")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AdvancedOptions) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AdvancedOptions) UnmarshalBinary(b []byte) error {
	var res AdvancedOptions
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
