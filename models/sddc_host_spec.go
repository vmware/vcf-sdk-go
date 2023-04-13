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

// SDDCHostSpec Spec contains parameters for Host
//
// swagger:model SddcHostSpec
type SDDCHostSpec struct {

	// Host Association: Location/Datacenter
	// Required: true
	Association *string `json:"association"`

	// Host Credentials
	// Required: true
	Credentials *SDDCCredentials `json:"credentials"`

	// Host Hostname
	// Example: esx-1
	// Required: true
	// Max Length: 63
	// Min Length: 3
	Hostname *string `json:"hostname"`

	// Host Private Management IP
	// Required: true
	IPAddressPrivate *IPAllocation `json:"ipAddressPrivate"`

	// Host key
	Key string `json:"key,omitempty"`

	// Host server ID
	ServerID string `json:"serverId,omitempty"`

	// Host SSH thumbprint (RSA SHA256)
	SSHThumbprint string `json:"sshThumbprint,omitempty"`

	// Host SSL thumbprint (SHA256)
	SslThumbprint string `json:"sslThumbprint,omitempty"`

	// v switch
	// Required: true
	VSwitch *string `json:"vSwitch"`

	// List of Host Vmknic Spec
	VmknicSpecs []*HostVmknicSpec `json:"vmknicSpecs"`

	// vswitch
	Vswitch string `json:"vswitch,omitempty"`
}

// Validate validates this Sddc host spec
func (m *SDDCHostSpec) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAssociation(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCredentials(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHostname(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIPAddressPrivate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVSwitch(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVmknicSpecs(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SDDCHostSpec) validateAssociation(formats strfmt.Registry) error {

	if err := validate.Required("association", "body", m.Association); err != nil {
		return err
	}

	return nil
}

func (m *SDDCHostSpec) validateCredentials(formats strfmt.Registry) error {

	if err := validate.Required("credentials", "body", m.Credentials); err != nil {
		return err
	}

	if m.Credentials != nil {
		if err := m.Credentials.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("credentials")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("credentials")
			}
			return err
		}
	}

	return nil
}

func (m *SDDCHostSpec) validateHostname(formats strfmt.Registry) error {

	if err := validate.Required("hostname", "body", m.Hostname); err != nil {
		return err
	}

	if err := validate.MinLength("hostname", "body", *m.Hostname, 3); err != nil {
		return err
	}

	if err := validate.MaxLength("hostname", "body", *m.Hostname, 63); err != nil {
		return err
	}

	return nil
}

func (m *SDDCHostSpec) validateIPAddressPrivate(formats strfmt.Registry) error {

	if err := validate.Required("ipAddressPrivate", "body", m.IPAddressPrivate); err != nil {
		return err
	}

	if m.IPAddressPrivate != nil {
		if err := m.IPAddressPrivate.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ipAddressPrivate")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("ipAddressPrivate")
			}
			return err
		}
	}

	return nil
}

func (m *SDDCHostSpec) validateVSwitch(formats strfmt.Registry) error {

	if err := validate.Required("vSwitch", "body", m.VSwitch); err != nil {
		return err
	}

	return nil
}

func (m *SDDCHostSpec) validateVmknicSpecs(formats strfmt.Registry) error {
	if swag.IsZero(m.VmknicSpecs) { // not required
		return nil
	}

	for i := 0; i < len(m.VmknicSpecs); i++ {
		if swag.IsZero(m.VmknicSpecs[i]) { // not required
			continue
		}

		if m.VmknicSpecs[i] != nil {
			if err := m.VmknicSpecs[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("vmknicSpecs" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("vmknicSpecs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this Sddc host spec based on the context it is used
func (m *SDDCHostSpec) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCredentials(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateIPAddressPrivate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVmknicSpecs(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SDDCHostSpec) contextValidateCredentials(ctx context.Context, formats strfmt.Registry) error {

	if m.Credentials != nil {
		if err := m.Credentials.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("credentials")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("credentials")
			}
			return err
		}
	}

	return nil
}

func (m *SDDCHostSpec) contextValidateIPAddressPrivate(ctx context.Context, formats strfmt.Registry) error {

	if m.IPAddressPrivate != nil {
		if err := m.IPAddressPrivate.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ipAddressPrivate")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("ipAddressPrivate")
			}
			return err
		}
	}

	return nil
}

func (m *SDDCHostSpec) contextValidateVmknicSpecs(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.VmknicSpecs); i++ {

		if m.VmknicSpecs[i] != nil {
			if err := m.VmknicSpecs[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("vmknicSpecs" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("vmknicSpecs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *SDDCHostSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SDDCHostSpec) UnmarshalBinary(b []byte) error {
	var res SDDCHostSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
