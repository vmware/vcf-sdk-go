// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// HealthSummary health summary
//
// swagger:model HealthSummary
type HealthSummary struct {

	// Indicate if bundle is available in file system.
	BundleAvailable string `json:"bundleAvailable,omitempty"`

	// Name of the Support Bundle.
	BundleName string `json:"bundleName,omitempty"`

	// CompletionTimestamp.
	CompletionTimestamp string `json:"completionTimestamp,omitempty"`

	// CreationTimestamp.
	CreationTimestamp string `json:"creationTimestamp,omitempty"`

	// Name of the Operation.
	Description string `json:"description,omitempty"`

	// Id of the Support Bundle task.
	ID string `json:"id,omitempty"`

	// Health summary collection status.
	// Example: One among: IN_PROGRESS, PENDING, COMPLETED_WITH_SUCCESS, COMPLETED_WITH_FAILURE
	Status string `json:"status,omitempty"`
}

// Validate validates this health summary
func (m *HealthSummary) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this health summary based on context it is used
func (m *HealthSummary) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *HealthSummary) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HealthSummary) UnmarshalBinary(b []byte) error {
	var res HealthSummary
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
