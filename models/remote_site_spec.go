// Code generated by go-swagger; DO NOT EDIT.

// Copyright © 2023 VMware, Inc. All Rights Reserved.
// Code licenced under: BSD-2

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// RemoteSiteSpec Spec contains parameters for Remote site products
//
// swagger:model RemoteSiteSpec
type RemoteSiteSpec struct {

	// Remote region vCenter address
	PscAddress string `json:"pscAddress,omitempty"`

	// Remote region vCenter SSL thumbprint (SHA256)
	SslThumbprint string `json:"sslThumbprint,omitempty"`

	// Remote region vCenter credentials
	VcCredentials *SDDCCredentials `json:"vcCredentials,omitempty"`
}

// Validate validates this remote site spec
func (m *RemoteSiteSpec) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateVcCredentials(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RemoteSiteSpec) validateVcCredentials(formats strfmt.Registry) error {
	if swag.IsZero(m.VcCredentials) { // not required
		return nil
	}

	if m.VcCredentials != nil {
		if err := m.VcCredentials.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vcCredentials")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("vcCredentials")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this remote site spec based on the context it is used
func (m *RemoteSiteSpec) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateVcCredentials(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RemoteSiteSpec) contextValidateVcCredentials(ctx context.Context, formats strfmt.Registry) error {

	if m.VcCredentials != nil {
		if err := m.VcCredentials.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vcCredentials")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("vcCredentials")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RemoteSiteSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RemoteSiteSpec) UnmarshalBinary(b []byte) error {
	var res RemoteSiteSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
