// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2023 VMware, Inc.
// SPDX-License-Identifier: BSD-2-Clause

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ResourceCertificateSpec This specification contains resource certificate details. Either resourceId or resourceFqdn should be provided. Either certificateChain or both resourceCertificate and caCertificate should be provided.
//
// swagger:model ResourceCertificateSpec
type ResourceCertificateSpec struct {

	// CA Certificate
	// Example: -----BEGIN CERTIFICATE-----\nMIIFq...\n-----END CERTIFICATE-----
	CaCertificate string `json:"caCertificate,omitempty"`

	// Certificate Chain
	// Example: -----BEGIN CERTIFICATE-----\nMIIFq...\n-----END CERTIFICATE-----\n-----BEGIN CERTIFICATE-----\nMIIFq...\n-----END CERTIFICATE-----
	CertificateChain string `json:"certificateChain,omitempty"`

	// Resource Certificate
	// Example: -----BEGIN CERTIFICATE-----\nMIIFq...\n-----END CERTIFICATE-----
	ResourceCertificate string `json:"resourceCertificate,omitempty"`

	// Resource FQDN
	// Example: sfo-vc01.rainpole.io
	ResourceFqdn string `json:"resourceFqdn,omitempty"`

	// Resource ID
	// Example: BE8A5E04-92A0-43F6-A166-AA041F4327CC
	ResourceID string `json:"resourceId,omitempty"`
}

// Validate validates this resource certificate spec
func (m *ResourceCertificateSpec) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this resource certificate spec based on context it is used
func (m *ResourceCertificateSpec) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ResourceCertificateSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ResourceCertificateSpec) UnmarshalBinary(b []byte) error {
	var res ResourceCertificateSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}