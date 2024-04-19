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
	"github.com/go-openapi/validate"
)

// ImportVdsSpec This specification contains the parameters required to import vSphere Distributed Switches to the inventory for a cluster.
//
// swagger:model ImportVdsSpec
type ImportVdsSpec struct {

	// List of details of vSphere Distributed Switches to be imported to the inventory
	// Required: true
	ListOfVdsDetails []*VdsDetail `json:"listOfVdsDetails"`
}

// Validate validates this import vds spec
func (m *ImportVdsSpec) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateListOfVdsDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ImportVdsSpec) validateListOfVdsDetails(formats strfmt.Registry) error {

	if err := validate.Required("listOfVdsDetails", "body", m.ListOfVdsDetails); err != nil {
		return err
	}

	for i := 0; i < len(m.ListOfVdsDetails); i++ {
		if swag.IsZero(m.ListOfVdsDetails[i]) { // not required
			continue
		}

		if m.ListOfVdsDetails[i] != nil {
			if err := m.ListOfVdsDetails[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("listOfVdsDetails" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("listOfVdsDetails" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this import vds spec based on the context it is used
func (m *ImportVdsSpec) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateListOfVdsDetails(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ImportVdsSpec) contextValidateListOfVdsDetails(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ListOfVdsDetails); i++ {

		if m.ListOfVdsDetails[i] != nil {

			if swag.IsZero(m.ListOfVdsDetails[i]) { // not required
				return nil
			}

			if err := m.ListOfVdsDetails[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("listOfVdsDetails" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("listOfVdsDetails" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ImportVdsSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ImportVdsSpec) UnmarshalBinary(b []byte) error {
	var res ImportVdsSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
