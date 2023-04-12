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

// VcenterSpec This specification contains the installation and configuration of vCenter in a workload domain
//
// swagger:model VcenterSpec
type VcenterSpec struct {

	// vCenter datacenter name
	// Required: true
	DatacenterName *string `json:"datacenterName"`

	// Name of the vCenter virtual machine
	// Required: true
	Name *string `json:"name"`

	// Network spec details of the vCenter virtual machine
	// Required: true
	NetworkDetailsSpec *NetworkDetailsSpec `json:"networkDetailsSpec"`

	// vCenter root shell password
	// Required: true
	RootPassword *string `json:"rootPassword"`

	// VCenter storage size
	// Example: One among:lstorage, xlstorage
	StorageSize string `json:"storageSize,omitempty"`

	// VCenter VM size
	// Example: One among:xlarge, large, medium, small, tiny
	VMSize string `json:"vmSize,omitempty"`
}

// Validate validates this vcenter spec
func (m *VcenterSpec) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDatacenterName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNetworkDetailsSpec(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRootPassword(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VcenterSpec) validateDatacenterName(formats strfmt.Registry) error {

	if err := validate.Required("datacenterName", "body", m.DatacenterName); err != nil {
		return err
	}

	return nil
}

func (m *VcenterSpec) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *VcenterSpec) validateNetworkDetailsSpec(formats strfmt.Registry) error {

	if err := validate.Required("networkDetailsSpec", "body", m.NetworkDetailsSpec); err != nil {
		return err
	}

	if m.NetworkDetailsSpec != nil {
		if err := m.NetworkDetailsSpec.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("networkDetailsSpec")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("networkDetailsSpec")
			}
			return err
		}
	}

	return nil
}

func (m *VcenterSpec) validateRootPassword(formats strfmt.Registry) error {

	if err := validate.Required("rootPassword", "body", m.RootPassword); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this vcenter spec based on the context it is used
func (m *VcenterSpec) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateNetworkDetailsSpec(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VcenterSpec) contextValidateNetworkDetailsSpec(ctx context.Context, formats strfmt.Registry) error {

	if m.NetworkDetailsSpec != nil {
		if err := m.NetworkDetailsSpec.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("networkDetailsSpec")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("networkDetailsSpec")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *VcenterSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VcenterSpec) UnmarshalBinary(b []byte) error {
	var res VcenterSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
