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

// SupportBundleScope support bundle scope
//
// swagger:model SupportBundleScope
type SupportBundleScope struct {

	// Domains and Clusters for SOS operation.
	Domains []*Domains `json:"domains"`

	// Include free hosts.
	IncludeFreeHosts bool `json:"includeFreeHosts,omitempty"`
}

// Validate validates this support bundle scope
func (m *SupportBundleScope) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDomains(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SupportBundleScope) validateDomains(formats strfmt.Registry) error {
	if swag.IsZero(m.Domains) { // not required
		return nil
	}

	for i := 0; i < len(m.Domains); i++ {
		if swag.IsZero(m.Domains[i]) { // not required
			continue
		}

		if m.Domains[i] != nil {
			if err := m.Domains[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("domains" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("domains" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this support bundle scope based on the context it is used
func (m *SupportBundleScope) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDomains(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SupportBundleScope) contextValidateDomains(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Domains); i++ {

		if m.Domains[i] != nil {

			if swag.IsZero(m.Domains[i]) { // not required
				return nil
			}

			if err := m.Domains[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("domains" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("domains" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *SupportBundleScope) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SupportBundleScope) UnmarshalBinary(b []byte) error {
	var res SupportBundleScope
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
