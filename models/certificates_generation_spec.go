// Code generated by go-swagger; DO NOT EDIT.

// Copyright © 2023 VMware, Inc. All Rights Reserved.
// Code licenced under: BSD-2

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

// CertificatesGenerationSpec This specification contains certificates generation request details
//
// swagger:model CertificatesGenerationSpec
type CertificatesGenerationSpec struct {

	// CA type. Only supports Microsoft and OpenSSL CAs
	// Example: One among: OpenSSL, Microsoft
	// Required: true
	CaType *string `json:"caType"`

	// Resource(s) in a domain
	Resources []*Resource `json:"resources"`
}

// Validate validates this certificates generation spec
func (m *CertificatesGenerationSpec) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCaType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResources(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CertificatesGenerationSpec) validateCaType(formats strfmt.Registry) error {

	if err := validate.Required("caType", "body", m.CaType); err != nil {
		return err
	}

	return nil
}

func (m *CertificatesGenerationSpec) validateResources(formats strfmt.Registry) error {
	if swag.IsZero(m.Resources) { // not required
		return nil
	}

	for i := 0; i < len(m.Resources); i++ {
		if swag.IsZero(m.Resources[i]) { // not required
			continue
		}

		if m.Resources[i] != nil {
			if err := m.Resources[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("resources" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("resources" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this certificates generation spec based on the context it is used
func (m *CertificatesGenerationSpec) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateResources(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CertificatesGenerationSpec) contextValidateResources(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Resources); i++ {

		if m.Resources[i] != nil {
			if err := m.Resources[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("resources" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("resources" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *CertificatesGenerationSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CertificatesGenerationSpec) UnmarshalBinary(b []byte) error {
	var res CertificatesGenerationSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
