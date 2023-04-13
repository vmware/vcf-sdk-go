// Code generated by go-swagger; DO NOT EDIT.

// Copyright © 2023 VMware, Inc. All Rights Reserved.
// Code licenced under: BSD-2

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

// BackupSchedule Backup schedule configuration
//
// swagger:model BackupSchedule
type BackupSchedule struct {

	// List of days of the week to schedule backup
	// Example: One among: SUNDAY, MONDAY, TUESDAY, WEDNESDAY, THURSDAY, FRIDAY, SATURDAY
	DaysOfWeek []string `json:"daysOfWeek"`

	// Backup frequency
	// Example: One among: WEEKLY, HOURLY
	// Required: true
	Frequency *string `json:"frequency"`

	// Hour of the day to schedule backup
	HourOfDay int32 `json:"hourOfDay,omitempty"`

	// Minute of the hour to schedule backup
	MinuteOfHour int32 `json:"minuteOfHour,omitempty"`

	// Resource type to configure backup schedule
	// Example: SDDC_MANAGER
	// Required: true
	ResourceType *string `json:"resourceType"`

	// Backup retention policy
	RetentionPolicy *BackupRetentionPolicy `json:"retentionPolicy,omitempty"`

	// Enable/disable backups on state change. If enabled, SDDC Manager will take a backup after the successful completion of an operation that changes its state. This mode requires that scheduled backups be enabled.
	// Example: True
	TakeBackupOnStateChange bool `json:"takeBackupOnStateChange,omitempty"`

	// Enable/disable scheduled backups
	// Example: True
	TakeScheduledBackups bool `json:"takeScheduledBackups,omitempty"`
}

// Validate validates this backup schedule
func (m *BackupSchedule) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFrequency(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResourceType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRetentionPolicy(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BackupSchedule) validateFrequency(formats strfmt.Registry) error {

	if err := validate.Required("frequency", "body", m.Frequency); err != nil {
		return err
	}

	return nil
}

func (m *BackupSchedule) validateResourceType(formats strfmt.Registry) error {

	if err := validate.Required("resourceType", "body", m.ResourceType); err != nil {
		return err
	}

	return nil
}

func (m *BackupSchedule) validateRetentionPolicy(formats strfmt.Registry) error {
	if swag.IsZero(m.RetentionPolicy) { // not required
		return nil
	}

	if m.RetentionPolicy != nil {
		if err := m.RetentionPolicy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("retentionPolicy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("retentionPolicy")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this backup schedule based on the context it is used
func (m *BackupSchedule) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateRetentionPolicy(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BackupSchedule) contextValidateRetentionPolicy(ctx context.Context, formats strfmt.Registry) error {

	if m.RetentionPolicy != nil {
		if err := m.RetentionPolicy.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("retentionPolicy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("retentionPolicy")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *BackupSchedule) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BackupSchedule) UnmarshalBinary(b []byte) error {
	var res BackupSchedule
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
