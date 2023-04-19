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

// MicrosoftCertificateAuthoritySpec This specification contains Microsoft CA configuration details
//
// swagger:model MicrosoftCertificateAuthoritySpec
type MicrosoftCertificateAuthoritySpec struct {

	// CA server password
	// Required: true
	Secret *string `json:"secret"`

	// CA server url
	// Required: true
	ServerURL *string `json:"serverUrl"`

	// CA server template name
	// Required: true
	TemplateName *string `json:"templateName"`

	// CA server username
	// Required: true
	Username *string `json:"username"`
}

// Validate validates this microsoft certificate authority spec
func (m *MicrosoftCertificateAuthoritySpec) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSecret(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateServerURL(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTemplateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUsername(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MicrosoftCertificateAuthoritySpec) validateSecret(formats strfmt.Registry) error {

	if err := validate.Required("secret", "body", m.Secret); err != nil {
		return err
	}

	return nil
}

func (m *MicrosoftCertificateAuthoritySpec) validateServerURL(formats strfmt.Registry) error {

	if err := validate.Required("serverUrl", "body", m.ServerURL); err != nil {
		return err
	}

	return nil
}

func (m *MicrosoftCertificateAuthoritySpec) validateTemplateName(formats strfmt.Registry) error {

	if err := validate.Required("templateName", "body", m.TemplateName); err != nil {
		return err
	}

	return nil
}

func (m *MicrosoftCertificateAuthoritySpec) validateUsername(formats strfmt.Registry) error {

	if err := validate.Required("username", "body", m.Username); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this microsoft certificate authority spec based on context it is used
func (m *MicrosoftCertificateAuthoritySpec) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *MicrosoftCertificateAuthoritySpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MicrosoftCertificateAuthoritySpec) UnmarshalBinary(b []byte) error {
	var res MicrosoftCertificateAuthoritySpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
