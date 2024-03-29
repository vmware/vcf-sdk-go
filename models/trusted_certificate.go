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

// TrustedCertificate The Trusted Certificate details.
//
// swagger:model TrustedCertificate
type TrustedCertificate struct {

	// Certificate alias
	// Example: vcf_59:24:D5:18:04:A0:26:B0:A4:05:EA:82:60:95:82:A2:4B:F6:31:FB:81:93:01:F3:29:7D:34:9C:D3:05:39:90
	// Required: true
	Alias *string `json:"alias"`

	// Certificate in PEM format
	// Example: -----BEGIN CERTIFICATE-----\nMIIFq...\n-----END CERTIFICATE-----
	// Required: true
	Certificate *string `json:"certificate"`
}

// Validate validates this trusted certificate
func (m *TrustedCertificate) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAlias(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCertificate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TrustedCertificate) validateAlias(formats strfmt.Registry) error {

	if err := validate.Required("alias", "body", m.Alias); err != nil {
		return err
	}

	return nil
}

func (m *TrustedCertificate) validateCertificate(formats strfmt.Registry) error {

	if err := validate.Required("certificate", "body", m.Certificate); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this trusted certificate based on context it is used
func (m *TrustedCertificate) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *TrustedCertificate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TrustedCertificate) UnmarshalBinary(b []byte) error {
	var res TrustedCertificate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
