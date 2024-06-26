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

// VvolDatastoreSpec This specification contains cluster storage configuration for vVol
//
// swagger:model VvolDatastoreSpec
type VvolDatastoreSpec struct {

	// Name of the datastore
	// Required: true
	Name *string `json:"name"`

	// Vasa Provider spec
	// Required: true
	VasaProviderSpec *VasaProviderSpec `json:"vasaProviderSpec"`
}

// Validate validates this vvol datastore spec
func (m *VvolDatastoreSpec) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVasaProviderSpec(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VvolDatastoreSpec) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *VvolDatastoreSpec) validateVasaProviderSpec(formats strfmt.Registry) error {

	if err := validate.Required("vasaProviderSpec", "body", m.VasaProviderSpec); err != nil {
		return err
	}

	if m.VasaProviderSpec != nil {
		if err := m.VasaProviderSpec.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vasaProviderSpec")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("vasaProviderSpec")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this vvol datastore spec based on the context it is used
func (m *VvolDatastoreSpec) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateVasaProviderSpec(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VvolDatastoreSpec) contextValidateVasaProviderSpec(ctx context.Context, formats strfmt.Registry) error {

	if m.VasaProviderSpec != nil {

		if err := m.VasaProviderSpec.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vasaProviderSpec")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("vasaProviderSpec")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *VvolDatastoreSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VvolDatastoreSpec) UnmarshalBinary(b []byte) error {
	var res VvolDatastoreSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
