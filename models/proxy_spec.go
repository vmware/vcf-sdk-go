// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2023 VMware, Inc.
// SPDX-License-Identifier: BSD-2-Clause

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ProxySpec Spec contains parameters for proxy settings
//
// swagger:model ProxySpec
type ProxySpec struct {

	// IP address/FQDN of proxy server
	// Example: 10.0.0.250
	Host string `json:"host,omitempty"`

	// Port of proxy server
	// Example: 3128
	Port int32 `json:"port,omitempty"`
}

// Validate validates this proxy spec
func (m *ProxySpec) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this proxy spec based on context it is used
func (m *ProxySpec) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ProxySpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ProxySpec) UnmarshalBinary(b []byte) error {
	var res ProxySpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
