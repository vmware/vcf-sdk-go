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

// NsxTSpec This specification contains the parameters required to install and configure NSX-T in a workload domain
//
// swagger:model NsxTSpec
type NsxTSpec struct {

	// NSX manager form factor
	FormFactor string `json:"formFactor,omitempty"`

	// The IP address pool specification
	IPAddressPoolSpec *IPAddressPoolSpec `json:"ipAddressPoolSpec,omitempty"`

	// NSX license value
	LicenseKey string `json:"licenseKey,omitempty"`

	// NSX manager admin password (basic auth and SSH)
	NsxManagerAdminPassword string `json:"nsxManagerAdminPassword,omitempty"`

	// NSX manager Audit password
	NsxManagerAuditPassword string `json:"nsxManagerAuditPassword,omitempty"`

	// Specification details of the NSX Manager virtual machine
	// Required: true
	NsxManagerSpecs []*NsxManagerSpec `json:"nsxManagerSpecs"`

	// Virtual IP address which would act as proxy/alias for NSX-T Managers
	// Required: true
	Vip *string `json:"vip"`

	// FQDN for VIP so that common SSL certificates can be installed across all managers
	// Required: true
	VipFqdn *string `json:"vipFqdn"`
}

// Validate validates this nsx t spec
func (m *NsxTSpec) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateIPAddressPoolSpec(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNsxManagerSpecs(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVip(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVipFqdn(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NsxTSpec) validateIPAddressPoolSpec(formats strfmt.Registry) error {
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

func (m *NsxTSpec) validateNsxManagerSpecs(formats strfmt.Registry) error {

	if err := validate.Required("nsxManagerSpecs", "body", m.NsxManagerSpecs); err != nil {
		return err
	}

	for i := 0; i < len(m.NsxManagerSpecs); i++ {
		if swag.IsZero(m.NsxManagerSpecs[i]) { // not required
			continue
		}

		if m.NsxManagerSpecs[i] != nil {
			if err := m.NsxManagerSpecs[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("nsxManagerSpecs" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("nsxManagerSpecs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *NsxTSpec) validateVip(formats strfmt.Registry) error {

	if err := validate.Required("vip", "body", m.Vip); err != nil {
		return err
	}

	return nil
}

func (m *NsxTSpec) validateVipFqdn(formats strfmt.Registry) error {

	if err := validate.Required("vipFqdn", "body", m.VipFqdn); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this nsx t spec based on the context it is used
func (m *NsxTSpec) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateIPAddressPoolSpec(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNsxManagerSpecs(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NsxTSpec) contextValidateIPAddressPoolSpec(ctx context.Context, formats strfmt.Registry) error {

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

func (m *NsxTSpec) contextValidateNsxManagerSpecs(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.NsxManagerSpecs); i++ {

		if m.NsxManagerSpecs[i] != nil {
			if err := m.NsxManagerSpecs[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("nsxManagerSpecs" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("nsxManagerSpecs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *NsxTSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NsxTSpec) UnmarshalBinary(b []byte) error {
	var res NsxTSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}