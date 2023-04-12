// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
)

// MapOfstringAndListOfAsyncPatch map ofstring and list of async patch
//
// swagger:model MapOfstringAndListOfAsyncPatch
type MapOfstringAndListOfAsyncPatch map[string]List

// Validate validates this map ofstring and list of async patch
func (m MapOfstringAndListOfAsyncPatch) Validate(formats strfmt.Registry) error {
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

// ContextValidate validate this map ofstring and list of async patch based on the context it is used
func (m MapOfstringAndListOfAsyncPatch) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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
