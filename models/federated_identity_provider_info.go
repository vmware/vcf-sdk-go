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

// FederatedIdentityProviderInfo The identity management info when the provider is via broker federation
//
// swagger:model FederatedIdentityProviderInfo
type FederatedIdentityProviderInfo struct {

	// The directory info of the Identity Provider
	DirectoryList *IdentityProviderDirectory `json:"directoryList,omitempty"`

	// The user-friendly name for the Identity Provider
	Name string `json:"name,omitempty"`

	// The OIDC profile of the Identity Provider
	OidcInfo *OidcInfo `json:"oidcInfo,omitempty"`

	// The source of the Identity Provider
	// Example: One among: OKTA, AZURE
	Source string `json:"source,omitempty"`

	// The lifetime in seconds of the sync client bear token
	// Example: One among: ACTIVE, INACTIVE
	SyncClientTokenTTL int64 `json:"syncClientTokenTTL,omitempty"`
}

// Validate validates this federated identity provider info
func (m *FederatedIdentityProviderInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDirectoryList(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOidcInfo(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FederatedIdentityProviderInfo) validateDirectoryList(formats strfmt.Registry) error {
	if swag.IsZero(m.DirectoryList) { // not required
		return nil
	}

	if m.DirectoryList != nil {
		if err := m.DirectoryList.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("directoryList")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("directoryList")
			}
			return err
		}
	}

	return nil
}

func (m *FederatedIdentityProviderInfo) validateOidcInfo(formats strfmt.Registry) error {
	if swag.IsZero(m.OidcInfo) { // not required
		return nil
	}

	if m.OidcInfo != nil {
		if err := m.OidcInfo.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("oidcInfo")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("oidcInfo")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this federated identity provider info based on the context it is used
func (m *FederatedIdentityProviderInfo) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDirectoryList(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateOidcInfo(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FederatedIdentityProviderInfo) contextValidateDirectoryList(ctx context.Context, formats strfmt.Registry) error {

	if m.DirectoryList != nil {

		if swag.IsZero(m.DirectoryList) { // not required
			return nil
		}

		if err := m.DirectoryList.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("directoryList")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("directoryList")
			}
			return err
		}
	}

	return nil
}

func (m *FederatedIdentityProviderInfo) contextValidateOidcInfo(ctx context.Context, formats strfmt.Registry) error {

	if m.OidcInfo != nil {

		if swag.IsZero(m.OidcInfo) { // not required
			return nil
		}

		if err := m.OidcInfo.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("oidcInfo")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("oidcInfo")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *FederatedIdentityProviderInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FederatedIdentityProviderInfo) UnmarshalBinary(b []byte) error {
	var res FederatedIdentityProviderInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
