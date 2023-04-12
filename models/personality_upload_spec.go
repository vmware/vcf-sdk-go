// Code generated by go-swagger; DO NOT EDIT.

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

// PersonalityUploadSpec Personality upload specification. This spec is used in personality upload API.
//
// swagger:model PersonalityUploadSpec
type PersonalityUploadSpec struct {

	// Personality name
	Name string `json:"name,omitempty"`

	// Personality upload mode
	// Example: RAW, RAW_WITH_FILE_UPLOAD_ID, REFERRED
	// Required: true
	UploadMode *string `json:"uploadMode"`

	// Personality upload spec for upload from RAW files
	UploadSpecRawMode *PersonalityUploadSpecRaw `json:"uploadSpecRawMode,omitempty"`

	// Personality upload spec for upload using file upload id
	UploadSpecRawWithFileUploadIDMode *PersonalityUploadSpecRawWithFileUploadID `json:"uploadSpecRawWithFileUploadIdMode,omitempty"`

	// Personality upload spec for upload from REFERRED vcenter cluster
	UploadSpecReferredMode *PersonalityUploadSpecReferred `json:"uploadSpecReferredMode,omitempty"`
}

// Validate validates this personality upload spec
func (m *PersonalityUploadSpec) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateUploadMode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUploadSpecRawMode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUploadSpecRawWithFileUploadIDMode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUploadSpecReferredMode(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PersonalityUploadSpec) validateUploadMode(formats strfmt.Registry) error {

	if err := validate.Required("uploadMode", "body", m.UploadMode); err != nil {
		return err
	}

	return nil
}

func (m *PersonalityUploadSpec) validateUploadSpecRawMode(formats strfmt.Registry) error {
	if swag.IsZero(m.UploadSpecRawMode) { // not required
		return nil
	}

	if m.UploadSpecRawMode != nil {
		if err := m.UploadSpecRawMode.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("uploadSpecRawMode")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("uploadSpecRawMode")
			}
			return err
		}
	}

	return nil
}

func (m *PersonalityUploadSpec) validateUploadSpecRawWithFileUploadIDMode(formats strfmt.Registry) error {
	if swag.IsZero(m.UploadSpecRawWithFileUploadIDMode) { // not required
		return nil
	}

	if m.UploadSpecRawWithFileUploadIDMode != nil {
		if err := m.UploadSpecRawWithFileUploadIDMode.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("uploadSpecRawWithFileUploadIdMode")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("uploadSpecRawWithFileUploadIdMode")
			}
			return err
		}
	}

	return nil
}

func (m *PersonalityUploadSpec) validateUploadSpecReferredMode(formats strfmt.Registry) error {
	if swag.IsZero(m.UploadSpecReferredMode) { // not required
		return nil
	}

	if m.UploadSpecReferredMode != nil {
		if err := m.UploadSpecReferredMode.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("uploadSpecReferredMode")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("uploadSpecReferredMode")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this personality upload spec based on the context it is used
func (m *PersonalityUploadSpec) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateUploadSpecRawMode(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUploadSpecRawWithFileUploadIDMode(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUploadSpecReferredMode(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PersonalityUploadSpec) contextValidateUploadSpecRawMode(ctx context.Context, formats strfmt.Registry) error {

	if m.UploadSpecRawMode != nil {
		if err := m.UploadSpecRawMode.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("uploadSpecRawMode")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("uploadSpecRawMode")
			}
			return err
		}
	}

	return nil
}

func (m *PersonalityUploadSpec) contextValidateUploadSpecRawWithFileUploadIDMode(ctx context.Context, formats strfmt.Registry) error {

	if m.UploadSpecRawWithFileUploadIDMode != nil {
		if err := m.UploadSpecRawWithFileUploadIDMode.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("uploadSpecRawWithFileUploadIdMode")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("uploadSpecRawWithFileUploadIdMode")
			}
			return err
		}
	}

	return nil
}

func (m *PersonalityUploadSpec) contextValidateUploadSpecReferredMode(ctx context.Context, formats strfmt.Registry) error {

	if m.UploadSpecReferredMode != nil {
		if err := m.UploadSpecReferredMode.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("uploadSpecReferredMode")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("uploadSpecReferredMode")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PersonalityUploadSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PersonalityUploadSpec) UnmarshalBinary(b []byte) error {
	var res PersonalityUploadSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
