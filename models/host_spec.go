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

// HostSpec This specification contains information required to add vSphere host to a cluster
//
// swagger:model HostSpec
type HostSpec struct {

	// Availability Zone Name
	// (This is required while performing a stretched cluster expand operation)
	AzName string `json:"azName,omitempty"`

	// Host name of the vSphere host
	HostName string `json:"hostName,omitempty"`

	// Network Details of the vSphere host
	HostNetworkSpec *HostNetworkSpec `json:"hostNetworkSpec,omitempty"`

	// ID of a vSphere host in the free pool
	// Required: true
	ID *string `json:"id"`

	// IP address of the vSphere host
	IPAddress string `json:"ipAddress,omitempty"`

	// License key of a vSphere host in the free pool
	// (This is required except in cases where the ESXi host has already been licensed outside of the VMware Cloud Foundation system)
	LicenseKey string `json:"licenseKey,omitempty"`

	// SSH password of the vSphere host
	Password string `json:"password,omitempty"`

	// Serial Number of the vSphere host
	SerialNumber string `json:"serialNumber,omitempty"`

	// SSH thumbprint(fingerprint) of the vSphere host
	// Note:This field will be mandatory in future releases.
	SSHThumbprint string `json:"sshThumbprint,omitempty"`

	// Username of the vSphere host
	Username string `json:"username,omitempty"`
}

// Validate validates this host spec
func (m *HostSpec) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateHostNetworkSpec(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HostSpec) validateHostNetworkSpec(formats strfmt.Registry) error {
	if swag.IsZero(m.HostNetworkSpec) { // not required
		return nil
	}

	if m.HostNetworkSpec != nil {
		if err := m.HostNetworkSpec.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("hostNetworkSpec")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("hostNetworkSpec")
			}
			return err
		}
	}

	return nil
}

func (m *HostSpec) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this host spec based on the context it is used
func (m *HostSpec) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateHostNetworkSpec(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HostSpec) contextValidateHostNetworkSpec(ctx context.Context, formats strfmt.Registry) error {

	if m.HostNetworkSpec != nil {
		if err := m.HostNetworkSpec.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("hostNetworkSpec")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("hostNetworkSpec")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *HostSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HostSpec) UnmarshalBinary(b []byte) error {
	var res HostSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
