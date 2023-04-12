// Code generated by go-swagger; DO NOT EDIT.

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

// PageOfTask Represents a page of elements of a single type
//
// swagger:model PageOfTask
type PageOfTask struct {

	// The list of elements included in this page
	Elements []*Task `json:"elements"`

	// Pageable elements pagination metadata information
	PageMetadata *PageMetadata `json:"pageMetadata,omitempty"`
}

// Validate validates this page of task
func (m *PageOfTask) Validate(formats strfmt.Registry) error {
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

func (m *PageOfTask) validateElements(formats strfmt.Registry) error {
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

func (m *PageOfTask) validatePageMetadata(formats strfmt.Registry) error {
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

// ContextValidate validate this page of task based on the context it is used
func (m *PageOfTask) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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

func (m *PageOfTask) contextValidateElements(ctx context.Context, formats strfmt.Registry) error {

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

func (m *PageOfTask) contextValidatePageMetadata(ctx context.Context, formats strfmt.Registry) error {

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
func (m *PageOfTask) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PageOfTask) UnmarshalBinary(b []byte) error {
	var res PageOfTask
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
