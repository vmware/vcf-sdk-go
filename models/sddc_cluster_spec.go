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

// SDDCClusterSpec Spec contains parameters for vCenter Cluster
//
// swagger:model SddcClusterSpec
type SDDCClusterSpec struct {

	// vCenter cluster EVC mode
	ClusterEvcMode string `json:"clusterEvcMode,omitempty"`

	// vCenter Cluster Name
	// Required: true
	ClusterName *string `json:"clusterName"`

	// Host failures to tolerate
	// Maximum: 3
	// Minimum: 0
	HostFailuresToTolerate *int32 `json:"hostFailuresToTolerate,omitempty"`

	// Hour at which the scheduled compliance check runs
	// Maximum: 23
	// Minimum: 0
	HostProfileComplianceCheckHour *int32 `json:"hostProfileComplianceCheckHour,omitempty"`

	// Minute at which the scheduled compliance check runs
	// Maximum: 59
	// Minimum: 0
	HostProfileComplianceCheckMinute *int32 `json:"hostProfileComplianceCheckMinute,omitempty"`

	// vCenter Cluster Host IDs
	Hosts []string `json:"hosts"`

	// Cluster Personality Name
	PersonalityName string `json:"personalityName,omitempty"`

	// Resource Pool Specs
	ResourcePoolSpecs []*ResourcePoolSpec `json:"resourcePoolSpecs"`

	// Virtual Machine folders map
	// Example: One among:MANAGEMENT, NETWORKING
	VMFolders map[string]string `json:"vmFolders,omitempty"`
}

// Validate validates this Sddc cluster spec
func (m *SDDCClusterSpec) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateClusterName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHostFailuresToTolerate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHostProfileComplianceCheckHour(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHostProfileComplianceCheckMinute(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResourcePoolSpecs(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SDDCClusterSpec) validateClusterName(formats strfmt.Registry) error {

	if err := validate.Required("clusterName", "body", m.ClusterName); err != nil {
		return err
	}

	return nil
}

func (m *SDDCClusterSpec) validateHostFailuresToTolerate(formats strfmt.Registry) error {
	if swag.IsZero(m.HostFailuresToTolerate) { // not required
		return nil
	}

	if err := validate.MinimumInt("hostFailuresToTolerate", "body", int64(*m.HostFailuresToTolerate), 0, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("hostFailuresToTolerate", "body", int64(*m.HostFailuresToTolerate), 3, false); err != nil {
		return err
	}

	return nil
}

func (m *SDDCClusterSpec) validateHostProfileComplianceCheckHour(formats strfmt.Registry) error {
	if swag.IsZero(m.HostProfileComplianceCheckHour) { // not required
		return nil
	}

	if err := validate.MinimumInt("hostProfileComplianceCheckHour", "body", int64(*m.HostProfileComplianceCheckHour), 0, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("hostProfileComplianceCheckHour", "body", int64(*m.HostProfileComplianceCheckHour), 23, false); err != nil {
		return err
	}

	return nil
}

func (m *SDDCClusterSpec) validateHostProfileComplianceCheckMinute(formats strfmt.Registry) error {
	if swag.IsZero(m.HostProfileComplianceCheckMinute) { // not required
		return nil
	}

	if err := validate.MinimumInt("hostProfileComplianceCheckMinute", "body", int64(*m.HostProfileComplianceCheckMinute), 0, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("hostProfileComplianceCheckMinute", "body", int64(*m.HostProfileComplianceCheckMinute), 59, false); err != nil {
		return err
	}

	return nil
}

func (m *SDDCClusterSpec) validateResourcePoolSpecs(formats strfmt.Registry) error {
	if swag.IsZero(m.ResourcePoolSpecs) { // not required
		return nil
	}

	for i := 0; i < len(m.ResourcePoolSpecs); i++ {
		if swag.IsZero(m.ResourcePoolSpecs[i]) { // not required
			continue
		}

		if m.ResourcePoolSpecs[i] != nil {
			if err := m.ResourcePoolSpecs[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("resourcePoolSpecs" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("resourcePoolSpecs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this Sddc cluster spec based on the context it is used
func (m *SDDCClusterSpec) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateResourcePoolSpecs(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SDDCClusterSpec) contextValidateResourcePoolSpecs(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ResourcePoolSpecs); i++ {

		if m.ResourcePoolSpecs[i] != nil {
			if err := m.ResourcePoolSpecs[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("resourcePoolSpecs" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("resourcePoolSpecs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *SDDCClusterSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SDDCClusterSpec) UnmarshalBinary(b []byte) error {
	var res SDDCClusterSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
