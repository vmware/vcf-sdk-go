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

// TagAssignmentResult Represents a Tag Assignment Result
//
// swagger:model TagAssignmentResult
type TagAssignmentResult struct {

	// Default Error messages for Assign/Detach Tags operation
	DefaultErrorMessages []string `json:"defaultErrorMessages"`

	// Success value of Assign/Detach Tags operation
	Success bool `json:"success,omitempty"`
}

// Validate validates this tag assignment result
func (m *TagAssignmentResult) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this tag assignment result based on context it is used
func (m *TagAssignmentResult) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *TagAssignmentResult) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TagAssignmentResult) UnmarshalBinary(b []byte) error {
	var res TagAssignmentResult
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
