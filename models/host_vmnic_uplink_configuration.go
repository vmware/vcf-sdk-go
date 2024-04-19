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

// HostVmnicUplinkConfiguration This specification contains vmnic-uplink network configuration of host in a cluster
//
// swagger:model HostVmnicUplinkConfiguration
type HostVmnicUplinkConfiguration struct {

	// VmNic ID
	ID string `json:"id,omitempty"`

	// Uplink of the VDS associated with the vmnic
	Uplink string `json:"uplink,omitempty"`

	// VDS name
	VdsName string `json:"vdsName,omitempty"`
}

// Validate validates this host vmnic uplink configuration
func (m *HostVmnicUplinkConfiguration) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this host vmnic uplink configuration based on context it is used
func (m *HostVmnicUplinkConfiguration) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *HostVmnicUplinkConfiguration) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HostVmnicUplinkConfiguration) UnmarshalBinary(b []byte) error {
	var res HostVmnicUplinkConfiguration
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}