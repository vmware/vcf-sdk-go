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

// HostDetail Host Configuration
//
// swagger:model HostDetail
type HostDetail struct {

	// Availability zone to which the host belongs when the cluster is stretched
	AzName string `json:"azName,omitempty"`

	// Host name of the vSphere host
	HostName string `json:"hostName,omitempty"`

	// Network configuration of the host
	HostNetworkConfiguration *HostNetworkConfiguration `json:"hostNetworkConfiguration,omitempty"`

	// ID of the host
	ID string `json:"id,omitempty"`
}

// Validate validates this host detail
func (m *HostDetail) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateHostNetworkConfiguration(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HostDetail) validateHostNetworkConfiguration(formats strfmt.Registry) error {
	if swag.IsZero(m.HostNetworkConfiguration) { // not required
		return nil
	}

	if m.HostNetworkConfiguration != nil {
		if err := m.HostNetworkConfiguration.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("hostNetworkConfiguration")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("hostNetworkConfiguration")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this host detail based on the context it is used
func (m *HostDetail) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateHostNetworkConfiguration(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HostDetail) contextValidateHostNetworkConfiguration(ctx context.Context, formats strfmt.Registry) error {

	if m.HostNetworkConfiguration != nil {

		if swag.IsZero(m.HostNetworkConfiguration) { // not required
			return nil
		}

		if err := m.HostNetworkConfiguration.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("hostNetworkConfiguration")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("hostNetworkConfiguration")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *HostDetail) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HostDetail) UnmarshalBinary(b []byte) error {
	var res HostDetail
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
