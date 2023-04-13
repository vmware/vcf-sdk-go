// Code generated by go-swagger; DO NOT EDIT.

// Copyright © 2023 VMware, Inc. All Rights Reserved.
// Code licenced under: BSD-2

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

// NSXTEdgeClusterUpgradeSpec Host transport node cluster upgrade input
//
// swagger:model NsxtEdgeClusterUpgradeSpec
type NSXTEdgeClusterUpgradeSpec struct {

	// Resource ID of the edge transport node cluster
	// Required: true
	EdgeClusterID *string `json:"edgeClusterId"`

	// disable/enable parallel upgrade of edges within the cluster
	EdgeParallelUpgrade bool `json:"edgeParallelUpgrade,omitempty"`
}

// Validate validates this Nsxt edge cluster upgrade spec
func (m *NSXTEdgeClusterUpgradeSpec) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEdgeClusterID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NSXTEdgeClusterUpgradeSpec) validateEdgeClusterID(formats strfmt.Registry) error {

	if err := validate.Required("edgeClusterId", "body", m.EdgeClusterID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this Nsxt edge cluster upgrade spec based on context it is used
func (m *NSXTEdgeClusterUpgradeSpec) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *NSXTEdgeClusterUpgradeSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NSXTEdgeClusterUpgradeSpec) UnmarshalBinary(b []byte) error {
	var res NSXTEdgeClusterUpgradeSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
