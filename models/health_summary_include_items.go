// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// HealthSummaryIncludeItems health summary include items
//
// swagger:model HealthSummaryIncludeItems
type HealthSummaryIncludeItems struct {

	// Collect VCF LCM Upgrade Pre-check Reports, Default value is False
	PrecheckReport bool `json:"precheckReport,omitempty"`

	// Collect Vcf Summary Reports
	SummaryReport bool `json:"summaryReport,omitempty"`
}

// Validate validates this health summary include items
func (m *HealthSummaryIncludeItems) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this health summary include items based on context it is used
func (m *HealthSummaryIncludeItems) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *HealthSummaryIncludeItems) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HealthSummaryIncludeItems) UnmarshalBinary(b []byte) error {
	var res HealthSummaryIncludeItems
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
