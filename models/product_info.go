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

// ProductInfo Product Info and License Key Info
//
// swagger:model ProductInfo
type ProductInfo struct {

	// Error Response while fetching product info
	ErrorResponse *Error `json:"errorResponse,omitempty"`

	// License key of product
	// Example: XXXXX-XXXXX-XXXXX-XXXXX-XXXXX
	LicenseKey string `json:"licenseKey,omitempty"`

	// License key expiration date
	// Example: 2021-04-08T04:37:51.662Z
	LicenseKeyExpirationDate string `json:"licenseKeyExpirationDate,omitempty"`

	// License key status
	// Example: One among: EXPIRED, ACTIVE, NEVER_EXPIRES
	LicenseKeyStatus string `json:"licenseKeyStatus,omitempty"`

	// Licensing Mode
	// Example: One among: SUBSCRIPTION, PERPETUAL
	LicensingMode string `json:"licensingMode,omitempty"`

	// Product type
	// Example: One among: VCENTER, NSXT, VSAN, WCP, ESXI
	Type string `json:"type,omitempty"`
}

// Validate validates this product info
func (m *ProductInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateErrorResponse(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ProductInfo) validateErrorResponse(formats strfmt.Registry) error {
	if swag.IsZero(m.ErrorResponse) { // not required
		return nil
	}

	if m.ErrorResponse != nil {
		if err := m.ErrorResponse.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("errorResponse")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("errorResponse")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this product info based on the context it is used
func (m *ProductInfo) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateErrorResponse(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ProductInfo) contextValidateErrorResponse(ctx context.Context, formats strfmt.Registry) error {

	if m.ErrorResponse != nil {

		if swag.IsZero(m.ErrorResponse) { // not required
			return nil
		}

		if err := m.ErrorResponse.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("errorResponse")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("errorResponse")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ProductInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ProductInfo) UnmarshalBinary(b []byte) error {
	var res ProductInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}