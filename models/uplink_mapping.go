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

// UplinkMapping The map of vSphere Distributed Switch uplinks to the NSX switch uplinks.
//
// swagger:model UplinkMapping
type UplinkMapping struct {

	// The uplink name of the NSX switch
	// Required: true
	NsxUplinkName *string `json:"nsxUplinkName"`

	// The uplink name of the vSphere Distributed Switch
	// Required: true
	VdsUplinkName *string `json:"vdsUplinkName"`
}

// Validate validates this uplink mapping
func (m *UplinkMapping) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateNsxUplinkName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVdsUplinkName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UplinkMapping) validateNsxUplinkName(formats strfmt.Registry) error {

	if err := validate.Required("nsxUplinkName", "body", m.NsxUplinkName); err != nil {
		return err
	}

	return nil
}

func (m *UplinkMapping) validateVdsUplinkName(formats strfmt.Registry) error {

	if err := validate.Required("vdsUplinkName", "body", m.VdsUplinkName); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this uplink mapping based on context it is used
func (m *UplinkMapping) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *UplinkMapping) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UplinkMapping) UnmarshalBinary(b []byte) error {
	var res UplinkMapping
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
