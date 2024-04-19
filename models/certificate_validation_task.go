// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2023 VMware, Inc.
// SPDX-License-Identifier: BSD-2-Clause

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// CertificateValidationTask The Certificate Validation Task.
//
// swagger:model CertificateValidationTask
type CertificateValidationTask struct {

	// Validation Task Completed
	// Example: One among: true, false
	// Required: true
	Completed *bool `json:"completed"`

	// Validation Task End Time
	// Example: 2021-02-02T00:00:00.000Z
	EndTimestamp string `json:"endTimestamp,omitempty"`

	// Validation Task Start Time
	// Example: 2021-02-02T00:00:00.000Z
	StartTimestamp string `json:"startTimestamp,omitempty"`

	// Validation ID
	// Example: B1930850-7D1A-4BAA-89D7-52FD5DAD196A
	// Required: true
	ValidationID *string `json:"validationId"`

	// Resource Certificate Validations
	// Required: true
	Validations []*CertificateValidation `json:"validations"`
}

// Validate validates this certificate validation task
func (m *CertificateValidationTask) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCompleted(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValidationID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValidations(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CertificateValidationTask) validateCompleted(formats strfmt.Registry) error {

	if err := validate.Required("completed", "body", m.Completed); err != nil {
		return err
	}

	return nil
}

func (m *CertificateValidationTask) validateValidationID(formats strfmt.Registry) error {

	if err := validate.Required("validationId", "body", m.ValidationID); err != nil {
		return err
	}

	return nil
}

func (m *CertificateValidationTask) validateValidations(formats strfmt.Registry) error {

	if err := validate.Required("validations", "body", m.Validations); err != nil {
		return err
	}

	for i := 0; i < len(m.Validations); i++ {
		if swag.IsZero(m.Validations[i]) { // not required
			continue
		}

		if m.Validations[i] != nil {
			if err := m.Validations[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("validations" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("validations" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this certificate validation task based on the context it is used
func (m *CertificateValidationTask) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateValidations(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CertificateValidationTask) contextValidateValidations(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Validations); i++ {

		if m.Validations[i] != nil {

			if swag.IsZero(m.Validations[i]) { // not required
				return nil
			}

			if err := m.Validations[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("validations" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("validations" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *CertificateValidationTask) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CertificateValidationTask) UnmarshalBinary(b []byte) error {
	var res CertificateValidationTask
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
