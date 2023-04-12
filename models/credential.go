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

// Credential Represents a credential and the resource to which it is associated with
//
// swagger:model Credential
type Credential struct {

	// Account type
	// Example: One among: USER, SYSTEM, SERVICE
	// Required: true
	AccountType *string `json:"accountType"`

	// Configured auto-rotate policy of a credential. Empty if not configured
	AutoRotatePolicy *AutoRotateCredentialPolicy `json:"autoRotatePolicy,omitempty"`

	// The timestamp at which credential was created
	// Required: true
	CreationTimestamp *string `json:"creationTimestamp"`

	// Credential type
	// Example: One among: SSO, SSH, API, FTP, AUDIT
	// Required: true
	CredentialType *string `json:"credentialType"`

	// Password expiration details
	Expiry *ExpirationDetails `json:"expiry,omitempty"`

	// Credential ID
	// Required: true
	ID *string `json:"id"`

	// The timestamp at which credential was last modified
	// Required: true
	ModificationTimestamp *string `json:"modificationTimestamp"`

	// Password
	Password string `json:"password,omitempty"`

	// The resource which owns the credential
	// Required: true
	Resource *AuthenticatedResource `json:"resource"`

	// Username
	// Required: true
	Username *string `json:"username"`
}

// Validate validates this credential
func (m *Credential) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAccountType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAutoRotatePolicy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreationTimestamp(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCredentialType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExpiry(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateModificationTimestamp(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResource(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUsername(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Credential) validateAccountType(formats strfmt.Registry) error {

	if err := validate.Required("accountType", "body", m.AccountType); err != nil {
		return err
	}

	return nil
}

func (m *Credential) validateAutoRotatePolicy(formats strfmt.Registry) error {
	if swag.IsZero(m.AutoRotatePolicy) { // not required
		return nil
	}

	if m.AutoRotatePolicy != nil {
		if err := m.AutoRotatePolicy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("autoRotatePolicy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("autoRotatePolicy")
			}
			return err
		}
	}

	return nil
}

func (m *Credential) validateCreationTimestamp(formats strfmt.Registry) error {

	if err := validate.Required("creationTimestamp", "body", m.CreationTimestamp); err != nil {
		return err
	}

	return nil
}

func (m *Credential) validateCredentialType(formats strfmt.Registry) error {

	if err := validate.Required("credentialType", "body", m.CredentialType); err != nil {
		return err
	}

	return nil
}

func (m *Credential) validateExpiry(formats strfmt.Registry) error {
	if swag.IsZero(m.Expiry) { // not required
		return nil
	}

	if m.Expiry != nil {
		if err := m.Expiry.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("expiry")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("expiry")
			}
			return err
		}
	}

	return nil
}

func (m *Credential) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *Credential) validateModificationTimestamp(formats strfmt.Registry) error {

	if err := validate.Required("modificationTimestamp", "body", m.ModificationTimestamp); err != nil {
		return err
	}

	return nil
}

func (m *Credential) validateResource(formats strfmt.Registry) error {

	if err := validate.Required("resource", "body", m.Resource); err != nil {
		return err
	}

	if m.Resource != nil {
		if err := m.Resource.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("resource")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("resource")
			}
			return err
		}
	}

	return nil
}

func (m *Credential) validateUsername(formats strfmt.Registry) error {

	if err := validate.Required("username", "body", m.Username); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this credential based on the context it is used
func (m *Credential) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAutoRotatePolicy(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateExpiry(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateResource(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Credential) contextValidateAutoRotatePolicy(ctx context.Context, formats strfmt.Registry) error {

	if m.AutoRotatePolicy != nil {
		if err := m.AutoRotatePolicy.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("autoRotatePolicy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("autoRotatePolicy")
			}
			return err
		}
	}

	return nil
}

func (m *Credential) contextValidateExpiry(ctx context.Context, formats strfmt.Registry) error {

	if m.Expiry != nil {
		if err := m.Expiry.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("expiry")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("expiry")
			}
			return err
		}
	}

	return nil
}

func (m *Credential) contextValidateResource(ctx context.Context, formats strfmt.Registry) error {

	if m.Resource != nil {
		if err := m.Resource.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("resource")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("resource")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Credential) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Credential) UnmarshalBinary(b []byte) error {
	var res Credential
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
