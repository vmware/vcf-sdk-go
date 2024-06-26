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

// StretchClusterNetworkProfile This specification contains the host switch configurations for the NSX transport nodes belonging to Secondary AZ.
//
// swagger:model StretchClusterNetworkProfile
type StretchClusterNetworkProfile struct {

	// The network profile description
	Description string `json:"description,omitempty"`

	// The network profile name
	// Required: true
	Name *string `json:"name"`

	// The list of NSX host switch configurations
	// Required: true
	NSXTHostSwitchConfigs []*NSXTHostSwitchConfig `json:"nsxtHostSwitchConfigs"`
}

// Validate validates this stretch cluster network profile
func (m *StretchClusterNetworkProfile) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNSXTHostSwitchConfigs(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *StretchClusterNetworkProfile) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *StretchClusterNetworkProfile) validateNSXTHostSwitchConfigs(formats strfmt.Registry) error {

	if err := validate.Required("nsxtHostSwitchConfigs", "body", m.NSXTHostSwitchConfigs); err != nil {
		return err
	}

	for i := 0; i < len(m.NSXTHostSwitchConfigs); i++ {
		if swag.IsZero(m.NSXTHostSwitchConfigs[i]) { // not required
			continue
		}

		if m.NSXTHostSwitchConfigs[i] != nil {
			if err := m.NSXTHostSwitchConfigs[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("nsxtHostSwitchConfigs" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("nsxtHostSwitchConfigs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this stretch cluster network profile based on the context it is used
func (m *StretchClusterNetworkProfile) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateNSXTHostSwitchConfigs(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *StretchClusterNetworkProfile) contextValidateNSXTHostSwitchConfigs(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.NSXTHostSwitchConfigs); i++ {

		if m.NSXTHostSwitchConfigs[i] != nil {

			if swag.IsZero(m.NSXTHostSwitchConfigs[i]) { // not required
				return nil
			}

			if err := m.NSXTHostSwitchConfigs[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("nsxtHostSwitchConfigs" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("nsxtHostSwitchConfigs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *StretchClusterNetworkProfile) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StretchClusterNetworkProfile) UnmarshalBinary(b []byte) error {
	var res StretchClusterNetworkProfile
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
