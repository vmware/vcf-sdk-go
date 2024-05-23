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

// VRA Spec contains parameters of a VMware Aria Automation instance
//
// swagger:model Vra
type VRA struct {

	// The ID of the VMware Aria Automation instance
	ID string `json:"id,omitempty"`

	// The Fully Qualified Domain Name of the VMware Aria Automation load balancer
	// Example: vra-lb.vrack.vsphere.local
	LoadBalancerFqdn string `json:"loadBalancerFqdn,omitempty"`

	// The IP address of the VMware Aria Automation load balancer
	// Example: 10.0.0.15
	LoadBalancerIPAddress string `json:"loadBalancerIpAddress,omitempty"`

	// The nodes of the VMware Aria Automation instance
	Nodes []*VrealizeProductNode `json:"nodes"`

	// The status of the VMware Aria Automation instance
	// Example: ACTIVE, ERROR
	Status string `json:"status,omitempty"`

	// The version of the VMware Aria Automation instance
	// Example: 8.1.0-13036238
	Version string `json:"version,omitempty"`
}

// Validate validates this Vra
func (m *VRA) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateNodes(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VRA) validateNodes(formats strfmt.Registry) error {
	if swag.IsZero(m.Nodes) { // not required
		return nil
	}

	for i := 0; i < len(m.Nodes); i++ {
		if swag.IsZero(m.Nodes[i]) { // not required
			continue
		}

		if m.Nodes[i] != nil {
			if err := m.Nodes[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("nodes" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("nodes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this Vra based on the context it is used
func (m *VRA) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateNodes(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VRA) contextValidateNodes(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Nodes); i++ {

		if m.Nodes[i] != nil {

			if swag.IsZero(m.Nodes[i]) { // not required
				return nil
			}

			if err := m.Nodes[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("nodes" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("nodes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *VRA) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VRA) UnmarshalBinary(b []byte) error {
	var res VRA
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
