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

// ConfigDriftSpec Represents a Configuration Drift workflow
//
// swagger:model ConfigDriftSpec
type ConfigDriftSpec struct {

	// Applicability of the Configuration Drift
	Applicability *Applicability `json:"applicability,omitempty"`

	// Description of the Configuration Drift
	// Example: This drift operation will create an additional distributed virtual port group on a cluster for VCF management VM communication, and migrate VCF-managed VMs to this new port group
	Description string `json:"description,omitempty"`

	// The UUID of the Configuration Drift
	// Example: 3f39d4a1-78d2-11e8-af85-f1cf26258cdc
	ID string `json:"id,omitempty"`

	// Name of the Configuration Drift
	// Example: Distributed Virtual Portgroup configuration drift
	Name string `json:"name,omitempty"`

	// The infrastructure resource on which this Configuration Drift operates
	// Example: system | domain | cluster | host | edge_cluster
	ResourceType string `json:"resourceType,omitempty"`

	// The type of Configuration Drift
	// Example: fix | feature
	Type string `json:"type,omitempty"`
}

// Validate validates this config drift spec
func (m *ConfigDriftSpec) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateApplicability(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConfigDriftSpec) validateApplicability(formats strfmt.Registry) error {
	if swag.IsZero(m.Applicability) { // not required
		return nil
	}

	if m.Applicability != nil {
		if err := m.Applicability.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("applicability")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("applicability")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this config drift spec based on the context it is used
func (m *ConfigDriftSpec) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateApplicability(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConfigDriftSpec) contextValidateApplicability(ctx context.Context, formats strfmt.Registry) error {

	if m.Applicability != nil {

		if swag.IsZero(m.Applicability) { // not required
			return nil
		}

		if err := m.Applicability.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("applicability")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("applicability")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ConfigDriftSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConfigDriftSpec) UnmarshalBinary(b []byte) error {
	var res ConfigDriftSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}