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

// CredentialsValidation Represents a validation with a list of one or more validation checks that are performed as part of the validation
//
// swagger:model CredentialsValidation
type CredentialsValidation struct {

	// Description of the validation
	Description string `json:"description,omitempty"`

	// Execution status of the validation
	// Example: One among: COMPLETED, FAILED, IN_PROGRESS
	ExecutionStatus string `json:"executionStatus,omitempty"`

	// ID of the validation
	ID string `json:"id,omitempty"`

	// List of one or more validation checks that are performed as part of the validation
	ValidationChecks []*CredentialValidationCheck `json:"validationChecks"`
}

// Validate validates this credentials validation
func (m *CredentialsValidation) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateValidationChecks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CredentialsValidation) validateValidationChecks(formats strfmt.Registry) error {
	if swag.IsZero(m.ValidationChecks) { // not required
		return nil
	}

	for i := 0; i < len(m.ValidationChecks); i++ {
		if swag.IsZero(m.ValidationChecks[i]) { // not required
			continue
		}

		if m.ValidationChecks[i] != nil {
			if err := m.ValidationChecks[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("validationChecks" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("validationChecks" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this credentials validation based on the context it is used
func (m *CredentialsValidation) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateValidationChecks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CredentialsValidation) contextValidateValidationChecks(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ValidationChecks); i++ {

		if m.ValidationChecks[i] != nil {
			if err := m.ValidationChecks[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("validationChecks" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("validationChecks" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *CredentialsValidation) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CredentialsValidation) UnmarshalBinary(b []byte) error {
	var res CredentialsValidation
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
