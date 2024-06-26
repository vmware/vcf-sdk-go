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
)

// Vcenter vCenter representation
//
// swagger:model Vcenter
type Vcenter struct {

	// Information about the domain this vCenter is part of
	Domain *DomainReference `json:"domain,omitempty"`

	// FQDN of the vCenter
	Fqdn string `json:"fqdn,omitempty"`

	// ID of the vCenter
	ID string `json:"id,omitempty"`

	// IP address of the vCenter
	IPAddress string `json:"ipAddress,omitempty"`

	// Version of the vCenter
	Version string `json:"version,omitempty"`
}

// Validate validates this vcenter
func (m *Vcenter) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDomain(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Vcenter) validateDomain(formats strfmt.Registry) error {
	if swag.IsZero(m.Domain) { // not required
		return nil
	}

	if m.Domain != nil {
		if err := m.Domain.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("domain")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("domain")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this vcenter based on the context it is used
func (m *Vcenter) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDomain(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Vcenter) contextValidateDomain(ctx context.Context, formats strfmt.Registry) error {

	if m.Domain != nil {

		if swag.IsZero(m.Domain) { // not required
			return nil
		}

		if err := m.Domain.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("domain")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("domain")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Vcenter) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Vcenter) UnmarshalBinary(b []byte) error {
	var res Vcenter
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
