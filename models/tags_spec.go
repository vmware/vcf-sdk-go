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

// TagsSpec Represents tags spec
//
// swagger:model TagsSpec
type TagsSpec struct {

	// Tag IDs
	// Required: true
	TagIds []string `json:"tagIds"`
}

// Validate validates this tags spec
func (m *TagsSpec) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTagIds(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TagsSpec) validateTagIds(formats strfmt.Registry) error {

	if err := validate.Required("tagIds", "body", m.TagIds); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this tags spec based on context it is used
func (m *TagsSpec) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *TagsSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TagsSpec) UnmarshalBinary(b []byte) error {
	var res TagsSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
