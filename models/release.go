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
	"github.com/go-openapi/validate"
)

// Release Model for releases with their description and product version
//
// swagger:model Release
type Release struct {

	// Release bill of materials
	// Required: true
	Bom []*ProductVersion `json:"bom"`

	// Release description with all major features.
	// Required: true
	Description *string `json:"description"`

	// Release eol information e.g. 2020-06-08T02:20:15.844Z in yyyy-MM-dd'T'HH:mm:ss[.SSS]XXX ISO 8601 format
	Eol string `json:"eol,omitempty"`

	// [Deprecated] Whether bundle is applicable or not.
	IsApplicable bool `json:"isApplicable,omitempty"`

	// Minimum compatible VCF version, used to represent compatibility of SDDC Manager and VMware BOM components
	// Required: true
	MinCompatibleVcfVersion *string `json:"minCompatibleVcfVersion"`

	// [Deprecated] Incompatibility reason for not applicable.
	NotApplicableReason string `json:"notApplicableReason,omitempty"`

	// List of patch bundles in this release
	PatchBundles []*PatchBundle `json:"patchBundles"`

	// Name of the product e.g. "VCF"
	// Required: true
	Product *string `json:"product"`

	// Release date e.g. 2020-06-08T02:20:15.844Z in yyyy-MM-dd'T'HH:mm:ss[.SSS]XXX ISO 8601 format
	// Required: true
	ReleaseDate *string `json:"releaseDate"`

	// Release sku specific patch and bill of materials
	Sku []*SkuBomDetails `json:"sku"`

	// Collection of release updates
	Updates []*ReleaseUpdate `json:"updates"`

	// Version of the release
	// Required: true
	Version *string `json:"version"`
}

// Validate validates this release
func (m *Release) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBom(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMinCompatibleVcfVersion(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePatchBundles(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProduct(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReleaseDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSku(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdates(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVersion(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Release) validateBom(formats strfmt.Registry) error {

	if err := validate.Required("bom", "body", m.Bom); err != nil {
		return err
	}

	for i := 0; i < len(m.Bom); i++ {
		if swag.IsZero(m.Bom[i]) { // not required
			continue
		}

		if m.Bom[i] != nil {
			if err := m.Bom[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("bom" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("bom" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Release) validateDescription(formats strfmt.Registry) error {

	if err := validate.Required("description", "body", m.Description); err != nil {
		return err
	}

	return nil
}

func (m *Release) validateMinCompatibleVcfVersion(formats strfmt.Registry) error {

	if err := validate.Required("minCompatibleVcfVersion", "body", m.MinCompatibleVcfVersion); err != nil {
		return err
	}

	return nil
}

func (m *Release) validatePatchBundles(formats strfmt.Registry) error {
	if swag.IsZero(m.PatchBundles) { // not required
		return nil
	}

	for i := 0; i < len(m.PatchBundles); i++ {
		if swag.IsZero(m.PatchBundles[i]) { // not required
			continue
		}

		if m.PatchBundles[i] != nil {
			if err := m.PatchBundles[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("patchBundles" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("patchBundles" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Release) validateProduct(formats strfmt.Registry) error {

	if err := validate.Required("product", "body", m.Product); err != nil {
		return err
	}

	return nil
}

func (m *Release) validateReleaseDate(formats strfmt.Registry) error {

	if err := validate.Required("releaseDate", "body", m.ReleaseDate); err != nil {
		return err
	}

	return nil
}

func (m *Release) validateSku(formats strfmt.Registry) error {
	if swag.IsZero(m.Sku) { // not required
		return nil
	}

	for i := 0; i < len(m.Sku); i++ {
		if swag.IsZero(m.Sku[i]) { // not required
			continue
		}

		if m.Sku[i] != nil {
			if err := m.Sku[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("sku" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("sku" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Release) validateUpdates(formats strfmt.Registry) error {
	if swag.IsZero(m.Updates) { // not required
		return nil
	}

	for i := 0; i < len(m.Updates); i++ {
		if swag.IsZero(m.Updates[i]) { // not required
			continue
		}

		if m.Updates[i] != nil {
			if err := m.Updates[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("updates" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("updates" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Release) validateVersion(formats strfmt.Registry) error {

	if err := validate.Required("version", "body", m.Version); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this release based on the context it is used
func (m *Release) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateBom(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePatchBundles(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSku(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUpdates(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Release) contextValidateBom(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Bom); i++ {

		if m.Bom[i] != nil {
			if err := m.Bom[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("bom" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("bom" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Release) contextValidatePatchBundles(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.PatchBundles); i++ {

		if m.PatchBundles[i] != nil {
			if err := m.PatchBundles[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("patchBundles" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("patchBundles" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Release) contextValidateSku(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Sku); i++ {

		if m.Sku[i] != nil {
			if err := m.Sku[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("sku" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("sku" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Release) contextValidateUpdates(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Updates); i++ {

		if m.Updates[i] != nil {
			if err := m.Updates[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("updates" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("updates" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *Release) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Release) UnmarshalBinary(b []byte) error {
	var res Release
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
