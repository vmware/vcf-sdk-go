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
	"github.com/go-openapi/validate"
)

// ClusterSpec This specification contains the parameters required to add a cluster to a workload domain
//
// swagger:model ClusterSpec
type ClusterSpec struct {

	// Advanced options used for cluster creation
	AdvancedOptions *AdvancedOptions `json:"advancedOptions,omitempty"`

	// ID of the Cluster Image to be used with the Cluster
	ClusterImageID string `json:"clusterImageId,omitempty"`

	// Cluster storage configuration; e.g. VSAN, NFS, vVol(nfs/iscsi/fc), vSAN Remote
	// Required: true
	DatastoreSpec *DatastoreSpec `json:"datastoreSpec"`

	// List of vSphere host information from the free pool to consume in the workload domain
	// Required: true
	HostSpecs []*HostSpec `json:"hostSpecs"`

	// Name of the new cluster that will be added to the specified workload domain
	// Required: true
	Name *string `json:"name"`

	// Network configuration for the cluster
	// Required: true
	NetworkSpec *NetworkSpec `json:"networkSpec"`

	// Skip thumbprint validation for ESXi and VxRail Manager during add cluster/host operation.
	// This property is deprecated.
	SkipThumbprintValidation bool `json:"skipThumbprintValidation,omitempty"`

	// Contains the VxRail Manager details
	VxRailDetails *VxRailDetails `json:"vxRailDetails,omitempty"`
}

// Validate validates this cluster spec
func (m *ClusterSpec) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAdvancedOptions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDatastoreSpec(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHostSpecs(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNetworkSpec(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVxRailDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClusterSpec) validateAdvancedOptions(formats strfmt.Registry) error {
	if swag.IsZero(m.AdvancedOptions) { // not required
		return nil
	}

	if m.AdvancedOptions != nil {
		if err := m.AdvancedOptions.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("advancedOptions")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("advancedOptions")
			}
			return err
		}
	}

	return nil
}

func (m *ClusterSpec) validateDatastoreSpec(formats strfmt.Registry) error {

	if err := validate.Required("datastoreSpec", "body", m.DatastoreSpec); err != nil {
		return err
	}

	if m.DatastoreSpec != nil {
		if err := m.DatastoreSpec.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("datastoreSpec")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("datastoreSpec")
			}
			return err
		}
	}

	return nil
}

func (m *ClusterSpec) validateHostSpecs(formats strfmt.Registry) error {

	if err := validate.Required("hostSpecs", "body", m.HostSpecs); err != nil {
		return err
	}

	for i := 0; i < len(m.HostSpecs); i++ {
		if swag.IsZero(m.HostSpecs[i]) { // not required
			continue
		}

		if m.HostSpecs[i] != nil {
			if err := m.HostSpecs[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("hostSpecs" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("hostSpecs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ClusterSpec) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *ClusterSpec) validateNetworkSpec(formats strfmt.Registry) error {

	if err := validate.Required("networkSpec", "body", m.NetworkSpec); err != nil {
		return err
	}

	if m.NetworkSpec != nil {
		if err := m.NetworkSpec.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("networkSpec")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("networkSpec")
			}
			return err
		}
	}

	return nil
}

func (m *ClusterSpec) validateVxRailDetails(formats strfmt.Registry) error {
	if swag.IsZero(m.VxRailDetails) { // not required
		return nil
	}

	if m.VxRailDetails != nil {
		if err := m.VxRailDetails.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vxRailDetails")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("vxRailDetails")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this cluster spec based on the context it is used
func (m *ClusterSpec) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAdvancedOptions(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDatastoreSpec(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateHostSpecs(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNetworkSpec(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVxRailDetails(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClusterSpec) contextValidateAdvancedOptions(ctx context.Context, formats strfmt.Registry) error {

	if m.AdvancedOptions != nil {

		if swag.IsZero(m.AdvancedOptions) { // not required
			return nil
		}

		if err := m.AdvancedOptions.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("advancedOptions")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("advancedOptions")
			}
			return err
		}
	}

	return nil
}

func (m *ClusterSpec) contextValidateDatastoreSpec(ctx context.Context, formats strfmt.Registry) error {

	if m.DatastoreSpec != nil {

		if err := m.DatastoreSpec.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("datastoreSpec")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("datastoreSpec")
			}
			return err
		}
	}

	return nil
}

func (m *ClusterSpec) contextValidateHostSpecs(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.HostSpecs); i++ {

		if m.HostSpecs[i] != nil {

			if swag.IsZero(m.HostSpecs[i]) { // not required
				return nil
			}

			if err := m.HostSpecs[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("hostSpecs" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("hostSpecs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ClusterSpec) contextValidateNetworkSpec(ctx context.Context, formats strfmt.Registry) error {

	if m.NetworkSpec != nil {

		if err := m.NetworkSpec.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("networkSpec")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("networkSpec")
			}
			return err
		}
	}

	return nil
}

func (m *ClusterSpec) contextValidateVxRailDetails(ctx context.Context, formats strfmt.Registry) error {

	if m.VxRailDetails != nil {

		if swag.IsZero(m.VxRailDetails) { // not required
			return nil
		}

		if err := m.VxRailDetails.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vxRailDetails")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("vxRailDetails")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ClusterSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClusterSpec) UnmarshalBinary(b []byte) error {
	var res ClusterSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
