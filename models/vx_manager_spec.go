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

// VxManagerSpec Spec contains parameters for VxManager
//
// swagger:model VxManagerSpec
type VxManagerSpec struct {

	// Default admin credentials VxManager
	// Required: true
	DefaultAdminUserCredentials *SddcCredentials `json:"defaultAdminUserCredentials"`

	// Default root credentials VxManager
	// Required: true
	DefaultRootUserCredentials *SddcCredentials `json:"defaultRootUserCredentials"`

	// VxRail Manager SSH thumbprint (RSA SHA256)
	SSHThumbprint string `json:"sshThumbprint,omitempty"`

	// VxRail Manager SSL thumbprint (SHA256)
	SslThumbprint string `json:"sslThumbprint,omitempty"`

	// VxManager host name
	// Required: true
	VxManagerHostName *string `json:"vxManagerHostName"`
}

// Validate validates this vx manager spec
func (m *VxManagerSpec) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDefaultAdminUserCredentials(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDefaultRootUserCredentials(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVxManagerHostName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VxManagerSpec) validateDefaultAdminUserCredentials(formats strfmt.Registry) error {

	if err := validate.Required("defaultAdminUserCredentials", "body", m.DefaultAdminUserCredentials); err != nil {
		return err
	}

	if m.DefaultAdminUserCredentials != nil {
		if err := m.DefaultAdminUserCredentials.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("defaultAdminUserCredentials")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("defaultAdminUserCredentials")
			}
			return err
		}
	}

	return nil
}

func (m *VxManagerSpec) validateDefaultRootUserCredentials(formats strfmt.Registry) error {

	if err := validate.Required("defaultRootUserCredentials", "body", m.DefaultRootUserCredentials); err != nil {
		return err
	}

	if m.DefaultRootUserCredentials != nil {
		if err := m.DefaultRootUserCredentials.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("defaultRootUserCredentials")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("defaultRootUserCredentials")
			}
			return err
		}
	}

	return nil
}

func (m *VxManagerSpec) validateVxManagerHostName(formats strfmt.Registry) error {

	if err := validate.Required("vxManagerHostName", "body", m.VxManagerHostName); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this vx manager spec based on the context it is used
func (m *VxManagerSpec) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDefaultAdminUserCredentials(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDefaultRootUserCredentials(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VxManagerSpec) contextValidateDefaultAdminUserCredentials(ctx context.Context, formats strfmt.Registry) error {

	if m.DefaultAdminUserCredentials != nil {
		if err := m.DefaultAdminUserCredentials.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("defaultAdminUserCredentials")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("defaultAdminUserCredentials")
			}
			return err
		}
	}

	return nil
}

func (m *VxManagerSpec) contextValidateDefaultRootUserCredentials(ctx context.Context, formats strfmt.Registry) error {

	if m.DefaultRootUserCredentials != nil {
		if err := m.DefaultRootUserCredentials.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("defaultRootUserCredentials")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("defaultRootUserCredentials")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *VxManagerSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VxManagerSpec) UnmarshalBinary(b []byte) error {
	var res VxManagerSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
