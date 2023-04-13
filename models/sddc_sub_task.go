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

// SDDCSubTask Represents a SDDC sub-task
//
// swagger:model SddcSubTask
type SDDCSubTask struct {

	// Sub-Task Creation Time
	CreationTimestamp string `json:"creationTimestamp,omitempty"`

	// Sub-Task Description
	Description string `json:"description,omitempty"`

	// List of errors in case of a failure
	Errors []*Error `json:"errors"`

	// Localizable SDDC Task description
	LocalizableDescriptionPack *MessagePack `json:"localizableDescriptionPack,omitempty"`

	// Localizable SDDC Task name
	LocalizableNamePack *MessagePack `json:"localizableNamePack,omitempty"`

	// Sub-Task Name
	Name string `json:"name,omitempty"`

	// Processing category description, e.g., VC Deployment, VSAN configuration etc
	ProcessingStateDescription string `json:"processingStateDescription,omitempty"`

	// Processing category name, e.g., VC Deployment, VSAN configuration etc
	ProcessingStateName string `json:"processingStateName,omitempty"`

	// SDDC ID
	// Example: 123e4567-e89b-42d3-a456-556642440000
	SDDCID string `json:"sddcId,omitempty"`

	// Task Status
	// Example: INITIALIZED, COMPLETED_WITH_SUCCESS, COMPLETED_WITH_FAILURE, PREVALIDATION_COMPLETED_WITH_SUCCESS, PREVALIDATION_COMPLETED_WITH_FAILURE, POSTVALIDATION_COMPLETED_WITH_SUCCESS, POSTVALIDATION_COMPLETED_WITH_FAILURE,IN_PROGRESS, PREVALIDATION_IN_PROGRESS, POSTVALIDATION_IN_PROGRESS, INTERNAL_ERROR
	Status string `json:"status,omitempty"`

	// Last Update Time of Sub-Task
	UpdateTimestamp string `json:"updateTimestamp,omitempty"`
}

// Validate validates this Sddc sub task
func (m *SDDCSubTask) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateErrors(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLocalizableDescriptionPack(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLocalizableNamePack(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SDDCSubTask) validateErrors(formats strfmt.Registry) error {
	if swag.IsZero(m.Errors) { // not required
		return nil
	}

	for i := 0; i < len(m.Errors); i++ {
		if swag.IsZero(m.Errors[i]) { // not required
			continue
		}

		if m.Errors[i] != nil {
			if err := m.Errors[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("errors" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("errors" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *SDDCSubTask) validateLocalizableDescriptionPack(formats strfmt.Registry) error {
	if swag.IsZero(m.LocalizableDescriptionPack) { // not required
		return nil
	}

	if m.LocalizableDescriptionPack != nil {
		if err := m.LocalizableDescriptionPack.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("localizableDescriptionPack")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("localizableDescriptionPack")
			}
			return err
		}
	}

	return nil
}

func (m *SDDCSubTask) validateLocalizableNamePack(formats strfmt.Registry) error {
	if swag.IsZero(m.LocalizableNamePack) { // not required
		return nil
	}

	if m.LocalizableNamePack != nil {
		if err := m.LocalizableNamePack.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("localizableNamePack")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("localizableNamePack")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this Sddc sub task based on the context it is used
func (m *SDDCSubTask) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateErrors(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLocalizableDescriptionPack(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLocalizableNamePack(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SDDCSubTask) contextValidateErrors(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Errors); i++ {

		if m.Errors[i] != nil {
			if err := m.Errors[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("errors" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("errors" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *SDDCSubTask) contextValidateLocalizableDescriptionPack(ctx context.Context, formats strfmt.Registry) error {

	if m.LocalizableDescriptionPack != nil {
		if err := m.LocalizableDescriptionPack.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("localizableDescriptionPack")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("localizableDescriptionPack")
			}
			return err
		}
	}

	return nil
}

func (m *SDDCSubTask) contextValidateLocalizableNamePack(ctx context.Context, formats strfmt.Registry) error {

	if m.LocalizableNamePack != nil {
		if err := m.LocalizableNamePack.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("localizableNamePack")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("localizableNamePack")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SDDCSubTask) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SDDCSubTask) UnmarshalBinary(b []byte) error {
	var res SDDCSubTask
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
