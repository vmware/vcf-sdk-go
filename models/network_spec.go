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

// NetworkSpec This specification contains cluster network configuration
//
// swagger:model NetworkSpec
type NetworkSpec struct {

	// NSX configuration to add to the cluster
	// Required: true
	NsxClusterSpec *NsxClusterSpec `json:"nsxClusterSpec"`

	// Distributed switches to add to the cluster
	// Required: true
	VdsSpecs []*VdsSpec `json:"vdsSpecs"`
}

// Validate validates this network spec
func (m *NetworkSpec) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateNsxClusterSpec(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVdsSpecs(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NetworkSpec) validateNsxClusterSpec(formats strfmt.Registry) error {

	if err := validate.Required("nsxClusterSpec", "body", m.NsxClusterSpec); err != nil {
		return err
	}

	if m.NsxClusterSpec != nil {
		if err := m.NsxClusterSpec.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("nsxClusterSpec")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("nsxClusterSpec")
			}
			return err
		}
	}

	return nil
}

func (m *NetworkSpec) validateVdsSpecs(formats strfmt.Registry) error {

	if err := validate.Required("vdsSpecs", "body", m.VdsSpecs); err != nil {
		return err
	}

	for i := 0; i < len(m.VdsSpecs); i++ {
		if swag.IsZero(m.VdsSpecs[i]) { // not required
			continue
		}

		if m.VdsSpecs[i] != nil {
			if err := m.VdsSpecs[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("vdsSpecs" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("vdsSpecs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this network spec based on the context it is used
func (m *NetworkSpec) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateNsxClusterSpec(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVdsSpecs(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NetworkSpec) contextValidateNsxClusterSpec(ctx context.Context, formats strfmt.Registry) error {

	if m.NsxClusterSpec != nil {
		if err := m.NsxClusterSpec.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("nsxClusterSpec")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("nsxClusterSpec")
			}
			return err
		}
	}

	return nil
}

func (m *NetworkSpec) contextValidateVdsSpecs(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.VdsSpecs); i++ {

		if m.VdsSpecs[i] != nil {
			if err := m.VdsSpecs[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("vdsSpecs" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("vdsSpecs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *NetworkSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NetworkSpec) UnmarshalBinary(b []byte) error {
	var res NetworkSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
