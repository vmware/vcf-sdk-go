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
)

// Task Represents a task
//
// swagger:model Task
type Task struct {

	// Task completion timestamp
	CompletionTimestamp string `json:"completionTimestamp,omitempty"`

	// Task creation timestamp
	CreationTimestamp string `json:"creationTimestamp,omitempty"`

	// List of errors in case of a failure
	Errors []*Error `json:"errors"`

	// Task ID
	ID string `json:"id,omitempty"`

	// Represents task can be cancellable or not.
	IsCancellable bool `json:"isCancellable,omitempty"`

	// Indicates whether a task is eligible for retry or not.
	IsRetryable bool `json:"isRetryable,omitempty"`

	// Localizable Task description
	LocalizableDescriptionPack *MessagePack `json:"localizableDescriptionPack,omitempty"`

	// Task name
	Name string `json:"name,omitempty"`

	// Resolution state
	// Example: Can have only one of the two values: RESOLVED or UNRESOLVED
	ResolutionStatus string `json:"resolutionStatus,omitempty"`

	// List of resources associated with task
	Resources []*Resource `json:"resources"`

	// Task status
	// Example: One among: PENDING, IN_PROGRESS, In Progress, SUCCESSFUL, Successful, FAILED, Failed, CANCELLED, Cancelled, COMPLETED_WITH_WARNING, SKIPPED
	Status string `json:"status,omitempty"`

	// List of sub-tasks of the task
	SubTasks []*SubTask `json:"subTasks"`

	// Operation that is represented by the Task in machine readable format.  The value is controlled by the owners/producers of the Task. The convention is <resource>_<operation>
	// Example: Sample values: HOST_COMMISSION, HOST_DECOMMISSION
	Type string `json:"type,omitempty"`
}

// Validate validates this task
func (m *Task) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateErrors(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLocalizableDescriptionPack(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResources(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSubTasks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Task) validateErrors(formats strfmt.Registry) error {
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

func (m *Task) validateLocalizableDescriptionPack(formats strfmt.Registry) error {
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

func (m *Task) validateResources(formats strfmt.Registry) error {
	if swag.IsZero(m.Resources) { // not required
		return nil
	}

	for i := 0; i < len(m.Resources); i++ {
		if swag.IsZero(m.Resources[i]) { // not required
			continue
		}

		if m.Resources[i] != nil {
			if err := m.Resources[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("resources" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("resources" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Task) validateSubTasks(formats strfmt.Registry) error {
	if swag.IsZero(m.SubTasks) { // not required
		return nil
	}

	for i := 0; i < len(m.SubTasks); i++ {
		if swag.IsZero(m.SubTasks[i]) { // not required
			continue
		}

		if m.SubTasks[i] != nil {
			if err := m.SubTasks[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("subTasks" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("subTasks" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this task based on the context it is used
func (m *Task) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateErrors(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLocalizableDescriptionPack(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateResources(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSubTasks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Task) contextValidateErrors(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Errors); i++ {

		if m.Errors[i] != nil {

			if swag.IsZero(m.Errors[i]) { // not required
				return nil
			}

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

func (m *Task) contextValidateLocalizableDescriptionPack(ctx context.Context, formats strfmt.Registry) error {

	if m.LocalizableDescriptionPack != nil {

		if swag.IsZero(m.LocalizableDescriptionPack) { // not required
			return nil
		}

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

func (m *Task) contextValidateResources(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Resources); i++ {

		if m.Resources[i] != nil {

			if swag.IsZero(m.Resources[i]) { // not required
				return nil
			}

			if err := m.Resources[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("resources" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("resources" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Task) contextValidateSubTasks(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.SubTasks); i++ {

		if m.SubTasks[i] != nil {

			if swag.IsZero(m.SubTasks[i]) { // not required
				return nil
			}

			if err := m.SubTasks[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("subTasks" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("subTasks" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *Task) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Task) UnmarshalBinary(b []byte) error {
	var res Task
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
