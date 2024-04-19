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
	"github.com/go-openapi/validate"
)

// IdentityProviderPrecheckResult Represents Identity Management pre-check result
//
// swagger:model IdentityProviderPrecheckResult
type IdentityProviderPrecheckResult struct {

	// Details about status type and error messages
	PrecheckInfo []*PrecheckInfo `json:"precheckInfo"`

	// One of SUCCESS, WARNING, FAILURE
	// Required: true
	Status *string `json:"status"`
}

// Validate validates this identity provider precheck result
func (m *IdentityProviderPrecheckResult) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePrecheckInfo(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IdentityProviderPrecheckResult) validatePrecheckInfo(formats strfmt.Registry) error {
	if swag.IsZero(m.PrecheckInfo) { // not required
		return nil
	}

	for i := 0; i < len(m.PrecheckInfo); i++ {
		if swag.IsZero(m.PrecheckInfo[i]) { // not required
			continue
		}

		if m.PrecheckInfo[i] != nil {
			if err := m.PrecheckInfo[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("precheckInfo" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("precheckInfo" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *IdentityProviderPrecheckResult) validateStatus(formats strfmt.Registry) error {

	if err := validate.Required("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this identity provider precheck result based on the context it is used
func (m *IdentityProviderPrecheckResult) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidatePrecheckInfo(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IdentityProviderPrecheckResult) contextValidatePrecheckInfo(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.PrecheckInfo); i++ {

		if m.PrecheckInfo[i] != nil {

			if swag.IsZero(m.PrecheckInfo[i]) { // not required
				return nil
			}

			if err := m.PrecheckInfo[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("precheckInfo" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("precheckInfo" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *IdentityProviderPrecheckResult) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IdentityProviderPrecheckResult) UnmarshalBinary(b []byte) error {
	var res IdentityProviderPrecheckResult
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}