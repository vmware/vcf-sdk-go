// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2023 VMware, Inc.
// SPDX-License-Identifier: BSD-2-Clause

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// SecuritySpec Spec contains security settings
//
// swagger:model SecuritySpec
type SecuritySpec struct {

	// ESXi certificates mode
	// Example: One among:Custom, VMCA
	EsxiCertsMode string `json:"esxiCertsMode,omitempty"`

	// Root Certificate Authority certificate list
	RootCaCerts []*RootCaCerts `json:"rootCaCerts"`
}

// Validate validates this security spec
func (m *SecuritySpec) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRootCaCerts(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SecuritySpec) validateRootCaCerts(formats strfmt.Registry) error {
	if swag.IsZero(m.RootCaCerts) { // not required
		return nil
	}

	for i := 0; i < len(m.RootCaCerts); i++ {
		if swag.IsZero(m.RootCaCerts[i]) { // not required
			continue
		}

		if m.RootCaCerts[i] != nil {
			if err := m.RootCaCerts[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("rootCaCerts" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("rootCaCerts" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this security spec based on the context it is used
func (m *SecuritySpec) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateRootCaCerts(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SecuritySpec) contextValidateRootCaCerts(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.RootCaCerts); i++ {

		if m.RootCaCerts[i] != nil {

			if swag.IsZero(m.RootCaCerts[i]) { // not required
				return nil
			}

			if err := m.RootCaCerts[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("rootCaCerts" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("rootCaCerts" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *SecuritySpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SecuritySpec) UnmarshalBinary(b []byte) error {
	var res SecuritySpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
