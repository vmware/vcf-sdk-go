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

// EdgeCluster EdgeCluster representation
//
// swagger:model EdgeCluster
type EdgeCluster struct {

	// List of clusters associated with the edge cluster
	Clusters []*ClusterReference `json:"clusters"`

	// List of edge nodes associated with the edge cluster
	EdgeNodes []*EdgeNodeReference `json:"edgeNodes"`

	// ID of the egde cluster
	ID string `json:"id,omitempty"`

	// Whether or not this edge cluster's tier 0 is managed by system
	IsTier0ManagedBySystem bool `json:"isTier0ManagedBySystem,omitempty"`

	// Name of the edge cluster
	Name string `json:"name,omitempty"`

	// NSX cluster associated with the edge cluster
	NSXTCluster *NsxTClusterReference `json:"nsxtCluster,omitempty"`

	// Whether or not host/TEP network checks were done for this edge cluster
	SkipTepRoutabilityCheck bool `json:"skipTepRoutabilityCheck,omitempty"`
}

// Validate validates this edge cluster
func (m *EdgeCluster) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateClusters(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEdgeNodes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNSXTCluster(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EdgeCluster) validateClusters(formats strfmt.Registry) error {
	if swag.IsZero(m.Clusters) { // not required
		return nil
	}

	for i := 0; i < len(m.Clusters); i++ {
		if swag.IsZero(m.Clusters[i]) { // not required
			continue
		}

		if m.Clusters[i] != nil {
			if err := m.Clusters[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("clusters" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("clusters" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *EdgeCluster) validateEdgeNodes(formats strfmt.Registry) error {
	if swag.IsZero(m.EdgeNodes) { // not required
		return nil
	}

	for i := 0; i < len(m.EdgeNodes); i++ {
		if swag.IsZero(m.EdgeNodes[i]) { // not required
			continue
		}

		if m.EdgeNodes[i] != nil {
			if err := m.EdgeNodes[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("edgeNodes" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("edgeNodes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *EdgeCluster) validateNSXTCluster(formats strfmt.Registry) error {
	if swag.IsZero(m.NSXTCluster) { // not required
		return nil
	}

	if m.NSXTCluster != nil {
		if err := m.NSXTCluster.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("nsxtCluster")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("nsxtCluster")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this edge cluster based on the context it is used
func (m *EdgeCluster) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateClusters(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateEdgeNodes(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNSXTCluster(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EdgeCluster) contextValidateClusters(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Clusters); i++ {

		if m.Clusters[i] != nil {

			if swag.IsZero(m.Clusters[i]) { // not required
				return nil
			}

			if err := m.Clusters[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("clusters" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("clusters" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *EdgeCluster) contextValidateEdgeNodes(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.EdgeNodes); i++ {

		if m.EdgeNodes[i] != nil {

			if swag.IsZero(m.EdgeNodes[i]) { // not required
				return nil
			}

			if err := m.EdgeNodes[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("edgeNodes" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("edgeNodes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *EdgeCluster) contextValidateNSXTCluster(ctx context.Context, formats strfmt.Registry) error {

	if m.NSXTCluster != nil {

		if swag.IsZero(m.NSXTCluster) { // not required
			return nil
		}

		if err := m.NSXTCluster.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("nsxtCluster")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("nsxtCluster")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *EdgeCluster) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EdgeCluster) UnmarshalBinary(b []byte) error {
	var res EdgeCluster
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
