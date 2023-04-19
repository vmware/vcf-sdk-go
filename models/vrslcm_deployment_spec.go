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

// VRSLCMDeploymentSpec Spec contains parameters for creating a new vRealize Suite Lifecycle Manager instance
//
// swagger:model VrslcmDeploymentSpec
type VRSLCMDeploymentSpec struct {

	// The password for an admin API/UI user of vRealize Suite Lifecycle Manager
	// Required: true
	APIPassword *string `json:"apiPassword"`

	// Fully Qualified Domain Name
	// Example: vrslcm.vrack.vsphere.local
	// Required: true
	Fqdn *string `json:"fqdn"`

	// The IP to use for deploying a new standalone Tier-1 router in NSX-T. This router will be used subsequently for vrealize load balancers.
	// Required: true
	NSXTStandaloneTier1IP *string `json:"nsxtStandaloneTier1Ip"`

	// The password for a root user of vRealize Suite Lifecycle Manager appliance
	// Required: true
	SSHPassword *string `json:"sshPassword"`
}

// Validate validates this Vrslcm deployment spec
func (m *VRSLCMDeploymentSpec) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAPIPassword(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFqdn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNSXTStandaloneTier1IP(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSSHPassword(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VRSLCMDeploymentSpec) validateAPIPassword(formats strfmt.Registry) error {

	if err := validate.Required("apiPassword", "body", m.APIPassword); err != nil {
		return err
	}

	return nil
}

func (m *VRSLCMDeploymentSpec) validateFqdn(formats strfmt.Registry) error {

	if err := validate.Required("fqdn", "body", m.Fqdn); err != nil {
		return err
	}

	return nil
}

func (m *VRSLCMDeploymentSpec) validateNSXTStandaloneTier1IP(formats strfmt.Registry) error {

	if err := validate.Required("nsxtStandaloneTier1Ip", "body", m.NSXTStandaloneTier1IP); err != nil {
		return err
	}

	return nil
}

func (m *VRSLCMDeploymentSpec) validateSSHPassword(formats strfmt.Registry) error {

	if err := validate.Required("sshPassword", "body", m.SSHPassword); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this Vrslcm deployment spec based on context it is used
func (m *VRSLCMDeploymentSpec) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *VRSLCMDeploymentSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VRSLCMDeploymentSpec) UnmarshalBinary(b []byte) error {
	var res VRSLCMDeploymentSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
