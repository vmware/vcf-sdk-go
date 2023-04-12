// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NsxtUpgradeOptions NSXT Upgrade options
//
// swagger:model NsxtUpgradeOptions
type NsxtUpgradeOptions struct {

	// Flag for disabling/enabling parallel upgrade of edge transportnode clusters
	IsEdgeClustersUpgradeParallel bool `json:"isEdgeClustersUpgradeParallel,omitempty"`

	// Flag for performing edge-only upgrade
	IsEdgeOnlyUpgrade bool `json:"isEdgeOnlyUpgrade,omitempty"`

	// Flag for disabling/enabling parallel upgrade of host transportnode clusters
	IsHostClustersUpgradeParallel bool `json:"isHostClustersUpgradeParallel,omitempty"`
}

// Validate validates this nsxt upgrade options
func (m *NsxtUpgradeOptions) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this nsxt upgrade options based on context it is used
func (m *NsxtUpgradeOptions) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *NsxtUpgradeOptions) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NsxtUpgradeOptions) UnmarshalBinary(b []byte) error {
	var res NsxtUpgradeOptions
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}