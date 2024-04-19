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

// NSXTTransportZoneInfo NSX transport zone representation.
//
// swagger:model NsxtTransportZoneInfo
type NSXTTransportZoneInfo struct {

	// ID of the NSX transport zone
	ID string `json:"id,omitempty"`

	// Name of the NSX transport zone
	Name string `json:"name,omitempty"`

	// List of tags associated with the NSX transport zone
	NSXTTags *NSXTTag `json:"nsxtTags,omitempty"`

	// Type of the NSX transport zone
	Type string `json:"type,omitempty"`
}

// Validate validates this Nsxt transport zone info
func (m *NSXTTransportZoneInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateNSXTTags(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NSXTTransportZoneInfo) validateNSXTTags(formats strfmt.Registry) error {
	if swag.IsZero(m.NSXTTags) { // not required
		return nil
	}

	if m.NSXTTags != nil {
		if err := m.NSXTTags.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("nsxtTags")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("nsxtTags")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this Nsxt transport zone info based on the context it is used
func (m *NSXTTransportZoneInfo) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateNSXTTags(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NSXTTransportZoneInfo) contextValidateNSXTTags(ctx context.Context, formats strfmt.Registry) error {

	if m.NSXTTags != nil {

		if swag.IsZero(m.NSXTTags) { // not required
			return nil
		}

		if err := m.NSXTTags.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("nsxtTags")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("nsxtTags")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *NSXTTransportZoneInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NSXTTransportZoneInfo) UnmarshalBinary(b []byte) error {
	var res NSXTTransportZoneInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}