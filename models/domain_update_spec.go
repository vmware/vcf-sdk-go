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
)

// DomainUpdateSpec Contains the parameters required to perform update operation on domain
//
// swagger:model DomainUpdateSpec
type DomainUpdateSpec struct {

	// Cluster Specification for the domain
	ClusterSpec *ClusterSpec `json:"clusterSpec,omitempty"`

	// Enable deletion for the domain
	MarkForDeletion bool `json:"markForDeletion,omitempty"`

	// Name of the domain
	Name string `json:"name,omitempty"`

	// NSX Specification for the domain
	NsxTSpec *NsxTSpec `json:"nsxTSpec,omitempty"`
}

// Validate validates this domain update spec
func (m *DomainUpdateSpec) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateClusterSpec(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNsxTSpec(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DomainUpdateSpec) validateClusterSpec(formats strfmt.Registry) error {
	if swag.IsZero(m.ClusterSpec) { // not required
		return nil
	}

	if m.ClusterSpec != nil {
		if err := m.ClusterSpec.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("clusterSpec")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("clusterSpec")
			}
			return err
		}
	}

	return nil
}

func (m *DomainUpdateSpec) validateNsxTSpec(formats strfmt.Registry) error {
	if swag.IsZero(m.NsxTSpec) { // not required
		return nil
	}

	if m.NsxTSpec != nil {
		if err := m.NsxTSpec.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("nsxTSpec")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("nsxTSpec")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this domain update spec based on the context it is used
func (m *DomainUpdateSpec) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateClusterSpec(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNsxTSpec(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DomainUpdateSpec) contextValidateClusterSpec(ctx context.Context, formats strfmt.Registry) error {

	if m.ClusterSpec != nil {

		if swag.IsZero(m.ClusterSpec) { // not required
			return nil
		}

		if err := m.ClusterSpec.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("clusterSpec")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("clusterSpec")
			}
			return err
		}
	}

	return nil
}

func (m *DomainUpdateSpec) contextValidateNsxTSpec(ctx context.Context, formats strfmt.Registry) error {

	if m.NsxTSpec != nil {

		if swag.IsZero(m.NsxTSpec) { // not required
			return nil
		}

		if err := m.NsxTSpec.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("nsxTSpec")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("nsxTSpec")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DomainUpdateSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DomainUpdateSpec) UnmarshalBinary(b []byte) error {
	var res DomainUpdateSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
