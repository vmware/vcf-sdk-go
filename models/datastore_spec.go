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

// DatastoreSpec This specification contains cluster storage configuration
//
// swagger:model DatastoreSpec
type DatastoreSpec struct {

	// Cluster storage configuration for NFS
	NfsDatastoreSpecs []*NfsDatastoreSpec `json:"nfsDatastoreSpecs"`

	// Cluster storage configuration for VMFS
	VmfsDatastoreSpec *VmfsDatastoreSpec `json:"vmfsDatastoreSpec,omitempty"`

	// Cluster storage configuration for vSAN
	VSANDatastoreSpec *VSANDatastoreSpec `json:"vsanDatastoreSpec,omitempty"`

	// Cluster storage configuration for vSAN Remote Datastore
	VSANRemoteDatastoreClusterSpec *VSANRemoteDatastoreClusterSpec `json:"vsanRemoteDatastoreClusterSpec,omitempty"`

	// Cluster storage configuration for VVOL
	VvolDatastoreSpecs []*VvolDatastoreSpec `json:"vvolDatastoreSpecs"`
}

// Validate validates this datastore spec
func (m *DatastoreSpec) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateNfsDatastoreSpecs(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVmfsDatastoreSpec(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVSANDatastoreSpec(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVSANRemoteDatastoreClusterSpec(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVvolDatastoreSpecs(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DatastoreSpec) validateNfsDatastoreSpecs(formats strfmt.Registry) error {
	if swag.IsZero(m.NfsDatastoreSpecs) { // not required
		return nil
	}

	for i := 0; i < len(m.NfsDatastoreSpecs); i++ {
		if swag.IsZero(m.NfsDatastoreSpecs[i]) { // not required
			continue
		}

		if m.NfsDatastoreSpecs[i] != nil {
			if err := m.NfsDatastoreSpecs[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("nfsDatastoreSpecs" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("nfsDatastoreSpecs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *DatastoreSpec) validateVmfsDatastoreSpec(formats strfmt.Registry) error {
	if swag.IsZero(m.VmfsDatastoreSpec) { // not required
		return nil
	}

	if m.VmfsDatastoreSpec != nil {
		if err := m.VmfsDatastoreSpec.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vmfsDatastoreSpec")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("vmfsDatastoreSpec")
			}
			return err
		}
	}

	return nil
}

func (m *DatastoreSpec) validateVSANDatastoreSpec(formats strfmt.Registry) error {
	if swag.IsZero(m.VSANDatastoreSpec) { // not required
		return nil
	}

	if m.VSANDatastoreSpec != nil {
		if err := m.VSANDatastoreSpec.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vsanDatastoreSpec")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("vsanDatastoreSpec")
			}
			return err
		}
	}

	return nil
}

func (m *DatastoreSpec) validateVSANRemoteDatastoreClusterSpec(formats strfmt.Registry) error {
	if swag.IsZero(m.VSANRemoteDatastoreClusterSpec) { // not required
		return nil
	}

	if m.VSANRemoteDatastoreClusterSpec != nil {
		if err := m.VSANRemoteDatastoreClusterSpec.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vsanRemoteDatastoreClusterSpec")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("vsanRemoteDatastoreClusterSpec")
			}
			return err
		}
	}

	return nil
}

func (m *DatastoreSpec) validateVvolDatastoreSpecs(formats strfmt.Registry) error {
	if swag.IsZero(m.VvolDatastoreSpecs) { // not required
		return nil
	}

	for i := 0; i < len(m.VvolDatastoreSpecs); i++ {
		if swag.IsZero(m.VvolDatastoreSpecs[i]) { // not required
			continue
		}

		if m.VvolDatastoreSpecs[i] != nil {
			if err := m.VvolDatastoreSpecs[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("vvolDatastoreSpecs" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("vvolDatastoreSpecs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this datastore spec based on the context it is used
func (m *DatastoreSpec) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateNfsDatastoreSpecs(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVmfsDatastoreSpec(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVSANDatastoreSpec(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVSANRemoteDatastoreClusterSpec(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVvolDatastoreSpecs(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DatastoreSpec) contextValidateNfsDatastoreSpecs(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.NfsDatastoreSpecs); i++ {

		if m.NfsDatastoreSpecs[i] != nil {

			if swag.IsZero(m.NfsDatastoreSpecs[i]) { // not required
				return nil
			}

			if err := m.NfsDatastoreSpecs[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("nfsDatastoreSpecs" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("nfsDatastoreSpecs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *DatastoreSpec) contextValidateVmfsDatastoreSpec(ctx context.Context, formats strfmt.Registry) error {

	if m.VmfsDatastoreSpec != nil {

		if swag.IsZero(m.VmfsDatastoreSpec) { // not required
			return nil
		}

		if err := m.VmfsDatastoreSpec.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vmfsDatastoreSpec")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("vmfsDatastoreSpec")
			}
			return err
		}
	}

	return nil
}

func (m *DatastoreSpec) contextValidateVSANDatastoreSpec(ctx context.Context, formats strfmt.Registry) error {

	if m.VSANDatastoreSpec != nil {

		if swag.IsZero(m.VSANDatastoreSpec) { // not required
			return nil
		}

		if err := m.VSANDatastoreSpec.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vsanDatastoreSpec")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("vsanDatastoreSpec")
			}
			return err
		}
	}

	return nil
}

func (m *DatastoreSpec) contextValidateVSANRemoteDatastoreClusterSpec(ctx context.Context, formats strfmt.Registry) error {

	if m.VSANRemoteDatastoreClusterSpec != nil {

		if swag.IsZero(m.VSANRemoteDatastoreClusterSpec) { // not required
			return nil
		}

		if err := m.VSANRemoteDatastoreClusterSpec.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vsanRemoteDatastoreClusterSpec")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("vsanRemoteDatastoreClusterSpec")
			}
			return err
		}
	}

	return nil
}

func (m *DatastoreSpec) contextValidateVvolDatastoreSpecs(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.VvolDatastoreSpecs); i++ {

		if m.VvolDatastoreSpecs[i] != nil {

			if swag.IsZero(m.VvolDatastoreSpecs[i]) { // not required
				return nil
			}

			if err := m.VvolDatastoreSpecs[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("vvolDatastoreSpecs" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("vvolDatastoreSpecs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *DatastoreSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DatastoreSpec) UnmarshalBinary(b []byte) error {
	var res DatastoreSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
