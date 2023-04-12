// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// EdgeClusterNsxtEntityCriterion Represents a criterion for querying the Edge Cluster
//
// swagger:model EdgeClusterNsxtEntityCriterion
type EdgeClusterNsxtEntityCriterion struct {

	// Arguments required for a particular criterion
	Arguments map[string]string `json:"arguments,omitempty"`

	// Description of the criterion
	Description string `json:"description,omitempty"`

	// Name of the criterion
	// Example: One among: TIER_0_GATEWAY_ASSOCIATED_WITH_EDGE_CLUSTER, TIER_1_GATEWAY_ASSOCIATED_WITH_EDGE_CLUSTER
	Name string `json:"name,omitempty"`
}

// Validate validates this edge cluster nsxt entity criterion
func (m *EdgeClusterNsxtEntityCriterion) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this edge cluster nsxt entity criterion based on context it is used
func (m *EdgeClusterNsxtEntityCriterion) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *EdgeClusterNsxtEntityCriterion) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EdgeClusterNsxtEntityCriterion) UnmarshalBinary(b []byte) error {
	var res EdgeClusterNsxtEntityCriterion
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
