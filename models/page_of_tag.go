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
)

// PageOfTag Represents a page of elements of a single type
//
// swagger:model PageOfTag
type PageOfTag struct {

	// The list of elements included in this page
	Elements []*Tag `json:"elements"`

	// Pageable elements pagination metadata information
	PageMetadata *PageMetadata `json:"pageMetadata,omitempty"`
}

// Validate validates this page of tag
func (m *PageOfTag) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateElements(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePageMetadata(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PageOfTag) validateElements(formats strfmt.Registry) error {
	if swag.IsZero(m.Elements) { // not required
		return nil
	}

	for i := 0; i < len(m.Elements); i++ {
		if swag.IsZero(m.Elements[i]) { // not required
			continue
		}

		if m.Elements[i] != nil {
			if err := m.Elements[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("elements" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("elements" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PageOfTag) validatePageMetadata(formats strfmt.Registry) error {
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

// ContextValidate validate this page of tag based on the context it is used
func (m *PageOfTag) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateElements(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePageMetadata(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PageOfTag) contextValidateElements(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Elements); i++ {

		if m.Elements[i] != nil {
			if err := m.Elements[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("elements" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("elements" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PageOfTag) contextValidatePageMetadata(ctx context.Context, formats strfmt.Registry) error {

	if m.PageMetadata != nil {
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
func (m *PageOfTag) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PageOfTag) UnmarshalBinary(b []byte) error {
	var res PageOfTag
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
