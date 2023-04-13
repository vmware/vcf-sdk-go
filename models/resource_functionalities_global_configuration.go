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

// ResourceFunctionalitiesGlobalConfiguration Defines a resource functionality caller specification
//
// swagger:model ResourceFunctionalitiesGlobalConfiguration
type ResourceFunctionalitiesGlobalConfiguration struct {

	// The localized error message
	ErrorMessage string `json:"errorMessage,omitempty"`

	// Global flag representing if all of the VMware Cloud Foundation operations are allowed or not
	IsAllowed bool `json:"isAllowed,omitempty"`
}

// Validate validates this resource functionalities global configuration
func (m *ResourceFunctionalitiesGlobalConfiguration) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this resource functionalities global configuration based on context it is used
func (m *ResourceFunctionalitiesGlobalConfiguration) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ResourceFunctionalitiesGlobalConfiguration) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ResourceFunctionalitiesGlobalConfiguration) UnmarshalBinary(b []byte) error {
	var res ResourceFunctionalitiesGlobalConfiguration
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
