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

// SupportBundleIncludeItems support bundle include items
//
// swagger:model SupportBundleIncludeItems
type SupportBundleIncludeItems struct {

	// Perform SOS Health checks.
	HealthCheck bool `json:"healthCheck,omitempty"`

	// Collect Vcf Summary Reports.
	SummaryReport bool `json:"summaryReport,omitempty"`
}

// Validate validates this support bundle include items
func (m *SupportBundleIncludeItems) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this support bundle include items based on context it is used
func (m *SupportBundleIncludeItems) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SupportBundleIncludeItems) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SupportBundleIncludeItems) UnmarshalBinary(b []byte) error {
	var res SupportBundleIncludeItems
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
