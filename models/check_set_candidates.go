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

// CheckSetCandidates Represents a check-set id that can be used to run checks for a resource
//
// swagger:model CheckSetCandidates
type CheckSetCandidates struct {

	// Id of the check-set
	CheckSetID string `json:"checkSetId,omitempty"`

	// Name of the check-set
	CheckSetName string `json:"checkSetName,omitempty"`

	// Type of the check-set
	CheckSetType string `json:"checkSetType,omitempty"`
}

// Validate validates this check set candidates
func (m *CheckSetCandidates) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this check set candidates based on context it is used
func (m *CheckSetCandidates) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CheckSetCandidates) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CheckSetCandidates) UnmarshalBinary(b []byte) error {
	var res CheckSetCandidates
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
