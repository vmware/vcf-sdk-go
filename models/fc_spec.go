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

// FcSpec Cluster storage configuration for VMFS on FC
//
// swagger:model FcSpec
type FcSpec struct {

	// Datastore name used for cluster creation
	// Required: true
	DatastoreName *string `json:"datastoreName"`
}

// Validate validates this fc spec
func (m *FcSpec) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDatastoreName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FcSpec) validateDatastoreName(formats strfmt.Registry) error {

	if err := validate.Required("datastoreName", "body", m.DatastoreName); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this fc spec based on context it is used
func (m *FcSpec) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *FcSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FcSpec) UnmarshalBinary(b []byte) error {
	var res FcSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
