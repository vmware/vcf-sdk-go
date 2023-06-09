// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2023 VMware, Inc.
// SPDX-License-Identifier: BSD-2-Clause

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PhysicalNic Represents a physical NIC
//
// swagger:model PhysicalNic
type PhysicalNic struct {

	// Device name of the physical NIC
	DeviceName string `json:"deviceName,omitempty"`

	// Mac address of the physical NIC
	MacAddress string `json:"macAddress,omitempty"`

	// Speed in bytes of the physical NIC
	Speed int64 `json:"speed,omitempty"`

	// Unit of physical NIC speed
	// Enum: [KB MB GB TB PB]
	Unit string `json:"unit,omitempty"`
}

// Validate validates this physical nic
func (m *PhysicalNic) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateUnit(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var physicalNicTypeUnitPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["KB","MB","GB","TB","PB"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		physicalNicTypeUnitPropEnum = append(physicalNicTypeUnitPropEnum, v)
	}
}

const (

	// PhysicalNicUnitKB captures enum value "KB"
	PhysicalNicUnitKB string = "KB"

	// PhysicalNicUnitMB captures enum value "MB"
	PhysicalNicUnitMB string = "MB"

	// PhysicalNicUnitGB captures enum value "GB"
	PhysicalNicUnitGB string = "GB"

	// PhysicalNicUnitTB captures enum value "TB"
	PhysicalNicUnitTB string = "TB"

	// PhysicalNicUnitPB captures enum value "PB"
	PhysicalNicUnitPB string = "PB"
)

// prop value enum
func (m *PhysicalNic) validateUnitEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, physicalNicTypeUnitPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *PhysicalNic) validateUnit(formats strfmt.Registry) error {
	if swag.IsZero(m.Unit) { // not required
		return nil
	}

	// value enum
	if err := m.validateUnitEnum("unit", "body", m.Unit); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this physical nic based on context it is used
func (m *PhysicalNic) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PhysicalNic) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PhysicalNic) UnmarshalBinary(b []byte) error {
	var res PhysicalNic
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
