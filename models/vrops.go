// Code generated by go-swagger; DO NOT EDIT.

// Copyright © 2023 VMware, Inc. All Rights Reserved.
// Code licenced under: BSD-2

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

// VROPS This specification contains information related to the existing vRealize Operations instance deployment
//
// swagger:model Vrops
type VROPS struct {

	// The ID of the vRealize Operations instance
	// Required: true
	ID *string `json:"id"`

	// Fully Qualified Domain Name for the vRealize Operations load balancer
	// Required: true
	LoadBalancerFqdn *string `json:"loadBalancerFqdn"`

	// IP for the vRealize Operations load balancer
	// Required: true
	LoadBalancerIP *string `json:"loadBalancerIp"`

	// The nodes of the vRealize Operations instance
	// Required: true
	Nodes []*VROPSNode `json:"nodes"`

	// The state of the current product instance
	// Required: true
	Status *string `json:"status"`

	// The version of the vRealize Operations instance
	// Required: true
	Version *string `json:"version"`
}

// Validate validates this Vrops
func (m *VROPS) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLoadBalancerFqdn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLoadBalancerIP(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNodes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVersion(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VROPS) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *VROPS) validateLoadBalancerFqdn(formats strfmt.Registry) error {

	if err := validate.Required("loadBalancerFqdn", "body", m.LoadBalancerFqdn); err != nil {
		return err
	}

	return nil
}

func (m *VROPS) validateLoadBalancerIP(formats strfmt.Registry) error {

	if err := validate.Required("loadBalancerIp", "body", m.LoadBalancerIP); err != nil {
		return err
	}

	return nil
}

func (m *VROPS) validateNodes(formats strfmt.Registry) error {

	if err := validate.Required("nodes", "body", m.Nodes); err != nil {
		return err
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

func (m *VROPS) validateStatus(formats strfmt.Registry) error {

	if err := validate.Required("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

func (m *VROPS) validateVersion(formats strfmt.Registry) error {

	if err := validate.Required("version", "body", m.Version); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this Vrops based on the context it is used
func (m *VROPS) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateNodes(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VROPS) contextValidateNodes(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Nodes); i++ {

		if m.Nodes[i] != nil {
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
func (m *VROPS) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VROPS) UnmarshalBinary(b []byte) error {
	var res VROPS
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
