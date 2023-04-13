// Code generated by go-swagger; DO NOT EDIT.

// Copyright © 2023 VMware, Inc. All Rights Reserved.
// Code licenced under: BSD-2

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

// SkuBomDetails Model for sku with their name, description, patchBundles and bom
//
// swagger:model SkuBomDetails
type SkuBomDetails struct {

	// Sku specific bill of materials
	// Required: true
	Bom []*ProductVersion `json:"bom"`

	// Description to be shown in release page
	Description string `json:"description,omitempty"`

	// SKU name
	Name string `json:"name,omitempty"`

	// List of patch bundles in this release
	SkuSpecificPatchBundles []*PatchBundle `json:"skuSpecificPatchBundles"`
}

// Validate validates this sku bom details
func (m *SkuBomDetails) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBom(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSkuSpecificPatchBundles(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SkuBomDetails) validateBom(formats strfmt.Registry) error {

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

func (m *SkuBomDetails) validateSkuSpecificPatchBundles(formats strfmt.Registry) error {
	if swag.IsZero(m.SkuSpecificPatchBundles) { // not required
		return nil
	}

	for i := 0; i < len(m.SkuSpecificPatchBundles); i++ {
		if swag.IsZero(m.SkuSpecificPatchBundles[i]) { // not required
			continue
		}

		if m.SkuSpecificPatchBundles[i] != nil {
			if err := m.SkuSpecificPatchBundles[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("skuSpecificPatchBundles" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("skuSpecificPatchBundles" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this sku bom details based on the context it is used
func (m *SkuBomDetails) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateBom(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSkuSpecificPatchBundles(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SkuBomDetails) contextValidateBom(ctx context.Context, formats strfmt.Registry) error {

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

func (m *SkuBomDetails) contextValidateSkuSpecificPatchBundles(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.SkuSpecificPatchBundles); i++ {

		if m.SkuSpecificPatchBundles[i] != nil {
			if err := m.SkuSpecificPatchBundles[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("skuSpecificPatchBundles" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("skuSpecificPatchBundles" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *SkuBomDetails) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SkuBomDetails) UnmarshalBinary(b []byte) error {
	var res SkuBomDetails
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
