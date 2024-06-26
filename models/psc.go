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
)

// Psc Psc representation
//
// swagger:model Psc
type Psc struct {

	// Information about the domain this PSC is part of
	Domain *DomainReference `json:"domain,omitempty"`

	// FQDN of the Psc
	Fqdn string `json:"fqdn,omitempty"`

	// ID of the Psc
	ID string `json:"id,omitempty"`

	// IP Address of the Psc
	IPAddress string `json:"ipAddress,omitempty"`

	// Indicates if the Psc is a replica
	IsReplica bool `json:"isReplica,omitempty"`

	// SSO Domain name of the Psc
	SSODomainName string `json:"ssoDomainName,omitempty"`

	// SSO sub domain name of the Psc
	SSOSubDomainName string `json:"ssoSubDomainName,omitempty"`
}

// Validate validates this psc
func (m *Psc) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDomain(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Psc) validateDomain(formats strfmt.Registry) error {
	if swag.IsZero(m.Domain) { // not required
		return nil
	}

	if m.Domain != nil {
		if err := m.Domain.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("domain")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("domain")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this psc based on the context it is used
func (m *Psc) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDomain(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Psc) contextValidateDomain(ctx context.Context, formats strfmt.Registry) error {

	if m.Domain != nil {

		if swag.IsZero(m.Domain) { // not required
			return nil
		}

		if err := m.Domain.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("domain")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("domain")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Psc) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Psc) UnmarshalBinary(b []byte) error {
	var res Psc
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
