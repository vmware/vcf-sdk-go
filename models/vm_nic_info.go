// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// VMNicInfo VMNics attached to an ESXi host reference
//
// swagger:model VMNicInfo
type VMNicInfo struct {

	// Status of VMNic if active or inactive
	IsActive bool `json:"isActive,omitempty"`

	// Status of VMNic if auto negotiate is supported or not
	IsAutoNegotiateSupported bool `json:"isAutoNegotiateSupported,omitempty"`

	// Status of VMNic if in use or available
	IsInUse bool `json:"isInUse,omitempty"`

	// VMNic link speed in MB
	LinkSpeedMB int32 `json:"linkSpeedMB,omitempty"`

	// Name of the VMNic
	Name string `json:"name,omitempty"`
}

// Validate validates this VM nic info
func (m *VMNicInfo) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this VM nic info based on context it is used
func (m *VMNicInfo) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *VMNicInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VMNicInfo) UnmarshalBinary(b []byte) error {
	var res VMNicInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
