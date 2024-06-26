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

// AssessmentTaskInfo A summary of an assessment execution
//
// swagger:model AssessmentTaskInfo
type AssessmentTaskInfo struct {

	// Timestamp corresponding to the time when the assessment was finished
	CompletionTimestamp int64 `json:"completionTimestamp,omitempty"`

	// The domains that have been assessed in the run associated to the given result
	DomainInputs []*DomainInfo `json:"domainInputs"`

	// The id of the assessment run
	ID string `json:"id,omitempty"`

	// Metadata about the assessment run that is triggered and passed to the result
	Metadata *AssessmentMetadata `json:"metadata,omitempty"`

	// The related task state of the run associated to the given result
	State string `json:"state,omitempty"`

	// Timestamp corresponding to the time when the assessment was initiated
	Timestamp int64 `json:"timestamp,omitempty"`
}

// Validate validates this assessment task info
func (m *AssessmentTaskInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDomainInputs(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMetadata(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AssessmentTaskInfo) validateDomainInputs(formats strfmt.Registry) error {
	if swag.IsZero(m.DomainInputs) { // not required
		return nil
	}

	for i := 0; i < len(m.DomainInputs); i++ {
		if swag.IsZero(m.DomainInputs[i]) { // not required
			continue
		}

		if m.DomainInputs[i] != nil {
			if err := m.DomainInputs[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("domainInputs" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("domainInputs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *AssessmentTaskInfo) validateMetadata(formats strfmt.Registry) error {
	if swag.IsZero(m.Metadata) { // not required
		return nil
	}

	if m.Metadata != nil {
		if err := m.Metadata.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("metadata")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("metadata")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this assessment task info based on the context it is used
func (m *AssessmentTaskInfo) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDomainInputs(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMetadata(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AssessmentTaskInfo) contextValidateDomainInputs(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.DomainInputs); i++ {

		if m.DomainInputs[i] != nil {

			if swag.IsZero(m.DomainInputs[i]) { // not required
				return nil
			}

			if err := m.DomainInputs[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("domainInputs" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("domainInputs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *AssessmentTaskInfo) contextValidateMetadata(ctx context.Context, formats strfmt.Registry) error {

	if m.Metadata != nil {

		if swag.IsZero(m.Metadata) { // not required
			return nil
		}

		if err := m.Metadata.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("metadata")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("metadata")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AssessmentTaskInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AssessmentTaskInfo) UnmarshalBinary(b []byte) error {
	var res AssessmentTaskInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
