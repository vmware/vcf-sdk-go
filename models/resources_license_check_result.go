// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2023 VMware, Inc.
// SPDX-License-Identifier: BSD-2-Clause

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ResourcesLicenseCheckResult Resource License check result
//
// swagger:model ResourcesLicenseCheckResult
type ResourcesLicenseCheckResult struct {

	// Task end timestamp
	EndTimestamp string `json:"endTimestamp,omitempty"`

	// ID of the resources license check task
	ID string `json:"id,omitempty"`

	// List of licensing infos of resources
	ResourceLicensingInfos []*ResourceLicensingInfo `json:"resourceLicensingInfos"`

	// Task start timestamp
	StartTimestamp string `json:"startTimestamp,omitempty"`

	// Task status
	// Example: One among: PENDING, IN_PROGRESS, In Progress, SUCCESSFUL, Successful, FAILED, Failed, CANCELLED, Cancelled, COMPLETED_WITH_WARNING, SKIPPED
	Status string `json:"status,omitempty"`
}

// Validate validates this resources license check result
func (m *ResourcesLicenseCheckResult) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateResourceLicensingInfos(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ResourcesLicenseCheckResult) validateResourceLicensingInfos(formats strfmt.Registry) error {
	if swag.IsZero(m.ResourceLicensingInfos) { // not required
		return nil
	}

	for i := 0; i < len(m.ResourceLicensingInfos); i++ {
		if swag.IsZero(m.ResourceLicensingInfos[i]) { // not required
			continue
		}

		if m.ResourceLicensingInfos[i] != nil {
			if err := m.ResourceLicensingInfos[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("resourceLicensingInfos" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("resourceLicensingInfos" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this resources license check result based on the context it is used
func (m *ResourcesLicenseCheckResult) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateResourceLicensingInfos(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ResourcesLicenseCheckResult) contextValidateResourceLicensingInfos(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ResourceLicensingInfos); i++ {

		if m.ResourceLicensingInfos[i] != nil {

			if swag.IsZero(m.ResourceLicensingInfos[i]) { // not required
				return nil
			}

			if err := m.ResourceLicensingInfos[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("resourceLicensingInfos" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("resourceLicensingInfos" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ResourcesLicenseCheckResult) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ResourcesLicenseCheckResult) UnmarshalBinary(b []byte) error {
	var res ResourcesLicenseCheckResult
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
