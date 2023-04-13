// Code generated by go-swagger; DO NOT EDIT.

// Copyright © 2023 VMware, Inc. All Rights Reserved.
// Code licenced under: BSD-2

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ErrorCode error code
//
// swagger:model ErrorCode
type ErrorCode struct {

	// bundle name
	BundleName string `json:"bundleName,omitempty"`

	// code
	Code string `json:"code,omitempty"`

	// component
	Component string `json:"component,omitempty"`
}

// Validate validates this error code
func (m *ErrorCode) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this error code based on context it is used
func (m *ErrorCode) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ErrorCode) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ErrorCode) UnmarshalBinary(b []byte) error {
	var res ErrorCode
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
