// Code generated by go-swagger; DO NOT EDIT.

// Copyright © 2023 VMware, Inc. All Rights Reserved.
// Code licenced under: BSD-2

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

// PersonalitySpec Personality Specification for vLCM based upgrade
//
// swagger:model PersonalitySpec
type PersonalitySpec struct {

	// Hardware Support Specifications for Firmware upgrade
	HardwareSupportSpecs []*HardwareSupportSpec `json:"hardwareSupportSpecs"`

	// Personality ID for vLCM based Upgrade
	// Required: true
	PersonalityID *string `json:"personalityId"`
}

// Validate validates this personality spec
func (m *PersonalitySpec) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateHardwareSupportSpecs(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePersonalityID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PersonalitySpec) validateHardwareSupportSpecs(formats strfmt.Registry) error {
	if swag.IsZero(m.HardwareSupportSpecs) { // not required
		return nil
	}

	for i := 0; i < len(m.HardwareSupportSpecs); i++ {
		if swag.IsZero(m.HardwareSupportSpecs[i]) { // not required
			continue
		}

		if m.HardwareSupportSpecs[i] != nil {
			if err := m.HardwareSupportSpecs[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("hardwareSupportSpecs" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("hardwareSupportSpecs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PersonalitySpec) validatePersonalityID(formats strfmt.Registry) error {

	if err := validate.Required("personalityId", "body", m.PersonalityID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this personality spec based on the context it is used
func (m *PersonalitySpec) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateHardwareSupportSpecs(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PersonalitySpec) contextValidateHardwareSupportSpecs(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.HardwareSupportSpecs); i++ {

		if m.HardwareSupportSpecs[i] != nil {
			if err := m.HardwareSupportSpecs[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("hardwareSupportSpecs" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("hardwareSupportSpecs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *PersonalitySpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PersonalitySpec) UnmarshalBinary(b []byte) error {
	var res PersonalitySpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
