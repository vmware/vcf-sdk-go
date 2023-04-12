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

// EdgeClusterShrinkageSpec This specification contains the parameters required to shrink a NSX-T edge cluster.
//
// swagger:model EdgeClusterShrinkageSpec
type EdgeClusterShrinkageSpec struct {

	// List of VCF Edge Node ID's to be removed for shrinkage.
	// Required: true
	EdgeNodeIds []string `json:"edgeNodeIds"`
}

// Validate validates this edge cluster shrinkage spec
func (m *EdgeClusterShrinkageSpec) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEdgeNodeIds(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EdgeClusterShrinkageSpec) validateEdgeNodeIds(formats strfmt.Registry) error {

	if err := validate.Required("edgeNodeIds", "body", m.EdgeNodeIds); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this edge cluster shrinkage spec based on context it is used
func (m *EdgeClusterShrinkageSpec) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *EdgeClusterShrinkageSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EdgeClusterShrinkageSpec) UnmarshalBinary(b []byte) error {
	var res EdgeClusterShrinkageSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
