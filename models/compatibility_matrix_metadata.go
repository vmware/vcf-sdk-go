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

// CompatibilityMatrixMetadata Model for metadata of compatibility matrix
//
// swagger:model CompatibilityMatrixMetadata
type CompatibilityMatrixMetadata struct {

	// Error message for the compatibility matrix content
	ErrorMessage string `json:"errorMessage,omitempty"`

	// existence of the compatibility matrix content
	IsMissing bool `json:"isMissing,omitempty"`

	// staleness of the compatibility matrix content
	IsStale bool `json:"isStale,omitempty"`

	// Last modified date of the compatibility matrix content
	LastModifiedDate int64 `json:"lastModifiedDate,omitempty"`

	// Info message for the compatibility matrix content
	Message string `json:"message,omitempty"`

	// Warning message for the compatibility matrix content
	WarningMessage string `json:"warningMessage,omitempty"`
}

// Validate validates this compatibility matrix metadata
func (m *CompatibilityMatrixMetadata) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this compatibility matrix metadata based on context it is used
func (m *CompatibilityMatrixMetadata) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CompatibilityMatrixMetadata) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CompatibilityMatrixMetadata) UnmarshalBinary(b []byte) error {
	var res CompatibilityMatrixMetadata
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
