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

// NsxTQueryResponse Represents a NSX cluster query response.
//
// swagger:model NsxTQueryResponse
type NsxTQueryResponse struct {

	// Query info
	QueryInfo *QueryInfo `json:"queryInfo,omitempty"`

	// Query result
	Result *PageOfNsxTCluster `json:"result,omitempty"`
}

// Validate validates this nsx t query response
func (m *NsxTQueryResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateQueryInfo(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResult(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NsxTQueryResponse) validateQueryInfo(formats strfmt.Registry) error {
	if swag.IsZero(m.QueryInfo) { // not required
		return nil
	}

	if m.QueryInfo != nil {
		if err := m.QueryInfo.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("queryInfo")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("queryInfo")
			}
			return err
		}
	}

	return nil
}

func (m *NsxTQueryResponse) validateResult(formats strfmt.Registry) error {
	if swag.IsZero(m.Result) { // not required
		return nil
	}

	if m.Result != nil {
		if err := m.Result.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("result")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("result")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this nsx t query response based on the context it is used
func (m *NsxTQueryResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateQueryInfo(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateResult(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NsxTQueryResponse) contextValidateQueryInfo(ctx context.Context, formats strfmt.Registry) error {

	if m.QueryInfo != nil {

		if swag.IsZero(m.QueryInfo) { // not required
			return nil
		}

		if err := m.QueryInfo.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("queryInfo")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("queryInfo")
			}
			return err
		}
	}

	return nil
}

func (m *NsxTQueryResponse) contextValidateResult(ctx context.Context, formats strfmt.Registry) error {

	if m.Result != nil {

		if swag.IsZero(m.Result) { // not required
			return nil
		}

		if err := m.Result.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("result")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("result")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *NsxTQueryResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NsxTQueryResponse) UnmarshalBinary(b []byte) error {
	var res NsxTQueryResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
