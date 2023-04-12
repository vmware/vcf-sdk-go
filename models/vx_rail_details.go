// Code generated by go-swagger; DO NOT EDIT.

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

// VxRailDetails Contains the VxRail Manager details
//
// swagger:model VxRailDetails
type VxRailDetails struct {

	// VxRail Manager admin credentials
	AdminCredentials *UnmanagedResourceCredential `json:"adminCredentials,omitempty"`

	// Map of Context class with list of key and value pairs
	ContextWithKeyValuePair map[string]List `json:"contextWithKeyValuePair,omitempty"`

	// DNS Name/Hostname of the VxRail Manager
	DNSName string `json:"dnsName,omitempty"`

	// IP Address of the VxRail Manager
	IPAddress string `json:"ipAddress,omitempty"`

	// Network details of the VxRail Manager
	Networks []*Network `json:"networks"`

	// Nic Profile Type
	NicProfile string `json:"nicProfile,omitempty"`

	// VxRail Manager root credentials
	RootCredentials *UnmanagedResourceCredential `json:"rootCredentials,omitempty"`

	// SSH thumbprint of the VxRail Manager
	SSHThumbprint string `json:"sshThumbprint,omitempty"`

	// SSL thumbprint of the VxRail Manager
	SslThumbprint string `json:"sslThumbprint,omitempty"`
}

// Validate validates this vx rail details
func (m *VxRailDetails) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAdminCredentials(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateContextWithKeyValuePair(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNetworks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRootCredentials(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VxRailDetails) validateAdminCredentials(formats strfmt.Registry) error {
	if swag.IsZero(m.AdminCredentials) { // not required
		return nil
	}

	if m.AdminCredentials != nil {
		if err := m.AdminCredentials.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("adminCredentials")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("adminCredentials")
			}
			return err
		}
	}

	return nil
}

func (m *VxRailDetails) validateContextWithKeyValuePair(formats strfmt.Registry) error {
	if swag.IsZero(m.ContextWithKeyValuePair) { // not required
		return nil
	}

	for k := range m.ContextWithKeyValuePair {

		if err := m.ContextWithKeyValuePair[k].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("contextWithKeyValuePair" + "." + k)
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("contextWithKeyValuePair" + "." + k)
			}
			return err
		}

	}

	return nil
}

func (m *VxRailDetails) validateNetworks(formats strfmt.Registry) error {
	if swag.IsZero(m.Networks) { // not required
		return nil
	}

	for i := 0; i < len(m.Networks); i++ {
		if swag.IsZero(m.Networks[i]) { // not required
			continue
		}

		if m.Networks[i] != nil {
			if err := m.Networks[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("networks" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("networks" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *VxRailDetails) validateRootCredentials(formats strfmt.Registry) error {
	if swag.IsZero(m.RootCredentials) { // not required
		return nil
	}

	if m.RootCredentials != nil {
		if err := m.RootCredentials.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("rootCredentials")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("rootCredentials")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this vx rail details based on the context it is used
func (m *VxRailDetails) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAdminCredentials(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateContextWithKeyValuePair(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNetworks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRootCredentials(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VxRailDetails) contextValidateAdminCredentials(ctx context.Context, formats strfmt.Registry) error {

	if m.AdminCredentials != nil {
		if err := m.AdminCredentials.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("adminCredentials")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("adminCredentials")
			}
			return err
		}
	}

	return nil
}

func (m *VxRailDetails) contextValidateContextWithKeyValuePair(ctx context.Context, formats strfmt.Registry) error {

	for k := range m.ContextWithKeyValuePair {

		if err := m.ContextWithKeyValuePair[k].ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("contextWithKeyValuePair" + "." + k)
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("contextWithKeyValuePair" + "." + k)
			}
			return err
		}

	}

	return nil
}

func (m *VxRailDetails) contextValidateNetworks(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Networks); i++ {

		if m.Networks[i] != nil {
			if err := m.Networks[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("networks" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("networks" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *VxRailDetails) contextValidateRootCredentials(ctx context.Context, formats strfmt.Registry) error {

	if m.RootCredentials != nil {
		if err := m.RootCredentials.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("rootCredentials")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("rootCredentials")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *VxRailDetails) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VxRailDetails) UnmarshalBinary(b []byte) error {
	var res VxRailDetails
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
