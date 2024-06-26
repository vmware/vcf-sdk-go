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

// NsxTManager NSX Manager representation
//
// swagger:model NsxTManager
type NsxTManager struct {

	// FQDN of the manager
	Fqdn string `json:"fqdn,omitempty"`

	// ID of the manager
	ID string `json:"id,omitempty"`

	// IP address of the manager
	IPAddress string `json:"ipAddress,omitempty"`

	// NSX Manager VM name in vCenter
	Name string `json:"name,omitempty"`
}

// Validate validates this nsx t manager
func (m *NsxTManager) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this nsx t manager based on context it is used
func (m *NsxTManager) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *NsxTManager) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NsxTManager) UnmarshalBinary(b []byte) error {
	var res NsxTManager
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
