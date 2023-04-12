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

// PersonalityUploadSpecRaw Personality upload specification for uploading the personality from raw files exported from avCenter cluster. This mode of uplaoding personality is useful when the source vCenter cluster isoutside the target VCF deployment.
//
// swagger:model PersonalityUploadSpecRaw
type PersonalityUploadSpecRaw struct {

	// Personality ISO File Path
	PersonalityISOFilePath string `json:"personalityISOFilePath,omitempty"`

	// Personality Info JSON File Path
	// Required: true
	PersonalityInfoJSONFilePath *string `json:"personalityInfoJSONFilePath"`

	// Personality JSON File Path
	// Required: true
	PersonalityJSONFilePath *string `json:"personalityJSONFilePath"`

	// Personality Zip File Path
	// Required: true
	PersonalityZIPFilePath *string `json:"personalityZIPFilePath"`
}

// Validate validates this personality upload spec raw
func (m *PersonalityUploadSpecRaw) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePersonalityInfoJSONFilePath(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePersonalityJSONFilePath(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePersonalityZIPFilePath(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PersonalityUploadSpecRaw) validatePersonalityInfoJSONFilePath(formats strfmt.Registry) error {

	if err := validate.Required("personalityInfoJSONFilePath", "body", m.PersonalityInfoJSONFilePath); err != nil {
		return err
	}

	return nil
}

func (m *PersonalityUploadSpecRaw) validatePersonalityJSONFilePath(formats strfmt.Registry) error {

	if err := validate.Required("personalityJSONFilePath", "body", m.PersonalityJSONFilePath); err != nil {
		return err
	}

	return nil
}

func (m *PersonalityUploadSpecRaw) validatePersonalityZIPFilePath(formats strfmt.Registry) error {

	if err := validate.Required("personalityZIPFilePath", "body", m.PersonalityZIPFilePath); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this personality upload spec raw based on context it is used
func (m *PersonalityUploadSpecRaw) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PersonalityUploadSpecRaw) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PersonalityUploadSpecRaw) UnmarshalBinary(b []byte) error {
	var res PersonalityUploadSpecRaw
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
