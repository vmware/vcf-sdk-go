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

// EdgeClusterUpdateSpec This specification contains the parameters required to perform an update operation on an NSX-T edge cluster. The nested expansion and shrinkage specifications need to be populated in a mutually exclusive manner
//
// swagger:model EdgeClusterUpdateSpec
type EdgeClusterUpdateSpec struct {

	// Parameters required to perform edge cluster expansion, optional otherwise
	EdgeClusterExpansionSpec *EdgeClusterExpansionSpec `json:"edgeClusterExpansionSpec,omitempty"`

	// Parameters required to perform edge cluster shrinkage, optional otherwise
	EdgeClusterShrinkageSpec *EdgeClusterShrinkageSpec `json:"edgeClusterShrinkageSpec,omitempty"`

	// Edge cluster operation Type.
	// Example: One among: EXPANSION, SHRINKAGE
	// Required: true
	Operation *string `json:"operation"`
}

// Validate validates this edge cluster update spec
func (m *EdgeClusterUpdateSpec) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEdgeClusterExpansionSpec(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEdgeClusterShrinkageSpec(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOperation(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EdgeClusterUpdateSpec) validateEdgeClusterExpansionSpec(formats strfmt.Registry) error {
	if swag.IsZero(m.EdgeClusterExpansionSpec) { // not required
		return nil
	}

	if m.EdgeClusterExpansionSpec != nil {
		if err := m.EdgeClusterExpansionSpec.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("edgeClusterExpansionSpec")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("edgeClusterExpansionSpec")
			}
			return err
		}
	}

	return nil
}

func (m *EdgeClusterUpdateSpec) validateEdgeClusterShrinkageSpec(formats strfmt.Registry) error {
	if swag.IsZero(m.EdgeClusterShrinkageSpec) { // not required
		return nil
	}

	if m.EdgeClusterShrinkageSpec != nil {
		if err := m.EdgeClusterShrinkageSpec.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("edgeClusterShrinkageSpec")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("edgeClusterShrinkageSpec")
			}
			return err
		}
	}

	return nil
}

func (m *EdgeClusterUpdateSpec) validateOperation(formats strfmt.Registry) error {

	if err := validate.Required("operation", "body", m.Operation); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this edge cluster update spec based on the context it is used
func (m *EdgeClusterUpdateSpec) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateEdgeClusterExpansionSpec(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateEdgeClusterShrinkageSpec(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EdgeClusterUpdateSpec) contextValidateEdgeClusterExpansionSpec(ctx context.Context, formats strfmt.Registry) error {

	if m.EdgeClusterExpansionSpec != nil {
		if err := m.EdgeClusterExpansionSpec.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("edgeClusterExpansionSpec")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("edgeClusterExpansionSpec")
			}
			return err
		}
	}

	return nil
}

func (m *EdgeClusterUpdateSpec) contextValidateEdgeClusterShrinkageSpec(ctx context.Context, formats strfmt.Registry) error {

	if m.EdgeClusterShrinkageSpec != nil {
		if err := m.EdgeClusterShrinkageSpec.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("edgeClusterShrinkageSpec")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("edgeClusterShrinkageSpec")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *EdgeClusterUpdateSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EdgeClusterUpdateSpec) UnmarshalBinary(b []byte) error {
	var res EdgeClusterUpdateSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
