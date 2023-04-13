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

// SDDCHotPatchInfo Sddc Hot patch info rest model that is located in the 2.0 manifest, as part of the async patch catalog.
//
// swagger:model SddcHotPatchInfo
type SDDCHotPatchInfo struct {

	// Sddc Hot patch min target version
	// Required: true
	SDDCHotPatchMinTargetVcfVersion *string `json:"sddcHotPatchMinTargetVcfVersion"`

	// Product version
	// Required: true
	Version *string `json:"version"`
}

// Validate validates this Sddc hot patch info
func (m *SDDCHotPatchInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSDDCHotPatchMinTargetVcfVersion(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVersion(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SDDCHotPatchInfo) validateSDDCHotPatchMinTargetVcfVersion(formats strfmt.Registry) error {

	if err := validate.Required("sddcHotPatchMinTargetVcfVersion", "body", m.SDDCHotPatchMinTargetVcfVersion); err != nil {
		return err
	}

	return nil
}

func (m *SDDCHotPatchInfo) validateVersion(formats strfmt.Registry) error {

	if err := validate.Required("version", "body", m.Version); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this Sddc hot patch info based on context it is used
func (m *SDDCHotPatchInfo) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SDDCHotPatchInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SDDCHotPatchInfo) UnmarshalBinary(b []byte) error {
	var res SDDCHotPatchInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
