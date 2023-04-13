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

// HealthSummaryConfig health summary config
//
// swagger:model HealthSummaryConfig
type HealthSummaryConfig struct {

	// Run SOS operations, even if there is a Workload running.
	Force bool `json:"force,omitempty"`

	// Skip known_hosts file for HealthSummary.
	SkipKnownHostCheck bool `json:"skipKnownHostCheck,omitempty"`
}

// Validate validates this health summary config
func (m *HealthSummaryConfig) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this health summary config based on context it is used
func (m *HealthSummaryConfig) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *HealthSummaryConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HealthSummaryConfig) UnmarshalBinary(b []byte) error {
	var res HealthSummaryConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
