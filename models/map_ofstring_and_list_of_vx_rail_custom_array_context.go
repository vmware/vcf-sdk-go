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
)

// MapOfstringAndListOfVxRailCustomArrayContext map ofstring and list of vx rail custom array context
//
// swagger:model MapOfstringAndListOfVxRailCustomArrayContext
type MapOfstringAndListOfVxRailCustomArrayContext map[string]List

// Validate validates this map ofstring and list of vx rail custom array context
func (m MapOfstringAndListOfVxRailCustomArrayContext) Validate(formats strfmt.Registry) error {
	var res []error

	for k := range m {

		if err := m[k].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName(k)
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName(k)
			}
			return err
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validate this map ofstring and list of vx rail custom array context based on the context it is used
func (m MapOfstringAndListOfVxRailCustomArrayContext) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	for k := range m {

		if err := m[k].ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName(k)
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName(k)
			}
			return err
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}