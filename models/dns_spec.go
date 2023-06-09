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

// DNSSpec Spec contains parameters of Domain Name System
//
// swagger:model DnsSpec
type DNSSpec struct {

	// Tenant domain
	// Example: vmware.com
	// Required: true
	Domain *string `json:"domain"`

	// Primary nameserver to be configured for vCenter/PSC/ESXi's/NSX
	// Example: 172.0.0.4
	Nameserver string `json:"nameserver,omitempty"`

	// Secondary nameserver to be configured for vCenter/PSC/ESXi's/NSX
	// Example: 172.0.0.5
	SecondaryNameserver string `json:"secondaryNameserver,omitempty"`

	// Tenant Sub domain
	// Example: vcf.vmware.com
	// Required: true
	Subdomain *string `json:"subdomain"`
}

// Validate validates this Dns spec
func (m *DNSSpec) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDomain(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSubdomain(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DNSSpec) validateDomain(formats strfmt.Registry) error {

	if err := validate.Required("domain", "body", m.Domain); err != nil {
		return err
	}

	return nil
}

func (m *DNSSpec) validateSubdomain(formats strfmt.Registry) error {

	if err := validate.Required("subdomain", "body", m.Subdomain); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this Dns spec based on context it is used
func (m *DNSSpec) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DNSSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DNSSpec) UnmarshalBinary(b []byte) error {
	var res DNSSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
