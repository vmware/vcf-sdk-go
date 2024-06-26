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

// CSR The Certificate Signing Request (CSR) details
//
// swagger:model Csr
type CSR struct {

	// The CSR decoded content
	// Example: DECODED CSR AS STRING
	// Required: true
	CSRDecodedContent *string `json:"csrDecodedContent"`

	// The CSR encoded content
	// Example: -----BEGIN CERTIFICATE REQUEST-----\nMIIEH...\n-----END CERTIFICATE REQUEST-----
	// Required: true
	CSREncodedContent *string `json:"csrEncodedContent"`

	// Resource associated with CSR
	// Required: true
	Resource *Resource `json:"resource"`
}

// Validate validates this Csr
func (m *CSR) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCSRDecodedContent(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCSREncodedContent(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResource(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CSR) validateCSRDecodedContent(formats strfmt.Registry) error {

	if err := validate.Required("csrDecodedContent", "body", m.CSRDecodedContent); err != nil {
		return err
	}

	return nil
}

func (m *CSR) validateCSREncodedContent(formats strfmt.Registry) error {

	if err := validate.Required("csrEncodedContent", "body", m.CSREncodedContent); err != nil {
		return err
	}

	return nil
}

func (m *CSR) validateResource(formats strfmt.Registry) error {

	if err := validate.Required("resource", "body", m.Resource); err != nil {
		return err
	}

	if m.Resource != nil {
		if err := m.Resource.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("resource")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("resource")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this Csr based on the context it is used
func (m *CSR) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateResource(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CSR) contextValidateResource(ctx context.Context, formats strfmt.Registry) error {

	if m.Resource != nil {

		if err := m.Resource.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("resource")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("resource")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CSR) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CSR) UnmarshalBinary(b []byte) error {
	var res CSR
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
