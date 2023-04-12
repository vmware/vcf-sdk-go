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

// Personality Personality contains bits to install/update the appropriate Cloud Foundation software components in your management domain or workload domain.
//
// swagger:model Personality
type Personality struct {

	// Personality created by
	CreatedBy string `json:"createdBy,omitempty"`

	// Personality description
	// Example: ESXi 7.0 Dell Personality
	// Required: true
	Description *string `json:"description"`

	// Personality displayName
	// Example: ESXi 7.0 Dell Personality
	// Required: true
	DisplayName *string `json:"displayName"`

	// Personality image checksum
	// Required: true
	ImageChecksum *string `json:"imageChecksum"`

	// Personality image size
	// Required: true
	ImageSize *string `json:"imageSize"`

	// Personality kb articles
	KbArticles *URL `json:"kbArticles,omitempty"`

	// Personality id
	// Required: true
	PersonalityID *string `json:"personalityId"`

	// Personality name
	// Required: true
	PersonalityName *string `json:"personalityName"`

	// Personality Release date
	ReleaseDate int64 `json:"releaseDate,omitempty"`

	// Personality software spec
	// Required: true
	SoftwareInfo *SoftwareInfo `json:"softwareInfo"`

	// Personality tags
	Tags []string `json:"tags"`

	// Personality Version
	// Required: true
	Version *string `json:"version"`

	// Personality depot path
	// Required: true
	VsphereExportedIsoPath *string `json:"vsphereExportedIsoPath"`

	// Personality depot path
	// Required: true
	VsphereExportedJSONPath *string `json:"vsphereExportedJsonPath"`

	// Personality depot path
	// Required: true
	VsphereExportedZipPath *string `json:"vsphereExportedZipPath"`
}

// Validate validates this personality
func (m *Personality) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDisplayName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateImageChecksum(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateImageSize(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateKbArticles(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePersonalityID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePersonalityName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSoftwareInfo(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVersion(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVsphereExportedIsoPath(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVsphereExportedJSONPath(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVsphereExportedZipPath(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Personality) validateDescription(formats strfmt.Registry) error {

	if err := validate.Required("description", "body", m.Description); err != nil {
		return err
	}

	return nil
}

func (m *Personality) validateDisplayName(formats strfmt.Registry) error {

	if err := validate.Required("displayName", "body", m.DisplayName); err != nil {
		return err
	}

	return nil
}

func (m *Personality) validateImageChecksum(formats strfmt.Registry) error {

	if err := validate.Required("imageChecksum", "body", m.ImageChecksum); err != nil {
		return err
	}

	return nil
}

func (m *Personality) validateImageSize(formats strfmt.Registry) error {

	if err := validate.Required("imageSize", "body", m.ImageSize); err != nil {
		return err
	}

	return nil
}

func (m *Personality) validateKbArticles(formats strfmt.Registry) error {
	if swag.IsZero(m.KbArticles) { // not required
		return nil
	}

	if m.KbArticles != nil {
		if err := m.KbArticles.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("kbArticles")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("kbArticles")
			}
			return err
		}
	}

	return nil
}

func (m *Personality) validatePersonalityID(formats strfmt.Registry) error {

	if err := validate.Required("personalityId", "body", m.PersonalityID); err != nil {
		return err
	}

	return nil
}

func (m *Personality) validatePersonalityName(formats strfmt.Registry) error {

	if err := validate.Required("personalityName", "body", m.PersonalityName); err != nil {
		return err
	}

	return nil
}

func (m *Personality) validateSoftwareInfo(formats strfmt.Registry) error {

	if err := validate.Required("softwareInfo", "body", m.SoftwareInfo); err != nil {
		return err
	}

	if m.SoftwareInfo != nil {
		if err := m.SoftwareInfo.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("softwareInfo")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("softwareInfo")
			}
			return err
		}
	}

	return nil
}

func (m *Personality) validateVersion(formats strfmt.Registry) error {

	if err := validate.Required("version", "body", m.Version); err != nil {
		return err
	}

	return nil
}

func (m *Personality) validateVsphereExportedIsoPath(formats strfmt.Registry) error {

	if err := validate.Required("vsphereExportedIsoPath", "body", m.VsphereExportedIsoPath); err != nil {
		return err
	}

	return nil
}

func (m *Personality) validateVsphereExportedJSONPath(formats strfmt.Registry) error {

	if err := validate.Required("vsphereExportedJsonPath", "body", m.VsphereExportedJSONPath); err != nil {
		return err
	}

	return nil
}

func (m *Personality) validateVsphereExportedZipPath(formats strfmt.Registry) error {

	if err := validate.Required("vsphereExportedZipPath", "body", m.VsphereExportedZipPath); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this personality based on the context it is used
func (m *Personality) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateKbArticles(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSoftwareInfo(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Personality) contextValidateKbArticles(ctx context.Context, formats strfmt.Registry) error {

	if m.KbArticles != nil {
		if err := m.KbArticles.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("kbArticles")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("kbArticles")
			}
			return err
		}
	}

	return nil
}

func (m *Personality) contextValidateSoftwareInfo(ctx context.Context, formats strfmt.Registry) error {

	if m.SoftwareInfo != nil {
		if err := m.SoftwareInfo.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("softwareInfo")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("softwareInfo")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Personality) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Personality) UnmarshalBinary(b []byte) error {
	var res Personality
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}