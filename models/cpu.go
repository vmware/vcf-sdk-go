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

// CPU Represents information about CPUs on a host
//
// swagger:model Cpu
type CPU struct {

	// Number of CPU cores
	Cores int32 `json:"cores,omitempty"`

	// Information about each of the CPU cores
	CPUCores []*CPUCore `json:"cpuCores"`

	// Total CPU frequency in MHz
	FrequencyMHz float64 `json:"frequencyMHz,omitempty"`

	// Used CPU frequency in MHz
	UsedFrequencyMHz float64 `json:"usedFrequencyMHz,omitempty"`
}

// Validate validates this Cpu
func (m *CPU) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCPUCores(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CPU) validateCPUCores(formats strfmt.Registry) error {
	if swag.IsZero(m.CPUCores) { // not required
		return nil
	}

	for i := 0; i < len(m.CPUCores); i++ {
		if swag.IsZero(m.CPUCores[i]) { // not required
			continue
		}

		if m.CPUCores[i] != nil {
			if err := m.CPUCores[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("cpuCores" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("cpuCores" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this Cpu based on the context it is used
func (m *CPU) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCPUCores(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CPU) contextValidateCPUCores(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.CPUCores); i++ {

		if m.CPUCores[i] != nil {
			if err := m.CPUCores[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("cpuCores" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("cpuCores" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *CPU) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CPU) UnmarshalBinary(b []byte) error {
	var res CPU
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
