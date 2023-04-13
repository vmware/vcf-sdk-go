// Code generated by go-swagger; DO NOT EDIT.

// Copyright © 2023 VMware, Inc. All Rights Reserved.
// Code licenced under: BSD-2

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NSXTManagerCluster Nsxt Manager Cluster Upgrade Resource
//
// swagger:model NsxtManagerCluster
type NSXTManagerCluster struct {

	// Manager cluster id
	ID string `json:"id,omitempty"`

	// Manager cluster name
	Name string `json:"name,omitempty"`

	// Current resource health status
	ResourceHealth string `json:"resourceHealth,omitempty"`

	// Upgrade status with respect to the bundle
	UpgradeStatus string `json:"upgradeStatus,omitempty"`

	// The current version of the manager cluster.If a partial upgrade is done, this will be the least version among the managers in the manager cluster
	Version string `json:"version,omitempty"`
}

// Validate validates this Nsxt manager cluster
func (m *NSXTManagerCluster) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this Nsxt manager cluster based on context it is used
func (m *NSXTManagerCluster) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *NSXTManagerCluster) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NSXTManagerCluster) UnmarshalBinary(b []byte) error {
	var res NSXTManagerCluster
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
