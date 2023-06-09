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

// BundleDownloadSpec Bundle Download Specification. This specification gets used in the Bundle Download API
//
// swagger:model BundleDownloadSpec
type BundleDownloadSpec struct {

	// Flag for enabling Download Now. If true, scheduledTimestamp is ignored
	DownloadNow bool `json:"downloadNow,omitempty"`

	// Bundle Download Scheduled Time
	ScheduledTimestamp string `json:"scheduledTimestamp,omitempty"`
}

// Validate validates this bundle download spec
func (m *BundleDownloadSpec) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this bundle download spec based on context it is used
func (m *BundleDownloadSpec) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *BundleDownloadSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BundleDownloadSpec) UnmarshalBinary(b []byte) error {
	var res BundleDownloadSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
