// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// BackupRetentionPolicy Backup retention policy for SDDC Manager comprising three attributes. Backup files are deleted if they do not satisfy any of the three attributes.
//
// swagger:model BackupRetentionPolicy
type BackupRetentionPolicy struct {

	// This attribute controls the number of daily backup files to retain, measured in days. Range 0 to 30 days. The system will filter the existing backup files, and retain one for every day for the specified number of days, counting back from the most recent backup.
	// Example: 20
	NumberOfDaysOfDailyBackups int32 `json:"numberOfDaysOfDailyBackups,omitempty"`

	// This attribute controls the number of hourly backup files to retain, measured in days. Range 0 to 14 days. The system will filter the existing backup files, and retain one for every hour for the specified number of days, counting back from the most recent backup.
	// Example: 10
	NumberOfDaysOfHourlyBackups int32 `json:"numberOfDaysOfHourlyBackups,omitempty"`

	// This attribute controls the number of recent backup files to retain. Range 1 to 600 backup files.
	// Example: 15
	// Required: true
	NumberOfMostRecentBackups *int32 `json:"numberOfMostRecentBackups"`
}

// Validate validates this backup retention policy
func (m *BackupRetentionPolicy) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateNumberOfMostRecentBackups(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BackupRetentionPolicy) validateNumberOfMostRecentBackups(formats strfmt.Registry) error {

	if err := validate.Required("numberOfMostRecentBackups", "body", m.NumberOfMostRecentBackups); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this backup retention policy based on context it is used
func (m *BackupRetentionPolicy) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *BackupRetentionPolicy) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BackupRetentionPolicy) UnmarshalBinary(b []byte) error {
	var res BackupRetentionPolicy
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}