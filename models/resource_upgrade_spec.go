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

// ResourceUpgradeSpec Resource Upgrade Specification
//
// swagger:model ResourceUpgradeSpec
type ResourceUpgradeSpec struct {

	// Custom ISO Specifications for VUM Clusters Upgrade with Custom ISO
	CustomISOSpec *CustomISOSpec `json:"customISOSpec,omitempty"`

	// Flag for requesting Quick Boot for ESXi upgrade
	EnableQuickboot bool `json:"enableQuickboot,omitempty"`

	// Flag for requesting Evacuation of Offline VMs for ESXi upgrade
	EvacuateOfflineVms bool `json:"evacuateOfflineVms,omitempty"`

	// Personality Specifications for vLCM based upgrade
	PersonalitySpec *PersonalitySpec `json:"personalitySpec,omitempty"`

	// Resource ID for Upgrade
	// Required: true
	ResourceID *string `json:"resourceId"`

	// Upgrade Scheduled Time
	ScheduledTimestamp string `json:"scheduledTimestamp,omitempty"`

	// If Resource Type is UNASSIGNED_HOST, set flag for enabling shutting down VM's before Upgrade
	ShutdownVms bool `json:"shutdownVms,omitempty"`

	// If Resource Type is UNASSIGNED_HOST, set the target version for Upgrade
	ToVersion string `json:"toVersion,omitempty"`

	// Flag for enabling Upgrade Now. If true, scheduledTimestamp is ignored
	UpgradeNow bool `json:"upgradeNow,omitempty"`
}

// Validate validates this resource upgrade spec
func (m *ResourceUpgradeSpec) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCustomISOSpec(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePersonalitySpec(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResourceID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ResourceUpgradeSpec) validateCustomISOSpec(formats strfmt.Registry) error {
	if swag.IsZero(m.CustomISOSpec) { // not required
		return nil
	}

	if m.CustomISOSpec != nil {
		if err := m.CustomISOSpec.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("customISOSpec")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("customISOSpec")
			}
			return err
		}
	}

	return nil
}

func (m *ResourceUpgradeSpec) validatePersonalitySpec(formats strfmt.Registry) error {
	if swag.IsZero(m.PersonalitySpec) { // not required
		return nil
	}

	if m.PersonalitySpec != nil {
		if err := m.PersonalitySpec.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("personalitySpec")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("personalitySpec")
			}
			return err
		}
	}

	return nil
}

func (m *ResourceUpgradeSpec) validateResourceID(formats strfmt.Registry) error {

	if err := validate.Required("resourceId", "body", m.ResourceID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this resource upgrade spec based on the context it is used
func (m *ResourceUpgradeSpec) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCustomISOSpec(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePersonalitySpec(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ResourceUpgradeSpec) contextValidateCustomISOSpec(ctx context.Context, formats strfmt.Registry) error {

	if m.CustomISOSpec != nil {

		if swag.IsZero(m.CustomISOSpec) { // not required
			return nil
		}

		if err := m.CustomISOSpec.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("customISOSpec")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("customISOSpec")
			}
			return err
		}
	}

	return nil
}

func (m *ResourceUpgradeSpec) contextValidatePersonalitySpec(ctx context.Context, formats strfmt.Registry) error {

	if m.PersonalitySpec != nil {

		if swag.IsZero(m.PersonalitySpec) { // not required
			return nil
		}

		if err := m.PersonalitySpec.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("personalitySpec")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("personalitySpec")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ResourceUpgradeSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ResourceUpgradeSpec) UnmarshalBinary(b []byte) error {
	var res ResourceUpgradeSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
