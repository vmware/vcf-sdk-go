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

// VmnicToUplink This specification contains vmnic to uplink configurations for vSphere host.
//
// swagger:model VmnicToUplink
type VmnicToUplink struct {

	// VmNic ID of vSphere host to be associated with VDS, once added to cluster
	// Required: true
	ID *string `json:"id"`

	// The uplink name of the vSphere Distributed Switch to be associated
	// Required: true
	Uplink *string `json:"uplink"`
}

// Validate validates this vmnic to uplink
func (m *VmnicToUplink) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUplink(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VmnicToUplink) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *VmnicToUplink) validateUplink(formats strfmt.Registry) error {

	if err := validate.Required("uplink", "body", m.Uplink); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this vmnic to uplink based on context it is used
func (m *VmnicToUplink) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *VmnicToUplink) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VmnicToUplink) UnmarshalBinary(b []byte) error {
	var res VmnicToUplink
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
