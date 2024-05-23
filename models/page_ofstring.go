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
)

// PageOfstring Represents a page of elements of a single type
//
// swagger:model PageOfstring
type PageOfstring struct {

	// The list of elements included in this page
	Elements []string `json:"elements"`

	// Pageable elements pagination metadata information
	PageMetadata *PageMetadata `json:"pageMetadata,omitempty"`
}

// Validate validates this page ofstring
func (m *PageOfstring) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePageMetadata(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PageOfstring) validatePageMetadata(formats strfmt.Registry) error {
	if swag.IsZero(m.PageMetadata) { // not required
		return nil
	}

	if m.PageMetadata != nil {
		if err := m.PageMetadata.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("pageMetadata")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("pageMetadata")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this page ofstring based on the context it is used
func (m *PageOfstring) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidatePageMetadata(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PageOfstring) contextValidatePageMetadata(ctx context.Context, formats strfmt.Registry) error {

	if m.PageMetadata != nil {

		if swag.IsZero(m.PageMetadata) { // not required
			return nil
		}

		if err := m.PageMetadata.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("pageMetadata")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("pageMetadata")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PageOfstring) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PageOfstring) UnmarshalBinary(b []byte) error {
	var res PageOfstring
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
