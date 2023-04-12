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
	"github.com/go-openapi/validate"
)

// SddcNsxtSpec Spec contains parameters for NSX-T deployment and configurations
//
// swagger:model SddcNsxtSpec
type SddcNsxtSpec struct {

	// NSX-T IP address pool specification
	IPAddressPoolSpec *IPAddressPoolSpec `json:"ipAddressPoolSpec,omitempty"`

	// NSX-T admin password. The password must be at least 12 characters long. Must contain at-least 1 uppercase, 1 lowercase, 1 special character and 1 digit. In addition, a character cannot be repeated 3 or more times consectively.
	NsxtAdminPassword string `json:"nsxtAdminPassword,omitempty"`

	// NSX-T audit password. The password must be at least 12 characters long. Must contain at-least 1 uppercase, 1 lowercase, 1 special character and 1 digit. In addition, a character cannot be repeated 3 or more times consectively.
	NsxtAuditPassword string `json:"nsxtAuditPassword,omitempty"`

	// NSX-T Manager license
	NsxtLicense string `json:"nsxtLicense,omitempty"`

	// NSX-T Manager size
	// Example: One among:medium, large
	// Required: true
	NsxtManagerSize *string `json:"nsxtManagerSize"`

	// NSX-T Managers
	// Required: true
	NsxtManagers []*NsxtManagerSpec `json:"nsxtManagers"`

	// NSX-T OverLay Transport zone
	OverLayTransportZone *NsxtTransportZone `json:"overLayTransportZone,omitempty"`

	// If true, allow root login for NSX-T Manager and deny if false. Deprecated as of 4.5, the root login will be always enabled for NSX-T Data Center Appliance.
	RootLoginEnabledForNsxtManager bool `json:"rootLoginEnabledForNsxtManager,omitempty"`

	// NSX-T Manager root password. Password should have 1) At least eight characters, 2) At least one lower-case letter, 3) At least one upper-case letter 4) At least one digit 5) At least one special character, 6) At least five different characters , 7) No dictionary words, 6) No palindromes
	// Required: true
	RootNsxtManagerPassword *string `json:"rootNsxtManagerPassword"`

	// If true, enable SSH for NSX-T Manager and disable if false. Deprecated as of 4.5, the SSH connection will be always enabled for NSX-T Data Center Appliance.
	SSHEnabledForNsxtManager bool `json:"sshEnabledForNsxtManager,omitempty"`

	// Transport VLAN ID
	TransportVlanID int32 `json:"transportVlanId,omitempty"`

	// Virtual IP address which would act as proxy/alias for NSX-T Managers
	// Required: true
	Vip *string `json:"vip"`

	// FQDN for VIP so that common SSL certificates can be installed across all managers
	// Required: true
	VipFqdn *string `json:"vipFqdn"`

	// NSX-T VLAN transport zone.
	// This property is deprecated, and it is a no-operation
	VlanTransportZone *NsxtTransportZone `json:"vlanTransportZone,omitempty"`
}

// Validate validates this sddc nsxt spec
func (m *SddcNsxtSpec) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateIPAddressPoolSpec(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNsxtManagerSize(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNsxtManagers(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOverLayTransportZone(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRootNsxtManagerPassword(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVip(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVipFqdn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVlanTransportZone(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SddcNsxtSpec) validateIPAddressPoolSpec(formats strfmt.Registry) error {
	if swag.IsZero(m.IPAddressPoolSpec) { // not required
		return nil
	}

	if m.IPAddressPoolSpec != nil {
		if err := m.IPAddressPoolSpec.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ipAddressPoolSpec")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("ipAddressPoolSpec")
			}
			return err
		}
	}

	return nil
}

func (m *SddcNsxtSpec) validateNsxtManagerSize(formats strfmt.Registry) error {

	if err := validate.Required("nsxtManagerSize", "body", m.NsxtManagerSize); err != nil {
		return err
	}

	return nil
}

func (m *SddcNsxtSpec) validateNsxtManagers(formats strfmt.Registry) error {

	if err := validate.Required("nsxtManagers", "body", m.NsxtManagers); err != nil {
		return err
	}

	for i := 0; i < len(m.NsxtManagers); i++ {
		if swag.IsZero(m.NsxtManagers[i]) { // not required
			continue
		}

		if m.NsxtManagers[i] != nil {
			if err := m.NsxtManagers[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("nsxtManagers" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("nsxtManagers" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *SddcNsxtSpec) validateOverLayTransportZone(formats strfmt.Registry) error {
	if swag.IsZero(m.OverLayTransportZone) { // not required
		return nil
	}

	if m.OverLayTransportZone != nil {
		if err := m.OverLayTransportZone.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("overLayTransportZone")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("overLayTransportZone")
			}
			return err
		}
	}

	return nil
}

func (m *SddcNsxtSpec) validateRootNsxtManagerPassword(formats strfmt.Registry) error {

	if err := validate.Required("rootNsxtManagerPassword", "body", m.RootNsxtManagerPassword); err != nil {
		return err
	}

	return nil
}

func (m *SddcNsxtSpec) validateVip(formats strfmt.Registry) error {

	if err := validate.Required("vip", "body", m.Vip); err != nil {
		return err
	}

	return nil
}

func (m *SddcNsxtSpec) validateVipFqdn(formats strfmt.Registry) error {

	if err := validate.Required("vipFqdn", "body", m.VipFqdn); err != nil {
		return err
	}

	return nil
}

func (m *SddcNsxtSpec) validateVlanTransportZone(formats strfmt.Registry) error {
	if swag.IsZero(m.VlanTransportZone) { // not required
		return nil
	}

	if m.VlanTransportZone != nil {
		if err := m.VlanTransportZone.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vlanTransportZone")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("vlanTransportZone")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this sddc nsxt spec based on the context it is used
func (m *SddcNsxtSpec) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateIPAddressPoolSpec(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNsxtManagers(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateOverLayTransportZone(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVlanTransportZone(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SddcNsxtSpec) contextValidateIPAddressPoolSpec(ctx context.Context, formats strfmt.Registry) error {

	if m.IPAddressPoolSpec != nil {
		if err := m.IPAddressPoolSpec.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ipAddressPoolSpec")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("ipAddressPoolSpec")
			}
			return err
		}
	}

	return nil
}

func (m *SddcNsxtSpec) contextValidateNsxtManagers(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.NsxtManagers); i++ {

		if m.NsxtManagers[i] != nil {
			if err := m.NsxtManagers[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("nsxtManagers" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("nsxtManagers" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *SddcNsxtSpec) contextValidateOverLayTransportZone(ctx context.Context, formats strfmt.Registry) error {

	if m.OverLayTransportZone != nil {
		if err := m.OverLayTransportZone.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("overLayTransportZone")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("overLayTransportZone")
			}
			return err
		}
	}

	return nil
}

func (m *SddcNsxtSpec) contextValidateVlanTransportZone(ctx context.Context, formats strfmt.Registry) error {

	if m.VlanTransportZone != nil {
		if err := m.VlanTransportZone.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vlanTransportZone")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("vlanTransportZone")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SddcNsxtSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SddcNsxtSpec) UnmarshalBinary(b []byte) error {
	var res SddcNsxtSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
