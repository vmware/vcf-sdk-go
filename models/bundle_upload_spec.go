// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2023 VMware, Inc.
// SPDX-License-Identifier: BSD-2-Clause

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

// BundleUploadSpec Bundle Upload Specification. This specification is used in the Bundle Upload API
//
// swagger:model BundleUploadSpec
type BundleUploadSpec struct {

	// Bundle Upload File Path
	// Required: true
	BundleFilePath *string `json:"bundleFilePath"`

	// [Deprecated] Path to the software compatibility sets file
	CompatibilitySetsFilePath string `json:"compatibilitySetsFilePath,omitempty"`

	// Bundle Upload Manifest File Path
	// Required: true
	ManifestFilePath *string `json:"manifestFilePath"`

	// Partner extensions for bundle upload
	PartnerExtensionSpec *PartnerExtensionSpec `json:"partnerExtensionSpec,omitempty"`

	// Bundle Upload Signature File Path
	SignatureFilePath string `json:"signatureFilePath,omitempty"`
}

// Validate validates this bundle upload spec
func (m *BundleUploadSpec) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBundleFilePath(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateManifestFilePath(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePartnerExtensionSpec(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BundleUploadSpec) validateBundleFilePath(formats strfmt.Registry) error {

	if err := validate.Required("bundleFilePath", "body", m.BundleFilePath); err != nil {
		return err
	}

	return nil
}

func (m *BundleUploadSpec) validateManifestFilePath(formats strfmt.Registry) error {

	if err := validate.Required("manifestFilePath", "body", m.ManifestFilePath); err != nil {
		return err
	}

	return nil
}

func (m *BundleUploadSpec) validatePartnerExtensionSpec(formats strfmt.Registry) error {
	if swag.IsZero(m.PartnerExtensionSpec) { // not required
		return nil
	}

	if m.PartnerExtensionSpec != nil {
		if err := m.PartnerExtensionSpec.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("partnerExtensionSpec")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("partnerExtensionSpec")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this bundle upload spec based on the context it is used
func (m *BundleUploadSpec) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidatePartnerExtensionSpec(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BundleUploadSpec) contextValidatePartnerExtensionSpec(ctx context.Context, formats strfmt.Registry) error {

	if m.PartnerExtensionSpec != nil {

		if swag.IsZero(m.PartnerExtensionSpec) { // not required
			return nil
		}

		if err := m.PartnerExtensionSpec.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("partnerExtensionSpec")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("partnerExtensionSpec")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *BundleUploadSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BundleUploadSpec) UnmarshalBinary(b []byte) error {
	var res BundleUploadSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
