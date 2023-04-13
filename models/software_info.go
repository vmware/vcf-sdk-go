// Code generated by go-swagger; DO NOT EDIT.

// Copyright © 2023 VMware, Inc. All Rights Reserved.
// Code licenced under: BSD-2

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

// SoftwareInfo SoftwareSpec contains base os, components, hardware support, addOn metadata to install/update the appropriate Cloud Foundation software components in your management domain or workload domain.
//
// swagger:model SoftwareInfo
type SoftwareInfo struct {

	// Personality addOn
	AddOn *AddOnInfo `json:"addOn,omitempty"`

	// Personality base image
	// Required: true
	BaseImage *BaseImageInfo `json:"baseImage"`

	// Personality components
	Components map[string]ComponentInfo `json:"components,omitempty"`

	// Personality hardware support
	HardwareSupport *HardwareSupportInfo `json:"hardwareSupport,omitempty"`
}

// Validate validates this software info
func (m *SoftwareInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAddOn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBaseImage(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateComponents(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHardwareSupport(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SoftwareInfo) validateAddOn(formats strfmt.Registry) error {
	if swag.IsZero(m.AddOn) { // not required
		return nil
	}

	if m.AddOn != nil {
		if err := m.AddOn.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("addOn")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("addOn")
			}
			return err
		}
	}

	return nil
}

func (m *SoftwareInfo) validateBaseImage(formats strfmt.Registry) error {

	if err := validate.Required("baseImage", "body", m.BaseImage); err != nil {
		return err
	}

	if m.BaseImage != nil {
		if err := m.BaseImage.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("baseImage")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("baseImage")
			}
			return err
		}
	}

	return nil
}

func (m *SoftwareInfo) validateComponents(formats strfmt.Registry) error {
	if swag.IsZero(m.Components) { // not required
		return nil
	}

	for k := range m.Components {

		if err := validate.Required("components"+"."+k, "body", m.Components[k]); err != nil {
			return err
		}
		if val, ok := m.Components[k]; ok {
			if err := val.Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("components" + "." + k)
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("components" + "." + k)
				}
				return err
			}
		}

	}

	return nil
}

func (m *SoftwareInfo) validateHardwareSupport(formats strfmt.Registry) error {
	if swag.IsZero(m.HardwareSupport) { // not required
		return nil
	}

	if m.HardwareSupport != nil {
		if err := m.HardwareSupport.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("hardwareSupport")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("hardwareSupport")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this software info based on the context it is used
func (m *SoftwareInfo) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAddOn(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateBaseImage(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateComponents(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateHardwareSupport(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SoftwareInfo) contextValidateAddOn(ctx context.Context, formats strfmt.Registry) error {

	if m.AddOn != nil {
		if err := m.AddOn.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("addOn")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("addOn")
			}
			return err
		}
	}

	return nil
}

func (m *SoftwareInfo) contextValidateBaseImage(ctx context.Context, formats strfmt.Registry) error {

	if m.BaseImage != nil {
		if err := m.BaseImage.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("baseImage")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("baseImage")
			}
			return err
		}
	}

	return nil
}

func (m *SoftwareInfo) contextValidateComponents(ctx context.Context, formats strfmt.Registry) error {

	for k := range m.Components {

		if val, ok := m.Components[k]; ok {
			if err := val.ContextValidate(ctx, formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *SoftwareInfo) contextValidateHardwareSupport(ctx context.Context, formats strfmt.Registry) error {

	if m.HardwareSupport != nil {
		if err := m.HardwareSupport.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("hardwareSupport")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("hardwareSupport")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SoftwareInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SoftwareInfo) UnmarshalBinary(b []byte) error {
	var res SoftwareInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
