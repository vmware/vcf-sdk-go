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

// ErrorCodePack error code pack
//
// swagger:model ErrorCodePack
type ErrorCodePack struct {

	// arguments
	Arguments []interface{} `json:"arguments"`

	// bundle name
	BundleName string `json:"bundleName,omitempty"`

	// class loader
	ClassLoader ClassLoader `json:"classLoader,omitempty"`

	// code
	Code string `json:"code,omitempty"`

	// component
	Component string `json:"component,omitempty"`

	// error code
	ErrorCode *ErrorCode `json:"errorCode,omitempty"`

	// string arguments
	StringArguments []string `json:"stringArguments"`
}

// Validate validates this error code pack
func (m *ErrorCodePack) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateErrorCode(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ErrorCodePack) validateErrorCode(formats strfmt.Registry) error {
	if swag.IsZero(m.ErrorCode) { // not required
		return nil
	}

	if m.ErrorCode != nil {
		if err := m.ErrorCode.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("errorCode")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("errorCode")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this error code pack based on the context it is used
func (m *ErrorCodePack) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateErrorCode(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ErrorCodePack) contextValidateErrorCode(ctx context.Context, formats strfmt.Registry) error {

	if m.ErrorCode != nil {

		if swag.IsZero(m.ErrorCode) { // not required
			return nil
		}

		if err := m.ErrorCode.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("errorCode")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("errorCode")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ErrorCodePack) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ErrorCodePack) UnmarshalBinary(b []byte) error {
	var res ErrorCodePack
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
