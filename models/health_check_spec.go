// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2023 VMware, Inc.
// SPDX-License-Identifier: BSD-2-Clause

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

// HealthCheckSpec HealthCheck
//
// swagger:model HealthCheckSpec
type HealthCheckSpec struct {

	// Health Check id
	// Required: true
	ID *string `json:"id"`

	// Health Check name
	// Required: true
	Name *string `json:"name"`

	// Name of the resource, if the ResourceType is CLUSTER, then the resourceName == ClusterName
	ResourceName string `json:"resourceName,omitempty"`

	// Resource Type
	// Required: true
	ResourceType *string `json:"resourceType"`

	// Health check status
	// Required: true
	Status *string `json:"status"`
}

// Validate validates this health check spec
func (m *HealthCheckSpec) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResourceType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HealthCheckSpec) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *HealthCheckSpec) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *HealthCheckSpec) validateResourceType(formats strfmt.Registry) error {

	if err := validate.Required("resourceType", "body", m.ResourceType); err != nil {
		return err
	}

	return nil
}

func (m *HealthCheckSpec) validateStatus(formats strfmt.Registry) error {

	if err := validate.Required("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this health check spec based on context it is used
func (m *HealthCheckSpec) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *HealthCheckSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HealthCheckSpec) UnmarshalBinary(b []byte) error {
	var res HealthCheckSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
