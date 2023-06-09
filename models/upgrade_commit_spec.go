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

// UpgradeCommitSpec Upgrade Commit Specification
//
// swagger:model UpgradeCommitSpec
type UpgradeCommitSpec struct {

	// Upgrade Scheduled Time
	ScheduledTimestamp string `json:"scheduledTimestamp,omitempty"`

	// Flag for enabling Upgrade Now. If true, scheduledTimestamp is ignored
	UpgradeNow bool `json:"upgradeNow,omitempty"`
}

// Validate validates this upgrade commit spec
func (m *UpgradeCommitSpec) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this upgrade commit spec based on context it is used
func (m *UpgradeCommitSpec) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *UpgradeCommitSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UpgradeCommitSpec) UnmarshalBinary(b []byte) error {
	var res UpgradeCommitSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
