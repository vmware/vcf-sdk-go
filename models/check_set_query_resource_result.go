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

// CheckSetQueryResourceResult Represents a result of a check-set query call for a given resource
//
// swagger:model CheckSetQueryResourceResult
type CheckSetQueryResourceResult struct {

	// Check set candidates to select from
	CheckSets []*CheckSetCandidates `json:"checkSets"`

	// Domain information about the domain the resource belongs to
	Domain *ResourceDomainInfo `json:"domain,omitempty"`

	// extra context
	ExtraContext interface{} `json:"extraContext,omitempty"`

	// Id of the resource
	ResourceID string `json:"resourceId,omitempty"`

	// Name of the resource
	ResourceName string `json:"resourceName,omitempty"`

	// Type of the resource
	ResourceType string `json:"resourceType,omitempty"`

	// Current version of the resource
	ResourceVersion string `json:"resourceVersion,omitempty"`
}

// Validate validates this check set query resource result
func (m *CheckSetQueryResourceResult) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCheckSets(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDomain(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CheckSetQueryResourceResult) validateCheckSets(formats strfmt.Registry) error {
	if swag.IsZero(m.CheckSets) { // not required
		return nil
	}

	for i := 0; i < len(m.CheckSets); i++ {
		if swag.IsZero(m.CheckSets[i]) { // not required
			continue
		}

		if m.CheckSets[i] != nil {
			if err := m.CheckSets[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("checkSets" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("checkSets" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *CheckSetQueryResourceResult) validateDomain(formats strfmt.Registry) error {
	if swag.IsZero(m.Domain) { // not required
		return nil
	}

	if m.Domain != nil {
		if err := m.Domain.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("domain")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("domain")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this check set query resource result based on the context it is used
func (m *CheckSetQueryResourceResult) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCheckSets(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDomain(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CheckSetQueryResourceResult) contextValidateCheckSets(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.CheckSets); i++ {

		if m.CheckSets[i] != nil {

			if swag.IsZero(m.CheckSets[i]) { // not required
				return nil
			}

			if err := m.CheckSets[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("checkSets" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("checkSets" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *CheckSetQueryResourceResult) contextValidateDomain(ctx context.Context, formats strfmt.Registry) error {

	if m.Domain != nil {

		if swag.IsZero(m.Domain) { // not required
			return nil
		}

		if err := m.Domain.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("domain")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("domain")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CheckSetQueryResourceResult) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CheckSetQueryResourceResult) UnmarshalBinary(b []byte) error {
	var res CheckSetQueryResourceResult
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
