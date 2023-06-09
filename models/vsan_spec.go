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

// VSANSpec Spec contains parameters of Virtual SAN
//
// swagger:model VsanSpec
type VSANSpec struct {

	// Datastore Name
	// Required: true
	DatastoreName *string `json:"datastoreName"`

	// HCL File
	HclFile string `json:"hclFile,omitempty"`

	// License File
	LicenseFile string `json:"licenseFile,omitempty"`

	// VSAN feature Deduplication and Compression flag, one flag for both features
	VSANDedup bool `json:"vsanDedup,omitempty"`

	// Virtual SAN config name
	// Required: true
	VSANName *string `json:"vsanName"`
}

// Validate validates this Vsan spec
func (m *VSANSpec) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDatastoreName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVSANName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VSANSpec) validateDatastoreName(formats strfmt.Registry) error {

	if err := validate.Required("datastoreName", "body", m.DatastoreName); err != nil {
		return err
	}

	return nil
}

func (m *VSANSpec) validateVSANName(formats strfmt.Registry) error {

	if err := validate.Required("vsanName", "body", m.VSANName); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this Vsan spec based on context it is used
func (m *VSANSpec) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *VSANSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VSANSpec) UnmarshalBinary(b []byte) error {
	var res VSANSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
