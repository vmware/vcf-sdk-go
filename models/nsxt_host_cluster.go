// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2023 VMware, Inc.
// SPDX-License-Identifier: BSD-2-Clause

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

// NSXTHostCluster NSX Host Cluster Upgrade Resource
//
// swagger:model NsxtHostCluster
type NSXTHostCluster struct {

	// Available Hardware Support Managers for Firmware upgrade
	AvailableHardwareSupportManagers []*HardwareSupportPackages `json:"availableHardwareSupportManagers"`

	// Configured Hardware Support Managers for Firmware upgrade
	ConfiguredHardwareSupportManagers []*HardwareSupportPackage `json:"configuredHardwareSupportManagers"`

	// Id of the domain the Host cluster belongs to
	DomainID string `json:"domainId,omitempty"`

	// Host cluster VCF cluster id
	ID string `json:"id,omitempty"`

	// Cluster name
	Name string `json:"name,omitempty"`

	// Flag to determine if primary cluster for vLCM
	PrimaryCluster bool `json:"primaryCluster,omitempty"`

	// Current resource health status
	ResourceHealth string `json:"resourceHealth,omitempty"`

	// Total number of hosts in the Host cluster
	TotalUnits int32 `json:"totalUnits,omitempty"`

	// Upgrade status with respect to the bundle
	UpgradeStatus string `json:"upgradeStatus,omitempty"`

	// Number of hosts which are already upgraded
	UpgradedUnits int32 `json:"upgradedUnits,omitempty"`

	// The current version of the host node cluster.If a partial upgrade is done, this will be the least version among the hosts
	Version string `json:"version,omitempty"`

	// vlcm enabled
	VlcmEnabled bool `json:"vlcmEnabled,omitempty"`
}

// Validate validates this Nsxt host cluster
func (m *NSXTHostCluster) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAvailableHardwareSupportManagers(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateConfiguredHardwareSupportManagers(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NSXTHostCluster) validateAvailableHardwareSupportManagers(formats strfmt.Registry) error {
	if swag.IsZero(m.AvailableHardwareSupportManagers) { // not required
		return nil
	}

	for i := 0; i < len(m.AvailableHardwareSupportManagers); i++ {
		if swag.IsZero(m.AvailableHardwareSupportManagers[i]) { // not required
			continue
		}

		if m.AvailableHardwareSupportManagers[i] != nil {
			if err := m.AvailableHardwareSupportManagers[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("availableHardwareSupportManagers" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("availableHardwareSupportManagers" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *NSXTHostCluster) validateConfiguredHardwareSupportManagers(formats strfmt.Registry) error {
	if swag.IsZero(m.ConfiguredHardwareSupportManagers) { // not required
		return nil
	}

	for i := 0; i < len(m.ConfiguredHardwareSupportManagers); i++ {
		if swag.IsZero(m.ConfiguredHardwareSupportManagers[i]) { // not required
			continue
		}

		if m.ConfiguredHardwareSupportManagers[i] != nil {
			if err := m.ConfiguredHardwareSupportManagers[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("configuredHardwareSupportManagers" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("configuredHardwareSupportManagers" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this Nsxt host cluster based on the context it is used
func (m *NSXTHostCluster) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAvailableHardwareSupportManagers(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateConfiguredHardwareSupportManagers(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NSXTHostCluster) contextValidateAvailableHardwareSupportManagers(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.AvailableHardwareSupportManagers); i++ {

		if m.AvailableHardwareSupportManagers[i] != nil {

			if swag.IsZero(m.AvailableHardwareSupportManagers[i]) { // not required
				return nil
			}

			if err := m.AvailableHardwareSupportManagers[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("availableHardwareSupportManagers" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("availableHardwareSupportManagers" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *NSXTHostCluster) contextValidateConfiguredHardwareSupportManagers(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ConfiguredHardwareSupportManagers); i++ {

		if m.ConfiguredHardwareSupportManagers[i] != nil {

			if swag.IsZero(m.ConfiguredHardwareSupportManagers[i]) { // not required
				return nil
			}

			if err := m.ConfiguredHardwareSupportManagers[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("configuredHardwareSupportManagers" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("configuredHardwareSupportManagers" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *NSXTHostCluster) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NSXTHostCluster) UnmarshalBinary(b []byte) error {
	var res NSXTHostCluster
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
