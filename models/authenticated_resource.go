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

// AuthenticatedResource Represents the credential resource in the system
//
// swagger:model AuthenticatedResource
type AuthenticatedResource struct {

	// Domain name
	// Required: true
	DomainName *string `json:"domainName"`

	// Resource ID
	// Required: true
	ResourceID *string `json:"resourceId"`

	// Resource IP address
	// Required: true
	ResourceIP *string `json:"resourceIp"`

	// Resource name
	// Required: true
	ResourceName *string `json:"resourceName"`

	// Resource type
	// Example: One among: ESXI, VCENTER, PSC, NSX_MANAGER, NSX_CONTROLLER, NSX_EDGE, NSXT_MANAGER, NSXT_EDGE, VRLI, VROPS, VRA, WSA, VRSLCM, VXRAIL_MANAGER, BACKUP
	// Required: true
	ResourceType *string `json:"resourceType"`
}

// Validate validates this authenticated resource
func (m *AuthenticatedResource) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDomainName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResourceID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResourceIP(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResourceName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResourceType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AuthenticatedResource) validateDomainName(formats strfmt.Registry) error {

	if err := validate.Required("domainName", "body", m.DomainName); err != nil {
		return err
	}

	return nil
}

func (m *AuthenticatedResource) validateResourceID(formats strfmt.Registry) error {

	if err := validate.Required("resourceId", "body", m.ResourceID); err != nil {
		return err
	}

	return nil
}

func (m *AuthenticatedResource) validateResourceIP(formats strfmt.Registry) error {

	if err := validate.Required("resourceIp", "body", m.ResourceIP); err != nil {
		return err
	}

	return nil
}

func (m *AuthenticatedResource) validateResourceName(formats strfmt.Registry) error {

	if err := validate.Required("resourceName", "body", m.ResourceName); err != nil {
		return err
	}

	return nil
}

func (m *AuthenticatedResource) validateResourceType(formats strfmt.Registry) error {

	if err := validate.Required("resourceType", "body", m.ResourceType); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this authenticated resource based on context it is used
func (m *AuthenticatedResource) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AuthenticatedResource) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AuthenticatedResource) UnmarshalBinary(b []byte) error {
	var res AuthenticatedResource
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
