// Code generated by go-swagger; DO NOT EDIT.

// Copyright © 2023 VMware, Inc. All Rights Reserved.
// Code licenced under: BSD-2

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

// DatastoreMountSpec Specification for datastore mount configuration.
//
// swagger:model DatastoreMountSpec
type DatastoreMountSpec struct {

	// Cluster storage configuration; e.g. HCI Mesh remote vSAN
	// Required: true
	DatastoreSpec *DatastoreSpec `json:"datastoreSpec"`
}

// Validate validates this datastore mount spec
func (m *DatastoreMountSpec) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDatastoreSpec(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DatastoreMountSpec) validateDatastoreSpec(formats strfmt.Registry) error {

	if err := validate.Required("datastoreSpec", "body", m.DatastoreSpec); err != nil {
		return err
	}

	if m.DatastoreSpec != nil {
		if err := m.DatastoreSpec.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("datastoreSpec")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("datastoreSpec")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this datastore mount spec based on the context it is used
func (m *DatastoreMountSpec) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDatastoreSpec(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DatastoreMountSpec) contextValidateDatastoreSpec(ctx context.Context, formats strfmt.Registry) error {

	if m.DatastoreSpec != nil {
		if err := m.DatastoreSpec.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("datastoreSpec")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("datastoreSpec")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DatastoreMountSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DatastoreMountSpec) UnmarshalBinary(b []byte) error {
	var res DatastoreMountSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
