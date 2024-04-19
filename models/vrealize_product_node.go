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

// VrealizeProductNode Spec contains information for a VMware Aria product node
//
// swagger:model VrealizeProductNode
type VrealizeProductNode struct {

	// The Fully Qualified Domain Name for the VMware Aria node (virtual appliance)
	// Example: vrealize.node.vrack.vsphere.local
	// Required: true
	Fqdn *string `json:"fqdn"`

	// The ID of the node
	ID string `json:"id,omitempty"`

	// IP Address of VMware Aria product appliance
	// Example: 10.0.0.17
	// Required: true
	IPAddress *string `json:"ipAddress"`

	// The type of the VMware Aria product node
	// Example: MASTER, REPLICA, DATA, REMOTE_COLLECTOR, WORKER
	// Enum: [MASTER REPLICA DATA REMOTE_COLLECTOR WORKER]
	Type string `json:"type,omitempty"`
}

// Validate validates this vrealize product node
func (m *VrealizeProductNode) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFqdn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIPAddress(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VrealizeProductNode) validateFqdn(formats strfmt.Registry) error {

	if err := validate.Required("fqdn", "body", m.Fqdn); err != nil {
		return err
	}

	return nil
}

func (m *VrealizeProductNode) validateIPAddress(formats strfmt.Registry) error {

	if err := validate.Required("ipAddress", "body", m.IPAddress); err != nil {
		return err
	}

	return nil
}

var vrealizeProductNodeTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["MASTER","REPLICA","DATA","REMOTE_COLLECTOR","WORKER"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		vrealizeProductNodeTypeTypePropEnum = append(vrealizeProductNodeTypeTypePropEnum, v)
	}
}

const (

	// VrealizeProductNodeTypeMASTER captures enum value "MASTER"
	VrealizeProductNodeTypeMASTER string = "MASTER"

	// VrealizeProductNodeTypeREPLICA captures enum value "REPLICA"
	VrealizeProductNodeTypeREPLICA string = "REPLICA"

	// VrealizeProductNodeTypeDATA captures enum value "DATA"
	VrealizeProductNodeTypeDATA string = "DATA"

	// VrealizeProductNodeTypeREMOTECOLLECTOR captures enum value "REMOTE_COLLECTOR"
	VrealizeProductNodeTypeREMOTECOLLECTOR string = "REMOTE_COLLECTOR"

	// VrealizeProductNodeTypeWORKER captures enum value "WORKER"
	VrealizeProductNodeTypeWORKER string = "WORKER"
)

// prop value enum
func (m *VrealizeProductNode) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, vrealizeProductNodeTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *VrealizeProductNode) validateType(formats strfmt.Registry) error {
	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this vrealize product node based on context it is used
func (m *VrealizeProductNode) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *VrealizeProductNode) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VrealizeProductNode) UnmarshalBinary(b []byte) error {
	var res VrealizeProductNode
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
