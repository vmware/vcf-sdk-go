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

// VersionAliasesForBundleComponentType Version Alias representation
//
// swagger:model VersionAliasesForBundleComponentType
type VersionAliasesForBundleComponentType struct {

	// Bundle Component Type
	// Example: One among: VCENTER, PSC, NSX_T_MANAGER, NSX_MANAGER, ESX_HOST, VCF_VRA_UPGRADE, VCF_VRLI_UPGRADE, VCF_VROPS_UPGRADE, VCF_VRSLCM_UPGRADE, SDDC_MANAGER_VCF, LCM, SDDC_MANAGER, SDDC_MANAGER_UI, OPERATIONS_MANAGER, COMMON_SERVICES, SOLUTIONS_MANAGER, MULTI_SITE_SERVICE,
	// Required: true
	BundleComponentType *string `json:"bundleComponentType"`

	// Version Aliases
	// Required: true
	VersionAliases []*BaseAlias `json:"versionAliases"`
}

// Validate validates this version aliases for bundle component type
func (m *VersionAliasesForBundleComponentType) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBundleComponentType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVersionAliases(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VersionAliasesForBundleComponentType) validateBundleComponentType(formats strfmt.Registry) error {

	if err := validate.Required("bundleComponentType", "body", m.BundleComponentType); err != nil {
		return err
	}

	return nil
}

func (m *VersionAliasesForBundleComponentType) validateVersionAliases(formats strfmt.Registry) error {

	if err := validate.Required("versionAliases", "body", m.VersionAliases); err != nil {
		return err
	}

	for i := 0; i < len(m.VersionAliases); i++ {
		if swag.IsZero(m.VersionAliases[i]) { // not required
			continue
		}

		if m.VersionAliases[i] != nil {
			if err := m.VersionAliases[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("versionAliases" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("versionAliases" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this version aliases for bundle component type based on the context it is used
func (m *VersionAliasesForBundleComponentType) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateVersionAliases(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VersionAliasesForBundleComponentType) contextValidateVersionAliases(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.VersionAliases); i++ {

		if m.VersionAliases[i] != nil {
			if err := m.VersionAliases[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("versionAliases" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("versionAliases" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *VersionAliasesForBundleComponentType) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VersionAliasesForBundleComponentType) UnmarshalBinary(b []byte) error {
	var res VersionAliasesForBundleComponentType
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
