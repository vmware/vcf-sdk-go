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

// Resource Represents the resource in the system
//
// swagger:model Resource
type Resource struct {

	// Resource FQDN
	// Example: sfo-vc01.rainpole.io
	Fqdn string `json:"fqdn,omitempty"`

	// Name of the resource
	Name string `json:"name,omitempty"`

	// Resource ID
	// Example: BE8A5E04-92A0-43F6-A166-AA041F4327CC
	// Required: true
	ResourceID *string `json:"resourceId"`

	// Subject alternative name(s)
	// Example: sfo-vc01.rainpole.io
	Sans []string `json:"sans"`

	// Resource type
	// Example: One among: SDDC_MANAGER, PSC, VCENTER, NSX_MANAGER, NSXT_MANAGER, VRA, VRLI, VROPS, VRSLCM, VXRAIL_MANAGER
	// Required: true
	Type *string `json:"type"`
}

// Validate validates this resource
func (m *Resource) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateResourceID(formats); err != nil {
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

func (m *Resource) validateResourceID(formats strfmt.Registry) error {

	if err := validate.Required("resourceId", "body", m.ResourceID); err != nil {
		return err
	}

	return nil
}

func (m *Resource) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this resource based on context it is used
func (m *Resource) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Resource) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Resource) UnmarshalBinary(b []byte) error {
	var res Resource
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
