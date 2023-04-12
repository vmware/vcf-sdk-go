// Code generated by go-swagger; DO NOT EDIT.

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

// UnmanagedResourceCredential Represents credentials of am unmanaged resource (i.e a resource that is not managed by VCF)
//
// swagger:model UnmanagedResourceCredential
type UnmanagedResourceCredential struct {

	// Credential type
	// Example: One among: SSH
	// Required: true
	CredentialType *string `json:"credentialType"`

	// Password
	Password string `json:"password,omitempty"`

	// Username
	// Required: true
	Username *string `json:"username"`
}

// Validate validates this unmanaged resource credential
func (m *UnmanagedResourceCredential) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCredentialType(formats); err != nil {
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

func (m *UnmanagedResourceCredential) validateCredentialType(formats strfmt.Registry) error {

	if err := validate.Required("credentialType", "body", m.CredentialType); err != nil {
		return err
	}

	return nil
}

func (m *UnmanagedResourceCredential) validateUsername(formats strfmt.Registry) error {

	if err := validate.Required("username", "body", m.Username); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this unmanaged resource credential based on context it is used
func (m *UnmanagedResourceCredential) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *UnmanagedResourceCredential) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UnmanagedResourceCredential) UnmarshalBinary(b []byte) error {
	var res UnmanagedResourceCredential
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}