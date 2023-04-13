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
)

// BundleUpdateSpec Bundle Update Specification.
//
// swagger:model BundleUpdateSpec
type BundleUpdateSpec struct {

	// Bundle Download Specification.
	BundleDownloadSpec *BundleDownloadSpec `json:"bundleDownloadSpec,omitempty"`

	// Path to the software compatibility sets file
	CompatibilitySetsFilePath string `json:"compatibilitySetsFilePath,omitempty"`
}

// Validate validates this bundle update spec
func (m *BundleUpdateSpec) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBundleDownloadSpec(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BundleUpdateSpec) validateBundleDownloadSpec(formats strfmt.Registry) error {
	if swag.IsZero(m.BundleDownloadSpec) { // not required
		return nil
	}

	if m.BundleDownloadSpec != nil {
		if err := m.BundleDownloadSpec.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("bundleDownloadSpec")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("bundleDownloadSpec")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this bundle update spec based on the context it is used
func (m *BundleUpdateSpec) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateBundleDownloadSpec(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BundleUpdateSpec) contextValidateBundleDownloadSpec(ctx context.Context, formats strfmt.Registry) error {

	if m.BundleDownloadSpec != nil {
		if err := m.BundleDownloadSpec.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("bundleDownloadSpec")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("bundleDownloadSpec")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *BundleUpdateSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BundleUpdateSpec) UnmarshalBinary(b []byte) error {
	var res BundleUpdateSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
