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

// VMNic This specification contains vmnic configurations for vSphere host
//
// swagger:model VmNic
type VMNic struct {

	// VmNic ID of vSphere host to be associated with VDS, once added to cluster
	ID string `json:"id,omitempty"`

	// This flag determines if the vmnic must be on N-VDS
	MoveToNvds bool `json:"moveToNvds,omitempty"`

	// Uplink to be associated with vmnic
	Uplink string `json:"uplink,omitempty"`

	// VDS name to associate with vSphere host
	VdsName string `json:"vdsName,omitempty"`
}

// Validate validates this Vm nic
func (m *VMNic) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this Vm nic based on context it is used
func (m *VMNic) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *VMNic) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VMNic) UnmarshalBinary(b []byte) error {
	var res VMNic
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
